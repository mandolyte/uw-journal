# tc-create-app.899

## Issue:

When the user branch is created brand new, adding a new row does not assign a new ID.

## DoD:

Adding a new row assigns a new ID when following the steps below.

## Details:

Steps to replicate:

1.  Delete your existing user branch.
2.  Open tCCreate>unfoldingWord org>tN>en>Luke
3.  Click on the "+" button on any row to add a new row.  
    Note that the modal has no ID displayed.
	If this row is added, the ID field remains blank.
	
# Replication Notes

**Using QA and develop code**

- Pre-conditions:
	- go to https://qa.door43.org/unfoldingWord/en_tn
	- then branches
	- delete my branch (cecil.new-tc-create-1)
- Next, go thru above steps
	- open https://develop--tc-create-app.netlify.app/
	- go to uw/tn/english/luke
	- configure to show ID column
	- go to Luke 1:5 id l008 (should be row 25, so scroll to bottom)
	- click add row
	- ! id is **not** missing?!


OK do the above in production.
- first with ID off... still worked
- second, with ID off, use an earlier row instead of row 25; still worked
- third, closed browser tab (above I just refreshed), with ID off, use early row; still works
- fourth, closed both develop and production app; open prod; with ID off, use early row; still works

Consulted with @elsy since I can't make it fail now.

How about the intro row? using production
- delete my branch
- open https://create.translationcore.com/
- go to Luke
- click "add row" for front:intro
- ID is shown...

## Code Trace

The first clue that an ID is missing is in the "Cancel/Add" dialog.

This comes from datatable-translatable/src/components/actions-menu, in the file AddRow.js.

The content comes from this location around line 80 to 87:
```html
        <DialogContent>
          <Divider />
          <br />
          {newRowComponent}
        </DialogContent>
```

So `newRowComponent` is:
```js
    const newRowComponent = columnNames.map((name, i) => {
      let text = '';

      if (!newRow) {
        const _newRow = rowGenerate({ rowIndex });
        setNewRow(_newRow);
        return text;
      } else {
        text = (
          <DialogContentText key={name + i}>
            <strong>{name}:</strong>
            {' ' + newRow[i]}
          </DialogContentText>
        );
      }
      return text;
    });
```

`newRow` is an array of values and if not yet created, then rowGenerate is called to create it. This function is passed in...

AddRow is called from ActionsMenu.js:
```js
      <AddRow
        rowData={rowData}
        rowIndex={rowIndex}
        columnNames={columnNames}
        rowGenerate={rowGenerate}
        rowAddBelow={rowAddBelow}
        button={addRowButton}
        generateRowId={generateRowId}
      />
```

RowGenerate comes from DataTableContext:
```js
  const { state, actions } = useContext(DataTableContext);

  const {
    rowGenerate,
    rowAddBelow,
    rowDelete,
    rowMoveAbove,
    rowMoveBelow,
  } = actions;
```

`actions` looks like this in the context provider:
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
      console.log("row deleted");
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
    setChanged,
  }), [columnNames, delimiters, targetRows, data]);

```

And `rowGenerate` is imported:
```js
import {
  rowMoveAbove, rowMoveBelow, rowAddBelow, rowDelete, cellEdit,
  parseDataTable, correlateData, rowGenerate, stringify, getColumnsFilterOptions,
} from '../../core/datatable';
```

Finally, the code itself:
```js
export const rowGenerate = ({
  rows, columnNames, rowIndex,
}) => {
  let rowsIndex = {};
  let lengthIndex = {};
  const rowData = rows[rowIndex];

  rows.forEach(_row => {
    _row.forEach((value, index) => {
      const column = columnNames[index];

      if (!rowsIndex[column]) {
        rowsIndex[column] = {};
      }

      if (!rowsIndex[column][value]) {
        rowsIndex[column][value] = 0;
      }
      rowsIndex[column][value]++;
      const valueLength = value.length;

      if (!lengthIndex[column]) {
        lengthIndex[column] = {};
      }

      if (!lengthIndex[column][valueLength]) {
        lengthIndex[column][valueLength] = 0;
      }
      lengthIndex[column][valueLength]++;
    });
  });

  const rowCount = rows.length;
  let newRow = rowData.map((value, index) => {
    const column = columnNames[index];
    const values = Object.keys(rowsIndex[column]).length;
    const valuesRatio = values / rowCount;
    const duplicateValue = (valuesRatio < 0.5);

    const valuesLengths = Object.keys(lengthIndex[column]);
    const valuesLengthsLength = valuesLengths.length;
    const needRandomId = (valuesRatio > 0.99 && valuesLengthsLength <= 2);

    let newValue = '';

    if (duplicateValue) {
      newValue = value;
    } else if (needRandomId) {
      const { length } = value;
      newValue = randomId({ length });
    }
    return newValue;
  });
  return newRow;
};
```