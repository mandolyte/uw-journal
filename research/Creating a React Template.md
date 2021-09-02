# Creating a React Template

Docs at https://create-react-app.dev/docs/custom-templates

Here is the "typescript template" repo:
https://github.com/facebook/create-react-app/tree/main/packages/cra-template-typescript

The typescript template contains:
- README.md
- package.json
- template.json
- (folder) template

The template folder contains:
- public
- src
- README.md
- gitignore (no leading dot)

## Template features needed

Based on the mockups at 
https://www.figma.com/file/yTcdc2Shb0rwaDlxD11zU8/tcCreate-BPF?node-id=2%3A4141, here is an initial list of features to be included in the template:
- An app bar with a side drawer (which I assume at least contains account settings, feedback report, and logout)
- A tabbed UI component (perhaps include three of them?)
- A "add" action to the right of the tabs
- A canvas for the cards with content

Since existing templates actually can be run as-is, I think we should also consider something useful as part of the template (in addition to working login, logout, account settings, and feedback).

For example, maybe there is an "avatar" tab with the action "Add Avatar".

### Template folder content

Based on gatewayEdit, the template should have:
- a `.husky`folder to track builds
- a Next.js "pages" folder for all pages in the app template
- a public folder for images, etc.
- a scripts folder for commonly used scripts (including the one needed for Zulip notifications)
- the 'src' folder with:
	- a common folder for functions that aren't specific to a page or component
	- components for components that will be within a page
	- context for all react contexts
	- hooks for all react custom hooks
	- styles for all styling CSS
	- utils for utility type code (why are both common and utils needed?)


### Try it!

Searched NPM with `cra-template-next` and picked one. Ran `yarn create react-app my-app --template nextjs-flow`.

On github, the repo (from NPM) is: 
https://github.com/xabierlameiro/cra-template-nextjs-flow

This single command did everything, even installing all the needed packages.

I then:
```sh
$ cd my-app
$ yarn dev
```
which then on http://localhost:3000 showed a simple output page.
