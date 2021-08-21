 # tc-create-app#897

## Issue Text
**Issue**:
Currently the app crashes when adding a new record on a deleted row in TSVs.  Users should be prevented from adding a new row from a deleted row.  The new, delete and up down buttons should be disabled until data is added to that row. 

**DoD**:
Users are prevented from attempting to add a new record from a deleted row.

**Details**:
1. Open tN for a book that has deleted rows.(eg. en-tn-LUK)
2. Navigate to the end of the book. This can be done using a filter on chapter and navigating to the last page.
3. click on the + to add a new row. The app crashes.

## Reproduce the error

1. go to https://develop--tc-create-app.netlify.app/
2. org: translate_test 
3. resource: tn
4. book: Jude
5. set rows per page to 100
6. scroll to end

There are deleted rows in this file.

First used last non-deleted row to add a row... that worked. Put this text into the occurrence note: "This is a test.".

Next went to first delete row, which comes immediately after the row I just added. Turned on the console log. Click the plus button. 

Console error:
```
2.d31d0275.chunk.js:1 TypeError: Cannot read property 'map' of undefined
    at t.rowGenerate (2.d31d0275.chunk.js:1)
    at rowGenerate (2.d31d0275.chunk.js:1)
    at 2.d31d0275.chunk.js:1
    at Array.map (<anonymous>)
```

Refers to a function named "rowGenerate"... 
datatable-translatable: core/datatable.js, rowGenerate()
- DataTable.context.js: actions: useMemo() of an object with a bunch of small closures. Ref is:
```js
    rowGenerate: ({ rowIndex }) => rowGenerate({
      rows: targetRows, columnNames, rowIndex,
    })

```

- AddRow.js: this calls rowGenerate()

- Above is called from ActionsMenu.js. It deconstructs the value from DataTable Context to get the functions to run. Among which is `rowAddBelow`.

Branch will be `feature-cn-897-add-row-below-deleted-rows`


Ended up just omitting the row toolbar entirely when the row is a deleted one. How do I determine if a row is deleted? See code graveyard below.

# Testing Notes

Testing Notes:
- login
- org unfoldingWord
- resource tN
- language English
- book Jude (as of this writing Jude has deleted rows); no longer true...
- delete three rows (both 1:2 rows and first 1:3 row)
- set rows per page to 100
- scroll down to end of file
- Notice that for deleted rows (no values on target side), the row action menu is absent.

Test row recovery:
- for a deleted row, enter a value for support reference
- click out
- at this point, the row will be recovered and moved to the end of the target file data, above its deleted rows. The action menu will be restored for the row.

## 2021-07-13

On the review of the PR, Zach noted that the actions menu is also missing from new inserted rows.

New test for deleted rows:
```js
  // If a row is deleted (as opposed to being a new add/insert), then
  // source side will *not* be empty, but the target side will be empty.
  let cellvals = []
  cellvals = rowData[1].split('\t')
  // is this a deleted row?
  if ( cellvals[0] !== "" && cellvals[1] === "" ) {
    return (
      <>
      </>
    )
  }
```

# Yalc script to reset

```sh
#!/bin/sh

BRANCH="feature-cn-897-add-row-below-deleted-rows"
CURDIR=`pwd`
DIRNAME=`dirname $CURDIR`
PROJDIR=`basename $DIRNAME`

if [ "$PROJDIR" != "tc-create-app" ]
then
  echo "Script must be run from ./tc-create-app/scripts"
  exit
fi
cd ..
echo Assumptions:
echo All project folders are at same level
echo All branch names for each project folder are the same 

echo _________________________________
echo Working on datatable-translatable
echo
cd ../datatable-translatable
git switch master
git pull 
git switch $BRANCH
git pull
yarn install
yalc publish


echo ________________________
echo Working on tc-create-app
echo
cd ../tc-create-app
echo First, remove any existing yalc links
yalc remove --all
git switch develop
git pull 

yalc link datatable-translatable
yarn install
yarn start

```

# Code Graveyard

```js
        // rowData is the combined array of source and target
        // if target rows have been deleted, then source will contain more
        // rows that the target. The combined array will have for each
        // cell just the source data ending in a tab character.
        // Normally the target value for each cell follows the tab character.
        // If that target value is missing, then we are on a deleted row 
        // and we cannot add a row at this location.

        // From the console log, here is a sample:
        // 0: "rowHeader"
        // 1: "JUD\t"
        // 2: "1\t"
        // 3: "2\t"
        // 4: "q2qo\t"
        // 5: "figs-metaphor\t"
        // 6: "ὑμῖν…πληθυνθείη\t"
        // 7: "1\t"
        // 8: "May…be multiplied to you\t"
        // 9: "These ideas are spoken of as if ...

        // Note that there is the constant value "rowHeader" in slot zero.
        // Let's test by splitting index 1; in the sample above there should
        // be two elements, both being the string "JUD" if the row is not
        // deleted.
        let cellvals = []
        cellvals = rowData[1].split('\t')
        console.log("cellvals:",cellvals)
        if ( cellvals[0] != cellvals[1] ) {
          return 'Cannot add a row below a deleted row!'
        }
```

