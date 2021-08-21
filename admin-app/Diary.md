# Diary

## 2021-03-11 (PM)

Complete card refresh on a repo rename action. Removed lots of code from `dcsApis.js`.

*This about completes the PoC!!*

## 2021-03-11 (AM)

Did some work on creating a manifest. Essential swagger is:
`https://qa.door43.org/api/v1/swagger#/repository/repoCreateFile`

Sample URL and JSON Body:
url: `https://qa.door43.org/api/v1/repos/translate_test/en_tw/contents/manifest.yaml?access_token=<TOKEN>`

body:

```json
{
      "author": {
        "email": "info@unfoldingword.org",
        "name": "unfoldingWord"
      },
      "branch": "master",
      "committer": {
        "email": "info@unfoldingword.org",
        "name": "unfoldingWord"
      },
      "content": "dGVzdGluZyAxIDIgMw==",
      "dates": {
        "author": "2021-03-11T14:02:16.158Z",
        "committer": "2021-03-11T14:02:16.158Z"
      },
      "message": "Initialize Manifest - must be updated",
      "new_branch": "master"
}
```
	
Note the content must be encode in base64. The key code snippet is:
`  const content = base64.encode(utf8.encode(manifests.ta_manifest));`

Normal response is status 201

```json
{
  "content": {
    "name": "manifest.yaml",
    "path": "manifest.yaml",
    "sha": "d0957c1537a95debd049dc13abf5e9f5302e4441",
    "type": "file",
    "size": 13,
    "encoding": "base64",
    "content": "dGVzdGluZyAxIDIgMw==",
    "target": null,
    "url": "https://qa.door43.org/api/v1/repos/translate_test/en_tw/contents/manifest.yaml?ref=master",
    "html_url": "https://qa.door43.org/translate_test/en_tw/src/branch/master/manifest.yaml",
    "git_url": "https://qa.door43.org/api/v1/repos/translate_test/en_tw/git/blobs/d0957c1537a95debd049dc13abf5e9f5302e4441",
    "download_url": "https://qa.door43.org/translate_test/en_tw/raw/branch/master/manifest.yaml",
    "submodule_git_url": null,
    "_links": {
      "self": "https://qa.door43.org/api/v1/repos/translate_test/en_tw/contents/manifest.yaml?ref=master",
      "git": "https://qa.door43.org/api/v1/repos/translate_test/en_tw/git/blobs/d0957c1537a95debd049dc13abf5e9f5302e4441",
      "html": "https://qa.door43.org/translate_test/en_tw/src/branch/master/manifest.yaml"
    }
  },
  "commit": {
    "url": "https://qa.door43.org/api/v1/repos/translate_test/en_tw/git/commits/38b0eb0971c28611e39f401137adb19e3efa2c4f",
    "sha": "38b0eb0971c28611e39f401137adb19e3efa2c4f",
    "created": "0001-01-01T00:00:00Z",
    "html_url": "https://qa.door43.org/translate_test/en_tw/commit/38b0eb0971c28611e39f401137adb19e3efa2c4f",
    "author": {
      "name": "unfoldingWord",
      "email": "info@unfoldingword.org",
      "date": "2021-03-11T14:02:16Z"
    },
    "committer": {
      "name": "unfoldingWord",
      "email": "info@unfoldingword.org",
      "date": "2021-03-11T14:02:16Z"
    },
    "parents": [
      {
        "url": "https://qa.door43.org/api/v1/repos/translate_test/en_tw/git/commits/c42a8a51df79dfd4ef720ab4bcaf5456b5a2c994",
        "sha": "c42a8a51df79dfd4ef720ab4bcaf5456b5a2c994",
        "created": "0001-01-01T00:00:00Z"
      }
    ],
    "message": "Initialize Manifest - must be updated\n",
    "tree": {
      "url": "https://qa.door43.org/api/v1/repos/translate_test/en_tw/git/trees/bf26c1129e796487961cf0ebe1741a5871fe2e5b",
      "sha": "bf26c1129e796487961cf0ebe1741a5871fe2e5b",
      "created": "0001-01-01T00:00:00Z"
    }
  },
  "verification": {
    "verified": false,
    "reason": "gpg.error.not_signed_commit",
    "signature": "",
    "signer": null,
    "payload": ""
  }
}
```


## 2021-03-10

**Completed**:
1. De-forking from create-app (now I can have my own issues and PRs)
2. Fixed behavior of Create Repo button so that it forces refresh of "card"
3. Added custom APIs to dcsApis.js that do exactly and only what I want.

**To Do**:
1. Trim dcsApis.js to only what I'm using.
2. Remove all create-app code that isn't being used.
3. Fix behavior of Rename Repo button so that forces refresh of "card"
4. Plan how to handle missing manifests.

## 2021-03-09

First up this morning... fix the problem with the server side references. See `docs/Runtime Detection of Dev vs Prd` for technique used by tc-create.

In that document, is this code:
```js
export const SERVER_URL = process.env.REACT_APP_DOOR43_SERVER_URL;
```

NOTE! https://nextjs.org/docs/migrating/from-create-react-app#environment-variables, wherein we learn:
```
Next.js has support for `.env` [Environment Variables](https://nextjs.org/docs/basic-features/environment-variables) similar to Create React App. The main different is the prefix used to expose environment variables on the client-side.

-   Change all environment variables with the `REACT_APP_` prefix to `NEXT_PUBLIC_`.
-   Server-side environment variables will be available at build-time and in [API Routes](https://nextjs.org/docs/api-routes/introduction).
```

So added a ".env" file, removed entry for it in ".gitignore", removed all references to hard coded values.

Second, fixed spacing (see styling added to Paper in `RepoHealthCheck.js`).

Third, added dismissal of info "Alert" message on success/failure of create or rename.


https://qa.door43.org/api/v1/repos/translate_test/en_tn 
returns:
```json
{"id":56744,"owner":{"id":24384,"login":"translate_test","full_name":"","email":"","avatar_url":"https://qa.door43.org/user/avatar/translate_test/-1","language":"","is_admin":false,"last_login":"1970-01-01T00:00:00Z","created":"2020-06-18T17:22:49Z","repo_languages":null,"username":"translate_test"},"name":"en_tn","full_name":"translate_test/en_tn","description":"","empty":false,"private":false,"fork":false,"template":false,"parent":null,"mirror":false,"size":12873,"html_url":"https://qa.door43.org/translate_test/en_tn","ssh_url":"git@qa.door43.org:translate_test/en_tn.git","clone_url":"https://qa.door43.org/translate_test/en_tn.git","original_url":"","website":"","stars_count":0,"forks_count":0,"watchers_count":12,"open_issues_count":0,"open_pr_counter":0,"release_counter":0,"default_branch":"master","archived":false,"created_at":"2020-06-22T13:40:58Z","updated_at":"2021-03-05T22:36:33Z","permissions":{"admin":true,"push":true,"pull":true},"has_issues":true,"internal_tracker":{"enable_time_tracker":true,"allow_only_contributors_to_track_time":true,"enable_issue_dependencies":true},"has_wiki":true,"has_pull_requests":true,"has_projects":false,"ignore_whitespace_conflicts":false,"allow_merge_commits":true,"allow_rebase":true,"allow_rebase_explicit":true,"allow_squash_merge":true,"avatar_url":"","internal":false,"language":"","subject":"","books":null,"title":"","checking_level":"","catalog":null}
```



## 2021-03-08

The demo is working OK, but because I have the tokenid hard-coded, it was broke this morning. There are two more requirements that this PoC need to take on. They are:
- Prove that we can programmatically create a new repo with the correct manifest
- Prove that we can determine if TQ repo is old markdown or new TSV

On the first one, I've asked Rich about it [here](https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE/topic/tc-create-app.23713/near/229308290)

On the second one, I've asked Robert about it [here](https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE/topic/TSV.20formats/near/229309849)

Here are a few things I'd like to do to make the app more resilient.
- Use the token for the authenticated user (instead of hard-coding one generated on QA, which is lost every weekend)
- Fix the colors/styles
- Fix spacing problems. Perhaps enclose the repo health check component in a card?

## 2021-03-04

To rename (example)
```sh
curl -X PATCH "https://qa.door43.org/api/v1/repos/richqa/old_name?access_token=<TOKEN>" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{  \"name\": \"new_name\"}"
```
Use "patch" method with body containing new name.



Using https://qa.door43.org/api/v1/swagger#/organization/createOrgRepo with org = `translate_test`:

```json
{
  "auto_init": true,
  "default_branch": "master",
  "description": "Init New Repo by Admin App",
  "gitignores": "macOS",
  "issue_labels": "",
  "license": "CC-BY-SA-4.0.md",
  "name": "ab_tn",
  "private": false,
  "readme": "",
  "template": true,
  "trust_model": "default"
}
```

This returned a 401 "Error: Unauthorized" and response body of:
```json
{
  "message": "token is required",
  "url": "https://qa.door43.org/api/swagger"
}
```

Generate a Token at:
https://qa.door43.org/user/settings/applications
![[Pasted image 20210304111252.png]]

Try again... with token this time... Worked!

```sh
curl -X POST "https://qa.door43.org/api/v1/orgs/translate_test/repos?token=...elided..." -H  "accept: application/json" -H  "Content-Type: application/json" -d "{  \"auto_init\": true,  \"default_branch\": \"master\",  \"description\": \"Init New Repo by Admin App\",  \"gitignores\": \"macOS\",  \"issue_labels\": \"\",  \"license\": \"CC-BY-SA-4.0.md\",  \"name\": \"ab_tn\",  \"private\": false,  \"readme\": \"\",  \"template\": true,  \"trust_model\": \"default\"}"
```

URL:
```sh
https://qa.door43.org/api/v1/orgs/translate_test/repos?token=...elided...
```

Response Body:
```json
{
  "id": 62983,
  "owner": {
    "id": 24384,
    "login": "translate_test",
    "full_name": "",
    "email": "",
    "avatar_url": "https://qa.door43.org/user/avatar/translate_test/-1",
    "language": "",
    "is_admin": false,
    "last_login": "1970-01-01T00:00:00Z",
    "created": "2020-06-18T17:22:49Z",
    "repo_languages": null,
    "username": "translate_test"
  },
  "name": "ab_tn",
  "full_name": "translate_test/ab_tn",
  "description": "Init New Repo by Admin App",
  "empty": false,
  "private": false,
  "fork": false,
  "template": true,
  "parent": null,
  "mirror": false,
  "size": 0,
  "html_url": "https://qa.door43.org/translate_test/ab_tn",
  "ssh_url": "git@qa.door43.org:translate_test/ab_tn.git",
  "clone_url": "https://qa.door43.org/translate_test/ab_tn.git",
  "original_url": "",
  "website": "",
  "stars_count": 0,
  "forks_count": 0,
  "watchers_count": 0,
  "open_issues_count": 0,
  "open_pr_counter": 0,
  "release_counter": 0,
  "default_branch": "master",
  "archived": false,
  "created_at": "2021-03-04T16:14:24Z",
  "updated_at": "2021-03-04T16:14:25Z",
  "permissions": {
    "admin": true,
    "push": true,
    "pull": true
  },
  "has_issues": true,
  "internal_tracker": {
    "enable_time_tracker": true,
    "allow_only_contributors_to_track_time": true,
    "enable_issue_dependencies": true
  },
  "has_wiki": true,
  "has_pull_requests": true,
  "has_projects": false,
  "ignore_whitespace_conflicts": false,
  "allow_merge_commits": true,
  "allow_rebase": true,
  "allow_rebase_explicit": true,
  "allow_squash_merge": true,
  "avatar_url": "",
  "internal": false,
  "language": "",
  "subject": "",
  "books": null,
  "title": "",
  "checking_level": "",
  "catalog": null
}
```

Response headers:
```
 access-control-allow-headers: Authorization,DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range 
 access-control-allow-methods: GETPOSTPUTOPTIONSPATCHDELETE 
 access-control-allow-origin: * 
 access-control-expose-headers: Content-Length,Content-Range 
 connection: keep-alive 
 content-length: 1438 
 content-type: application/json; charset=UTF-8 
 date: Thu04 Mar 2021 16:14:25 GMT 
 server: nginx/1.16.1 
 x-content-type-options: nosniff 
 x-frame-options: SAMEORIGIN 
 ```
 
 

## 2021-03-03

```json
{
  "auto_init": true,
  "default_branch": "master",
  "description": "Init New Repo by Admin App",
  "gitignores": "macOS",
  "issue_labels": "",
  "license": "CC-BY-SA-4.0.md",
  "name": "testrepo",
  "private": false,
  "readme": "",
  "template": true,
  "trust_model": "default"
}
```


```sh
curl -X POST "https://qa.door43.org/api/v1/user/repos" -H  "accept: application/json" -H  "authorization: Basic Y2VjaWwubmV3OjM1XkhoRF5IRiRkKiNOKkE=" -H  "Content-Type: application/json" -d "{  \"auto_init\": true,  \"default_branch\": \"master\",  \"description\": \"Init New Repo by Admin App\",  \"gitignores\": \"macOS\",  \"issue_labels\": \"\",  \"license\": \"CC-BY-SA-4.0.md\",  \"name\": \"kn_gst\",  \"private\": false,  \"readme\": \"\",  \"template\": true,  \"trust_model\": \"default\"}"
```


```json
```

So this worked. It created a README with just the one line of description provided. No license was created.

## 2021-02-16

Below is a list of changes made to today:

Rename the app to "translationCore: Admin". Line 11 in `Layout.js`

In `WorkspaceContainer.js`, I removed all the default cards and added a fake one. For the fake card, I imported `Card` from `translation-helps-rcl`. I'll have copy this code into this code since the card expects the typical translation sorts of content.

The component `TranslationSettings.js` has the needed org and language.
These are stored in Local Storage, not in indexedDB.

