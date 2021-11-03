# Issue 498

Original text:
v1.0.5-rc.7  
Open any tN project and note the Rows per Page  
Probably not a big deal for most users, but does break automation scripts

## 2021-02-04

**Info from the `datatable-translatable` component**

The table in question is a "MUIDataTable" from the material-ui components.
```js
import MUIDataTable from 'mui-datatables';
```

This code in `src/components/datatable/DataTable.js` shows that there is an 'options' parameter that has a default for the number of rows.

```js
function DataTable({
  options = {},
  delimiters,
  config,
  onSave,
  onValidate,
  sourceFile,
  generateRowId: _generateRowId,
  ...props
}) {
  const {
    columnsFilter,
    columnsShowDefault,
    rowHeader,
  } = config;
  const dataTableElement = useRef();
  const [rowsPerPage, setRowsPerPage] = useState(options.rowsPerPage || 25);
```

Note that the default if none are specified is 25. Thus this option is being passed here with a non-null value, apparently of 10.

Note that the DataTable() function is not exported. The default and only exported function is this:

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

Found this in TranslatableTSV.js:
```js
  const options = {
    page: 0,
    rowsPerPage: 10,
    rowsPerPageOptions: [10, 25, 50, 100],
  };

  const datatable = useMemo(() => {
    _config.rowHeader = rowHeader;
    return (
      <DataTable
        sourceFile={sourceFile.content}
        targetFile={targetFile.content}
        onSave={onSave}
        onValidate={onValidate}
        delimiters={delimiters}
        config={_config}
        generateRowId={generateRowId}
        options={options}
      />
    );
  }, [sourceFile.content, targetFile.content, onSave, onValidate, generateRowId, options, rowHeader]);
```

Changed to `rowsPerPage` to 25 and tested ... looks good.