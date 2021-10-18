# TSV Support OKR

## 2021-10-12 "all tsv" branch merged... test!
- Testing with: https://develop--tc-create-app.netlify.app/
- For each resource type, first remove my branch from QA DCS (for example, https://qa.door43.org/unfoldingWord/en_tn/branches).
- Use TIT to test

NOTE: after the fix to tn, testing locally

- twl: worked, cv worked (no errors); lots of debug output so I file an issue at https://github.com/unfoldingWord/uw-content-validation/issues/224.
- tn: was hanging; fixed and made PR
- tq: worked, cv worked (no errors found)
- sq: worked, cv worked (no errors found)
- sn: worked, cv worked (no errors found)
- obs-tq: worked, cv worked (no errors found)
- obs-tn: worked, cv worked with errors found. I reported them on zulip [here](https://unfoldingword.zulipchat.com/#narrow/stream/214041-CONTENT--.20UR/topic/OBS.20TSV.20Migration/near/257378508)
- obs-sn: worked, but row header is missing; asked Elsy to confirm and write an issue for it; also saw a lot of debug messages and left RH an issue for that
- obs-sq: 

## 2021-10-11 Test Each TSV resource

### As an editor (Use `unfoldingWord` Org)
- TN (9 col) - fails with "incomingTable.map is not a function"; should avoid using tsv parser with old 9 col format. Let the old code run as-is. Done. Now works. CV works.
- TWL - works with master branch of datatable-translatable, but not with what is in NPM; cv works
- TQ - works; cv works; **asked @elsy about missing scripture pane**
- SQ - works; cv works
- SN - works; cv works
- OBS-TQ - after removing resource context from scripture-resources-rcl, it works.
- OBS-TN - 


https://qa.door43.org/api/v1/repos/unfoldingWord/en_obs-tq/contents/tq_OBS.tsv

## 2021-10-11 Add obs-tq

Copied the Bible TQ files:
```
$ cp TranslatableTqTSV.js TranslatableObsTqTSV.js
$ cp RowHeaderTq.js RowHeaderObsTq.js
```

*update RowHeaderzObsTq:*
- all RowHeaderTq to RowHeaderObsTq

*update TranslatableObsTq:*
- all RowHeaderTq to RowHeaderObsTq
- all "TqTSV" to "ObsTqTSV"

*update Translatable.js*
- add file name pattern
- add imports

## 2021-10-11 Create an "all tsv branch"

Today, I plan to merge all the TSV work into a single branch. This is getting a bit fragmented and very confusing.

Let's say the new combined branch is named `feature-cn-all-tsv`

Here are the branches to merge. Do not merge if a branch has already been merged into develop.

- starting from develop (which has some work already in it)
- feature-cn-409-spt-7col-en-tq (no) DONE (already merged)
- feature-cn-661-add-obs-sq (yes) DONE
- feature-cn-666-add-obs-tq (no) DONE
- feature-cn-788-impl-tsv-parser (yes) DONE
- feature-cn-953-add-sn-sq (yes) DONE
- feature-cn-twl-spt (no) DONE (already merged)
- refactor-cn-1008-cv-9col-tn (no) DONE (already merged)
- feature-cn-1030-add-obs-sn (yes) DONE ("already up to date")
- feature-cn-1031-add-obs-tn (yes) DONE


*First, switch to develop branch and do a pull.* The state of develop is as follows:
- CV code is behind; needs options argument
- TWL has refactored CV code, but is missing options code
- TSV (9col) has refactored CV code
- TQ does not have CV
- TN has CV but needs refactored code
- SQ has refactored CV code
- SN has refactored CV code
- Translatable has pattern matching for TN, TQ, SQ, SN, TWL, and 9col TSV (total of 6 resources)

*Second, create my super branch.* `set-branch feature-cn-all-tsv`

*Third, merge in 661.*
```
# from feature-cn-all-tsv
git switch feature-cn-661-add-obs-sq
git pull
git switch feature-cn-all-tsv
git merge feature-cn-661-add-obs-sq
```
Results:
- only conflict was the public buld_number; resolved that in favor of incoming
- this branch had both obs-sq and obs-sn

*Fourth, merge 666.*
```
# from feature-cn-all-tsv
git switch feature-cn-666-add-obs-tq
git pull
git switch feature-cn-all-tsv
git merge feature-cn-666-add-obs-tq
```
Results: nothing new came. I switched back to the branch and see nothing about obs-tq in it. OK, that's a bit weird.

*Fifth, merge 788 (tsv parser implementation).* Since stuff already in my branch uses the tsv parser, not expecting anything new...
```
# from feature-cn-all-tsv
git switch feature-cn-788-impl-tsv-parser
git pull
git switch feature-cn-all-tsv
git merge feature-cn-788-impl-tsv-parser
Auto-merging public/build_number
CONFLICT (content): Merge conflict in public/build_number
Auto-merging package.json
CONFLICT (content): Merge conflict in package.json
Automatic merge failed; fix conflicts and then commit the result.
```
Results: 
- resolve conflicts (2)

*Sixth, merge feature-cn-953-add-sn-sq.* Note, this had been merged, but I pushed more commits to it afterwards.
```
# from feature-cn-all-tsv
git switch feature-cn-953-add-sn-sq
git pull
git switch feature-cn-all-tsv
git merge feature-cn-953-add-sn-sq
```
Results: merged build_number conflict

*Seventh, merge feature-cn-1030-add-obs-sn.*
```
# from feature-cn-all-tsv
git switch feature-cn-1030-add-obs-sn
git pull
git switch feature-cn-all-tsv
git merge feature-cn-1030-add-obs-sn
```
Results: "Already up to date":
```
$ git merge feature-cn-1030-add-obs-sn
Already up to date.
```
Also, note that obs-sq is also present.

*Eighth, merge feature-cn-1031-add-obs-tn*
```
# from feature-cn-all-tsv
git switch feature-cn-1031-add-obs-tn
git pull
git switch feature-cn-all-tsv
git merge feature-cn-1031-add-obs-tn
```
Results:
- My local branch for 1031 had un-pushed commits; so I pushed them
- Lots of changes in this branch, including passing of validate options... good.

**Summary**
The new branch has the following TSV resources:
- TWL
- TN 9 col ("TSV")
- TQ
- TN 7 col
- SQ 
- SN
- obs-tn
- obs-sq
- obs-sn

Missing is obs-tq



## 2021-09-27 Analysis of Current State

Setup:
- tc-create branch `feature-cn-953-add-sn-sq`
- datatable-translatable branch `feature-cn-788-impl-tsv-parser` (via yalc)
- running locally with datatable linked in via yalc

Resources listed:
![image](/images/Pasted%20image%2020210927114907.png)

For each of the below, using organization 'unfoldingWord'... thus as an "editor".

### Bible Study Questions

- used Titus; rendered OK.
- Noticed there was a "on demand validation" button... does not work, crashes; Remove it?

### Bible Study Notes

- used Titus (actually, it is the only one available)
- Also had "on demand validation" button and it also crashes if clicked. Remove it?

### OBS Translation Questions

**Not in TSV format yet!** Markdown worked OK.

### OBS Translation Notes

**Not in TSV format yet!** Markdown worked OK.

### OBS Study Notes

**Not in TSV format yet!** Markdown worked OK.
