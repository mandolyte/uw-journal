# Open Datasets

Started with the STEPBible data [here](https://github.com/STEPBible/STEPBible-Data/blob/master/TIPNR%20-%20Translators%20Individualised%20Proper%20Names%20with%20all%20References%20-%20STEPBible.org%20CC%20BY.txt)

... which:
- is a TSV file, 
- hand created, 
- with many to one relationships internally
- size is 3,125,646 bytes (over 3MB)

---
I wrote some code to process the file to create CSV files, that:
- had a common key
- separated the many and the one
- The two places files are almost 400KB
- The two persons files are almost 1.5MB
- Less that 2MB total

I did this for both the places and persons data.

---
Next I imported into SQLite. 
- Schema:
```
$ sqlite3 properNames.db 
SQLite version 3.34.1 2021-01-20 14:10:07
Enter ".help" for usage hints.
sqlite> .schema
CREATE TABLE IF NOT EXISTS "persons"(
  "UniqueName" TEXT,
  "UnifiedName" TEXT,
  "Description" TEXT,
  "Parents" TEXT,
  "Siblings" TEXT,
  "Partners" TEXT,
  "Offspring" TEXT,
  "TribeNation" TEXT,
  "Summary" TEXT
);
CREATE TABLE IF NOT EXISTS "persons_significance"(
  "UniqueName" TEXT,
  "Qualifier" TEXT,
  "Significance" TEXT,
  "Strongs" TEXT,
  "EsvName" TEXT,
  "References" TEXT
);
CREATE TABLE IF NOT EXISTS "places"(
  "UniqueName" TEXT,
  "Strongs" TEXT,
  "OpenBible" TEXT,
  "Founder" TEXT,
  "PeopleGroup" TEXT,
  "GoogleMapURL" TEXT,
  "PalopenmapsURL" TEXT
);
CREATE TABLE IF NOT EXISTS "places_significance"(
  "UniqueName" TEXT,
  "Qualifier" TEXT,
  "Significance" TEXT,
  "Strongs" TEXT,
  "EsvName" TEXT,
  "References" TEXT
);
sqlite> 
```
- Size of database is 2,039,808 bytes (about 2MB)

---
Compared to the JSON version of the file [here](https://github.com/PatristicTextArchive/tipnr_data/blob/master/tipnr_persons.json)
which is nearly 4MB in size!

![[Pasted image 20230216122042.png]]

---
So the Sqlite3 database is the smallest and it performs well. But how can it made useful to a web app? There are some pointers in that direction...

---

First, Sqlite3 now exists in Web Assembly (WASM), making it available to web applications

---

And there is the "Origin Private File System":
> This document defines fundamental infrastructure for file system APIs. In addition, it defines an API that makes it possible for websites to get access to a file system directory ...

From https://fs.spec.whatwg.org/

---

And finally there is the OPFS Explorer for developer types, a [Chrome Extension](https://chrome.google.com/webstore/detail/opfs-explorer/acndjpgkpaclldomagafnognkcgjignd) that allows you see the origin file system

---

An example that get part way there is:
- Project at https://github.com/rhashimoto/wa-sqlite
- Demo at https://rhashimoto.github.io/wa-sqlite/demo/
