**Release Steps** *for gatewayTranslate v0.9.0*

**Preconditions**

- [] The develop branch has all changes merged for the next release
- [] The develop branch has the semver for the next release
- [] QA has signed off on the work (including tested against production)
- [] Release notes are prepared

**Tasks**

- [] create a release branch from develop, named by the release semver: `release-v0.9.0`
- [] create PR to merge release branch into main and merge
- [] tag develop: git tag *v0.9.0* && git push --tags
- [] cut a release from the release branch
- [] add to the release the notes and any assets
- [] announce on forum.door43.org

**Expected Outcomes**

- At this point, in time the main, develop, and release branches will be identical. 
- In addition, the release itself will be the same code plus the notes and assets.
- There will be two netlify links, one for develop and one for main. The code in both will be identical (after netlify produces the builds and deploys them)
