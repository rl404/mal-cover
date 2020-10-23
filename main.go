package main

import "log"

func main() {
	// Read and parse config.
	err := GetConfig()
	if err != nil {
		log.Fatal("failed reading config", err)
		return
	}

	// Start HTTP serve.
	err = StartHTTP()
	if err != nil {
		log.Fatal("failed starting HTTP", err)
	}
}