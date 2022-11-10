# Goals and Assumptions
Goal: maximize re-usability of alignment component by minimizing its dependencies on its environment.

Assumptions:
- In the description below, I assume that the alignment component (AC) lives in its own page in gateway-translate which is a NextJs application.  Being in its own page makes it easier to reason about.
- Let's also assume that the way to route to the aligner page is by selecting it in the sidebar navigation.

# Workflow

After user logs in, opens a file, and (perhaps) edits it. Then they wish to align it.

1. In side navigation, click the alignment button
2. The app routes to the alignment page
3. The page shows a list of files that have been opened.
4. The user selects a book
5. The app retrieves the original language text for the selected book (if hasn't been retrieved already) and loads it into the PK cache (converting it to PERF along the way)
6. The app accesses the original language PERF for the selected book from the PK cache
7. The app accesses the edited PERF for the selected book from the PK cache
8. The app instantiates the alignment component with the two PERF files, plus a save function, the bookId, and the versification for the book.
9. Thus the component would look something like this:
   ```html
   <Aligner 
	   originalLanguage={olPerf} 
	   content={editedPerf} 
	   bookid={bookId}
	   onSave={onSave}
	   versification={vrs}
	/>
```
10. The AC would then show, defaulting to chapter 1 verse 1.
11. In addition to the alignment area, the follow UI elements are needed:
   - chapter and verse navigation
   - a Cancel button: routes to main workspace without saving
   - a Save button: save the updated PERF to the PK cache
   - a Close button (which saves and routes to the main workspace)
12. Possibly, the cancel and close may need to fire functions provided by the app. In that case, they would be passed in similar to the "onSave" above.
13. Possibly, instead of passing in PERF, it may be desirable to pass in Epitelete instances

