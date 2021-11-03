# bugfix-cn-244-no-write-permissions

## 2020-07-28 

Modified `orgs.ts` to loop thru the orgs, get its teams, and then test current user membership.

If I find any non-read permissions, I set the permission attribute of org (also added) to 'write'. Otherwise, if none found, I set to 'read'.

Then in a callback in `Organization.js` I test the permission attribute for org prior to setting state. If it is 'read' I show an alert and do not execute `onOrganization(org)` in the callback.

This tested fine. But I have to modify the "CurrentUserOrganizatoins.md" demo page from:
```js
const config = {
  server: 'https://bg.door43.org',
  tokenid: 'PlaygroundTesting',
};
```
To this:
```js
const config = {
  server: 'https://git.door43.org',
  tokenid: 'PlaygroundTesting',
};
```
This is because, I have a read only test case in production that I can use to simulate the issue.

**Now to test locally**

1. use yalc to locally publish my changes to GRT: $ yalc publish
2. To use the RCL, in tc-create-app: $ yalc link gitea-react-toolkit

Note: to test against production, change roor ".env" file to point to `git.door43.org`.


## 2020-07-27 after 11am

At unblocking session, Birch noted that they could see the org. This means they had access to the org, but **not** write access. So his suggestion was that the folks who encountered the issue were set up in a team that only had read access. To test:
- Birch created a team and added me
- He noted that the default access was read-only
- I tested using production https://create.translationcore.com/
- In production, translate_test was the only org available to me
- I selected it, then clicked uta, then readme.md for the file
- Got the spinner!
Bingo - this exactly duplicates the problem I need to fix.

Suggested that perhaps we should check the org to see if we have write permissions before letting them continue. Alternatively, only show orgs to which I have write access (then change the message to state we don't have write-access to any orgs if the list is empty -- otherwise, just show the ones they do have write access to).

**Step 0.** looks over swagger apis to find one that will show me whether I have write access

This API shows whether I am a member:
curl -X GET "https://git.door43.org/api/v1/orgs/translate_test/members/cecil.new" -H "accept: application/json"
returns:
```sh
access-control-allow-headers: Authorization,DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range 
access-control-allow-methods: GET, POST, PUT, OPTIONS, PATCH, DELETE 
access-control-allow-origin: * 
access-control-expose-headers: Content-Length,Content-Range 
connection: keep-alive 
date: Mon, 27 Jul 2020 15:39:17 GMT 
server: nginx/1.16.1 
x-content-type-options: nosniff 
x-frame-options: SAMEORIGIN 
```
- Response code 204 means "is a member"; 404 means "not a member"
- Above was a 204 - so I'm a member; but no indication of access


**This API shows the teams of an org:**
It shows two teams: the owners team (with owner [all] permissions) and a "read_test" team with only "read" permissions.

Request URL: https://git.door43.org/api/v1/orgs/translate_test/teams?access_token=PlaygroundTesting

Results:
```json
[
  {
    "id": 159,
    "name": "Owners",
    "description": "",
    "organization": null,
    "permission": "owner",
    "units": [
      "repo.code",
      "repo.issues",
      "repo.pulls",
      "repo.releases",
      "repo.wiki",
      "repo.ext_wiki",
      "repo.ext_issues"
    ]
  },
  {
    "id": 172,
    "name": "read_test",
    "description": "this is a read access only team",
    "organization": null,
    "permission": "read",
    "units": [
      "repo.code",
      "repo.issues",
      "repo.pulls",
      "repo.releases",
      "repo.wiki",
      "repo.ext_wiki",
      "repo.ext_issues"
    ]
  }
]
```

**This API tests my membership in a team**
Test each team to see if I'm in it or not.

Request URL: https://git.door43.org/api/v1/teams/159/members/cecil.new?access_token=PlaygroundTesting

id '159' yields error "must be a member":
```json
{
  "message": "Must be a team member",
  "url": "https://git.door43.org/api/swagger"
}
```

Request URL: https://git.door43.org/api/v1/teams/172/members/cecil.new?access_token=PlaygroundTesting

id '172' yields:
```json
{
  "id": 0,
  "login": "cecil.new",
  "full_name": "",
  "email": "cecil.new@noreply.door43.org",
  "avatar_url": "https://git.door43.org/avatars/f8a86db8b875c63428a3d89dae47cf84",
  "language": "",
  "is_admin": false,
  "last_login": "0001-01-01T00:00:00Z",
  "created": "2019-09-19T13:37:42Z",
  "username": "cecil.new"
}
```

Thus I only have read access and that is not sufficient to use tC Create.


## 2020-07-27 up to 11am

I removed myself from the `translate_test` org and now I have read only access. Now if I change the `ensure` demo, it will fail since I only have read access.

Looks like the raw update and create bits are funneled thru `ensure`. Here is the console output:

```
createContent() error: Error: Request failed with status code 403
    at createError (createError.js:16)
    at settle (settle.js:17)
    at XMLHttpRequest.handleLoad (xhr.js:59)
react_devtools_backend.js:2273 console.trace
```

The trace goes up from the above createContent() function to ensureContent() function.

To replicate above **use the modified ensure demo code below**.

Trace:
- ensureContent in core/gitea-api/contents/contents.ts
    - ensureFile() in file/helpers.js
        - useFile.js in file


Here are the two closures in useFile.js that use ensureFile():
- two callbacks: load() and createFile()
- I don't see any error handling here... does it passthru?
```js
  const load = useCallback(async () => {
    if (config && repository && filepath) {
      const _file = await ensureFile({
        filepath, defaultContent, authentication, config, repository, branch,
      });
      const content = await getContentFromFile(_file);

      update({
        ..._file, branch, content, filepath: _file.path,
      });
    }
  }, [authentication, branch, config, defaultContent, filepath, repository, update]);

  const createFile = useCallback(async ({ branch: _branch, filepath: _filepath, defaultContent: _defaultContent }) => {
    if (config && repository) {
      const _file = await ensureFile({
        authentication, config, repository,
        branch: _branch,
        filepath: _filepath,
        defaultContent: _defaultContent,
      });

      if (_file) {
        repositoryActions.updateBranch(_branch);
        onFilepath(_filepath);
      };
    }
  }, [authentication, config, repository]);
```

Note: createFile is used as the onSubmit() function in the File Picker form (line):
```js
  const components = {
    create: repository && (
      <FileForm
        branch={branch}
        defaultContent={defaultContent}
        onSubmit={createFile}
      />
    ),
    browse: repository && blobComponents.browse,
    fileCard: repository && file && (
      <FileCard
        authentication={authentication}
        repository={repository}
        file={{ ...file, ...actions }}
      />
    ),
  };
```
The above create component is conditionally called:
```js
  if (file) {
    component = components.fileCard;
  } else if (!filepath) {
    if (create) {
      component = components.create;
    } else {
      component = components.browse;
    }
  }

```

OTOH, load() is added to the actions and returned to File context.
```js
  const actions = {
    update,
    load,
    read,
    save,
    close,
    dangerouslyDelete,
  };
```

It is also used by the save() callback:
```js
  const save = useCallback(async (content) => {
    if (writeable) {
      await saveFile({
        authentication, repository, branch, file, content,
      });
      await load();
    }
  }, [writeable, authentication, repository, branch, file, load]);
```

and the useEffect:

```js
  useEffect(() => {
    const notLoaded = (!file && filepath && !deleted);
    const loadNew = (file && filepath && file.filepath !== filepath);

    if (notLoaded || loadNew) {
      load();
    }
  }, [deleted, filepath, load, file]);
```

When I search for useFile, it only shows in two spots:
- gitea-react-toolkit: the File.context.js
- tc-create-app: in TargetFile.context.js

Since target file is where this issue will be happening, let's go there...

Here is the code, starting at line 22 of TargetFile.context.js:
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
```

Looks like TargetFileContext does not use it other than to provide it:
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

  return (
    <TargetFileContext.Provider value={context}>
      {children}
    </TargetFileContext.Provider>
  );
};
```

So now I have trace usage of this target file context... 

It is used in exactly on place, namely, `Workspace.js`:
```js
          <TargetFileContextProvider>
            <Translatable />
          </TargetFileContextProvider>
```

So now to Translatable and there:
```js
  const { state: targetFile, actions: targetFileActions } = useContext(
    TargetFileContext
  );
```

Looks like actions is renamed 'targetFileActions'.

The save() is used in a useMemo():
```js
      if (sourceFile.filepath.match(/\.md$/)) {
        let translatableProps = {
          original: sourceFile.content,
          translation: targetFile.content,
          onTranslation: targetFileActions.save,
        };
        _translatable = <MarkDownTranslatable {...translatableProps} />;
      } else if (sourceFile.filepath.match(/\.tsv$/)) {
        _translatable = <TranslatableTSV />;
      }
    }
    return _translatable;
  }, [filepath, sourceFile, targetFile, targetFileActions.save]);
```

1. Looks like there is a component named `TranslatableTSV` that might use the target file context as well. It isn't passed as a property however.
2. The save() is passed along as a property to `MarkDownTranslatable`

**Will have to branch now. Using level 3 to track...**

### MarkDownTranslatable

This is its own repo... so clone, yarn install, etc.

From above, the save() is passed as the "onTranslation" property. In `src/components/translatable/Translatable.js`:
```js
  const saveTranslation = useCallback(() => {
    onTranslation(editedTranslation);
  }, [onTranslation, editedTranslation]);
```

This callback is passed as a property named "onSave" to the Actions component and returned by Translatable:
```js
  return (
    <div className={classes.root}>
      <Paper>
        <Actions
          sectionable={sectionable}
          onSectionable={setSectionable}
          blockable={blockable}
          onBlockable={setBlockable}
          preview={preview}
          onPreview={setPreview}
          changed={changed}
          onSave={saveTranslation}
        />
      </Paper>
      {component}
    </div>
  );
```

So on to Actions... here we find at the beginning:

```js
  const _onSave = useCallback(onSave, [onSave]);
```

Not sure what this doing other than caching/renaming it?? But now need to look for both `_onSave` and `onSave`.

`_onSave` is used in a useMemo() to create a "saveAction()". It is used as the "onClick" response to the "save" button:

```js
  const saveAction = useMemo(() => (
    <Tooltip title={localString("Save")} arrow>
      <IconButton className={classes.action} aria-label="Save" disabled={!changed} onClick={_onSave}>
        {saveIcon}
      </IconButton>
    </Tooltip>
  ), [_onSave, changed, classes.action, saveIcon]);
```

Then finally it is rendered:
```js
  return (
    <div className={classes.actions}>
      {sectionsAction}
      {blocksAction}
      {previewAction}
      {saveAction}
    </div>
  );
```


### TranslatableTSV





## 2020-07-24

Began work by setting up project folder:
- cloned: tc-create-app, gitea-react-toolkit, and my journals repo (which will have this file).
- created vs code workspace for the above

In GRT's repositories demo page for repositories, here is new code that points to a test org I can use:

```js
import { Paper } from '@material-ui/core';
import { Repositories } from 'gitea-react-toolkit';

<Paper>
  <Repositories
    urls={[
      'https://git.door43.org/api/v1/repos/translate_test/en_ta',
      'https://git.door43.org/api/v1/repos/translate_test/ru_ta',
      'https://git.door43.org/api/v1/repos/translate_test/en_tn',
      'https://git.door43.org/api/v1/repos/translate_test/ru_tn',
    ]}
    onRepository={(data) => {
      alert(JSON.stringify(data, null, 2));
    }}
  />
</Paper>
```

And on `ensure` demo:
```js
import { Core, ensureContent } from 'gitea-react-toolkit';

const props = {
  config: {
    server: 'https://git.door43.org',
    tokenid: 'PlaygroundTesting',
  },
  owner: 'translate_test',
  repo: 'ru_ta',
  branch: 'testing',
  filepath: 'README-remove-this.md',
  content: 'Testing ensureContent',
  message: 'Testing ensureContent via Gitea-React-Toolkit',
  author: {
    email: "cecil.new@gmail.com",
    username: "cecil.new",
  },
};

<Core
  props={props}
  promise={ensureContent}
  authenticate
/>
```

