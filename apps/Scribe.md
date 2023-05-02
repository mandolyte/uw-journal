# Sync Proposal

## Requirements

The sync solution for Scribe has the following requirements:

1. Synchronization is done via a server. In this proposal, the server is the Door43 Content Server (DCS) which is a git-based document management solution built on Gitea.
2. Offline first (meaning that most work will be done offline and syncing will be done when a connection is initiated successfully by the user)
3. Synchronization is bi-directional:
   - It must enable updates from other translators who are working on the same book
   - It must enable work to be uploaded to a server as a "backup". This is not intended to be shared.
   - It must enable work to be merged with work from other translators, even on the same book

## Assumptions

The main assumption is that all files produced will be in the same form, for example, as a Scripture Burrito. Furthermore, the file naming will follow a convention, for example, "TIT.usfm" for the book of Titus in USFM format.

# Client Side Requirements

The client is responsible for enforcing the overall shape of the work, for example as a Scripture Burrito. For a given project, whether files are started from scratch or from existing work, they must be in the same form so that work can be combine, shared, or merged with work from collaborators.

# Server Side Requirements

- All translation projects by a team must reside in the same Gitea organization.
- A given translation project must have its own repository.
- All collaborators on a given project will be part of a "team" with rights appropriate with their roles.
- Collaboration on a given project will be effected by normal "git" techniques, namely:
	- Each collaborator will sync and have their own branch
	- The default (master) branch will be the branch where all work is consolidated and merged

# Example Workflow

Scenario: An OL team begins to translate work from a GL.

Prerequisites:
- Accounts for all team members are setup on DCS
- An organization on DCS is created for the team
- All collaborators are added to a team
- A (empty) repo is created for the team

Workflow:
- Each member of the team acquires the GL work which will be translated and does any conversion, re-organization needed to enable use by the OL team.
- As assigned, each member begins the translation work
- At certain points, the translated work is synced to the server. This means:
	- The work is uploaded to a branch in the project repo
	- The branch is named so that collaborators (or the team admin) can see work and understand who produced it and, perhaps, what should be in the branch. *Note: this is mostly about how to name the branches in a standard fashion.*
- Work in branches is understood to be "work-in-progress" and may be in an indeterminate state.
- Once work has reached maturity (this is up to the OL team to determine), then the work is merged into the master branch.
- Work in the master branch is considered mature and ready to share with other collaborators.
- A team member may update their work-in-progress branch with work from the master branch.
	- If the work in master conflicts with their own work, then this must be resolved using conflict management tools in DCS
	- If there is no conflict, then even work on the same book may be used to update their branch.
	- If their branch was updated, then the branch must be used to update the local copy of the work.
	- *It is important that the user branch is synced before attempting to update from master to avoid losing any unsaved changes on the local copy.*

