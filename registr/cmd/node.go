package main

import (
	"log"

	"para-bank/registr/internal/network"
)

func main() {
	// Ініціалізація вузла блокчейну
	network.StartNode()
	log.Println("Blockchain node started")
}
