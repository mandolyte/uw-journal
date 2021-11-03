# Serialize ULT

This is an experiment to import and serialize all the books of the bible using ULT in `https://github.com/mandolyte/proskomma-experiments` in `./data/en_ult` (which is a download from uW).

## Helpful Tidbits

How to import:
```js
    const text = await dcs.fetchBook('Door43-Catalog', 'en_ult', state.bookId);
    setContentStatus("Book Retrieved");
    pk.importDocument(
      {lang: "eng", abbr: state.bookId},
      "usfm",
      text
    );
```
The above is taken from the bible-ref-pk-demo code.

How to serialize:
```js
      let query = '{ docSets { id } }';
      let result = await pk.gqlQuery(query);
      const docSetId = result.data.docSets[0].id;
      const serialized = pk.serializeSuccinct(docSetId);
      // console.log(JSON.stringify(serialized, null, 2));
      const pk2 = new (Proskomma);
      pk2.loadSuccinctDocSet(serialized);
```
The above is taken from `proskomma-js` package in `test/code/serialize.js`.


```js
fse.writeFileSync('/home/mark/Desktop/succ.json', JSON.stringify(pk.serializeSuccinct('xxx_yyy'))); # Arg is docSetId
pk.loadSuccinctDocSet(fse.readJsonSync('/home/mark/Desktop/succ.json'));
```


## Journal

### 2021-04-07

Today, I fixed the script to remove indentation parameter from JSON stringify() function. That saved some file size.

Also tried compressing to see what the download size would be:

```sh
$ tar -czvf filename.tar.gz en_ult-succinct.json
$ tar -czvf filename2.tar.gz *.usfm
```
These two are 12.8MB and 4.6MB respectively.

*Working on loading the succinct file and then running a query against it.*

Script name is `load_query_ult.js`.



### 2021-04-06

Using 'mandolyte/proskomma-experiments' repo on github. There is a folder 'data/en_ult' that will be my first test.

I will borrow this test code as the starter:
```js
const pk2 = pkWithDocs([
  ['../test_data/usx/web_rut.usx', {
    lang: 'eng',
    abbr: 'webbe',
  }],
  ['../test_data/usx/web_psa150.usx', {
    lang: 'eng',
    abbr: 'webbe',
  }],
  ['../test_data/usfm/ust_psa.usfm', {
    lang: 'eng',
    abbr: 'ust',
  }],
  ['../test_data/usx/not_nfc18_phm.usx', {
    lang: 'eng',
    abbr: 'nnfc18',
  }],
]);
```

I will put the script in the same folder to keep things direct and simple.

Script is named "serialize_ult.js".

To run: `time node serialize_ult.js`

Time and size:
- using 0.4.2: GEN thru DEU: 25s, 8MB
- using 0.4.9: GEN thru DEU: 28s, 8MB

Time and size for 66 books: 11m13s, 109MB

```
$ time node serialize_ult.js 
Import time: 637
Query time: 0
Serialization time: 2
Stringify time: 0
File write time: 0

real    10m43.409s
user    0m0.000s
sys     0m0.015s
```
Size of file is 112307652 or about 112MB.