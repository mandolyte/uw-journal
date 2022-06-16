# Save function trace

## tc-create-app
Note: this is post-refactor. All the below are in tc-create unless otherwise noted.

*Starting with Translatable.js*

These lines:
```js
      if (filepath.match(/\.md$/)) {
        let translatableProps = {
          original: sourceFileContent,
          translation: targetFileContent,
          onTranslation: saveTranslation,
          onContentIsDirty: setContentIsDirty,
        };
        console.log('Markdown file selected');
        _translatable = <MarkdownContextProvider><MarkDownTranslatable {...translatableProps} /></MarkdownContextProvider>;
      } else if (filepath.match(/\.tsv$/)) {
        console.log('TSV file selected');
        _translatable = <TranslatableTSV onSave={saveTranslation} onEdit={autoSaveOnEdit} onContentIsDirty={setContentIsDirty} />;
```
For both markdown and tsv files, the function `saveTranslation` is passed.

*Note that auto save is there for TSV but not for markdown!*

The `saveTranslation` function comes from a hook:
```js
  const {
    actions: {
      autoSaveOnEdit,
      saveTranslation,
    },
    component: authenticationDialog,
  } = useRetrySave();
```

From the hook we find (including the auto save too):
```js
  const saveTranslation = useDeepCompareCallback(async (content) => {
    setSavingTargetFileContent(content);

    try {
      await save(content);
    } catch (error) {
      const { isRecoverable } = parseError({ error });

      // assumption is that it is an authentication issue.
      if (isRecoverable) {
        openAuthenticationDialog();
      } else {
        setSaveFailed(true);
      };
    };
  }, [save]);

  const autoSaveOnEdit = useCallback(async (content) => {
    await saveCache(content);
  }, [saveCache]);
```

The save is done by a "save" function. This, and auto save, from the AppContext:
```js
  const {
    giteaReactToolkit: {
      authenticationHook,
      targetFileHook,
    },
  } = useContext(AppContext);

  const { save, saveCache } = targetFileHook.actions || {};
```

The above hints that the functionality lies in the Gitea React Toolkit (GRT).
These lines in `App.context.js` show where it comes from:
```js
  const giteaReactToolkit = useGiteaReactToolkit({ state, actions });

  const value = {
    state,
    actions,
    giteaReactToolkit,
  };

  return <AppContext.Provider value={value}>{children}</AppContext.Provider>;
};
```

So the GRT hook is returning a `targetFileHook` which returns the save function (and the auto save cache function).

In the GRT hook code, we find:
```js
  const targetFileHook = useFile({
    config,
    authentication,
    repository: targetRepository,
    filepath: (readyForTargetFile ? filepath : undefined),
    onFilepath: setFilepath,
    defaultContent,
    onOpenValidation: _onOpenValidation,
    onLoadCache: _onLoadCache,
    onSaveCache: _onSaveCache,
  });
```

The useFile hook is in the GRT package.

## gitea-react-toolkit

Picking up with the above useFile hook
at `gitea-react-toolkit\src\components\file\useFile.js`.

The save and saveCache are part of the set of actions returned by the hook:
```js
  const actions = {
    update,
    load,
    read,
    save,
    saveCache,
    onSaveCache,
    onLoadCache,
    close,
    dangerouslyDelete,
    setIsChanged,
    onConfirmClose,
  };
```

The save function is defined in this file as:
```js
  const save = useDeepCompareCallback(async (_content) => {
    await saveFile({ authentication, repository, branch, file, content: _content });
    // (save() will not happen for "OFFLINE" system files)
    await saveCache(); // Empty cache if user has saved this file
    await load();
    contentActions.reset();
  }, [writeable, authentication, repository, branch, file, load, saveFile, saveCache]);
```

Here we see a function named `saveFile` that is handling the actual save of the content to DCS.

That function is imported:
```js
import {
  saveFile, ensureFile, deleteFile, getContentFromFile,
} from './helpers';
```

In the helpers file, we have:
```js
export const saveFile = async ({
  authentication, repository, branch, file, content, message,
}) => {
  const {
    user: author, config, token: { name: tokenid },
  } = authentication;
  const { owner: { username: owner }, name: repo } = repository;
  const { path: filepath, sha } = file;
  const _message = message || `Edit '${filepath}' using '${tokenid}'`;

  const response = await updateContent({
    config, owner, repo, branch, filepath,
    content, message: _message, author, sha,
  });
  return response;
};
```

**Note! the code assumes the file exists. Notice it uses "updateContent" unconditionally.**

Thus here is where we need to do some checking. Since we defer branch creation until this point, here is where those changes will need to be made.

There is a function in this file named "ensureFile()"; possibly it could be used with the updated content provided as the "default". Here is the function signature:

```js
export const ensureFile = async ({
  config, authentication, repository, branch, filepath, defaultContent, message, onOpenValidation,
}) => { 
```

This function uses "ensureContent" to actually do the work.

The `ensureContent()` function is here:
`src\core\gitea-api\repos\contents\contents.ts`. Note it is written in Typescript.

This function has a lot of comments to help understand what it is doing. 
Let's step thru it.

First here is the function signature and first line of code:
```js
export const ensureContent = async ({
  config, owner, repo, branch, filepath, content, message, author, onOpenValidation,
}: ModifyContentOptions): Promise<ContentObject> => {
  let contentObject: ContentObject;
```

Note that a "ContentObject" is returned. This is an interface, defined here as:
```ts
interface ContentObject {
  path: string;
  sha: string;
  content: string;
  html_url: string;
}
```

Next is the outermost try block with comments removed:
```ts
  try { 
    contentObject = await readContent({
      owner, repo, ref: branch, filepath, config,
    });
    if (!contentObject) throw new Error('File does not exist in branch');

    const _content:string = await getContentFromFile(contentObject);
    let notices: string[] = [];
    if ( onOpenValidation ) {
      notices = onOpenValidation(filepath, _content,contentObject.html_url);
    }
  } catch {
```

It first attempts to read the content from the branch. Remember that this function is trying to save the file. It is trying to ensure that the branch and file *already* exist. And, if not, then put the "default content" there.

So if there is content, that means the user has already begun to edit the file and thus it will be returned. If an "onOpenValidation" function is passed it is used to validate that the file is reasonable and won't cause the app to crash.

If the read fails, then an error is thrown. So proceed to the catch and the inner try block (with comments removed and new ones inserted)
```ts
    try { 
		
	  // does the file exist in the repo's master branch?
	  // if so, then translation work has been done before
	  // and this is an update to it. An error is thrown if none
      const _contentObject = await readContent({
        owner, repo, filepath, config,
      });

	  // this function decodes it from base64
      const _content = await getContentFromFile(_contentObject);
      let notices: string[] = [];
      if ( onOpenValidation ) {
        notices = onOpenValidation(filepath, _content,_contentObject.html_url);
      }
      if ( notices.length === 0 ) {
        // only update if no notices
        contentObject = await updateContent({
          config, owner, repo, branch, filepath, content: _content, message, author, sha: _contentObject.sha,
        });
      } else {
        contentObject = _contentObject;
      }
    } catch { // try to create the file if it doesn't exist in default or new branch
      contentObject = await createContent({
        config, owner, repo, branch, filepath, content, message, author,
      });
    };
  };
```

Here the code attempts to read from the master branch of the repo. If this file has been worked on before, it will exist there. So it is read, decoded from base64, light validation performed (optional), then the file is updated in the branch.

If the file is not in the master branch, then this is first time this file has been worked on. And an error is thrown. The to the inner catch,  which reads:

```ts
    } catch { // try to create the file if it doesn't exist in default or new branch
      contentObject = await createContent({
        config, owner, repo, branch, filepath, content, message, author,
      });
    };
```

Finally, an attempt is made create the branch (if it doesn't exist) and the file.

The "createContent" function is in this file and it uses Gitea API POSTs to do the work. The function returns a ContentObject.

The "updateContent" function is also in this file and uses PUTs to update the file.
