package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type slideshow struct {
	Author string  `json:"author"`
	Date   string  `json:"date"`
	Slides []slide `json:"slides"`
	Title  string  `json:"title"`
}

type slide struct {
	Items []string `json:"items"`
	Title string   `json:"title"`
	Type  string   `json:"type"`
}

func main() {
	resp, err := http.Get("https://httpbin.org/json")
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}
	var s map[string]slideshow
	err = json.Unmarshal(body, &s)
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}
	a := s["slideshow"]

	fmt.Println("======== JSON Recieved: ========")
	fmt.Print(string(body))

	fmt.Println("======== s[\"SlideShow\"] : ========")
	fmt.Println(s["slideshow"])

	fmt.Println("======== Pretty Print of structures : ========")
	fmt.Println("Slideshow author = " + a.Author)
	fmt.Println("Slideshow date = " + a.Date)
	fmt.Println("Slideshow title = " + a.Title)
	printSlides(&a)
}

func printSlides(s *slideshow) {
	for i, sl := range s.Slides {
		fmt.Printf("%vth slide : \n", i)
		fmt.Println("Title :" + sl.Title)
		fmt.Println("Things:")
		for i, t := range sl.Items {
			fmt.Printf("thing %v: ", i)
			fmt.Println(t)
		}
		fmt.Println("Slide type :" + sl.Type)
	}
}
