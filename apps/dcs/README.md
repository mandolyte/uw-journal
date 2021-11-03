# DCS
## Links
Repo (fork of gogs/Gitea):
https://github.com/unfoldingWord/dcs


### Code Links

Modules for git:
https://github.com/unfoldingWord/dcs/tree/main/modules/git

To the function for listing entries by tree:

https://github.com/unfoldingWord/dcs/blob/main/modules/repofiles/tree.go#L34

https://github.com/unfoldingWord/dcs/blob/main/modules/git/tree_nogogit.go#L102

https://github.com/unfoldingWord/dcs/blob/main/modules/repofiles/tree.go#L17

Hacking on Gitea:
https://docs.gitea.io/en-us/hacking-on-gitea/

```bash
TAGS="bindata sqlite sqlite_unlock_notify" make build
```

## Coding Practices

There is a zulip post on how @rich flags his changes or adds new code.


# Zulip Notes

```
Rich Mahn: Cecil New said:

Finally, note that I might be able to help Rich since I know Go pretty well; But I realize that I since I'm not familiar with Gitea's inner workings, I might end up just slowing Rich down. For what it's worth...

@Cecil New I wouldn't mind help, as I would love others to know the Gitea code and partake, but it will involve writing some of our own interfaces with git to find everything we need. 

I don't believe I do that for generating the Door43Metadata table which the Catalog Next API relies on for the Endpoint results. 

Yet that is the best place to start, just seeing what I did for Door43Metadata model/module/nofication objects (the latter is not to notify users, although it could, but just a hook to be told when a release is created/edited/deleted): https://github.com/unfoldingWord/dcs/search?q=Door43Metadata

Also just see about getting Gitea working locally for yourself, and can substitute github.com/unfoldingword/dcs for github.com/go-gitea/gitea: https://docs.gitea.io/en-us/hacking-on-gitea/

Run it and set it up, probably easiest with a MySQL database. It also uses Docker, so you can spin it up with the Dockerfile in the root dir, and spin up a MySQL docker container: https://docs.gitea.io/en-us/install-with-docker/

Rich Mahn: To see some interfaces with git, here's where a bunch of commands are for Gitea:

https://github.com/unfoldingWord/dcs/tree/main/modules/git

unfoldingWord/dcs
Door43 Content Service. Contribute to unfoldingWord/dcs development by creating an account on GitHub.
```