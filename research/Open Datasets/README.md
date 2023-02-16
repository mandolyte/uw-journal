# Links to open datasets

- Link to Discord channel: https://discord.com/channels/867746700390563850/1074708329445728287
*  [STEPBible-Data/TIPNR - Translators Individualised Proper Names with all References](https://github.com/STEPBible/STEPBible-Data/blob/master/TIPNR%20-%20Translators%20Individualised%20Proper%20Names%20with%20all%20References%20-%20STEPBible.org%20CC%20BY.txt)
    * JSON version that may be easier to work with: [PatristicTextArchive/tipnr_data](https://github.com/PatristicTextArchive/tipnr_data) 
* Robert Rouse has collected data on individuals and people groups (as well as other entity types) at [Viz.Bible | Bible Data](https://viz.bible/bible-data/). [Digital Manna](https://www.digitalmanna.org/) has taken this data for their [Digital Manna's Encyclopedia](https://www.digitalmanna.org/encyclopedia/people) and improved it. 
* [Freely-Given-org/Bible_speaker_identification](https://github.com/Freely-Given-org/Bible_speaker_identification) combines data from STEPBible, Robert Rouse, and SIL Glyssen. 
* [Clear-Bible/speaker-quotations](https://github.com/Clear-Bible/speaker-quotations) identifies biblical people associated with reported speech, based on data from [Faith Comes By Hearing](https://www.faithcomesbyhearing.com/). There are also person identifiers in the [MACULA Greek](https://github.com/Clear-Bible/macula-greek) and [Hebrew](https://github.com/Clear-Bible/macula-hebrew/) syntax trees, taken from UBS MARBLE data.
* https://github.com/Freely-Given-org/Bible_speaker_identification/tree/main/outsideSources This was trying to take the best of three Bible person DBs from SIL, Viz.Bible, and StepBible.

## 2023-02-14

Counts:
```
$ grep "^\$===" properNames.txt | sort | uniq -c
    115 $========== OTHER
      1 $========== OTHER
   3142 $========== PERSON(s)
   1022 $========== PLACE
$ 
```

Created a repo to continue work at:
https://github.com/mandolyte/learnathon-2023