# Testing Autosave
The scenarios below provided by ZP. My comments are *italicized*.


Scenario 1: 
HAPPY PATH :)

using DEVELOP url:
- *Used https://deploy-preview-1116--tc-create-app.netlify.app/*
- *QA DCS https://qa.door43.org/unfoldingWord/en_tn/branches --- I have no branches at the start of test*

1. Deleted branch (to be sure of clean starting point) *done*
1. logged in with org unfoldingWord and resource tN and language English *done*
1. selected book of Jude *done*
1. made an edit, but did not save *done*
1. turned off wifi *On win10, I went into Airplane Mode*
1. close
	- clicked the "x", I got a popup warning that my changes are not saved *done*
	- ![](./images/Pasted%20image%2020211223082241.png)
	- clicked cancel on popup *done*
	- tried to save *done*
	- got popup saying that changes could not be saved *done*
	- ![](./images/Pasted%20image%2020211223082331.png)
	- The save button was deactivated at this point *done*
1. Turned wifi back on *done*
1. refreshed page and it returned to logon page *done*
1. logged as before and selected book of Jude *done*

VERIFY: auto-save content is displayed. *verified*

*Question: at this point, the save button is _not_ active, even tho I have unsaved content. Which is to say, on QA DCS, my branch exists with the old version of Jude, not my changed version. If I really want to save the content, I'll have to make another change. then the button will become active, enabling me to save. Is this OK?* 

---

Scenario 2: 
DCS USER BRANCH CHANGED (IN "ANOTHER SESSION" - simulated by deleting branch)

using DEVELOP url:

1. logged in with org unfoldingWord and resource tN and language English *I simply continued where I was at the end of scenario 1.*
1. selected book of Jude *done*
1. made an edit and saved it *I added a space and then backspaced to remove, leaving the file as it had been restored at end of scenario 1. Confirmed my change is in my branch.*
1. went to QA DCS and deleted the branch I just updated/created in app, clicked the "x" to return to select resource *Instead, I opened an incognito window, logged in, and made a change to Jude and saved it. Confirmed that DCS has the change. Next I closed by incognito session window. Then I returned to the other, original window, clicked the close file ("x") button, then selected en_tn, then ...*
1. selected Jude again *done*

VERIFY: the auto-save content is there *Verified... but the content shown could have come from QA DCS itself. Not sure it picked up (or needed to pick up) any autosaved content from my incognito window.*

---

Scenario 3: 
DCS USER BRANCH CHANGED (IN "ANOTHER SESSION" - simulated by editing directly on DCS)

using DEVELOP url: 
*I used a different computer to make the change.*

1. logged in with org unfoldingWord and resource tN and language English *done*
1. selected book of Jude *done*
1. made an edit and saved it *done*
2. close tC Create *done*
1. went to QA DCS and edited the same file -- in the USER'S BRANCH. *used a different computer to make the change*
1. open tC Create; open the same file:


VERIFY: the app prompts you to DISCARD or KEEP the autosaved file.
**It did not prompt to discard or keep. It simply went to the latest file in my branch which I had changed on the other computer**

1. Discard: The file is closed. VERIFY: open the file again and it should show your server content.
1. --OR-- Keep.  Then edit and save your file in tC Create.  Then open tC Create and open the same file.  VERIFY: the newly saved content is shown.


---

Scenario 4:
TWO FILES

using DEVELOP url:

1. logged in with org unfoldingWord and resource tN and language English *done*
1. selected book of 1 JN *done*
1. made an edit ... DO NOT SAVE 
1. select book of 2 JN *after choosing 2JN from nav drawer, I receive a prompt: ![](./images/Pasted%20image%2020211223085954.png) I clicked OK to proceed to 2JN.*
	VERIFY: 2 JN looks OK *verified... 2JN came up ok.*
1. make an edit .... SAVE 2 JN *done*
1. selected book of 1 JN *done*
	VERIFY: prior auto-save content for 1 JN is correct. *Verified: my change is there.*


.
