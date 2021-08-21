# SPIKE: Add TWL resource using TSV 
Issue requirements:
```
Using the TWL TSV, only present the TWs for the current book in a TSV view, showing the verse text like in the TN TSVs.

From stakeholder feedback: Show the translated text of the TW article in subsequent occurrences in the book. In other words, if "Truth" occurs in 1:5, 2:6, and 3:7, the user translates the article in 1:5 and when the user gets to 2:6 the newly translated article is displayed.
```

Ideas and approaches below...

## A. Treat similar to TN

In this approach, the row bar has the TWL content and the source and target TW articles are loaded into two "dummy" columns.

Pros: familiar
Cons:
- there are two files being managed and so you can't makes edits to both
- would consume a lot of memory and would have excessive load times
- would still need a way to edit the TWL file itself

## B. Use TWL as a secondary file picker

In this approach you would invent a fake resource called "book package TW". This selection would use the TWL for a book to present a list of all the articles. When one was selected, it would use the current TW markdown editor to edit.

The TWL files themselves would be edited using current TSV handling.

## C. Use Perma-Links

When editing TWL files, allow user to select an article to edit. That would open a new tab with the source/target using current TW markdown editor.

## D. Use a context sensitive card interface

This borrows ideas from Create 2.0. Revamp the workspace to house separate but connected cards. This would need the following:
- a scripture reference picker (could use the one from Create 2.0)
- for the given reference, show the rows from the TWL file
- when user select a TWL row, show a translate view (source/target) of the article

