# Journal

## Quick facts

Link to translation database:
https://td.unfoldingword.org/uw/languages/

Click "Data Sources", then Language Names JSON to pull entire dataset in JSON format. Direct link:
https://td.unfoldingword.org/exports/langnames.json

ISO tables:
https://iso639-3.sil.org/code_tables/download_tables


## Table of Contents

- Daily Journal
- Appendices
    - Generate the styleguide locally
    - Unpublishing
    - Creating a React+Material-UI app


## Daily Journal

### 2019-10-28

This repo is now transferred to `unfoldingWord-box3`. Today I have worked on getting to work again. Here are the things I did:

- Removed my copy of the repo and cloned the one from `unfoldingWord`.
- Ran "yarn" to (re) install all the dependencies (note these are *not* part of the repo, but are a slew of node modules)
- Bumped the rev in `package.json`
- Ran `yarn start` to verify things worked locally
- Committed and pushed code
- Ran yarn publish


### 2019-10-08

After meeting with Jesse, my first task is to do some refactoring, namely:
1. create two folders under `components`: one singular "language" and one plural "languages"
2. Under `language`, move the current "Languages" component and rename it `languageAttributes`.
3. Under `languages`, move the current "LanguageList" component

Next add a new "Language" component that simply returns the uW format. In chat, Jesse said to use: 
(am) Amharic – አማርኛ (Africa Gateway), which is:

- language code in parentheses
- the anglicized name
- the localized name
- the gateway region

*Note also the idea of being able to override the default uW display.*

Finally, the LanguageList component should be refactored to use the Language component. In other words, the LanguageList component would be made up of Language components.

### 2019-10-07

Today, I begin working an "edit" component. In many ways it will be simplar to the "languages" component. The requirements are captured in the DCS project `cecil.new/tD-DataRestructure`.

### 2019-10-04

Goal: add a new component that will create a list of languages. This should take some options, potentially:
- a format: 
    - "uw" for a unfoldingWord standard format (whatever that is)
    - "bcp47" for a standards compliant full language tag
- a filter: to be applied to the selected format
- a line limit: with a default of, say, 10

So something like this:
```
<LanguageList format="bcp47" filter="" limit=10 />
```

To test interactively:
```
$ yarn start
yarn run v1.17.3
$ styleguidist server
i ｢wds｣: Project is running at http://localhost:6060/
i ｢wds｣: webpack output is served from undefined
i ｢wds｣: Content not from webpack is served from C:\Users\mando\Projects\mandoly
te\uw-languages-rcl
You can now view your style guide in the browser:

  Local:            http://localhost:6060/
  On your network:  http://192.168.200.25:6060/

```

Looks good... now commit, push, and run `yarn publish`

### 2019-10-02

Adding commas to separate items in array of strings. In `langnames.js`, 
added join method:
```
      aliases: langname.alt.join(),
```


### 2019-09-28

Looking at testing: https://reactjs.org/docs/testing-recipes.html

```
$ yarn add @testing-library/react @testing-library/jest-dom
yarn add v1.17.3
[1/4] Resolving packages...
[2/4] Fetching packages...
info fsevents@1.2.9: The platform "win32" is incompatible with this module.
info "fsevents@1.2.9" is an optional dependency and failed compatibility check.
Excluding it from installation.
info fsevents@1.2.4: The platform "win32" is incompatible with this module.
info "fsevents@1.2.4" is an optional dependency and failed compatibility check.
Excluding it from installation.
[3/4] Linking dependencies...
warning "react-styleguidist > terser-webpack-plugin@1.4.1" has unmet peer depend
ency "webpack@^4.0.0".
warning "react-styleguidist > webpack-dev-server@3.8.1" has unmet peer dependenc
y "webpack@^4.0.0".
warning "react-styleguidist > webpack-dev-server > webpack-dev-middleware@3.7.1"
 has unmet peer dependency "webpack@^4.0.0".
[4/4] Building fresh packages...

success Saved lockfile.
success Saved 20 new dependencies.
info Direct dependencies
├─ @testing-library/jest-dom@4.1.0
└─ @testing-library/react@9.2.0
info All dependencies
├─ @jest/types@24.9.0
├─ @sheerun/mutationobserver-shim@0.3.2
├─ @testing-library/dom@6.4.1
├─ @testing-library/jest-dom@4.1.0
├─ @testing-library/react@9.2.0
├─ @types/istanbul-lib-report@1.1.1
├─ @types/istanbul-reports@1.1.1
├─ @types/react-dom@16.9.1
├─ @types/testing-library__dom@6.4.0
├─ @types/testing-library__react@9.1.1
├─ @types/yargs-parser@13.1.0
├─ @types/yargs@13.0.3
├─ css.escape@1.5.1
├─ css@2.2.4
├─ diff-sequences@24.9.0
├─ indent-string@4.0.0
├─ min-indent@1.0.0
├─ redent@3.0.0
├─ strip-indent@3.0.0
└─ wait-for-expect@1.3.0
Done in 20.04s.
$

```

### 2019-09-25 

After unpublishing (see notes on this below), I corrected the URL to be
"plural" (in package.json, the "s" was missing).

Now commit and push changes:
```
$ git -a -m "fixed urls and added journal"
Unknown option: -a
usage: git [--version] [--help] [-C <path>] [-c name=value]
           [--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]
           [-p | --paginate | --no-pager] [--no-replace-objects] [--bare]
           [--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]
           <command> [<args>]
$ git commit -a -m "fixed urls and added journal"
[master d52807b] fixed urls and added journal
 2 files changed, 34 insertions(+), 3 deletions(-)
 create mode 100644 Journal.md
$ git push
Counting objects: 7, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (7/7), done.
Writing objects: 100% (7/7), 937 bytes | 0 bytes/s, done.
Total 7 (delta 4), reused 0 (delta 0)
remote: Resolving deltas: 100% (4/4), completed with 2 local objects.
To git@github.com:mandolyte/uw-languages-rcl.git
   c06c0c2..d52807b  master -> master
$
```

Now run `yarn publish`:
```
$ yarn publish
yarn publish v1.17.3
[1/4] Bumping version...
info Current version: 0.0.1
question New version:
[2/4] Logging in...
[3/4] Publishing...
$ rm -fr ./dist & babel ./src --out-dir ./dist -s inline
Successfully compiled 4 files with Babel.
error Couldn't publish package: "https://registry.yarnpkg.com/uw-languages-rcl: uw-languages-rcl cannot be republished until 24 hours have passed."
info Visit https://yarnpkg.com/en/docs/cli/publish for documentation about this command.
$
```
Will have to wait until tomorrow!

## Appendices

### To generate just the styleguide

```
$ yarn styleguide:build
yarn run v1.17.3
$ styleguidist build
Building style guide...
Style guide published to:
/home/mando/Projects/mandolyte/uw-languages-rcl/styleguide
Done in 11.45s.
$
```

### Unpublishing

*NOTE: you must first login!*

Transcript:
```
$ npm login
Username: mandolyte
Password:
Email: (this IS public) cecil.new@gmail.com
Logged in as mandolyte on https://registry.npmjs.org/.
$ npm unpublish uw-languages-rcl -f
npm WARN using --force I sure hope you know what you are doing.
-uw-languages-rcl
$
```

### Creating a React Material-ui application

*Where "my-app" is the name of your new React project*

```
$ npx create-react-app my-app
$ cd my-app
$ npm install @material-ui/core
```
