# Runtime detection: Development vs Production


1. In the project root folder, include a file named `.env`
2. This file has environment variables that are automatically added to the run time instance of the code.
3. tc-create 's content for `.env` is as follows:
```sh
GENERATE_SOURCEMAP=false
REACT_APP_DOOR43_SERVER_URL="https://qa.door43.org"
CYPRESS_CACHE_FOLDER="./node_modules/CypressBinary"
```
4. This means that by default it will target QA as the DCS instance to use.
5. This also means that when netlify deploys the code it will overwrite this file with suitable values for the expected target DCS instance.
6. The server URL is read into the code via this in `state.defaults.js`: 
```js
export const SERVER_URL = process.env.REACT_APP_DOOR43_SERVER_URL;
```
7. This makes the server (aka base_url or baseUrl) available to the code for use in authenticating and DCS/Gitea API calls.