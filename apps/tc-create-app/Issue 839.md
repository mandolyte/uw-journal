# Copy and Paste Issue
All pastes in tc-create need to be "unformatted".


From [here](https://stackoverflow.com/questions/58980235/stop-pasting-html-style-in-a-contenteditable-div-only-paste-the-plain-text)

```js
/* Derrived from: https://stackoverflow.com/a/6035265/1762224 */
const onPastePlainText = (e) => {
  var pastedText = undefined;
  if (window.clipboardData && window.clipboardData.getData) { // IE
    pastedText = window.clipboardData.getData('Text');
  } else if (e.clipboardData && e.clipboardData.getData) {
    pastedText = e.clipboardData.getData('text/plain');
  }
  e.target.textContent = pastedText;
  e.preventDefault();
  return false;
}

document.querySelector('.ediatable-div').addEventListener('paste', onPastePlainText);
```

Found this in markdown-translatable:
```js
export function useHandlePaste(el, preview) {
  const handlePaste = useCallback((e) => {
    e.preventDefault();
    const pastedData = e.clipboardData.getData('text/plain');
    const doc = new DOMParser().parseFromString(pastedData, 'text/html');
    const text = doc.body.textContent || '';
    document.execCommand('insertHTML', false, text);
  }, []);

  useEffect(() => {
    if (el) {
      el.addEventListener('paste', handlePaste);
    };
    return () => {
      if (el) {
        el.removeEventListener('paste', handlePaste);
      }
    };
  }, [el, handlePaste, preview]);
}
```

```html
<div class="block" contenteditable="true" dir="" style="width: 100%; white-space: pre-wrap;">
```

**Before**
```js
  return (
    <pre className={classes.pre}>
      <code
        className={classes.markdown}
        style={{ ..._style, fontSize }}
        dir='auto'
        contentEditable={editable}
        dangerouslySetInnerHTML={dangerouslySetInnerHTML}
        onBlur={handleRawBlur}
        onKeyPress={handleKeyPress}
        onKeyUp={handleKeyUp}
        onCut={handleCutPaste}
        onPaste={handleCutPaste}
        data-test="blockeditable-editable-markdown-pre"
      />
    </pre>
  );

```

**After**
```js
  return (
    <div 
        className={classes.markdown}
        style={{ ..._style, fontSize, whiteSpace: 'pre-wrap' }}
        dir='auto'
        contentEditable={editable}
        dangerouslySetInnerHTML={dangerouslySetInnerHTML}
        onBlur={handleRawBlur}
        onKeyPress={handleKeyPress}
        onKeyUp={handleKeyUp}
        onCut={handleCutPaste}
        onPaste={handleCutPaste}
        data-test="blockeditable-editable-markdown-pre"
      >
    </div>
```
