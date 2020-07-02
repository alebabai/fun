package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	if flag.NArg() == 1 {
		dir = flag.Arg(0)
	} else if flag.NArg() > 1 {
		log.Fatal("too many arguments")
	}

	err = Generate(dir)
	if err != nil {
		log.Fatal(err)
	}
}
