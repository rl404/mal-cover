package main

import "log"

func main() {
	// Read and parse config.
	if err := setConfig(); err != nil {
		log.Fatal("failed reading config", err)
		return
	}

	// Start HTTP serve.
	if err := startHTTP(); err != nil {
		log.Fatal("failed starting HTTP", err)
	}
}
