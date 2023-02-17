https://sqlite.org/forum/forumpost/728ff91b6c5733fe64c4d001304c83fb8d9a140c85a6d5396bf05f9c3c92be10

```
Btw, setting Next.js is a piece of cake.

Install Node.js

npx create-next-app@latest
Replace pages/index with

import { useEffect } from "react";

export default function Home() {
  // This runs code only in the browser.
  useEffect(() => {
    const dbWorker = new Worker(new URL("../lib/dbWorker", import.meta.url));
  }, []);
  return null;
}
// lib/dbWorker.js, 
import sqlite3 from "./sqlite3.mjs";

sqlite3().then((sqlite3: any) => {
  const capi = sqlite3.capi /*C-style API*/,
    oo = sqlite3.oo1; /*high-level OO API*/
  // console.log(
  //   "sqlite3 version",
  //   capi.sqlite3_libversion(),
  //   capi.sqlite3_sourceid()
  // );
});
Then npm run dev, and that's all.
```
Example: https://github.com/evoluhq/evolu

const db = new oo.DB("/mydb.sqlite3",'ct');


sqlite3.capi.sqlite3_js_vfs_create_file("opfs", "my-db.db", arrayBuffer);
const db = new sqlite3.oo1.OpfsDb("my-db.db");


## SQL Viewer Examples

https://sqliteviewer.app/
You can drag and drop a db file and it will work with it.

https://inloop.github.io/sqlite-viewer/
Similar features

https://github.com/nalgeon/sqlime
Uses plain js

https://sqlite.org/fiddle/index.html
First attempt to get wasm working

This looks nice, but I think it only works with client node apps:
https://github.com/theastroscout/SQLite

This project is active and is actually underneath some wrappers:
https://www.npmjs.com/package/better-sqlite3

Google has a demo:
https://github.com/GoogleChrome/developer.chrome.com/blob/main/site/en/blog/sqlite-wasm-in-the-browser-backed-by-the-origin-private-file-system/index.md
... with demo here: https://sqlite-wasm-opfs.glitch.me/
and this case, you can see the files in the OPFS explorer!

This one looks interested in my goals:
https://github.com/overtone-app/sqlite-wasm-esm

