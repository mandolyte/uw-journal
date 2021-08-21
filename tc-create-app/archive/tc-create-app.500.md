# Issue 500 Adding Content Valiation (on open)

## 2021-02-02

Current solution for showing on open issues. In Workspace component:

```js
    if (sourceRepoMemo && sourceFilepath && filepath) {
      if (sourceFilepath === filepath) {
        _component = (
          <TargetFileContextProvider 
            validated={validated} onValidated={setValidated} 
            onCriticalErrors={setCriticalErrors}
          >
            {
              (validated && <Translatable />) 
              || 
              (criticalErrors.length > 0 && 
                <Alert severity="error" onClose={() => {
                    setCriticalErrors([]);
                    setSourceRepository(undefined);
                  }}
                
                  action={
                    <Button color="inherit" size="small" >
                      CLOSE
                    </Button>
                  }
                
                >
                  <AlertTitle>This file cannot be opened by tC Create. Please contact your administrator to address the following error(s).</AlertTitle>
                  {
                    criticalErrors.map( (msg,idx) => {
                      return (
                        <>
                        <Typography key={idx}>
                          On <Link href={msg[0]} target="_blank" rel="noopener">
                            line {msg[1]}
                          </Link>
                          &nbsp;{msg[2]}&nbsp;{msg[3]}&nbsp;{msg[4]}&nbsp;{msg[5]}
                        </Typography>
                        </>
                      )
                    })
                  }
                </Alert>
              )
            }
          </TargetFileContextProvider>
        );
      }
    }
```
```js
import React, { useContext, useMemo, useState, useCallback } from 'react';

import { Dialog, DialogTitle, DialogContent, DialogContentText, DialogActions, Button } from '@material-ui/core';

  const handleClose = useCallback( () => {
    setCriticalErrors([]);
    setSourceRepository(undefined);
  }, [setCriticalErrors, setSourceRepository]);

  const component = useMemo(() => {
    let _component = <ApplicationStepper />;

    if (sourceRepoMemo && sourceFilepath && filepath) {
      if (sourceFilepath === filepath) {
        _component = (
          <TargetFileContextProvider 
            validated={validated} onValidated={setValidated} 
            onCriticalErrors={setCriticalErrors}
          >
            {
              (validated && <Translatable />) 
              || 
              (criticalErrors.length > 0 && 
                <Dialog
                  disableBackdropClick
                  open={!validated}
                  onClose={handleClose}
                  aria-labelledby="alert-dialog-title"
                  aria-describedby="alert-dialog-description"
                >
                  <DialogTitle id="alert-dialog-title">
                  This file cannot be opened by tC Create. Please contact your administrator to address the following error(s).
                  </DialogTitle>
                  <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                    {
                      criticalErrors.map( (msg,idx) => {
                        return (
                          <>
                          <Typography key={idx}>
                            On <Link href={msg[0]} target="_blank" rel="noopener">
                              line {msg[1]}
                            </Link>
                            &nbsp;{msg[2]}&nbsp;{msg[3]}&nbsp;{msg[4]}&nbsp;{msg[5]}
                          </Typography>
                          </>
                        )
                    })}
                    </DialogContentText>
                  </DialogContent>
                  <DialogActions>
                    <Button onClick={handleClose} color="primary">
                      Close
                    </Button>
                  </DialogActions>
                </Dialog>
              )
            }
          </TargetFileContextProvider>
        );
      }
    }
    return _component;
  }, [sourceRepoMemo, sourceFilepath, filepath, validated, criticalErrors, handleClose]);
```

## 2020-11-19

In the below, last two lines are the before and after

```sh
updated github workflow:
name: Install, Build & Run Cypress

on: [push]
env:
  CHILD_CONCURRENCY: 1
  NODE_ENV: test
  COVERALLS_REPO_TOKEN: ${{ secrets.COVERALLS_REPO_TOKEN }}
jobs:
  test:
    runs-on: macos-latest
    strategy:
      matrix:
        node-version: [10.15.1]
    steps:
      - name: Extract Branch Name
        run: echo "::set-env name=BRANCH::${GITHUB_REF##*/}"
        run: echo "BRANCH=${GITHUB_REF##*/}" >> $GITHUB_ENV
```

## 2020-11-12

Interesting... I tried this:
```js
  if ( state ) {
    console.log("target file context value=", context);
    const _name  = state.name.split('_');
    const langId = _name[0];
    const bookID = _name[2].split('-')[1].split('.')[0];
    const rawResults = validate(langId, bookID, state.content).then(
      (value) => {
        console.log("inside promise:", value);
        return (
          <TargetFileContext.Provider value={context}>
            {children}
          </TargetFileContext.Provider>
        );
      },
      (value) => {
        console.log("[TargetFile.context.js] rejected promise in validate on open");
      }
    );
    console.log("target file validation=", rawResults);
  }
  /*
  return (
    <TargetFileContext.Provider value={context}>
      {children}
    </TargetFileContext.Provider>
  );
  */
```
But then the component didn't return anything to render. Sort of expected that.


```js
      validate(langId, bookID, state.content).then(
        (value) => {
          let criticalNotices = [];
          for ( let i=0; i<value.noticeList.length; i++ ) {
            let notice = value.noticeList[i];
            if ( notice.priority >= 746 ) {
              let msgArray = [];
              msgArray.push(`${link}#L${notice.lineNumber}`);
              msgArray.push(`${notice.lineNumber}`);
              msgArray.push(notice.message);
              msgArray.push(notice.fieldName ? notice.fieldName : '');
              msgArray.push(notice.details ? notice.details : '');
              msgArray.push(notice.rowID ? `with row id=${notice.rowID}` : '');
              let msg = `On {<Link href="">}line ${notice.lineNumber}{</Link>}, ${notice.message}.`;
              if ( notice.fieldName !== undefined ) msg = msg + '\n    ' + notice.fieldName;
              if ( notice.details !== undefined ) msg = msg + ' ' + notice.details;
              if ( notice.rowID !== undefined ) msg = msg + ' with row id=' + notice.rowID;
              criticalNotices.push(msgArray);
            }
          }
```

## 2020-11-11

This CV check should be light, mostly to ensure that the file is parseable and will not crash the app.

Based on conclusions in this issue:

- We will go with a model dialog box
- The box will show high priority errors
  - include the message and other info needed to understand the error
  - the "line" number with link to open DCS to location
- Provide two buttons:
  - Close which will not advance the stepper
  - Continue which will advance the stepper

One goal for new content is to avoid creating the branch with default (bad) content. Otherwise, the branch will need to be removed before continuing to work on the file after the source is corrected.


Begin tracing...

In target file context, this:
```js
  const {
    state, actions, component, components, config,
  } = useFile({
    config: (authentication && authentication.config),
    authentication,
    repository: targetRepository,
    filepath,
    onFilepath: setFilepath,
    defaultContent: (sourceFile && sourceFile.content),
  });

  const context = {
    state,
    actions,
    component,
    components,
    config,
  };

  console.log("target file context value=", context);
```
I added the logger. This implies to me that the file context is created along with the branch with the source content copied over if the branch/file does not yet exist.
The log message should confirm this. First I deleted my branch in uw en_tn:
![mybranch](Screenshot_2020-11-11_092120.png "My uW en_tn branch")

Next started tc-create and went to Ruth UTN. This should create a branch from scratch and populate the target content with the source.

Console:
![console](Screenshot_2020-11-11_093134.png)

So by the time, it returns to tc-create the branch is there with content. Added notes to the issue. For now, just validate the target content.



Here is essential steps to run the validation and turn it into a CSV:
```js
      const _name  = targetFile.name.split('_');
      const langId = _name[0];
      const bookID = _name[2].split('-')[1].split('.')[0];
      const rawResults = await cv.checkTN_TSVText(langId, bookID, 'dummy', rows, '');
      const nl = rawResults.noticeList;
      let hdrs =  ['Priority','Chapter','Verse','Line','Row ID','Details','Char Pos','Excerpt','Message','Location'];
      let data = [];
      data.push(hdrs);
      Object.keys(nl).forEach ( key => {
        const rowData = nl[key];
        csv.addRow( data, [
            String(rowData.priority),
            String(rowData.C),
            String(rowData.V),
            String(rowData.lineNumber),
            String(rowData.rowID),
            String(rowData.fieldName || ""),
            String(rowData.characterIndex || ""),
            String(rowData.extract),
            String(rowData.message),
            String(rowData.location),
        ]);
      });
```


## Issue 500

Doing a light check: 
- are headers present?
- do all rows have the correct number of cells?

To create test case:
- go to https://qa.door43.org/unfoldingWord/en_tn
- click branches
- remove mine `cecil.new-tc-create-1`, if present
- select book of Jude
- click the pencil edit icon
- remove the header row and remove a tab character from row 17
- change the branch name to `cecil.new-tc-create-1`
- click the propose change button

now the book of Jude in my tc-create branch will be mangled.




## Scratch

```js
          <TargetFileContextProvider 
            validated={validated} onValidated={setValidated} 
            onCriticalErrors={setCriticalErrors}
          >
            {
              (validated && <Translatable />) 
              || 
              (criticalErrors.length > 0 &&
                <Alert severity="error" onClose={() => {
                  // not sure what to put here...
                }}>
                  <AlertTitle>Error</AlertTitle>
                  {
                    criticalErrors.map( (msg,idx) => {
                      return (
                        <>
                        <Typography key={idx}>
                          On <Link href={msg[0]} target="_blank" rel="noopener">
                            line {msg[1]}
                          </Link>
                          &nbsp;{msg[2]}&nbsp;{msg[3]}&nbsp;{msg[4]}&nbsp;{msg[5]}
                        </Typography>
                        </>
                      )
                  })}
                  <Typography>Please correct resource or close and select another</Typography>
                </Alert>
              )
            }
          </TargetFileContextProvider>
```