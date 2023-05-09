# Save Function Design

## Scope

This design is intended to cover *only* the function that stores data back to DCS in a user branch. It does not cover any UI/X look and feel. It will impliment using the `dcs-js` package.

## Context

During the edit of a Scripture text, the user may wish to save current work at any time. The app will respond to a save request by the user by doing the following:
- Inputs: 
	- Type of content: ['usfm' , 'perf']
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


## Functions to Code

### existsUserBranch

This function will test to see if a user branch exists. We need a naming convention for these branches since they will be created by the app if they don't exist.

So how about `gt-{bookId}-{username}`?

Parameters: bookId and username.

Returns: true or false. 

Exception: on any system failures.

Code: https://github.com/unfoldingWord/dcs-js/blob/master/apis/repository-api.ts#L18344 or https://github.com/unfoldingWord/dcs-js/blob/master/apis/repository-api.ts#L11825

### createUserBranch

This function will create the user branch above if it does not exist.

Parameters: bookId and username

Returns: true (if created ok) or false (if create fails; also throws exception)

Exception: on any create failure

Code: https://github.com/unfoldingWord/dcs-js/blob/master/apis/repository-api.ts#L1567

### saveUserBranch

This function saves the supplied content to the user branch.

Parameters: 
- bookId
- username
- content type ("perf" or "usfm")
- content

Returns: true or false (on failure to save)

Exception: on failure to save an exception will be thrown

Code: https://github.com/unfoldingWord/dcs-js/blob/master/apis/repository-api.ts#L18759
https://github.com/unfoldingWord/dcs-js/blob/master/documentation/classes/RepositoryApi.md#repoupdatefile

**repoUpdateFile**(`owner`, `repo`, `filepath`, `body`, `options?`): `Promise`<`AxiosResponse`<[`FileResponse`](https://github.com/unfoldingWord/dcs-js/blob/master/documentation/interfaces/FileResponse.md)>>

