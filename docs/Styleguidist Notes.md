

https://stackoverflow.com/questions/53558916/babel-7-referenceerror-regeneratorruntime-is-not-defined


## Option 2: Library

**When to use:** ✔ for libraries ✔ _no_ global scope pollution ✔ includes _all_ polyfills, not selective ✔ bigger bundle size neglectable

```json
"plugins": [
  [
    "@babel/plugin-transform-runtime",
    {
      "regenerator": true,
      "corejs": 3
    }
  ]
]
```

Install compile-time and run-time dependencies:

```javascript
npm i --save-dev @babel/plugin-transform-runtime // only for build phase

npm i @babel/runtime // runtime babel helpers + just regenerator runtime
// OR (choose one!)
npm i @babel/runtime-corejs3 
// also contains other JS polyfills (not only regenerator runtime)
// depends on core-js-pure ("ponyfills"/polyfills that don't pollute global scope)
```

used yarn add xxx --dev for devDependencies

