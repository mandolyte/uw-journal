# Implementing TSV Parser
Branch name in tc-create is "feature-cn-788-impl-tsv-parser".

## 2021-07-28

Yesterday, implemented the new parser in tc-create's onOpenValidation code.

Today, working on DataTableContextProvider (DataTable.context.js).

Here is current parameters:
```js
export function DataTableContextProvider({
  children,
  sourceFile,
  targetFile,
  delimiters,
  config: {
    compositeKeyIndices,
    columnsFilter,
  },
}) {
```

Instead of passing delimiters, let's pass in the entire parser of choice.

This is called indirectly in tc-create. In the "translatable" TSV components in tc-create, we find code like this:
```js
    return (
      <DataTable
        sourceFile={sourceFile.content}
        targetFile={targetFile.content}
        onSave={onSave}
        onContentIsDirty={onContentIsDirty}
        delimiters={delimiters}
        config={_config}
        generateRowId={generateRowId}
        options={options}
      />
    );
```
which is imported like this:
`import { DataTable } from 'datatable-translatable';`

However, the default function exported is not "DataTable", but:
```js
export default function DataTableWrapper(props) {
  return (
    <MarkdownContextProvider>
      <DataTableContextProvider {...props}>
        <DataTable {...props} />
      </DataTableContextProvider>
    </MarkdownContextProvider>
  );
}
```

So 

## 2021-07-27
Places (with code snippets) that need to be considered to implement the new TSV parser are all below in section 2021-07-26.

Two packages impacted:
- tc-create-app (for on open validation of TSV files)
- datatable-translatable (for parsing of source and target TSV files)

Here is the list of functions needed in the parsing package:

First from datatable.js:
- parse a TSV string to a 2D array (`parseDataTable` )
- convert 2D array back to a string (`stringify` )
- access the header row (`getColumnNames`)
- access the rows (`getRows`)

In DataTable.context is where these are used:
- There are two `useEffect()` hooks which control when the above functions are used in  `DataTable.context.js` (actually just `parseDataTable`).
- The `stringify` is used one place in a "reducer" item named "targetFileSave".

## 2021-07-26
At this point, I have a local (not published to NPM) TSV package. Let's try to use it.

### Translatable
First, to Translatable.js... this has the following lines that are processed for TSVs:

```js
      } else if (sourceFile.filepath.match(/^tq_...\.tsv$/)) {
        _translatable = <TranslatableTqTSV onSave={saveOnTranslation} onContentIsDirty={setContentIsDirty} />;
      } else if (sourceFile.filepath.match(/^twl_...\.tsv$/)) {
        _translatable = <TranslatableTwlTSV onSave={saveOnTranslation} onContentIsDirty={setContentIsDirty} />;
      } else if (sourceFile.filepath.match(/\.tsv$/)) {
        _translatable = <TranslatableTSV onSave={saveOnTranslation} onContentIsDirty={setContentIsDirty} />;
```

The component is save into the variable `_translatable` which by default has:
```js
    let _translatable = (
      <div style={{ textAlign: 'center' }}>
        <CircularProgress />{' '}
      </div>
    );

```

In turn, this is part of a useMemo() stored in `translatableComponent`. This handles both TSV and Markdown and it includes an "save" function.

This component is combined with a files header component into this:
```js
    <div className={classes.root}>
      {filesHeader}
      <div id='translatableComponent'>
        {translatableComponent}
      </div>
      {authenticationModal}
    </div>
```

### TranslatableTwlTSV
Let's the twl route for now using `TranslatableTwlTSV`. In this code, we see this:
```js
  const datatable = useMemo(() => {
    _config.rowHeader = rowHeader;
    return (
      <DataTable
        sourceFile={sourceFile.content}
        targetFile={targetFile.content}
        onSave={onSave}
        onContentIsDirty={onContentIsDirty}
        delimiters={delimiters}
        config={_config}
        generateRowId={generateRowId}
        options={options}
      />
    );
  }, [sourceFile.content, targetFile.content, onSave, onContentIsDirty, generateRowId, options, rowHeader]);
```

Here the source and target content are passed to DataTable. This implies that no TSV parsing is happening in tc-create.

### TranslatableTSV

Here is the datatable component for TNs (9 col):
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

### TranslatableTq

Here is DataTable:
```js
  const datatable = useMemo(() => {
    _config.rowHeader = rowHeader;
    return (
      <DataTable
        sourceFile={sourceFile.content}
        targetFile={targetFile.content}
        onSave={onSave}
        onContentIsDirty={onContentIsDirty}
        delimiters={delimiters}
        config={_config}
        generateRowId={generateRowId}
        options={options}
      />
    );
  }, [sourceFile.content, targetFile.content, onSave, onContentIsDirty, generateRowId, options, rowHeader]);
```

### TargetFile.context
The useFile hook has an "onOpenValidation" function that is part of the input. This is passed as a parameter to this context provider:

```js
function TargetFileContextProvider({
  onOpenValidation, 
  children
}) {
```

Where is TargetFileContext used? It is used in Workspace.js. 

Here, it:
- passes in the onOpenValidation function
- and then renders, either the results of the validation or, if no errors found, will render the Translatable component.

```js
          <TargetFileContextProvider 
            onOpenValidation={_onOpenValidation}
          >
            {
              (criticalErrors.length > 0 && 
                <Dialog
                  disableBackdropClick
                  open={(criticalErrors.length > 0)}
                  onClose={handleClose}
                  aria-labelledby="alert-dialog-title"
                  aria-describedby="alert-dialog-description"
                >
                  <DialogTitle id="alert-dialog-title">
                  This file cannot be opened by tC Create as there are errors in the target file. 
                  Please contact your administrator to address the following error(s)                  
                  </DialogTitle>
                  <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                    {
                      criticalErrors.map( (msg,idx) => {
                        return (
                          <>
                          <Typography key={idx}>
                            On <Link href={msg[0]} target="_blank" rel="noopener">
                              line {msg[1]}
                            </Link>
                            &nbsp;{msg[2]}&nbsp;{msg[3]}&nbsp;{msg[4]}&nbsp;{msg[5]}
                          </Typography>
                          </>
                        )
                    })}
                    </DialogContentText>
                  </DialogContent>
                  <DialogActions>
                    <Button onClick={handleClose} color="primary">
                      Close
                    </Button>
                  </DialogActions>
                </Dialog>
              ) 
              ||
              <Translatable />
            }
          </TargetFileContextProvider>
```


This function takes in the filename, the content, and the URL of the file, runs the validation function and manages the state:
```js
  const _onOpenValidation = (filename,content,url) => {
    const notices = onOpenValidation(filename, content, url);
    if (notices.length > 0) {
      setCriticalErrors(notices);
    } else {
      setCriticalErrors([]);
    }
    return notices;
  };
```

The "real" onOpenValidation is in the core folder and is imported into Workspace:
```js
import { onOpenValidation } from './core/onOpenValidations';
```

This function does some setup and then runs more specific functions:
```js
export const onOpenValidation = (filename, content, url) => {
  const link = url.replace('/src/', '/blame/');
  let criticalNotices = [];

  if ( filename.match(/^tn_...\.tsv$/) ) {
    criticalNotices = onOpenValidationTn7(content,link);
  } else if ( filename.match(/tn_..-...\.tsv$/) ) {
    criticalNotices = onOpenValidationTn9(content, link);
  } else if ( filename.match(/^twl_...\.tsv$/) ) {
    criticalNotices = onOpenValidationTwl(content, link);
  }
  return criticalNotices;
}
```

There is a generic TSV parse phase that it is configurable that each of the specific functions above run with the appropriate arguments.

```js
 const rows = content.replaceAll('\r','').split('\n');
```

After some checks that can be done at the row level, then cells are obtained via
```js
      let cols = rows[i].split('\t');
```

Appears that the tsv parser package can be used in this routine.

### DataTable

In the core folder is datatable.js with lots of common routines. These are the ones to review:

```js
export const parseDataTable = ({ table, delimiters }) => {
  const rows = parseRows({ table, delimiter: delimiters.row })
    .map(row =>
      parseCells({ row, delimiter: delimiters.cell }),
    );
  const dataTable = {
    columnNames: getColumnNames(rows),
    rows: getRows(rows),
  };
  return dataTable;
};

export const stringify = ({
  columnNames, rows, delimiters,
}) => {
  let string = '';

  if (columnNames && rows) {
    let dataTable = [columnNames, ...rows];

    string = dataTable.map(cells => cells.join(delimiters.cell))
      .join(delimiters.row);
  }
  return string;
};

export const getColumnNames = (rows) => rows[0];
export const getRows = (rows) => rows.slice(1);

export const parseRows = ({ table, delimiter }) => table.split(delimiter).filter(row => row !== '');
export const parseCells = ({ row, delimiter }) => row.split(delimiter);
```

There is also a test case in `__test__/bidi-tsv.js`:
```js
import { parseDataTable, stringify } from "../src/core/datatable";
import fs from 'fs';
import path from 'path';

describe('Bidirectional TSV Tests', () => {

  it('should convert en_tn_57-TIT tsv string to JSON and back', () => {
    generateTest('en_tn_57-TIT');
  });

  it('should convert tn_57-TIT tsv string to JSON and back', () => {
    generateTest('tn_57-TIT');
  });
})

function generateTest(fileName) {
  const delimiters = { row: '\n', cell: '\t' };
  const tsv = fs.readFileSync(path.join(__dirname, './fixtures', `${fileName}.tsv`), { encoding: 'utf-8' });
  const parsedTable = parseDataTable({ table: tsv, delimiters });
  expect(parsedTable).toMatchSnapshot();
  const { columnNames, rows } = parsedTable;
  expect(stringify({ columnNames, rows, delimiters })).toBe(tsv);
}
```

In DataTable.context.js are two useEffect()s:
```js
  // parse sourceFile when updated
  useEffect(() => {
    if (delimiters) {
      const { rows } = parseDataTable({ table: sourceFile, delimiters });
      setSourceRows(rows);
    }
  }, [sourceFile, delimiters]);
  // parse targetFile when updated
  useEffect(() => {
    if (delimiters) {
      const { columnNames, rows } = parseDataTable({ table: targetFile, delimiters });
      setColumnNames(columnNames);
      setTargetRows(rows);
      setChanged(false);
    }
  }, [targetFile, delimiters]);

```

In the actions "useMemo()" is:
```js
    targetFileSave: () => stringify({
      columnNames, rows: targetRows, delimiters,
    }),
```
