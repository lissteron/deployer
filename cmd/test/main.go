package main

import (
	"io/ioutil"
	"log"

	"github.com/lissteron/deployer/pkg/github"
)

func main() {

	log.SetFlags(log.Lshortfile)

	b, err := ioutil.ReadFile("git.json")
	if err != nil {
		log.Fatalln(err)
	}

	/*	if err := json.Unmarshal(b, &event); err != nil {
		log.Fatalln(err)
	}*/

	var event github.PushEvent
	if err := event.UnmarshalJSON(b); err != nil {
		log.Fatalln(err)
	}

	log.Println(event)

}
