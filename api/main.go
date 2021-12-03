package main

import (
	"encoding/json"
	"flag"
	"github.com/richmondang/custom-tf/api/server"
	"io/ioutil"
	"log"
)

func main() {

	//Run Server
	// if err := server.runServer("localhost:8080"); err != nil {
	// 	log.Println(err)
	// }

	seed := flag.String("seed", "", "a file location with some data in JSON form to seed the server content")
	flag.Parse()

	volumes := map[string]server.Volume{}

	if *seed != "" {
		seedData, err := ioutil.ReadFile(*seed)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(seedData, &volumes)
		if err != nil {
			log.Fatal(err)
		}
	}

	volumeService := server.NewService("localhost:3001", volumes)
	err := volumeService.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
