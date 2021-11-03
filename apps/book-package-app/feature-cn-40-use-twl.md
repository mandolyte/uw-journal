# Issue 40

This issue is to switch to locating the Translation Word articles, 
- *from* parsing the USFM
- *to* the TWL files

This change is anticipated to be mostly, if not solely, in the rcl code.

In the RCL, this obtains the data from the USFM:
```js
        if ( bookId === 'obs' ) {
          result = await fetchObsTw({ bookId: bookId,  clearFlag: clearFlag });
        } else {
          result = await fetchBookPackageTw(
            {username: 'unfoldingword', languageId:'en', 
            bookId: bookId, chapters: chapter, clearFlag: clearFlag
          });

```

I also noted that for OBS, it already uses a TWL. So I'll be able to use it for Bible books.

## 2021-07-14
Steps:
1. cloned both app and rcl and created a vs code workspace for the two of them
2. In rcl, did yarn install, and yarn start.
3. Confirmed that tW for bible is not working:
![[Pasted image 20210714083604.png]]
4. Also noticed that tw for OBS is working even tho it is pointing the wrong repo now. The repo, even tho archived, must still be accessible.
![[Pasted image 20210714083731.png]]

First, let's correct the OBS link and see if we get a count different from 58,487. It is!
![[Pasted image 20210714084550.png]]

Next up... the bible books!
- main function in the helpers file is `fetchBookPackageTw`.
- currently, it fetches the USFM and parses out the t-tw 

Copied the OBS code and adjusted for bible books.

Now Titus is reporting 33,341 (which is correct... confirmed with Russ Perry).

- Chapter 1 has 21,400
- Chapter 2 has 14,050
- Chapter 3 has 15,713

Cannot just sum these since they will contain dups. But if I check chapters 1,2,3:
![[Pasted image 20210714092535.png]]

Boom! done.

Now remove the old code and deploy.