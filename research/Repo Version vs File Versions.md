# Versions: repo vs files

## Fundamentals

Files are stored in Github by their SHA value. Two files with the same SHA value are identical. So files are only stored once. This means that across releases of a repo, each release points to the same files if the files are unchanged.

This gives us a way to:
- compute the number of changes made to a file and thus compute "file versions"
- to identify the version of a file that exists in each repo release version

## Assumptions

1. Translators require stability in the source from which they are translating.
2. The master branch must be considered work-in-progress, for the most part (see step 4).
3. Branches and the Pull Request process are used to feed the master branch.
4. At some point, the content of master will be considered "reviewed and approved".
5. This is when a release is cut.
6. Files in a release are considered to  use by translators.

## Approach

This explores a specific repository where all the source files are in the top level directory of the repo: `unfoldingword/en_tn`.
This git API call returns for a given release tag (v46 in this case), the files and their SHA value:
```
https://qa.door43.org/api/v1/repos/unfoldingword/en_tn/git/trees/v46
```

A snippet of the JSON that is returned:
```json
    {
      "path": "en_tn_01-GEN.tsv",
      "mode": "100644",
      "type": "blob",
      "size": 719284,
      "sha": "b28f3c1fda66c648a15f3bfe337d83c8322a87f5",
      "url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_tn/git/blobs/b28f3c1fda66c648a15f3bfe337d83c8322a87f5"
    },
    {
      "path": "en_tn_02-EXO.tsv",
      "mode": "100644",
      "type": "blob",
      "size": 615668,
      "sha": "639b6c9c46fbf4848592bb2846ec2ba422b288eb",
      "url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_tn/git/blobs/639b6c9c46fbf4848592bb2846ec2ba422b288eb"
    },

```

A spreadsheet can be output from this. After sorting on File (1) and Release (2), then you get (an excerpt):
```csv
Repo,Release,File,SHA
en_twl,v1,twl_1CH.tsv,8dfa031ee35acf863f014b7ccf29170f970d6eb4
en_twl,v2,twl_1CH.tsv,8dfa031ee35acf863f014b7ccf29170f970d6eb4
en_twl,v1,twl_1CO.tsv,933b5f5c2227b6b496a7d0da7f62d2c29ff1d28b
en_twl,v2,twl_1CO.tsv,35e3c41f4e5700f4f51aa2e251eb7e97f4e4cd08
en_twl,v1,twl_1JN.tsv,a76f6347f2065b76d940b51cc273f64a155be513
en_twl,v2,twl_1JN.tsv,a76f6347f2065b76d940b51cc273f64a155be513
en_twl,v1,twl_1KI.tsv,76e014e0fb7f06a525fa5d2e05cdeb203a50354f
en_twl,v2,twl_1KI.tsv,76e014e0fb7f06a525fa5d2e05cdeb203a50354f
en_twl,v1,twl_1PE.tsv,70f1ec4f54b9b57200f0a64893e20982c5da2b5e
en_twl,v2,twl_1PE.tsv,70f1ec4f54b9b57200f0a64893e20982c5da2b5e
en_twl,v1,twl_1SA.tsv,d57e644fecfb71ed89572ff3b158100381a4cec8
en_twl,v2,twl_1SA.tsv,d57e644fecfb71ed89572ff3b158100381a4cec8
```

Notice that the TWL for 1CH did not change between v1 and v2 (they both have the same SHA, so its the same file). So across both releases of this TWL repo, 1CH is at version 1.
On the other hand, 1CO *did* change. So version 1 of 1CO is repo release v1 and version 2 of 1CO is in repo release v2.

Thus, this data could be processed to add a "revision" column:
```csv
Repo,Release,File,SHA,Revision
en_twl,v1,twl_1CH.tsv,8dfa031ee35acf863f014b7ccf29170f970d6eb4,1
en_twl,v1,twl_1CO.tsv,933b5f5c2227b6b496a7d0da7f62d2c29ff1d28b,1
en_twl,v2,twl_1CO.tsv,35e3c41f4e5700f4f51aa2e251eb7e97f4e4cd08,2
en_twl,v1,twl_1JN.tsv,a76f6347f2065b76d940b51cc273f64a155be513,1
en_twl,v1,twl_1KI.tsv,76e014e0fb7f06a525fa5d2e05cdeb203a50354f,1
en_twl,v1,twl_1PE.tsv,70f1ec4f54b9b57200f0a64893e20982c5da2b5e,1
en_twl,v1,twl_1SA.tsv,d57e644fecfb71ed89572ff3b158100381a4cec8,1
```


- latest release as source
- fixed version as source per book
- let user see other versions
- let users compare two source versions

Big Things:
1. sorting semver strings
2. performance concerns
3. joining to the manifest project ids in order to track a file even tho its name changed

Version 1 TN files were markdown
Version 13 introduced the current TSV naming convention: `en_tn_01-GEN.tsv`
A future version will introduce the naming convention: `tn_GEN.tsv`

