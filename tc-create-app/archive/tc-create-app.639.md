# tc-create-app.639

## 2021-03-25

At this point, I have solved the problem by relocating code so that the data state variable is visible to the code it needs to be visible to.

One nagging thing left to do. A value on the screen is being left behind even when it isn't in the underlying data. @klappy says it could be due to memoization, but most likely due to rendering not being updated.

If I change row-per-page to force re-rendering then it goes away.

**here is the case of where I deleted 3 rows**
![[Pasted image 20210325110833.png]]

**here is after I edit the last deleted row to add a support reference**
![[Pasted image 20210325111004.png]]

Notice that the "add-a-spt-ref" text is now where it should be in the edited row that jumped above (row id b4zh). But it also appears on the last row where it was before it jumped.

If you remove the text, which a user will be tempted to do, then that is treated as an edit and the row will jump and become a new row. *If you now delete that row, then all be well.*

Found some links that might be useful in fixing this:
https://medium.com/@dev.cprice/wild-react-useforceupdate-e4459f2c1272
https://lodash.com/docs#isEqual (this is used in our code connected with the mui-datatable - see around line 30 of `DataTable.js`.

React has a `forceUpdate()`: 
https://reactjs.org/docs/react-component.html#forceupdate

Also looked at the API for mui-datatables:
https://www.npmjs.com/package/mui-datatables#api

So far, nothing promising.


## 2021-03-23

Issues are about how to pass the in-memory data in a "live" way. Closures will pick the value of data at the time of invocation, but not the *current* value.

## 2021-03-17

Puzzlement... the in-memory representation does not indicate deleted rows.

```js
export const cellEdit = ({
  rows, rowIndex, columnIndex, value, data,
}) => {
  let _rows = rows.map(cells => [...cells]);
  // if row index points beyond end of array, 
  // then add as many empty rows as needed to 
  // make it a valid, even if empty row
  if ( rowIndex >= rows.length ) {
    console.log("Undo delete process begins")
    console.log("[datatable.js] cellEdit() number of row=", rows.length, " rowIndex=", rowIndex);
    for (let i=-1; i < (rowIndex - rows.length); i++) {
      let _row = new Array(rows[0].length);
      // don't really want to do the below
      for (let j=0; j < _row.length; j++) {
        _row[j] = "---";
      }
      _rows.push( _row );
    }
    // now do an "undo" by filling in values from source
    console.log("cellEdit() data=",data);
    
    for (let i=0; i < _rows[rowIndex].length; i++) {
      console.log("copying: from, to", rowIndex, i, " values:", data[rowIndex][i], _rows[rowIndex][i])
      //_rows[rowIndex][i] = sourceRows[rowIndex][i];
    }
    console.log("Undo delete process ends")
  }
  
  _rows[rowIndex][columnIndex] = value;
  return _rows;
};

```

## 2021-03-16
Per discussion today, it did not crash like it does now when an attempt is made to edit a deleted row. Also noteworthy: *a new row added to the source after target branch file is created will look like a deleted row. It will not have an id that corresponds to any in the target... thus looks the same as a deleted row.*

### Reproduce the issue locally
After running the set up below, click in a target cell on the deleted row, say occurrence note. Enter some data, then lose focus.

Crash happens. Here is the console with complete details/trace:
```
datatable.js:160 Uncaught TypeError: Cannot set property '8' of undefined
    at cellEdit (datatable.js:160)
    at rowsReducer (DataTable.context.js:198)
    at updateReducer (react-dom.development.js:15120)
    at Object.useReducer (react-dom.development.js:15904)
    at useReducer (react.development.js:1467)
    at DataTableContextProvider (DataTable.context.js:227)
```

Then more details:
```
The above error occurred in the <DataTableContextProvider> component:
    in DataTableContextProvider (created by DataTableWrapper)
    in MarkdownContextProvider (created by DataTableWrapper)
    in DataTableWrapper (at TranslatableTSV.js:171)
    in TranslatableTSV (at TranslatableTSV.js:194)
    in ResourcesContextProvider (at TranslatableTSV.js:185)
    in TranslatableTSVWrapper (at Translatable.js:139)
    in MarkdownContextProvider (at Translatable.js:164)
    in div (at Translatable.js:162)
    in Translatable (at Workspace.js:38)
    in TargetFileContextProvider (at Workspace.js:33)
    in Workspace (at App.js:94)
    in div (at App.js:93)
    in FileContextProvider (at App.js:77)
    in RepositoryContextProvider (at App.js:71)
    in OrganizationContextProvider (at App.js:66)
    in AuthenticationContextProvider (at App.js:60)
    in ThemeProvider (at App.js:59)
    in div (at App.js:58)
    in AppComponent (at App.js:140)
    in AppContextProvider (at App.js:139)
    in App (at src/index.js:8)
```

The message about cannot set property '8' ... the 8 apparently is the index into the row. So if I try to edit the ID column it will complain about property '3'.


### Project setup: first time steps only
First, get master for tc-create, run it and reproduce the problem locally.
- git switch master
- git pull
- yalc remove --all
- yarn install
- yarn start

Login:
- as myself
- go to 3JN
- set rows to 50 (there's only 47 notes)
- delete the first note on 3JN 1:1, rowId=rni7
- after it resorts and updates the display, scroll to bottom to find the deleted row
- Observe that the target values are empy now
