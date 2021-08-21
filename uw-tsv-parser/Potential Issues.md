# Potential Issues


Yesterday, I realized something that gave me pause about this package. It has to do with how the spec deals with tab characters. In order to allow tab characters to be in the text, they are first encoded as two characters, namely, a backslash followed by a lower case "t".

Furthermore:
- unbeknownst to most is that tc-create uses a tab character to separate the source and target values in a single cell. The source and target are obviously different files, but they are combined into a single table (2D array). This aligns the two TSV files, ordering the source to be like the target (and putting all the no-match rows at the end). *It may be possible to change how the source and target are aligned, but it would be a major change.*
- tab key presses are processed by the browser and cannot be entered into a field in tc-create. *(don't know if this behavior can be overridden or not)*

There are some side effects to this:
- users must avoid using tabs when they create new content; instead they must enter the two character encoding `\t`. If a tab character is present in new content, it will interfere with source and target alignment.
- plus, users cannot enter the encoding for a tab character either when using tc-create, since the backslash will be encoded as backslash-backslash; it will not survive when written to DCS as `\t`.

So with the above in mind, *do users have any need to actually use tab characters in the Markdown cells of a TSV?*