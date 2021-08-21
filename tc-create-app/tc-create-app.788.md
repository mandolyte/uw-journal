# tc-create-app.788

Incorporating the new TSV parser package.

- ~~tN (9 col) worked (both editor and translator modes checked)~~
- ~~twl works (both checked)~~
**NOTE!**: must check that round trips work... so retest!

Need to test tN 7col format.


## 2021-08-12

Working on the save side of things. 

In `DataTable.context.js`, in the reducer (~line 152) is:
```js
    targetFileSave: () => stringify({
      columnNames, rows: targetRows, delimiters,
    }),
```

The `stringify` function converts to a string; is located in `src/core/datatable.js` which is largely helper functions. The code is:
```js
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

```

This needs to use any provider that is supplied. This is supplied via a parameter named "parser" to the context:
```js
export function DataTableContextProvider({
  children,
  sourceFile,
  targetFile,
  delimiters,
  parser,
  config: {
    compositeKeyIndices,
    columnsFilter,
  },
```



## Steps for tN 7col format

1. login to qa and go to: https://qa.door43.org/?repo-search-tab=organizations
2. create a new org: test_tn_7col_format: https://qa.door43.org/org/test_tn_7col_format/dashboard
3. fork unfoldingword/en_tn to it
![[Pasted image 20210729090238.png]]
4. Verify newFormat branch is present: https://qa.door43.org/test_tn_7col_format/en_tn/src/branch/newFormat
5. Now, merge newFormat into master
	- git clone git@qa.door43.org:test_tn_7col_format/en_tn.git
	- cd en_tn
	- git config merge.renameLimit 999999
	- git merge origin/newFormat
	- This resulted in both formats being in the master branch
	- Remove all the old format: `rm en_tn_*.tsv`
	- Now commit: git commit -a -m "merged newformat branch"
	- Now push: git push
7. Now my org has en_tn with 7col tN data; verify in browser

Next, update (temporarily only!!) `core/state.defaults.js` to use my org for tN data:

```js
      SERVER_URL + '/api/v1/repos/unfoldingWord/en_ta',
      SERVER_URL + '/api/v1/repos/unfoldingWord/en_tw',
      SERVER_URL + '/api/v1/repos/unfoldingWord/en_twl',
      SERVER_URL + '/api/v1/repos/test_tn_7col_format/en_tn', // <-- here
      SERVER_URL + '/api/v1/repos/unfoldingWord/en_tq',
```

DO NOT COMMIT THIS CHANGE!!!

Now login with local tc-create:
- select my new org
- ![[Pasted image 20210729094430.png]]
- select resource (now from my org!)
- ![[Pasted image 20210729094336.png]]
- select language and file, then it hangs. This is in the console:
![[Pasted image 20210729111052.png]]

To fix:
- added new TranslatableTnTSV and RowHeaderTn
- added pattern match in Translatable.js

```js
// Used to fake 7col unfoldingWord tN resource:
// SERVER_URL + '/api/v1/repos/test_tn_7col_format/en_tn',
```
