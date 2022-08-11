# Djot

Links: 
- Playground: https://djot.net/playground/
- Syntax Specification: https://htmlpreview.github.io/?https://github.com/jgm/djot/blob/master/doc/syntax.html#reference-link-definition
- Repo: https://github.com/jgm/djot (note: the readme has some basic info)
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

Here is what I'll use in the playground:
```
JUDE{.book}

Greeting{.heading}

1{.verse} 
Jude, a servant[^1]{.footnote} of Jesus Christ and brother of James,
[^a]{.reference}To those who are called, 
[^b]{.reference}beloved in God the Father and
[^c]{.reference}kept
for[^2]{.footnote} Jesus Christ:

2{.verse} May 
[^d]{.reference}mercy, 
[^e]{.reference}peace, and love be multiplied to you.
```