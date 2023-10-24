# Right to Left Languages

## Links

Font playground:
https://simple-usfm-alignment-prototype.netlify.app/

Document deliverable:
https://docs.google.com/document/d/1aOiyL-98j-ldFG4sTAvvEWJEkSiV24d0WAmxYl7VKU8/edit#

Font detection: https://font-detect-rhl.netlify.app/

Graphite fonts: https://software.sil.org/fonts

How to build RTL apps:
- [https://hacks.mozilla.org/2015/09/building-rtl-aware-web-apps-and-websites-part-1/](https://hacks.mozilla.org/2015/09/building-rtl-aware-web-apps-and-websites-part-1/) 
- [https://hacks.mozilla.org/2015/10/building-rtl-aware-web-apps-websites-part-2/](https://hacks.mozilla.org/2015/10/building-rtl-aware-web-apps-websites-part-2/)

## To Do

1. Contact Alex Agha-Khan and ask these interview questions in Appendix A
2. Determine acceptable lag time for RTL/Complex script character to show after typing.



# Appendix A - Alex

### Pre-interview question

-   What RtLs can you read? Farsi, Arabic, Urdu, Dari
-   What RtLs can you type? Same

Classical Arabic is like calligrophy. They go into the above line and into the line below.

Farsi has no diacritics.

Arabic words are written on top of each other. Both vowels and consanants use diacritics.

Cursor needs to be on the right. Not left.

Does the browser support the language, keyboard, and font and size.

"aap" - three letter sequence Urdo

Question: is there a keyboard viewer

Urdo has a "beloved" font. 

Farsi is simplistic.


### Interview questions

Set up a tC Create environment in Firefox with the appropriate RtL files for the tester

-   Does it read correctly? Is the font legible (spacing, correct diacritics)?
-   Does wrapping work correctly?
-   Try typing and deleting. What issues do you see?
-   Try scrolling. Does it function correctly? If not, what is wrong?
-   Try pasting. What issues do you see?
-   Try undo. What issues do you see? 
-   Do the source and target need to be swapped?
-   What are the most annoying issues with RtL editors?
    

List discovered issues in order of priority (need to have vs nice to have)**