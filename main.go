package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {

	argument := os.Args[1]

	if len(argument) < 2 {
		log.Fatalf("Use: %s filename.txt", os.Args[0])
		return
	}

	jsonData := make(map[string]string)

	parsedUrl, err := url.ParseQuery(argument)

	if err != nil {
		log.Fatalln(err)
	}

	for key, values := range parsedUrl {
		if len(values) > 0 {
			jsonData[key] = values[0]
		}
	}

	jsonByte, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(jsonByte))
}
