# Repository Release and Publishing Book Packages

## Introduction

### Book Packages
A book package is composed of resource files associated with a book of the bible. These resource files are contained in resource specific Gitea DCS repositories. 

A "published" book package consists of the full set of approved (released) resources associated with a book of the bible. Sometimes this set of versioned resources is called a Book Package (BP) Snapshot.

### Resources
There are two types of resources:
- resources that are book specific
- resources that are not

Resources that are book specific have a filename that includes the bookId of the book. For example, the Translation Word List for the book of Matthew is `twl_MAT.tsv`.

Translation Words (TW) and Translation Academy (TA) are the resources that are shared across book packages, i.e., these are not book specific. These are also 1-to-many:
- The TW articles for a book are listed in the Translation Word List resource for that book.
- The TA articles for a book are listed in the Support Reference column in the Translation Notes resource.

Resources are released as they are reviewed and approved. A release is a snapshot of the working "master" branch. *Not every file in the master branch is reviewed and approved!* When the master branch is released it may contain a mix of approved and unapproved content.

## Problem Statement
The concept of a *published* book package cuts across the resource repositories. The requirement is to:
- Identify the approved content for a book package.
- Approved content will be referenced in a release for a resource repository.

So the problem to solve is how to enable a project administrator identify, save, and maintain the set of files for a book bookage and make it available for use.

## Proposed Solution
For a given project (i.e., a given organization and language), create a new type of repository in which to maintain published book package data. Below find:
- a brief example
- a detailed example
- validations required to ensure the integrity of the BP

### Brief Example
Suppose the standard repo name to contain published BP data is the language ID and the resource ID "pub". Then for for uW English, the published BP repo would be `unfoldingWord/en_pub`.

This repository would contain 66 files, one for each book of the bible. Each file would be, say, JSON format and for each resource would have a "partial" URL (or Resource Container URIs?) to the released resource (if available). *Note, in the below we assume JSON format. But it could be TSV as well.*

### Detailed Example
Suppose the uW team determined that they wanted to publish the book package for Titus. Further, let's suppose that of all the resources possible for a book package, they only wished to publish the following:
- The Translation Notes (TN)
- The Translation Word List (TWL)
- The Translation Word articles referenced in the TWL file
- The Literal Translation (ULT)

Missing from the above are:
- The Simplified Translation (UST)
- The Translation Academy articles referenced in the TN file
- The Translation Questions
- The Study Notes
- The Study Questions

Prior to publishing the book package for Titus, they will have had to release the associated repositories with content for Titus. Let's assume that the following releases contain the content for Titus:
- TN: https://git.door43.org/unfoldingWord/en_tn/releases/tag/v63
- TWL: https://git.door43.org/unfoldingWord/en_twl/releases/tag/v15
- TW: https://git.door43.org/unfoldingWord/en_tw/releases/tag/v33
- ULT: https://git.door43.org/unfoldingWord/en_ult/releases/tag/v38

*Note: these are real URLs; feel free to click and explore!*

The admin would have a UI/X that allows the above to be identified. Then they would click to publish. At that point, some validations could be performed to ensure the integrity of the BP (see next section).

If validated, then a JSON file for the above would be created in the "en_pub" repository for Titus. It might be named `pub_TIT.json`

It's content might be:
```json
{
	"TN": "unfoldingWord/en_tn/raw/tag/v63/en_tn_57-TIT.tsv",
	"TWL": "unfoldingWord/en_twl/src/tag/v15/twl_TIT.tsv",
	"TW": [ array of links derived from the above TWL file],
	"LT": "unfoldingWord/en_ult/src/tag/v38/57-TIT.usfm"
}
```


### Validations
To ensure the integrity of the published book package, the following validations may be considered:
- The release for a resource is a *production* release.
- The file exists for the book in the release.
- *Content Validation* could be performed on the file.
- For TW, ensure that all TW articles referenced in the TWL exist *in the same org and language*.
- For TA, ensure that all the support references referenced in the TN exist *in same org and language*.
- For LT, require evidence of alignment in the book.