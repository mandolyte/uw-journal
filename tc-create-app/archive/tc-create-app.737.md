---
id: b04135f3-5ac7-4556-9457-c003b1f6eb32
title: '737'
desc: ''
updated: 1618670374070
created: 1618670374070
isDir: false
---
# tc-create-app#737

Original text of issue:

```
Subject: # On-open validator reports errors on ru tN tsv files
Body:
Log in and select the Russian tN project for Titus.  
Note that the on-open validator reports an invalid header, and that line 199 only contains 7 columns. Upon examining the file, I don't see any problem with the header, and all rows have 8 tabs in them.
```

Here is the main comment I had made:

```
It isn't obvious what the issue is with the header. But here is the issue with line 199:

-   in the past, we've had issues with a blank line at the end of the file
-   thus in some cases we trim whitespace off the end of the file to prevent this from causing issues
-   in the this issue, the file has empty columns on the last row.
-   which means it ends with a bunch of tab characters... which are "whitespace"... which are then trimmed
-   which means it has fewer columns than it is supposed to
-   and then the validator reports it

This could be patched pretty easily. However, at this point there are several problems with how we parse TSV files and I propose instead that we take a bit more time to create an official TSV parser for us to use across our tools.

In particular, I propose a lax implementation of the spec that Robert Hunt calls out [here](https://en.wikipedia.org/wiki/Tab-separated_values#Conventions_for_lossless_conversion_to_TSV).

I say "lax" because, we could easily fill out extra columns if they are missing from a row and attempt to fix rows that have too many columns (there are one or two things that can be attempted, but some cases may not be fixable).
```

## 2021-04-08

Step 0. remove my branch in `https://qa.door43.org/translate_test/ru_tn/branches`

Step 1. Reproduce the error

- launched the vs code workspace for tc-create
- ran `yalc remove --all`
- switched to develop branch of tc-create
- ran `git pull`
- ran `yarn install`
- ran `yarn start`
- login (cecil.new), translate_test, tN, Russian, Titus

Now go to `https://qa.door43.org/translate_test/ru_tn/branches` and find my branch.

Found that the header row is too long. The extra character is at the end of the row. But I am unable to see what it is. It isn't whitespace since trimming does not affect it.

So I used this:

```js
        // NOTE: there are cases where invisible characters are at the end 
        // of the row. This line ensures that the header row only has the
        // number of characters needed. Only then are they compared.
        rows[0] = rows[0].slice(0,tsvHeader.length);
        if (tsvHeader !== rows[0]) {
			...
        }
```

To take care of the other problem where whitespace was being trimmed off the end of the file, _including the empty cells on the end!_, I changed this:

```js
let tsvFile = state.content.trimEnd();
```

to this:

```js
let tsvFile \= state.content;
```

and then added to the loop the ability to skip over empty rows:

```js
            if ( rows[i] === undefined || rows[i] === '' ) {
              continue;
            }
```

