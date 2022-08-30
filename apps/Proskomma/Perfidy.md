# Perfidy

## 2022-07-27

This link: https://github.com/Proskomma/proskomma-json-tools/blob/main/test/test_data/validation/valid_flat_document.json
has examples of inserted meta_content.

- under the paragraph array is a content array.
- the array has mixed types
- if the type is a string, then that is just the bare content, which are words in the text without any annotation
- if the type is an object, then the object will have a content array attribute where the words in the text will be found
- the other attributes in the object contain other information (such as meta content) associated with the content

## 2022-07-26

### Unique Word Checks
Issue: https://github.com/Proskomma/perfidy/issues/27

**Steps**

1. Make a copy of `wordFrequency.js` in src/transforms with name 'uniqueWords.js'
2. To `src/lib/stepTemplates.js` add import and entry for this new transform.
3. Result is `src/pipelines/uniqueWordsPipeline.json`

### a new pipeline to get PERF from USFM
Wrote my first pipeline:

```json
[
  {
    "id": 1,
    "title": "USFM from DCS",
    "type": "Source",
    "sourceLocation": "http",
    "outputType": "text",
    "httpUrl": "https://git.door43.org/unfoldingWord/en_ust/raw/branch/master/65-3JN.usfm"
  },
  {
    "id": 2,
    "title": "Selectors",
    "type": "Source",
    "sourceLocation": "local",
    "localValue": "{\"org\": \"dcs\", \"lang\": \"en\", \"abbr\": \"ust\"}",
    "outputType": "json"
  },
  {
    "id": 3,
    "title": "USFM to PERF",
    "name": "usfm2perf",
    "type": "Transform",
    "inputs": [
      {
        "name": "usfm",
        "type": "text",
        "source": "Source 1"
      },
      {
        "name": "selectors",
        "type": "json",
        "source": "Source 2"
      }
    ],
    "outputs": [
      {
        "name": "perf",
        "type": "json"
      }
    ],
    "description": "USFM=>PERF: Conversion via Proskomma"
  },
  {
    "id": 5,
    "title": "Convert USFM to PERF",
    "type": "Display",
    "inputType": "json",
    "inputSource": "Transform 3 perf"
  }
]

```

About as basic as it gets. It only converts usfm to perf. Then I could use the write to file to save it. So now I have 3JN in PERF format. Which might be useful.



## 2022-07-25

**Question #1** Isn't the goal of "results from any other transform" a bit broad? An extreme example would be the identity transform. Surely we would not want to enable this be metaContent. I think there are some transforms that are intended to be "about" the PERF (i.e., the data is *meta* data). So shouldn't limit metaContent to "meta data"?

**Question #2** I suppose meta content could be open ended, but perhaps there should be a standard format for meta content that can be validated by the PERF schema? Meta content doesn't need to have a fixed format to be "standard". I'm thinking something like the JSON used by many table components would work ok. So something like this:
```js
{
	name: "searchResults" // this allows multiple sets of meta content to be added into the PERF
	columns: ["Search Term", "Book", "Chapter", "Verse", "Occurence"]
	rows: [
		["wept", "GEN", "21", "16", 1],
		// lots of other hits
	]
}
```
So the schema validation would looks inside the meta content section of the PERF and expect to find three attributes, with types:
- name: `string`
- columns: `string[]`
- rows: `any[any[]]`

Also, do you have links to existing ideas, partially developed code, etc that I can study?

**Answers** are over in discord. The main idea is that the metaContent will be the actual snippets (chapter and verse) that are a result of a transform.

