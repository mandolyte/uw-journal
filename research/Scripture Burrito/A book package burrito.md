# Book Package Burrito

This Github repo *is* a burrito:
https://github.com/bible-technology/sb_textStories

Notes:
1. It contains metadata.json file in the root which has the metadata for the info in this repo.
2. This JSON file points to the "content" folder in the "ingredients" [section](https://github.com/bible-technology/sb_textStories/blob/master/metadata.json#L226):

```json
  "ingredients": {
    "LICENSE.md": {
      "mimeType": "text/markdown",
      "checksum": {
        "md5": "97c254986effcb3f9b27ed2e8a19577b"
      },
      "size": 1653
    },
    "content/01.md": {
      "mimeType": "text/markdown",
      "checksum": {
        "md5": "eeaabe633a42492a6cf80c2f4223b116"
      },
      "size": 4396,
      "scope": {
        "GEN": ["1-2"]
      }
    },
```

In the document [Book Package Artifacts](https://docs.google.com/document/d/1kkY55d3lPETFBRFvhIeKpdpoIl7OHq8X3ap6PErg34M/edit), this is used to desribe a Book Package for Titus:

- LT: https://git.door43.org/unfoldingWord/en_ult/raw/tag/v40/57-TIT.usfm
- ST: https://git.door43.org/unfoldingWord/en_ust/raw/tag/v39/57-TIT.usfm
- TN: https://git.door43.org/unfoldingWord/en_tn/raw/tag/v65/en_tn_57-TIT.tsv
- TQ: https://git.door43.org/unfoldingWord/en_tq/raw/tag/v37/tq_TIT.tsv
- TWL:  https://git.door43.org/unfoldingWord/en_twl/raw/tag/v17/twl_TIT.tsv
- SQ: no releases yet – either empty or row omitted
- SN: no releases yet – either empty or row omitted

For a set of texts (LT and ST) with the resources (TN, TQ, and TWL) and, implied, the TA and TW resources pointed to withing the TN and TWL resources, these are the exact versions that are approved and published.

*So what would this data look like as a "book package burrito" (BP Burrito, hereafter)?*

The OBS burrito above contains these files/folders (ignoring the gh artifacts):
- metadata.json
- README.md
- LICENSE.md
- and the "content" folder.

## Option 1

Include the same resources and URLs in the metadata.json and avoid any actual content. Instead, expect the app to retrieve the linked data if/when needed.

## Option 2

Going beyond the above to actually store the published content with the burrito in a folder (ingredients?).
