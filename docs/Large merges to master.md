# Merging large hierarchies or large number of files/folders

## Summary
Using the newFormat branch as an example:
(note: the below is going the wrong way. after I committed, I switched to master and ran `git merge newFormat` and it work smoothly. Then after confirming all looked ok, then I pushed to remote master)

```
$ git clone ...
$ git switch newFormat
$ git config merge.renameLimit 999999  # bump the limit
$ git merge master
$ # here is where I had to "resolve" conflicts by manually
$ # removing files with git rm -fr deu 1sa ...             
$ git config --unset merge.renameLimit
$ git commit -a -m "merged master into newFormat"
$ git push # optional?
```

Then:
```
$ git switch master #back to master
$ git merge newFormat #conflicts should be gone and this should go fast
$ git status
On branch master
Your branch is ahead of 'origin/master' by 40 commits.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean

mando@DESKTOP-0V8P6MM MINGW64 ~/Projects/git.door43.org/test_tn_7col_format/en_tn (master)
$ git push
```


## Research

Info at:
https://stackoverflow.com/questions/4722423/how-to-merge-two-branches-with-different-directory-hierarchies-in-git

Recommends first doing this command:
```
git merge --abort                    # undo the started merge 
git switch master                    # switch back to master
git pull                             # get the latest (in case it has changed)
git switch newFormat                 # back to the branch
git config merge.renameLimit 999999  # bump the limit
git merge master                     # try again
git config --unset merge.renameLimit # reset limits
```

After the merge, I still had a lot of conflicts. For about 20 directories, I got this message:
```
CONFLICT (modify/delete): 1ti/01/15.md deleted in HEAD and modified in master. Version master of 1ti/01/15.md left in tree.
```

So I used `git rm -fr 1ti` to remove them manually as a way of "resolving the conflicts".
Note: actually removed multiple folders using a list of them on the command.

Then I committed and pushed:
```
$ git commit -a -m "merged master into newFormat branch"
[newFormat 350ba134a] merged master into newFormat branch

$ git push
Enumerating objects: 10, done.
Counting objects: 100% (10/10), done.
Delta compression using up to 4 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (4/4), 498 bytes | 249.00 KiB/s, done.
Total 4 (delta 3), reused 0 (delta 0), pack-reused 0
remote:
remote: Visit the existing pull request:
remote:   https://qa.door43.org/unfoldingWord/en_tq/pulls/65
remote:
remote: . Processing 1 references
remote: Processed 1 references in total
To https://qa.door43.org/unfoldingWord/en_tq.git
   f8391baef..350ba134a  newFormat -> newFormat

$
```

*Summary*
```
$ git clone ...
$ git switch newFormat
$ git config merge.renameLimit 999999  # bump the limit
$ git merge master
$ # here is where I had to "resolve" conflicts by manually
$ # removing files with git rm -fr deu 1sa ...             
$ git config --unset merge.renameLimit
$ git commit -a -m "merged master into newFormat"
$ git push
```


# Then commit and push your updated feature branch

git commit -a -m "merged develop/master into feature"
git push
