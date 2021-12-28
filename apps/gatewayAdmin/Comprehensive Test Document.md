# Comprehensive Test Document
This document is intended to describe a comprehensive test for gateway-admin.

## Test Pre-requisites

First, a new, and empty, organization must be created in QA DCS, say, `ga_test`.

## Test 1 - Add Books 
Next, login to QA DCS and go to Settings to use the organization you have created.

After clicking "save and continue", you should see an empty page.

Click the "Add Book" button and add two books, one OT and one NT (say Ruth and Titus). For each book card, click the resources button to show book/repo details.

Expected results: all repos for both books should show the "+" icon under the action column.

## Test 2 - Add a repo
From the Ruth card, click the add repo button to create the en_glt repo.

Expected result on the server:
- in the ga_test org, the repo en_glt will be created.
- the manifest will be created
- the manifest will have an "invalid" sticker since at least one project is required

Expected results on the client:
- the repo will now show as existing with a new status saying "Book not in manifest"
- the action will be a plus sign having the tooltip "add book"
- the above will be on all the cards, not just the one used to create the repo

## Test 3 - Add a book to the manifest
On the Ruth card, click the Add Book action.

Expected result on server:
- in `ga_test/en_glt`, the manifest will now have a project entry for Ruth.
- the manifest will have a "valid" sticker

Expected result on client:
- the status will change to "File not found". 
- The tooltip will say "Use tC Create to create file"; but no action will be done if clicked.
- the status on the other card(s) will remain as "Book not in manifest".

## Test 4 - Add more books to the manifest
On the Titus card, click the Add book action.

Add another book, say, Psalms; click to add it to the manifest.

Expected result on server:
- All three books will be in the manifest and will be sorted in bible order. 

Expected result on client:
- all three cards will show "File not found" and the action tooltip saying "Use tC Create to create file".