package main

import (
	"github.com/piprate/json-gold/ld"
	"log"
)

func main() {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	expanded, err := proc.Expand("https://raw.githubusercontent.com/c2bo/didcomm-test-context/main/test.jsonld", options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)
		return
	}
	ld.PrintDocument("JSON-LD expansion succeeded", expanded)

}
