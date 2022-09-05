package main

import (
	"flag"
	"log"

	"github.com/Dsmit05/fixturegen/fixture"
)

var path = flag.String("source", "", "Input Go source file")

func main() {
	flag.Parse()

	gen, err := fixture.NewGenerator(*path)
	if err != nil {
		log.Fatal(err)
	}

	err = gen.Start()
	if err != nil {
		log.Fatal(err)
	}
}
