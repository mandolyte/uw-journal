# Issue Last Row
## Test 1

Steps:
1. go to tN for Ruth in English, in uW
2. go to page 2 of tN
3. this will be row 26: Ruth 1:9 and ID=pm6y
4. click into the occurrence note to indicate an edit
5. console shows:

```
searchClose, pagination, saveRowId= false true header-1-9-pm6y
pagination true
```

6. click search and enter "moab", which is in page 1 row 1.
7. next click to close search

Result:
- it scrolls to the correct location
- then after 10s, pagination is turned back on
- then display is set to last row of page 1
- why?


## Test 2
(alt search b518)
Steps:
1. go to tN for Ruth in English, in uW
2. go to page 3 of tN
3. this will be row 51: Ruth 1:17 and ID=abc2
4. click into the occurrence note to indicate an edit
5. console shows:

```
searchClose, pagination, saveRowId= false true header-1-17-abc2
pagination true
```

6. click search and enter "moab", which is in page 1 row 1.
7. next click to close search

Result:
- it scrolls to the correct location
- then after 10s, pagination is turned back on
- then display is set to last row of page 1
- why?

##  Test 3
Steps:
1. go to tN for Ruth in English, in uW
2. go to page 3 of tN (rows 76 to 100 of 269)
3. scroll down to rowId "ab11", which is for Ruth 2:3
4. click into the occurrence note to indicate an edit
5. console shows:

```
searchClose, pagination, saveRowId= false true header-2-3-ab11
```

6. click search and enter "b518"
7. click occurrence note for b518 and observe that rowId is *not* updated; we are sticking with the "last edited row" before the search, not something done during a search.
8. next click to close search

console shows:
```
searchClose, pagination, saveRowId= true false header-2-3-ab11
DataTable.js:497 Setting pagination to true!
DataTable.js:499 returning to page: 3
DataTable.js:501 Next, in 1s scroll to  header-2-3-ab11
```

Result:
- pagination is turned on
- page is reset back to page 3
- on a delay, the DOM is searched for the element with id `header-2-3-ab11`
- it scrolls to the correct location

