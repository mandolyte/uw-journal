## 2022-11-02

Links to USFM in the wild:
https://raw.githubusercontent.com/Proskomma/perfidy/main/dataSources/lsg_tit.usfm
https://qa.door43.org/unfoldingWord/fr_ulb/raw/branch/master/57-TIT.usfm
https://git.door43.org/cecil.new/en_ult/raw/branch/master/sample.usfm

Greek:
Titus: https://qa.door43.org/unfoldingWord/el-x-koine_ugnt/raw/branch/master/57-TIT.usfm

Hebrew:
Ruth: https://qa.door43.org/unfoldingWord/hbo_uhb/raw/branch/master/08-RUT.usfm

## 2022-11-09

### On alignment

Definitions:
1. **source text:** this is the text which is the standard to which other texts are aligned; for example, the Greek or Hebrew original language texts.
2. **target text:** this the text that will be aligned to the source

*Inputs:*
- the alignment source text
- the text to be aligned

*Outputs:*
- the updated text

**Titus data for verse 1:1**

*Greek*
```
\v 1
\w Παῦλος|lemma="Παῦλος" strong="G39720" x-morph="Gr,N,,,,,NMS,"\w*,
```

*ULT*
```
\v 1 \zaln-s |x-strong="G39720" x-lemma="Παῦλος" x-morph="Gr,N,,,,,NMS," x-occurrence="1" x-occurrences="1" x-content="Παῦλος"\*\w Paul|x-occurrence="1" x-occurrences="1"\w*\zaln-e\*,
```

*UST*
```
\v 1 \zaln-s |x-strong="G39720" x-lemma="Παῦλος" x-morph="Gr,N,,,,,NMS," x-occurrence="1" x-occurrences="1" x-content="Παῦλος"\*\w I|x-occurrence="1" x-occurrences="3"\w*,
\w Paul|x-occurrence="1" x-occurrences="1"\w*,
```

*es-419_glt*
```
\v 1
\zaln-s |x-strong="G39720" x-lemma="Παῦλος" x-morph="Gr,N,,,,,NMS," x-occurrence="1" x-occurrences="1" x-content="Παῦλος"\*\w Pablo|x-occurrence="1" x-occurrences="1"\w*\zaln-e\*,
```

*hi_glt*
```
\v 1 \zaln-s | x-strong="G39720" x-lemma="Παῦλος" x-morph="Gr,N,,,,,NMS," x-occurrence="1" x-occurrences="1" x-content="Παῦλος"\*\w पौलुस|x-occurrence="1" x-occurrences="1"\w*\zaln-e\*,
```


## 2022-10-11

### Layers of Goals

- Resuability by the OCE community (see note 1)
- Usable by multiple uW applications (see note 2)

### Notes

1. On Resusability by the OCE community:
	- Based on common data formats (see note 3), where applicable.
	- Framework independence
2. On common data formats. As much as possible, data formats that can be widely used by any programming language; priority on human readability (doesn't need a piece of special software to "view" it) and, perhaps even human editibility. Such data formats include at present:
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
