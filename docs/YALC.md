# YALC info

Took some time to document how to create and manage local RCL/Module repos. This avoids have to publish to NPM just to use a package in the context of an app (or another RCL).

**Using YALC**
To install and see help:
```
$ yarn global add yalc
$ yalc help
yalc [command] [options] [package1 [package2...]]

Commands:
  publish        Publish package in yalc local repo
  installations  Work with installations file: show/clean
  push           Publish package in yalc local repo and push to all
                 installations
  add            Add package from yalc repo to the project
  link           Link package from yalc repo to the project
  update         Update packages from yalc repo
  remove         Remove packages from the project
  retreat        Remove packages from project, but leave in lock file (to be
                 restored later)
  check          Check package.json for yalc packages
  dir            Show yalc system directory

Options:
  --help  Show help                                                    [boolean]
```
In RCL:
```
$ yalc publish
```
*NOTE: this command runs `prepublishOnly` script.*

To see where the “local repo” is:
```
$ yalc dir
C:\Users\mando\AppData\Local\Yalc
```
To use the locally sourced RCL, go to the app (or other RCL):
```
$ yalc link markdown-translatable
Package markdown-translatable@1.0.0-rc.1+05065799 linked ==> C:\Users\mando\Projects\bugfix-mandolyte-120-oq-highlight\tc-create-app\node_modules\markdown-translatable.
```
To confirm linkage:
```
$ cd node_modules/
$ ls -al | grep markdown-translatable
lrwxrwxrwx 1 mando 197611    100 Jun 26 07:44 markdown-translatable -> /c/Users/mando/Projects/bugfix-mandolyte-120-oq-highlight/tc-create-app/.yalc/markdown-translatable/
```
Or to see all linked modules: `ls -al | grep ^l`

Note that the “link” command copies from the “local repo” to a subdirectory “.yalc” in the app folder. Thus, add these two entries to “.gitignore”
```
/.yalc
yalc.lock
```

Then `yarn start` as usual to launch the app. If start is already running, then it will detect the new link sources and recompile.

**Workflow**
1. make changes to RCL
2. `yalc publish`(this builds and puts the app into the local store)
3. in using APP or RCL, `yalc link yourRclName` (this copies from the local store to `.yalc` and make it linkable in node_modules). Once copied, you can simply use `yalc update` to update to any published changes.
4. if `yarn start` is running it will auto rebuild; otherwise run `yarn start`

**To unlink**

- Run `yalc remove my-package`, it will remove package info from package.json and yalc.lock
- Run `yalc remove --all` to remove all packages from project.