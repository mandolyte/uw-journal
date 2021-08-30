# Admin App

This is a Proof of Concept. It will be in unfoldingWord-box3 for now. I began with a fork of `create-app` and then renamed it.

Issue: https://github.com/unfoldingWord/tc-create-app/issues/713

See [[Design]].
Requirements from above (roughly reformatted).

1. Admin does a health check of files prior to assigning translation work.  
2. Log into Organization
	- initialization and 
	- then run validation check.
3. If repo exists with correct naming convention, then check complete. No error.  
4. If repo missing or incorrectly named, then error symbol on resource. 
	1. Admin to have option to 
	2. 1. Create a new Repo 
	3. 2. Select existing repo  
5. If ‘Select existing repo’ selected, Admin to be displayed all repos for the particular resource type within the organization.  
6. If Admin selects repo but existing repo name incorrect, then repo to be renamed.  
7. If a manifest file is missing in repo, an appropriate error message will be displayed.

Restatement of Requirements:

**Purpose**: Admin does a health check of files prior to assigning translation work.

**Initialization Steps**:
1. Log in
2. Select Organization
3. Select Language
4. Select Resource Type 
	- Alternate 1: include an "all resource types" selection
	- Alternate 2: just automatically check all resource types

**Repo Check Steps** (based on what is selected above):
1. Does repo exist?
2. If yes, then proceed to manifest check steps. Otherwise continue.
3. Offer two options: 
	- Create the repo
	- Rename an existing repo to conform to standard naming
4. If 'create', then create it; then done
5. If 'rename' then list all repos in org for given resource type and let admin select one to rename.

**Manifest Check Steps**
1. Does manifest exist in repo?
2. If yes, then done.
3. Show appropriate error message.

**Format check:** (added by @birch via [here](https://docs.google.com/document/d/1crQ3CoRBo0gySGW8IBQ4x_NyZriUTl6-UX_mLAR40E8/edit))
1. Check to see if the format is in TSV
2. Prove that we can launch a process to convert it.


## 2021-02-15

Step 1: after forking and renaming, get it running as "create-app" (then I'll start stripping it down)

- `yarn install`
- `yarn dev`

This worked OK... Now to start stripping it back to just login and organization/language selection.

