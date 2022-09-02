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

### Flavor Types
*Burritos come in four flavor **types**:*
1. Scripture
2. Gloss - includes narratives, stories, etc.
3. Parascriptural - includes anything indexed by book, chapter, verse that is not Scripture, e.g. commentaries or syntactic notes.
4. Peripheral - any other resource related to Scripture

### Flavors
*They come in different **flavors**:*
Burritos exist in a number of flavors. Flavors are distinguished by their FlavorType and reference system.

A reference system identifies the way that a resource is referenced and navigated. For instance, a resource may use BCV (book, chapter, verse).

### Ingredients
*Burritos have ingredients:*
Burritos contain ingredients. An ingredient is a file-like resource with a mime-type and, optionally, a scope or role.

This specification places no constraints on the file layout used for ingredients but strongly recommends they be placed in a ingredients/ directory. It is further recommended that application-specific files be placed within a sub-directory under this.

### Variants
Variants provide a mechanism for distinguishing source burritos from derived burritos. 
- A source variant is a user modifiable burrito. 
- A derived variant is programmatically derived from a source variant.

## Goals[](https://docs.burrito.bible/en/v1.0.0-rc1/introduction/overview.html#goals "Permalink to this headline")

1.  Scripture Burrito is designed first and foremost for **data interchange** between ecosystems, although creators and consumers may also choose to use some or all of the format internally.
2.  Scripture Burrito is **a Bible-lifespan format**. In other words, it is intended to be used from the start of the translation, through checking and community testing, into publication via multiple toolchains, and then through revision processes.
3.  Scripture Burrito supports **non-text formats as first-class content**. In other words, the model is not “text plus multimedia”. In some cases text may play a secondary role or even be absent (eg in the case of oral translation or sign-language projects).
4.  Scripture Burrito assumes the existence of **ecosystem servers** that provide ids for users, organizations and projects, and stores information to enable that server-hosted context to be discovered.
5.  Scripture Burrito is intended to allow **lossless roundtripping of projects between ecosystems**. This depends to some extent on references to ecosystem servers that enable reconnection with different ecosystem-specific contexts.
6.  Scripture Burrito supports **Scripture content** (original languages and translations), but also **Scriptural content** (eg glosses) and **Scripture-related content** (eg commentaries, translation manuals).

