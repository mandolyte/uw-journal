# Notes

## Links
- The CNTR website: https://greekcntr.org/collation/index.htm
- Github: https://github.com/Center-for-New-Testament-Restoration 
- In-development website: https://github.com/Center-for-New-Testament-Restoration/website-app
	- the above includes two frameworks in process: NextJs and Solid
	- there is a data zip and data folder with the JSON files from the source CNTR database

## Data Files

See sample data for Matthew 1:1 below.
1. There is one JSON file per verse (7957 of them per Textus Receptus at least)
2. Each file has a "verse", such as 40001025. This number encodes the following:
  - 40 is the 40th book of the Bible (Matthew in this example)
  - 001 encodes the chapter
  - 025 is the verse in the chapter
3. There is an array of "collations" after the verse
4. There is one collation per word in the verse. Each word has a number of attributes that classify the word.
5. 




# Sample for Matthew 1:1
```js
{
  "verse": "40001001",
  "collation": [
    {
      "CollationID": 40001001001,
      "VerseID": 40001001,
      "VariantID": null,
      "Relation": null,
      "Pattern": null,
      "Align": null,
      "Span": null,
      "Probability": 100,
      "Classic": "biblos",
      "Koine": "βιβλοσ",
      "Medieval": "βίβλος",
      "LexemeID": 9760,
      "Role": "N",
      "Morphology": "....NFS",
      "GlossPre": null,
      "GlossHelper": "the",
      "GlossWord": "scroll",
      "GlossPost": null
    },
    {
      "CollationID": 40001001002,
      "VerseID": 40001001,
      "VariantID": null,
      "Relation": null,
      "Pattern": null,
      "Align": null,
      "Span": null,
      "Probability": 100,
      "Classic": "genesews",
      "Koine": "γενεσεωσ",
      "Medieval": "γενέσεως",
      "LexemeID": 10780,
      "Role": "N",
      "Morphology": "....GFS",
      "GlossPre": "of",
      "GlossHelper": "the",
      "GlossWord": "birth",
      "GlossPost": null
    },
    {
      "CollationID": 40001001003,
      "VerseID": 40001001,
      "VariantID": null,
      "Relation": null,
      "Pattern": null,
      "Align": null,
      "Span": null,
      "Probability": 100,
      "Classic": "ihsou",
      "Koine": "=ιυ",
      "Medieval": "Ἰησοῦ",
      "LexemeID": 24240,
      "Role": "N",
      "Morphology": "....GMS",
      "GlossPre": "of",
      "GlossHelper": null,
      "GlossWord": "Jesus",
      "GlossPost": null
    },
    {
      "CollationID": 40001001004,
      "VerseID": 40001001,
      "VariantID": null,
      "Relation": null,
      "Pattern": null,
      "Align": null,
      "Span": null,
      "Probability": 100,
      "Classic": "cristou",
      "Koine": "=χυ",
      "Medieval": "Χριστοῦ",
      "LexemeID": 55470,
      "Role": "N",
      "Morphology": "....GMS",
      "GlossPre": null,
      "GlossHelper": null,
      "GlossWord": "Christ",
      "GlossPost": null
    },
    {
      "CollationID": 40001001005,
      "VerseID": 40001001,
      "VariantID": null,
      "Relation": null,
      "Pattern": null,
      "Align": null,
      "Span": null,
      "Probability": 100,
      "Classic": "uiou",
      "Koine": "υιου",
      "Medieval": "υἱοῦ",
      "LexemeID": 52070,
      "Role": "N",
      "Morphology": "....GMS",
      "GlossPre": null,
      "GlossHelper": null,
      "GlossWord": "son",
      "GlossPost": null
    },
    {
      "CollationID": 40001001006,
      "VerseID": 40001001,
      "VariantID": null,
      "Relation": null,
      "Pattern": null,
      "Align": null,
      "Span": null,
      "Probability": 100,
      "Classic": "dauid",
      "Koine": "δαυειδ",
      "Medieval": "Δαυὶδ",
      "LexemeID": 11380,
      "Role": "N",
      "Morphology": "....gms",
      "GlossPre": "of",
      "GlossHelper": null,
      "GlossWord": "David",
      "GlossPost": null
    },
    {
      "CollationID": 40001001007,
      "VerseID": 40001001,
      "VariantID": null,
      "Relation": null,
      "Pattern": null,
      "Align": null,
      "Span": null,
      "Probability": 100,
      "Classic": "uiou",
      "Koine": "υιου",
      "Medieval": "υἱοῦ",
      "LexemeID": 52070,
      "Role": "N",
      "Morphology": "....GMS",
      "GlossPre": null,
      "GlossHelper": null,
      "GlossWord": "son",
      "GlossPost": null
    },
    {
      "CollationID": 40001001008,
      "VerseID": 40001001,
      "VariantID": null,
      "Relation": null,
      "Pattern": null,
      "Align": null,
      "Span": null,
      "Probability": 100,
      "Classic": "abraam",
      "Koine": "αβρααμ",
      "Medieval": "Ἀβραάμ",
      "LexemeID": 110,
      "Role": "N",
      "Morphology": "....gms",
      "GlossPre": "of",
      "GlossHelper": null,
      "GlossWord": "Abraham",
      "GlossPost": null
    }
  ],
  "variant": [
    { "VariantID": null, "Relation": null, "Pattern": null },
    { "VariantID": null, "Relation": null, "Pattern": null },
    { "VariantID": null, "Relation": null, "Pattern": null },
    { "VariantID": null, "Relation": null, "Pattern": null },
    { "VariantID": null, "Relation": null, "Pattern": null },
    { "VariantID": null, "Relation": null, "Pattern": null },
    { "VariantID": null, "Relation": null, "Pattern": null },
    { "VariantID": null, "Relation": null, "Pattern": null }
  ],
  "matrix": {
    "0G0SR": [
      { "Column": 0, "Word": "¶Βίβλος", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενέσεως", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "˚Ἰησοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "˚Χριστοῦ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "Δαυὶδ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "Ἀβραάμ:", "Hand": "", "Version": "", "Error": null }
    ],
    "0G1WH": [
      { "Column": 0, "Word": "Βίβλος", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενέσεως", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "Ἰησοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "Χριστοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "Δαυεὶδ", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "Ἀβραάμ.", "Hand": "", "Version": "", "Error": null }
    ],
    "0G2NA": [
      { "Column": 0, "Word": "βιβλος", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενεσεως", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "ιησου", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "χριστου", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "δαυιδ", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "αβρααμ", "Hand": "", "Version": "", "Error": null }
    ],
    "0G3SBL": [
      { "Column": 0, "Word": "Βίβλος", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενέσεως", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "Ἰησοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "χριστοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "Δαυὶδ", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "Ἀβραάμ.", "Hand": "", "Version": "", "Error": null }
    ],
    "0G4BHP": [
      { "Column": 0, "Word": "¶βίβλος", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενέσεως", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "˚Ἰησοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "˚Χριστοῦ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "Δαυεὶδ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "Ἀβραάμ:", "Hand": "", "Version": "", "Error": null }
    ],
    "0G4SR": [
      { "Column": 0, "Word": "¶Βίβλος", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενέσεως", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "˚Ἰησοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "˚Χριστοῦ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "Δαυὶδ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "Ἀβραάμ:", "Hand": "", "Version": "", "Error": null }
    ],
    "0G4TH": [
      { "Column": 0, "Word": "Βίβλος", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενέσεως", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "Ἰησοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "χριστοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "Δαυεὶδ", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "Ἁβραάμ.", "Hand": "", "Version": "", "Error": null }
    ],
    "0G5RP": [
      { "Column": 0, "Word": "Βίβλος", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενέσεως", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "Ἰησοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "χριστοῦ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "Δαυίδ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "Ἀβραάμ.", "Hand": "", "Version": "", "Error": null }
    ],
    "0G6ST": [
      { "Column": 0, "Word": "Βίβλος", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενέσεως", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "Ἰησοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "Χριστοῦ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "Δαβὶδ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "Ἁβραάμ.", "Hand": "", "Version": "", "Error": null }
    ],
    "0G7KJTR": [
      { "Column": 0, "Word": "Βίβλος", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενέσεως", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "Ἰησοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "Χριστοῦ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "Δαυίδ,", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υἱοῦ", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "Ἀβραάμ.", "Hand": "", "Version": "", "Error": null }
    ],
    "1G10001": [
      { "Column": 0, "Word": "βιβλοσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενεσεωσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "=ιυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "=χυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "=υυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "δαυιδ", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "~υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "αβρααμ", "Hand": "", "Version": "", "Error": null }
    ],
    "1G20001": [
      { "Column": 0, "Word": "βιβλοσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενεσεωσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "=ιυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "=χυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "=δαδ", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "αβρααμ", "Hand": "", "Version": "", "Error": null }
    ],
    "1G20003": [
      { "Column": 0, "Word": "βιβλοσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενεσεωσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "=ιυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "=χυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "δαυειδ", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "αβρααμ", "Hand": "", "Version": "", "Error": null }
    ],
    "1G20032": [
      { "Column": 0, "Word": "βιβλοσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενεσεωσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "=ιυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "=χυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "δα%υ^ε^ι%δ", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "αβρααμ", "Hand": "", "Version": "", "Error": null }
    ],
    "2G061617": [
      { "Column": 0, "Word": "βιβλοσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενεσεωσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 2, "Word": "ιησου", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "χριστου", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "δαυετ", "Hand": "", "Version": "", "Error": "dau[e:i]t" },
      { "Column": 6, "Word": "υιου", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "αβρααμ", "Hand": "", "Version": "", "Error": null }
    ],
    "2G064853": [
      { "Column": 0, "Word": "βιβλοσ", "Hand": "", "Version": "", "Error": null },
      { "Column": 1, "Word": "γενεσενσ", "Hand": "", "Version": "", "Error": "genese[w:n]s" },
      { "Column": 2, "Word": "=ιυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 3, "Word": "=χυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 4, "Word": "=υυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 5, "Word": "=δαδ", "Hand": "", "Version": "", "Error": null },
      { "Column": 6, "Word": "=υυ", "Hand": "", "Version": "", "Error": null },
      { "Column": 7, "Word": "αβρα^α^μ^", "Hand": "", "Version": "", "Error": null }
    ]
  }
}
```
