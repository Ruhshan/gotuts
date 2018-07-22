package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titeles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>mnews>Keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	var s SitemapIndex
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)

	for _, Location := range s.Locations{
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
	}

}
