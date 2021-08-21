# Issue 646

Link: https://github.com/unfoldingWord/tc-create-app/issues/646

Steps to always reproduce:

1.  Close any open tC Create sessions
2.  Start a new tC Create session
3.  Log in and open a Titus tN project
4.  Open a new tab and start a second tC Create session
5.  Log in with the same user name as in step 3, and select the same org and language
6.  Open a Revelation project in tN
7.  Go back to the first tC Create session and open the drawer
8.  Choose the 2Peter project
9.  Note that the project does not load

## 2021-04-01

Setup: first go to https://qa.door43.org/translate_test/en_tn/branches
and delete my branch, making this a clean start.


After step 6...
- return to first tab
- open console: nothing amiss yet
- click the application tab and check persisted data:
![[Pasted image 20210401095333.png]]
- note that it is set to the book of Revelation... other key/values look normal
- click the network tab in console to record network traffic
- open the drawer and select 2PE
- debugger paused the app with:
- ![[Pasted image 20210401095633.png]]
- *unable to get any where at this point -- completely frozen*
- after a meeting, things became unstuck; but it doesn't work: seeing 403:
- ![[Pasted image 20210401104157.png]]
Here are the four links above:
https://bg.door43.org/api/v1/repos/translate_test/en_tn/contents/en_tn_62-2PE.tsv?noCache=0.8745014752012958&ref=cecil.new-tc-create-1

https://bg.door43.org/api/v1/repos/translate_test/en_tn/contents/en_tn_62-2PE.tsv?noCache=0.7530780172658673

https://bg.door43.org/api/v1/repos/translate_test/en_tn/contents/en_tn_62-2PE.tsv

https://bg.door43.org/api/v1/repos/translate_test/en_tn/contents/en_tn_62-2PE.tsv

This is followed by an error message that may be in our code:
`Uncaught (in promise) Error: Error creating file.`

![[Pasted image 20210401114551.png]]

