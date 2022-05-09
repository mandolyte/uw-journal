# PK Editor

Asked this question on Discord in the `proskomma-editors` channel. There is also a PDF write up of latest plans.

```
Cecil — Today at 12:27 PM
Does this describe the flow of things?
1. Book is imported by the editing tool
2. A block is queried for by the editor
3. The block is transformed by PK into a JSON format
4. The editor [may|will] need to transform the PK JSON into its own JSON format
5. The editor then transforms the JSON into USFM text
6. User makes changes and saves the block
7. The editor converts the USFM to JSON (possibly twice, see above)
8. The editor gives it back to PK as a mutation query (or directly as some edit function call?)
9. PK merges the alignment data back in, if possible; is a warning returned if the alignments are broken?
10. PK merges the block back in.

Have to ask one question... why wouldn't USFM itself be the format instead of JSON? 
```