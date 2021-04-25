package main

import (
	"log"
	"os"
)

func errorLog(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
