# Links

Here are relevant links:

- develop: https://develop--gateway-translate.netlify.app/
- "staging": https://develop--gateway-translate.netlify.app/?server=prod
- prod: https://gateway-translate.netlify.app/

Here are links to the underlying github branches:
- prod: https://github.com/unfoldingWord/gateway-translate/tree/main
- develop: https://github.com/unfoldingWord/gateway-translate (now the default branch)
- release branch for v0.9.0: https://github.com/unfoldingWord/gateway-translate/tree/release-v0.9.0
- release: https://github.com/unfoldingWord/gateway-translate/releases/tag/v0.9.0

## Steps to Release to Production Use

**Release Processes**

Note 1: Prerequisites and one time activities are included at the end of this post.

Note 2: Pull Requests (PRs) are the mechanism used to do merges and they should be reviewed. In the below, for the sake of brevity, PRs are not mentioned.

**Preconditions**

a. The develop branch has all changes merged that are supposed to be in the next release
b. QA has signed off on the work
c. Release notes are prepared

**Tasks**
 
1. create a release branch from develop, named by the release semver
2. merge release branch into main
3. tag develop: `git tag <semver> && git push --tags`
4. cut a release from the release branch
5. add to the release the notes and any assets
6. verify Netlify links for main (production) and develop are working

**Expected Outcome**

1. At this point, in time the main, develop, and release branches will be identical. In addition, the release itself will be the same code.
2. There will be two netlify links, one for develop and one for main. The code in both will be identical (after netlify produces the builds and deploys them)

**Process for Bug Fixes**

If a bug is found in production,then:
  - branch from the release branch a patch release
  - make, test changes
  - make release notes
  - merge into main
  - tag it
  - cut a patch release from main, named by the release semver
  - add to the release the notes and any assets

To incorporate bug fix into develop, do:
  - branch from develop
  - merge in the patch branch
  - resolve any conflicts and re-test
  - merge into develop


*Prerequisites and One-Time Tasks*

1. In addition to the "main" branch, create a "develop" branch.
2. Set the "develop" branch as the "default" branch.
3. All new work should be branched from "default"
4. In repo settings, setup branch protections for both branches
5. Configure netlify to monitor the repo and automate the builds
  - preview builds for user branches
  - a develop build (https://develop--gateway-admin.netlify.app/)
  - a production build (https://gateway-admin.netlify.app/) 