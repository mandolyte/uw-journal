# Thought Experiment - 4D Text

## Basic underpinnings

Given Hebrew and Greek texts under version control, then could the text at some version, say v1, be converted to a four dimensional array, where the dimensions (major to minor) are:
1. Book
2. Chapter
3. Verse
4. Word

This 4D array would be "irregular", since the minor dimensions are not same across books, chapters, and verses.

This array could be statically built for applications to use. Normal life cycle control methods would be used to update downstream when text is updated, released, and bumped to a new version.

With this in place any individual word could specified by a 4-tuple (b,c,v,w), that is, four integers representing book, chapter, verse, and word in verse.

With this scheme, ranges could be built. For example, John 11:35 "ὁ  Ἰησοῦς ἐδάκρυσεν" ("Jesus wept") would be (69,11,35,0)-(69,11,35,2). Discontiguous ranges could be specified with a comma delimited list of ranges: (b1,c1,v1,w1)-(b2,c2,v2,w2),(b3,c3,v3,w3)-(b4,c4,v4,w4).

Since the data is an array, access would be fast.

Summary of underpinnings:
- Text at some versioned release is transformed into a static 4D array.
- An API using 4-tuples is used to access any word, range of words, or any set of discontinuous ranges.

## Overlays

An overlay is set of instrutions that augments the text. For example, in the `en_ult` text for this verse is this:
`x-strong="G24240" x-lemma="Ἰησοῦς"`, which associates the Strong's number to the second word (Jesus) in the text.

So the overlay instruction would have the following:
- Since a single word, the range specification would be a single 4-tuple: (69,11,35,1)
- The attribute indicating this is a Strong's number, say: "strong"
- And the value for the attribute, in this case: "G24240"

Any number of overlays for different sorts of needs could simultaneously exist. Examples used today in the Original Language text are:
- Translation Words
- Strongs

## Possible implementation problems

The primary one that comes to mind is how much memory is needed for an array containing every single word in the Hebrew and Greek texts. 
Without trying it, not sure if the problem is a real one or not. If it was, then the concepts could be applied at a lower level, for example, one 4D array per book. The downside is that this could not be statically built into the apps, but would have be fetched, thus losing some performance.
