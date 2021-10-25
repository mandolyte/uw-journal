# Creating a uW Next.js Template Application

## Links

Some history:
https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE--.20UR/topic/Next.2Ejs.20Template/near/251546182

The Clickup task: 
https://app.clickup.com/t/6f6mdz

Figma Mockups:
https://www.figma.com/file/yTcdc2Shb0rwaDlxD11zU8/tcCreate-BPF?node-id=2%3A4141

## 2021-10-21


## 2021-10-18

Experiment: 
- Take a zip of gatewayEdit and trim out the unneeded things
- What should show on the main page? A uW logo?
- Then using `example` options [here](https://nextjs.org/docs/api-reference/create-next-app#options), try to use the API to create a new app from the template.

Steps: 
1. Create a box3 repo to house the template. Used `https://github.com/unfoldingWord-box3/next-js-template`
2. Download the zip and extract contents into the box3 org folder
3. Made some small tweaks and pushed

## 2021-10-05

Initial thoughts about what to include in the template (from the Zulip link):
- An app bar with a side drawer (which I assume at least contains account settings, feedback report, and logout) 
- A tabbed UI component (perhaps include three of them?) 
- A "add" action to the right of the tabs A canvas for the cards with content 
- Since existing templates actually can be run as-is, I think we should also consider something useful as part of the template (in addition to working login, logout, account settings, and feedback). For example, maybe there is an "avatar" tab with the action "Add Avatar".

