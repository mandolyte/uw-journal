# TSV Requirements

The new spec is [here](https://forum.door43.org/t/draft-parascriptural-tab-separated-value-format-specification-v2/502)

- [ ] The new TSV repos will include the OBS material. They will be treated like another bible book with the id "OBS".
- [ ] The filenames in the repos will be of the form BBB_xx.tsv, where 'BBB' is the book id (including OBS) and xx is the two letter abbreviation of the resource type.
- [ ] The resource abbreviations are: tn, tq, sn, sq, and twl.
- [ ] The repos will be named by language id and resource id. For example, en_tn, en_tq, en_sn, en_sq, and en_twl
- [ ] Note that the files will be listed in DCS in alphabetical order, since the numeric prefix for biblical order is no longer used.
- [ ] The TSV spec we will be implementing is described [here](https://en.wikipedia.org/wiki/Tab-separated_values#Conventions_for_lossless_conversion_to_TSV). And is defined [here](https://www.iana.org/assignments/media-types/text/tab-separated-values)
- [ ] The above spec requires that each line have the same number of tabs. *Should we consider re-laxing this requirement?*
- [ ] While not part of the IANA spec, there are several generally accepted escape sequences which we will support. These are: `\n, \t, \r, and \\`
- [ ] The older formats had one column each for Book, Chapter, and Verse. The new formats will not have the Book at all. Instead, it must be obtained from the name of the file. Chapter and Verse are combined into a single column called "Reference" with the form "c:v". For OBS, the meaning of verse is "frame".

## TQ and SQ

These TSV files have FIVE columns: Reference, ID, Tags, Question, Response. Since these are identical, one set of changes will work for both.

*Since these resource types are relatively stable, the cut in of the new format and the update to tc-create can be done independently. But, the cut in must happen first.*

## TWL

This resource is used to encode links to TWs that have been copied out of the UHB and UGNT (and eventually the TW links in those two resources will be deprecated).

These TSV files have SIX columns: `Reference`, `ID`, `Tags`, `OrigWords`, `Occurrence`, `TWLink`.

*Since this format has no predecessor (no equivalent today), it can be cut in at any time. And tc-create support can follow.*

## TN and SN

These TSV files have SEVEN columns: `Reference`, `ID`, `Tags`, `SupportReference`, `Quote`, `Occurrence`, `Note`.

*Since tc-create does support TN today in its old format (9 columns) *and* we support content validation of this resource type, the cut-in of this new format (7 columns) must be coordinated with both tc-create and uw-content-validation.*

