# Djot

Links: 
- Playground: https://djot.net/playground/
- Syntax Specification: https://htmlpreview.github.io/?https://github.com/jgm/djot/blob/master/doc/syntax.html#reference-link-definition
- Repo: Lua Version https://github.com/jgm/djot (note: the readme has some basic info)
- Repo: Js/Ts Version https://github.com/jgm/djot.js
- Cheatsheet: https://github.com/jgm/djot/blob/main/doc/cheatsheet.md


## Idea: use Djot to mark up scripture text

Here is raw text from https://www.esv.org/Jude/
```
JUDE
Greeting
1 Jude, a servant1 of Jesus Christ and brother of James,

aTo those who are called, bbeloved in God the Father and ckept for2 Jesus Christ:

2 May dmercy, epeace, and love be multiplied to you.
```

I created this repo to test this:
https://github.com/mandolyte/experiment-djot

See the folder "jude", where:
- `jude.djot`: this is the markdown file with djot additions
- `run_pandoc.sh`: this the script to run the pandoc command to create the `jude.html` file
- `styles.css`: a little CSS just to confirm things are working as they should

