# Releasing Book Packages

## Assumptions

Based on response to [[On Versioning]] in [Zulip](https://unfoldingword.zulipchat.com/#narrow/stream/207526-Tools---UR/topic/On.20Versioning/near/295839989), the specs below specify use of simple incrementing integers for repo release versions.

## Process for book resources

**Notes:**

1. This does *not* include Translation Notes or Translation Academy. See next section for those two resources.
2. This does *not* include the OBS or its resources.
3. Thus by elimination, the process steps below *only* apply to the following resources:
	1. Literal Translation (LT)
	2. Simplified Translation (ST)
	3. Translation Notes (TN), either 9col or 7col
	4. Translation Questions (TQ), only TSV format, not markdown
	5. Translation Word Lists (TWL)
	6. Study Questions (SQ)
	7. Study Notes (SN)

**Steps:**

1. Get the latest release version. If none, set to zero.
2. Increment the version by one. 
3. Create a branch from master with new version number as the name. For example, v40. 
4. Remove book files not being released by the admin in the v40 version. 
5. Copy in all files from the prior "v39" release that are not being released by the admin in the v40 version.
6. Update the manifest "projects" section to contain on the files in the v40 version. 
7. Finally, cut "v40" release from the "v40" branch.


## Process for OBS

**Notes:**

This process applies to the OBS itself, for example, `en_obs` and its resources, which are:
1. Translation Notes, for example: `en_obs-tn`
2. Translation Questions, for example: `en_obs-tq`
3. Translation Word Lists, for example: `en_obs-twl`
4. Study Notes, for example: `en_obs-sn`
5. Study Questions, for example: `en_obs-sq`

This process does *not* apply to the Translation Words and Translation Academy resources which are shared with Scripture resources. In other words, the OBS TN uses the same TA repository as does Scripture TN; and the OBS TWL uses the same TW repository as does Scripture TWL.

The OBS and the OBS resources are all "single" respositories. In other words, when any of the repositories are released, there is no need to trim out any files. Specifially:
- The OBS text is in a single folder
- The TN is a single file named "tn_OBS.tsv"
- The TQ is a single file named "tq_OBS.tsv"
- The TWL is a single file named "twl_OBS.tsv"
- The SN is a single file named "sn_OBS.tsv"
- The SQ is a single file named "sq_OBS.tsv"

Similarly, there is no need to adjust the manifest for any of these resources.

**Steps:**

1. Get the latest release version. If none, set to zero.
2. Increment the version by one.
3. Cut release *directly* from the master branch.

## Process for TW and TA

**Notes:**

For the reasons listed below, it is proposed to treat TW and TA in a manner similary to OBS resources, namely, cut the release from the master branch without any adjustments.

1. Articles may be released prior to their incorporation into the TWL and TN files, respectively.
2. Conversely, the TN and TWL may be released before the articles are translated. So the connection between the resources is fairly loose.
3. There may be OBS articles that are not used in Scripture and the reverse is true as well.

**Steps:**

1. Get the latest release version. If none, set to zero.
2. Increment the version by one.
3. Cut release *directly* from the master branch.
