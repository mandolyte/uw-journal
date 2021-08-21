# tc-create-app.829

This issue is a spike to research how to integrate Catalog Next (CN) as the source side for tc-create-app.

See [[Catalog Next]] for a design for a CN component.

This document is to show how to integrate the CN component into tc create.


# Testing

Below are six different use cases to be tested. Altho, since Editor Case A isn't possible, there are actually only five.

In addition to the content shown in the source and target sides, there is also the functionality of the file header chips that must work properly. Both source and target are shown here:

![[Pasted image 20210526080834.png]]

The functions are:
- view the license for the repo
- view the DCS webpage for the item
- source: compare vs master
- target: compare vs branch


## Editor Role

This is when unfoldingWord English language resources are being edited.

### Case A: New Item

In this case, an entirely new item is being created. This case is not possible with tc-create. New items must be created using other tools and added to the repository using Gitea DCS itself.

### Case B: Begin Update of Existing Item

This is the common case. It is recognizable when the source and target repos are identical. At present, the source is *always* a repo owned by `unfoldingword`. The following occurs:
- A user branch is created (if it doesn't already exist)
- The current item in the master branch is copied to the user branch, which becomes the starting point for the edits.
- Note that if a second party is also editing, saving, merging the same item, then it is possible that the source side (left) may different from time to time while the first is editing. There is currently no mechanism in place that will notify the first user that this has occurred. Teams must rely on communication to coordinate changes.

### Case C: Continue Update of an Existing Item

In this case, the edits may take several sessions/days/weeks. When the user returns to the item to continue editing:
- No user branch is created, since it already exists
- No content is copied from the master branch to the user branch, since the user edits are being captured there for later merging into the master branch

## Translator Role

This is when other organizations create content, translating from unfoldingWord owned and published resources. The role of translator is recognized by selection of a non-uW organization.

### Case A: New Item

In this case, an entirely new item is being created. The new item, of course, is only new to the translating organization. It will be selected from the `unfoldingword` organization for translation. The source content is currently taken from the latest production uW content. The steps are:
- The source is selected from the uW master branch file list. *There is an edge case where uW might have content in the master branch that is not yet released. This will cause an error when the published version of the content is retrieved.*
- The content for the source is retrieved from the latest published content.
- Note that since a new version of the content may be published while the translator is working, there is at present no mechanism in place to detect and notify the translator that the source from which they are translating may have changed (it may or may not -- it is only a possibility).
- A user branch is created in the translating org's repo and the default content from the uW published content is created there by default. That is, the English content is the starting point for the translation.
- At the end, the source side contains latest uW published content and the target side is identical.
- Eventually, once the translation is done, it will be merged into the org's repo's master branch.

### Case B: Begin Update of Existing Translated Item

In this case, translated content already exists and this will be an update to the existing content.
- A user branch is created in the repo
- The content of the repo's master branch will be used as the starting point will be copied to the user's branch.
- At the end, the source side will have the latest published content from uW and the target side will have the org's repo's master branch content.

### Case C: Continue Update of an Existing Translated Item

In this case, the edits may take several sessions/days/weeks. When the user returns to the item to continue editing:
- No user branch is created, since it already exists
- No content is copied from the master branch to the user branch, since the user edits are being captured there for later merging back into the master branch


# Implementation

## 2021-05-11

Below is the GRT "ensureContent()" function. If no content is found, it will provide content based on the provided parameter "content".

```js
export const ensureContent = async ({
  config, owner, repo, branch, filepath, content, message, author, onOpenValidation,
}: ModifyContentOptions): Promise<ContentObject> => {
  let contentObject: ContentObject;

  try { // try to read the file
    // NOTE: when a source file is fetched for translation, the following readConent
    // should always succeed since the file was selected from a UI which
    // is showing files that exist.
    //
    // OTOH, if the file is the target this will return null (the first time), 
    // throwing the error. When the error is thrown, the catch will fire.
    contentObject = await readContent({
      owner, repo, ref: branch, filepath, config,
    });
    if (!contentObject) throw new Error('File does not exist in branch');
    //
    // add on open validation checks here for source side or existing branch content
    //
    const _content:string = await getContentFromFile(contentObject);
    let notices: string[] = [];
    if ( onOpenValidation ) {
      notices = onOpenValidation(filepath, _content,contentObject.html_url);
    }
  } catch {
    try { // try to update the file in case it is in the default branch
      // NOTE: if the file is in the master branch of the target
      // the following readConcent will succeed
      // Otherwise it returns null; if null then the getContentFromFile
      // will throw an error (probably from trying to decode null or
      // if by url, a 404 is returned and get throws an error)
      // In this case, the catch() will create the content from the source
      // the "updateContent" will cause the existing target content in master
      // to be used. createContent will be called at some point during the update
      const _contentObject = await readContent({
        owner, repo, filepath, config,
      });
      //
      // add on open validation checks here for when target repo has data, but there is no user branch
      //
      // the below can throw an error, so it will go to the catch for create to be done
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

  return contentObject;
};
```

The above is only called in one place: 
`src\components\file\helpers.js` in the ensureFile() function. It is the ensureFile() function which provides ensureContent() with the "default content":
```js
  const file = await ensureContent({
    config: _config, owner, repo, branch, filepath,
    content: defaultContent, message: _message, author,
    onOpenValidation,
  });
```

The default content is a parameter passed to ensureFile(), which is called in two places in the useFile() hook.

The default content is passed to the useFile() hook:
```js
function useFile({
  authentication,
  repository,
  filepath,
  onFilepath,
  defaultContent,
  config: _config,
  create=false,
  onOpenValidation,
}) {
```

The useFile() hook is called from two places:
- in `tc-create-app` TargetFile.context.js
- in `gitea-react-toolkit` File.Context.js

In the target file context case (first one above), the default content is acquired from the "source" file context:
```js
  const { state: sourceFile } = useContext(FileContext); // line 18
```
using this parameter to the useFile() hook:
```js
    defaultContent: (sourceFile && sourceFile.content), // line 28
```
*NOTE! this means that default content comes solely via the Source File Context, which is created in App.js.*

In File Context itself, the default content is passed as an parameter, which in turn is passed down to the useFile hook.
```js
export function FileContextProvider({
  config: _config,
  authentication: _authentication,
  repository: _repository,
  filepath,
  onFilepath,
  defaultContent,
  create,
  onOpenValidation,
  children,
})
```
The useFile hook:
```js
  const {
    state, actions, component, components, config,
  } = useFile({
    config: _config || contextConfig,
    authentication: _authentication || contextAuthentication,
    repository: _repository || contextRepository,
    filepath, onFilepath, defaultContent, create, onOpenValidation,
  });
```

So how does FileContext come by the "default content"? In App.js:
```js
              <FileContextProvider
                authentication={authentication}
                repository={sourceRepository}
                filepath={filepath}
                onFilepath={setFilepath}
              >
```
Note that default content is NOT passed in.





# Research

Sizes of Zips:
![[Pasted image 20210507095702.png]]

## Source FileContext

The source file context is used in `App.js`:

```js
              <FileContextProvider
                authentication={authentication}
                repository={sourceRepository}
                filepath={filepath}
                onFilepath={setFilepath}
                onOpenValidation={_onOpenValidation}
              >
```
This is nested within: Repository Context, Organization Context, and Authentication Context.

Note that language comes later in the "stepper"; this confirms that this file context is about the uW English language resources. For example, in the file selection step, which is initiated by this context, the list of files show are from 
- owner: unfoldingWord
- language: English
- resource type: *user selected*

This maps to a specific repository on DCS. For example: 
`https://git.door43.org/unfoldingword/en_tn`.

The `<FileContextProvider/>` itself is in the `gitea-react-toolkit`.


## Application Stepper

In this file `src\components\application-stepper\ApplicationStepper.js`, is this line 31:
```js
 const { state: sourceFile, component: fileComponent } = useContext(FileContext);
```
This line provides a UI component named here as "fileComponent". Whereas there are a lot things in the object returned useContext() here, the destructuring only picks out two, namely:
sourcefile and fileComponent.

The steps are defined as:
```js
  const steps = [
    {
      label: 'Login',
      instructions: 'Login to Door43',
      component: () => (authenticationComponent),
    },
    {
      label: 'Organization',
      instructions: 'Select Your Organization',
      component: () => (organizationComponent),
    },
    {
      label: 'Resource',
      instructions: 'Select Resource to Translate',
      component: () => (repositoryComponent),
    },
    {
      label: 'Language',
      instructions: 'Select Your Language',
      component: () => (
        <LanguageSelect
          language={language}
          onLanguage={setLanguage}
        />
      ),
    },
    {
      label: 'File',
      instructions: 'Select File to Translate',
      component: () => (fileComponent),
    },
  ];
```

Note the last one... it uses the `fileComponent` as the UI to show.

Then in returned element, this where the component is used on the File step:
```js
                <Divider className={classes.divider} />
                {steps[activeStep].component()}
                <Divider className={classes.divider} />
```

So "fileComponent" comes from FileContext. So what is this "fileComponent"?

## gitea-react-toolkit

### File Context

This code is located at `src\components\file\File.context.js`.

The context includes this hook:
```js
  const {
    state, actions, component, components, config,
  } = useFile({
    config: _config || contextConfig,
    authentication: _authentication || contextAuthentication,
    repository: _repository || contextRepository,
    filepath, onFilepath, defaultContent, create, onOpenValidation,
  });
```

And the context value property is set:
```js
  const context = {
    state,
    actions,
    component,
    components,
    config,
  };
```
And then assigned:
```js
  return (
    <FileContext.Provider value={context}>
      {children}
    </FileContext.Provider>
  );
};
```

From this then "fileComponent" is part of "component". On to useFile() to find it ...

### useFile

There is state at the beginning:
```js
  const [file, setFile] = useState();
```
And this "if" at the end:
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
Since `file` has no default value, it will be initially null or undefined. And since initially there will be no filepath yet, we get to either create or browse.

The components are defined from this:
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

Since the parameter is provided to useFile from file context and App.js does not specify it, then per the useFile args, create is set to the default 'false':
```js
function useFile({
  authentication,
  repository,
  filepath,
  onFilepath,
  defaultContent,
  config: _config,
  create=false,
  onOpenValidation,
})
```

So we finally determine that the component shown on the UI comes from "browse". Which is, given that the property "repository" is set:
`repository && blobComponents.browse` comes from "blobComponents"... 
So where is that??

It comes from another hook inside this useFile hook:
```js
  const {
    state: blobState, actions: blobActions, components: blobComponents,
  } = useBlob({
    blob, onBlob: setBlob, config, repository, filepath,
  });
```

### useBlob

This hook is here: `src\components\tree-blob\useBlob.js`.

The browse component is this:
```js
  const browse = useMemo(() => {
    return (tree || url) ? (
      <Tree
        tree={tree}
        url={url}
        config={config}
        selected={true}
        onBlob={update}
        comparer={tsvManifestFileComparer}
      />
    ) : (<></>);
  }, [tree, url, config, update, repository?.name]);
```

The url property comes from line 20:
```js
  const url = _url || (repository && repository.tree_url);
```
So it is the tree_url member of the repository object.

On to Tree...

### Tree

In this code we finally have a DCS API call (still indirect):
```js
  const updateTree = async () => {
    const __tree = await fetchTree({ url, config, comparer });
    setTree(__tree);
  };
```

This will return an array of objects and a map() is used to create all the UI elements (which are either a folder item (a `<TreeObject/>`) or a file (a `<BlobObject/>`)).

The function `fetchTree` is in `./helpers`:
```js
import { get } from '../../core';

export const fetchTree = async ({ url, config, comparer }) => {
  const _config = {
    cache: {
      maxAge: 1 * 2 * 1000, // 2 sec cache override
    },
    ...config,
  };
  const response = await get({ url, config: _config });
  let { tree } = response;

  if (comparer)
  {
    tree = tree.sort(comparer);
  }

  return tree;
};
```

Turns out that tree_url is added by `repoTreeUrl()` in:
`gitea-react-toolkit\src\core\gitea-api\repos\git\trees.ts`
```js
// /api/v1/repos/unfoldingWord/en_ta/git/trees/master
export const repoTreeUrl = ({
  full_name, branch = '', default_branch = '',
}) => {
  const url = path.join(apiPath, 'repos', (full_name || ''), 'git', 'trees', branch || default_branch || '');
  return url;
};
```



# Alternates

```
# this gets the repo for master
https://qa.door43.org/api/v1/repos/unfoldingword/en_twl

# this gets the v2 released repo
https://qa.door43.org/api/catalog/v5/entry/unfoldingword/en_twl/v2


# this gets the repo tree for master
https://qa.door43.org/api/v1/repos/unfoldingword/en_twl/git/trees/master

# this get the v2 released repo tree
https://qa.door43.org/api/v1/repos/unfoldingword/en_twl/git/trees/7a3ccb5d879a7c2e6853cecbc67261534de5373d

https://git.door43.org/unfoldingWord/en_twl/src/commit/7a3ccb5d879a7c2e6853cecbc67261534de5373d




# this returns the "id" of the all the releases of a repo
https://qa.door43.org/api/v1/repos/unfoldingword/en_twl/releases

v2 (a tag_name) has id 10762


# this gets a release by id
https://qa.door43.org/api/v1/repos/unfoldingword/en_twl/releases/10762


# this gets the zip
https://qa.door43.org/unfoldingWord/en_twl/archive/v2.zip
```
