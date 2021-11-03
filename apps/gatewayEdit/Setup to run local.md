# Setup to run local

First, cloned the repo. Then used `yarn install` to fetch dependencies. Got this error:
```sh
error tailwindcss@2.0.3: The engine "node" is incompatible with this module. Expected version ">=12.13.0". Got "10.16.3"
error Found incompatible module.
info Visit https://yarnpkg.com/en/docs/cli/install for documentation about this command.

mando@DESKTOP-0V8P6MM MINGW64 ~/Projects/unfoldingWord/create-app (main)
$ node --version
v10.16.3
```

Looks like I need to upgrade the version of node I am using.

Site: https://nodejs.org/en/download/

Now have:
```sh
$ node --version
v14.16.0
```


```sh
$ yarn start
yarn run v1.22.5
$ next start
Error: Could not find a production build in the 'C:\Users\mando\Projects\unfoldingWord\create-app\.next' directory. Try building your app with 'next build' before starting the production server.
```

Appears that the `next` command must run as a sub-command of npx or yarn. Thus:
```sh
$ yarn next --version
yarn run v1.22.5
$ C:\Users\mando\Projects\unfoldingWord\create-app\node_modules\.bin\next --version
Next.js v10.0.7
Done in 0.42s.
```

Must use `yarn dev` to get the hot re-load feature. Here is sample output from it:
```sh
$  yarn dev
yarn run v1.22.5
$ next
ready - started server on 0.0.0.0:3000, url: http://localhost:3000
event - compiled successfully
event - build page: /
wait  - compiling...
event - build page: /
```

Seeing these messages in terminal:
```sh
useLocalStorage(owner) - init error:' ReferenceError: localStorage is not defined
    at C:\Users\mando\Projects\unfoldingWord\create-app\.next\server\pages\_app.js:625:20
```

