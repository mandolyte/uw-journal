# tc-create-app#1075
Issue [here](https://github.com/unfoldingWord/tc-create-app/issues/1075)

NOTE: put relevant material here into the issue!!

At present the issue has a number of links to articles on infinite scrolling, pros and cons, etc. See appendix A.

I also have a suggestion that JoelC do a proof of concept (see appendix B).

Wondered if an unbounded dataset was possible... (see appendix C).

## 2021-10-26

To answer questions raised yesterday, namely:
1. Will this work with a MUI datatable?
2. Will this work with pagination turned on?

Let's try to use the Styleguidist demo for datatable-translatable to find out...


## 2021-10-25

Articles on a JS function named "scroll to view":
- https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollIntoView
- https://stackoverflow.com/questions/5007530/how-do-i-scroll-to-an-element-using-javascript
- https://www.w3schools.com/jsref/met_element_scrollintoview.asp
- Example: https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_element_scrollintoview

Questions:
1. Will this work with a MUI datatable?
2. Will this work with pagination turned on?

*Links to Table Components:*
- https://flatlogic.com/blog/top-19-remarkable-javascript-table-libraries-and-plugins/#five
- https://bvaughn.github.io/react-virtualized/#/components/Grid
- https://wpdatatables.com/javascript-data-table/
- https://griddlegriddle.github.io/Griddle/
- https://react-table.tanstack.com/ (recommended by Manny)






# Appendix A - on infinite scrolling

Here is a MUI supported way to do infinite scrolling:
https://mui.com/components/data-grid/rows/#infinite-loading

@Joel-C-Johnson here is some reading to do as part of the spike:
- https://uxplanet.org/ux-infinite-scrolling-vs-pagination-1030d29376f1
- https://crocoblock.com/blog/pagination-vs-infinite-scroll/
- https://www.uxmatters.com/mt/archives/2018/11/paging-scrolling-and-infinite-scroll.php
- https://www.knowband.com/blog/ecommerce-blog/pagination-vs-infinite-scrolling/
- https://uxdesign.cc/ui-cheat-sheet-pagination-infinite-scroll-and-the-load-more-button-e5c452e279a8
- https://webo.digital/blog/pagination-vs-infinite-scroll-which-is-better-for-your-content/
- https://www.oneupweb.com/blog/infinite-scroll-pagination-better/


# Appendix B - PoC on Infinite Scrolling
@Joel-C-Johnson let's chat about doing a PoC. What I'm thinking of is something like this:
- A main workspace containing two infinite scrolling windows, divided horizontally, top and bottom
- The top window is a scripture window with divided into 3 or more columns
- The bottom window is the TSV (say Translation Notes)

Both using the [MUI data-grid component](https://mui.com/components/data-grid/rows/#infinite-loading); and both synced on scroll and some provision to sync on demand (in the case where they scroll manually to some location and need to "re-center", if you will)

# Appendix C - unbounded datasets

The goals would be to minimize overhead of rendered data and make the amount of data available essentially unbounded. This means the DOM would be constantly updated as it scrolled, rendering the rows on the fly; enough to fill the display space plus perhaps one row overflow at top and bottom (outside the view port).

- First, there is a data source defined by an interface. So that whether the data is actually stored in-memory, in indexedDB, proskomma, then a function, say getRow(n), will get the nth row of data.
- Second, the UI has "callbacks" that will fire when a row scrolls out of view or into view. The callbacks would reconstruct the viewport on they fly. In other words, the viewport and the rendered data are nearly the same.
- Third, no pagination... the user can scroll wherever they wish (and then sync back up with the below feature)
- Fourth, *must* support the "scroll to view" function so that multiple tables can be synced with each other
