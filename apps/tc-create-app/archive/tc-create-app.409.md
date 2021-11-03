# tc-create-app#409 Add TQ as TSV
Note this issue was blocked due to some acceptance issues that are not resolved as of 2021-02-11.

## 2021-02-11

The new spec is at:
https://forum.door43.org/t/draft-parascriptural-tab-separated-value-format-specification-v2/502

Here is the relevant language for Translation Questions:
```
These TSV files have FIVE columns: Reference, ID, Tags, Question, Response.

Tags - no (optional) tags are defined yet for TQ
Joel Ruark wrote: Grammar; Syntax; Vocabulary; Semantics; Discourse (just kinda thinking out loud here how we might categorize different types of tQ’s)
OBS_sq.tsv uses tags meaning (for “What the Story Says”) and application (for “What the Story Means to Us”)
Question - The Markdown formatted question itself.
The text should be Markdown formatted, which means the following are also acceptable:
Plaintext - if you have no need for extra markup, just use plain text in this column
HTML - if you prefer to use inline HTML for markup, that works because it is supported in Markdown
Response - The Markdown formatted (as above for Question) answer or response.
TQ repo
This will be in en_tq repo.

Contains one file per “book” (e.g., GEN_tq.tsv, PSA_tq.tsv, 3JN_tq.tsv, OBS_tq.tsv)

TO BE DECIDED: how to encode OBS in the manifest, and where to put the OBS version number, e.g., rc://en/obs/book/obs?v4 ???
```

**Validate in Production**
The newFormat is there per latest spec. QA DCS isn't ready yet. Should be there on Monday.

**NOTE!** Added a new document covering TSV, capturing requirements from the above spec. Also in Github under this issue.


Notes on datatable and markdown translatable components [here](DataTable-Markdown-translatable.md)

## Merge newFormat into Master in QA

This needs be done every week since QA is overwritten by Production
First:
- remove the cloned local folder for en_tq:
    - cd $HOME/Projects/qa.door43.org
    - rm -rf en_tq
- re-clone it: git clone https://qa.door43.org/unfoldingWord/en_tq.git
- then these commands in the cloned folder
```
git switch newFormat                 # back to the branch
git config merge.renameLimit 999999  # bump the limit
git merge master                     # try again
```

At this point, if there are conflicts, the conflicts must be removed by hand. These will be directories that must be removed.
As of 2021-01-25, these need to be removed:

git rm -fr 1ti 2sa 2ti 3jn deu eph est exo ezr isa jer jhn jon luk neh num oba psa rut

Next commit:
git commit -a -m "merged master into newFormat branch"
[newFormat 9e082877a] merged master into newFormat branch

Finally, push:
git push

Next on QA DCS:
- log in as tcc001
- create PR for the updated newFormat branch
- logout 
- log in as myself
- Go to PR, approve and squach/merge it
- Go to "files" tab to confirm that master now has TSV files

## Initial Issue Description

*Issue text* Coming out of the spike #360

This will require different display columns than the TSV for TNs.
The goal is to show the MD for the questions using the data in the TSV to index the questions according to the verse.

MD should be displayed the way TN displays but only display default column that are not empty.

The `<br>` to separate the question and answer should be separate lines with whitespace between them.

It will necessary to work with @RobH123 on learning the new TSV format.

## Notes

Working branches for new TSV format are always named "newFormat".

Here is the manifest entry for Titus:
```yaml
-
--
  | title: 'Titus'
  | versification: 'ufw'
  | identifier: 'tit'
  | sort: 56
  | path: './TIT_tq.tsv'
  | categories: [ 'bible-nt' ]
```

HTML link to Titus: https://qa.door43.org/unfoldingWord/en_tq/src/branch/newFormat/TIT_tq.tsv

Raw link: https://qa.door43.org/unfoldingWord/en_tq/raw/branch/newFormat/TIT_tq.tsv

Link to Specification (looks like it needs some updating):
https://forum.door43.org/t/draft-parascriptural-tab-separated-value-format-specification-v2/502

## TQ Requirements

First, the TQ TSV format only uses three columns:
- Reference
- ID
- Annotation

So in tc-create, only three columns should be available. Alternatively, if they are available, then it should be an error for them to have content.

The annotation columns is a specialized form of Markdown. The differences and conventions are:
- The question and answer are in the same Markdown.
- The question is followed by 5 characters `\n\n>`. The first two characters `\n` indicate a newline that ends the question. The second two characters `\n` indicates a blank line that should separate the question and answer. The fifth character indicates the beginning of a Markdown block quote.
- The above string must be pre-processed for preview:
    - The question and answer must be separated by a blank line
    - The answer must be treated as an HTML block quote (indented lines of text with a left vertical bar)


**Possible Issues**

1. The bi-directional translation between HTML and Markdown must be updated to change `\n` to/from `<br>`.
2. The blockquote visual aid (vertical bar) does not work in currently. It only indents.
3. The answer must be written in Markdown mode since there is no way to indicate a block quote in the HTML preview mode.
4. Does the CV package handle tQ TSV? And if so, does it generate a notice if Tags, SupportReference, Quote, or Occurrence have content?


## Scratch Space

Trace of:
```
VM131:38 Uncaught TypeError: Cannot read property 'split' of undefined
    at rowHeader (eval at evalInContext (evalInContext.js:19), <anonymous>:38:32)
    at HeaderCell (HeaderCell.js:20)
```

HeaderCell
- rowHeader is passed as property to HeaderCell
- - getColumns() in datatable/helpers.js; rowHeader is passed
- - - DataTable: rowHeader passed as an element of a config property
- - - - DataTableWrapper: rowHeader passed props
- - - - - DataTableContextProvider: passes it as props





PUT /api/v1/repos/unfoldingWord/en_tq/contents/1JN_tq.tsv HTTP/1.1
Host: bg.door43.org
Connection: keep-alive
Content-Length: 14939
sec-ch-ua: "Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"
Accept: application/json, text/plain, */*
DNT: 1
Authorization: token 2e22af20bbe730a418da654acba66bf96728b25c
sec-ch-ua-mobile: ?0
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36
Content-Type: application/json
Origin: http://localhost:3000
Sec-Fetch-Site: cross-site
Sec-Fetch-Mode: cors
Sec-Fetch-Dest: empty
Referer: http://localhost:3000/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9


GET /api/v1/repos/unfoldingWord/en_tq/contents/1JN_tq.tsv HTTP/1.1
Host: qa.door43.org
Connection: keep-alive
sec-ch-ua: "Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"
sec-ch-ua-mobile: ?0
DNT: 1
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Sec-Fetch-Site: none
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Fetch-Dest: document
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9
Cookie: _ga=GA1.2.1009756751.1599235306; __utmc=7818919; __utmz=7818919.1601562318.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utma=7818919.1009756751.1599235306.1603460311.1605538382.3; _gid=GA1.2.120751528.1611150592; lang=en-US; dcs_session=30c6adf0cec03593



      console.log("if condition values:")
      console.log(". filepath:",filepath);
      console.log(". sourceFile.filepath:",sourceFile.filepath);
      console.log(". targetFile.filepath:",targetFile.filepath);
