# Save Function Design

## Scope

This design is intended to cover *only* the function that stores data back to DCS in a user branch. It does not cover any UI/X look and feel. It will impliment using the `dcs-js` package.

## Context

During the edit of a Scripture text, the user may wish to save current work at any time. The app will respond to a save request by the user by doing the following:
- Inputs: 
	- Flag for Save or "Auto-Save" (former requires lengthy conversion from PERF to USFM; whereas the latter is intended to save the PERF without conversion)
	- Content to be saved
	- BookId to be saved
	- DCS user name
- Actions:
	- If save, convert PERF to USFM (unless Epitelete does this)
	- Test for existence of user branch, creating if it doesn't exist
	- Save content to user branch
- Outputs:
	- Console log showing save details
	- Boolean return indicating success (or failure) 
	- Message with failure details (to be used in an error popup)