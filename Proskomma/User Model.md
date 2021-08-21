# User Model
See details in documentation at:
https://doc.proskomma.bible/en/dev/user_model/building_blocks.html

At the end, there is a "store" - which is a single variable containing all the data. 
- a store may have multiple docsets
- a docset may have multiple documents
- documents are broken down into sequences, blocks, and items.

## docsets
A store has one or more "docsets". A docset which is a collection of documents. Each docset has one or more "selectors" that are used to uniquely identify the docset. 

The selectors form a composite, primary key for a collection (that is, a docset). The selectors are an array of strings. The default selectors are language code and an abbreviation. For example: "eng" and "lt".

Tags may also be assigned to a docset. A tag is a string. A tag may also have qname-like format (str:val). The qname format is intended to support key-value semantics.

## document

## sequence
Quoting:
>A sequence is one continuous piece of text content. Each document has exactly one main sequence and zero or more other sequences which are linked, directly or indirectly, to the main sequence by grafts. For scripture, the main sequence contains the canonical text of one book of the Bible.

The notion of a graft is similar to how our TWL files work. A document about a word is connected to a BCV, word, and occurrence number. Like a graft, then, this content about the text at a specific location.