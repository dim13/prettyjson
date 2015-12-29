package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func main() {
	var in, out bytes.Buffer

	_, err := in.ReadFrom(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Indent(&out, in.Bytes(), "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	out.WriteTo(os.Stdout)
}
