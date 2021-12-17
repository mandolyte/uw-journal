# Diary

## 2021-12-10 Card Refresh on Org/Lang Change

Here is org eo_gl with only one repo, en_glt created.
![[Pasted image 20211210075426.png]]

Now switch to es-419_gl:
![[Pasted image 20211210075623.png]]

The status values are the same... so the card is NOT refreshing as it should.

If I logout and log back in, doesn't help.

If I do a page refresh, which takes me back to the login page, then it works:
![[Pasted image 20211210080103.png]]


## 2021-12-08 Bugs in Add Book Impl

Yesterday, I created in QA DCS an new org called `eo_gl` (Esperanto). I then created two repos: en_gst and en_glt (language code `eo` is not showing in language list.)

**BUGS**
- Language codes are not downselected to what is available in the org.
- Whereas at this moment, when I view `https://qa.door43.org/eo_gl` only the two repos I created yesterday are showing, the admin app seems to think that TN, TW, and TQ are also OK. Even showing/implying that "Ruth" is available in them!? (see snippet below)
- Both TW and TA are showing the status "Fetch Error" - should be "Repo not found"
- Need to implement a way to view the error if the "view" icon button is clicked.
- In the card for Titus, it shows that the entry is in the manifest... probably because it is the default. Need to add the entry of the card's book into the manifest before creating it.

![[Pasted image 20211208075829.png]]



## 2021-12-07

How to make an "icon button":
```js
<Tooltip title="Validate">
  <IconButton className={classes.iconButton} onClick={onValidate} aria-label="Validate">
    <PlaylistAddCheck />
  </IconButton>
</Tooltip>
```

## 2021-11-17 TA Approach

The TA articles are in the support reference column of the tN TSV file. At present there are two flavors of this TSV: a 9 column legacy and a new 7 column version. The pattern of the filename will reveal which it is. Not every translation note has a support reference. 

### Legacy 9 column tN TSV
The 9column version has entries like this: `writing-newevent` in the 5th column. This is simply the name of the article. It is located at:
https://git.door43.org/unfoldingWord/en_ta/src/branch/master/translate/writing-newevent/01.md
NOTE! this is a folder containing three files:
- 01.md which is the body of the article
- sub-title.md
- title.md

### New 7 column tN TSV
The 7 column version will have the support reference in column 4. The new version will entries like this: `rc://*/ta/man/translate/writing-newevent`

The location and organization remains the same.

### Computing the URL to the article

Given this example:
https://git.door43.org/unfoldingWord/en_ta/raw/branch/master/translate/writing-newevent/01.md

The template will be something like:
`${server}/${owner}/${langid}_ta/raw/branch/master/translate/${articleName}/01.md`

### Testing for existence

Here is the git trees link:
https://qa.door43.org/api/v1/repos/unfoldingword/en_ta/git/trees/master?recursive=true&per_page=99999

There are only 944 entries; so this approach will work.

Here are the entries that make up the files for the above example:
```json
    {
      "path": "translate/writing-newevent/01.md",
      "mode": "100644",
      "type": "blob",
      "size": 7848,
      "sha": "de9fe4e9a5e0826bf0240286aa429e83be965df2",
      "url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ta/git/blobs/de9fe4e9a5e0826bf0240286aa429e83be965df2"
    },
    {
      "path": "translate/writing-newevent/sub-title.md",
      "mode": "100755",
      "type": "blob",
      "size": 43,
      "sha": "46c8b56a345f555538bbdf90c46494701a32dbaf",
      "url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ta/git/blobs/46c8b56a345f555538bbdf90c46494701a32dbaf"
    },
    {
      "path": "translate/writing-newevent/title.md",
      "mode": "100755",
      "type": "blob",
      "size": 27,
      "sha": "d5530743b439cfa858845c95d9cab3c94ae44cd4",
      "url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_ta/git/blobs/d5530743b439cfa858845c95d9cab3c94ae44cd4"
    },
```

Notice that the paths are from the root.

TBD: Need to find out how regex is supported by MistQL.

## 2021-11-16 TW Approach

- TW words from the TWL file. 
- The TW articles come from the, for example, the "en_tw" repo.
- Here is a link to identify all the articles:

https://qa.door43.org/api/v1/repos/unfoldingword/en_tw/git/trees/master?recursive=true&per_page=99999

NOTE: the total count of files is 1022, so this will easily get them all in one fell swoop.

The articles are under the bible directory, and in a subfolder, one of three. Here is an entry from the above link:
```js
    {
      "path": "bible/names/rebekah.md",
      "mode": "100644",
      "type": "blob",
      "size": 1783,
      "sha": "9d9a1dcbc0f40a9367edbbe98b1fd120877396ef",
      "url": "https://qa.door43.org/api/v1/repos/unfoldingWord/en_tw/git/blobs/9d9a1dcbc0f40a9367edbbe98b1fd120877396ef"
    },
```

Here is one of the thirty TWL TSV rows that point to the "rebekah" article (vertically):
```
22:23
cbmf
name
רִבְקָ֑ה
1
rc://*/tw/dict/bible/names/rebekah
```

Steps:
1. remove the rc prefix, namely: `rc://.*dict/` (a regex). This leaves `bible/names/rebekah`
2. add a `.md` suffix. This leaves `bible/names/rebekah.md`. 
3. This will now match exactly the form returned in the git trees api call above.

Consider using mistql to query the JSON to locate entry (or its absence). Sample query: 'tree | filter path == "bible/names/rebekah.md"'

## 2021-11-09

The last few days, serious work finally begun. Some decisions have been made and recorded in Discord on what we need to deliver first.
- no need for a tabbed interface yet... the sole functionality to be delivered is the "light repo validation".  Message link [here](https://discord.com/channels/867746700390563850/906161562287480902/907296957993734165)
- the light validation will be similar to what the cv app does

This initial work will be journalled over in the note for issue # 3

## 2021-08-24

One thing we need is a nice Next.js based shell/template to use for the tabbed UI the mockups have. Here is the mockup:

![[Pasted image 20210824113538.png]]

I certainly could use some Next.js experience. So folder "Next.js" will contain notes, etc.
