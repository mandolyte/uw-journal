# PDF RCL

## 2022-03-23

Notes from yesterday's work from the team.
- Mark Howe pointed us to the latest `proskomma-render-pdf`
- in `scripts/make_pdf.js` we find the following...
These two imports:
```js
import {doRender} from '../index.js';
import {Proskomma} from 'proskomma';
```

Create an instance of PK:
```js
const pk = new Proskomma();
```

Followed by an import of usfm:
```js
        pk.importDocument(
            {lang: "xxx", abbr: "yyy"},
            contentType,
            content,
            {}
        );
```

Then a "doRender":
```js
const config2 = await doRender(pk, config);
```

Here is the entire index.js where "doRender()" lives:
```js
import {ScriptureParaModel, ScriptureParaModelQuery} from "proskomma-render";
import MainDocSet from "./MainDocSet.js";

const doRender = async (pk, config, docSetIds, documentIds) => {
    let ts = Date.now();
    const doMainRender = (config, result) => {
        const model = new ScriptureParaModel(result, config);
        model.addDocSetModel('default', new MainDocSet(result, model.context, config));
        model.render();
        console.log(`Main rendered in  ${(Date.now() - ts) / 1000} sec`);
        console.log(model.logString());
    }
    const thenFunction = result => {
        console.log(`Query processed in  ${(Date.now() - ts) / 1000} sec`);
        ts = Date.now();
        doMainRender(config, result);
        return config;
    }
    const result = await ScriptureParaModelQuery(pk, docSetIds || [], documentIds || []);
    return thenFunction(result);
};

export {doRender}
```

## 2022-03-21

Link:
https://docs.google.com/document/d/1ghjWzLQLdjQIILrfqZzgehDHzoTHSdOlpG--_Jjyo8M/edit

Clickup card: https://app.clickup.com/t/1w4qjzd

