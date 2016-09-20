package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"os"
)

func spaces(n int) string {
	z := make([]rune, n)
	for i := range z {
		z[i] = ' '
	}
	return string(z)
}

func main() {
	var n = flag.Int("n", 2, "indent spaces")
	flag.Parse()
	indent := spaces(*n)

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
