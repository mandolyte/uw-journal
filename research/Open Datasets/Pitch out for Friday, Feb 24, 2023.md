# Team: Knowledge

Goal: Explore Open Datasets for "places" data

---

## Datasets used this week
- StepBible TIPNR - place data only, not person data
- Pleiades - historical place data (secular)
- unfoldingWord - Translation Word Lists, Translation Words, versification data
- Open Bible Geocoding - linked location/place data


---


## Methodology 

- Transform data to tabular (if in JSON)
- Import into Sqlite3 
- Iterate to evaluate key data to connect disparate datasets
- Clean and normalize the data
- Develop SQL queries to establish/verify linkages


---



## Demo 1 

Show linkages between Step Bible data and Translation Word Lists.

---

## Demo 2

Show linkages between  Step Bible and Pleiades data using a latitude and longitude approach.

---

## Demo 3

Show linkages between Step Bible and Open Bible Geo Data.

---

## Learnings

- A number of open datasets exist, each with their own "keys"
- Often these datasets include the "keys" to other datasets and this enables the datasets to be linked
- "book, chapter, verse" can often be used to link data in unforeseen ways (for example, nearly all unfoldingWord can be connected via BCV)
- JSON support in Sqlite3 can be used to create custom datasets tailored to an app's need

---

## Why not real time? (a)

- Doable, but with downsides:
	-  A lot of code
	-  Overhead of data retrieval
	-  In some cases, significant data munging in order to link



---

## Why not real time? (b)

- Datasets are relatively static:
	- StepBible data hasn't been updated in 10 months
	- OpenBible data hasn't been updated in 2 years
- Many of the datasets are in CSV/TSV format and sometimes are not even uniform


---

## Possible Next Steps

- Many of these datasets also have people data, but would take some amount of work "align" the datasets
- Explore use of Sqlite3 WASM & OPFS to distribute these  datasets