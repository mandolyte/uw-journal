# Translation Database

From Zulip:
Jesse Griffin: @Cecil New looks like you have an account already, the django siteadmin page is [https://td.unfoldingword.org/siteadmin/](https://td.unfoldingword.org/siteadmin/)

Jesse Griffin: However, I think you should be able to make this change in [https://td.unfoldingword.org/uw/languages/3013/edit/](https://td.unfoldingword.org/uw/languages/3013/edit/) @Cecil New

---
With the above, I was able to change the lang code for Kannada from the three-character code to the two-character code already being used. `kan` to `kn`.

Note: my account name is `ceciln`. Password is in bitwarden. The account is for `unfoldingword.org`:
![[Pasted image 20210721094027.png]]

SIL language data source:
https://github.com/silnrsi/langtags/blob/master/doc/langtags.md



First, you must have access to django site administration at: https://td.unfoldingword.org/siteadmin/login/?next=/siteadmin/
I suspect this is also an account to a postgres database that is underneath this
And I think the source for this data (at least originally) was https://github.com/silnrsi/langtags/blob/master/doc/langtags.md
Ok, maybe you don't need all that... Let's try this:
1. Go to https://td.unfoldingword.org/uw/languages/
2. Find the language
3. Click the language code (it should be a URL, but maybe only if you have access rights?)
4. In the address bar, alter the link by adding "/edit" at the end
5. Go there to that new URL
6. The edit form is presented.
7. Make changes and click Save 
And yes, you must have an account. If you go thru steps, when you do the edit modification, then a login is presented if you have not previously authenticated.

![[Pasted image 20221220092941.png]]

