# Catalog Next
This document was begun as the deliverable to a spike issue:
https://github.com/unfoldingWord/tc-create-app/issues/829

## To Do
- capture future features like user choice of version and that version being sticky
- correct below to clearly specify that in near term, it is latest prod that will be used

## Links

The Catalog Next (CN) overview:
https://forum.door43.org/t/catalog-next-populating-accessing-the-dcs-resource-catalog/899/8

The Catalog Next Swagger:
https://qa.door43.org/api/catalog/swagger#/v5/v5SearchOwner

Clarifying discussion in Zulip at:
- [here](https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE--.20UR/topic/tc-create-app.23818/near/238464863)
- [here](https://unfoldingword.zulipchat.com/#narrow/stream/207526-Tools--.20UR/topic/Stakeholder.20Meeting/near/238465718)

## Key Differences with GRT

There are some key differences in how CN works compared to the gitea-react-toolkit (GRT)

1. There is no requirement for authentication: this simplifies things a little
2. A release is read only: this simplifies things a lot
3. A "source file context" with hooks based on the File Context provided by GRT is one approach to populating the source side with catalog next content.
4. However, the "default content" for the user branch will be different depending on whether the user is a translator or a uW English language resource maintainer ("editor"). Please see [[Edit vs Translation]] for details. This same information is in Zulip [here](https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE--.20UR/topic/tc-create-app.23829/near/238499028). and in issue 829.

## Prior Work

1. The caching proposal below is similar to the approach used by the Content Validation App, which was developed by Bruce McLean. You can watch the console when the app starts to see some of the concepts described below in action. Specifically:
- ![[Pasted image 20210512134301.png]]
- All the above repos are fetched concurrently at startup while the first page of the app is being show to the user. The timing, shown later in the console log, shows that all of them took 1.348s:
- ![[Pasted image 20210512134447.png]]
2. The `gatewayEdit` code base includes a hook that wraps the NPM "Local Forage" component and manages data in the browser's `indexedDB` database.

## Proposal - High Level

Create a headless component library that integrates CN with DCS.

The features of this CN component are:

1. For production users, restrict application access to production releases for our resources.
2. For developers, allow specification of the CN release to use for resources
3. For improved performance (and reduced network and server resources), provide automatic management of caching based on resource releases. A description of the caching features is below.

## Proposal - Details

### Initialization

The CN component requires initialization when the application starts. This initialization will proceed in the background, so that app is not required to wait until it is done. 

The initialization parameters are:
- SERVER: this points to the copy of DCS to use. Default will be 
`https://qa.door43.org`; and for production: `https://git.door43.org`.
- OWNER: this is the owner of the resources. The default is 'unfoldingWord'. 
- LANGID: this provides the language to be used in identifying resources. Default will be `en`.
- RESOURCES: this is a comma delimited list of resource IDs. The complete list is:
'ta', 'tw', 'twl', 'tn', 'tq', 'obs', 'obs-tq', 'obs-tn', 'obs-sn', 'obs-sq'. *Issue: how to specify Original Language resources (Greek and Hebrew texts)*.
- STAGE: this is the value for CN component which determines the release level to work with. It must be one of the following: 'latest', 'preprod', 'prod', 'draft'. The default is 'latest'. *For the source stage 'prod' (i.e., latest production) will be used.*
- APIVERSION: this may be hard coded for now to be `api/catalog/v5`.

### API

The CN component will provide a set of functions to access the content of a resource.

Given a server, owner, language, resource type list, and stage:
- initialize(): 
	- will clear the "unzipped catalog"; this storage contains items already extracted from the zip. 
	- will run `initalizeLocalCatalog()` for each resource type asynchronously. 

- initializeLocalCatalog(): will retrieve the corresponding Zip file for the resource.
- If stage is not provided, then stage 'prod' will be used. *This is the common case.*

Given an owner, language, and resource type:
- getManifest(): return the manifest as a JSON object. *This would be used to access arbitrary attributes in the manifest.*
- getProjects(): return the list of projects from the manifest. *This would be used to provide a list of projects from which the user may select.*
- getFilenames(): return a list of the files that correspond to the project(s) in the sort order (if any) provided by the manifest. *This would be used to provide a file list from which the user may select.* NOTE: for non-Scripture resources (i.e., TW and TA), a list of filenames will have to be extracted from the zip itself. This list should itself be stored in the unzipped catalog.

Given an owner, language, resource type, and a filename:
- getContent(): returns the content of the file. This function may throw an error if the content does not pass validation.

## Local Catalog Features

There are two levels to the local catalog:
- one that is cleared at the start of the app: "unzipped catalog"; this stores extracted items from the zip file so that it is only done once during a session.
- one that contains the released resource CN zip files: "zip catalog"

When the app requests a file from the catalog (either the manifest or a project file), then:

- The unzipped catalog is checked first and the file is returned if found
- If not found, then the file is extracted from the zip file and stored in the unzipped catalog; then it is returned.

There may be times when a user logs in and gets to the file selection step before the initialization has fetched the zip file. In this case, the app must show an "in process" spinner until the zip is present for use.

*Issue*: Running apps that use the CN component in multiple tabs of the same browser may interfere with each other at the local storage level. For example, consider when a new latest is available to the app in the second tab that wasn't available when the app started in the first tab.

# Examples

## Example #1 

Purpose: Find the latest production release for a resource.

Given:
- SERVER: `https://qa.door43.org`
- APIVERSION: `/api/catalog/v5`
- OWNER: `unfoldingword`
- LANGID: `en`
- RESOURCEID: `twl`
- STAGE: `prod`
Then, this URL may be constructed, which returns the JSON (shortened) below:
https://qa.door43.org/api/catalog/v5/search/unfoldingword/en_twl?stage=prod

https://qa.door43.org/api/catalog/v5/search/unfoldingword/en_twl?stage=prod&includeHistory=true

with history ^^^

```json
{
  "ok": true,
  "data": [
    {
      "id": 5076,
      "url": "https://qa.door43.org/api/catalog/v5/entry/unfoldingWord/en_twl/v2",
      "name": "en_twl",
      "owner": "unfoldingWord",
      "full_name": "unfoldingWord/en_twl",
      "repo": {
        "id": 60478,
        "owner": {
          "id": 613,
...
          ],
          "repo_subjects": [
            "Aligned Bible",
            "Aramaic Grammar",
            "Bible",
            "Greek Grammar",
            "Greek New Testament",
            "Hebrew Grammar",
            "Hebrew Old Testament",
...
          ],
          "restricted": false,
          "username": "unfoldingWord"
        },
        "name": "en_twl",
        "full_name": "unfoldingWord/en_twl",
        "description": "Links from the original language words to Translation Words articles.\r\n\r\nPreviously customized links in UHB and UGNT were used (but that didn't enable them to be customized for Gateway Languages).",
...        "language": "en",
        "subject": "TSV Translation Words Links",
        "books": [
          "gen",
          "exo",
...
        ],
        "title": "unfoldingWord® Translation Words Links",
        "checking_level": "3",
        "catalog": {
          "prod": {
            "branch_or_tag_name": "v2",
            "release_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_twl/releases/10762",
            "released": "2021-04-27T00:10:33Z",
            "zipball_url": "https://qa.door43.org/unfoldingWord/en_twl/archive/v2.zip",
            "tarball_url": "https://qa.door43.org/unfoldingWord/en_twl/archive/v2.tar.gz"
          },
          "preprod": null,
          "draft": null,
          "latest": {
            "branch_or_tag_name": "master",
            "release_url": null,
            "released": "2021-05-10T23:16:13Z",
            "zipball_url": "https://qa.door43.org/unfoldingWord/en_twl/archive/master.zip",
            "tarball_url": "https://qa.door43.org/unfoldingWord/en_twl/archive/master.tar.gz"
          }
        },
        "mirror_interval": ""
      },
      "release": {
        "id": 10762,
        "tag_name": "v2",
...    }
  ]
}
```
In this JSON, the "repo.catalog" property contains a single "prod" entry (note: history is turned off by default, so that you only get a single entry for stage=prod). This "prod" entry has the `zipball_url` which can be used to download the production release zip file.

## Example #2

Purpose: Download the latest production release for a resource.

Given the results in example #1, the catalog entry (shortened) will be:
```json 
        "catalog": {
          "prod": {
            "branch_or_tag_name": "v2",
            "release_url": "https://qa.door43.org/api/v1/repos/unfoldingWord/...",
            "released": "2021-04-27T00:10:33Z",
            "zipball_url": "https://qa.door43.org/unfoldingWord/en_twl/archive/v2.zip",
            "tarball_url": "https://qa.door43.org/unfoldingWord/en_twl/archive/v2.tar.gz"
          },
          "preprod": null,
          "draft": null,
          "latest": {
            "branch_or_tag_name": "master",
            "release_url": null,
            "released": "2021-05-10T23:16:13Z",
            "zipball_url": "https://qa.door43.org/unfoldingWord/en_twl/archive/master.zip",
            "tarball_url": "https://qa.door43.org/unfoldingWord/en_twl/archive/master.tar.gz"
          }
        },
```

Then, this property in the catalog entry  `repo.catalog.prod.zipball_url` contains this URL
https://qa.door43.org/unfoldingWord/en_twl/archive/v2.zip
which can be used to download the zip file.

