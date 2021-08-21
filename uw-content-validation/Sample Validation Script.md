# Sample script to use in tc-create

1. Go to: https://create.translationcore.com/
2. Login
3. For Organization, select unfoldingWord
4. For Resource, select Translation Notes
5. For Language, select English
6. For File, select the book of Ruth
7. Notice the new validate icon button on the toolbar (text + check mark)
8. Click it
9. The first time validation is used in a session it will cache everything it retrieves; So first time is slowest
10. An alert will appear indicating there are no problems
11. Click on the side navigation menu.
12. Notice there are three levels of validation: High priority, Medium priority, and "all".
	- High priority are those numbered 800 and above
	- Medium priority are those numbered 600 and above
	- And "ALL" returns all notices
13. Select the "all" option
14. Close side navigation and then click the validation again
15. This time a CSV of notices found will be downloaded; open the CSV and place it side-by-side with tc-create
16. In tc-create, use the column selection button to add ID
17. Then click the search button to locate row ID `ab01`
18. Fix the problem in one of the two possible ways
19. Do not save the file
20. Now click to validate again
21. Open the new CSV; notice the two notices about `ab01` are now resolved.