# Content Validation App Journal

Original thinking of updating the book package app to include CV was discarded. Notes for that are in [[feature-cn-28-add-cv-to-bpa]].

# 2020-10-11

Next attempt... use two useEffects one for input to component and one on
the state variable.

# 2020-10-09


Note: updates for this components data come from two directions.
1. The first is from the argument 'results'. This data comes from when a 
language is selected in step 1.
2. The second are updates made by the user to either the org or repo values.

If the results are updated, they take precedence since it means the user
has selected a new language. New results will be cached.

The user updates are reflected in the cached data in indexedDB.

The data presented to the user, in either case comes from the 
state variable "data" and is updated by the function "setData".

When this function runs, here is the logic flow.
- First check the cache for prior results (repoLangResultsKey).
    - If no prior results, 
        - store original/latest results using key repoLangResultsKey
        - format the data as needed
        - store in cache using key repoValidationKey
        - update using setData()*
    - If prior results, stringify both and compare
        - if the same, the still on same language, so:
            - get data from cache
            - call setData()* with it
        - if not the same, then treat as above case of no prior results
- (*) In the above, wherever it says to use setData(), the following must be done
since data is a state variable and persists between steps.
    - get the current value of data and stringify it
    - take the value given to be set and stringify it
    - only use setData() if the values are different.
- In the above there are actually two keys in indexedDB:
    - original/latest language results key: this will only be updated if
    the results actually change
    - the reformatted array of data for the material table.
    - the key for the first is 'repoLangResultsKey'.
    - the key for the table cache is 'repoValidationKey'.


A code graveyard...

```js
    // this is going to be messy! Since components cannot be async, 
    // I cannot use the await keyword. Therefore, lots of nested 
    // promise stuff... there's probably a way to use a are React useEffect
    // but it isn't obvious to me... my apologies

    repoValidationStore.getItem(repoLangResultsKey).then( (value) => {
        if ( value ) {
            // this means we've had prior language results
            // have they changed?
            let _oldResults = JSON.stringify(value);
            let _newResults = JSON.stringify(results);
            if ( _oldResults === _newResults ) {
                // still on same language
                // In this case, the table cache takes precidence
                repoValidationStore.getItem(repoValidationKey)
                .then( (value) => {
                    // has data changed? if not do nothing
                    // can't use stringify directly because data in table has extra
                    // column(s) added by material table software.
                    if ( value && data ) {
                        let _oldResults = JSON.stringify(value);
                        let _newResults = JSON.stringify(data);
                        if ( _oldResults === _newResults ) {
                            // data and cache are same
                            // don't update table
                        } else {
                            console.log("Point A");
                            console.log("cache and table are not the same?! - table, cache:", _newResults, _oldResults);
                            setData(value);
                        }
                    }
                })
            } else {
                // OK, new language results, takes precidence!
                // Step 1. update the cached lang results
                repoValidationStore.setItem(repoLangResultsKey, results)
                .then( () => {
                    // Step 2. reformat the data and update table cache
                    let _data = [];
                    for (let i=0; i<results.length; i++) {
                        _data.push({
                            repoType: results[i].repoType,
                            lang: results[i].language,
                            org: results[i].username,
                            repo: results[i].repository,
                            message: results[i].message,
                        });
                    }
                    repoValidationStore.setItem(repoValidationKey,_data)
                    .then( () => {
                        // this is done unconditionally since we have 
                        // new language validation results
                        console.log("Point B");
                        setData(_data);
                    })
                })
            }
        } else {
            // no prior results, use passed in results
            repoValidationStore.setItem(repoLangResultsKey, results)
            .then( () => {
                let _data = [];
                for (let i=0; i<results.length; i++) {
                    _data.push({
                        repoType: results[i].repoType,
                        lang: results[i].language,
                        org: results[i].username,
                        repo: results[i].repository,
                        message: results[i].message,
                    });
                }
                repoValidationStore.setItem(repoValidationKey,_data)
                .then( () => {
                    // this is done unconditionally since we have 
                    // new language validation results
                    console.log("Point C");
                    setData(_data);
                })
            })
        }
    })


    let repoVisual = (
        <Paper>
        <MaterialTable
            icons={tableIcons}
            title="Repo Validation"
            columns={columns}
            data={data}
            options={ {sorting: true, pageSize: 7} }
            cellEditable={{
                onCellEditApproved: (newValue, oldValue, rowData, columnDef) => {
                    console.log("onCellEditApproved()", newValue, oldValue, rowData, columnDef);
                    return new Promise((resolve, reject) => {
                    let _data = data;
                    // first find the matching row
                    for (let i=0; i<_data.length; i++) {
                        if ( _data[i].repoType === rowData.repoType &&
                             _data[i].org      === rowData.org &&
                             _data[i].repo     === rowData.repo 
                        ) {
                            if ( columnDef.field === 'repo' ) {
                                _data[i].repo = newValue;
                            } else if ( columnDef.field === 'org' ) {
                                _data[i].org  = newValue;
                            }
                            getApi.setPathForRepo(_data[i].lang, _data[i].repoType, _data[i].org, _data[i].repo);
                            let errors = [];
                            getApi.verifyRepo(_data[i].org,_data[i].repo,errors,_data[i].repoType,_data[i].lang)
                            .then((errors) => {
                                console.log("getApi.verifyRepo() errors=",errors);
                                _data[i].message = errors[0].message;
                                repoValidationStore.setItem(repoValidationKey, _data).then( () => {
                                    console.log("repoValidationStore.setItem()");
                                });
                                console.log("Point D");
                                setData(_data);
                            });
                            break;
                        }
                    }
                    setTimeout(resolve, 1000);
                  });
                }
              }}      
        />
        </Paper>
    );

    return repoVisual;
```


```js
        console.log("useEffect() 1");
        repoValidationStore.removeItem(repoValidationKey).then(
            () => {
                console.log("removeItem():", repoValidationKey);
            }
        )
        console.log("useEffect() 2");
```

```js
                onCellEditApproved: (newValue, oldValue, rowData, columnDef) => {
                    console.log("onCellEditApproved()", newValue, oldValue, rowData, columnDef);
                    return new Promise((resolve, reject) => {
                    let _data = data;
                    // first find the matching row
                    for (let i=0; i<_data.length; i++) {
                        if ( _data[i].repoType === rowData.repoType &&
                             _data[i].org      === rowData.org &&
                             _data[i].repo     === rowData.repo 
                        ) {
                            if ( columnDef.field === 'repo' ) {
                                _data[i].repo = newValue;
                            } else if ( columnDef.field === 'org' ) {
                                _data[i].org  = newValue;
                            }
                            getApi.setPathForRepo(_data[i].lang, _data[i].repoType, _data[i].org, _data[i].repo);
                            let errors = [];
                            getApi.verifyRepo(_data[i].org,_data[i].repo,errors,_data[i].repoType,_data[i].lang)
                            .then((errors) => {
                                console.log("getApi.verifyRepo() errors=",errors);
                                _data[i].message = errors[0].message;
                                repoValidationStore.setItem(repoValidationKey, _data).then( () => {
                                    console.log("repoValidationStore.setItem()");
                                });
                                console.log("Point D");
                                setData(_data);
                            });
                            break;
                        }
                    }
```

```js
// caches repo validation results
const repoValidationStore = localforage.createInstance({
    driver: [localforage.INDEXEDDB],
    name: 'repo-validation-cache',
});
  
const repoValidationKey = 'repoValidationKey';
```

```
                                /*
                                repoValidationStore.setItem(repoValidationKey, _data).then( () => {
                                    console.log("repoValidationStore.setItem()");
                                });
                                console.log("Point D");
                                */
```

# 2020-09-29

Here is master list of problems:

- Russian, ru_gl
1. The Russian org ru_gl does not have expected repos for ULT and UST. The expected repo names for these are "ru_glt" and "ru_gst", respectively. However, I do see two repos claiming to be these, namely: "ru_rlob" and "ru_rsob".
2. The Russaion tQ does not exist, namely "ru_tq". But I see what appears to be tQ with the name "ru_tq_2lv".
3. ~~It is looking for original languages in the GL repo. It should not, since the two are in the unfoldingword org and nowhere else.~~
4. There is no manifest for `ru_tn`

- Hindi, translationCore-Create-BCS
1. Notice that the Hindi tQ repo uses an uppercase "Q" in the name. Since URLs are case insensitive, there is no impact on web usage. But we are having misses on keys in the browser database. Could someone rename this please! Or alternatively, we could make all lookups lowercase.
2. ~~It is looking for original languages in the GL repo. It should not, since the two are in the unfoldingword org and nowhere else.~~
3. There is no bible content , i.e., no LT nor ST.

- Kannada, translationCore-Create-BCS
1. ~~It is looking for original languages in the GL repo. It should not, since the two are in the unfoldingword org and nowhere else.~~
2. Missing resources: tQ, LT, and ST.

- Latin American Spanish, ES-419_gl
1. ~~It is looking for original languages in the GL repo. It should not, since the two are in the unfoldingword org and nowhere else.~~
2. There is no bible content (no GST, GLT). But the org is full of repos that appear to have single books in them. For example, Titus. Note also that it appears to be USFM, but does NOT have the '.usfm' extension on the file name. Will these be combined at some point into their expected repos? (namely: es-419_gst or es-491_glt).
3. In org es-419_gl, there are is no tQ repository (es-419_tq), but there are lots of individual repos that contain questions for a single book. For example, 1 John. Note that data is actually JSON, not Markdown. This might be a good candidate to go straight to the TSV format.


After manifest and some other changes, redo timings:

First, cleared indexedDB and refreshed page.
- Before doing anything, the original languages are fetched. Then continuing as below on 25 Sep.
- First time: selected unfoldingword/English. Before any selections made, more repos are prefetched:
    - TA, TN, TW, TQ
    - selected JUD and 2PE; JUD took 29s and 2PE took 48s
- Second time: same selections; JUD took 15s and 2PE took 16s.
- Third time: added 1JN; JUD took 31s, 2PE took 33s, 1JN took 60s
- Fourth time: same selections; JUD took 31s, 2PE took 34s, 1JN took 35s
- Fifth time: added Titus; JUD took 37s, 2PE took 43s, 1JN took 53s, and Titus took 52s
- Sixth time: same selections; JUD took 36s, 2PE took 40s, 1JN took 41s, and Titus took 40s


# 2020-09-28

More added! this time to preload ULT and UST... redo timings:

First, cleared indexedDB and refreshed page.
- Before doing anything, the original languages are fetched. Then continuing as below on 25 Sep.
- First time: selected unfoldingword/English. Before any selections made, more repos are prefetched:
    - TA, TN, TW, TQ
    - selected JUD and 2PE; JUD took 20s and 2PE took 42s
- Second time: same selections; JUD took 11s and 2PE took 14s.
- Third time: added 1JN; JUD took 24s, 2PE took 28s, 1JN took 57s
- Fourth time: same selections; JUD took 23s, 2PE took 27s, 1JN took 29s
- Fifth time: added Titus; JUD took 27s, 2PE took 36s, 1JN took 49s, and Titus took 48s
- Sixth time: same selections; JUD took 26s, 2PE took 31s, 1JN took 33s, and Titus took 31s



Redid timings with the branch `feature-mcleanb-9/getFileOptimize`.

First, cleared indexedDB and refreshed page.
- Before doing anything, the original languages are fetched. Then continuing as below on 25 Sep.
- First time: selected unfoldingword/English. Before any selections made, more repos are prefetched:
    - TA, TN, TW, TQ
    - selected JUD and 2PE; JUD took 25s and 2PE took 43s
- Second time: same selections; JUD took 10s and 2PE took 12s.
- Third time: added 1JN; JUD took 20s, 2PE took 25s, 1JN took 52s
- Fourth time: same selections; JUD took 22s, 2PE took 25s, 1JN took 27s
- Fifth time: added Titus; JUD took 20s, 2PE took 25s, 1JN took 30s, and Titus took 43s
- Sixth time: same selections; JUD took 23s, 2PE took 28s, 1JN took 29s, and Titus took 28s

# 2020-09-25

Timings after some performance improvements.

- first time: selected unfoldingword/English, then JUD and 2PE; 2PE took 34s and JUD 25s.
- second time: same selections; both took 7s.
- third time: added one book 1JN;  both JUD and 2PE took 7s; 1JN took 38s.
- fourth time: same selections; all took about 17s.
- fifth time: added one book Titus; all took 16s except Titus which took 43s.

Tentative conclusions: the first time a book is selected it takes much longer, but subsequent times is much shorter.

So the question: What is happening the first time for *any* new book that makes it take so much longer?
Tentative answer: network fetch time.
Possible solution: fetch and persist all repos, not just TA, TW, and TQ