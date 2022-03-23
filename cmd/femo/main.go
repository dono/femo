package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/atotto/clipboard"
	"github.com/ktr0731/go-fuzzyfinder"
)

type Entity struct {
	Title string `json:"title"`
	ID    string `json:"id"`
	Value string `json:"value"`
}

func main() {
	var (
		src = flag.String("src", "", "source json file")
	)
	flag.Parse()

	if *src == "" {
		log.Fatal("source file does not exist")
	}

	b, err := ioutil.ReadFile(*src)
	if err != nil {
		log.Fatal(err)
	}

	var entities []Entity
	if err = json.Unmarshal(b, &entities); err != nil {
		log.Fatal(err)
	}

	idx, err := fuzzyfinder.Find(
		entities,
		func(i int) string {
			return entities[i].Title
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("%s\n\nID: %s\nValue: %s\n",
				entities[i].Title,
				entities[i].ID,
				entities[i].Value,
			)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(entities[idx].Value)

	err = clipboard.WriteAll(entities[idx].Value)
	if err != nil {
		log.Fatal(err)
	}
}
