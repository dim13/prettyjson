package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	n := flag.Int("n", 2, "indent spaces")
	c := flag.Bool("c", false, "compact output")
	flag.Parse()

	in := new(bytes.Buffer)
	_, err := in.ReadFrom(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	out := new(bytes.Buffer)
	if *c {
		err = json.Compact(out, in.Bytes())
	} else {
		err = json.Indent(out, in.Bytes(), "", strings.Repeat(" ", *n))
	}
	if err != nil {
		log.Fatal(err)
	}

	_, err = out.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
