# Issue 518 "SPIKE: Admin App - Conflict resolution"

- Issue: https://github.com/unfoldingWord/tc-create-app/issues/518
- Drive location for diagram: https://drive.google.com/file/d/1HQqLWeaTj6aLCzJrIGQCAfA-zgwcpjmh/view?usp=sharing_eip&ts=5fb3d6ec
- Workflow diagram: https://app.diagrams.net/#G1HQqLWeaTj6aLCzJrIGQCAfA-zgwcpjmh
- Rich's PoC work: https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE/topic/DCS.20POC.20-.20Merge.20branches.20with.20API/near/215030799
- The PoC artifacts at https://unfoldingword-dev/dcs-api-merge-poc 

# 2020-12-01 

If conflict resolution must work with how tc-create works today, then tc-create itself could be used as the starting point for the code, since:
- Needs login
- Needs to select org
- Needs to select resource type
- Needs to select lanugage
- Needs to select file

Of course, the file may not exist in master yet. In which case, no conflicts are possible and thus this tool will not be needed/used.

In addition this tool will need to know the user's branch. This can be computed, since branch names are always the following:
- user name, for example "birch"
- a hyphen
- the string "tc-create-1"

Prerequisites:
- PR submitted
- need an API call to query PR, especially whether "mergeable" (see below)
- if mergeable, then no need for this tool; thus mergeable needs to be false

So the flow for the conflict resolution app will be:
- login
- select org
- select resource type
- select lanugage
- select file, by:
  - computing the user's branch
  - presenting a picklist of files in the user's branch
- if same file not in master, then throw alert and let them pick another
- retrieve the file from master and user branch
- present them side by side (see @bruce spiedel's mockup)
- when resolution is complete, overwrite/squash master with resolved file.

Alternate workflow:
- login
- select org
- select resource type
- select language
- select PR

How to query a PR:
- the curl and URL below query a repo for all open pulls
```
curl -X GET "https://qa.door43.org/api/v1/repos/unfoldingword/en_tn/pulls?state=open&sort=recentupdate&access_token=c8b93b7ccf7018eee9fec586733a532c5f858cdd" -H "accept: application/json"

https://qa.door43.org/api/v1/repos/unfoldingword/en_tn/pulls?state=open&sort=recentupdate&access_token=c8b93b7ccf7018eee9fec586733a532c5f858cdd
```
- the JSON result for the above (at this writing) has two PRs. See `response_1606841847488.json`
- the PR number, the title (and body), and the mergeable attribute are all top level attributes of each PR (there is an array of them).
```
[
  {
    "id": 3587,
    "url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/2294",
    "number": 2294,
...
    "title": "WIP: Correct alignment of Greek articles not rendered in English, with pronouns implied by Greek cases.",
    "body": "Fixed alignment of articles in 3Jn, Tit, and Eph according to new standard. ",
...
    "html_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/2294",
    "diff_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/2294.diff",
    "patch_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/2294.patch",
    "mergeable": false,
    "merged": false,
```
The files changed in the PR are not listed, but may be found in the return of diff_url. Each one shows up like this:
```
diff --git a/50-EPH.usfm b/50-EPH.usfm
diff --git a/57-TIT.usfm b/57-TIT.usfm
diff --git a/65-3JN.usfm b/65-3JN.usfm
```
where "a/" is the base and "b/" is the changed file.




# UI/X description for Alt-view

To identify what is need to resolve:
- Org
- resource type or repo
- either:
  - book or article name with path lookup from manifest
  - path
- working branch
- branch to be merged into working branch


# An alternate view of the translation workflow (partial)

## Overview
In this alternate view, the translate process at a high level could be:

The translator:
- selects a file to translate or to update
- makes changes to the file
- save changes to the file
- this may continue for for some time
- proposes to replace the orginal with their file
- if discovered that someone else is also working with the same file, then the first translator's changes are merged into the second translator's changes, reconciling any conflicts that may arise.

The reviewer:
- collaborates with translator to resolve any questions
- approves the changes, which...
- replaces the orginal with the approved (updated) version

This alternate view is, in essence, a document life-cycle. This is not how Gitea works, but it is possible to layer this life-cycle over the top of Gitea.

## Solution Description

1. Let each branch be associated to the translator and to the file being edited. That is, the name of the branch must be a combination of user and file name. So if the user is working on three files, then that will involve three branches.
2. If two users are working on same file, for example, different chapters in the same book, then this results in two branches being created.
3. There is a reviewer role who will be responsible for approving changes and updating the master branch. 

See scenarios below for details.

## Scenarios

In the following, unless otherwise noted, the book of Ruth is used as book being created or updated. Alice and Bob are the translators and Charles is the reviewer. The acronym PR stands for Pull Request (a Git term). 

If the resource being worked on is a book-oriented resource, then the branch created is a combination of the bookd id and the user's id. For example, "rut-alice".

If the resource is a tA or tN, then the full name of the article is used. For example, "figs-metaphor-bob".

### Initial Creation of a Translated File

_Description_ In this case there are no conflicts to resolve with other translators and no actionable feedback from the reviewer.

_Steps_ In this case, the translator uses tc-create to create a file; saves the work; submits a PR. The reviewer merges the changes into the master branch.

_Details_ Alice selects Ruth to translate. A branch named "rut-alice" is created. Updates are made and saved. After some time passes, the translation is ready for review. Alice submits a PR. The reviewer is assigned along with any comments by Alice about the PR. The reviewer, having no feedback, merges Alice's branch into the master branch. Since this book is new and it is the only content in Alice's branch, it simply creates the file in the master branch. The working branch "rut-alice" is deleted.

### Revision to a Translated File

_Description_ This describes the case where there are no conflicts to resolve with other translators and no actionable feedback from the reviewer.

_Steps_ In this case, the translator uses tc-create to open and update a file; saves the work; submits a PR. The reviewer merges the changes into the master branch.

_Details_ Alice selects Ruth to update. A branch named "rut-alice" is created. Updates are made and saved. After some time passes, the updated translation is ready for review. Alice submits a PR. The reviewer is assigned along with any comments by Alice about the PR. The reviewer, having no feedback, merges Alice's branch into the master branch. Since this book exists in the master branch and this book is the sole content in the Alice's working branch, the effect is simply to replace the file in the master branch.

### Two translators working to Translate the same File, but no conflicts

_Description_ This describes the case where there are two translators are working on the same file, but are working in different chapters with no conflicts; and with no actionable feedback from the reviewer.

_Steps_ In this case, the translators uses tc-create to open and update the same file and save their work. At some point, one of them will attempt to submit a PR. At this point, it is discovered that that someone else is also working in the same file; this causes the PR to fail to be created. At this point, the first user's changes are merged into the still on-going changes of the second user. Since there are no conflicts, then this will be fairly straighforward since both users' changes will be accepted into the second user's branch. At this point, the second user's branch will incorporate all the first user's changes.

Later the other translator will attempt to submit a PR. Since master has not changed since they began, the PR will succeed and the changes will replace the file in the master branch.

_Details_ Alice selects Ruth to translate. A branch named "rut-alice" is created. Bob does the same, creating a branch named "rut-bob". Updates are made and saved by both Alice and Bob. After some time passes, Bob attempts to submits a PR. However, it is discovered that Alice is also making changes to the same file; the PR fails to be created. Then Bob's changes are merged into Alice's changes, so that Alice's file now has both hers and Bob's changes. Bob's branch is deleted. She continues to work and submits a PR when finished.

Since the file in the master branch is unchanged since Alice started, the PR succeeds and the new content replaces the old in the master branch. Alice's branch is deleted.

### Two translators working to Translate the same File, with conflicts

_Description_ This path describes the case where there are two translators are working on the same file and end up translating the same portion of the file differently; and with no actionable feedback from the reviewer.

_Steps_ In this case, the translators uses tc-create to open and update the same file; save their work. At some point, one of them will attempt to submit a PR. At this point, it is discovered that that someone else is also working in the same file. At this point, the first user's changes are merged into the still on-going changes of the second user. Since there are  conflicts, then this could be fairly complex, especially if the desired result is a combination of both translators' changes. 

Once the second user's branch has the desired reconciled content, then the first user's branch is deleted.

Later the other translator will attempt to submit a PR. Since master has not changed since they began, the PR will succeed and the changes will replace the file in the master branch.

_Details_ Alice selects Ruth to translate. A branch named "rut-alice" is created. Bob does the same, creating a branch named "rut-bob". Updates are made and saved by both Alice and Bob. In particular, both have made changes to the same part of the the resource file. After some time passes, Bob attempts to submits a PR. However, it is discovered that Alice is also making changes to the same file. So Bob's changes are merged into Alice's changes, reconciling the conflicts so that Alice's file now has both hers and Bob's changes, with any conflicts resolved. Bob's branch is deleted. Alice continues to work and submits a PR when finished.

Since the file in the master branch is unchanged since Alice started, the PR succeeds and the new content replaces the old in the master branch. Alice's branch is deleted.

### Reviewer has changes

If, at any time after PR is submitted, the reviewer wants changes made, then the reviewer collaborates with the translator to make the desired changes in their branch. The content associated with the PR will update as well. This continues until the reviewer and translator are satisfied. At this point, the PR is used to merge the user's branch into the master branch.

# Essentials

The term "merge" may be used in two senses in this document. 
- Branches can be merged
- Changes to a single file can be merged. I will try to use the word "resolve" for this sense.

**Fundamental Requirements**
1. User needs to know if a file in their branch has changed on master
2. User needs to be able to merge master into their branch when preceding is true and resolve any conflicts
3. User needs to be able to "submit a PR" for their changes.

# Flow of Changes to a File

## At start

- When a user begins making changes, a branch is created from the master branch.
- Later, when the user wishes to save the changes, they will click the "Save" button.
- This will save the changes to the user's branch.
- The user will check occasionally to see if anyone else has updated the same files in master that they are updating. And if so, then the user must be able to merge changes to master into their own branch. Furthermore, if such a merge entails 
- Later there will be an attempt to merge the user's branch to the master branch.

NOTE 1: if the master has not been changed subsequent to the user's branch creation, 
then the user's branch will always be mergeable. (branch merged)

NOTE 2: if the master branch has been changed, but the change does not involve any file they 
have edited in their branch, then the user's branch will always be mergeable. (branch merged)

The remainder of this describes what happens when the above two NOTES are not the case, namely, 
the master version of a file has changed and the user has also made changes to the same file.
In this case, the two sets of changes must be resolved. This may or may not involve conflicts
(i.e., changes made to the same lines).

The following describes the process for a single file, but it must be done for all files 
where the changes must be merged (file merge).

*Gitea APIs needed:*
Detect whether master and branch are mergeable without conflict. If there are conflicts, then
by definition, some subset of the files changed will have to be resolved.

## Case 1: No conflicts

This will be the case when changes are not to the same or not to adjacent lines. 
If this is the case, then the two versions of the file are considered "mergeable".

*Gitea APIs needed:*
Merge master into branch.

In this case, the user can update their files with the newer files from master and may
continue making changes with the latest changes made by others.

## Case 2: Conflicts

If Gitea cannot automatically merge due to conflicts (same or adjacent lines have been 
changed in both versions of the file), then the user must be able to resolve the 
conflicts. There are three possible actions:

1. The user selects a change from master version to replace their own
2. The user retains their change, discarding the version of the change from master
3. The user opts to combine in some fashion the two changes

In the third case, we expect the user to abandon the merge, make the synthesized change
in their branch, then attempt the merge again. This process essentially reduces the
possible actions to two in the end.

*Gitea APIs needed:*
A comparison difference result which can be used in an application/RCL to guide the user
thru resolving the conflicts.

## At end

Once the user is done, then the changes must be submitted as a Pull Request (PR).
If the user continues making changes in their branch, then the PR will automatically
include new changes. This is an important training issue.

*Gitea APIs needed:*
A way to create the PR for the user's branch.

## API Summary

1. Detect whether master and branch are mergeable without conflict. If there are conflicts, then by definition, some subset of the files changed will have to be resolved.
2. Merge master into branch.
3. A comparison difference result which can be used in an application/RCL to guide the user thru resolving the conflicts.
4. A way to create the PR for the user's branch.


# Study of Rich's PoC

## Setup

Here are the instructions to setup:

For merge with NO Conflicts:

Online:
Run here: https://repl.it/@richmahn/dcs-api-merge-poc-merge-no-conflicts
Offline:
Clone git@github.com:unfoldingword-dev/dcs-api-merge-poc.git, cd to it, and check out the merge-no-conflicts branch, run bash main.sh.

For merge with conflicts:

Online:
Run here: https://repl.it/@richmahn/dcs-api-merge-poc-merge-conflicts
Offline:
Clone git@github.com:unfoldingword-dev/dcs-api-merge-poc.git, cd to it, and check out the merge-conflicts branch, run bash main.sh.

## Notes

This website converts curl commands into Javascript fetch:
https://kigiri.github.io/fetch/

For example, this:
```
curl POST "https://github.com/api/v1/repos/$org/$repo/pulls?access_token=$token" -H "accept: application/json" -H "Content-Type: application/json" -d '{ "base": "master", "body": "Merging user branch into master", "head": "$branch", "title": "$branch into master"}'
```

becomes this:
```
fetch("https://github.com/api/v1/repos/$org/$repo/pulls?access_token=$token", {
  body: "{ \"base\": \"master\", \"body\": \"Merging user branch into master\", \"head\": \"$branch\", \"title\": \"$branch into master\"}",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json"
  },
  method: "POST"
})
```

Looks like the magic all happens when a PR is submitted. Here is the code from Rich's script to merge master into user's branch:
```sh
echo "
=========
MAKING PR FOR master INTO USER BRANCH $branch:

"
response=$(curl --silent -X POST "$host/api/v1/repos/$org/$repo/pulls?access_token=$token" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"base\": \"$branch\", \"body\": \"Merging master into user branch\", \"head\": \"master\", \"title\": \"master into $branch\"}")
echo "$response"

pr_num=$(echo "$response" | jq -r '.number')
pr_url=$(echo "$response" | jq -r '.url')
diff_url=$(echo "$response" | jq -r '.diff_url')
patch_url=$(echo "$response" | jq -r '.path_url')
mergeable=$(echo "$response" | jq -r '.mergeable')

echo -e "
PR URL: $pr_url
DIFF URL: $diff_url

MERGEABLE: $mergeable"

read -p "
Press ENTER to continue"

if [[ $mergeable == "false" ]]; then
  echo -e "\nIS NOT MERGEABLE!"
  exit
fi
echo -e "\nIS MERGEABLE!"
```

The post returns a JSON object with the following available:
- the PR number
- the URL to the PR
- the URL to the difference report
- the URL to the patch
- a boolean indicating whether the two are mergeable or not.

# Experiment

## steps
1. First, create a repo:
2. Second, create a file in the repo (master branch):
3. Third, create a branch of the repo:
4. Fourth, update the file in branch:
5. Fifth, update the file in master branch:
6. Sixth, submit PR via swagger:

## records

1. First, create a repo:
```
https://qa.door43.org/repo/create
Created "mergetest" using form at above link.
It defaulted with a README.md and a LICENSE.md
```
2. Second, create a file in the repo (master branch):
File named: base.md
```
- line 1
- line 2
- line 3
- line 4
- line 5
- line 6
- line 7
```
Used DCS directly (New File button) and committed directly to master branch.

3. Third, create a branch of the repo: cecil.new-patch-1
Note: branch name auto-generated and I accepted it.
4. Fourth, update the file in branch:
```
Note: I had to combine the above by editing the existing file and committing to the new branch.
New file:

- line 1
- line 2
- line 3a
    - item 1
    - item 2
- line 4a
- line 5
- line 6
- line 7

Note: When I committed to new branch a pull request was automatically started as well.
However, I never "pushed" the button to actually create the PR so it was not created.
```
5. Fifth, update the file in master branch:
```
- line 1
- line 2
- line 3b
    - item 1a
    - item 2a
- line 4b
- line 5
- line 6
- line 7
```
Above comitted directly to master. There should be conflicts with below PR.
6. Sixth, create PR via swagger:
Form requires: owner=cecil.new, repo=mergetest

Note: be sure to use the QA version of swagger:
Swagger at: https://qa.door43.org/api/v1/swagger#/repository/repoCreatePullRequest

Note: had to authorize at top of Swagger page... used "basic".

Body for PR:
```json
{ 
    "base": "cecil.new-patch-1",
    "body": "Merging master into user branch",
    "head": "master", 
    "title": "master into cecil.new-patch-1"
}
```
pasted above into swagger replacing model template value.
![Snapshot](Screenshot_2020-11-23_091336.png)

Here is the curl command it generated. NOTE! had to use access token (in query) using token from Rich's script.
```
curl -X POST "https://git.door43.org/api/v1/repos/cecil.new/mergetest/pulls" -H "accept: application/json" -H "authorization: Basic Y2VjaWwubmV3OjM1XkhoRF5IRiRkKiNOKkE=" -H "Content-Type: application/json" -d "{ \"base\": \"cecil.new-patch-1\", \"body\": \"Merging master into user branch\", \"head\": \"master\", \"title\": \"master into cecil.new-patch-1\"}"
```

Response headers were:
```
 access-control-allow-headers: Authorization,DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range 
 access-control-allow-methods: GETPOSTPUTOPTIONSPATCHDELETE 
 access-control-allow-origin: * 
 access-control-expose-headers: Content-Length,Content-Range 
 connection: keep-alive 
 content-type: application/json; charset=UTF-8 
 date: Mon23 Nov 2020 14:39:37 GMT 
 server: nginx/1.16.1 
 transfer-encoding: chunked 
 x-content-type-options: nosniff 
 x-frame-options: SAMEORIGIN 
 ```


```json
{
  "id": 3907,
  "url": "https://qa.door43.org/cecil.new/mergetest/pulls/1",
  "number": 1,
  "user": {
    "id": 31488,
    "login": "dcs-poc",
    "full_name": "",
    "email": "dcs-poc@noreply.door43.org",
    "avatar_url": "https://qa.door43.org/user/avatar/dcs-poc/-1",
    "language": "en-US",
    "is_admin": false,
    "last_login": "2020-10-27T19:11:06Z",
    "created": "2020-10-27T19:10:30Z",
    "repo_languages": null,
    "username": "dcs-poc"
  },
  "title": "master into cecil.new-patch-1",
  "body": "Merging master into user branch",
  "labels": [],
  "milestone": null,
  "assignee": null,
  "assignees": null,
  "state": "open",
  "comments": 0,
  "html_url": "https://qa.door43.org/cecil.new/mergetest/pulls/1",
  "diff_url": "https://qa.door43.org/cecil.new/mergetest/pulls/1.diff",
  "patch_url": "https://qa.door43.org/cecil.new/mergetest/pulls/1.patch",
  "mergeable": false,
  "merged": false,
  "merged_at": null,
  "merge_commit_sha": null,
  "merged_by": null,
  "base": {
    "label": "cecil.new-patch-1",
    "ref": "cecil.new-patch-1",
    "sha": "6e4315c018756a5b897278d772fd79849d56aa70",
    "repo_id": 60806,
    "repo": {
      "id": 60806,
      "owner": {
        "id": 13993,
        "login": "cecil.new",
        "full_name": "",
        "email": "cecil.new@noreply.door43.org",
        "avatar_url": "https://qa.door43.org/user/avatar/cecil.new/-1",
        "language": "en-US",
        "is_admin": true,
        "last_login": "2020-11-18T13:44:28Z",
        "created": "2019-09-19T13:37:42Z",
        "repo_languages": null,
        "username": "cecil.new"
      },
      "name": "mergetest",
      "full_name": "cecil.new/mergetest",
      "description": "mergetest",
      "empty": false,
      "private": false,
      "fork": false,
      "template": false,
      "parent": null,
      "mirror": false,
      "size": 19,
      "html_url": "https://qa.door43.org/cecil.new/mergetest",
      "ssh_url": "git@qa.door43.org:cecil.new/mergetest.git",
      "clone_url": "https://qa.door43.org/cecil.new/mergetest.git",
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
      "created_at": "2020-11-23T13:47:58Z",
      "updated_at": "2020-11-23T13:59:01Z",
      "permissions": {
        "admin": false,
        "push": false,
        "pull": false
      },
      "has_issues": true,
      "internal_tracker": {
        "enable_time_tracker": true,
        "allow_only_contributors_to_track_time": true,
        "enable_issue_dependencies": true
      },
      "has_wiki": true,
      "has_pull_requests": true,
      "ignore_whitespace_conflicts": false,
      "allow_merge_commits": true,
      "allow_rebase": true,
      "allow_rebase_explicit": true,
      "allow_squash_merge": true,
      "avatar_url": "",
      "language": "",
      "subject": "",
      "books": null,
      "title": "",
      "checking_level": "",
      "catalog": null
    }
  },
  "head": {
    "label": "master",
    "ref": "master",
    "sha": "d27ae2022ab57a09bdd382efc014a869ac442262",
    "repo_id": 60806,
    "repo": {
      "id": 60806,
      "owner": {
        "id": 13993,
        "login": "cecil.new",
        "full_name": "",
        "email": "cecil.new@noreply.door43.org",
        "avatar_url": "https://qa.door43.org/user/avatar/cecil.new/-1",
        "language": "en-US",
        "is_admin": true,
        "last_login": "2020-11-18T13:44:28Z",
        "created": "2019-09-19T13:37:42Z",
        "repo_languages": null,
        "username": "cecil.new"
      },
      "name": "mergetest",
      "full_name": "cecil.new/mergetest",
      "description": "mergetest",
      "empty": false,
      "private": false,
      "fork": false,
      "template": false,
      "parent": null,
      "mirror": false,
      "size": 19,
      "html_url": "https://qa.door43.org/cecil.new/mergetest",
      "ssh_url": "git@qa.door43.org:cecil.new/mergetest.git",
      "clone_url": "https://qa.door43.org/cecil.new/mergetest.git",
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
      "created_at": "2020-11-23T13:47:58Z",
      "updated_at": "2020-11-23T13:59:01Z",
      "permissions": {
        "admin": false,
        "push": false,
        "pull": false
      },
      "has_issues": true,
      "internal_tracker": {
        "enable_time_tracker": true,
        "allow_only_contributors_to_track_time": true,
        "enable_issue_dependencies": true
      },
      "has_wiki": true,
      "has_pull_requests": true,
      "ignore_whitespace_conflicts": false,
      "allow_merge_commits": true,
      "allow_rebase": true,
      "allow_rebase_explicit": true,
      "allow_squash_merge": true,
      "avatar_url": "",
      "language": "",
      "subject": "",
      "books": null,
      "title": "",
      "checking_level": "",
      "catalog": null
    }
  },
  "merge_base": "86d2077ef0b7167628d7f643ed9a1bcd6696d7cf",
  "due_date": null,
  "created_at": "2020-11-23T14:39:36Z",
  "updated_at": "2020-11-23T14:39:36Z",
  "closed_at": null
}
```

Created the PR (number 1). From JSON returned, the vital bits:
```
  "html_url": "https://qa.door43.org/cecil.new/mergetest/pulls/1",
  "diff_url": "https://qa.door43.org/cecil.new/mergetest/pulls/1.diff",
  "patch_url": "https://qa.door43.org/cecil.new/mergetest/pulls/1.patch",
  "mergeable": false,
```

Here is the `1.diff` file:
```
diff --git a/base.md b/base.md
index cac86d3..a74022a 100644
--- a/base.md
+++ b/base.md
@@ -1,7 +1,9 @@
 - line 1
 - line 2
-- line 3
-- line 4
+- line 3b
+    - item 1a
+    - item 2a
+- line 4b
 - line 5
 - line 6
 - line 7
\ No newline at end of file
```

Here is the `1.patch` file:
```
From d27ae2022ab57a09bdd382efc014a869ac442262 Mon Sep 17 00:00:00 2001
From: "cecil.new" <cecil.new@noreply.door43.org>
Date: Mon, 23 Nov 2020 13:59:01 +0000
Subject: [PATCH] update directly in master

---
 base.md | 6 ++++--
 1 file changed, 4 insertions(+), 2 deletions(-)

diff --git a/base.md b/base.md
index cac86d3..a74022a 100644
--- a/base.md
+++ b/base.md
@@ -1,7 +1,9 @@
 - line 1
 - line 2
-- line 3
-- line 4
+- line 3b
+    - item 1a
+    - item 2a
+- line 4b
 - line 5
 - line 6
 - line 7
\ No newline at end of file
-- 
2.24.3

```