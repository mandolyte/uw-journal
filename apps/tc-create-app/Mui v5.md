# Mui v5
## Initial look using datatable-translatable
Refs:
- youtube video walkthru to upgrade  https://www.youtube.com/watch?v=3GY0j-a5h5o
- current styleguidist demo at: https://datatable-translatable.netlify.app/

**Questions**
1. is there a newer version of mui-datatables? Yes. https://www.npmjs.com/package/mui-datatables is at 4.2.2, updated a month ago in April, 2022. In the compatibility chart, it shows that it must be at v4 or above.

2. what are the refs to mui in the package.json

Here is the relevant info in package.json (lines elided around lines of interest):
```json
  "dependencies": {
    "mui-datatables": "3.3.1",
  },
  "devDependencies": {
    "@material-ui/core": "^4.6.1",
    "@material-ui/icons": "^4.5.1",
  },
  "peerDependencies": {
    "@material-ui/core": "^4.6.1",
    "@material-ui/icons": "^4.5.1",
  },
```

3. Are any changes in the react support? In MUI sandbox examples, they are using react 18. Per youtube video, react must be at least 17.

## 2022-05-10

Since yesterday ended with errors in markdown-translatable, let's see how far I get with it. It is the "bottom" of the include tree. So perhaps the better part of wisdom to work up from the leaves.

List of things to do:
- set-branch upgrade-to-muiv5
- Change peer dependencies. Currently:
```json
  "peerDependencies": {
    "@material-ui/core": "^4.6.1",
    "@material-ui/icons": "^4.5.1",
    "react": "^16.8.6",
    "react-dom": "^16.8.6"
  },
```
Should be:
```json
  "peerDependencies": {
	"@mui/icons-material": "^5.6.2",
    "react": "^17.0.2",
    "react-dom": "^17.0.2"
  },
```

- Changed a number of things in devDependencies to match what I had done in datatable-translatable.
- Notice these two in dependencies:
```json
    "react-headroom": "^3.1.0",
    "react-markdown": "4.0.6",
```
- the headroom is at v3.2.0 - not too far behind; will update to this
- the markdown is at v8.0.3 - we are way behind; will update to this
- rm -rf node_modules
- yarn install went smoothly first time
- yarn start had lots of compile errors (wrong imports). For instance:
```
./src/components/actions/Actions.js
Module not found: Can't resolve '@material-ui/core' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\src\components\actions'
```
- The mui folks have split out a lot of stuff to separate modules. So I have also added:
	- @mui/material (buttons and such)
	- @mui/styles (makeStyles, styling, etc.)
	- Also had to add @emotion/react

More errors:
```
./node_modules/@mui/styled-engine/index.js
Module not found: Can't resolve '@emotion/styled' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\node_modules\@mui\styled-engine'
./src/components/block-translatable/BlockTranslatable.js
Module not found: Can't resolve '@material-ui/core' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\src\components\block-translatable'
./src/components/section-translatable/SectionTranslatable.js
Module not found: Can't resolve '@material-ui/core' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\src\components\section-translatable'
./src/components/translatable/Translatable.js
Module not found: Can't resolve '@material-ui/core' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\src\components\translatable'
./src/components/block-editable/useStyles.js
Module not found: Can't resolve '@material-ui/core/styles' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\src\components\block-editable'
./src/components/block-translatable/BlockTranslatable.js
Module not found: Can't resolve '@material-ui/core/styles' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\src\components\block-translatable'
./src/components/document-translatable/DocumentTranslatable.js
Module not found: Can't resolve '@material-ui/core/styles' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\src\components\document-translatable'
./src/components/section-translatable/styles.js
Module not found: Can't resolve '@material-ui/core/styles' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\src\components\section-translatable'
./src/components/translatable/Translatable.js
Module not found: Can't resolve '@material-ui/core/styles' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\src\components\translatable'
./src/components/section-translatable/SectionTranslatable.js
Module not found: Can't resolve '@material-ui/icons' in 'C:\Users\mando\Projects\github.com\unfoldingword\markdown-translatable\src\components\section-translatable'
```


## 2022-05-09

### Yarn Install work
Looks like `@mui/core` is deprecated. However, in the v5 sandbox demos, they don't include it. Thus perhaps the component include what they need and I don't need to include core at all. *Note: says it was replaced by @mui-base. See https://www.npmjs.com/package/@mui/core*

icons package is now: @mui/icons-material and is at v5.6.2

This finally compiled with yarn install:
```json
  "dependencies": {
    "deep-freeze": "^0.0.1",
    "lodash.isequal": "^4.5.0",
    "markdown-translatable": "1.3.1-rc.1",
    "mui-datatables": "4.2.2",
    "use-deep-compare-effect": "^1.3.1"
  },
  "devDependencies": {
    "@babel/cli": "^7.5.0",
    "@babel/core": "^7.4.5",
    "@babel/preset-env": "^7.4.5",
    "@babel/preset-react": "^7.0.0",
    "@mui/icons-material": "^5.6.2",
    "@types/jest": "24.0.22",
    "@unfoldingword/eslint-config": "^1.3.0",
    "coveralls": "^3.0.9",
    "eslint-plugin-jest": "^24.0.0",
    "eslint-plugin-prettier": "^3.1.4",
    "eslint-plugin-react": "^7.20.6",
    "jest": "^26.6.3",
    "react": "^17.0.2",
    "react-docgen": "5.4.0",
    "react-dom": "^17.0.2",
    "react-scripts": "^4.0.3",
    "react-styleguidist": "^11.2.0",
    "react-test-renderer": "^17.0.2",
    "webpack": "^4.0.0"
  },
  "peerDependencies": {
    "@material-ui/core": "^5.0.0",
    "@material-ui/icons": "^5.0.0",
    "react": "^17.0.2",
    "react-dom": "^17.0.2"
  },
```

### Yarn Start work

Looks like I need to add @emotion/react to package.json; added:
```
    "@emotion/react": "11.9.0",
```

Now I'm getting errors from markdown-translatable...
