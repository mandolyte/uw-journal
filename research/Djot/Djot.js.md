Link: https://github.com/jgm/djot.js

It uses typescript: so `npm install -g typescript`

It uses make, so `sudo apt-get install build-essential`

It uses jest, so `npm install jest --global` 

The usual `npm i` to install the packages used.

```
$ make test
tsc
yarn test
yarn run v1.22.19
$ jest
 PASS  ./find.spec.ts
 PASS  ./attributes.spec.ts
 PASS  ./inline.spec.ts
 PASS  ./block.spec.ts
 PASS  ./ast.spec.ts
 PASS  ./html.spec.ts

Test Suites: 6 passed, 6 total
Tests:       41 passed, 41 total
Snapshots:   0 total
Time:        5.187 s
Ran all test suites.
Done in 5.67s.
$ 
```
