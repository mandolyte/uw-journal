# Issue 76
Link: https://github.com/unfoldingWord/tc-create-app/issues/76

## Issue Description
Users want to see the exact options that are available for the RC links needed for supportreferences in the TNs. Users should be able to type ahead any of the Support References that are already in the document. Users need to be able to select from a list, enter custom text or a blank field.

DoD: Users can pick a TA article from a type-ahead drop-down for TNs.

## Technical Description

1. The dropdown feature will only be available on fields designated as "filterable".
2. The choices shown in dropdown will be those already in the file.
3. The user must be able to enter a new value that isn't shown

NOTE: In some cases the total list of allowed values is actually known. And the approach described above does not account showing the complete list and preventing something new, and possibly invalid, to be entered. See Appendix A for a different approach that would work this for those columns that have a finite known list of valid values.

# Appendix A 
**Assumption**
It is assumed that a GL may not create a language specific Translation Academy articles. If new articles are permitted, then the below would need to be adjusted slightly.

Consider the support refence column, which has 177 valid values as of this writing. Here is sketch of how tc-create could restrict the choices to these values.

- Obtain the list of valid values at run time.
- When a tN resource is being edited, pass an optional parameter to the instance of datatable-translatable.
- This parameter would be an object where:
	- The key is the column header value, namely, `SupportReference`
	- The value would be an array of the valid values
- When the user attempts to enter a value for support reference, the component would use the autocomplete feature of the dropdown to permit easy selection of the desired value.

