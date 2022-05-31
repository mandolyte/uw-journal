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

- A set of documents is imported into Proskomma (PK) by the application.
	- The application owns the Proskomma instance
	- This will be one or more books of the Bible in USFM format.
	- The instance will be passed to Epitelete 
- The editor will request a book, which will be returned in PERF JSON format to the editor.
	- The requesting method is `Epitelete.readPerf` with bookCode as the sole argument
	- *Alternatively,* the editor may have stored the PERF JSON and may wish to resume editing. Then the method to use is: `Epitelete.loadPerf`. The argument is the PERF JSON that the app had saved in the prior editing session.
- The editor will convert the PERF into HTML which then may be rendered for the user to view or make changes. 
	- The method to do this is: `readHtml` . This method is not in Epitelete itself, but on a subclass of Epitelete that adds this method.
- When there is a change at the PERF Sequence level, then the editor will call the `writeHtml()` method of the Epitelete subclass.
- The editor will call this function with three arguments:
	- the document book code
	- the sequence id
	- the HTML content
- This function will extract the sequence from the HTML and will run the super class method "writePerf()"
- This function actually updates the original PERF JSON with the changed content. It does not "write" it anywhere. It will do the following steps:
	- Validate the sequence 
	- If valid, then the sequence will be merged into the document and the document.
	- The document as a whole will be validated and, if valid, returned to the editor
	- If not valid, then an exception will be thrown
- At the end of the editing, a new method (not yet defined) will be used to supply the PERF to the PK instance to update the document in PK. For example, it might be named "savePerf".

## Standalone Mode
We will need to use Epitelete in standalone mode. Standalone mode means to use it without reference to a PK instance.

In this usage, use the `loadPerf` method instead of `readPerf`. From that point forward, usage is the same as above since `writePerf` does not reference the PK instance.

Of course, the app can do both:
- It will use `readPerf` when it needs a different book.
- It will use `loadPerf` when it needs to resume a prior editing session.

Another use of standalone mode is in a script. Here only `loadPerf` and `writePerf` are needed. If no PK instance is provided, the constructor will need to handle *null* for the PK instance. And if `readPerf` is used by mistake, an exception will be thrown.




