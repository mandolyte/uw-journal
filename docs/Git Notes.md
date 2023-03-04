# Git Notes

## add other repo branch into my repo branch

https://stackoverflow.com/questions/21353656/merge-git-repo-into-branch-of-another-repo

You can't merge a repository into a branch. You can merge a branch from another repository into a branch in your local repository. Assuming that you have two repositories, foo and bar both located in your current directory:

$ ls
foo bar
Change into the foo repository:

$ cd foo
Add the bar repository as a remote and fetch it:

$ git remote add bar ../bar
$ git remote update
Create a new branch baz in the foo repository based on whatever your current branch is:

$ git switch -c baz
Merge branch somebranch from the bar repository into the current branch:

$ git merge --allow-unrelated-histories bar/somebranch
(--allow-unrelated-histories is not required prior to git version 2.9)


## set-branch script

```sh
#!/bin/bash
if [[ "$1" == "" ]]
then
	echo branch name is required
	exit
fi
git checkout -b $1
git push --set-upstream origin $1

```

## set account

```
  git config --global user.email "cecil.new@gmail.com"
  git config --global user.name "mandolyte"
```


## Using Git Branching
Some details and learnings below. To setup a feature branch, say ‘feature-cn-tooltips’ (the word feature, then my initials, then a descriptive short name for the feature).
- First do a checkout on the branch. This will create it locally and set it to be the default.
- git checkout -b feature-cn-tooltips
- Next create it on remote side
- git push --set-upstream origin feature-cn-tooltips
- Now make changes as normal, test them, etc.
- Next commit them and push them, etc.
- Finally, when ready, request PR (substitute repo before using)
https://github.com/unfoldingWord/gitea-react-toolkit/pull/new/feature-cn-tooltips

To clone a specific branch:
$ git clone --single-branch --branch enhancement-klappy-hooks-contexts git@github.com:unfoldingWord/scripture-resources-rcl.git
Cloning into 'scripture-resources-rcl'...
remote: Enumerating objects: 296, done.
remote: Counting objects: 100% (296/296), done.
remote: Compressing objects: 100% (217/217), done.
remote: Total 1860 (delta 116), reused 244 (delta 78), pack-reused 1564
Receiving objects: 100% (1860/1860), 1.84 MiB | 1021.00 KiB/s, done.
Resolving deltas: 100% (1104/1104), done.
$ mv scripture-resources-rcl/ hooks-contexts ## rename it
$ cd hooks-contexts/
$ git branch --list
* enhancement-klappy-hooks-contexts
$



Create a branch:
$ git branch feature-cn-tooltips

List branches:
$ git branch
* develop
  feature-cn-tooltips
$ git branch --list
* develop
  Feature-cn-tooltips

Delete a branch:
$ git branch -d feature-cn-tooltips
Deleted branch feature-cn-tooltips (was e47c07d).
$ git branch
* develop

Create and switch to a branch (notice which has the asterisk):
$ git checkout -b feature-cn-tooltips
Switched to a new branch 'feature-cn-tooltips'
$ git branch
  develop
* feature-cn-tooltips
On git status, the current branch is shown:
$ git status
On branch feature-cn-tooltips
Changes to be committed:
On first push, must create branch on remote/server:
$ git push
fatal: The current branch feature-cn-tooltips has no upstream branch.
To push the current branch and set the remote as upstream, use      

    git push --set-upstream origin feature-cn-tooltips

$ git push --set-upstream origin feature-cn-tooltips
Counting objects: 8, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (8/8), done.
Writing objects: 100% (8/8), 1.29 KiB | 660.00 KiB/s, done.
Total 8 (delta 5), reused 0 (delta 0)
After push, note the instructions on how to do a PR:
$ git push --set-upstream origin feature-cn-tooltips
Total 0 (delta 0), reused 0 (delta 0)
remote: 
remote: Create a pull request for 'feature-cn-tooltips' on GitHub by visiting:
remote:      https://github.com/unfoldingWord/gitea-react-toolkit/pull/new/feature-cn-tooltips
remote:
To github.com:unfoldingWord/gitea-react-toolkit.git
 * [new branch]      feature-cn-tooltips -> feature-cn-tooltips
Branch 'feature-cn-tooltips' set up to track remote branch 'feature-cn-tooltips' from 'origin'.

Git Notes
Rebase vs Merge into Feature Branch
Reference:https://superuser.com/questions/224085/git-merge-master-into-a-branch
Deleting a branch
// delete branch locally
git branch -d localBranchName

// delete branch remotely
git push origin --delete remoteBranchName
Branching from a branch
From: https://stackoverflow.com/questions/4470523/create-a-branch-in-git-from-another-branch
But to answer your question:
$ git checkout -b myFeature dev
Above creates myFeature branch off the dev branch. Do your work and commit.
$ git commit -am "Your message"
Now merge your changes to dev without a fast-forward
$ git checkout dev 
$ git merge --no-ff myFeature
Now push changes to the server
$ git push origin dev 
$ git push origin myFeature

