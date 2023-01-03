package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	indent := flag.Int("n", 2, "indent level")
	compact := flag.Bool("c", false, "compact output")
	flag.Parse()

	dec, enc := json.NewDecoder(os.Stdin), json.NewEncoder(os.Stdout)
	if !*compact {
		spaces := strings.Repeat(" ", *indent)
		enc.SetIndent("", spaces)
	}

	var v any
	if err := dec.Decode(&v); err != nil {
		log.Fatal(err)
	}
	if err := enc.Encode(&v); err != nil {
		log.Fatal(err)
	}
}
