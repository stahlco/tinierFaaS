package main

import (
	"github.com/google/uuid"
	"log"
)

func main() {
	log.SetPrefix("manager: ")
	log.SetFlags(log.Lshortfile | log.Ltime)

	log.Printf("Manager started")

	//Creating Backend

	id := uuid.New().String()

	var backend Backend
	//Creating Reverse Proxy
}
