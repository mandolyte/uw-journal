# Issue 1243

## basic issue description


## 2022-04-21

first, cleared user branch:
https://qa.door43.org/es-419_gl/es-419_tn/

second, logged in with tc-create:
- org es-419_gl
- resource TN
- lang es-419


In the repo, there exist these files in master:
en_tn_01-GEN.tsv
en_tn_02-EXO.tsv
en_tn_08-RUT.tsv
en_tn_56-2TI.tsv
en_tn_57-TIT.tsv
en_tn_65-3JN.tsv
es-419_tn_57-TIT.tsv

I need to pick a book that **isn't** in master yet.

So lets go with Jude...

*Result:* problem occurs, namely, endless spinner

Here are the messages in console log (trimmed down):
```
xhr.js:166          GET https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv?noCache=0.3698261406395298&ref=cecil.new-tc-create-1 404 (Not Found)
xhr.js:166          GET https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv?noCache=0.3486190488840297 404 (Not Found)
xhr.js:166          GET https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv?noCache=0.19908593899115035&ref=cecil.new-tc-create-1 404 (Not Found)
_callee3$ @ useFileContent.js:251
tryCatch @ runtime.js:64
xhr.js:166          POST https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv 404 (Not Found)
xhr.js:166          GET https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv?noCache=0.8337425881622968 404 (Not Found)
_callee3$ @ useFileContent.js:251
_next @ useFile.js:114
Promise.then (async)
asyncGeneratorStep @ useFile.js:102
_next @ useFile.js:114
Promise.then (async)
asyncGeneratorStep @ useFile.js:102
_next @ useFile.js:114
(anonymous) @ useFile.js:121
(anonymous) @ useFile.js:110
(anonymous) @ useFile.js:597
xhr.js:166          POST https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv 422 (Unprocessable Entity)
_callee3$ @ useFileContent.js:251
tryCatch @ runtime.js:64
invoke @ runtime.js:281
(anonymous) @ runtime.js:117
asyncGeneratorStep @ useFileContent.js:28
_next @ useFileContent.js:50
Promise.then (async)
asyncGeneratorStep @ useFileContent.js:38
_next @ useFileContent.js:50
Promise.then (async)
asyncGeneratorStep @ useFileContent.js:38
_next @ useFileContent.js:50
(anonymous) @ useFileContent.js:57
(anonymous) @ useFileContent.js:46
(anonymous) @ useFileContent.js:274
_next @ useFile.js:114
(anonymous) @ useFile.js:121
(anonymous) @ useFile.js:110
(anonymous) @ useFile.js:597
xhr.js:166          POST https://qa.door43.org/api/v1/repos/Es-419_gl/es-419_tn/contents/en_tn_66-JUD.tsv 422 (Unprocessable Entity)
```

