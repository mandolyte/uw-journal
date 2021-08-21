package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	now := time.Now().UTC()
	outputFile := flag.String("o", "", "Output path/filename, required")
	flag.Parse()

	if *outputFile == "" {
		log.Fatalln("Output file name is missing")
	}

	log.Printf("Start\n")

	var results [][]string
	header := []string{"gw", "lc", "ld", "ang"}
	results = append(results, header)

	results = append(results, getLanguages()...)
	f, err := os.Create(*outputFile)
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(results) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}

	stop := time.Since(now)
	log.Printf("Done. %v", fmt.Sprintf("Elapsed Time: %v\n", stop))
}

// example url:
// https://qa.door43.org/api/v1/repos/unfoldingword/en_twl/releases
func getLanguages() [][]string {
	fullUrl := "https://td.unfoldingword.org/exports/langnames.json"
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// convert json to a map
	var jsonArray []map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonArray)
	if err != nil {
		panic(err)
	}

	var languages [][]string
	for _, lval := range jsonArray {
		var language []string
		gw := fmt.Sprintf("%v", lval["gw"])
		lc := fmt.Sprintf("%v", lval["lc"])
		ld := fmt.Sprintf("%v", lval["ld"])
		ang := fmt.Sprintf("%v", lval["ang"])
		language = append(language, gw, lc, ld, ang)
		//tags = append(tags, fmt.Sprintf("%v", val["tag_name"]))
		languages = append(languages, language)
	}

	return languages
}

/* Sample data

[
  {
    "cc": [
      "DJ",
      "ER",
      "ET",
      "US",
      "CA"
    ],
    "hc": "ET",
    "pk": 6,
    "alt": [
      "Afaraf",
      "Danakil",
      "Denkel",
      "Adal",
      "Afar Af",
      "Qafar",
      "Baadu (Ba'adu)"
    ],
    "lc": "aa",
    "lr": "Africa",
    "ld": "ltr",
    "ang": "Afar",
    "gw": false,
    "ln": "Afaraf"
  },
  {
    "cc": [
      "NG"
    ],
    "hc": "NG",
    "pk": 7,
    "alt": [

    ],
    "lc": "aaa",
    "lr": "Africa",
    "ld": "ltr",
    "ang": "Ghotuo",
    "gw": false,
    "ln": "Ghotuo"
  },

*/
