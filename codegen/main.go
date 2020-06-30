package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := flag.String(
		"dir",
		wd,
		"specify directory to process",
	)
	flag.Parse()
	err = Generate(*dir)
	if err != nil {
		log.Fatal(err)
	}
}
