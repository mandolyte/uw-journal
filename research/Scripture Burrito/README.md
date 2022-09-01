# Scripture Burrito

## Links

Github: https://github.com/bible-technology/scripture-burrito
Docs: https://docs.burrito.bible/en/v1.0.0-rc1/

## Notes from the docs

**Scope**
The scope of the Scripture Burrito specification includes any Bible-centric content and the entire ecosystem, including publication. The intent is not to create new data formats, but rather define a portable way of interchanging existing formats between applications and ecosystems. This necessarily includes existing and future client/server architectures and the attendant need to uniquely identify users, organizations and content.

**Features**

- a single container format to span the entire Scripture life cycle, from translation, through community testing and checking, to publication
- support for canonical and quasi-canonical translations in multiple media (initially text, audio, video, print on demand and braille)
- support for a wide range of resources relating to scripture, such as lexicons, cross-references, translation manuals, and commentaries
- mechanisms for linking related content, both at a burrito-to-burrito and ingredient-to-ingredient level
- identification of people, organizations and content via namespaced ids relating to multiple authenticating servers

## Concepts

*What is a burrito?*
A burrito is a wrapper that contains content and metadata. That wrapper may be made available in various digital formats, such as a zip file, an Amazon S3 bucket or a series of API calls. The term “burrito” describes the wrapper, not the distribution mechanism.

The metadata describes the contents of the burrito, including directory structure and ingredients.


*Burritos come in four flavor **types**:*
1. Scripture
2. Gloss - includes narratives, stories, etc.
3. Parascriptural - includes anything indexed by book, chapter, verse that is not Scripture, e.g. commentaries or syntactic notes.
4. Peripheral - any other resource related to Scripture

*They come in different **flavors**:*
Burritos exist in a number of flavors. Flavors are distinguished by their FlavorType and reference system.

A reference system identifies the way that a resource is referenced and navigated. For instance, a resource may use BCV (book, chapter, verse).

*Burritos have ingredients:*
Burritos contain ingredients. An ingredient is a file-like resource with a mime-type and, optionally, a scope or role.

This specification places no constraints on the file layout used for ingredients but strongly recommends they be placed in a ingredients/ directory. It is further recommended that application-specific files be placed within a sub-directory under this.

