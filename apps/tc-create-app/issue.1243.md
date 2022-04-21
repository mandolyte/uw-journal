# Issue 1243

## basic issue description


## 2022-04-21

first, cleared user branch:
https://qa.door43.org/es-419_gl/es-419_tn/

second, logged in with tc-create:
- org es-419_gl
- resource TN
- lang es-419


In the repo, there exist these files in master:
en_tn_01-GEN.tsv
en_tn_02-EXO.tsv
en_tn_08-RUT.tsv
en_tn_56-2TI.tsv
en_tn_57-TIT.tsv
en_tn_65-3JN.tsv
es-419_tn_57-TIT.tsv

I need to pick a book that **isn't** in master yet.

So lets go with Jude...

*Result:* problem occurs, namely, endless spinner

Here are the messages in console log (trimmed down):
```
xhr.js:166          GET https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv?noCache=0.3698261406395298&ref=cecil.new-tc-create-1 404 (Not Found)
xhr.js:166          GET https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv?noCache=0.3486190488840297 404 (Not Found)
xhr.js:166          GET https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv?noCache=0.19908593899115035&ref=cecil.new-tc-create-1 404 (Not Found)
_callee3$ @ useFileContent.js:251
tryCatch @ runtime.js:64
xhr.js:166          POST https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv 404 (Not Found)
xhr.js:166          GET https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv?noCache=0.8337425881622968 404 (Not Found)
_callee3$ @ useFileContent.js:251
_next @ useFile.js:114
Promise.then (async)
asyncGeneratorStep @ useFile.js:102
_next @ useFile.js:114
Promise.then (async)
asyncGeneratorStep @ useFile.js:102
_next @ useFile.js:114
(anonymous) @ useFile.js:121
(anonymous) @ useFile.js:110
(anonymous) @ useFile.js:597
xhr.js:166          POST https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv 422 (Unprocessable Entity)
_callee3$ @ useFileContent.js:251
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ useFileContent.js:28
_next @ useFileContent.js:50
Promise.then (async)
asyncGeneratorStep @ useFileContent.js:38
_next @ useFileContent.js:50
Promise.then (async)
asyncGeneratorStep @ useFileContent.js:38
_next @ useFileContent.js:50
(anonymous) @ useFileContent.js:57
(anonymous) @ useFileContent.js:46
(anonymous) @ useFileContent.js:274
_next @ useFile.js:114
(anonymous) @ useFile.js:121
(anonymous) @ useFile.js:110
(anonymous) @ useFile.js:597
xhr.js:166          POST https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv 422 (Unprocessable Entity)
contents.js:149 
Uncaught (in promise) Error: Error creating file.
```

Outcome:
1. my user branch is created
2. the file for Jude is created, but it is empty

In gitea-react-toolkit...

updated `src\core\gitea-api\repos\contents\contents.ts`:
- the createContent() function, adding the actual error to the throw new Error:

```js
// POST /api/v1/repos/{owner}/{repo}/contents/{filepath}
export const createContent = async ({
  config, owner, repo, branch, filepath, content, message, author,
}: ModifyContentOptions): Promise<ContentObject> => {
  const url = Path.join(apiPath, 'repos', owner, repo, 'contents', filepath);
  let contentObject: ContentObject;

  try {
    // TODO: Check to see if branch exists to set branch or new_branch in payload
    try {
      const _payload = payload({
        branch, content, message, author,
      });
      const response = await post({
        url, payload: _payload, config,
      });
      contentObject = response.content;
    } catch {
      const _payload = payload({
        new_branch: branch, content, message, author,
      });
      const response = await post({
        url, payload: _payload, config,
      });
      contentObject = response.content;
    }
  } catch (error) {
    throw new Error('Error creating file. Error:\n'+error);
  };
  return contentObject;
};
```

- to entureContent(), added a try/catch at the end with some console logging:

```js
export const ensureContent = async ({
  config, owner, repo, branch, filepath, content, message, author, onOpenValidation,
}: ModifyContentOptions): Promise<ContentObject> => {
  let contentObject: ContentObject ;

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
      try {
        console.log("contents.ts/ensureContent() inner catch: config, owner, repo, branch, content:",
          config, owner, repo, branch, content
        );
        contentObject = await createContent({
          config, owner, repo, branch, filepath, content, message, author,
        });
      } catch (e) {
        console.log("ensureContent()/createContent() failed. Errors:", e);
      }
    };
  };

  return contentObject;
};
```

