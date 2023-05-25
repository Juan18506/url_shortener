package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Setting flags
	addr := flag.String("addr", ":8080", "HTTP network address")
	jsonFileName := flag.String("json", "test.json", "Name of the url JSON file")
	flag.Parse()

	// Open JSON file
	jsonFile, err := os.Open(fmt.Sprintf("./json/%s", *jsonFileName))
	if err != nil {
		log.Fatal(err)
	}

	// Read JSON and store the values in jsonData var
	var jsonData []PathUrl
	err = json.NewDecoder(jsonFile).Decode(&jsonData)
	if err != nil {
		log.Fatal(err)
	}

	// Build the jsonHandler using mux as the fallback
	jsonHandler := mapHandler(buildMap(jsonData), defaultMux())

	// Start server
	log.Printf("Starting server on port %s...", *addr)
	http.ListenAndServe(*addr, jsonHandler)
}
