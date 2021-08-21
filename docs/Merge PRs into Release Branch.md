# Merge PRs into Release Branch

## 2021-02-18

At this writing, there are a number of PRs in the Zenhub review pipeline. We just released 1.1.0 and these were considered out of scope... so the are just sitting there.

Some of the PRs are actually against other packages, not tc-create itself. So I'll ignore those for the moment.

These are the ones against tc-create itself (branch and PR):
- `bug-cn-498-set-tn-rows-page-to-25` https://github.com/unfoldingWord/tc-create-app/issues/685 
- `dependabot/npm_and_yarn/react-select-4.1.0` https://github.com/unfoldingWord/tc-create-app/issues/687
- `dependabot/npm_and_yarn/scripture-resources-rcl-3.2.0` https://github.com/unfoldingWord/tc-create-app/issues/700
- `dependabot/npm_and_yarn/cypress-6.5.0` https://github.com/unfoldingWord/tc-create-app/issues/701

All but the first are "dependabot" catches.

**PR 685**

1. Switch to the develop branch first in tc-create
2. Do a `git pull` to be sure it is up to date
3. Create release branch 1.1.1: `set-branch v1.1.1` (see below for the script)
4. At this point:
	- the release branch is same as develop
	- the current branch is the release branch
5. While sitting in the release branch, merge PR 685 into v1.1.1
```sh
$ git merge bug-cn-498-set-tn-rows-page-to-25
```
6. If there are conflicts (there will be at least one due to `./public/build_number`) . In the build_number case, always accept the current change, that is, the value that is in tc-create already. It will be higher. Or maybe, always take the higher value?
7. After resolving the conflict(s), then commit: `git commit -a -m "resolve merge conflict"`
8. Note: since the preceding is a commit, the commit hook will run and will bump and also commit the build_number.
9. Test by 
	- examine the code to see if has the change (confirmed)
	- using `yarn start` to verify change (confirmed - there are now 25 rows per page)

Repeat for each of the branches to be included in v1.1.1. Here are some notes taken along the way.

**PR 687**
This is a dependabot change, upgrading relact-select from 3.0.4 to 4.1.0. Notice the major number changed; so it may be a breaking change. Confirmed that current version in `package.json` is `^3.0.4`. The merge command returns error message:
```sh 
$ git merge dependabot/npm_and_yarn/react-select-4.1.0
merge: dependabot/npm_and_yarn/react-select-4.1.0 - not something we can merge
```
Merge needs to have the branch local. So check it out explicitly:
```sh
$ git checkout dependabot/npm_and_yarn/react-select-4.1.0
Switched to a new branch 'dependabot/npm_and_yarn/react-select-4.1.0'
Branch 'dependabot/npm_and_yarn/react-select-4.1.0' set up to track remote branch 'dependabot/npm_and_yarn/react-select-4.1.0' from 'origin'.
```
Now switch back to v1.1.1 by `git switch v1.1.1`.

Now do the merge:
```sh
$ git merge dependabot/npm_and_yarn/react-select-4.1.0
Auto-merging yarn.lock
CONFLICT (content): Merge conflict in yarn.lock
Auto-merging package.json
Automatic merge failed; fix conflicts and then commit the result.
```

Notice the merge conflict, this time in `yarn.lock`. Since this file is updated with a `yarn install` command, simply accept current change and commit. Then the merge conflict will go away. Next need to rerun the install command to ensure that what is in `package.json` is actually being used. Might need to delete it prior to running so it can completely rebuild from scratch?


**PR 700**
1. checkout the branch: `git checkout dependabot/npm_and_yarn/scripture-resources-rcl-3.2.0`
2. switch back to v1.1.1: `git switch v1.1.1`
3. merge: `git merge dependabot/npm_and_yarn/scripture-resources-rcl-3.2.0`
4. result is yarn lock conflict
	1. in vs code, accept current change and save file
	2. commit (to make merge conflict go away)
	3. run yarn install (to update entry in yarn lock for scripture-resources to v3.2.0)
	4. commit again since yarn lock was updated
5. Since 4.4 above did a commit, we are now clean read for next PR


**PR 701**
1. checkout the branch: `git checkout dependabot/npm_and_yarn/cypress-6.5.0`
2. switch back to v1.1.1: `git switch v1.1.1`
3. merge: `git merge dependabot/npm_and_yarn/cypress-6.5.0`
4. result was auto merged:

```sh
$ git merge dependabot/npm_and_yarn/cypress-6.5.0
Auto-merging yarn.lock
Auto-merging package.json
Merge made by the 'recursive' strategy.
 package.json |   2 +-
 yarn.lock    | 113 ++++++++++++++++++++++++++++++++++++++---------------------
 2 files changed, 74 insertions(+), 41 deletions(-)
```
6. confirm changes: `git status`. Hmmm:

```sh
$ git status
On branch v1.1.1
Your branch is ahead of 'origin/v1.1.1' by 10 commits.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean
$
```
Asked Rich Mahn about [here](https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE/topic/Git.20merge.20question/near/226827774)

7. pushed:

```sh
$ git push
Enumerating objects: 73, done.
Counting objects: 100% (61/61), done.
Delta compression using up to 4 threads
Compressing objects: 100% (27/27), done.
Writing objects: 100% (33/33), 125.42 KiB | 742.00 KiB/s, done.
Total 33 (delta 20), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (20/20), completed with 6 local objects.
To github.com:unfoldingWord/tc-create-app.git
   810ff2b..d66152e  v1.1.1 -> v1.1.1

mando@DESKTOP-0V8P6MM MINGW64 ~/Projects/unfoldingWord/tc-create-app (v1.1.1)
$
```


## Resources

### Set Branch

```sh
#!/bin/sh

if [ "$1x" = "x" ]; then
	echo Feature or bugfix branch name required
	exit
fi

git checkout -b $1
git push --set-upstream origin $1

echo Done.
```

