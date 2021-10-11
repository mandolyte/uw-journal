# README
This records notes about Next.js learnings...

## 2021-10-06

Notes on creating a Next.js application template for uW will be in [here](./Creating%20a%20uW%20Next.js%20Template%20Application)

## 2021-08-31

Reading https://nextjs.org/docs/basic-features/data-fetching

This:
```text
Note that getStaticProps runs only on the server-side. It will never be run on the client-side. It won’t even be included in the JS bundle for the browser. That means you can write code such as direct database queries without them being sent to browsers.
```

but later, this:
```text
fallback: true is useful if your app has a very large number of static pages that depend on data (think: a very large e-commerce site). You want to pre-render all product pages, but then your builds would take forever.

Instead, you may statically generate a small subset of pages and use fallback: true for the rest. When someone requests a page that’s not generated yet, the user will see the page with a loading indicator. Shortly after, getStaticProps finishes and the page will be rendered with the requested data. From now on, everyone who requests the same page will get the statically pre-rendered page.

This ensures that users always have a fast experience while preserving fast builds and the benefits of Static Generation.
```
These two seem to contradict each other.

## 2021-08-30 

Working on the tutorial from:
https://nextjs.org/learn/basics/create-nextjs-app

- React components are exported from the pages folder.
- "pages" are associated with a route based on their filename.
- the file `pages/index.js` is associated with the route "/"; whereas `pages/posts/first-post.js` has the route `/posts/first-post` (note the absence of the "js" extension).
- the first-post could be as simple as:
```js
export default function FirstPost() {
  return <h1>First Post</h1>
}
```
   but note that it *must* be marked as **default**
- There doesn't seem to be a need to clutter the folders with an index js file to export things either. So I'm guessing that common code bits will be outside the "pages" folder.
- In next js you can use `{' '}` to create a space (that won't get discarded when rendered in HTML).
- here is what the next "link" looks like:

```js
<h1 className="title">
  Read{' '}
  <Link href="/posts/first-post">
    <a>this page!</a>
  </Link>
</h1>
```

- note the Link element and use of the "a" element to provide the text for the link.
- next js also has an image package which lazy loads, resizes, and optimizes for you. Example:

```js
import Image from 'next/image'

const YourComponent = () => (
  <Image
    src="/images/profile.jpg" // Route of the image file
    height={144} // Desired size with correct aspect ratio
    width={144} // Desired size with correct aspect ratio
    alt="Your Name"
  />
)
```
- they also have a "Head" element that can be used instead of HTML "head"
- `_app.js`: used to add styling CSS; this is also common top level component which will be used across all the pages. For example:

```js
export default function App({ Component, pageProps }) {
  return <Component {...pageProps} />
}
```
   You can use this `App` component to keep state when navigating between pages (that is nice!).
- `_document.js`: used to modify the `<html>` element (such as to add a lang attribute)
- getStaticProps can be used to acquire data at build/compile time. Example:

```js
export async function getStaticProps() {
  const allPostsData = getSortedPostsData()
  return {
    props: {
      allPostsData
    }
  }
}
```
- getServerSideProps can be used to acquire data each time the page renders. Example:

```js
export async function getServerSideProps(context) {
  /* async fetch or whatever here */
  return {
    props: {
      // props for your component
    }
  }
}
```

- lots more info on data fetching and example for next js at https://nextjs.org/docs/basic-features/data-fetching


Today, I completed thru https://nextjs.org/learn/basics/dynamic-routes/dynamic-routes-details. Next up is the "API Routes" lesson.