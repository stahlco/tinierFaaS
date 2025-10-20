package main

import (
	"log"
	"tinierFaaS/pkg/manager"
)

const PORT = ":8001"

type server struct {
	ms *manager.ManagerService
}

func main() {
	log.SetPrefix("manager: ")
	log.SetFlags(log.Ltime | log.Lshortfile)

	log.Printf("starting tinierFaaS-Manager...")

	// Execute rproxy.bin
}
