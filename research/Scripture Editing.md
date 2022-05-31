# Scripture Editing

## 2022-05-07
- Downloaded the uW en_ult
- Wrote this script (gettags.sh):

```sh
#!/bin/sh
OUTPUT=allbooktags.txt
rm -f $OUTPUT

for i in `ls *.usfm` 
do
    sed -e "s/ .*$//" < $i \
        | sort \
        | grep -E "^[\\]" \
        | uniq -c \
        | sed -e "s/^ *//" \
        | sed -e "s/^/$i /" \
        >> $OUTPUT
done
```

- Worked with the data in Google Sheets [here](https://docs.google.com/spreadsheets/d/1xzUJWwNuNX9ZpZanqp5WuocLJAkFoq4ZHwbSOmwhKS4/edit#gid=0)