# Edit vs Translation

The tc-create app has both a "source" and a "target" for content. The source side is on the left side of the page and the target on the right. The term "edit" and "editor" are used to describe those who maintain the English language content of the uW resources. 
- *For editors, both source and target are the exact same repo.* 
- *For translators, the source and target are never the same repo; and generally never the same owner. In other words, translator work is saved in a different org.*

Below are both current process and future catalog next based process.

## Current Process

### Translator Chronology

1. Translator selects a file to translate from the list of files from the source uW master repo branch.
2. Branch is created in target repo and the file is created in that branch populated with the English content from the  uW source master repo branch.
3. The source side is populated from the uW source master branch.
4. The target side may be changed and saved to the translator's branch until done.
5. Translator finishes, saves, and a PR is used to merge this new file into the target master branch.
6. Later, the translator wants to "edit" the translated file.
7. Again the file is selected from the list of files from the  uW source master repo branch.
8. This time, since the target master branch has content for this file, it will be used to populate the translator's branch.
9. Changes are made, saved, and a PR is used to merge the updates back to the target master branch.


### Editor Chronology

This case describes an update to the uW *English* source repo.

*Reminder: for an editor the source and target are the same org and repo. Thus the adjectives "source" and "target" are dropped in the description below.*

1. Editor selects a file to translate from the list of files from the uW master branch.
2. Branch is created in repo and file is created in the editor's branch and is populated the content from the uW master branch.
3. Both the left (source) side is populated from the repo master branch.
4. The editor's side content may be changed and saved until the editors's work is done.
5. Editor finishes, saves, and a PR is used to merge the updates into the uW master branch.

## Future Catalog Next Process

### Translator Chronology

1. Translator selects a file to translate from the list of files in the uW catalog. The list shown will always be from the latest English production published content in the uW org.
2. Branch is created in target repo and the file is created in a  branch with the English content from the catalog.
3. Translator may continue to make changes and save them until the work is done.
4. Translator finishes, saves, and a PR is used to merge this new file into the target master branch.
5. Later, the translator wants to "edit" the translated file.
6. Again the file is selected from the list of files in the uW catalog.
7. This time, since target master branch has content for this file, it will be used to populate the translator's target branch.
8. The source side is populated from the uW production English catalog.
9. Changes are made, saved, and a PR is used to merge the updates back to the target master branch.

### Editor Chronology

This case describes an update to the uW *English* source repo.

*Reminder: for an editor the source and target are the same org and repo. Thus the adjectives "source" and "target" for "repo" are dropped in the description below.*

1. Editor selects a file to translate from the list of files in the uW production English catalog.  See Note #1 about added content, such as a new Translation Word article.
2. Branch is created in repo and the file is created in a  branch with the English content from the repo.
3. The source side is populated from the uW production English catalog.
4. The Editor may continue to make changes and save them until the work is done.
5. Editor finishes, saves, and a PR is used to merge the changed file into the master branch.

**Note #1** To add new content, the editor must something like this:
1. Create the content manually.
2. Add it directly to the uW English repo master branch. 
3. Once even a partial file is in the master branch and is released to catalog next, then it will be available for further editing within tc-create.
