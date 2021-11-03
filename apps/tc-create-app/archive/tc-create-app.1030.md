# Add OBS Study Notes

Create branch: set-branch feature-cn-1030-add-obs-sn)d-obs-sn

Create the translatable file and the row header file by copying from Bible SN files:
Now create these copy to create as follows:
```
cp TranslatableSnTsv.js TranslatableObsSnTsv.js
cp RowHeaderSn.js RowHeaderObsSn.js
```

*In TranslatableObsSnTsv.js*
- Change all SnTSV to ObsSnTSV
- Change all RowHeaderSn to RowHeaderObsSn
- Columns for en_obs-sn are: `Reference\tID\tTags\tSupportReference\tQuote\tOccurrence\tNote`
- No need to change this:
```js
// columns Reference	ID	Tags	SupportReference	Quote	Occurrence	Note
const _config = {
  compositeKeyIndices: [0, 1],
  columnsFilter: ['Reference', 'ID', 'Tags', 'Quote', 'Occurrence'],
  columnsShowDefault: [
    'Reference','SupportReference','Note',
  ],
}
;
```
- No need (?) to change the checker function, since this is same as what was copied from. Here is the validation code:

```js
  const _onValidate = useCallback(async (rows) => {
    // NOTE! the content on-screen, in-memory does NOT include
    // the headers. This must be added.
    let data = [];
    const header = "Reference\tID\tTags\tSupportReference\tQuote\tOccurrence\tNote\n";
    if ( targetFile && rows ) {
      data = await contentValidate(rows, header, cv.checkNotesTSV7Table, langId, bookId, 'TN2', validationPriority);
      if ( data.length < 2 ) {
        alert("No Validation Errors Found");
        setOpen(false);
        return;
      }
    
      let ts = new Date().toISOString();
      let fn = 'Validation-' + targetFile.name + '-' + ts + '.csv';
      csv.download(fn, csv.toCSV(data));    
    }

    setOpen(false);
  },[targetFile, validationPriority, langId, bookId]);

  const onValidate = useCallback( (rows) => {
    setOpen(true);
    setTimeout( () => _onValidate(rows), 1);
  }, [_onValidate]);

```

- Update Translatable.js to have the file name pattern and use this component

As of 2021-10-07, application is hanging; not sure what the issue may be yet, but found that this attempt to get the published content has a problem.
- using this URL: https://qa.door43.org/api/v1/repos/unfoldingword/en_obs-sn
- it reports that v1 is the latest production release, but it isn't.