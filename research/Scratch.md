# Scratch Space

Consider the following desires:
- offline support
- leverage Proskomma
- transition to create 2.0 
- new TSV formats
- well defined book packages (see [here](https://unfoldingword.zulipchat.com/#narrow/stream/214041-CONTENT/topic/New.20TSV.20formats/near/226974917))

Suppose we had a way to precisely document the translation work done on a book package. It would need something link Jesse had in the link above, namely:
```
Jonah Book Package version 1:

ULT v12
UST v13
UTA v20
UTW v9
UTN v35
```

**The Problem:** 
- Today our tools do not use versioned resources, using the master branch instead.
- The master branch changes with no way to alert translators that the basis of their work has changed

**Why do we work this way?** Today even the English resources are not considered complete; thus translators tend the favor the latest even with the drawbacks it has. But isn't this a temporary problem? At some point the English resources will be stable and normal release life cycle procedures could be instituted.

**The Dream** let's posit for the sake of discussion that the English resources will become stable about the same time that Create 2.0 is ready for release. This coincidence of events would allow us to design 2.0 to:
- be offline-first
- leverage Proskomma for searching, retrieving texts
- the TSV and Markdown resources could be retrieved as zip files and treated as a local filesystem (NOTE: Bruce McLean did this for CV app, so the code already exists that could be adapted)

Below are three scenarios:
1. A translator starting a book package who has internet access
2. A translator revising a book package who has internet access
3. An intermediary with internet who acts on behalf of a translator who does not have internet access (Book Package Import/Export).

**Imagine a Translator Workflow for Starting a Book Package**
This is a GL workflow.
1. Start Create 2.0 and log in (obviously starting online to begin a book package)
2. Prompt for book package definition
3. When definition complete, retrieve all the resources, storing them locally. 
	- Using Proskomma forms for all USFM type resources
	- Using normal Git release Zip files for TSV and Markdown
4. Can now go offline...
5. Display (by translator preferences) all the info needed to begin translation for the book package
6. Commit all changes (at verse or frame level) locally to a "commit log"
7. When back online, make an "upload" button available for the package
8. The upload button would process the commit log backwards and upload to the translator's branch(s) all the translated material.
9. A PR would be created if branch did not already exist. It would be set to draft mode to indicate it was a work-in-progress and not ready for review.
10. The PR would include in the comment the exact dependencies used for this book package, perhaps like a YAML manifest in the comment.
11. Go to offline step and repeat above until material is ready to submit to review
12. The translator submits PR for review by changing the PR status so it is no longer in draft mode.
13. This would alert the admin reviewer to take it from there to resolve conflicts and merge into the master branch. If changes are required, change the PR back to draft mode and allow the translator to continue to make changes.
14. At some point, the admin would decide that they are ready to publish and would cut a version of their translated work.
15. This version of the translated work would be used, in turn, by an OL translation team.

**Imagine a Translator Workflow for Revising a Book Package**
This workflow would work for both GL and OL translators
1. Start Create 2.0 and log in
2. Select an existing book package to revise
3. Update the resources that have changed
4. Retrieve old package dependencies
5. Retrieve new package dependencies
6. Compare and capture all changes to help translator consider what might need to be updated (or created, for example, in the case of new tW or tA articles).
7. Can now go offline...
8. At this point, the process is the same.

**Book Package Import and Export**
This workflow considers how to solve the problem of a translator with no internet access at all. In other words, a translator who cannot connect to define a book package as described above.
1. This scenario requires an intermediary with internet access and who is able to physically exchange data with the translator utilizing a USB or other storage device.
2. The intermediary does the initial steps to define the book package (either for creation or for a revision).
3. The intermediary then exports the book package data to local storage, which then can be copied to a portable storage device.
4. The portable storage device can then be delivered to the translator.
5. The translator imports the book package from the portable storage device.
6. At this point, they can begin translating.
7. At any point during the translation, they may export the book package work as a backup or as a way to transfer the package to someone else to continue.
8. Once completed, the package is exported and copied to portable storage.
9. The portable storage is delivered to an intermediary with internet access who can upload it.
10. From here the steps are the same as above.