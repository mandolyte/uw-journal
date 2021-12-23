# Add OBS

Here is content from issue [#4](https://github.com/unfoldingWord/gateway-admin/issues/4)

User wants to start the OBS package. The following will be done for OBS:

-   Check OBS existence
-   Check TN existence (obs-tn), plus TA from support references in the TN
-   Check TWL existence (obs-twl), plus TW from the resource links in the TWL
-   Check TQ existence (obs-tq)
-   Check SN existence (obs-sn)
-   Check SQ existence (obs-sq)

Otherwise, this issue is similar to [#3](https://github.com/unfoldingWord/gateway-admin/issues/3)

## To Do List

1. Add "Open Bible Stories" to the Add Book selection list
2. Add hooks for each of the above resources, eight in all
3. Add the hooks into the Admin Context
4. Since OBS will be a different set of rows to display than a normal bible book, then some changes are needed for the card that are OBS specific.


# Diary

## 2021-12-17

My plan is do this from the bottom up... thus first, let's make the hook.

Starting with `useObsRepoValidation`. I'll copy the `useLt...` as the starter...

My branch will be named: feature-cn-4-add-obs-resources.

Added into AdminContext:
```js
  const {
    state: {
      obsRepoTree,
      obsRepoTreeManifest,
      obsRepoTreeErrorMessage,
    },
  } = useObsRepoValidation({authentication, owner, server, languageId, refresh, setRefresh});
```
And then added the state to the context state variable.
```js
  // create the value for the context provider
  const context = {
    state: {
      obsRepoTree,
      obsRepoTreeManifest,
      obsRepoTreeErrorMessage,
      tnRepoTree,
      tnRepoTreeManifest,
... etc. 
```

In `common/BooksOfTheBible.js`, added this:
```js
export const BIBLE_AND_OBS = {
  ...ALL_BIBLE_BOOKS, 
  ...{obs: 'Open Bible Stories'},
}
```

and used it in this function:
```js
export function bookSelectList() {
  return Object.keys(BIBLE_AND_OBS).map(
    (bookId) => {
      return { id: bookId, name: BIBLE_AND_OBS[bookId] }
    }
  )
}
```

