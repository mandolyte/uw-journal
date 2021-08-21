# Local Build of DCS
These are instructions to build DCS (Gitea) locally for the purpose of development.
Clone DCS:
git clone git@github.com:unfoldingWord/dcs.git

Install Make (if needed):
```sh
http://www.equation.com/servlet/equation.cmd?fa=make
```
Copy `make.exe` to a folder on the path ($HOME/bin)

Run make:
```bash
TAGS="bindata sqlite sqlite_unlock_notify sqlite_json" make build
```



# Appendix: loading content
from: https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE--.20UR/topic/DCS.20Development/near/244929632

As for populating your local copy of DCS with repos from produciton, follow this guildline of my tx-dev pipeline (I'm still trying to get this into the official README):

https://github.com/unfoldingWord-dev/tx-dev/blob/develop/README_new.md#migrating-resources-from-produciton-dcs

I have a repo that spins up a DCS populated with all our main unfodingWord English resources, but not too sure how to easily get the data dir you can create with it (had to break it down in multiple tarballs due to Github's size contraints on a single file), but here is how you can see it yourself using Docker, and maybe copy the SQLite DB (gitea.db) and the repositoriesa nd other dirs (attachments, etc.) to where you have your dev copy getting its data:

```
git clone  git@github.com:unfoldingWord-dev/tx-dev-dcs.git

cd tx-dev-dcs

./untar_data.sh

docker-compose up -d
```