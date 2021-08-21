# feature-cn-28-add-cv-to-bpa

*Issue description:* Per brief discussion, the current thinking is to include a new table of results after/before the work count frequency table that shows results of the Content Validation RCL. Only the master branch must be considered.

Issue: https://github.com/unfoldingWord/book-package-app/issues/28

## 2020-09-23

Thinking about how to move file fetching to from the CV package into the CV app.

Assumptions:
- All needed repos can be fetched as zip files and stored in indexedDB
- Performance of using the stored zip files is significantly faster than fetching
  - If not we could try other zip reader packages; 
  - We only need to read, not write
- Fetching can be conditional based the commit hash of fetched repos

- Step one. copy the CV package core book package code into CV app and get it working again
  - all files fetched by CV app
  - file content given to CV pkg to validate
  - results are returned as done currently and then displayed for user
- Step two. *unconditionally* fetch all needed repos based selected org and language
  - Will need to swap first two Stepper steps: do org/lang picker first, then show books picker
  - Once org/lang selected then run the preload function to fetch the repos needed
- Step three. *conditionally* fetch all needed repos
  - Will need to store the commit hash of fetched repos in order to do conditionally fetch
  - Based on org/lang selection, first check commit hash of each repo
  - If that repo is already downloaded, do not fetch again
  - If that repo is not downloaded, then:
    - delete existing out-of-date data plus the stored associated commit hash
    - fetch the repo zip
    - store in indexedDB
    - store the commit hash

## 2020-09-21

Added support for GLs by:

- inserting new step 2 to prompt for language
- see details in issue #5
- used radio buttons
- English is default and still works.



## 2020-09-02

Today, I will try to quickly create an app wrapper for content validation library.

- content validation is now on NPM
- added to package JSON for the app
- found a UI component that wasn't exported... need to publish locally with fix.
- in content-validation module: `yalc publish`
- in app: `yalc link uw-content-validation`






## Hypotheses

### Alternative 1

Assumption(s): that the content validation RCL can be provided "in-memory" a text to validate.
Given:
- BPA already fetches all related content to for one or more books
- BPA already validates the following existence tests:
  - UTA links in UTN Support Reference column
  - UTW links in Original Hebrew and Greek 

Then BPA, as it fetches content for word count purposes, can also provide the same content to the validation RCL. The results of the validation can be returned and shown.

Reservations: the checks performed by the validator may require following embedded links in the content (such the markdown column in a UTN file).


### Alternative 2

Assumption: the content validator must be provided a bookId which is used to fetch all resources and validate interpendencies.

In this case, I propose that a new step be added on the end to perform the validation, since it will need to do all the fetches a second time. I'd rather not slow down current count results. It would be nice to know if the API allows:

- segregating results to each book and resource type
- line/row number of error so that a link can be constructed to take the user to the location in DCS
