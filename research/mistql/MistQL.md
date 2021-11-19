# MistQL
Installed CLI via: `npm install -g mistql`. This adds the command `mq`.

## Querying git trees for en_tw

API JSON downloaded from this API call:
https://qa.door43.org/api/v1/repos/unfoldingword/en_tw/git/trees/master?recursive=true&per_page=99999

**Extract just the bible folder**
From the raw input, just take the tree array. See `tree-entries.json` for what this will look like. Then query for an entry:
```
$ mq '@ | filter path == "bible/names/rebekah.md"' tree-entries.json
```
The above returns:
```json
[
  {
    "path": "bible/names/rebekah.md",
    "mode": "100644",
    "type": "blob",
    "size": 1783,
    "sha": "9d9a1dcbc0f40a9367edbbe98b1fd120877396ef",
    "url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_tw/git/blobs/9d9a1dcbc0f40a9367edbbe98b1fd120877396ef"
  }
]
```

A missing entry returns:
```
$ mq '@ | filter path == "bible/names/rebekah.mdx"' tree-entries.json
[]
```

A better way... just name the attribute, here "tree", and get all its children:
```
$ mq 'tree | filter path == "bible/names/rebekah.md"' git-trees-en_tw.json 
[
  {
    "path": "bible/names/rebekah.md",
    "mode": "100644",
    "type": "blob",
    "size": 1783,
    "sha": "9d9a1dcbc0f40a9367edbbe98b1fd120877396ef",
    "url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_tw/git/blobs/9d9a1dcbc0f40a9367edbbe98b1fd120877396ef"
  }
]
```

## Querying the TD Language file

Downloaded via: https://td.unfoldingword.org/exports/langnames.json

**Count the number of languages in the file**
```
$ mq "count @" langnames.json
8999
```

**Match on English** (output shortened)
```sh
$ mq '@ | filter ang == "English"' langnames.json
[
  {
    "ang": "English",
    "ln": "English",
    "pk": 1747,
    "lr": "Europe",
    "gw": true,
    "cc": [
      "AG",
      "AO",
...elided...
    ],
    "alt": [
      "Anglit",
      "Kiingereza",
...elided...
    ],
    "hc": "GB",
    "lc": "en",
    "ld": "ltr"
  }
]
```

**List Gateway Languages**
```sh
$ time mq '@ | filter gw == true | groupby ang | keys' langnames.json
[
  "",
  "Amharic",
  "Arabic",
  "Arabic - Dominant Culture Variant",
  "Assamese",
  "Bengali, Bangla",
  "Bislama",
  "Brazilian Portuguese",
  "Burmese",
  "Cebuano",
  "Chinese",
  "English",
  "French",
  "Gujarati",
  "Hausa",
  "Hindi",
  "Hindi - Dominant Culture Variant",
  "Ilocano",
  "Indonesian",
  "Kannada",
  "Khmer",
  "Laotian",
  "Latin American Spanish",
  "Malagasy, Plateau",
  "Malay",
  "Malay, Papuan",
  "Malayalam",
  "Marathi (Marāṭhī)",
  "Mongolian",
  "Nepali",
  "Odia",
  "Pashto, Pushto",
  "Persian (Farsi)",
  "Persian, Iranian",
  "Portuguese",
  "Punjabi, Eastern",
  "Russian",
  "Spanish",
  "Swahili",
  "Tagalog",
  "Tamil",
  "Telugu",
  "Thai",
  "Tok Pisin",
  "Turkish",
  "Urdu",
  "Urdu - Dominant Culture Variant",
  "Vietnamese"
]

real    0m0.531s
user    0m0.061s
sys     0m0.122s
```

**Count Gateway Languages**
```sh
$ time mq '@ | filter gw == true | groupby ang | keys | count' langnames.json
48

real    0m0.505s
user    0m0.152s
sys     0m0.105s
```

**How many GW languages are missing the "ang" value?**
```sh
$ time mq '@ | filter gw == true | filter ang == "" | count' langnames.json
2

real    0m0.487s
user    0m0.045s
sys     0m0.168s
```

**Which Gateway Languages are missing the "ang" value?**
```sh
$ time mq '@ | filter gw == true | filter ang == "" | groupby ln | keys' langnames.json
[
  "Bahasa Indonesia - Dominant Culture Variant",
  "Oral Mandarin Chinese"
]

real    0m0.502s
user    0m0.091s
sys     0m0.135s
```