# Check for possible merge conflicts

If multiple developers are working in the same repo, then the baseline for your feature branch may be behind the master branch (or develop branch for apps). Since you know your code best, it is expected that the developer will merge the master branch into their feature branch before proceeding to the checklist below. The general steps are below. Be sure to use “develop” instead of “master” if working on an app. Use “master” for libraries/RCLs.

# First switch to develop or master branch

git switch master  # for RCLs
git switch develop # for apps

# Then update your local copy of master with latest commits

git pull 

# Next switch back to your feature branch

git switch feature-branch

# Make sure you have the latest commits locally

git pull 

# Finally, merge in master to your feature branch

git merge master     # for RCLs
git merge develop    # for apps




# ! Deal with all conflicts !

## Merging large hierarchies or large number of files/folders

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
