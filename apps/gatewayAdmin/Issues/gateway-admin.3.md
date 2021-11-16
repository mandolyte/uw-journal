# Issue 3

Link: https://github.com/unfoldingWord/gateway-admin/issues/3


## 2021-11-11

Things to clean up from yesterday:
- Consider refactor to move the TN useEffect into its own custom hook. Named like this `useTnRepoValidation` (as opposed to content validation)
- Implement a nice table to list the resources. Something like this:

| Resource   | Repo Name | Status |
| ----------- | ------------ | ------- |
| Translation Notes | vi_tn | No manifest |
| Translation Word List | vi_twl | Repo does not exist |



## 2021-11-10

Today, let's create the "admin context". 
- Use GRT's File.context.js as a starter
- Worked fairly well. Note: branch is is `feature-cn-3-add-repo-validation`.



## 2021-11-09

Today, lets focus on the following:

- Use zip copy of template and get it working as gA.
- Layout a card for each book of the Bible


Actual work...

- Created `src/components/RepoValidation.js` from `WorkspaceContainer.js`. Made some mods to iterate over all books of the bible...
- Wasted a lot of time in trying to figure out how to use the workspace rcl.

## Design Notes 

*These are also in the issue...*

Create a React Context and custom hooks, which does the following:
1. Since repo validation is common to any book/OBS, then it should only be performed one time and the results updated as needed (ie, when org or language is changed).
2. The results of the validation will be stored in the Context and be available to the UI to be presented to the user.
3. There are two possible variations on how to do the hook(s).
   - The first variation is to have one hook with org, language, and resource type as dependencies.
   - The second variation is to have one hook per resource type, taking only org and language as dependencies.

For the moment, let's call the context `RepoValidationContext`.

The values for org and language will be managed as state variables (unless they themselves are in their own Context already -TBD). As they are changed by the  user, the hook will execute doing the "light" repo validation and storing the results in the context.

Once the results are updated/changed, then the UI will update.

