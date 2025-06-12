package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Ik-cyber/caching-proxy/config"
	"github.com/Ik-cyber/caching-proxy/proxy"
	"github.com/Ik-cyber/caching-proxy/utils"
)

func main() {
	// Initialize logger
	utils.InitLogger()

	// Parse the config path flag
	configPath := flag.String("config", "config/config.yaml", "Path to the configuration file")
	flag.Parse()

	// Load configuration
	cfg := config.LoadConfig(*configPath)

	// Initialize proxy (cache is created inside)
	p := proxy.NewProxy(cfg)

	// Start proxy server
	log.Printf("Proxy server running on :8080 using config: %s", *configPath)
	if err := http.ListenAndServe(":8080", http.HandlerFunc(p.HandleRequest)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
