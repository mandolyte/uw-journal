# tc-create-app#953
## Issue:

the https://git.door43.org/unfoldingWord/en_sq repository should be listed as a source text for English. _Note_: Currently (7/9/21) there are no releases of this repo.  Once a release is made other languages should be able to translate from it. If there are no releases then there will not be any source text. The warning should read, "No source text has been released for this resource, please choose a different resource." After clicking 'ok', the user should be returned to the resource list.

## DoD:

English users (at a minimum) can choose Study Questions as a resource and edit it (edit mode). If there's no released text then users in translate mode should be warned after they choose a language. 

## 2021-08-26

The branch that has the new TSV parser is not merged yet. So I'll try to do a branch of a branch.

So first switch to `feature-cn-788-impl-tsv-parser`. 

Then: `git checkout -b feature-cn-953-add-sn-sq`
Then: `git push --set-upstream orign feature-cn-953-add-sn-sq`

Now create these copy to create as follows:
```
cp TranslatableTnTsv.js TranslatableSnTsv.js
cp TranslatableTnTsv.js TranslatableSqTsv.js
cp RowHeaderTn.js RowHeaderSn.js
cp RowHeaderTn.js RowHeaderSq.js
```

*In TranslatableSqTsv.js*
- Change all TnTSV to SqTSV
- Change all RowHeaderTn to RowHeaderSq

- Modify this:
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
to be this:
```js
// columns Reference	ID	Tags	Quote	Occurrence	Question	Response
const _config = {
  compositeKeyIndices: [0, 1],
  columnsFilter: ['Reference', 'ID', 'Tags', 'Quote', 'Occurrence'],
  columnsShowDefault: [
    'Reference','Question',
  ],
}
;
```

- Change validation function from:
```js
 const rawResults = await cv.checkTN_TSV9Table(langId, 'TN', bookID, 'dummy', rows, '', {suppressNoticeDisablingFlag: false});
```
to:
```js
 // const rawResults = await checkQuestionsTSV7Table(languageCode, repoCode, bookID, filename, tableText, givenLocation, checkingOptions);

 const rawResults = await cv.checkQuestionsTSV7Table(langId, 'TQ2', bookID, 'dummy', rows, '', {suppressNoticeDisablingFlag: false});
```




*In TranslatableSnTsv.js*
- Change all TnTSV to SnTSV
- Change all RowHeaderTn to RowHeaderSn
- Update `_config`
- Update `generateRowId`
- Update validate:
```js
      // checkNotesTSV7Table with repo code "TN2"
      const rawResults = await cv.checkNotesTSV7Table(langId, 'TN2', bookID, 'dummy', rows, '', {suppressNoticeDisablingFlag: false});
```

*In RowHeaderSn.js*
- change all `RowHeaderTn` to `RowHeaderSn`






