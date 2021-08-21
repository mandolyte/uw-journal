# DCS Workflow

What if:
- The user branch is named with the following elements:
    - user name
    - filename
    - app name
- A PR is created, but set to draft mode as soon as the branch is created

Then the benefits:
- We would be able to easily tell if someone is already working on the file by looking at the branch names that have a draft PR associated
- We would know who else is working on the same file
- We would know if the file has changed since we started working on it by checking on any PRs closed/merged for the file since I started making my changes. (assumes that some timestamps are associated with PRs; if not then timestanp could be an addition element of the branch name)
