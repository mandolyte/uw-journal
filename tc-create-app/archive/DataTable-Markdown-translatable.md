# Datatable and Markdown Translatable Components

## Transform locations

*In datatable-translatable* [here](https://github.com/unfoldingWord/datatable-translatable/blob/master/src/components/cell/Cell.js)

*In markdown-translatable* [here](https://github.com/unfoldingWord/markdown-translatable/blob/master/src/core/markdown-converter.ts)

## General Notes

1. Every cell in a datatable is treated as Markdown. This was verified by opening a UTN resource, adding all columns to the view, and inserting markup into each cell (did this in markdown mode); then when preview was turned on, all rendered HTML per the markdown.
2. Furthermore, I went to https://datatable-translatable.netlify.app/#/Cell and modified the inputs in the demo to have markup and the display (html) was changed per markup. See [below](DataTable-Markdown-translatable.md#datatable-cell-demo-notes).
3. Given this, it seems redundant to have transforms in both datatable and markdown.
4. The process potentially could involve the following points of transformation:
- DCS to memory (state variable): does it change any data directly from the source before it stores it in memory?
- memory to display: it must change it to HTML so can be altered on the way there
    - Note that with the "preview" button, what is rendered may be just the markdown or the HTML equivalent
    - Both end up on screen as HTML (how could it not be this way), but suspect the markdown may be shown in a `<code/>` block in order to avoid any interpretation.
- display to memory: if user updates the HTML, then it may be altered on the way back to memory
- memory to DCS: if the user saves, then it may be altered on the way back to DCS

## Experiments

### Round trip with no changes to the code

#### 2021-01-27

**Steps**
1. Delete my branch in QA DCS at https://qa.door43.org/unfoldingWord/en_tq/branches
2. Run tc-create and select `TIT_tq.tsv`. The raw DCS value for first annotation is:
```
What was Paul’s purpose in his service to God?\n\n> His purpose was to establish the faith of God’s chosen people and to establish the knowledge of the truth.
```
3. The first annotation reads:
```md
What was Paul’s purpose in his service to God?\n\n> His purpose was to establish the faith of God’s chosen people and to establish the knowledge of the truth.
```
This matches original DCS. 
The raw data and the transformed HTML look the same, *except* the markdown shows between Q and A: `\n\n`; and the HTML shows: `\\n\\n`
4. In HTML mode, make a small change, say the last period to an exclamation mark.
5. Now save it and look at DCS. It is incorrect! Now reads:
```
What was Paul’s purpose in his service to God?\\n\\n> His purpose was to establish the faith of God’s chosen people and to establish the knowledge of the truth!
```
Notice that each original `\n` becomes `\\n`.




## Datatable Cell Demo Notes

The code looks like this:
```js
const value = "Original\tTranslation"
const cellDelimiter = '\t';

const tableMeta = {
  columnIndex: 0,
  rowIndex: 0,
  columnData: {name: 'Column A'},
  rowData: [
    value,
    ['a','a'].join(cellDelimiter),
    ['b','b'].join(cellDelimiter),
    ['c','c'].join(cellDelimiter),
    ['d','d'].join(cellDelimiter),
  ],
};

const onEdit = (object) => alert(JSON.stringify(object));

<Cell
  value={value}
  tableMeta={tableMeta}
  preview
  onEdit={onEdit}
/>
```

Each cell requires two values: one for the source and one for the target. Or "original" and "translation".
In the `<Cell>` component, you can update value to show some of the other rows in the object `tableMeta`:
```js
value={tableMeta.rowData[3]}
```

If you use markdown in the value, then the HTML rendered will be properly per the markdown.
For example:
```js
const value = "# Original\t**Translation**"
```