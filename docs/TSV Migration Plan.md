# TSV Migration Plan

The migration starts with the English uW content.

*Dependencies*:
- Availability of the Markdown to TSV converter for the resource type
- Version 1.3 of tc-create is being used (distinguishes between editor and translator roles)


Steps for uW:
1. Obtain approval of content team to migrate a resource type to TSV. This means that all branches with open edits are either merged into the master branch or will be discarded.
2. Delete all `tc-create-1` branches.
3. (Re)Publish the markdown content (make a new release with latest approved changes)
4. Convert all files in master branch to TSV
5. **Do not publish the TSV content until GLs are cleared** (see below)

Steps for GLs:
1. Clear all work in progress and delete the user branches after merging work into master
2. Convert all files in master branch to TSV

Step for uW (final step):
- publish the TSV content

Final State:
- GLs can now work in TSV format using latest released from uW as their source
- uW can now continue editing and publishing as needed


