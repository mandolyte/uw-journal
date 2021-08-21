# Git Notes

## Script to set remote branch

```sh
#!/bin/sh

if [ "$1x" = "x" ]; then
	echo Feature or bugfix branch name required
	exit
fi

git checkout -b $1
git push --set-upstream origin $1

echo Done.
```
