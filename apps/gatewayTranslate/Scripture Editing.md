# Scripture Editing

## 2022-10-11

### Layers of Goals

- Resuability by the OCE community (see note 1)
- Usable by multiple uW applications (see note 2)

### Notes

1. On Resusability by the OCE community:
	- Based on common data formats (see note 3), where applicable.
	- Framework independence

2. sdfsdf
3. On common data formats. As much as possible, data formats that can be widely used by any programming language; priority on human readability (doesn't need a piece of special software to "view" it) and, perhaps even human editibility. Such data formats include at present:
	- Markdown
	- USFM
	- CSV/TSV
	- JSON


## 2022-09-14

Netlify: https://perf-html-editor-app.netlify.app/
Github: https://github.com/klappy/perf-html-editor-app

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
