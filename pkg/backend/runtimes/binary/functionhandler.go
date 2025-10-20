package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	// What is a function handler -> is basically an http-server that runs an embedded shell-script
	// 1. Listens on Port 8000 -> container is on 8000 exposed
	port := ":8000"
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		// 2. Has a GET Endpoint for health checks
		case "GET":
			if req.URL.Path != "/health" {
				w.WriteHeader(http.StatusNotFound)
				log.Printf("received bad get request")
				return
			}
			w.WriteHeader(http.StatusOK)
			_, err := fmt.Fprint(w, "OK")
			if err != nil {
				log.Printf("error occured answering to health-check: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			log.Printf("reported healthy function handler")
			return
		// 3. Has a POST Endpoint for calling a function -> executed the fn.sh
		case "POST":
			data, err := io.ReadAll(req.Body)
			if err != nil {
				log.Printf("error occured when reading body from request with error: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// 4. Executing the shell script with given data
			cmd := exec.Command("./fn.sh")
			cmd.Stdin = bytes.NewReader(data)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("error occured executing combined command: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(out)
			if err != nil {
				log.Printf("error writing output of fn.sh to response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		default:
			log.Printf("received request with abnormal method, returning")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})

	log.Printf("Listening on port: %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("error occured listening on port: %s with error: %s", port, err)
	}
}
