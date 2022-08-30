# Manifest Version Attribute

In this post, I argue that a manifest should not have a version since it will be incorrect and misleading.

The manifest version attribute refers to an entire resource repository. Thus:
- In the master branch, if the version refers to the latest released version, then it is wrong since:
	- the master branch will have files not in the latest released branch (since we will be implenting book-package-releases, only the book packages selected will be released).
	- the master branch is work-in-progress, so soon it will have content that differs from the released branch, yet it has the same version
- Since the release branch accumulates content as book packages are released, then if the release branch manifest has a version attribute, it will be misleading since it will refer to content that was released previously. I grant that this is much less of a problem in understanding, so if the manifest in the release branch has a version, it is less serious.