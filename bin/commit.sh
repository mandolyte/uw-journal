#!/bin/sh
MSG="ckpt"
if [ "$1x" != "x" ]
then
	MSG="$1"
fi
git commit -a -m "$MSG"
git push

