

# tc-create-app.469

**Title:** A single click highlights an original language word when it is directly adjacent to a highlighted word #469

**Text:**
```
v1.0.5-rc.6  
Movie at [https://app.zenhub.com/files/191973535/a854ebfa-5912-4066-ab09-971d71dce3b9/download](https://app.zenhub.com/files/191973535/a854ebfa-5912-4066-ab09-971d71dce3b9/download)

1.  Open a tN project
2.  Scroll to a check that has an original language word highlighted
3.  Single click on a word adjacent to the highlighted word
4.  Note that the clicked on word is now also highlighted
5.  Single click on the word in step 3
6.  Note that the highlight on that word is removed
```

A comment later clarifies:
```
this is specific to 'kai' this doesn't happen with other words. Doesn't even have to be adjacent. Just Kai'
```

Here is the case cited (Titus 1:3):
```
1:3 ἐφανέρωσεν δὲ καιροῖς ἰδίοις τὸν λόγον αὐτοῦ ἐν κηρύγματι, ὃ ἐπιστεύθην ἐγὼ, κατ’ ἐπιταγὴν τοῦ Σωτῆρος ἡμῶν, Θεοῦ;
```

## 2021-02-17

This morning I had a thought... why is there this onClick() function in the first place? The words highlighted come from the Original Quote. Perhaps this was part of an abandoned attempt to make it bi-directional, that is, highlight with cursor/mouse and update the Original Quote value. But that does not work. I don't see any value in what this onClick() does. So this morning, I'm going to test commenting it out and see what happens...



## 2021-02-16

Question from yesterday: *Where does the Boolean "selected" get set?*

Appears to be in function `areSelected()`:

```js
export const areSelected = ({ words, selections }) => {
  console.log("areSelected() words, selections", words, selections);
  let selected = false;
  const _selections = words.map((word) => selectionFromWord(word));
  console.log("areSelected() _selections:", _selections);

  _selections.forEach((selection) => {
    //if (selections.includes(_s)) selected = true;
    const _selection = JSON.parse(selection);
    let _text = normalizeString(_selection.text);
    let _occ = _selection.occurrence;
    let _occs = _selection.occurrences;
    console.log("areSelected() parsed selection:", _selection);

    for (let i = 0; i < selections.length; i++) {
      const text = selections[i].text; //already normalized.
      const occ = selections[i].occurrence;
      const occs = selections[i].occurrences;
      console.log("areSelected() for loop, selections[i]:", selections[i]);
      if (text === _text && occ === _occ && occs === _occs) {
        console.log("areSelected() match! selected is true, break");
        selected = true;
        break;
      }
    }
  });
  return selected;
};
```

NOTE: I have put a bunch of console messages in above. The number of times this executes is extraordinary. Could be an efficiency opportunity.

OK, notice the following:
1. in the above function, selected is set to false at the beginning
2. it is only set to true in the inner for loop. It is never set to false again.
3. therefore, once set to true, it should return from the function and break out of the outer for loop
4. the outer for loop is a forEach(). So convert this to a normal for loop and test after the inner loop finishes whether the Boolean is true.
5. if it is true, then break out of it as well -- we're done
6. then return it as it does now.

## 2021-02-15

Found where highlighting is done. It is done via CSS. In `AlignedWordObject.js`, these:

```js
const useStyles = makeStyles((theme) => ({
  open: { backgroundColor: 'lightgoldenrodyellow' },
  closed: {},
  popover: { pointerEvents: 'none' },
  paper: { padding: theme.spacing(1) },
  selected: { backgroundColor: 'yellow' },
}));
```

So when the class is "selected", the background is yellow. Not sure what the "open" class is supposed to be, but it is also a (different) shade of yellow.

The class is used in this "words" component built in the same source file:

```js
  const words = children.map((verseObject, index) => (
    <span
      data-test="aligned-word-object"
      data-testselected={selected}
      onClick={onClick}
      key={index}
      className={selected ? classes.selected : undefined}
    >
      <WordObject
        verseObject={verseObject}
        disableWordPopover={disableWordPopover}
      />
    </span>
  ));
```

It is triggered when a Boolean named `selected` is true. So the WordObject itself does not know if it is selected; instead that knowledge is wrapped in its span element.

So where does the Boolean selected get set?

## 2021-02-12

This occurs in Titus 1:3. In the tN there are two notes for this verse. In either one, in the Greek text (parallel scripture viewer), just single click on the word and it will highlight. This behavior is unique to this Greek word.

Using https://www.branah.com/unicode-converter, I entered: `κατ᾽` and the Unicode is `\u03ba\u03b1\u03c4\u1fbd`. So that character on the end is \u1fbd. 

Per https://codepoints.net/U+1FBD?lang=en, it is a "smooth breathing" accent, the Greek Koronis.


I asked on the Content stream for other Greek words ending in that character (I was unable to find one just scanning the text). If clone the Greek repo, this grep will find others to test: grep `\’\| *.usfm` 

Many cases are found. Using the tail, one from Rev:
```
1:12 καὶ ἐπέστρεψα βλέπειν τὴν φωνὴν ἥτις ἐλάλει μετ’ ἐμοῦ. καὶ ἐπιστρέψας, εἶδον ἑπτὰ λυχνίας χρυσᾶς,
```


In the scripture-resources-rcl, [here](http://localhost:6060/#/Parallel%20Scripture%20?id=parallelscripture)
```js
const defaultResourceLinks = [
  'unfoldingWord/el-x-koine/ugnt/master/tit',
  'unfoldingWord/en/ult/v5/tit',
  'unfoldingWord/en/ust/v5/tit',
  'ru_gl/ru/rlob/master/tit',
  'https://git.door43.org/unfoldingWord/en_ust/src/branch/master',
];

const _resourceLinks = [...defaultResourceLinks];

const reference = { bookId: 'tit', chapter: 1, verse: 3 };
```

The above mods will show the verse in Titus with the problem.

Here is screen shot before the problem, with `καιροῖς ἰδίοις` as the original quote:

![[Pasted image 20210212090052.png]]

Here is screen shot after clicking on  `κατ᾽`:

![[Pasted image 20210212090304.png]]

This code in the "onClick()" does fire when you click the words in the verse. But only the words with the Koronis character at the end will highlight. So perhaps the bug is in the addSelections() code...

```js
  if (_selectionsContext) {
    const {
      state: selections,
      actions: { areSelected, addSelections, removeSelections },
    } = _selectionsContext;
    selected = areSelected(originalWords);
    onClick = () => {
      if (selected) {
        console.log("AlignedWordsObject(), onClick(), removeSelections() originalWords:", originalWords);
        removeSelections(originalWords);
      } else {
        console.log("AlignedWordsObject(), onClick(), addSelections() originalWords:", originalWords);
        addSelections(originalWords);
      }
    };
  }

  const words = children.map((verseObject, index) => (
    <span
      data-test="aligned-word-object"
      data-testselected={selected}
      onClick={onClick}
      key={index}
      className={selected ? classes.selected : undefined}
    >
      <WordObject
        verseObject={verseObject}
        disableWordPopover={disableWordPopover}
      />
    </span>
  ));

```

From above, addSelections() comes from the SelectionsContext. There, the actions come a custom hook useSelections():

```js
  let {state, actions} = useSelections({
      selections: selections,
      onSelections: onSelections,
      occurrence: occurrence,
      quote: quote,
      onQuote: onQuote,
      verseObjects: verseObjects,
  });
```

Here is the function in the hook:

```js
  const addSelections = (words) => {
    let _selections = helpers.addSelections({words, selections});
    update(_selections);
  };

```

Which calls a helper function in `./helpers.js`. Code is:

```js
export const addSelections = ({ words, selections }) => {
  let _selections = new Set(selections);

  words.forEach((word) => {
    const selection = selectionFromWord(word);
    _selections.add(selection);
  });
  return [..._selections];
};
```

![[Pasted image 20210212100258.png]]

Note that the "word" object for this word is a string, not an object. But same thing happens to other words... so can't be the problem.