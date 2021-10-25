---
id: efc9f450-8b74-4685-86e8-24b08b604453
title: '818'
desc: ''
updated: 1619438718811
created: 1618922356017
---

# Issue 818

This is a field reported problem and as of today (2021-04-20), it seems to happen only in Luke.

The steps to reproduce are complex. See Appendix A.

## 2021-05-03

Today I have the `onOpenValidation` being pushed all the way down to `ensureContent`. 

Now at present `TargetFileContextProvider` looks like:
```js
function TargetFileContextProvider({
  validated, onValidated, onCriticalErrors, children
})
```

The three validation parameters are passed in. I need to change this to merely pass in the single 
validation function. And have it work the same way it does now.

First off, let's commit our current changes to have a clean point. Have to do this in both tc-create and GRT (my branch is named the same in both).

Now:
```js
function TargetFileContextProvider({
  onOpenValidation, children
})
```
I have commented out the `useEffect` since it isn't needed anymore.

*Test:*
1. in QA DCS/ unfolding word tN, remove my branch
2. verify that normal happy path works ok using book of Jude
3. this wasn't true at first, but I suspected it may have been due to the default value for "validated" state. Actually, don't need this now and can go solely by whether the `criticalErrors` state is an empty array or not.
7. After this bit of refactoring now works on the happy path.
8. in Gitea, in my branch, cut first row of Jude, and paste back as third row and remove a column
9. commit it
10. in tc-create, go back to Jude... expect errors reported on open

Looks good. Both happy path and defect path work.

*Next:* 
Do the same for the source side file context.
1. Add state for critical errors to App.js (copy from Workspace.js)
2. And whatever else is needed.
3. In GRT, update FileContext to be like TargetContext

*Test:*
Let's damage master for Ruth...
- first, remove my branch (it already has a clean copy of Ruth).
- second, cut header from Ruth in the master branch; paste it back in below; remove a column
- third, (since this is master) create a patch branch and propose the PR
- continue to approve the PR by clicking "Squash and Merge"
- delete my patch branch after committing to master
- Now try it... it works

Does it create a branch? NO!! Super...

Remaining to do:
1. add dup ID validation
2. distinguish between source and target in error messages (see updated text in the issue)
3. clean up the code by removing all the console debug logging




## 2021-04-29

In `Workspace.js` there are a few things that make this work:
1. There are two states that manage the validation error dialog:
	- a Boolean state variable named `validated` with `setValidated`
	- an array state variable named `criticalErrors` with `setCriticalErrors`
	- from the AppContext, we obtain the action `setSourceRepository`
	- a `handleClose` action for the error dialog box to set the step back to source repo picker
2. I want to add validation to the `useFile` hook. This hook is used in `FileContextProvider` and in `TargetFileContextProvider`. The former is used in `App.js` and the latter in `Workspace.js`. *Do I need to replicate all the `Workspace.js` stuff above in `App.js`?*


## 2021-04-28

This image shows the progression of things:

![[Pasted image 20210428085645.png]]

First the source data runs thru load() and then ensureContent(). At this point it is done. Since the source file was picked from a UI dialog, it will always be there (if it didn't exist the user could not have selected it).

Thus we need to insert an on open validation here and throw an error if any problems are found on the source side.

Second, what is shown is after I removed my branch, so starting from a clean spot; note also I am using a book that is not in the es-419_gl/es-419_tn repo (thus the second readContent() below will fail too).
- the first readContent() fails (notice the GET 404 error); this one is after my branch: `GET https://bg.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv?noCache=0.8747914809271238&ref=cecil.new-tc-create-1 404 (Not Found)`
- then the second readContent() fails, which is going after master (ie, it will be editing an existing translated file): `GET https://bg.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv?noCache=0.19272907416648977 404 (Not Found)`
- The "decode" on a null will throw an error, so now in the catch(), createContent will begin to create the user branch and file with the content from the source. The POST shows as: `POST https://bg.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv 404 (Not Found)`
- Afterwards, the contentObject has the content.

So the plan for a) first time editing or b) source updated defectively:
- as soon as source content loaded, validate it. If defective, throw an error.
- may need a variable that I can test when this happens so that createContent() is never attempted as well.

Now the flow for a) not first time editing and b) source is ok, but target becomes defective.

*Test 1*: I'm in Jude, first back out and go back in; capture console log. This is the "happy path":
![[Pasted image 20210428092152.png]]

- the source from `unfoldingword` loads ok
- the target from my branch in `es-419_gl` works ok
- both above succeed with first `readContent()`

*Test 2*: delete my branch and select a file that exists in the target repo.
These are the files in the repo:
![[Pasted image 20210428092619.png]]
So let's go 3JN:
![[Pasted image 20210428093245.png]]

This is complicated!
- source from `unfoldingword` is ok
- content from my user branch fails (expected); but there are actually two GETs, both failing with 404
- first `readContent()` fails
- second `readContent()` succeeds (it is in the master branch)
- third, `updateContent()` runs and will cause my branch to be created. There appear to be 3 PUTs that fail, with 404, 404, and 500 respectively.
- fourth, nested within above, `createContent()` runs which POSTs the data. There are two POSTs, both which return 422... *what is that about?*
- finally, the content is there as shown when the "runOnOpenValidation()" runs.

I think validating after both `readContent()` functions will do the job. This will catch both problems in the source and problems in the target.

*How to pass down the validation function*

To determine this I need to trace backwards `ensureConent()`.
```
ensureContent() -- src/core/gitea-api/repos/contents/content.ts
	-> ensureFile() -- src/components/file/helpers.js
		-> load() -- src/components/file/useFile.js
		-> createFile() -- also useFile.js

useFile()
	-> FileContextProvider() -- GRT @ src/components/file/File.context.js
	-> TargetFileContextProvider() -- tCC @ src/components/core/TargetFile.context.js
```

Based on above, looks like the on-open validation function needs to be passed into File Context.

In that case, this in tcc@src/App.js:
```js
              <FileContextProvider
                authentication={authentication}
                repository={sourceRepository}
                filepath={filepath}
                onFilepath={setFilepath}
              >
```

would become something like this:
```js
              <FileContextProvider
                authentication={authentication}
                repository={sourceRepository}
                filepath={filepath}
                onFilepath={setFilepath}
				onOpenValidation={onOpenChecks}
              >
```

and in tcc@src/Workspace.js:
```js
          <TargetFileContextProvider 
            validated={validated} onValidated={setValidated} 
            onCriticalErrors={setCriticalErrors}
          >
```

would become something like:
```js
          <TargetFileContextProvider 
            validated={validated} 
			onValidated={setValidated} 
            onCriticalErrors={setCriticalErrors}
			onOpenValidation={onOpenChecks}
          >
```




## 2021-04-27
```
         1         2         3         4         5         6         7
12345678901234567890123456789012345678901234567890123456789012345678901234567890
```
GRT logic flow:
GRT->src/components/file/File.context.js
-> useFile.js:  load()
		-> helper.js:  ensureFile()
				-> core/gitea-api/repos/contents/contents.ts: ensureContent()



## 2021-04-26

As of last week, I have the bones ready:
- In GRT `src/components/file/helpers.js` there is now a stub function named `runOnOpenValidation`. This function will have the file metadata *and* content, plus the onValidation function provided by the app to run against the content.

The validation function must:
- throw an error in order to prevent content from being created in the target repo
- it must set state data: 
  - set a Boolean to cause a dialog to be shown
  - set an array of messages to dispaly on the dialog

It all starts in the app's TargetFile.context.js code:
```js
  const fake_onopen_validator = () => {console.log("FAKE!!")}

  const {
    state, actions, component, components, config,
  } = useFile({
    config: (authentication && authentication.config),
    authentication,
    repository: targetRepository,
    filepath,
    onFilepath: setFilepath,
    defaultContent: (sourceFile && sourceFile.content),
    onOpenValidation: fake_onopen_validator,
  });
```
The above shows the passing of a new property to the hook, which above is just a fake stub for one.

Then in useFile.js, the validation logic is inserted:
```js
  const load = useCallback(async () => {
    if (config && repository && filepath) {
      const _file = await ensureFile({
        filepath, defaultContent, authentication, config, repository, branch,
      });
      // let content;
      // content = await repositoryActions.fileFromZip(filepath);
      const content = await getContentFromFile(_file);
      onOpenValidation && runOnOpenValidation({
        file: _file, 
        content: content, 
        onValidate: onOpenValidation
      });

      update({
        ..._file, branch, content, filepath: _file.path,
      });
    }
  }, [authentication, branch, config, defaultContent, filepath, repository, update]);
```

Discovered that source does not use a file context... need to find where/how/when source content is fetched, so that I can insert the three checks required.

## 2021-04-23

Changes will be made to both tc-create and gitea-react-toolkit (GRT). Set branch to, say, "feature-cn-818-on-open-validation"

## 2021-04-22

Yesterday, during sprint planning this issue was reworked to address the root cause, namely, source data issues.

It now is:

```
In order to ensure that the user can continue to work the initial validation check must be done on the source side as we are already doing on the target. Errors in source side should say:
"tC Create is unable to continue. The master data has the following error: duplicate ID(s) ."
"tC Create is unable to continue. The master data has the following error: ."
"tC Create is unable to continue. The master data has the following error: ."

"Please take a screenshot and contact your administrator."

DoD:
On Open validation is run against both the source and target files.
```

The target context provider is created largely from a useFile hook. This hook takes as input the source file content, which if the target file does not exist yet, it will be created with the source content as the default content.

This means we have the following scenarios:

1. Target file does not exist
2. Target file does exist

Test: if target does not exist and source has a defect, then does anything get created in the target?

1. delete branch cecil.new-tc-create-1 in QA DCS `unfoldingword/en_tn`
2. edit the book of Jude by removing the header row. Submit patch as a PR. The approve PR and squash merge and delete my patch branch.
3. Check and confirm master changed. -- confirmed
4. Now use tc-create to edit this file
5. The validation error properly shows, however, the branch is created and the damaged file is created as well.

**Implication** the on-open check needs to be moved to the useFile hook.

### Solution Design

**Concept** add a new optional onOpenValidation function parameter to the `useFile()` function in gitea-react-toolkit. If provided, then the content of the file will be passed to this function for validation. In addition, the `useFile()` function will return a new property as part of the state object. The new property will be an object containing two properties: a boolean indicating pass/fail of the validation and an array of error messages.

**Scenarios**
1. If no validation function is provided, then the state validation property will be null or undefined.
2. If validation function is provided and the validation passes, then the object will return a value of true for the `valid` property and an empty array for the messages property.
3. If validation is function provided and the validation fails, then the object will return a value of false for the `valid` property and the messages array will list all the errors found. The format of the messages will be completely determined by the validation function provided.

**Notes**
- Existing code that does not provide the new validation parameter will not need to be changed; nor does it need to capture the validation state return value.
- Code that needs to validate the content must provide the new function to `useFile()` and it must examine the return value in order properly act on it.

The plan is to split the current on-open validation logic into two parts:
- The part that updates state of the app must remain in the app
- The part that is pure code looking at content will be moved to GRT itself to make it re-usable.

At the moment, TN (both 9 and 7 column versions) and TWL are targeted for on-open validation for a total of three formats. While superficially the checks are the same, they are each unique. The checks are:
1. Does the file have the correct header row?
2. Does each row in the file have the correct number of columns?
3. (new) Are there duplicate row IDs within chapter/verse?



## 2021-04-20

Today, I am reproducing the problem using door43-catalog version of Luke.

Exceptions:
- only added the ID column
- on row `l400`, I changed support ref to figs-l400

Now on page forward and back, the figs-l400 is now on the previous row (l365), not row l400:
![pic](assets/Screenshot_2021-04-20_085535.png)

Raw data after save:
![pic2](assets/Screenshot_2021-04-20_090305.png)


*Variation 1*
Let's use the search function to pull up row l400 and change that way.
- delete branch
- close file, returning to resource selection page
- go back to Luke
- add ID column
- click search and enter l400
- make the change

Still happened:
![pic3](assets/Screenshot_2021-04-20_091209.png)

Talked to @klappy and he suggested a binary search to find the exact breaking point. We know it is less than line 1100.



# Appendix A - steps to reproduce

These are taken from the Issue itself; written by @cckozie

Field reported by Larry S.
movie here: https://drive.google.com/file/d/1CgS3NKv5TfR7Qo4XFsKnitTR6udfiJMi/view?usp=sharing

Follow these steps to reproduce the problem:

* Delete the Luke tN file from your user's repo
* Log in to tC Create using that users credentials
* Select the Luke tN project
* Verify that the app is set for 25 rows/page
* Use the Next Page button to page forward until you get to rows 1076-1100
* Scroll down until you get to ID l400 (lower case L)
* Open the Columns modal and show all columns
* Change the SupportReference field on row l400 to some other text including "l400" in that text (I used Testing l400)
* Click out of the SupportReference field and wait for the UI to catch up
* Click the Next Page button to go to the next page
* After it pages forward, click the Previous Page button to go back to the page you were on
* Scroll down to the l400 row
* Notice that the SupportReference field is now back to what it was before you changed it
* Scroll up 1 row to l365
* Note that the text you entered for row l400 is now in row l365
* Click the Save button
* After the save is complete, click the button to go to the file in the repo
* Go into edit mode in the repo and scroll down to line 1100
* Note that the data has been saved to the wrong row as it was displayed after paging forward and back
* In my testing it seems that it only fails after row 1100, but it does fail consistently.

In another comment, @elsylambert says:
*I was able to replicate this issue with the steps mentioned above. The data was wrongly written on the row above as reported in the issue. However, I was able to do so only in Luke. I tested Acts, Rev, Jer, Psalms, Proverbs, Genesis and few other big books but none of them showed this behaviour.

Luke has 4490 rows, Psalms has 5219 rows, but only Luke shows consistently bad behaviour after 1100 rows.*
