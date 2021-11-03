# Issue 679

**Update 2021-02-09** New Definition of Done added to issue: 
```
The links on both the target and source side should look like links but should not be clickable.

**DoD:** Clicking a link on either side does nothing.
```

Link: https://github.com/unfoldingWord/tc-create-app/issues/679

*Original Text in Issue:* v1.1.0-rc.4 build 19-6e50a4c (Moved here from [#660](https://github.com/unfoldingWord/tc-create-app/issues/660))  
There is a problem with book reference links. I don't know what the cause of the not found error is, but being sent back to the login page after clicking on the "Back to our site" or the browser back button does not seem very user friendly.

When clicking on a book reference link in the source side (I've seen this in tn and obs-sn) this message is displayed: (there is a screen shot in the issue).

## 2021-02-10 (after consultation)

in markdown-translatable: ./src/core/markdown-converter.ts, added showdown option `openLinksInNewWindow`:
```js
const markdownToHtmlConverter \= new showdown.Converter({openLinksInNewWindow: true});
```
This will open links in new tab/window so that never will destroy your current window.

in tc-create-app: 
- index.css added:
```css
#translatableComponent a {
  pointer-events: none;
}
```

- Translatable.js added the divs and id around the content components:
```js
        <div id='translatableComponent'>
        {translatableComponent}
        </div>
```


## 2021-02-10 

It appears that only md-to-html is involved. However, source and target do not work the same. I need to make source work like target, namely:
- From same markdown, make the reference look clickable, but not really - just like the target works.

There is a function named `markdownToHtml` in  `core/markdown-converter.ts`. It is a wrapper that pre-processes the markdown content and then calls
`markdownToHtmlConverter.makeHtml`. The wrapper doesn't do anything with making md links HTML links.

So it could either be the input filters (an array of regex) or it might be in the converter function.

Here is the wrapper code mentioned above:

```js

export const markdownToHtml = ({ markdown, inputFilters = [] }) => {
  let _markdown = (markdown || '').slice(0);

  // Make "easy" blockquote:
  _markdown = _markdown.replace(/\n\>/g, '  \n\>');
  _markdown = _markdown.replace(/\<br\>\>/g, '  \<br\>\>');

  _markdown = filter({ string: _markdown, filters: inputFilters });

  let html = markdownToHtmlConverter.makeHtml(_markdown);
  html = html.replace(/<br\s.\\?>/ig, '<br/>');

  // Insert NBSP into empty blocks.
  // See above (htmlToMarkdown) where this NBSP is later stripped out.
  if (!html || html === '') {
    html = '<p>&#8203;</p>';
  }
  return html;
};
```

The converter is actually the instance of the third party package `showdown`:
```js 
import showdown from 'showdown';
// elided
const markdownToHtmlConverter = new showdown.Converter();
```

So let's check the filters...
- markdownToHtml() is called with the inputFilters array
	- markdown-html.ts (a test generator)
	- BlockEditable.js
		- in datatable-translatable: Cell.js
		- in markdown-translatable: BlockTranslatable.js (see snippet below). This one is interesting since it returns both source and target in a single Grid container, side by side. Notice the editable attribute: false for source and true for target.
			- SectionTranslatable.js: here an array of Block Translatables is constructed and placed with a React expansion component with Title, etc. A snippet is below.

				The filters are just passed down.
				- DocumentTranslatable.js: this is, in turn, an array of Section Translatables. A snippet is below.
					- Translatable.js: this one uses both document and section translatables. It switches between depending on whether the content is "sectionable".




**BlockTranslatable.js Snippet**
```js
  const originalBlock = useMemo(() => (
    <BlockEditable
      markdown={original}
      inputFilters={inputFilters}
      outputFilters={outputFilters}
      preview={preview}
      editable={false}
    />
  ), [original, inputFilters, outputFilters, preview]);

  const translationBlock = useMemo(() => (
    <BlockEditable
      markdown={translation}
      onEdit={onTranslation}
      inputFilters={inputFilters}
      outputFilters={outputFilters}
      preview={preview}
      editable={true}
    />
  ), [translation, onTranslation, inputFilters, outputFilters, preview]);
```

**SectionTranslatable.js Snippet**
```js
	_blocksTranslatables.push(
	<BlockTranslatable
	  key={key}
	  original={originalBlock}
	  translation={translationBlock}
	  inputFilters={inputFilters}
	  outputFilters={outputFilters}
	  onTranslation={_onTranslation}
	  preview={preview}
	/>
	);
```

**DocumentTranslable.js Snippet**
```js
  _sectionsTranslatables.push (
	<SectionTranslatable
	key={key}
	original={originalSection}
	translation={translationSection}
	inputFilters={inputFilters}
	outputFilters={outputFilters}
	onTranslation={__onTranslation}
	onExpanded={onExpanded}
	expanded={expanded}
	preview={preview}
	blockable={blockable}
	style={style}
	/>
);
```

## 2021-02-09
To replicate issue:
1. run: https://develop--tc-create-app.netlify.app/
2. login as tcc001 and get to obs-sn resource and select: `content/01/01.md`
3. scroll down to the heading "God's Spirit".
4. there find a couple of obs references (24:08) and (42.10).
5. these two have (invalid) URLs:
	1. https://develop--tc-create-app.netlify.app/24/08
	2. https://develop--tc-create-app.netlify.app/42/10

The raw markdown is:
```md
# God’s Spirit
This is the Holy Spirit (See: [24:08](24/08), [42:10](42/10)).
```

When one of these is clicked you get Netlify's 404 error page. Then if you click to return to the web site, you are returned to the login page. **NOTE! this is restarting tc-create from scratch.**

*However, if click the "keep me logged in" toggle, then you will be returned to where you left off.*

Since this link is to a completely different file in this resource, in order for the link to work, tc-create would have to support routing. And if you had not saved your work, it would be lost.

The link needed would have to look like one of the two methods below:

The query parameter method:
`https://develop--tc-create-app.netlify.app?org=translate_test&resource=obs_sn&language=en&file=content/24/08/08.md`

The URL method:
`https://develop--tc-create-app.netlify.app/translate_test/en_obs_sn/content/24/08/08.md`

AFAIK, tc-create does not support any kind of internal routing.

