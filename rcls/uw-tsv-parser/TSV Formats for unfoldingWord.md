# Notes on TSV 

## Links
Here is the unfinished draft of the spec (as of this writing - 2021-03-12):
https://forum.door43.org/t/parascriptural-tab-separated-value-format-specification-v2/870

Here is the Wikipedia article on TSV: https://en.wikipedia.org/wiki/Tab-separated_values

## Spec Highlights

In the section named "Various TSV File Formats", the details on how uW uses TSV is spelled out.

At low level, there is this:
```
`\n` (2-characters) is used for line breaks within fields as per the TSV spec [here](https://en.wikipedia.org/wiki/Tab-separated_values#Conventions_for_lossless_conversion_to_TSV) (we formerly used HTML `<br>`) – these `\n`s should be automatically converted to newLine in the low-level TSV read functions immediately after each row has been divided into fields, and vice versa when writing rows. In other words, a `\n` for example should never reach a markdown processor – it should already be a newLine character by then. It is recommended that all four escaped characters are implemented by this low-level software: `\n`, `\t`, `\r`, and `\\`.
```

