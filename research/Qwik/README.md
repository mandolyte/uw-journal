# README

## Links

Deploy to Netlify
https://www.netlify.com/blog/how-to-deploy-the-qwik-javascript-framework/

FAQ
https://qwik.builder.io/docs/faq/

Articles
- https://www.builder.io/blog/qwik-and-qwik-city-have-reached-beta?ck_subscriber_id=1697731695
- https://qwik.builder.io/docs/think-qwik/
- https://javascript.plainenglish.io/react-and-next-js-is-dead-something-new-is-finally-replacing-it-for-good-c792c48806f6

Docs
- Qwik City https://qwik.builder.io/qwikcity/overview/


## 2022-10-08

1. Appears that Qwik City is the `NextJS` equivalent.
2. They support ".mdx" files! [here](https://qwik.builder.io/qwikcity/routing/overview/#implementing-a-component) --> Pages created with markdown! ([[MDX-Notes]])
3. Can use React components (presumably on MPM):[here](https://qwik.builder.io/docs/faq/#can-i-enjoy-the-rich-react-ecosystem)
4. Nice! Quote:
   > Your component does not need to think about server/client differences when using data.
5. Can make restful API: [here](https://qwik.builder.io/qwikcity/data/endpoints/)
6. A "page" in their framework is an `index.tsx` in a sub-folder under `src/routes`
7. If a page has sub components, they can be added in the same folder and imported. This way you can split things up into nice manageable pieces. This [link](https://qwik.builder.io/qwikcity/content/component/) shows how it works.
8. There are three API (all look like hooks) to discover information that the component may need to know.
	1. useContent: https://qwik.builder.io/qwikcity/api/use-content/
	2. useDocumentHead: https://qwik.builder.io/qwikcity/api/use-document-head/
	3. useLocation: https://qwik.builder.io/qwikcity/api/use-location/
9. Details on using Netlify [here](https://qwik.builder.io/qwikcity/adaptors/netlify-edge/)
10. Docs do not mention yarn at all. So probably safer to use npm as they do.
11. They use Typescript all the time... as we should :-)


