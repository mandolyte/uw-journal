# Branch Management

## Overview and Definitions

In translation scenarios, there is always a source and target, a "from" language text and a "to" language text.

Even when a translator is simply reviewing or editing their own work, there is also a source and target. The source will be the work-in-progress in the master branch and the target will be the updates in their user branch. I call this the "editor" scenario.

Both source and target work occur in organizations and in a language and for a resource type. Importantly, the work exists in a **repo**, which equates to an organization, language, and resource type.

For an *editor*, there is only a single repo involved. They are updating content in a repo, not translating it.

For a *translator*, two repos are involved. One is the source repo; and the other is the target repo. The source repo is view-only. The target repo will receive updates made by the translators.


## Editor Scenario

In this scenario, a single repo is involved.

**Source**: content is read from the master branch. This is on-going work-in-progress that has been approved, but not yet released (unless no updates have been made since the last release).

**Target**: content is read from editor's user branch. If it doesn't exist yet, it is read from the master branch. If any changes are saved, a user branch will be created if it doesn't exist and the changes saved into their branch.

## Translator Scenario

In this scenario, two repos are involved.

**Source**: content is read from the source organization, language, and resource type, which equates to a repo in DCS. This is view only access. The content will come from the latest release for that repo; it will **not** from from the source repo's master branch (which is work in progress).

**Target**: content is read from target organization, language, and resource type, which equates to a (different) repo in DCS. In that repo, it is read from the translator's user branch. 
- If the translator does not have a user branch yet, it will be read from the target master branch. 
- If the content does not exist yet in the master branch (indicating that this is new translation work), then it will come from the source's latest released. (see note 1)
- Regardless of where the content for the target is obtained, once changes are made and saved, they will be saved to translator's user branch in the target repo.

## Notes

1. A translator may wish to 'stick' with a released version of the source during a work effort. This would require persisting this preference somewhere in DCS. Since the translator may use different computers, it cannot be reliably stored locally. At this time, there is no solution proposed to store preferences in DCS. Thus in the above description, the source is said to come from the latest released.