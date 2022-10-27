

## 2022-10-27

Working on https://github.com/unfoldingWord/tc-create-app/issues/1413
#issue1413 

Original post:
Edits are not saved in the TSV file after Save button is clicked unless the file is refreshed.
-Open en_tn> open file(book.tsv)

Add some text in the occurrence notes and click Save.
Click on the Chip on the Target side to go to the user branch. Screen Shot 2022-10-25 at 9.25.44 AM.png
Note that the changes are not saved on to the user branch.
Close the File and reopen or refresh the page.
Now check the user branch. Note that the changes are shown.
This is not the expected behaviour of Save button. User expects the edits to be reflected/saved in the user branch when the Save button is clicked.

### Steps to Replicate Issue

#### Production

First login and pick a TN, Jude. Use production: https://create.translationcore.com/

Second, make a change to the intro: put "xyz" after the word Jude.

Third click out.

Now examine local storage for auto-save activity.

Now auto-save activity. *Shouldn't there be?* Picture:
![[Pasted image 20221027113103.png]]

Fourth: click save.

Now examine:
1. Confirmed updated to my user branch.
2. Auto-save cache had indicator that changes were available, but when I click refresh, there was nothing. *Perhaps auto-save did work, but only when I clicked save. But then a successful save should have cleared it, which was in fact the case.*
3. Here is the network activity:
   ![[Pasted image 20221027113900.png]]

The first was "OPTIONS" (never heard of that one).
![[Pasted image 20221027114129.png]]

Second was the PUT to update the file; but had a 404... why?
![[Pasted image 20221027114157.png]]

Third was another PUT, which worked:
![[Pasted image 20221027114252.png]]

Fourth, another OPTIONS
![[Pasted image 20221027114359.png]]

Fifth, a GET to retrieve the latest copy:
![[Pasted image 20221027114516.png]]

Now, let's try again from the top to see if anything changes.

Change "xyz" to "abc" and click out.
This time there is autosave activity!
![[Pasted image 20221027114803.png]]

I closed *without saving* using the "x" near profile picture.

Then I went back to the file. 
- Data shown is from autosave. *Good!*
- But the save button is not enabled. *Bad!*

So I add a blank and then remove after the "abc". Now save button is enabled and I click it.

- Autosave cache is cleared. *Good.*
- DCS is updated. *Good.*

#### Development

First login and pick a TN, Jude. Use development(QA): 
https://develop--tc-create-app.netlify.app/

Second, make a change to the intro: put "xyz" after the word Jude.

Third click out.

Now examine local storage for auto-save activity.

No auto-save activity. *Shouldn't there be?* 

Fourth: click save.

Now examine:
1. Clicking the link on the page to the file *did not show the change!*, but...
2. Confirmed update to my user branch if I went to DCS.
4. Auto-save cache had indicator that changes were available, but when I click refresh, there was nothing. *Perhaps auto-save did work, but only when I clicked save. But then a successful save should have cleared it, which was in fact the case.*
5. The network activity was same as above.

Now, let's try again from the top to see if anything changes.

Change "xyz" to "abc" and click out. Again, this time there is autosave activity

I closed *without saving* using the "x" near profile picture.

Then I went back to the file. 
- Data shown is from autosave. *Good!*
- But the save button is not enabled. *Bad!*

So I add a blank and then remove after the "abc". Now save button is enabled and I click it.

- Autosave cache is cleared. *Good.*
- DCS is updated. *Good.*
