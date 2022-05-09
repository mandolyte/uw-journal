# Mui v5

## datatable-translatable

Refs:
- https://www.youtube.com/watch?v=3GY0j-a5h5o

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

## 2022-05-09

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

