# Pandoc, etc.

## Pandoc
NOTE: apt does not yield a recent version for my Debian Linux under ChromeOS

1. https://pandoc.org/installing.html
2. https://github.com/jgm/pandoc/releases/tag/2.19.2
3. Installed the ".deb" as normal (double-clicking)
4. then:
```
cecil@penguin:~/Projects/github.com/mandolyte/uw-journal$ pandoc --version
pandoc 2.19.2
Compiled with pandoc-types 1.22.2.1, texmath 0.12.5.2, skylighting 0.13,
citeproc 0.8.0.1, ipynb 0.2, hslua 2.2.1
Scripting engine: Lua 5.4
User data directory: /home/cecil/.local/share/pandoc
Copyright (C) 2006-2022 John MacFarlane. Web:  https://pandoc.org
This is free software; see the source for copying conditions. There is no
warranty, not even for merchantability or fitness for a particular purpose.
cecil@penguin:~/Projects/github.com/mandolyte/uw-journal$ 
```

## wkhtmltopdf

1. https://wkhtmltopdf.org/
2. https://wkhtmltopdf.org/downloads.html
3. then install the ".deb" the usual way by double-clicking on it.
```
$ pandoc 'Lesson 1.md'  | wkhtmltopdf - --default-header lesson1a.pdf
Loading pages (1/6)
Counting pages (2/6)                                               
Resolving links (4/6)                                                       
Loading headers and footers (5/6)                                           
Printing pages (6/6)
Done                                                                      
$ 
```

