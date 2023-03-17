In Discord:
https://discord.com/channels/867746700390563850/1085956782775603350/1085964155938414794

# 2023-03-17

If we solve first for the no-conflict merge case, then:
- **Prerequisite**: PRs must be created for every branch created
- **Prerequisite**: There are no unsaved changes to the file being translated
- When user selects a file for editing, then do the following:
	- Query the PR and find if the attribute "mergeable" is true or false
	- If true, then the master branch can be merged without any issues.
	- If false, then there are conflicts. The merge cannot be made and the user is notified that conflicts exist
*Tasks*
1. In the above, creating a PR for each branch created can be done independently of anything else.
2. Create the ui and component for querying the PR and performing no-conflict merge of master into their user branch.


## Creating a PR

QA DCS Swagger:
https://qa.door43.org/api/swagger#/repository/repoCreatePullRequest

Here is the POST body:

```
{
  "assignee": "cecil.new",
  "assignees": [
    "cecil.new"
  ],
  "base": "master",
  "body": "This is a PR for my work in Ruth",
  "due_date": "2023-03-17T12:50:55.437Z",
  "head": "gt-RUT-cecil.new",
  "labels": [
    0
  ],
  "milestone": 0,
  "title": "My PR by Cecil New"
}
```

Here is the curl:

```
curl -X 'POST' \
  'https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/pulls' \
  -H 'accept: application/json' \
  -H 'authorization: Basic Y2VjaWwubmV3OjM1XkhoRF5IRiRkKiNOKkE=' \
  -H 'Content-Type: application/json' \
  -d '{
  "assignee": "cecil.new",
  "assignees": [
    "cecil.new"
  ],
  "base": "master",
  "body": "This is a PR for my work in Ruth",
  "due_date": "2023-03-17T12:50:55.437Z",
  "head": "gt-RUT-cecil.new",
  "labels": [
    0
  ],
  "milestone": 0,
  "title": "My PR by Cecil New"
}'
```

Here is the link to the PR created:
https://qa.door43.org/unfoldingWord/en_ult/pulls/3346


The response body is in Appendix A. From the response body, `number` has the value `3346` which is the PR number. The URL is o.url (see above).

The properties of interest are:

```
"url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346",
"number": 3346,
"state": "open",
"html_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346",
"diff_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346.diff",
"patch_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346.patch",
"mergeable": true,
"merged": false,
```


## On Merging Master to User Branch

This link answers the question, has it been merged:
https://qa.door43.org/api/swagger#/repository/repoPullRequestIsMerged

If the PR has not been merged, then this URL:
https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/pulls/3346/merge

will return a 404 with this message:
```
{
  "errors": null,
  "message": "The target couldn't be found.",
  "url": "https://qa.door43.org/api/swagger"
}
```

Hopefully, this swagger API will update my branch from master:
https://qa.door43.org/api/swagger#/repository/repoUpdatePullRequest

Looks good... results:

First the URL:

```
https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/pulls/3346/update?style=merge
```

Next the curl for reference:

```
curl -X 'POST' \
  'https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/pulls/3346/update?style=merge' \
  -H 'accept: application/json' \
  -H 'authorization: Basic Y2VjaWwubmV3OjM1XkhoRF5IRiRkKiNOKkE=' \
  -d ''
```

The status returned was a 500; and message returned was:

```
{
  "message": "HeadBranch of PR 3346 is up to date",
  "url": "https://qa.door43.org/api/swagger"
}
```

Which is what I'd expect in this case since there no commits to master since I created the above PR. So I'll make a change and retry. I made a PR: `gt-DEU-superdav42 #3347` and merged it. Now there should be an update to make to my branch.

This time the status returned was a 200; and there was no message returned.


## On Merging User Branch into Master

This is the swagger for this:
https://qa.door43.org/api/swagger#/repository/repoMergePullRequest

The POST has the usual stuff, but in the body needs `MergeCommitID`. I think this the branch SHA value. Before I merged, that value can be found in Appendix A at this location:

```
  "head": {
    "label": "gt-RUT-cecil.new",
    "ref": "gt-RUT-cecil.new",
    "sha": "20c356f89d5c63655a002e1e3afd2c66ec927db7",
    "repo_id": 11419,
```

But since I merged master into my branch it is now different. See Appendix B for re-running to get the updated PR info. The new SHA is `0006fcc95e0d3f36090a6fefe6ad0abed15a36b3`.

Here is the body for the post for my first attempt:

```
{
  "Do": "merge",
  "MergeCommitID": "0006fcc95e0d3f36090a6fefe6ad0abed15a36b3",
  "MergeMessageField": "LGTM",
  "MergeTitleField": "this is the title, LGTM",
  "delete_branch_after_merge": false,
  "force_merge": false,
  "head_commit_id": "0006fcc95e0d3f36090a6fefe6ad0abed15a36b3",
  "merge_when_checks_succeed": false
}
```

Notice that there is a `head_commit_id`, so maybe `MergeCommitID` is supposed to be `master`??

Merge succeed: https://qa.door43.org/unfoldingWord/en_ult/pulls/3346


# Appendix A

The response body in full.

```
{
  "id": 10989,
  "url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346",
  "number": 3346,
  "user": {
    "id": 13993,
    "login": "cecil.new",
    "login_name": "",
    "full_name": "",
    "email": "cecil.new@noreply.door43.org",
    "avatar_url": "https://qa.door43.org/avatars/4906b5e03f108264e4c119874ce31c37",
    "language": "",
    "is_admin": false,
    "last_login": "0001-01-01T00:00:00Z",
    "created": "2019-09-19T13:37:42Z",
    "repo_languages": [
      "en"
    ],
    "repo_subjects": [
      "Aligned Bible"
    ],
    "repo_metadata_types": [
      "rc"
    ],
    "restricted": false,
    "active": false,
    "prohibit_login": false,
    "location": "",
    "website": "",
    "description": "",
    "visibility": "public",
    "followers_count": 0,
    "following_count": 0,
    "starred_repos_count": 0,
    "username": "cecil.new"
  },
  "title": "My PR by Cecil New",
  "body": "This is a PR for my work in Ruth",
  "labels": [],
  "milestone": null,
  "assignee": {
    "id": 13993,
    "login": "cecil.new",
    "login_name": "",
    "full_name": "",
    "email": "cecil.new@noreply.door43.org",
    "avatar_url": "https://qa.door43.org/avatars/4906b5e03f108264e4c119874ce31c37",
    "language": "",
    "is_admin": false,
    "last_login": "0001-01-01T00:00:00Z",
    "created": "2019-09-19T13:37:42Z",
    "repo_languages": [
      "en"
    ],
    "repo_subjects": [
      "Aligned Bible"
    ],
    "repo_metadata_types": [
      "rc"
    ],
    "restricted": false,
    "active": false,
    "prohibit_login": false,
    "location": "",
    "website": "",
    "description": "",
    "visibility": "public",
    "followers_count": 0,
    "following_count": 0,
    "starred_repos_count": 0,
    "username": "cecil.new"
  },
  "assignees": [
    {
      "id": 13993,
      "login": "cecil.new",
      "login_name": "",
      "full_name": "",
      "email": "cecil.new@noreply.door43.org",
      "avatar_url": "https://qa.door43.org/avatars/4906b5e03f108264e4c119874ce31c37",
      "language": "",
      "is_admin": false,
      "last_login": "0001-01-01T00:00:00Z",
      "created": "2019-09-19T13:37:42Z",
      "repo_languages": [
        "en"
      ],
      "repo_subjects": [
        "Aligned Bible"
      ],
      "repo_metadata_types": [
        "rc"
      ],
      "restricted": false,
      "active": false,
      "prohibit_login": false,
      "location": "",
      "website": "",
      "description": "",
      "visibility": "public",
      "followers_count": 0,
      "following_count": 0,
      "starred_repos_count": 0,
      "username": "cecil.new"
    }
  ],
  "state": "open",
  "is_locked": false,
  "comments": 0,
  "html_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346",
  "diff_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346.diff",
  "patch_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346.patch",
  "mergeable": true,
  "merged": false,
  "merged_at": null,
  "merge_commit_sha": null,
  "merged_by": null,
  "allow_maintainer_edit": false,
  "base": {
    "label": "master",
    "ref": "master",
    "sha": "efde31b9ae0dcf18004c3f34a92f6deee2dc75cb",
    "repo_id": 11419,
    "repo": {
      "id": 11419,
      "owner": {
        "id": 613,
        "login": "unfoldingWord",
        "login_name": "",
        "full_name": "unfoldingWord®",
        "email": "unfoldingword@noreply.door43.org",
        "avatar_url": "https://qa.door43.org/avatars/1bc81b740b4286613cdaa55ddfe4b1fc",
        "language": "",
        "is_admin": false,
        "last_login": "0001-01-01T00:00:00Z",
        "created": "2016-02-16T23:44:26Z",
        "repo_languages": [
          "el-x-koine",
          "en",
          "fr",
          "hbo"
        ],
        "repo_subjects": [
          "Aligned Bible",
          "Aramaic Grammar",
          "Bible",
          "Greek Grammar",
          "Greek Lexicon",
          "Greek New Testament",
          "Hebrew Grammar",
          "Hebrew Old Testament",
          "OBS Study Questions",
          "OBS Translation Notes",
          "OBS Translation Questions",
          "Open Bible Stories",
          "Study Notes",
          "Training Library",
          "Translation Academy",
          "Translation Words",
          "TSV OBS Study Notes",
          "TSV OBS Study Questions",
          "TSV OBS Translation Notes",
          "TSV OBS Translation Questions",
          "TSV OBS Translation Words Links",
          "TSV Study Notes",
          "TSV Study Questions",
          "TSV Translation Notes",
          "TSV Translation Questions",
          "TSV Translation Words Links"
        ],
        "repo_metadata_types": [
          "rc"
        ],
        "restricted": false,
        "active": false,
        "prohibit_login": false,
        "location": "",
        "website": "https://unfoldingword.org",
        "description": "",
        "visibility": "public",
        "followers_count": 0,
        "following_count": 0,
        "starred_repos_count": 0,
        "username": "unfoldingWord"
      },
      "name": "en_ult",
      "full_name": "unfoldingWord/en_ult",
      "description": "unfoldingWord® Literal Text (formerly ULB)",
      "empty": false,
      "private": false,
      "fork": false,
      "template": false,
      "parent": null,
      "mirror": false,
      "size": 399747,
      "languages_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/languages",
      "html_url": "https://qa.door43.org/unfoldingWord/en_ult",
      "ssh_url": "git@qa.door43.org:unfoldingWord/en_ult.git",
      "clone_url": "https://qa.door43.org/unfoldingWord/en_ult.git",
      "original_url": "",
      "website": "https://www.unfoldingword.org/ult",
      "stars_count": 9,
      "forks_count": 12,
      "watchers_count": 3,
      "open_issues_count": 52,
      "open_pr_counter": 1,
      "release_counter": 46,
      "default_branch": "master",
      "archived": false,
      "created_at": "2017-06-01T22:16:16Z",
      "updated_at": "2023-03-13T12:53:53Z",
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
      "allow_rebase": false,
      "allow_rebase_explicit": false,
      "allow_squash_merge": true,
      "allow_rebase_update": true,
      "default_delete_branch_after_merge": false,
      "default_merge_style": "merge",
      "avatar_url": "https://qa.door43.org/repo-avatars/11419-ba04b7a1942e0a5ba20adc6ac9372799",
      "internal": false,
      "mirror_interval": "",
      "mirror_updated": "0001-01-01T00:00:00Z",
      "repo_transfer": null,
      "metadata_type": "rc",
      "metadata_version": "0.2",
      "language": "en",
      "language_title": "English",
      "language_direction": "ltr",
      "language_is_gl": true,
      "subject": "Aligned Bible",
      "title": "unfoldingWord® Literal Text",
      "ingredients": [
        {
          "categories": null,
          "identifier": "frt",
          "path": "./A0-FRT.usfm",
          "sort": 0,
          "title": "Front Matter",
          "versification": "ufw",
          "alignment_count": 0
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "gen",
          "path": "./01-GEN.usfm",
          "sort": 1,
          "title": "Genesis",
          "versification": "ufw",
          "alignment_count": 23024
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "exo",
          "path": "./02-EXO.usfm",
          "sort": 2,
          "title": "Exodus",
          "versification": "ufw",
          "alignment_count": 18390
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "lev",
          "path": "./03-LEV.usfm",
          "sort": 3,
          "title": "Leviticus",
          "versification": "ufw",
          "alignment_count": 13293
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "num",
          "path": "./04-NUM.usfm",
          "sort": 4,
          "title": "Numbers",
          "versification": "ufw",
          "alignment_count": 17615
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "deu",
          "path": "./05-DEU.usfm",
          "sort": 5,
          "title": "Deuteronomy",
          "versification": "ufw",
          "alignment_count": 15522
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jos",
          "path": "./06-JOS.usfm",
          "sort": 6,
          "title": "Joshua",
          "versification": "ufw",
          "alignment_count": 10741
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jdg",
          "path": "./07-JDG.usfm",
          "sort": 7,
          "title": "Judges",
          "versification": "ufw",
          "alignment_count": 10614
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "rut",
          "path": "./08-RUT.usfm",
          "sort": 8,
          "title": "Ruth",
          "versification": "ufw",
          "alignment_count": 1424
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1sa",
          "path": "./09-1SA.usfm",
          "sort": 9,
          "title": "1 Samuel",
          "versification": "ufw",
          "alignment_count": 14629
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2sa",
          "path": "./10-2SA.usfm",
          "sort": 10,
          "title": "2 Samuel",
          "versification": "ufw",
          "alignment_count": 12416
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1ki",
          "path": "./11-1KI.usfm",
          "sort": 11,
          "title": "1 Kings",
          "versification": "ufw",
          "alignment_count": 14237
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2ki",
          "path": "./12-2KI.usfm",
          "sort": 12,
          "title": "2 Kings",
          "versification": "ufw",
          "alignment_count": 13308
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1ch",
          "path": "./13-1CH.usfm",
          "sort": 13,
          "title": "1 Chronicles",
          "versification": "ufw",
          "alignment_count": 11226
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2ch",
          "path": "./14-2CH.usfm",
          "sort": 14,
          "title": "2 Chronicles",
          "versification": "ufw",
          "alignment_count": 14282
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ezr",
          "path": "./15-EZR.usfm",
          "sort": 15,
          "title": "Ezra",
          "versification": "ufw",
          "alignment_count": 3989
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "neh",
          "path": "./16-NEH.usfm",
          "sort": 16,
          "title": "Nehemiah",
          "versification": "ufw",
          "alignment_count": 5745
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "est",
          "path": "./17-EST.usfm",
          "sort": 17,
          "title": "Esther",
          "versification": "ufw",
          "alignment_count": 3364
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "job",
          "path": "./18-JOB.usfm",
          "sort": 18,
          "title": "Job",
          "versification": "ufw",
          "alignment_count": 8649
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "psa",
          "path": "./19-PSA.usfm",
          "sort": 19,
          "title": "Psalms",
          "versification": "ufw",
          "alignment_count": 19943
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "pro",
          "path": "./20-PRO.usfm",
          "sort": 20,
          "title": "Proverbs",
          "versification": "ufw",
          "alignment_count": 7585
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ecc",
          "path": "./21-ECC.usfm",
          "sort": 21,
          "title": "Ecclesiastes",
          "versification": "ufw",
          "alignment_count": 3312
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "sng",
          "path": "./22-SNG.usfm",
          "sort": 22,
          "title": "Song of Solomon",
          "versification": "ufw",
          "alignment_count": 1374
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "isa",
          "path": "./23-ISA.usfm",
          "sort": 23,
          "title": "Isaiah",
          "versification": "ufw",
          "alignment_count": 5893
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jer",
          "path": "./24-JER.usfm",
          "sort": 24,
          "title": "Jeremiah",
          "versification": "ufw",
          "alignment_count": 8336
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "lam",
          "path": "./25-LAM.usfm",
          "sort": 25,
          "title": "Lamentations",
          "versification": "ufw",
          "alignment_count": 1622
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ezk",
          "path": "./26-EZK.usfm",
          "sort": 26,
          "title": "Ezekiel",
          "versification": "ufw",
          "alignment_count": 19166
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "dan",
          "path": "./27-DAN.usfm",
          "sort": 27,
          "title": "Daniel",
          "versification": "ufw",
          "alignment_count": 6445
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hos",
          "path": "./28-HOS.usfm",
          "sort": 28,
          "title": "Hosea",
          "versification": "ufw",
          "alignment_count": 2563
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jol",
          "path": "./29-JOL.usfm",
          "sort": 29,
          "title": "Joel",
          "versification": "ufw",
          "alignment_count": 1024
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "amo",
          "path": "./30-AMO.usfm",
          "sort": 30,
          "title": "Amos",
          "versification": "ufw",
          "alignment_count": 2247
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "oba",
          "path": "./31-OBA.usfm",
          "sort": 31,
          "title": "Obadiah",
          "versification": "ufw",
          "alignment_count": 325
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jon",
          "path": "./32-JON.usfm",
          "sort": 32,
          "title": "Jonah",
          "versification": "ufw",
          "alignment_count": 763
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "mic",
          "path": "./33-MIC.usfm",
          "sort": 33,
          "title": "Micah",
          "versification": "ufw",
          "alignment_count": 709
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "nam",
          "path": "./34-NAM.usfm",
          "sort": 34,
          "title": "Nahum",
          "versification": "ufw",
          "alignment_count": 592
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hab",
          "path": "./35-HAB.usfm",
          "sort": 35,
          "title": "Habakkuk",
          "versification": "ufw",
          "alignment_count": 737
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "zep",
          "path": "./36-ZEP.usfm",
          "sort": 36,
          "title": "Zephaniah",
          "versification": "ufw",
          "alignment_count": 826
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hag",
          "path": "./37-HAG.usfm",
          "sort": 37,
          "title": "Haggai",
          "versification": "ufw",
          "alignment_count": 644
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "zec",
          "path": "./38-ZEC.usfm",
          "sort": 38,
          "title": "Zechariah",
          "versification": "ufw",
          "alignment_count": 3227
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "mal",
          "path": "./39-MAL.usfm",
          "sort": 39,
          "title": "Malachi",
          "versification": "ufw",
          "alignment_count": 948
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "mat",
          "path": "./41-MAT.usfm",
          "sort": 41,
          "title": "Matthew",
          "versification": "ufw",
          "alignment_count": 19233
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "mrk",
          "path": "./42-MRK.usfm",
          "sort": 42,
          "title": "Mark",
          "versification": "ufw",
          "alignment_count": 11830
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "luk",
          "path": "./43-LUK.usfm",
          "sort": 43,
          "title": "Luke",
          "versification": "ufw",
          "alignment_count": 20345
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jhn",
          "path": "./44-JHN.usfm",
          "sort": 44,
          "title": "John",
          "versification": "ufw",
          "alignment_count": 16236
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "act",
          "path": "./45-ACT.usfm",
          "sort": 45,
          "title": "Acts",
          "versification": "ufw",
          "alignment_count": 19271
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "rom",
          "path": "./46-ROM.usfm",
          "sort": 46,
          "title": "Romans",
          "versification": "ufw",
          "alignment_count": 7609
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1co",
          "path": "./47-1CO.usfm",
          "sort": 47,
          "title": "1 Corinthians",
          "versification": "ufw",
          "alignment_count": 7262
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2co",
          "path": "./48-2CO.usfm",
          "sort": 48,
          "title": "2 Corinthians",
          "versification": "ufw",
          "alignment_count": 4845
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "gal",
          "path": "./49-GAL.usfm",
          "sort": 49,
          "title": "Galatians",
          "versification": "ufw",
          "alignment_count": 2334
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "eph",
          "path": "./50-EPH.usfm",
          "sort": 50,
          "title": "Ephesians",
          "versification": "ufw",
          "alignment_count": 2565
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "php",
          "path": "./51-PHP.usfm",
          "sort": 51,
          "title": "Philippians",
          "versification": "ufw",
          "alignment_count": 1743
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "col",
          "path": "./52-COL.usfm",
          "sort": 52,
          "title": "Colossians",
          "versification": "ufw",
          "alignment_count": 1671
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1th",
          "path": "./53-1TH.usfm",
          "sort": 53,
          "title": "1 Thessalonians",
          "versification": "ufw",
          "alignment_count": 1549
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2th",
          "path": "./54-2TH.usfm",
          "sort": 54,
          "title": "2 Thessalonians",
          "versification": "ufw",
          "alignment_count": 875
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1ti",
          "path": "./55-1TI.usfm",
          "sort": 55,
          "title": "1 Timothy",
          "versification": "ufw",
          "alignment_count": 1699
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2ti",
          "path": "./56-2TI.usfm",
          "sort": 56,
          "title": "2 Timothy",
          "versification": "ufw",
          "alignment_count": 1310
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "tit",
          "path": "./57-TIT.usfm",
          "sort": 57,
          "title": "Titus",
          "versification": "ufw",
          "alignment_count": 716
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "phm",
          "path": "./58-PHM.usfm",
          "sort": 58,
          "title": "Philemon",
          "versification": "ufw",
          "alignment_count": 361
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "heb",
          "path": "./59-HEB.usfm",
          "sort": 59,
          "title": "Hebrews",
          "versification": "ufw",
          "alignment_count": 5316
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jas",
          "path": "./60-JAS.usfm",
          "sort": 60,
          "title": "James",
          "versification": "ufw",
          "alignment_count": 1866
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1pe",
          "path": "./61-1PE.usfm",
          "sort": 61,
          "title": "1 Peter",
          "versification": "ufw",
          "alignment_count": 1825
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2pe",
          "path": "./62-2PE.usfm",
          "sort": 62,
          "title": "2 Peter",
          "versification": "ufw",
          "alignment_count": 1214
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1jn",
          "path": "./63-1JN.usfm",
          "sort": 63,
          "title": "1 John",
          "versification": "ufw",
          "alignment_count": 2207
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2jn",
          "path": "./64-2JN.usfm",
          "sort": 64,
          "title": "2 John",
          "versification": "ufw",
          "alignment_count": 261
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "3jn",
          "path": "./65-3JN.usfm",
          "sort": 65,
          "title": "3 John",
          "versification": "ufw",
          "alignment_count": 243
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jud",
          "path": "./66-JUD.usfm",
          "sort": 66,
          "title": "Jude",
          "versification": "ufw",
          "alignment_count": 487
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "rev",
          "path": "./67-REV.usfm",
          "sort": 67,
          "title": "Revelation",
          "versification": "ufw",
          "alignment_count": 10349
        }
      ],
      "checking_level": 3,
      "catalog": {
        "prod": {
          "branch_or_tag_name": "v45",
          "release_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/releases/3441599",
          "released": "2023-02-11T15:48:55Z",
          "zipball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/v45.zip",
          "tarball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/v45.tar.gz",
          "git_trees_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/git/trees/v45?recursive=1&per_page=99999",
          "contents_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/contents?ref=v45"
        },
        "preprod": null,
        "draft": null,
        "latest": {
          "branch_or_tag_name": "master",
          "release_url": null,
          "released": "2023-03-11T07:50:44Z",
          "zipball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/master.zip",
          "tarball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/master.tar.gz",
          "git_trees_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/git/trees/master?recursive=1&per_page=99999",
          "contents_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/contents?ref=master"
        }
      },
      "content_format": "usfm"
    }
  },
  "head": {
    "label": "gt-RUT-cecil.new",
    "ref": "gt-RUT-cecil.new",
    "sha": "20c356f89d5c63655a002e1e3afd2c66ec927db7",
    "repo_id": 11419,
    "repo": {
      "id": 11419,
      "owner": {
        "id": 613,
        "login": "unfoldingWord",
        "login_name": "",
        "full_name": "unfoldingWord®",
        "email": "unfoldingword@noreply.door43.org",
        "avatar_url": "https://qa.door43.org/avatars/1bc81b740b4286613cdaa55ddfe4b1fc",
        "language": "",
        "is_admin": false,
        "last_login": "0001-01-01T00:00:00Z",
        "created": "2016-02-16T23:44:26Z",
        "repo_languages": [
          "el-x-koine",
          "en",
          "fr",
          "hbo"
        ],
        "repo_subjects": [
          "Aligned Bible",
          "Aramaic Grammar",
          "Bible",
          "Greek Grammar",
          "Greek Lexicon",
          "Greek New Testament",
          "Hebrew Grammar",
          "Hebrew Old Testament",
          "OBS Study Questions",
          "OBS Translation Notes",
          "OBS Translation Questions",
          "Open Bible Stories",
          "Study Notes",
          "Training Library",
          "Translation Academy",
          "Translation Words",
          "TSV OBS Study Notes",
          "TSV OBS Study Questions",
          "TSV OBS Translation Notes",
          "TSV OBS Translation Questions",
          "TSV OBS Translation Words Links",
          "TSV Study Notes",
          "TSV Study Questions",
          "TSV Translation Notes",
          "TSV Translation Questions",
          "TSV Translation Words Links"
        ],
        "repo_metadata_types": [
          "rc"
        ],
        "restricted": false,
        "active": false,
        "prohibit_login": false,
        "location": "",
        "website": "https://unfoldingword.org",
        "description": "",
        "visibility": "public",
        "followers_count": 0,
        "following_count": 0,
        "starred_repos_count": 0,
        "username": "unfoldingWord"
      },
      "name": "en_ult",
      "full_name": "unfoldingWord/en_ult",
      "description": "unfoldingWord® Literal Text (formerly ULB)",
      "empty": false,
      "private": false,
      "fork": false,
      "template": false,
      "parent": null,
      "mirror": false,
      "size": 399747,
      "languages_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/languages",
      "html_url": "https://qa.door43.org/unfoldingWord/en_ult",
      "ssh_url": "git@qa.door43.org:unfoldingWord/en_ult.git",
      "clone_url": "https://qa.door43.org/unfoldingWord/en_ult.git",
      "original_url": "",
      "website": "https://www.unfoldingword.org/ult",
      "stars_count": 9,
      "forks_count": 12,
      "watchers_count": 3,
      "open_issues_count": 52,
      "open_pr_counter": 1,
      "release_counter": 46,
      "default_branch": "master",
      "archived": false,
      "created_at": "2017-06-01T22:16:16Z",
      "updated_at": "2023-03-13T12:53:53Z",
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
      "allow_rebase": false,
      "allow_rebase_explicit": false,
      "allow_squash_merge": true,
      "allow_rebase_update": true,
      "default_delete_branch_after_merge": false,
      "default_merge_style": "merge",
      "avatar_url": "https://qa.door43.org/repo-avatars/11419-ba04b7a1942e0a5ba20adc6ac9372799",
      "internal": false,
      "mirror_interval": "",
      "mirror_updated": "0001-01-01T00:00:00Z",
      "repo_transfer": null,
      "metadata_type": "rc",
      "metadata_version": "0.2",
      "language": "en",
      "language_title": "English",
      "language_direction": "ltr",
      "language_is_gl": true,
      "subject": "Aligned Bible",
      "title": "unfoldingWord® Literal Text",
      "ingredients": [
        {
          "categories": null,
          "identifier": "frt",
          "path": "./A0-FRT.usfm",
          "sort": 0,
          "title": "Front Matter",
          "versification": "ufw",
          "alignment_count": 0
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "gen",
          "path": "./01-GEN.usfm",
          "sort": 1,
          "title": "Genesis",
          "versification": "ufw",
          "alignment_count": 23024
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "exo",
          "path": "./02-EXO.usfm",
          "sort": 2,
          "title": "Exodus",
          "versification": "ufw",
          "alignment_count": 18390
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "lev",
          "path": "./03-LEV.usfm",
          "sort": 3,
          "title": "Leviticus",
          "versification": "ufw",
          "alignment_count": 13293
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "num",
          "path": "./04-NUM.usfm",
          "sort": 4,
          "title": "Numbers",
          "versification": "ufw",
          "alignment_count": 17615
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "deu",
          "path": "./05-DEU.usfm",
          "sort": 5,
          "title": "Deuteronomy",
          "versification": "ufw",
          "alignment_count": 15522
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jos",
          "path": "./06-JOS.usfm",
          "sort": 6,
          "title": "Joshua",
          "versification": "ufw",
          "alignment_count": 10741
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jdg",
          "path": "./07-JDG.usfm",
          "sort": 7,
          "title": "Judges",
          "versification": "ufw",
          "alignment_count": 10614
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "rut",
          "path": "./08-RUT.usfm",
          "sort": 8,
          "title": "Ruth",
          "versification": "ufw",
          "alignment_count": 1424
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1sa",
          "path": "./09-1SA.usfm",
          "sort": 9,
          "title": "1 Samuel",
          "versification": "ufw",
          "alignment_count": 14629
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2sa",
          "path": "./10-2SA.usfm",
          "sort": 10,
          "title": "2 Samuel",
          "versification": "ufw",
          "alignment_count": 12416
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1ki",
          "path": "./11-1KI.usfm",
          "sort": 11,
          "title": "1 Kings",
          "versification": "ufw",
          "alignment_count": 14237
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2ki",
          "path": "./12-2KI.usfm",
          "sort": 12,
          "title": "2 Kings",
          "versification": "ufw",
          "alignment_count": 13308
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1ch",
          "path": "./13-1CH.usfm",
          "sort": 13,
          "title": "1 Chronicles",
          "versification": "ufw",
          "alignment_count": 11226
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2ch",
          "path": "./14-2CH.usfm",
          "sort": 14,
          "title": "2 Chronicles",
          "versification": "ufw",
          "alignment_count": 14282
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ezr",
          "path": "./15-EZR.usfm",
          "sort": 15,
          "title": "Ezra",
          "versification": "ufw",
          "alignment_count": 3989
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "neh",
          "path": "./16-NEH.usfm",
          "sort": 16,
          "title": "Nehemiah",
          "versification": "ufw",
          "alignment_count": 5745
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "est",
          "path": "./17-EST.usfm",
          "sort": 17,
          "title": "Esther",
          "versification": "ufw",
          "alignment_count": 3364
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "job",
          "path": "./18-JOB.usfm",
          "sort": 18,
          "title": "Job",
          "versification": "ufw",
          "alignment_count": 8649
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "psa",
          "path": "./19-PSA.usfm",
          "sort": 19,
          "title": "Psalms",
          "versification": "ufw",
          "alignment_count": 19943
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "pro",
          "path": "./20-PRO.usfm",
          "sort": 20,
          "title": "Proverbs",
          "versification": "ufw",
          "alignment_count": 7585
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ecc",
          "path": "./21-ECC.usfm",
          "sort": 21,
          "title": "Ecclesiastes",
          "versification": "ufw",
          "alignment_count": 3312
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "sng",
          "path": "./22-SNG.usfm",
          "sort": 22,
          "title": "Song of Solomon",
          "versification": "ufw",
          "alignment_count": 1374
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "isa",
          "path": "./23-ISA.usfm",
          "sort": 23,
          "title": "Isaiah",
          "versification": "ufw",
          "alignment_count": 5893
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jer",
          "path": "./24-JER.usfm",
          "sort": 24,
          "title": "Jeremiah",
          "versification": "ufw",
          "alignment_count": 8336
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "lam",
          "path": "./25-LAM.usfm",
          "sort": 25,
          "title": "Lamentations",
          "versification": "ufw",
          "alignment_count": 1622
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ezk",
          "path": "./26-EZK.usfm",
          "sort": 26,
          "title": "Ezekiel",
          "versification": "ufw",
          "alignment_count": 19166
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "dan",
          "path": "./27-DAN.usfm",
          "sort": 27,
          "title": "Daniel",
          "versification": "ufw",
          "alignment_count": 6445
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hos",
          "path": "./28-HOS.usfm",
          "sort": 28,
          "title": "Hosea",
          "versification": "ufw",
          "alignment_count": 2563
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jol",
          "path": "./29-JOL.usfm",
          "sort": 29,
          "title": "Joel",
          "versification": "ufw",
          "alignment_count": 1024
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "amo",
          "path": "./30-AMO.usfm",
          "sort": 30,
          "title": "Amos",
          "versification": "ufw",
          "alignment_count": 2247
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "oba",
          "path": "./31-OBA.usfm",
          "sort": 31,
          "title": "Obadiah",
          "versification": "ufw",
          "alignment_count": 325
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jon",
          "path": "./32-JON.usfm",
          "sort": 32,
          "title": "Jonah",
          "versification": "ufw",
          "alignment_count": 763
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "mic",
          "path": "./33-MIC.usfm",
          "sort": 33,
          "title": "Micah",
          "versification": "ufw",
          "alignment_count": 709
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "nam",
          "path": "./34-NAM.usfm",
          "sort": 34,
          "title": "Nahum",
          "versification": "ufw",
          "alignment_count": 592
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hab",
          "path": "./35-HAB.usfm",
          "sort": 35,
          "title": "Habakkuk",
          "versification": "ufw",
          "alignment_count": 737
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "zep",
          "path": "./36-ZEP.usfm",
          "sort": 36,
          "title": "Zephaniah",
          "versification": "ufw",
          "alignment_count": 826
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hag",
          "path": "./37-HAG.usfm",
          "sort": 37,
          "title": "Haggai",
          "versification": "ufw",
          "alignment_count": 644
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "zec",
          "path": "./38-ZEC.usfm",
          "sort": 38,
          "title": "Zechariah",
          "versification": "ufw",
          "alignment_count": 3227
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "mal",
          "path": "./39-MAL.usfm",
          "sort": 39,
          "title": "Malachi",
          "versification": "ufw",
          "alignment_count": 948
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "mat",
          "path": "./41-MAT.usfm",
          "sort": 41,
          "title": "Matthew",
          "versification": "ufw",
          "alignment_count": 19233
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "mrk",
          "path": "./42-MRK.usfm",
          "sort": 42,
          "title": "Mark",
          "versification": "ufw",
          "alignment_count": 11830
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "luk",
          "path": "./43-LUK.usfm",
          "sort": 43,
          "title": "Luke",
          "versification": "ufw",
          "alignment_count": 20345
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jhn",
          "path": "./44-JHN.usfm",
          "sort": 44,
          "title": "John",
          "versification": "ufw",
          "alignment_count": 16236
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "act",
          "path": "./45-ACT.usfm",
          "sort": 45,
          "title": "Acts",
          "versification": "ufw",
          "alignment_count": 19271
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "rom",
          "path": "./46-ROM.usfm",
          "sort": 46,
          "title": "Romans",
          "versification": "ufw",
          "alignment_count": 7609
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1co",
          "path": "./47-1CO.usfm",
          "sort": 47,
          "title": "1 Corinthians",
          "versification": "ufw",
          "alignment_count": 7262
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2co",
          "path": "./48-2CO.usfm",
          "sort": 48,
          "title": "2 Corinthians",
          "versification": "ufw",
          "alignment_count": 4845
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "gal",
          "path": "./49-GAL.usfm",
          "sort": 49,
          "title": "Galatians",
          "versification": "ufw",
          "alignment_count": 2334
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "eph",
          "path": "./50-EPH.usfm",
          "sort": 50,
          "title": "Ephesians",
          "versification": "ufw",
          "alignment_count": 2565
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "php",
          "path": "./51-PHP.usfm",
          "sort": 51,
          "title": "Philippians",
          "versification": "ufw",
          "alignment_count": 1743
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "col",
          "path": "./52-COL.usfm",
          "sort": 52,
          "title": "Colossians",
          "versification": "ufw",
          "alignment_count": 1671
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1th",
          "path": "./53-1TH.usfm",
          "sort": 53,
          "title": "1 Thessalonians",
          "versification": "ufw",
          "alignment_count": 1549
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2th",
          "path": "./54-2TH.usfm",
          "sort": 54,
          "title": "2 Thessalonians",
          "versification": "ufw",
          "alignment_count": 875
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1ti",
          "path": "./55-1TI.usfm",
          "sort": 55,
          "title": "1 Timothy",
          "versification": "ufw",
          "alignment_count": 1699
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2ti",
          "path": "./56-2TI.usfm",
          "sort": 56,
          "title": "2 Timothy",
          "versification": "ufw",
          "alignment_count": 1310
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "tit",
          "path": "./57-TIT.usfm",
          "sort": 57,
          "title": "Titus",
          "versification": "ufw",
          "alignment_count": 716
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "phm",
          "path": "./58-PHM.usfm",
          "sort": 58,
          "title": "Philemon",
          "versification": "ufw",
          "alignment_count": 361
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "heb",
          "path": "./59-HEB.usfm",
          "sort": 59,
          "title": "Hebrews",
          "versification": "ufw",
          "alignment_count": 5316
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jas",
          "path": "./60-JAS.usfm",
          "sort": 60,
          "title": "James",
          "versification": "ufw",
          "alignment_count": 1866
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1pe",
          "path": "./61-1PE.usfm",
          "sort": 61,
          "title": "1 Peter",
          "versification": "ufw",
          "alignment_count": 1825
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2pe",
          "path": "./62-2PE.usfm",
          "sort": 62,
          "title": "2 Peter",
          "versification": "ufw",
          "alignment_count": 1214
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1jn",
          "path": "./63-1JN.usfm",
          "sort": 63,
          "title": "1 John",
          "versification": "ufw",
          "alignment_count": 2207
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2jn",
          "path": "./64-2JN.usfm",
          "sort": 64,
          "title": "2 John",
          "versification": "ufw",
          "alignment_count": 261
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "3jn",
          "path": "./65-3JN.usfm",
          "sort": 65,
          "title": "3 John",
          "versification": "ufw",
          "alignment_count": 243
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jud",
          "path": "./66-JUD.usfm",
          "sort": 66,
          "title": "Jude",
          "versification": "ufw",
          "alignment_count": 487
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "rev",
          "path": "./67-REV.usfm",
          "sort": 67,
          "title": "Revelation",
          "versification": "ufw",
          "alignment_count": 10349
        }
      ],
      "checking_level": 3,
      "catalog": {
        "prod": {
          "branch_or_tag_name": "v45",
          "release_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/releases/3441599",
          "released": "2023-02-11T15:48:55Z",
          "zipball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/v45.zip",
          "tarball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/v45.tar.gz",
          "git_trees_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/git/trees/v45?recursive=1&per_page=99999",
          "contents_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/contents?ref=v45"
        },
        "preprod": null,
        "draft": null,
        "latest": {
          "branch_or_tag_name": "master",
          "release_url": null,
          "released": "2023-03-11T07:50:44Z",
          "zipball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/master.zip",
          "tarball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/master.tar.gz",
          "git_trees_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/git/trees/master?recursive=1&per_page=99999",
          "contents_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/contents?ref=master"
        }
      },
      "content_format": "usfm"
    }
  },
  "merge_base": "efde31b9ae0dcf18004c3f34a92f6deee2dc75cb",
  "due_date": "2023-03-17T12:50:55Z",
  "created_at": "2023-03-17T12:55:30Z",
  "updated_at": "2023-03-17T12:55:32Z",
  "closed_at": null
}
```



# Appendix B - updated PR info

Swagger needed:
https://qa.door43.org/api/swagger#/repository/repoGetPullRequest

The curl:

```
curl -X 'GET' \
  'https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/pulls/3346' \
  -H 'accept: application/json' \
  -H 'authorization: Basic Y2VjaWwubmV3OjM1XkhoRF5IRiRkKiNOKkE='
```

The URL:
https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/pulls/3346

The relevant part from the reponse:

```
  "head": {
    "label": "gt-RUT-cecil.new",
    "ref": "gt-RUT-cecil.new",
    "sha": "0006fcc95e0d3f36090a6fefe6ad0abed15a36b3",
    "repo_id": 11419,
```

The full response:

```
{
  "id": 10989,
  "url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346",
  "number": 3346,
  "user": {
    "id": 13993,
    "login": "cecil.new",
    "login_name": "",
    "full_name": "",
    "email": "cecil.new@noreply.door43.org",
    "avatar_url": "https://qa.door43.org/avatars/4906b5e03f108264e4c119874ce31c37",
    "language": "",
    "is_admin": false,
    "last_login": "0001-01-01T00:00:00Z",
    "created": "2019-09-19T13:37:42Z",
    "repo_languages": [
      "en"
    ],
    "repo_subjects": [
      "Aligned Bible"
    ],
    "repo_metadata_types": [
      "rc"
    ],
    "restricted": false,
    "active": false,
    "prohibit_login": false,
    "location": "",
    "website": "",
    "description": "",
    "visibility": "public",
    "followers_count": 0,
    "following_count": 0,
    "starred_repos_count": 0,
    "username": "cecil.new"
  },
  "title": "My PR by Cecil New",
  "body": "This is a PR for my work in Ruth",
  "labels": [],
  "milestone": null,
  "assignee": {
    "id": 13993,
    "login": "cecil.new",
    "login_name": "",
    "full_name": "",
    "email": "cecil.new@noreply.door43.org",
    "avatar_url": "https://qa.door43.org/avatars/4906b5e03f108264e4c119874ce31c37",
    "language": "",
    "is_admin": false,
    "last_login": "0001-01-01T00:00:00Z",
    "created": "2019-09-19T13:37:42Z",
    "repo_languages": [
      "en"
    ],
    "repo_subjects": [
      "Aligned Bible"
    ],
    "repo_metadata_types": [
      "rc"
    ],
    "restricted": false,
    "active": false,
    "prohibit_login": false,
    "location": "",
    "website": "",
    "description": "",
    "visibility": "public",
    "followers_count": 0,
    "following_count": 0,
    "starred_repos_count": 0,
    "username": "cecil.new"
  },
  "assignees": [
    {
      "id": 13993,
      "login": "cecil.new",
      "login_name": "",
      "full_name": "",
      "email": "cecil.new@noreply.door43.org",
      "avatar_url": "https://qa.door43.org/avatars/4906b5e03f108264e4c119874ce31c37",
      "language": "",
      "is_admin": false,
      "last_login": "0001-01-01T00:00:00Z",
      "created": "2019-09-19T13:37:42Z",
      "repo_languages": [
        "en"
      ],
      "repo_subjects": [
        "Aligned Bible"
      ],
      "repo_metadata_types": [
        "rc"
      ],
      "restricted": false,
      "active": false,
      "prohibit_login": false,
      "location": "",
      "website": "",
      "description": "",
      "visibility": "public",
      "followers_count": 0,
      "following_count": 0,
      "starred_repos_count": 0,
      "username": "cecil.new"
    }
  ],
  "state": "open",
  "is_locked": false,
  "comments": 0,
  "html_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346",
  "diff_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346.diff",
  "patch_url": "https://qa.door43.org/unfoldingWord/en_ult/pulls/3346.patch",
  "mergeable": true,
  "merged": false,
  "merged_at": null,
  "merge_commit_sha": null,
  "merged_by": null,
  "allow_maintainer_edit": false,
  "base": {
    "label": "master",
    "ref": "master",
    "sha": "77747d060f9bfe97ab0b220683603a1a9188668a",
    "repo_id": 11419,
    "repo": {
      "id": 11419,
      "owner": {
        "id": 613,
        "login": "unfoldingWord",
        "login_name": "",
        "full_name": "unfoldingWord®",
        "email": "unfoldingword@noreply.door43.org",
        "avatar_url": "https://qa.door43.org/avatars/1bc81b740b4286613cdaa55ddfe4b1fc",
        "language": "",
        "is_admin": false,
        "last_login": "0001-01-01T00:00:00Z",
        "created": "2016-02-16T23:44:26Z",
        "repo_languages": [
          "el-x-koine",
          "en",
          "fr",
          "hbo"
        ],
        "repo_subjects": [
          "Aligned Bible",
          "Aramaic Grammar",
          "Bible",
          "Greek Grammar",
          "Greek Lexicon",
          "Greek New Testament",
          "Hebrew Grammar",
          "Hebrew Old Testament",
          "OBS Study Questions",
          "OBS Translation Notes",
          "OBS Translation Questions",
          "Open Bible Stories",
          "Study Notes",
          "Training Library",
          "Translation Academy",
          "Translation Words",
          "TSV OBS Study Notes",
          "TSV OBS Study Questions",
          "TSV OBS Translation Notes",
          "TSV OBS Translation Questions",
          "TSV OBS Translation Words Links",
          "TSV Study Notes",
          "TSV Study Questions",
          "TSV Translation Notes",
          "TSV Translation Questions",
          "TSV Translation Words Links"
        ],
        "repo_metadata_types": [
          "rc"
        ],
        "restricted": false,
        "active": false,
        "prohibit_login": false,
        "location": "",
        "website": "https://unfoldingword.org",
        "description": "",
        "visibility": "public",
        "followers_count": 0,
        "following_count": 0,
        "starred_repos_count": 0,
        "username": "unfoldingWord"
      },
      "name": "en_ult",
      "full_name": "unfoldingWord/en_ult",
      "description": "unfoldingWord® Literal Text (formerly ULB)",
      "empty": false,
      "private": false,
      "fork": false,
      "template": false,
      "parent": null,
      "mirror": false,
      "size": 399752,
      "languages_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/languages",
      "html_url": "https://qa.door43.org/unfoldingWord/en_ult",
      "ssh_url": "git@qa.door43.org:unfoldingWord/en_ult.git",
      "clone_url": "https://qa.door43.org/unfoldingWord/en_ult.git",
      "original_url": "",
      "website": "https://www.unfoldingword.org/ult",
      "stars_count": 9,
      "forks_count": 12,
      "watchers_count": 3,
      "open_issues_count": 52,
      "open_pr_counter": 2,
      "release_counter": 46,
      "default_branch": "master",
      "archived": false,
      "created_at": "2017-06-01T22:16:16Z",
      "updated_at": "2023-03-17T13:39:35Z",
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
      "allow_rebase": false,
      "allow_rebase_explicit": false,
      "allow_squash_merge": true,
      "allow_rebase_update": true,
      "default_delete_branch_after_merge": false,
      "default_merge_style": "merge",
      "avatar_url": "https://qa.door43.org/repo-avatars/11419-ba04b7a1942e0a5ba20adc6ac9372799",
      "internal": false,
      "mirror_interval": "",
      "mirror_updated": "0001-01-01T00:00:00Z",
      "repo_transfer": null,
      "metadata_type": "rc",
      "metadata_version": "0.2",
      "language": "en",
      "language_title": "English",
      "language_direction": "ltr",
      "language_is_gl": true,
      "subject": "Aligned Bible",
      "title": "unfoldingWord® Literal Text",
      "ingredients": [
        {
          "categories": null,
          "identifier": "frt",
          "path": "./A0-FRT.usfm",
          "sort": 0,
          "title": "Front Matter",
          "versification": "ufw",
          "alignment_count": 0
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "gen",
          "path": "./01-GEN.usfm",
          "sort": 1,
          "title": "Genesis",
          "versification": "ufw",
          "alignment_count": 23024
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "exo",
          "path": "./02-EXO.usfm",
          "sort": 2,
          "title": "Exodus",
          "versification": "ufw",
          "alignment_count": 18390
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "lev",
          "path": "./03-LEV.usfm",
          "sort": 3,
          "title": "Leviticus",
          "versification": "ufw",
          "alignment_count": 13293
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "num",
          "path": "./04-NUM.usfm",
          "sort": 4,
          "title": "Numbers",
          "versification": "ufw",
          "alignment_count": 17615
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "deu",
          "path": "./05-DEU.usfm",
          "sort": 5,
          "title": "Deuteronomy",
          "versification": "ufw",
          "alignment_count": 15520
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jos",
          "path": "./06-JOS.usfm",
          "sort": 6,
          "title": "Joshua",
          "versification": "ufw",
          "alignment_count": 10741
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jdg",
          "path": "./07-JDG.usfm",
          "sort": 7,
          "title": "Judges",
          "versification": "ufw",
          "alignment_count": 10614
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "rut",
          "path": "./08-RUT.usfm",
          "sort": 8,
          "title": "Ruth",
          "versification": "ufw",
          "alignment_count": 1424
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1sa",
          "path": "./09-1SA.usfm",
          "sort": 9,
          "title": "1 Samuel",
          "versification": "ufw",
          "alignment_count": 14629
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2sa",
          "path": "./10-2SA.usfm",
          "sort": 10,
          "title": "2 Samuel",
          "versification": "ufw",
          "alignment_count": 12416
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1ki",
          "path": "./11-1KI.usfm",
          "sort": 11,
          "title": "1 Kings",
          "versification": "ufw",
          "alignment_count": 14237
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2ki",
          "path": "./12-2KI.usfm",
          "sort": 12,
          "title": "2 Kings",
          "versification": "ufw",
          "alignment_count": 13308
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1ch",
          "path": "./13-1CH.usfm",
          "sort": 13,
          "title": "1 Chronicles",
          "versification": "ufw",
          "alignment_count": 11226
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2ch",
          "path": "./14-2CH.usfm",
          "sort": 14,
          "title": "2 Chronicles",
          "versification": "ufw",
          "alignment_count": 14282
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ezr",
          "path": "./15-EZR.usfm",
          "sort": 15,
          "title": "Ezra",
          "versification": "ufw",
          "alignment_count": 3989
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "neh",
          "path": "./16-NEH.usfm",
          "sort": 16,
          "title": "Nehemiah",
          "versification": "ufw",
          "alignment_count": 5745
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "est",
          "path": "./17-EST.usfm",
          "sort": 17,
          "title": "Esther",
          "versification": "ufw",
          "alignment_count": 3364
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "job",
          "path": "./18-JOB.usfm",
          "sort": 18,
          "title": "Job",
          "versification": "ufw",
          "alignment_count": 8649
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "psa",
          "path": "./19-PSA.usfm",
          "sort": 19,
          "title": "Psalms",
          "versification": "ufw",
          "alignment_count": 19943
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "pro",
          "path": "./20-PRO.usfm",
          "sort": 20,
          "title": "Proverbs",
          "versification": "ufw",
          "alignment_count": 7585
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ecc",
          "path": "./21-ECC.usfm",
          "sort": 21,
          "title": "Ecclesiastes",
          "versification": "ufw",
          "alignment_count": 3312
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "sng",
          "path": "./22-SNG.usfm",
          "sort": 22,
          "title": "Song of Solomon",
          "versification": "ufw",
          "alignment_count": 1374
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "isa",
          "path": "./23-ISA.usfm",
          "sort": 23,
          "title": "Isaiah",
          "versification": "ufw",
          "alignment_count": 5893
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jer",
          "path": "./24-JER.usfm",
          "sort": 24,
          "title": "Jeremiah",
          "versification": "ufw",
          "alignment_count": 8336
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "lam",
          "path": "./25-LAM.usfm",
          "sort": 25,
          "title": "Lamentations",
          "versification": "ufw",
          "alignment_count": 1622
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ezk",
          "path": "./26-EZK.usfm",
          "sort": 26,
          "title": "Ezekiel",
          "versification": "ufw",
          "alignment_count": 19166
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "dan",
          "path": "./27-DAN.usfm",
          "sort": 27,
          "title": "Daniel",
          "versification": "ufw",
          "alignment_count": 6445
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hos",
          "path": "./28-HOS.usfm",
          "sort": 28,
          "title": "Hosea",
          "versification": "ufw",
          "alignment_count": 2563
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jol",
          "path": "./29-JOL.usfm",
          "sort": 29,
          "title": "Joel",
          "versification": "ufw",
          "alignment_count": 1024
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "amo",
          "path": "./30-AMO.usfm",
          "sort": 30,
          "title": "Amos",
          "versification": "ufw",
          "alignment_count": 2247
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "oba",
          "path": "./31-OBA.usfm",
          "sort": 31,
          "title": "Obadiah",
          "versification": "ufw",
          "alignment_count": 325
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jon",
          "path": "./32-JON.usfm",
          "sort": 32,
          "title": "Jonah",
          "versification": "ufw",
          "alignment_count": 763
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "mic",
          "path": "./33-MIC.usfm",
          "sort": 33,
          "title": "Micah",
          "versification": "ufw",
          "alignment_count": 709
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "nam",
          "path": "./34-NAM.usfm",
          "sort": 34,
          "title": "Nahum",
          "versification": "ufw",
          "alignment_count": 592
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hab",
          "path": "./35-HAB.usfm",
          "sort": 35,
          "title": "Habakkuk",
          "versification": "ufw",
          "alignment_count": 737
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "zep",
          "path": "./36-ZEP.usfm",
          "sort": 36,
          "title": "Zephaniah",
          "versification": "ufw",
          "alignment_count": 826
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hag",
          "path": "./37-HAG.usfm",
          "sort": 37,
          "title": "Haggai",
          "versification": "ufw",
          "alignment_count": 644
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "zec",
          "path": "./38-ZEC.usfm",
          "sort": 38,
          "title": "Zechariah",
          "versification": "ufw",
          "alignment_count": 3227
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "mal",
          "path": "./39-MAL.usfm",
          "sort": 39,
          "title": "Malachi",
          "versification": "ufw",
          "alignment_count": 948
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "mat",
          "path": "./41-MAT.usfm",
          "sort": 41,
          "title": "Matthew",
          "versification": "ufw",
          "alignment_count": 19233
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "mrk",
          "path": "./42-MRK.usfm",
          "sort": 42,
          "title": "Mark",
          "versification": "ufw",
          "alignment_count": 11830
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "luk",
          "path": "./43-LUK.usfm",
          "sort": 43,
          "title": "Luke",
          "versification": "ufw",
          "alignment_count": 20345
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jhn",
          "path": "./44-JHN.usfm",
          "sort": 44,
          "title": "John",
          "versification": "ufw",
          "alignment_count": 16236
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "act",
          "path": "./45-ACT.usfm",
          "sort": 45,
          "title": "Acts",
          "versification": "ufw",
          "alignment_count": 19271
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "rom",
          "path": "./46-ROM.usfm",
          "sort": 46,
          "title": "Romans",
          "versification": "ufw",
          "alignment_count": 7609
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1co",
          "path": "./47-1CO.usfm",
          "sort": 47,
          "title": "1 Corinthians",
          "versification": "ufw",
          "alignment_count": 7262
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2co",
          "path": "./48-2CO.usfm",
          "sort": 48,
          "title": "2 Corinthians",
          "versification": "ufw",
          "alignment_count": 4845
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "gal",
          "path": "./49-GAL.usfm",
          "sort": 49,
          "title": "Galatians",
          "versification": "ufw",
          "alignment_count": 2334
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "eph",
          "path": "./50-EPH.usfm",
          "sort": 50,
          "title": "Ephesians",
          "versification": "ufw",
          "alignment_count": 2565
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "php",
          "path": "./51-PHP.usfm",
          "sort": 51,
          "title": "Philippians",
          "versification": "ufw",
          "alignment_count": 1743
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "col",
          "path": "./52-COL.usfm",
          "sort": 52,
          "title": "Colossians",
          "versification": "ufw",
          "alignment_count": 1671
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1th",
          "path": "./53-1TH.usfm",
          "sort": 53,
          "title": "1 Thessalonians",
          "versification": "ufw",
          "alignment_count": 1549
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2th",
          "path": "./54-2TH.usfm",
          "sort": 54,
          "title": "2 Thessalonians",
          "versification": "ufw",
          "alignment_count": 875
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1ti",
          "path": "./55-1TI.usfm",
          "sort": 55,
          "title": "1 Timothy",
          "versification": "ufw",
          "alignment_count": 1699
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2ti",
          "path": "./56-2TI.usfm",
          "sort": 56,
          "title": "2 Timothy",
          "versification": "ufw",
          "alignment_count": 1310
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "tit",
          "path": "./57-TIT.usfm",
          "sort": 57,
          "title": "Titus",
          "versification": "ufw",
          "alignment_count": 716
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "phm",
          "path": "./58-PHM.usfm",
          "sort": 58,
          "title": "Philemon",
          "versification": "ufw",
          "alignment_count": 361
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "heb",
          "path": "./59-HEB.usfm",
          "sort": 59,
          "title": "Hebrews",
          "versification": "ufw",
          "alignment_count": 5316
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jas",
          "path": "./60-JAS.usfm",
          "sort": 60,
          "title": "James",
          "versification": "ufw",
          "alignment_count": 1866
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1pe",
          "path": "./61-1PE.usfm",
          "sort": 61,
          "title": "1 Peter",
          "versification": "ufw",
          "alignment_count": 1825
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2pe",
          "path": "./62-2PE.usfm",
          "sort": 62,
          "title": "2 Peter",
          "versification": "ufw",
          "alignment_count": 1214
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1jn",
          "path": "./63-1JN.usfm",
          "sort": 63,
          "title": "1 John",
          "versification": "ufw",
          "alignment_count": 2207
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2jn",
          "path": "./64-2JN.usfm",
          "sort": 64,
          "title": "2 John",
          "versification": "ufw",
          "alignment_count": 261
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "3jn",
          "path": "./65-3JN.usfm",
          "sort": 65,
          "title": "3 John",
          "versification": "ufw",
          "alignment_count": 243
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jud",
          "path": "./66-JUD.usfm",
          "sort": 66,
          "title": "Jude",
          "versification": "ufw",
          "alignment_count": 487
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "rev",
          "path": "./67-REV.usfm",
          "sort": 67,
          "title": "Revelation",
          "versification": "ufw",
          "alignment_count": 10349
        }
      ],
      "checking_level": 3,
      "catalog": {
        "prod": {
          "branch_or_tag_name": "v45",
          "release_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/releases/3441599",
          "released": "2023-02-11T15:48:55Z",
          "zipball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/v45.zip",
          "tarball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/v45.tar.gz",
          "git_trees_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/git/trees/v45?recursive=1&per_page=99999",
          "contents_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/contents?ref=v45"
        },
        "preprod": null,
        "draft": null,
        "latest": {
          "branch_or_tag_name": "master",
          "release_url": null,
          "released": "2023-03-17T13:38:15Z",
          "zipball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/master.zip",
          "tarball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/master.tar.gz",
          "git_trees_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/git/trees/master?recursive=1&per_page=99999",
          "contents_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/contents?ref=master"
        }
      },
      "content_format": "usfm"
    }
  },
  "head": {
    "label": "gt-RUT-cecil.new",
    "ref": "gt-RUT-cecil.new",
    "sha": "0006fcc95e0d3f36090a6fefe6ad0abed15a36b3",
    "repo_id": 11419,
    "repo": {
      "id": 11419,
      "owner": {
        "id": 613,
        "login": "unfoldingWord",
        "login_name": "",
        "full_name": "unfoldingWord®",
        "email": "unfoldingword@noreply.door43.org",
        "avatar_url": "https://qa.door43.org/avatars/1bc81b740b4286613cdaa55ddfe4b1fc",
        "language": "",
        "is_admin": false,
        "last_login": "0001-01-01T00:00:00Z",
        "created": "2016-02-16T23:44:26Z",
        "repo_languages": [
          "el-x-koine",
          "en",
          "fr",
          "hbo"
        ],
        "repo_subjects": [
          "Aligned Bible",
          "Aramaic Grammar",
          "Bible",
          "Greek Grammar",
          "Greek Lexicon",
          "Greek New Testament",
          "Hebrew Grammar",
          "Hebrew Old Testament",
          "OBS Study Questions",
          "OBS Translation Notes",
          "OBS Translation Questions",
          "Open Bible Stories",
          "Study Notes",
          "Training Library",
          "Translation Academy",
          "Translation Words",
          "TSV OBS Study Notes",
          "TSV OBS Study Questions",
          "TSV OBS Translation Notes",
          "TSV OBS Translation Questions",
          "TSV OBS Translation Words Links",
          "TSV Study Notes",
          "TSV Study Questions",
          "TSV Translation Notes",
          "TSV Translation Questions",
          "TSV Translation Words Links"
        ],
        "repo_metadata_types": [
          "rc"
        ],
        "restricted": false,
        "active": false,
        "prohibit_login": false,
        "location": "",
        "website": "https://unfoldingword.org",
        "description": "",
        "visibility": "public",
        "followers_count": 0,
        "following_count": 0,
        "starred_repos_count": 0,
        "username": "unfoldingWord"
      },
      "name": "en_ult",
      "full_name": "unfoldingWord/en_ult",
      "description": "unfoldingWord® Literal Text (formerly ULB)",
      "empty": false,
      "private": false,
      "fork": false,
      "template": false,
      "parent": null,
      "mirror": false,
      "size": 399752,
      "languages_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/languages",
      "html_url": "https://qa.door43.org/unfoldingWord/en_ult",
      "ssh_url": "git@qa.door43.org:unfoldingWord/en_ult.git",
      "clone_url": "https://qa.door43.org/unfoldingWord/en_ult.git",
      "original_url": "",
      "website": "https://www.unfoldingword.org/ult",
      "stars_count": 9,
      "forks_count": 12,
      "watchers_count": 3,
      "open_issues_count": 52,
      "open_pr_counter": 2,
      "release_counter": 46,
      "default_branch": "master",
      "archived": false,
      "created_at": "2017-06-01T22:16:16Z",
      "updated_at": "2023-03-17T13:39:35Z",
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
      "allow_rebase": false,
      "allow_rebase_explicit": false,
      "allow_squash_merge": true,
      "allow_rebase_update": true,
      "default_delete_branch_after_merge": false,
      "default_merge_style": "merge",
      "avatar_url": "https://qa.door43.org/repo-avatars/11419-ba04b7a1942e0a5ba20adc6ac9372799",
      "internal": false,
      "mirror_interval": "",
      "mirror_updated": "0001-01-01T00:00:00Z",
      "repo_transfer": null,
      "metadata_type": "rc",
      "metadata_version": "0.2",
      "language": "en",
      "language_title": "English",
      "language_direction": "ltr",
      "language_is_gl": true,
      "subject": "Aligned Bible",
      "title": "unfoldingWord® Literal Text",
      "ingredients": [
        {
          "categories": null,
          "identifier": "frt",
          "path": "./A0-FRT.usfm",
          "sort": 0,
          "title": "Front Matter",
          "versification": "ufw",
          "alignment_count": 0
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "gen",
          "path": "./01-GEN.usfm",
          "sort": 1,
          "title": "Genesis",
          "versification": "ufw",
          "alignment_count": 23024
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "exo",
          "path": "./02-EXO.usfm",
          "sort": 2,
          "title": "Exodus",
          "versification": "ufw",
          "alignment_count": 18390
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "lev",
          "path": "./03-LEV.usfm",
          "sort": 3,
          "title": "Leviticus",
          "versification": "ufw",
          "alignment_count": 13293
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "num",
          "path": "./04-NUM.usfm",
          "sort": 4,
          "title": "Numbers",
          "versification": "ufw",
          "alignment_count": 17615
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "deu",
          "path": "./05-DEU.usfm",
          "sort": 5,
          "title": "Deuteronomy",
          "versification": "ufw",
          "alignment_count": 15520
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jos",
          "path": "./06-JOS.usfm",
          "sort": 6,
          "title": "Joshua",
          "versification": "ufw",
          "alignment_count": 10741
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jdg",
          "path": "./07-JDG.usfm",
          "sort": 7,
          "title": "Judges",
          "versification": "ufw",
          "alignment_count": 10614
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "rut",
          "path": "./08-RUT.usfm",
          "sort": 8,
          "title": "Ruth",
          "versification": "ufw",
          "alignment_count": 1424
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1sa",
          "path": "./09-1SA.usfm",
          "sort": 9,
          "title": "1 Samuel",
          "versification": "ufw",
          "alignment_count": 14629
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2sa",
          "path": "./10-2SA.usfm",
          "sort": 10,
          "title": "2 Samuel",
          "versification": "ufw",
          "alignment_count": 12416
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1ki",
          "path": "./11-1KI.usfm",
          "sort": 11,
          "title": "1 Kings",
          "versification": "ufw",
          "alignment_count": 14237
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2ki",
          "path": "./12-2KI.usfm",
          "sort": 12,
          "title": "2 Kings",
          "versification": "ufw",
          "alignment_count": 13308
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "1ch",
          "path": "./13-1CH.usfm",
          "sort": 13,
          "title": "1 Chronicles",
          "versification": "ufw",
          "alignment_count": 11226
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "2ch",
          "path": "./14-2CH.usfm",
          "sort": 14,
          "title": "2 Chronicles",
          "versification": "ufw",
          "alignment_count": 14282
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ezr",
          "path": "./15-EZR.usfm",
          "sort": 15,
          "title": "Ezra",
          "versification": "ufw",
          "alignment_count": 3989
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "neh",
          "path": "./16-NEH.usfm",
          "sort": 16,
          "title": "Nehemiah",
          "versification": "ufw",
          "alignment_count": 5745
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "est",
          "path": "./17-EST.usfm",
          "sort": 17,
          "title": "Esther",
          "versification": "ufw",
          "alignment_count": 3364
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "job",
          "path": "./18-JOB.usfm",
          "sort": 18,
          "title": "Job",
          "versification": "ufw",
          "alignment_count": 8649
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "psa",
          "path": "./19-PSA.usfm",
          "sort": 19,
          "title": "Psalms",
          "versification": "ufw",
          "alignment_count": 19943
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "pro",
          "path": "./20-PRO.usfm",
          "sort": 20,
          "title": "Proverbs",
          "versification": "ufw",
          "alignment_count": 7585
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ecc",
          "path": "./21-ECC.usfm",
          "sort": 21,
          "title": "Ecclesiastes",
          "versification": "ufw",
          "alignment_count": 3312
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "sng",
          "path": "./22-SNG.usfm",
          "sort": 22,
          "title": "Song of Solomon",
          "versification": "ufw",
          "alignment_count": 1374
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "isa",
          "path": "./23-ISA.usfm",
          "sort": 23,
          "title": "Isaiah",
          "versification": "ufw",
          "alignment_count": 5893
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jer",
          "path": "./24-JER.usfm",
          "sort": 24,
          "title": "Jeremiah",
          "versification": "ufw",
          "alignment_count": 8336
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "lam",
          "path": "./25-LAM.usfm",
          "sort": 25,
          "title": "Lamentations",
          "versification": "ufw",
          "alignment_count": 1622
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "ezk",
          "path": "./26-EZK.usfm",
          "sort": 26,
          "title": "Ezekiel",
          "versification": "ufw",
          "alignment_count": 19166
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "dan",
          "path": "./27-DAN.usfm",
          "sort": 27,
          "title": "Daniel",
          "versification": "ufw",
          "alignment_count": 6445
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hos",
          "path": "./28-HOS.usfm",
          "sort": 28,
          "title": "Hosea",
          "versification": "ufw",
          "alignment_count": 2563
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jol",
          "path": "./29-JOL.usfm",
          "sort": 29,
          "title": "Joel",
          "versification": "ufw",
          "alignment_count": 1024
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "amo",
          "path": "./30-AMO.usfm",
          "sort": 30,
          "title": "Amos",
          "versification": "ufw",
          "alignment_count": 2247
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "oba",
          "path": "./31-OBA.usfm",
          "sort": 31,
          "title": "Obadiah",
          "versification": "ufw",
          "alignment_count": 325
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "jon",
          "path": "./32-JON.usfm",
          "sort": 32,
          "title": "Jonah",
          "versification": "ufw",
          "alignment_count": 763
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "mic",
          "path": "./33-MIC.usfm",
          "sort": 33,
          "title": "Micah",
          "versification": "ufw",
          "alignment_count": 709
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "nam",
          "path": "./34-NAM.usfm",
          "sort": 34,
          "title": "Nahum",
          "versification": "ufw",
          "alignment_count": 592
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hab",
          "path": "./35-HAB.usfm",
          "sort": 35,
          "title": "Habakkuk",
          "versification": "ufw",
          "alignment_count": 737
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "zep",
          "path": "./36-ZEP.usfm",
          "sort": 36,
          "title": "Zephaniah",
          "versification": "ufw",
          "alignment_count": 826
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "hag",
          "path": "./37-HAG.usfm",
          "sort": 37,
          "title": "Haggai",
          "versification": "ufw",
          "alignment_count": 644
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "zec",
          "path": "./38-ZEC.usfm",
          "sort": 38,
          "title": "Zechariah",
          "versification": "ufw",
          "alignment_count": 3227
        },
        {
          "categories": [
            "bible-ot"
          ],
          "identifier": "mal",
          "path": "./39-MAL.usfm",
          "sort": 39,
          "title": "Malachi",
          "versification": "ufw",
          "alignment_count": 948
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "mat",
          "path": "./41-MAT.usfm",
          "sort": 41,
          "title": "Matthew",
          "versification": "ufw",
          "alignment_count": 19233
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "mrk",
          "path": "./42-MRK.usfm",
          "sort": 42,
          "title": "Mark",
          "versification": "ufw",
          "alignment_count": 11830
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "luk",
          "path": "./43-LUK.usfm",
          "sort": 43,
          "title": "Luke",
          "versification": "ufw",
          "alignment_count": 20345
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jhn",
          "path": "./44-JHN.usfm",
          "sort": 44,
          "title": "John",
          "versification": "ufw",
          "alignment_count": 16236
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "act",
          "path": "./45-ACT.usfm",
          "sort": 45,
          "title": "Acts",
          "versification": "ufw",
          "alignment_count": 19271
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "rom",
          "path": "./46-ROM.usfm",
          "sort": 46,
          "title": "Romans",
          "versification": "ufw",
          "alignment_count": 7609
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1co",
          "path": "./47-1CO.usfm",
          "sort": 47,
          "title": "1 Corinthians",
          "versification": "ufw",
          "alignment_count": 7262
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2co",
          "path": "./48-2CO.usfm",
          "sort": 48,
          "title": "2 Corinthians",
          "versification": "ufw",
          "alignment_count": 4845
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "gal",
          "path": "./49-GAL.usfm",
          "sort": 49,
          "title": "Galatians",
          "versification": "ufw",
          "alignment_count": 2334
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "eph",
          "path": "./50-EPH.usfm",
          "sort": 50,
          "title": "Ephesians",
          "versification": "ufw",
          "alignment_count": 2565
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "php",
          "path": "./51-PHP.usfm",
          "sort": 51,
          "title": "Philippians",
          "versification": "ufw",
          "alignment_count": 1743
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "col",
          "path": "./52-COL.usfm",
          "sort": 52,
          "title": "Colossians",
          "versification": "ufw",
          "alignment_count": 1671
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1th",
          "path": "./53-1TH.usfm",
          "sort": 53,
          "title": "1 Thessalonians",
          "versification": "ufw",
          "alignment_count": 1549
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2th",
          "path": "./54-2TH.usfm",
          "sort": 54,
          "title": "2 Thessalonians",
          "versification": "ufw",
          "alignment_count": 875
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1ti",
          "path": "./55-1TI.usfm",
          "sort": 55,
          "title": "1 Timothy",
          "versification": "ufw",
          "alignment_count": 1699
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2ti",
          "path": "./56-2TI.usfm",
          "sort": 56,
          "title": "2 Timothy",
          "versification": "ufw",
          "alignment_count": 1310
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "tit",
          "path": "./57-TIT.usfm",
          "sort": 57,
          "title": "Titus",
          "versification": "ufw",
          "alignment_count": 716
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "phm",
          "path": "./58-PHM.usfm",
          "sort": 58,
          "title": "Philemon",
          "versification": "ufw",
          "alignment_count": 361
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "heb",
          "path": "./59-HEB.usfm",
          "sort": 59,
          "title": "Hebrews",
          "versification": "ufw",
          "alignment_count": 5316
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jas",
          "path": "./60-JAS.usfm",
          "sort": 60,
          "title": "James",
          "versification": "ufw",
          "alignment_count": 1866
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1pe",
          "path": "./61-1PE.usfm",
          "sort": 61,
          "title": "1 Peter",
          "versification": "ufw",
          "alignment_count": 1825
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2pe",
          "path": "./62-2PE.usfm",
          "sort": 62,
          "title": "2 Peter",
          "versification": "ufw",
          "alignment_count": 1214
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "1jn",
          "path": "./63-1JN.usfm",
          "sort": 63,
          "title": "1 John",
          "versification": "ufw",
          "alignment_count": 2207
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "2jn",
          "path": "./64-2JN.usfm",
          "sort": 64,
          "title": "2 John",
          "versification": "ufw",
          "alignment_count": 261
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "3jn",
          "path": "./65-3JN.usfm",
          "sort": 65,
          "title": "3 John",
          "versification": "ufw",
          "alignment_count": 243
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "jud",
          "path": "./66-JUD.usfm",
          "sort": 66,
          "title": "Jude",
          "versification": "ufw",
          "alignment_count": 487
        },
        {
          "categories": [
            "bible-nt"
          ],
          "identifier": "rev",
          "path": "./67-REV.usfm",
          "sort": 67,
          "title": "Revelation",
          "versification": "ufw",
          "alignment_count": 10349
        }
      ],
      "checking_level": 3,
      "catalog": {
        "prod": {
          "branch_or_tag_name": "v45",
          "release_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/releases/3441599",
          "released": "2023-02-11T15:48:55Z",
          "zipball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/v45.zip",
          "tarball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/v45.tar.gz",
          "git_trees_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/git/trees/v45?recursive=1&per_page=99999",
          "contents_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/contents?ref=v45"
        },
        "preprod": null,
        "draft": null,
        "latest": {
          "branch_or_tag_name": "master",
          "release_url": null,
          "released": "2023-03-17T13:38:15Z",
          "zipball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/master.zip",
          "tarball_url": "https://qa.door43.org/unfoldingWord/en_ult/archive/master.tar.gz",
          "git_trees_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/git/trees/master?recursive=1&per_page=99999",
          "contents_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ult/contents?ref=master"
        }
      },
      "content_format": "usfm"
    }
  },
  "merge_base": "77747d060f9bfe97ab0b220683603a1a9188668a",
  "due_date": "2023-03-17T12:50:55Z",
  "created_at": "2023-03-17T12:55:30Z",
  "updated_at": "2023-03-17T13:39:34Z",
  "closed_at": null
}
```