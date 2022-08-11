# Tauri

As of 2022-07-22, this study is being done on Linux (Debian x64 on my Chromebook).

## General Notes 
1. Supports Netlify as a hosting service (not sure why it needs one? perhaps there is a server portion??)

## Steps
1. Had to install rust 
2. Using the quickstart at  https://tauri.app/v1/guides/getting-started/
3. It says that it supports any web framework as a front-end
4. So I am trying the react/proskomma app at https://github.com/mandolyte/bible-ref-pk-demo
5. I cloned it; then used 'yarn install' to retrieve all the dependencies
6. This is in `$HOME/Projects/github.com/mandolyte/bible-ref-pk-demo`
7. Installed the tauri cli: yarn add -D @tauri-apps/cli
8. Created the tauri app scaffold using `yarn tauri init`
9. The preceding is done in your web app folder (here `bible-ref-pk-demo`)
10. Transcript:
```
$ yarn tauri init
yarn run v1.22.15
$ /home/mando/Projects/github.com/mandolyte/bible-ref-pk-demo/node_modules/.bin/tauri init
✔ What is your app name? · bible-ref-pk-demo
✔ What should the window title be? · bible-ref-pk-demo
? Where are your web assets (HTML/CSS/JS) located, relative to the "<current dir>✔ Where are your web assets (HTML/CSS/JS) located, relative to the "<current dir>/src-tauri/tauri.conf.json" file that will be created? · ../build
? What is the url of your dev server? (http://localhost:3000) › http://localhost:✔ What is the url of your dev server? · http://localhost:3333
Done in 63.41s.
$ 
```
11. I then had to start the app server via the usual `yarn start`
12. Next I ran `yarn tauri dev`; this looked for my running server and then started downloading lots of stuff:
![[Pasted image 20220722085954.png]]
13. Yes... that 383 items being used in the build...
14. Took 11m:
![[Pasted image 20220722090758.png]]
15. It then opened  a local window that was blank... nothing worked, altho moving the pointer around the appearance changed to indicate it was located over something clickable. When I did click, I noticed that code in vs-code window would appear. So perhaps some debugging things happening.
16. I closed the window and re-ran the command and again a blank window appeared. This time it took a little over one minute to build.

Later after some simpler tests, I came back to this thinking I'd try to resize the window to see if it would render. *This time the page did render!*

Note: the "Print" button did not work. It failed when I tried to open a new page from within the app.
