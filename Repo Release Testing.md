# Repo Release Testing

**NOTE** having trouble with fa_tq and suspect it is due to it being a fork of en_tq. The same general test process should work with fa_obs-twl, which does not have any releases yet either.

Repo: https://qa.door43.org/fa_gl/fa_obs-twl
(org=fa_gl and lang=fa)

*Current state*: no releases have been done

### Pre-Release 
Start gatewayAdmin and use app drawer to go to the Account Settings page.

Select org `fa_gl` and language `fa - Persian (Farsi)`

Click "Save and Continue" button.

Use app drawer to go to the "Release Repository" page.

Select "OBS TWL" as the resource to release.

Select "Pre-Release" as the Release Type.

Enter "v1" (as recommended on page) for the version.

Enter "Version 1" for the Release Name.

Copy and pasted the below into the Release Notes.
```
*This is bolded text. Some links [https://www.unfoldingword.org/utq/](https://www.unfoldingword.org/utq/).*

The following books have undergone a Book Package consistency check:

-   Exodus (EXO)
-   Ruth (RUT)
-   Ezra (EZR)
-   Nehemiah (NEH)
```

After a few seconds, if all is well, this status message will be shown on the page just above the footer (may have to scroll down):
```
Status: Created release v1 of fa_obs-twl
```
Finally, click the Close button to return to the home page (book package page).

Verify using this [link(https://qa.door43.org/fa_gl/fa_obs-twl/releases)] to go to the release page for this repo. It will now have a "Version 1" release. Also look for the following:
- Your identity as the one who made the release
- The complete release notes
- Tag of "v1"
- Catalog lable reading "Catalog (prod)".

### Production

Repeat the steps above, except select "Production" as the Release Type.

Note: to delete a release, click the "edit" link in the Name (Title) area. At the bottom of the screen will be a button to delete the release.
