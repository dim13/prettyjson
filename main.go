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

	if err := json.Indent(&out, in.Bytes(), "", "  "); err != nil {
		log.Fatal(err)
	}

	out.WriteTo(os.Stdout)
}
