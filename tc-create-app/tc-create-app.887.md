# tc-create-app#887
## Issue Details
### Issue
The current GL list is hardcoded. The list should be retrieved from a central master list of languages.  The language drop down should be filtered to only include GLs.

### Dod
The language list displays the GLs from a uW master list of languages. 

### Details

From Eric Steggerda: For MENA GLs to start using tcCreate app, the app must display the list of Gateway Languages.

## Branch name
feature-cn-887-dynamic-language-fetch

## 2021-07-20

Have completed uw-languages-rcl and incorporating into tc-create.

## 2021-07-08

Have re-located the complete list of languages into the AppContext. It is part of the state and is named "languages". 

This also meant changes to the languages `helper.js` file.

## 2021-07-07

New plan... just fetch the JSON and leave all existing logic in place.

Show only languages for given org... this is the start (in Languages.js)

```js
  const appContext = useContext(AppContext);
  const getOrgLanguages = () => {
    console.log("appContext:", appContext)
  }

  getOrgLanguages();

```

## 2021-06-30

Forget about filtering for now... that will be another issue. For now, focus on making all the languages available. Since there is the CORS issue with `langnames.json`, I will just make the changes needed to show all the files in the UI. 

In `helpers.js`, two constants are exported:
- `gatewayLanguages`
- `languages`

The component `<LanguageSelect />` (in `LanguageSelect.js`) imports only "gatewayLanguages". Changes this to "languages".

## 2021-06-29

Per Zulip, we need to make a toggle to allow users to expand the language selection list to include all languages. So something like a toggle with text like "Show all languages"

## Research

This is a UI for the uW languages: 
https://td.unfoldingword.org/uw/languages/

This returns the JSON for same:
https://td.unfoldingword.org/exports/langnames.json



## JSON Details
Notice the "gw" (gateway) attribute is false for "Afar" ('aa'); whereas it is true for English.
```json
[
  {
    "cc": [
      "DJ",
      "ER",
      "ET",
      "US",
      "CA"
    ],
    "hc": "ET",
    "pk": 6,
    "alt": [
      "Afaraf",
      "Danakil",
      "Denkel",
      "Adal",
      "Afar Af",
      "Qafar",
      "Baadu (Ba'adu)"
    ],
    "lc": "aa",
    "lr": "Africa",
    "ld": "ltr",
    "ang": "Afar",
    "gw": false,
    "ln": "Afaraf"
  },
  {
    "cc": [
      "AG",
      "AO",
      "AR",
      "AU",
      "AW",
      "BB",
      "BD",
      "BE",
      "BM",
      "BN",
      "BW",
      "BZ",
      "CA",
      "CK",
      "CL",
      "CO",
      "CR",
      "CZ",
      "DE",
      "DO",
      "EC",
      "ES",
      "FJ",
      "FM",
      "FR",
      "GB",
      "GD",
      "GF",
      "GG",
      "GH",
      "GI",
      "GL",
      "GY",
      "HN",
      "HT",
      "IE",
      "IM",
      "IN",
      "IS",
      "IT",
      "JE",
      "JM",
      "KE",
      "KY",
      "LC",
      "LK",
      "LS",
      "LU",
      "MC",
      "MH",
      "MM",
      "MT",
      "MW",
      "MY",
      "MZ",
      "NA",
      "NF",
      "NL",
      "NP",
      "NZ",
      "PE",
      "PG",
      "PK",
      "RO",
      "SB",
      "SC",
      "SD",
      "SG",
      "SX",
      "SZ",
      "TH",
      "TT",
      "TZ",
      "UG",
      "US",
      "VE",
      "VG",
      "WS",
      "ZA",
      "ZM",
      "ZW",
      "CC"
    ],
    "hc": "GB",
    "pk": 1747,
    "alt": [
      "Anglit",
      "Kiingereza",
      "Gustavia English",
      "Samaná English",
      "Saint Lucian English",
      "Noongar",
      "Noonga",
      "Newcastle Northumber",
      "Neo-Nyungar (Noogar)",
      "Glaswegian",
      "Brummy",
      "Birmingham (Brummie)",
      "Bay Islands English",
      "Australian Standard English",
      "Aboriginal English",
      "African American Vernacular English (AAVE)"
    ],
    "lc": "en",
    "lr": "Europe",
    "ld": "ltr",
    "ang": "English",
    "gw": true,
    "ln": "English"
  },
	
```