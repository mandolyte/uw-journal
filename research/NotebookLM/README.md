# 2023-10-24

At this time there are still some restrictions that make it difficult to import moderate sized notes into NbLM. Here is a script I wrote that combines gdocs into a single PDF.

1. Step 1. Download a folder from Google Drive. This will download a zip of `.docx` files/folders.
2. Step 2. Run the script with adjustments as needed.

Here is the script:
```sh
#!/bin/sh

## variables
GDFOLDER="New-Family-Recipes"
FILELIST="_filelist"
TARGETFOLDER="recipes-md"
TARGETPDF="my-recipe-book.pdf"

## First, create list of all the files
rm -f $FILELIST
find $GDFOLDER -type f -name "*.docx" > $FILELIST

## Second, loop thru the file names and 
## convert each docx to a markdown with
## front matter containing title, which is
## the base filename
rm -rf $TARGETFOLDER
mkdir $TARGETFOLDER
while read -r line
do
	#echo filename is $line
	fn=`basename -- "${line}" .docx`
	#echo Basename is $fn
	mdname=${fn}.md
	#echo Markdown name is $mdname
	pandoc "$line" -s -o "$TARGETFOLDER/$mdname" --metadata title="$fn"
done < $FILELIST

## Next create a PDF each of the markdown content files
cd $TARGETFOLDER
find . -type f -name "*.md" > $FILELIST

while read -r line
do
	pandoc "$line" \
	-t html \
	--output="${line}.pdf"
done < $FILELIST 

## Finally, combine them all into a PDF in the parent directory
pdfunite *.md.pdf ../$TARGETPDF

```