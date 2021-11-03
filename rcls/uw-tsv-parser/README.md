# TSV Parser

## Issue 788
Link:
https://github.com/unfoldingWord/tc-create-app/issues/788

## Links
uW post by Robert Hunt which calls out the spec to follow:
https://forum.door43.org/t/parascriptural-tab-separated-value-format-specification-v2/870

The spec itself is described in this article:
https://en.wikipedia.org/wiki/Tab-separated_values

The IANA standard:
https://www.iana.org/assignments/media-types/text/tab-separated-values

## Basic Operation

First, the parser must support lossless conversions to TSV. The data in memory may contain tab characters, newlines, etc. But these are not permitted in the "as stored" format. Here are the conventions (from the Wikipedia article):
```
   \n for newline,
   \t for tab,
   \r for carriage return,
   \\ for backslash.
```

Note that the above are two characters, *not* escaped single characters.

### On Read: string to 2D array

*input*: a string containing the TSV file content
*output*: an object with three member attributes:
- header: which is an array containing the column headers
- data: which the 2D array with the TSV data
- errors: an array of integer pairs, the first being the row number and the second being the number of columns found (that does not match the number expected)

The returned object will be like this:
`{header: hvalue, data: dvalue, errors: evalue}`


The input file to be converted from TSV to a two dimensional array will be processed as follows:
- the file must provided as a string
- the string will be split into an array based on the newline character (an EOL of CR-LF will not be supported)
- this array will be then read one row at a time and each row split by the tab character
- the first row will determine the number of columns expected in all the rows.
- each column value will altered per lossless conventions above. For example, changing all occurrences of backslash-t to a tab character.
- then the columns will be added to output array
- if the number of columns in the row do not match the first row, then an error message will be generated containing the:
	- the row number (zero based)
	- the number of columns found (could be more or less than expected)
	- the number of columns expected
	- the content of the row (with commas substituted for tabs)
	- an explanatory message
	- the error message will be added to an array of error messages
- after the file is processed, 
	- if there were any parsing errors, an exception will be thrown
	- otherwise, the output 2D array is returned


### On Write: 2D array to string

*input*: a 2D array
*output*: an object containing two member attributes:
- data: being the TSV compliant string
- errors: being an array of integer triplets, where
	- the first integer is the row number
	- the second integer is the actual number of columns when it does not match the number in the header, i.e., the expected number; if the columns do match the value will be zero.
	- the third integer is the column number (zero based) in the row where the value is not a string.

The returned object will be like this:
`{data: dvalue, errors: evalue}`

The 2D array will be processed as follows:
- the number of columns in the first row is noted
- the type of each column must be a string; if not then an error message is generated containing:
	- the row number (zero based)
	- the column containing the data that is not a string
	- an explanatory message
	- the error message will be added to an array of error messages
- if the number of columns does not match the first row, then an error is generated containing:
	- the row number
	- the number of columns found and the number of columns expected
	- an explanatory message
	- the error message is added to the array of error messages
- the value in each column is altered per lossless conventions
- the columns are joined to make a string, with a tab character as the field delimiter and a newline character as the record separator.
- the string for the row is added to the output string
- after the 2D array is processed
	- if there were any conversion errors, an exception will be thrown
	- otherwise, the output string is returned

## Library Design

- There will be no React UI/X in this library.
- Styleguidist will be used to demonstrate each of the following:
	- String to 2D Array
		- a normal valid parsing case
		- a case having rows that are too long and too short
		- all cases will contain values requiring lossless conversion processing
	- 2D Array to String
		- a normal valid conversion
		- a case having non-string data
		- a case with rows too long and too short
		- all cases will contain values requiring lossless conversion processing

