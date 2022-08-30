# On Versioning

Consider these two links:
https://www.unfoldingword.org/for-translators/content#VERSIONING and 
https://door43.org/u/unfoldingWord/en_ta/master/03-translate.html#translate-source-version

I see two practices at present:
- Some are applying the above versioning practices to an entire repository
- Others are using simple incrementing integers for repositories

In this post, I'll make the argument that the use of simple incrementing integers is the correct way to version repositories and that the more complex versioning practice is only suitable for individually translated texts.

Consider this scenario:
- About nine months ago a GL team started translating Genesis from the ULT using release v32 (current at the time).
- Last week, a GL team member started on Obidiah using release v40, which was released 2 weeks ago.
- Yesterday, Obidiah GLT was completed and released and the GLT repo version was v40.1.
- Today, Genesis was completed and released and the GLT version would be v32.1

But this means the versions are going backwards, which is nonsensical.

Therefore, the documented version practice can only apply to individual texts, not entire repositories. Consequently:
- The language should be clarified to refer to individual texts (I think it is implied already in the wording, but clearly my interpretation has not been followed by some GL teams).
- Since Tools is working releasing book packages in gatewayAdmin, we need permission to standardize on how this works, namely, to use simple incrementing integers.

