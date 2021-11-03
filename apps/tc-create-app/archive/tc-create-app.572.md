# tc-create-app.572

This issue will be fixed `datatable-translatable`

This screen shot captures the area that Zach and I determined to be the correct place. We need a place that will update the data in-memory.

![[Screenshot 2021-02-05 071850.png]]

Here is the code above:

```js
    data = Object.values(rowIndex).map(row => {
      let _row;

      if (row.source) {
        _row = row.source.map((sourceCell, index) =>
          `${sourceCell}${delimiters.cell}${row.target ? row.target[index] : ''}`,
        );
      } else {
        _row = row.target.map((targetCell, index) =>
          `${delimiters.cell}${targetCell}`,
        );
      }
      return _row;
    });
  }

```

This code is in a function named `correlateData` and it returns the data array that combines both source and target into a single array of rows. Each cell has both source (if any) and target (if any) values delimited by a tab (`delimiters.cell`). The correlation is done by row id prior to this.

If the row id exists in the source, then the value from the source is added, followed by the tab character, followed by the value from the target or null if empty.

If the row id does not exist in the source, then the cell is created by adding the tab first (omitting non-existent source value), followed by the tab character.

So here we can add filtering to remove or update the in-memory state for the TSV used by the datatable. 

The definition of done is to remove any zero width spaces from the beginning and the end of the target value.

There are two places.

1. From: 
	- `row.target[index]` 
	- to `row.target[index].replace(/^\u200B+/, '').replace(/\u200B+$/,'')`
1. From:
	- `${targetCell}`
	- to `${targetCell.replace(/^\u200B+/, '').replace(/\u200B+$/,'')`