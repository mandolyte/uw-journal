# Issue 566

1. delete my branch at qa.door43.org/translate_test/en_tn/branches 
	cecil.new-tc-create-1
2. login with tc-create-app (local host)
3. org=translate_test
4. resource=tranlation notes
5. language=English
6. file=book of titus

Results:
- first, my session is paused on debugger with an uncaught exception
in 'helpers.js', called from TranslatableTSV.js and useFile.js. Message is:
"TypeError: Cannot read property 'content' of null"
- if I click to let it go on, it seems to act OK.
- if now run validation, no errors are found.
- at this point, my branch is re-created and contains Titus.
- refresh page to return to login page... do not log in yet.
- now in DCS, go to my branch and select Titus in DCS
- click the edit file button
- select all the text and delete it
- paste in text from @koz. This file is here named `koz_titus.tsv`.
- commit directly into my branch
- now login with tc-create and return to this file with the new content
- run validation and I get all the errors that @koz introduced:
```
Priority	Chapter	Verse	Line	Row ID	Details	Char Pos	Excerpt	Message	Location
889	front	intro	2	m2jl	SupportReference		​	Unable to find SupportReference TA link	 translate/​/01.md
851	front	intro	2	m2jl	OrigQuote		undefined	Unable to load original language verse text	
550	front	intro	2	m2jl			undefined	Invalid zero occurrence field when we have an original quote	
194	front	intro	2	m2jl	OccurrenceNote	110	…br>1.␣␣Pau…	Unexpected double spaces	
916	1	1	4	rtc9	OrigQuote			Unable to find original language quote in verse text	
916	1	1	5	xyz8	OrigQuote		εὐσέβειαν​	Unable to find original language quote in verse text	
```
- Now try to correct them
- Cannot find first error at all -- where is it???
- Fixed all the rest of them and for the last one just re-pasted into Origquote field (does not appear to be a valid error).
- click save
- now run validation
- it reports all the same errors, even tho the data has changed.
- close the file and return back to the file
- run validation
- it still reports the same errors (where is it getting the data from!?)
- refresh page to return to login page
- return to the file and run validation again
- it still reports the same errors!?
- delete all databases in indexedDB
- refresh page to return to login page
- return to the file and run validation again
- it still reports the same errors!
- delete all databases, close tab of locally running tc-create, restart with yarn start
- it still reports the same errors
- delete databases, close browser and return to page
- it still reports the same errors

About the first non-existent TA article issue...
Here are the three RC links in the front intro:
```
[[rc://en/ta/man/translate/translate-names]]
[[rc://en/ta/man/translate/figs-exclusive]]
[[rc://en/ta/man/translate/figs-you]]
```
The first one is the only one with "translate" as part of the name.
Let's change it say figs-you and see what happens.

