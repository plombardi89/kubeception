package main

import (
	"log"
	"os"
	"strings"
)

func isDisabled() bool {
	choices := make(map[string]bool)
	choices["1"] = true
	choices["yes"] = true
	choices["true"] = true

	if _, exists := choices[strings.ToLower(os.Getenv("KUBECEPTION_BOOTSTRAP_DISABLE"))]; exists {
		return true
	}

	return false
}

func main() {
	if isDisabled() {
		log.Println("cluster bootstrap disabled, exiting.")
		return
	}

	log.Println("kubernetes bootstrap started")

	// check if we are bootstrapping a master or worker

	// kubeadm init

	// install cluster networking

	// perform smoke tests and sanity

	// publish the kubeconfig into a secret

	log.Println("kubernetes bootstrap finished")
}
