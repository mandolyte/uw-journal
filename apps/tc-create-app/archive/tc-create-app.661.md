# Add OBS Study Questions

Create branch: set-branch feature-cn-661-add-obs-sq

Create the translatable file and the row header file by copying from Bible SN files:
Now create these copy to create as follows:
```
cp TranslatableSqTsv.js TranslatableObsSqTsv.js
cp RowHeaderSq.js RowHeaderObsSq.js
```

*In TranslatableObsSqTsv.js*
- Change all SqTSV to ObsSqTSV
- Change all RowHeaderSq to RowHeaderObsSq
- Columns for en_obs-sq are: `Reference\tID\tTags\tQuote\tOccurrence\tQuestion\tResponse`
- No need to change this:
```js
const _config = {
  compositeKeyIndices: [0, 1],
  columnsFilter: ['Reference', 'ID', 'Tags', 'Quote', 'Occurrence','Response'],
  columnsShowDefault: [
    'Reference','Question',
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
    const header = "Reference\tID\tTags\tQuote\tOccurrence\tQuestion\tResponse\n";
    if ( targetFile && rows ) {
      data = await contentValidate(rows, header, cv.checkQuestionsTSV7Table, langId, bookId, 'TQ2', validationPriority);
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

As of 2021-10-07, this is crashing.
```
xhr.js:175 POST https://qa.door43.org/api/v1/repos/unfoldingWord/en_obs-sq/contents/sq_OBS.tsv 422 (Unprocessable Entity)
dispatchXhrRequest @ xhr.js:175
xhrAdapter @ xhr.js:20
dispatchRequest @ dispatchRequest.js:40
Promise.then (async)
request @ Axios.js:64
Axios.<computed> @ Axios.js:89
wrap @ bind.js:11
_callee3$ @ http.js:385
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ http.js:26
_next @ http.js:48
(anonymous) @ http.js:55
(anonymous) @ http.js:44
post @ http.js:401
_callee$ @ contents.js:132
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ contents.js:28
_throw @ contents.js:54
Promise.then (async)
asyncGeneratorStep @ contents.js:38
_next @ contents.js:50
(anonymous) @ contents.js:57
(anonymous) @ contents.js:46
createContent @ contents.js:164
_callee5$ @ contents.js:461
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ contents.js:28
_throw @ contents.js:54
Promise.then (async)
asyncGeneratorStep @ contents.js:38
_next @ contents.js:50
Promise.then (async)
asyncGeneratorStep @ contents.js:38
_next @ contents.js:50
Promise.then (async)
asyncGeneratorStep @ contents.js:38
_next @ contents.js:50
Promise.then (async)
asyncGeneratorStep @ contents.js:38
_next @ contents.js:50
(anonymous) @ contents.js:57
(anonymous) @ contents.js:46
ensureContent @ contents.js:491
_callee$ @ helpers.js:117
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ helpers.js:63
_next @ helpers.js:85
(anonymous) @ helpers.js:92
(anonymous) @ helpers.js:81
ensureFile @ helpers.js:142
_callee2$ @ useFile.js:352
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ useFile.js:147
_next @ useFile.js:169
(anonymous) @ useFile.js:176
(anonymous) @ useFile.js:165
(anonymous) @ useFile.js:649
commitHookEffectListMount @ react-dom.development.js:19607
commitPassiveHookEffects @ react-dom.development.js:19644
callCallback @ react-dom.development.js:189
invokeGuardedCallbackDev @ react-dom.development.js:238
invokeGuardedCallback @ react-dom.development.js:291
flushPassiveEffectsImpl @ react-dom.development.js:22708
unstable_runWithPriority @ scheduler.development.js:656
runWithPriority$1 @ react-dom.development.js:11076
flushPassiveEffects @ react-dom.development.js:22676
performSyncWorkOnRoot @ react-dom.development.js:21591
(anonymous) @ react-dom.development.js:11130
unstable_runWithPriority @ scheduler.development.js:656
runWithPriority$1 @ react-dom.development.js:11076
flushSyncCallbackQueueImpl @ react-dom.development.js:11125
flushSyncCallbackQueue @ react-dom.development.js:11113
scheduleUpdateOnFiber @ react-dom.development.js:21053
dispatchAction @ react-dom.development.js:15633
_onOpenValidation @ App.js:69
_callee5$ @ contents.js:396
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ contents.js:28
_next @ contents.js:50
Promise.then (async)
asyncGeneratorStep @ contents.js:38
_next @ contents.js:50
Promise.then (async)
asyncGeneratorStep @ contents.js:38
_next @ contents.js:50
(anonymous) @ contents.js:57
(anonymous) @ contents.js:46
ensureContent @ contents.js:491
_callee$ @ helpers.js:117
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ helpers.js:63
_next @ helpers.js:85
(anonymous) @ helpers.js:92
(anonymous) @ helpers.js:81
ensureFile @ helpers.js:142
_callee2$ @ useFile.js:352
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ useFile.js:147
_next @ useFile.js:169
(anonymous) @ useFile.js:176
(anonymous) @ useFile.js:165
(anonymous) @ useFile.js:649
commitHookEffectListMount @ react-dom.development.js:19607
commitPassiveHookEffects @ react-dom.development.js:19644
callCallback @ react-dom.development.js:189
invokeGuardedCallbackDev @ react-dom.development.js:238
invokeGuardedCallback @ react-dom.development.js:291
flushPassiveEffectsImpl @ react-dom.development.js:22708
unstable_runWithPriority @ scheduler.development.js:656
runWithPriority$1 @ react-dom.development.js:11076
flushPassiveEffects @ react-dom.development.js:22676
performSyncWorkOnRoot @ react-dom.development.js:21591
(anonymous) @ react-dom.development.js:11130
unstable_runWithPriority @ scheduler.development.js:656
runWithPriority$1 @ react-dom.development.js:11076
flushSyncCallbackQueueImpl @ react-dom.development.js:11125
flushSyncCallbackQueue @ react-dom.development.js:11113
scheduleUpdateOnFiber @ react-dom.development.js:21053
dispatchAction @ react-dom.development.js:15633
_onOpenValidation @ App.js:69
_callee5$ @ contents.js:396
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ contents.js:28
_next @ contents.js:50
Promise.then (async)
asyncGeneratorStep @ contents.js:38
_next @ contents.js:50
Promise.then (async)
asyncGeneratorStep @ contents.js:38
_next @ contents.js:50
(anonymous) @ contents.js:57
(anonymous) @ contents.js:46
ensureContent @ contents.js:491
_callee$ @ helpers.js:117
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ helpers.js:63
_next @ helpers.js:85
(anonymous) @ helpers.js:92
(anonymous) @ helpers.js:81
ensureFile @ helpers.js:142
_callee2$ @ useFile.js:352
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ useFile.js:147
_next @ useFile.js:169
(anonymous) @ useFile.js:176
(anonymous) @ useFile.js:165
(anonymous) @ useFile.js:649
commitHookEffectListMount @ react-dom.development.js:19607
commitPassiveHookEffects @ react-dom.development.js:19644
callCallback @ react-dom.development.js:189
invokeGuardedCallbackDev @ react-dom.development.js:238
invokeGuardedCallback @ react-dom.development.js:291
flushPassiveEffectsImpl @ react-dom.development.js:22708
unstable_runWithPriority @ scheduler.development.js:656
runWithPriority$1 @ react-dom.development.js:11076
flushPassiveEffects @ react-dom.development.js:22676
performSyncWorkOnRoot @ react-dom.development.js:21591
(anonymous) @ react-dom.development.js:11130
unstable_runWithPriority @ scheduler.development.js:656
runWithPriority$1 @ react-dom.development.js:11076
flushSyncCallbackQueueImpl @ react-dom.development.js:11125
flushSyncCallbackQueue @ react-dom.development.js:11113
scheduleUpdateOnFiber @ react-dom.development.js:21053
dispatchAction @ react-dom.development.js:15633
(anonymous) @ useStateReducer.js:113
(anonymous) @ useStateReducer.js:140
Promise.then (async)
(anonymous) @ useStateReducer.js:137
(anonymous) @ App.context.js:45
commitHookEffectListMount @ react-dom.development.js:19607
commitPassiveHookEffects @ react-dom.development.js:19644
callCallback @ react-dom.development.js:189
invokeGuardedCallbackDev @ react-dom.development.js:238
invokeGuardedCallback @ react-dom.development.js:291
flushPassiveEffectsImpl @ react-dom.development.js:22708
unstable_runWithPriority @ scheduler.development.js:656
runWithPriority$1 @ react-dom.development.js:11076
flushPassiveEffects @ react-dom.development.js:22676
performSyncWorkOnRoot @ react-dom.development.js:21591
(anonymous) @ react-dom.development.js:11130
unstable_runWithPriority @ scheduler.development.js:656
runWithPriority$1 @ react-dom.development.js:11076
flushSyncCallbackQueueImpl @ react-dom.development.js:11125
flushSyncCallbackQueue @ react-dom.development.js:11113
flushPassiveEffectsImpl @ react-dom.development.js:22735
unstable_runWithPriority @ scheduler.development.js:656
runWithPriority$1 @ react-dom.development.js:11076
flushPassiveEffects @ react-dom.development.js:22676
(anonymous) @ react-dom.development.js:22555
workLoop @ scheduler.development.js:600
flushWork @ scheduler.development.js:556
performWorkUntilDeadline @ scheduler.development.js:160
Show 136 more frames
contents.js:149 Uncaught (in promise) Error: Error creating file.
    at _callee$ (contents.js:149)
    at tryCatch (runtime.js:64)
    at Generator.invoke [as _invoke] (runtime.js:281)
    at Generator.throw (runtime.js:117)
    at asyncGeneratorStep (contents.js:28)
    at _throw (contents.js:54)
```

