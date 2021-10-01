package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func SearchForVideoLinks(url string) {

	log.Println("Parsing : ", url)

	// Request the HTML page.
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Unable to get URL with status code error: %d %s", resp.StatusCode, resp.Status)
	}

	htmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	videoRegExp := regexp.MustCompile(`<video[^>]+`)
	sourceRegExp := regexp.MustCompile(`<source[^>]+`)

	videoMatchSlice := videoRegExp.FindAllStringSubmatch(string(htmlData), -1)
	sourceMatchSlice := sourceRegExp.FindAllStringSubmatch(string(htmlData), -1)

	for _, item := range videoMatchSlice {
		log.Println("Video found : ", item)
		for _, sourceItem := range sourceMatchSlice {
			log.Println("Source found : ", sourceItem)
		}
	}

}

func main() {

	SearchForVideoLinks("https://cdpn.io/caraya/fullpage/FckCd")
}
