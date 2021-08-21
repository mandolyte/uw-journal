# Installing

## Installing and testing the code
First, clone the repository locally, then change to the cloned folder and do the following:
```
npm install
npm test
npm run rawTest
TESTSCRIPT=cp_vp npm run testOne
npm run coverage
```

## Running the code
```
cd scripts
node do_graph.js ../test/test_data/usx/web_rut_1.usx example_query.txt
node do_graph.js ../test/test_data/usfm/hello.usfm example_query.txt
```

This folder now has a `scripts` subfolder. To make it work, it needs a `package.json` file and you must run `npm i` to create the node_modules folder.

I have also created a data folder where I'll put various USFM texts to query.

There is a "run" script that can be used.