package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/rishavmehra/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "Enter the Json file name, With some Stories")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)
}
