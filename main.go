package main

import (
	"os"
	"io"
	"log"
	"errors"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(errors.New("Provide path to the file"))
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
	err = os.Remove(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
