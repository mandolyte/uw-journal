## 2022-11-08

Here are the main dependencies for uw-editor

**uw-editor** is at 0.1.11
  "dependencies": {
    "@xelah/type-perf-html": "^1.0.0",
    "epitelete-html": "^0.2.10",
    "proskomma-react-hooks": "^2.4.0"
  },
  "peerDependencies": {
    "@mui/material": "^5.10.12",
    "@mui/styles": "^5.10.10",
    "@mui/icons-material": "^5.10.9",
    "@mui/styled-engine": "npm:@mui/styled-engine-sc@latest",
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "translation-helps-rcl": "3.3.4-rc.2"
  },

NPM for each:
**deps**
- xelah: is current
- latest is 0.2.15: is behind
- pk react hooks has *latest* at 2.4.0-rc.1 but can't be right since there is a 2.4.0 that exists already; have asked the PK team about it.
**peer**
- translation helps is current at 3.3.4-rc.2, but there is a *next* at 3.3.4-rc.3
- react 18.2.0: is current
- mui: material 5.10.12; styles 5.10.10; icons-material 5.10.9; styled-engine @latest



**gateway-translate**
  "dependencies": {
    "@emotion/is-prop-valid": "^1.2.0",
    "@emotion/react": "^11.10.4",
    "@emotion/styled": "^11.10.4",
    "@material-ui/core": "^4.12.3",
    "@material-ui/icons": "^4.11.2",
    "@material-ui/lab": "^4.0.0-alpha.57",
    "@mui/icons-material": "^5.10.6",
    "@mui/material": "^5.10.6",
    "@sendgrid/mail": "^7.7.0",
    "autoprefixer": "^10.4.7",
    "base-64": "^1.0.0",
    "core-js": "^3.22.8",
    "dcs-js": "^1.4.1",
    "deep-equal": "^2.0.5",
    "gitea-react-toolkit": "2.2.2",
    "localforage": "^1.10.0",
    "next": "^12.1.6",
    "postcss": "^8.4.14",
    "react": "^18.1.0",
    "react-dom": "^18.1.0",
    "react-json-view": "1.21.3",
    "regenerator-runtime": "^0.13.9",
    "resource-workspace-rcl": "^2.1.1",
    "scripture-resources-rcl": "^5.2.1",
    "single-scripture-rcl": "^3.2.0",
    "tailwindcss": "^3.1.2",
    "translation-helps-rcl": "^3.3.4-rc.3",
    "use-deep-compare-effect": "^1.8.1",
    "utf8": "3.0.0",
    "uw-editor": "0.1.11"