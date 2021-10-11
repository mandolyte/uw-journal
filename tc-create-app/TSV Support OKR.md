# TSV Support OKR

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
