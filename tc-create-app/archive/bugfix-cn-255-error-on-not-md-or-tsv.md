# bugfix-cn-255-on-not-md-or-tsv

Code fix was lost somehow...

- In this file:
https://github.com/unfoldingWord/tc-create-app/blob/develop/src/components/translatable/Translatable.js

- Look at commit history
- Find my last commit; was https://github.com/unfoldingWord/tc-create-app/commit/93e18219d10883ca0964de284939489dd61ad409#diff-214267ff898e8838818a312a414d3f02
- lines of code lost were:
```js
} else {
    _translatable = <h3 style={{ 'text-align': 'center'}} >Unsupported File. Please select .md or .tsv files.</h3>;
}
```

With code, this will now show the error if you select say a YAML file.