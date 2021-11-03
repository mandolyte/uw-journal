# book-package-app

# Issue 30 - add obs support

## 2020-10-15

Begin work on the new TWL file.
Link: https://git.door43.org/unfoldingWord/en_translation-annotations/src/branch/master/OBS/OBS_twl.tsv

It has the same seven columns as the other TSV files. Column four has "Support Reference", which 
has a resource container link to the tW article. For example: `rc://*/tw/dict/bible/kt/god`.

For context of unfoldingWord and English, this maps to:
https://git.door43.org/unfoldingWord/en_tw/raw/branch/master/bible/kt/god.md

Process notes:

- Fetch the OBS TSV file
- Loop thru column four, push each value onto an array
- Dedup by creating a Set for the array
- Loop thru Set, fetching and counting the words in each article.

## 2020-10-14

Today, started to put together the project pieces after studied more about 
TSV formats and the resources needed/expected for OBS support.

**Here are questions about the requirements below:**

- Should the files in the front and back folders be included?
- Please confirm that for tN and tQ, that only the words in the Annotation columns are to be counted.
- Please confirm that for the `twl` file, no words are to be counted in the file itself.

**Here are the requirements for uW's English OBS:**

- The OBS content itself is in: https://git.door43.org/unfoldingWord/en_obs/src/branch/master/content
    - The content consists of 50 markdown files and...
    - Two folders named front and back
    - The `front` folder contains:
        - title.md
        - intro.md
    - The `back` folder contains:
        - intro.md
- The resources for OBS are in https://git.door43.org/unfoldingWord/en_translation-annotations/src/branch/master/OBS
- These resources consist of:
    - OBS_tn.tsv
    - OBS_tq.tsv
    - OBS_twl.tsv

The translation notes resource is in TSV format. The last column is "Annotation". This is the only column on which BPA should do word counts. (TBC).

The translation questions resources is in TSV format. Again, the last column is "Annotation" and it the column to be counted.

The translation words link resource is in TSV format. This file aligns the tW articles to the text. No words are to be counted in this linking file itself.

There is a Support Reference column which has a resource container link to the article that should be counted. Articles should only be counted once and the articles can be referenced many times within this linking file.

An example link is: `rc://*/tw/dict/bible/kt/bless`. The only valid context for the book package app today is unfoldingWord's English content. Therefore the the location of the "bless" article will be a combination of:

- the link to the repo: `https://git.door43.org/unfoldingWord/en_tw/src/branch/master/bible` (note this is an HTML link, not the link used for the tool)
- the category folder: `kt` (there two other categories: `names` and `other`)
- the article itself: `bless.md`

The complete tool article link would be:
https://git.door43.org/unfoldingWord/en_tw/raw/branch/master/bible/kt/bless.md

NOTE! There are no translation academy articles associated the OBS, so no need to count any of this resource type.
Similarly, there are no Literal or Simplified Text resources.
