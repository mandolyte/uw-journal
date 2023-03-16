# Login

The login is for DCS and you may point it production with:
https://gateway-translate.netlify.app/ or 
https://gateway-translate.netlify.app/?server=prod (if you are switching back from QA)

And to QA with: https://gateway-translate.netlify.app/?server=qa

---

We have considered a "guest" login so that gatewayTranslate (gT) could be used in "read-only" (but see more on how read-only files are handled... *such files can still be saved locally!*)

---

# Open Books

- open Titus ULT and UST
- open Titus from es-419 GLT
- open by URL: [here](https://raw.githubusercontent.com/Proskomma/perfidy/main/dataSources/lsg_tit.usfm)
- open by upload: `TIT.usfm`

---
# Notes

- By upload and URL are read-only by default
- Cannot start from scratch yet

---

# Meanwhile, back on DCS

- Branch naming is under discussion, but at present they look like this: `gt-TIT-cecil.new`
- Pros/Cons

---

# Other Features

- Broken alignment indicator
- Read-only setting
- Data loss prevention
	- On closing a text
	- On closing the browser tab

---

# Plans for v1.0

- Book syncing 
- Search and replace (within a book)
- "Global" search and replace
- Alignment editor
- Preview
