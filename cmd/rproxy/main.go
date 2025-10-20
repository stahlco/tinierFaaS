package main

import "log"

func main() {
	log.SetPrefix("rproxy: ")
	log.SetFlags(log.Lshortfile | log.Ltime)

	log.Printf("rproxy is starting...")
}
