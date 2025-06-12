package main

import (
	"log"
	"net/http"

	"github.com/Ik-cyber/caching-proxy/config"
	"github.com/Ik-cyber/caching-proxy/proxy"
	"github.com/Ik-cyber/caching-proxy/utils"
)

func main() {
	utils.InitLogger()
	cfg := config.LoadConfig("config.yaml")
	p := proxy.NewProxy(cfg)

	http.HandleFunc("/", p.HandleRequest)
	log.Println("Proxy server running on :8080")
	http.ListenAndServe(":8080", nil)
}
