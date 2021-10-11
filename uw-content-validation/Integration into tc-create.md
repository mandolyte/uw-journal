# Integration into tc-create

Note: this document contains a lot of the notes on integrating several resource types into tc-create for content validation.

This resulted in some identified improvements in the API. See Zulip for more into (the "Content Validation package" thread).

**As of today (2010-10-04), further notes will be in CV Diary**

## API Notes
As of 2021-08-26:
-   for TN (7 col): checkNotesTSV7Table with repo code "TN2"
-   for SN: checkNotesTSV7Table with repo code "TN2"
-   for SQ: checkQuestionsTSV7Table with repo code "TQ2"
-   for TQ: checkQuestionsTSV7Table with repo code "TQ2"
-   for TWL: checkTWL_TSV6Table with repo code "TWL"
-   for OBS-TQ: may be in markdown
-   for OBS-TN: may be in markdown
-   for OBS-SN: may be in markdown
-   for OBS-SQ: may be in markdown

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

## 2021-10-01

The refactor is done and works. Branch is `feature-cn-953-add-sn-sq`.
It ended up a bit different than above due to some things that were only available to the Translatable component itself. But very similar.

Next, let's do TWL...

### TranslatableTwlTSV.js

First, it needs to pass in a validate function to datatable:
```js
        onValidate={onValidate}
```

Now it has:
```js
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
        parser={parser}
      />
    );
  }, [sourceFile.content, targetFile.content, onSave, onContentIsDirty, generateRowId, options, rowHeader]);
```

Now add that function and the `_onValidate` function.

Next add the imports.
```js
import * as cv from 'uw-content-validation';
import * as csv from '../../core/csvMaker';
import { contentValidate } from '../../core/contentValidate';
```

in the `_onValidate` function:

1. update the header to: `Reference\tID\tTags\tOrigWords\tOccurrence\tTWLink\n`
2. update the cv function to `checkTWL_TSV6Table`
3. update the resource code to 'TWL'
4. add validationPriority to things from useContext(AppContext)
5. add targetRepository to things from useContext(AppContext)
6. add this line after the useContext for AppContext: ` const langId = targetRepository.language;`
7. add `onValidate` to useMemo dependencies

After doing this and getting it to work, the behavior is weird. I have documented this and provided info to Robert Hunt and will work closely with him to resolve.

See Zulip threads [here](https://unfoldingword.zulipchat.com/#narrow/pm-with/237116-robert.hunt) and [here](https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE--.20UR/topic/Content.20Validation.20package/near/255769857). Some of the content relayed in the posts are in Appendix A.

## 2021-10-02

**Working on Bible Study Notes.** In this case, some of the work was done, but it crashes because it wasn't tied to SN, but to TN 

First, it needs to pass in a validate function to datatable:
```js
        onValidate={onValidate}
```
This was done.

Next add the imports.
```js
import * as cv from 'uw-content-validation';
import * as csv from '../../core/csvMaker';
import { contentValidate } from '../../core/contentValidate';
```
First two were there, only added the last one.

Update the  `_onValidate` function:

1. update the header to: `Reference\tID\tTags\tSupportReference\tQuote\tOccurrence\tNote\n`
2. update the cv function to `checkNotesTSV7Table`
3. update the resource code to 'TN2'
4. add validationPriority to things from useContext(AppContext)
5. add targetRepository to things from useContext(AppContext)
6. add this line after above: ` const langId = targetRepository.language;`
7. add `onValidate` to useMemo dependencies

Seems to work fine!

**Working on Bible Study Questions**
First, it needs to pass in a validate function to datatable:
```js
        onValidate={onValidate}
```
This was done.

Next add the imports.
```js
import * as cv from 'uw-content-validation';
import * as csv from '../../core/csvMaker';
import { contentValidate } from '../../core/contentValidate';
```
First two were there, only added the last one.

Update the  `_onValidate` function:

1. update the header to: `Reference\tID\tTags\tQuote\tOccurrence\tQuestion\tResponse\n`
2. update the cv function to `checkQuestionsTSV7Table`
3. update the resource code to 'TQ2'
4. add validationPriority to things from useContext(AppContext)
5. add targetRepository to things from useContext(AppContext)
6. add this line after above: ` const langId = targetRepository.language;`
7. add `onValidate` to useMemo dependencies


















# Appendix A - TWL
From first link:
@**Robert Hunt** I did some refactoring today in tc-create to prepare us to extend cv support to all the tsv types (and beyond). Right now we only have tN support.

I find that I am passing the appropriate cv function as a parameter to a set of common code that processes the results from the validation. This means I have to pass to this common code, all the parameters that change as well. All that's ok and works fine.

But I find that passing the resource type code is "almost" redundant. So I have a suggestion, namely to remove the type code and instead provide a function that is named per resource type instead. That way you can worry about these codes and I don't have to.

Here is an example...

First, this is what I do now:
```js
      data = await contentValidate(rows, header, cv.checkTN_TSV9Table, langId, bookId, 'TN', validationPriority);
```

If you implemented my suggestion, I'd do this instead:
```js
      data = await contentValidate(rows, header, cv.checkTN_TSV9Table, langId, bookId, validationPriority);
```

For another type, say for the 7 column tN TSV, I would do this:
```js
      data = await contentValidate(rows, header, cv.checkTN_TSV7Table, langId, bookId, validationPriority);
```

These API calls would be simple wrappers for your full set of arguments including the resource type code.

Hope this makes sense. Of course, this is a "breaking change", so it would be a v3.x.x change. Let me know if you want to know more.

**From second link**
@**Robert Hunt** some questions... I tried to add validation for TWL files today. I'm seeing a lot of DEBUG output in the console log. Here are a few lines:
```text
achedFetchFileFromServerWorker could not fetch unfoldingWord en_ugl master content/G14980/01.md: Error: Request failed with status code 404
utilities.js:29 uw-content-validation:   cachedGetFileUsingPartialURL downloaded Door43 unfoldingWord/en_ugl/raw/branch/master/content/G24780/01.md
utilities.js:41 uw-content-validation debug: removeDisabledNotices() cannot work without repoCode for {"priority":138,"message":"File ends with additional blank line(s)","characterIndex":5272,"excerpt":"…v␣19:6](rev␣19:6).\\n\\n","location":""} in list of 1 notices.
utilities.js:29 uw-content-validation:   cachedGetFileUsingPartialURL downloaded Door43 unfoldingWord/en_ugl/raw/branch/master/content/G24790/01.md
utilities.js:29 uw-content-validation:   cachedGetFileUsingPartialURL downloaded Door43 unfoldingWord/en_ugl/raw/branch/master/content/G29000/01.md
utilities.js:29 uw-content-validation:   cachedGetFileUsingPartialURL downloaded Door43 unfoldingWord/en_ugl/raw/branch/master/content/G29040/01.md
utilities.js:41 uw-content-validation debug: removeDisabledNotices() cannot work without repoCode for {"priority":138,"message":"File ends with additional blank line(s)","characterIndex":1929,"excerpt":"…1](luk␣1:51).␣␣\\n\\n\\n\\n\\n","location":""} in list of 1 notices.
utilities.js:29 uw-content-validation:   cachedGetFileUsingPartialURL downloaded Door43 unfoldingWord/en_ugl/raw/branch/master/content/G31670/01.md
utilities.js:41 uw-content-validation debug: removeDisabledNotices() cannot work without repoCode for {"priority":774,"message":"Unexpected ) closing character (no matching opener)","lineNumber":24,"characterIndex":9,"excerpt":"[μέγας]()),","location":""} in list of 2 notices.
utilities.js:29 uw-content-validation:   cachedGetFileUsingPartialURL downloaded Door43 unfoldingWord/en_ugl/raw/branch/master/content/G31730/01.md
utilities.js:41 uw-content-validation debug: removeDisabledNotices() cannot work without repoCode for {"priority":138,"message":"File ends with additional blank line(s)","characterIndex":2507,"excerpt":"…3:13)](1Co␣13:13).\\n\\n","location":""} in list of 1 notices.
utilities.js:29 uw-content-validation:   cachedGetFileUsingPartialURL downloaded Door43 unfoldingWord/en_tw/raw/branch/master/bible/kt/amen.md
utilities.js:29 uw-content-validation:   cachedGetFileUsingPartialURL downloaded Door43 unfoldingWord/en_uhal/raw/branch/master/content/H0543.md
utilities.js:41 uw-content-validation debug: removeDisabledNotices() cannot work without repoCode for {"priority":138,"message":"File ends with additional blank line(s)","characterIndex":585,"excerpt":"…\\n####␣Citations:\\n\\n\\n\\n","location":""} in list of 1 notices.
utilities.js:29 uw-content-validation:   cachedGetFileUsingPartialURL downloaded Door43 unfoldingWord/en_ugl/raw/branch/master/content/G02810/01.md
```
Here is some console output I captured:
[image.png](/user_uploads/15437/Edj8CnHevof9w8fkGmWTr2gp/image.png)
Based on the spreadsheet produced, I think it might be validating the linked TW resources.
[Validation-twl_JUD.tsv-2021-10-01T16_16_45.970Z.xlsx](/user_uploads/15437/KtA6nMOSAD06zESOMiJKW0cV/Validation-twl_JUD.tsv-2021-10-01T16_16_45.970Z.xlsx)
Here is code that I am executing:
```js
  const rawResults = await cvFunction(langId, resourceCode, bookID.toUpperCase(), 'dummy', tableString, '', {suppressNoticeDisablingFlag: false});
```
where:
- `cvFunction` is `checkTWL_TSV6Table`
- langId is 'en'
- resourceCode is 'TWL'
- bookId is "JUD"
- tableString is the content of `twl_JUD.tsv` file.

