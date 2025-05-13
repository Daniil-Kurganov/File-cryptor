package main

import (
	"file_crypter/src"
	"log"
)

func main() {
	log.SetFlags(0)
	src.HTTPServer()
}
