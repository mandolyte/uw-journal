# Issue 235 Adding Content Validation (on edit)

## 2020-10-31

Here I trace all the copying going on...

- upon file selection, DataTableContext has a useEffect which:
  - parses and loads the data for the targetFile
  - uses setChanged to update the "change" state to false
- once the targetRows are changed, a second useEffect detects this and uses setData
- after rendering the table, if a table cell is edited, then the action cellEdit is called
  - this sends the data to the reducer, passing row number, column number, and value; plus the action 'CELL_EDIT'
    - in the reducer, this action calls a function also named cellEdit (this time from a core module)
    - this cellEdit function is called with rows, rowIndex, columnIndex, value (note: includes the rows)
      - the core cellEdit function does the following:
        - makes a copy of the rows into a varialbe _rows (they are deep frozen which apparently makes them immutable)
        - updates the value with `_rows[rowIndex][columnIndex] = value;`
        - then returns the rows
    - the reducer cellEdit function receives the copied rows, deep freezes them, then returns them
    - the reducer then updates targetRows with the updated copied rows
  - the then actions cellEdit function uses setChanged to update the "change" state to true
- a useEffect again detects the change and uses setData to udpate the data state variable

For every keypress in a cell (within a time window), the following takes place.

- the targetRows are copied
- the cell is updated in copy
- the data is deep frozen
- the targetRows are updated 
- the state data is updated


## 2020-10-30

In DataTable, the cell edit function is provided by DataTableContext:

```js
  const { state, actions } = useContext(DataTableContext);
  const {
    columnNames, data, changed, columnsFilterOptions,
  } = state;
  const { cellEdit:_cellEdit } = actions;
```

In DataTableContext, there are these lines:
```js
  const [targetRows, targetRowsDispatch] = useReducer(rowsReducer, {});
  const setTargetRows = (rows) => targetRowsDispatch({ type: 'SET_ROWS', value: { rows } });
```
These lines apply a reducer to process the actions for the target rows. The second line defines a function to set the rows from the target data.

Here is the reducer:
```js
const rowsReducer = (rows, action) => {
  let _rows;
  const { type, value } = action;
  const {
    rowIndex, rowData, columnIndex,
  } = value;

  switch (type) {
  case 'SET_ROWS':
    return deepFreeze(value.rows);
  case 'ROW_MOVE_ABOVE':
    _rows = rowMoveAbove({ rows, rowIndex });
    return deepFreeze(_rows);
  case 'ROW_MOVE_BELOW':
    _rows = rowMoveBelow({ rows, rowIndex });
    return deepFreeze(_rows);
  case 'ROW_ADD_BELOW':
    _rows = rowAddBelow({
      rows, rowIndex, rowData,
    });
    return deepFreeze(_rows);
  case 'ROW_DELETE':
    _rows = rowDelete({ rows, rowIndex });
    return deepFreeze(_rows);
  case 'CELL_EDIT':
    _rows = cellEdit({
      rows, rowIndex, columnIndex, value: value.value,
    });
    return deepFreeze(_rows);
  default:
    throw new Error(`Unsupported action type: ${action.type}`);
  };
};

```

Here is the Context value:
```js
  const value = useMemo(() => deepFreeze({
    state:{
      columnNames,
      data,
      changed,
      columnsFilterOptions,
    },
    actions,
  }), [actions, changed, columnNames, columnsFilterOptions, data]);
```
which is an object with two properties: state and actions. State is shown (has the column names, the data, etc.). The actions is defined above as:

```js
  const actions = useMemo(() => ({
    rowMoveAbove: ({ rowIndex }) => {
      targetRowsDispatch({ type: 'ROW_MOVE_ABOVE', value: { rowIndex } });
      setChanged(true);
    },
    rowMoveBelow: ({ rowIndex }) => {
      targetRowsDispatch({ type: 'ROW_MOVE_BELOW', value: { rowIndex } });
      setChanged(true);
    },
    rowAddBelow: ({ rowIndex, rowData }) => {
      targetRowsDispatch({ type: 'ROW_ADD_BELOW', value: { rowIndex, rowData } });
      setChanged(true);
    },
    rowDelete: ({ rowIndex }) => {
      targetRowsDispatch({ type: 'ROW_DELETE', value: { rowIndex } });
      setChanged(true);
    },
    cellEdit: ({
      rowIndex, columnIndex, value,
    }) => {
      targetRowsDispatch({
        type: 'CELL_EDIT', value: {
          rowIndex, columnIndex, value,
        },
      });
      setChanged(true);
    },
    rowGenerate: ({ rowIndex }) => rowGenerate({
      rows: targetRows, columnNames, rowIndex,
    }),
    targetFileSave: () => stringify({
      columnNames, rows: targetRows, delimiters,
    }),
  }), [columnNames, delimiters, targetRows]);
```
Which is an memoized arrow function composed of a single object containing a lot of closures, one for each action that can be taken on the rows.

I'm interested in `cellEdit`:
```js
    cellEdit: ({
      rowIndex, columnIndex, value,
    }) => {
      targetRowsDispatch({
        type: 'CELL_EDIT', value: {
          rowIndex, columnIndex, value,
        },
      });
      setChanged(true);
    },
```
This takes three parameters: row, column, and value. It then invokes the target dispatcher. That will run the "CELL_EDIT" part of the reducer, which does this:
```js
  case 'CELL_EDIT':
    _rows = cellEdit({
      rows, rowIndex, columnIndex, value: value.value,
    });
    return deepFreeze(_rows);
```
This case calls a function `cellEdit()` which is imported from:
```js
import {
  rowMoveAbove, rowMoveBelow, rowAddBelow, rowDelete, cellEdit,
  parseDataTable, correlateData, rowGenerate, stringify, getColumnsFilterOptions,
} from '../../core/datatable';
```
From there:
```js
export const cellEdit = ({
  rows, rowIndex, columnIndex, value,
}) => {
  let _rows = rows.map(cells => [...cells]);
  _rows[rowIndex][columnIndex] = value;
  return _rows;
};
```
This does the following:
- makes a copy of the rows
- sets the value for the give row and colum index
- returns the update copy of rows

Upon return, the updated copy of the rows is "deep frozen", then the changed state variable is to true with `setChanged(true);`.

A useEffect detects when the rows have changed:
```js
  // correlate data by compositeKeyIndices when sourceRows or targetRows updated
  useDeepEffect(() => {
    if (Object.keys(sourceRows).length && Object.keys(targetRows).length) {
      const _data = correlateData({
        sourceRows, targetRows, compositeKeyIndices, delimiters,
      });
      setData(_data);
    }
  }, [sourceRows, targetRows, compositeKeyIndices, delimiters]);
```

It sets the state "data".

The state data is part of the context value (see above). Now when DataTable is used, there is a head-fake and what is really exported and then imported by the app is a wrapping function disguised as DataTable:
```js
export default function DataTableWrapper(props) {
  return (
    <DataTableContextProvider {...props}>
      <DataTable {...props} />
    </DataTableContextProvider>
  );
}
```
This wrapping function returns a the DataTableContext with a DataTable inside it.


Here is the current on-demand code in TranslatableTSV:
```js
  const onValidate = useCallback(async () => {
    // sample name: en_tn_08-RUT.tsv
    if ( targetFile ) {
      const _name  = targetFile.name.split('_');
      const langId = _name[0];
      const bookID = _name[2].split('-')[1].split('.')[0];
      const content = targetFile.content;
      const rawResults = await cv.checkTN_TSVText(langId, bookID, 'dummy', content, '');
      const nl = rawResults.noticeList;
      let hdrs =  ['Priority','Chapter','Verse','Line','Row ID','Details','Char Pos','Excerpt','Message','Location'];
      let data = [];
      data.push(hdrs);
      Object.keys(nl).forEach ( key => {
        const rowData = nl[key];
        csv.addRow( data, [
            String(rowData.priority),
            String(rowData.C),
            String(rowData.V),
            String(rowData.lineNumber),
            String(rowData.rowID),
            String(rowData.details),
            String(rowData.characterIndex),
            String(rowData.extract),
            String(rowData.message),
            String(rowData.location),
        ]);
      });

      let ts = new Date().toISOString();
      let fn = 'Validation-' + targetFile.name + '-' + ts + '.csv';
      csv.download(fn, csv.toCSV(data));
  
      console.log("validations:",rawResults);
    }
  },[targetFile]);
```


# Issue 235 Adding Content Validation (on demand)

## 2020-10-24

Borrowed from book-package-app:
```js
function download(filename: string, text: string) {
    let element = document.createElement('a');
    element.setAttribute('href', 'data:text/csv;charset=utf-8,' + encodeURIComponent(text));
    element.setAttribute('download', filename);

    element.style.display = 'none';
    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);
}
```

## 2020-10-22 

Today trying to figure out how to add a check icon button. It needs to be on the datatable translatable toolbar.

- The RCL toolbar component only has save and preview (right to left).
- The DataTableWrapper component has more:
    - filter
    - column selection
    - search

I think I need to add one between search and column selection.

This work had to be done in DataTable RCL.
