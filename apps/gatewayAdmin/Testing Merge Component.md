

Note: I used migration process to create a "en_tn2" repo for the testing below.

Testing approach:
- for each case, do the update from master first, then the other way
- after all PRs done, do the update from master for each case, from the top
- lastly, do all the merges to master for each case, from the top

### branch-is-same

```
  const server = "qa.door43.org"
  const owner = "dcs-poc-org"
  // Can use: en_tn, en_tn_main_branch
  const repo = "en_tn2"
  /* Can use: branch-is-same, branch-behind, 
      branch-ahead, branch-behind-and-ahead, 
      branch-conflicts
  */
  const userBranch = "branch-is-same"
  // for single org use of the dcs-poc user
  const tokenid = "c8b93b7ccf7018eee9fec586733a532c5f858cdd" 
  // set to true when settings above are ready
  const inputsReady = true 
```

#### Check Merge of Default Branch into User Branch

PR is https://qa.door43.org/dcs-poc-org/en_tn2/pulls/1

```json
{
    "mergeNeeded": false,
    "conflict": false,
    "error": false,
    "message": ""
}
```

#### Check Merge of User Branch into Default Branch

```json
{
    "mergeNeeded": false,
    "conflict": false,
    "error": false,
    "message": ""
}
```

#### Merge Default Branch into User Branch

```json
{
    "mergeNeeded": false,
    "conflict": false,
    "success": true,
    "userBranchDeleted": false,
    "error": false,
    "message": "no merge needed"
}
```

#### Merge User Branch into Default Branch
```json
{
    "mergeNeeded": false,
    "conflict": false,
    "success": true,
    "userBranchDeleted": false,
    "error": false,
    "message": "no merge needed"
}
```

### branch-behind

```
  const server = "qa.door43.org"
  const owner = "dcs-poc-org"
  // Can use: en_tn, en_tn_main_branch
  const repo = "en_tn2"
  /* Can use: branch-is-same, branch-behind, 
      branch-ahead, branch-behind-and-ahead, 
      branch-conflicts
  */
  const userBranch = "branch-behind"
  // for single org use of the dcs-poc user
  const tokenid = "c8b93b7ccf7018eee9fec586733a532c5f858cdd" 
  // set to true when settings above are ready
  const inputsReady = true 
```

#### Check Merge of Default Branch into User Branch

PR is https://qa.door43.org/dcs-poc-org/en_tn2/pulls/2

```json
{
    "mergeNeeded": true,
    "conflict": false,
    "error": false,
    "message": ""
}
```

#### Check Merge of User Branch into Default Branch

```json
{
    "mergeNeeded": false,
    "conflict": false,
    "error": false,
    "message": ""
}
```

#### Merge Default Branch into User Branch

```json
{
    "mergeNeeded": false,
    "conflict": false,
    "success": true,
    "userBranchDeleted": false,
    "error": false,
    "message": ""
}
```

#### Merge User Branch into Default Branch

```json
{
    "mergeNeeded": true,
    "conflict": false,
    "success": false,
    "userBranchDeleted": false,
    "error": true,
    "message": "git commit [dcs-poc-org/en_tn2:branch-behind -> dcs-poc-org/en_tn2:master]: exit status 1\nOn branch base\nYou are in a sparse checkout with 0% of tracked files present.\n\nnothing to commit, working tree clean\n\n"
}
```

### branch-ahead

```
  const server = "qa.door43.org"
  const owner = "dcs-poc-org"
  // Can use: en_tn, en_tn_main_branch
  const repo = "en_tn2"
  /* Can use: branch-is-same, branch-behind, 
      branch-ahead, branch-behind-and-ahead, 
      branch-conflicts
  */
  const userBranch = "branch-ahead"
  // for single org use of the dcs-poc user
  const tokenid = "c8b93b7ccf7018eee9fec586733a532c5f858cdd" 
  // set to true when settings above are ready
  const inputsReady = true 
```

#### Check Merge of Default Branch into User Branch

PR is https://qa.door43.org/dcs-poc-org/en_tn2/pulls/3

```json
{
    "mergeNeeded": false,
    "conflict": false,
    "error": false,
    "message": ""
}
```

#### Check Merge of User Branch into Default Branch
```json
{
    "mergeNeeded": true,
    "conflict": false,
    "error": false,
    "message": ""
}
```

#### Merge Default Branch into User Branch

```json
{
    "mergeNeeded": false,
    "conflict": false,
    "success": true,
    "userBranchDeleted": false,
    "error": false,
    "message": "no merge needed"
}
```

#### Merge User Branch into Default Branch

```json
{
    "mergeNeeded": false,
    "conflict": false,
    "success": true,
    "userBranchDeleted": true,
    "error": false,
    "message": ""
}
```

NOTE! the merge to master succeeded and the PR was closed.

### branch-behind-and-ahead

```
  const server = "qa.door43.org"
  const owner = "dcs-poc-org"
  // Can use: en_tn, en_tn_main_branch
  const repo = "en_tn2"
  /* Can use: branch-is-same, branch-behind, 
      branch-ahead, branch-behind-and-ahead, 
      branch-conflicts
  */
  const userBranch = "branch-behind-and-ahead"
  // for single org use of the dcs-poc user
  const tokenid = "c8b93b7ccf7018eee9fec586733a532c5f858cdd" 
  // set to true when settings above are ready
  const inputsReady = true 
```

#### Check Merge of Default Branch into User Branch

PR is https://qa.door43.org/dcs-poc-org/en_tn2/pulls/4

```json
{
    "mergeNeeded": true,
    "conflict": false,
    "error": false,
    "message": ""
}
```

#### Check Merge of User Branch into Default Branch

```json
{
    "mergeNeeded": true,
    "conflict": false,
    "error": false,
    "message": ""
}
```

#### Merge Default Branch into User Branch

```json
{
    "mergeNeeded": false,
    "conflict": false,
    "success": true,
    "userBranchDeleted": false,
    "error": false,
    "message": ""
}
```

#### Merge User Branch into Default Branch

```json
{
    "mergeNeeded": false,
    "conflict": false,
    "success": true,
    "userBranchDeleted": true,
    "error": false,
    "message": ""
}
```

NOTE! the merge to master succeeded and the PR was closed.

### branch-conflicts

```
  const server = "qa.door43.org"
  const owner = "dcs-poc-org"
  // Can use: en_tn, en_tn_main_branch
  const repo = "en_tn2"
  /* Can use: branch-is-same, branch-behind, 
      branch-ahead, branch-behind-and-ahead, 
      branch-conflicts
  */
  const userBranch = "branch-conflicts"
  // for single org use of the dcs-poc user
  const tokenid = "c8b93b7ccf7018eee9fec586733a532c5f858cdd" 
  // set to true when settings above are ready
  const inputsReady = true 
```

#### Check Merge of Default Branch into User Branch

PR is https://qa.door43.org/dcs-poc-org/en_tn2/pulls/5

```json
{
    "mergeNeeded": false,
    "conflict": true,
    "error": false,
    "message": ""
}
```

#### Check Merge of User Branch into Default Branch

```json
{
    "mergeNeeded": false,
    "conflict": true,
    "error": false,
    "message": ""
}
```

#### Merge Default Branch into User Branch

```json
{
    "mergeNeeded": false,
    "conflict": true,
    "success": false,
    "userBranchDeleted": false,
    "error": false,
    "message": ""
}
```

#### Merge User Branch into Default Branch

```json
{
    "mergeNeeded": false,
    "conflict": true,
    "success": false,
    "userBranchDeleted": false,
    "error": false,
    "message": ""
}
```
