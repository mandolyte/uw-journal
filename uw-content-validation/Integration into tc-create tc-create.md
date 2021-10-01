# Integration into tc-create

## Current Status

At this writing (2021-09-29), only CV of Translation Notes is in place.

Here are the details.

Imports:
```js
import * as cv from 'uw-content-validation';
import * as csv from '../../core/csvMaker';
```

Validation Priority, which is set in the "drawer" (click hamburger menu button), is needed from App Context:
```js
  const {
    state: { resourceLinks, expandedScripture, validationPriority },
    actions: { setResourceLinks },
  } = useContext(AppContext);
```

Then there is a closure (callback) used to execute the CV code. It is long...

Below are two needed functions. The smaller one at the end defines an "onValidate" closure. This is a wrapper for the "real" `onValidate` closure that is immediately below and which does all the work. 

*Idea: refactor into a common utility function and pass as parameters the things need per resource type.* See below the code snippets for the parameters.

```js
  const _onValidate = useCallback(async (rows) => {
    // sample name: en_tn_08-RUT.tsv
    // NOTE! the content on-screen, in-memory does NOT include
    // the headers. This must be added
    const header = "Book\tChapter\tVerse\tID\tSupportReference\tOrigQuote\tOccurrence\tGLQuote\tOccurrenceNote\n";
    if ( targetFile && rows ) {
      // first - create a string from the rows 2D array (table)
      let tableString = header;
      for (let i=0; i < rows.length; i++) {
        for (let j=0; j < rows[i].length; j++) {
          tableString += rows[i][j];
          if ( j < (rows[i].length - 1) ) {
            tableString += delimiters.cell;
          }
        }
        tableString += delimiters.row;
      }

      // second collect parameters needed by cv package
      const _name  = targetFile.name.split('_');
      const langId = _name[0];
      const bookID = _name[2].split('-')[1].split('.')[0];
      const rawResults = await cv.checkTN_TSV9Table(langId, 'TN', bookID, 'dummy', tableString, '', {suppressNoticeDisablingFlag: false});
      const nl = rawResults.noticeList;
      let hdrs =  ['Priority','Chapter','Verse','Line','Row ID','Details','Char Pos','Excerpt','Message','Location'];
      let data = [];
      data.push(hdrs);
      let inPriorityRange = false;
      Object.keys(nl).forEach ( key => {
        inPriorityRange = false; // reset for each
        const rowData = nl[key];
        if ( validationPriority === 'med' && rowData.priority > 599 ) {
          inPriorityRange = true;
        } else if ( validationPriority === 'high' && rowData.priority > 799 ) {
          inPriorityRange = true;
        } else if ( validationPriority === 'low' ) {
          inPriorityRange = true;
        }
        if ( inPriorityRange ) {
          csv.addRow( data, [
              String(rowData.priority),
              String(rowData.C),
              String(rowData.V),
              String(rowData.lineNumber),
              String(rowData.rowID),
              String(rowData.fieldName || ""),
              String(rowData.characterIndex || ""),
              String(rowData.extract || ""),
              String(rowData.message),
              String(rowData.location),
          ])
        }
      });

      if ( data.length < 2 ) {
        alert("No Validation Errors Found");
        setOpen(false);
        return;
      }

      let ts = new Date().toISOString();
      let fn = 'Validation-' + targetFile.name + '-' + ts + '.csv';
      csv.download(fn, csv.toCSV(data));
  
      //setOpen(false);
    }
    setOpen(false);
  },[targetFile, validationPriority]);

  const onValidate = useCallback( (rows) => {
    setOpen(true);
    setTimeout( () => _onValidate(rows), 1);
  }, [_onValidate]);

```

The datatable setup:
```js
  const datatable = useMemo(() => {
    _config.rowHeader = rowHeader;
    return (
      <DataTable
        sourceFile={sourceFile.content}
        targetFile={targetFile.content}
        onSave={onSave}
        onValidate={onValidate}
        onContentIsDirty={onContentIsDirty}
        delimiters={delimiters}
        config={_config}
        generateRowId={generateRowId}
        options={options}
      />
    );
  }, [sourceFile.content, targetFile.content, onSave, onValidate, onContentIsDirty, generateRowId, options, rowHeader]);
```

1. Since the headers are not stored in memory, they must be added. So this is one parameter that must be passed in.
2. The name of the CV function varies, so it must be passed in.

Note: hooks can only be used in React components. Thus the essential structure of the above must stay intact. This means I'll need to factor only the logic in `_onValidate`.

Here's what it might look like:
```js
  const _onValidate = useCallback(async (rows) => {
  // header must include the newline character at the end!
    const header = "Book\tChapter\tVerse\tID\tSupportReference\tOrigQuote\tOccurrence\tGLQuote\tOccurrenceNote\n";
	const cvFunction = cv.checkTN_TSV9Table;
    if ( targetFile && rows ) {
		contentValidation(rows, header, cvFunction)
    }
    setOpen(false);
  },[targetFile, validationPriority]);
```

Then in `core/contentValidate.js`, we put the guts... so it would be like this:
```js
export const contentValidate = (rows, header, cvFunction) => {
  // first - create a string from the rows 2D array (table)
  let tableString = header;
  for (let i=0; i < rows.length; i++) {
	for (let j=0; j < rows[i].length; j++) {
	  tableString += rows[i][j];
	  if ( j < (rows[i].length - 1) ) {
		tableString += '\t';
	  }
	}
	tableString += '\n';
  }

  // second collect parameters needed by cv package
      const _name  = targetFile.name.split('_');
      const langId = _name[0];
      const bookID = _name[2].split('-')[1].split('.')[0];
      const rawResults = await cv.checkTN_TSV9Table(langId, 'TN', bookID, 'dummy', tableString, '', {suppressNoticeDisablingFlag: false});
      const nl = rawResults.noticeList;
      let hdrs =  ['Priority','Chapter','Verse','Line','Row ID','Details','Char Pos','Excerpt','Message','Location'];
      let data = [];
      data.push(hdrs);
      let inPriorityRange = false;
      Object.keys(nl).forEach ( key => {
        inPriorityRange = false; // reset for each
        const rowData = nl[key];
        if ( validationPriority === 'med' && rowData.priority > 599 ) {
          inPriorityRange = true;
        } else if ( validationPriority === 'high' && rowData.priority > 799 ) {
          inPriorityRange = true;
        } else if ( validationPriority === 'low' ) {
          inPriorityRange = true;
        }
        if ( inPriorityRange ) {
          csv.addRow( data, [
              String(rowData.priority),
              String(rowData.C),
              String(rowData.V),
              String(rowData.lineNumber),
              String(rowData.rowID),
              String(rowData.fieldName || ""),
              String(rowData.characterIndex || ""),
              String(rowData.extract || ""),
              String(rowData.message),
              String(rowData.location),
          ])
        }
      });

      if ( data.length < 2 ) {
        alert("No Validation Errors Found");
        setOpen(false);
        return;
      }

      let ts = new Date().toISOString();
      let fn = 'Validation-' + targetFile.name + '-' + ts + '.csv';
      csv.download(fn, csv.toCSV(data));
}

```