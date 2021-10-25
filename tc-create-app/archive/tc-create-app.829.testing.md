# Issue 829 Testing

## Editor Testing

### Editor Case B

1. In QA DCS, went to unfoldingword/en_tn
2. Removed my branch
3. Selected the book of Jude
4. Clicked edit icon
5. Altered first note to begin with "# (MASTER) Intro to Jude"
6. Clicked "Propose change" at bottom of page
7. Clicked "New Pull Request"
8. Clicked "Create Pull Request"
9. Clicked "Squash and Merge"
10. Clicked "Squash and Merge" (different button this time)
11. Clicked "Delete Branch 'cecil.new-patch-1'"
12. Click "yes" to confirm delete
13. Clicked "Files" tab, then selected book of Jude again to confirm change now in master branch
14. Start tc-create
15. Login as cecil.new
16. Select unfolding word org (testing editor role)
17. Select translation notes
18. Select English
19. Select book of Jude
20. Confirmed results:

![[Pasted image 20210526111434.png]]

### Editor Case C

Continues from last step in Case B.
1. Make changes to target side and save them. Replaced "MASTER" with "MYBRANCH".
2. Click the save button
3. Click the "X" button to return to step 3
4. Select translation notes
5. Select book of Jude (will not be asked for language this time)
6. Confirm target contains my prior edits:

![[Pasted image 20210526111838.png]]

## Translator Testing

I will be acting as an es-419 translator.

### Translator Case A

1. Go to https://qa.door43.org/es-419_gl/es-419_tn and remove my branch
2. Go to Files tab and note a file that *does not exist yet* in the repo. Will use Book of Jude.
3. Start tc-create (or refresh page)
4. Login as cecil.new
5. Select "Es-419_gl" org
6. Select translation notes
7. Select Es-419 language
8. Select book of Jude
9. Confirm that source side comes from latest published (v47 in my case). The first note should begin with "Introduction to Jude", not "(MASTER)".

![[Pasted image 20210526112657.png]]

### Translator Case B

Return to QA DCS and note a file that already exists in the master branch. I will use Titus.

1. Login, select es-419_gl org, and then translator notes
2. Select Titus
3. Confirm that a) source side has published content and target side has translated content (ie, will be in Spanish).

![[Pasted image 20210526113810.png]]

### Translator Case C
Continues from last step in case B.
1. Added "(MYBRANCH)" to the beginning of the note 
2. Click save
3. Click "X" to return to step 3
4. Click translation notes
5. Select book of Titus (language prompt will be skipped and will remain es-419).
6. Confirm that a) source is published content as before; b) target content has my prior changes to this file.

![[Pasted image 20210526114023.png]]
