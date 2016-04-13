package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"os"
)

var nspaces = flag.Int("spaces", 4, "indent")

func spaces(n int) string {
	z := make([]rune, n)
	for i := range z {
		z[i] = ' '
	}
	return string(z)
}

func main() {
	flag.Parse()
	indent := spaces(*nspaces)

	in := bytes.Buffer{}
	if _, err := in.ReadFrom(os.Stdin); err != nil {
		log.Fatal(err)
	}

	out := bytes.Buffer{}
	if err := json.Indent(&out, in.Bytes(), "", indent); err != nil {
		log.Fatal(err)
	}

	if _, err := out.WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}
}