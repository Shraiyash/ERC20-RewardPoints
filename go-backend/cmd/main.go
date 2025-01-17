package main

import (
	"log"
	"net/http"
	"cashback-rewards/pkg"
	"cashback-rewards/config"
)

func main() {
	cfg := config.LoadConfig()

	client, err := pkg.ConnectToEthereum(cfg.EthereumNodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum: %v", err)
	}

	router := pkg.InitRoutes(client)

	log.Printf("Starting server on port %s...", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}