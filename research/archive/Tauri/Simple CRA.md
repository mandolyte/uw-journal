# Simple CRA

These commands (from ): https://reactjs.org/docs/create-a-new-react-app.html
```
npx create-react-app my-app
cd my-app
npm start
```

Then:
```
yarn add -D @tauri-apps/cli
$ yarn tauri init
yarn run v1.22.15
$ /home/mando/Projects/github.com/mandolyte/my-app/node_modules/.bin/tauri init
✔ What is your app name? · my-app
✔ What should the window title be? · my-app
? Where are your web assets (HTML/CSS/JS) located, relative to the "<current dir>✔ Where are your web assets (HTML/CSS/JS) located, relative to the "<current dir>/src-tauri/tauri.conf.json" file that will be created? · ../build
? What is the url of your dev server? (http://localhost:3000) › http://localhost:✔ What is the url of your dev server? · http://localhost:3000
Done in 6.37s.
$ 
```

And then:
```
yarn tauri dev
Compiling atk v0.15.1
   Compiling javascriptcore-rs v0.16.0
   Compiling tauri-codegen v1.0.4
   Compiling tauri-build v1.0.4
   Compiling gdk-pixbuf v0.15.11
   Compiling soup2 v0.2.1
   Compiling tauri-macros v1.0.4
   Compiling gdk v0.15.4
   Compiling app v0.1.0 (/home/mando/Projects/github.com/mandolyte/my-app/src-tauri)
   Compiling webkit2gtk v0.18.0
   Compiling tauri-runtime v0.10.2
   Compiling tauri-runtime-wry v0.10.2
    Finished dev [unoptimized + debuginfo] target(s) in 10m 01s
```
After 10m, the expected window opens and it looks good.

