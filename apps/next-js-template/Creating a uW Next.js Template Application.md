# Creating a uW Next.js Template 

Link: https://next-js-template.netlify.app/


## History

This file started out before we began work on the template. Now it exists in box3 [here](https://github.com/unfoldingWord-box3/next-js-template)

Have moved this file to a folder named after the repo.

## Links

Some history:
https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE--.20UR/topic/Next.2Ejs.20Template/near/251546182

The Clickup task: 
https://app.clickup.com/t/6f6mdz

Figma Mockups:
https://www.figma.com/file/yTcdc2Shb0rwaDlxD11zU8/tcCreate-BPF?node-id=2%3A4141

## 2021-11-01

As of today, login and an arrangement of text on paper objects are shown.

I self assigned issue 14 to create a uW image card.

*Task 1.* find a vector form of the uW logo. There is an eps format and Chrome can handle it.

*Task 2.* Create a repo for this "RCL" (React Component Library) in box3. Do this by "forking" the translation-helps-rcl. Let's call it: uw-card-rcl.

*Task 3.* Trim out everything except what is required for the "card" component.

*Task 4.* Modify the card component to take, as a parameter, its content.

*Task 5.* Create some demos for it. At least these:
- Take the uW logo as content and let the logo resize as the card is resized.
- Take some Material-UI text (Typography component) and let the text re-flow as the card is resized.




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

