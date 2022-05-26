# Perf Write
Repo Issue: https://github.com/Proskomma/epitelete/issues/15

**Original Statement**:
When we 'close the circle' the editor should send back HTML which eventually turns into PERF which can then be written back to the cache.

The method should take

-   the document bookCode and
-   PERF fragment for a sequence

as its arguments, and should assign that JSON to the appropriate place in `documents` before returning the amended `documents` object. The bookCode and sequenceId must already exist since, at present, we don't expect the editor to create new sequences or indeed documents this way.

Before assignment the method should validate the PERF. I just realised that I need to rework the validator for this use case... but, when we have the validator, the workflow will be

-   assemble the modified document
-   validate that document
-   if it validates, update `this.documents`
-   otherwise throw an error

## 2022-05-24

Cloned the repo: `git clone git@github.com:Proskomma/epitelete.git`

Switch to testing branch: `git switch testing-showPerfDataFormat`

Installed the dependencies: `npm install`

Run the test: `npm run rawTest`





# Big Picture

- A set of documents is imported into PK (application dependent).
- This will be one or more books of the Bible in USFM format.
- The editor will request via an API call (not GraphQL?) a portion of text, which will be returned in PERF JSON format to the editor.
- The editor will convert the PERF into HTML which then may be rendered for the user to view or make changes. (read HTML method on an editor specific subclass of Epitelete)
- When there is a change at the PERF Sequence level, then the editor will call the `perfWrite()` function (vis writeHtml())
- The editor will call this function with three arguments:
	- the document book code
	- the sequence id
	- the PERF sequence
- This function (actually a method of the Epitelete class), will:
	- Validate the sequence 
	- If valid, then the sequence will be merged into the document and the document returned to the editor
	- If not valid, then the errors will be returned to the editor (consult Klappy)

To Do:
- how to handle deleting and adding sequences
- may need to validate at both the sequence and later at the document level

Out of Scope:
- write the PERF back to Proskomma




