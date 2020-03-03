package main

import (
	"flag"
	"fmt"
	"github.com/sathishkumar-manogaran/GoLangPrograms/story-reading/handlers"
	"github.com/sathishkumar-manogaran/GoLangPrograms/story-reading/models"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "Port Number to start the application")
	fileName := flag.String("file", "gopher.json", "JSON file with stories")
	flag.Parse()
	fmt.Printf("File Name %s\n", *fileName)

	file, err := os.Open(*fileName)
	errorFunc(err)

	story, err := models.JsonStory(file)
	errorFunc(err)

	//fmt.Printf("%+v\n", story)

	handler := handlers.NewHandler(story)
	fmt.Printf("Starting application on port :%d\n", *port)
	log.Fatal(http.ListenAndServe("3000", handler))
}

func errorFunc(err error) {
	if err != nil {
		panic(err)
	}
}
