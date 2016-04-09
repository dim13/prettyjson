package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func main() {
	var in, out bytes.Buffer

	if _, err := in.ReadFrom(os.Stdin); err != nil {
		log.Fatal(err)
	}

	if err := json.Indent(&out, in.Bytes(), "", "\t"); err != nil {
		log.Fatal(err)
	}

	if _, err := out.WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
