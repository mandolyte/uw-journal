# On Row Id Generation

## DataTable

in DataTable.js, `generateRowId` function is passed as a parameter with alias `_generateRowId`

Then used:
```js
  const generateRowId = useCallback(_generateRowId, []);
```

and here:
```js
  const columns = useMemo(() => getColumns({
    columnNames, columnsFilter, columnsFilterOptions,
    columnsShow, delimiters, rowHeader,
    generateRowId, cellEdit, preview,
  }), [cellEdit, columnNames, columnsFilter, columnsFilterOptions, columnsShow, delimiters, generateRowId, preview, rowHeader]);
```

## AddRow
In actions-menu / AddRow.js, it is passed as a parameter.

Then used:

```js
  const handleRowAdd = () => {
    rowAddBelow({ rowIndex, rowData: newRow });
    handleClose();
    setTimeout(() => {
      const rowBelow = getRowElement(generateRowId, rowData, 1);

      if (rowBelow) {
        const top = getOffset(rowBelow).top - rowBelow.offsetHeight;
        document.documentElement.scrollTop = top - 20;
        document.body.scrollTop = top - 20;
      }
    }, 200);
  };
```

