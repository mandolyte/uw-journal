# CV Diary
## Overview
### API Notes
As of 2021-08-26:
-   for TN (7 col): checkNotesTSV7Table with repo code "TN2"
-   for SN: checkNotesTSV7Table with repo code "TN2"
-   for SQ: checkQuestionsTSV7Table with repo code "TQ2"
-   for TQ: checkQuestionsTSV7Table with repo code "TQ2"
-   for TWL: checkTWL_TSV6Table with repo code "TWL"
-   for OBS-TQ: may be in markdown
-   for OBS-TN: may be in markdown
-   for OBS-SN: may be in markdown
-   for OBS-SQ: may be in markdown

This will change! Identified simplification using resource specific wrapper functions are coming.

See [here](https://github.com/unfoldingWord/uw-content-validation/blob/new.2021.September.5/src/core/wrapper.js)

## Daily Journal

### 2021-10-05 add cv to tq

Using refactored cv work done earlier.

### 2021-10-05 (bp requirements)

How should a book package entry point (API call) work? First we have an example of how it should work at: https://unfoldingword-box3.github.io/content-validation-app/

While this Q&D (quick and dirty) app was limited for a specific initiative, it has the fundamentals needed to think thru how a BP API should work.

**Inputs**
- Organization
- Language
- Book of the Bible

*Note: resource is not an input! All resources must be in repositories, where the repositories have standard names.*

**Outputs**
- Notices when a resource repository does not exist
- Notices when resource files do not exist
- Other notices as normal

Here is an example of normal validation results:
![[Pasted image 20211005071919.png]]

Here is an example of results where things are missing:
![[Pasted image 20211005072103.png]]

Here is an example of results where entire repos are missing:
![[Pasted image 20211005072212.png]]



### 2021-10-04

Today testing in branch `feature-cn-953-add-sn-sq`. This includes refactoring tN for CV and adding CV to more resource types.

*On TWL CV* 

It appears to be checking recursively. Robert Hunt recommends adding these options: `disableLinkedTAArticlesCheckFlag`, `disableLinkedTWArticlesCheckFlag`, `disableLexiconLinkFetchingFlag`.

So now the code in TWL looks like:
```js
      data = await contentValidate(rows, header, cv.checkTWL_TSV6Table, langId, 
        bookId, 'TWL', validationPriority, 
        {suppressNoticeDisablingFlag: false,
          disableLinkedTAArticlesCheckFlag: true,
          disableLinkedTWArticlesCheckFlag: true,
          disableLexiconLinkFetchingFlag: true,
        }
      );
```

It is passing an options object, so I need to add this parameter to the "contentValidate()" function.

Results: Looks good for Book of Ruth; found no errors. I did ask Robert whether the disable flags did an existence check (I want that - just not further nested, recursive checking).

Using book of Acts: filed issue: https://github.com/unfoldingWord/uw-content-validation/issues/221

*On TN CV* 

Running checks on the Book of Acts TN file.

