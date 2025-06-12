package proxy

import (
	"bytes"
	"io"
	"net/http"

	"github.com/Ik-cyber/caching-proxy/models"
	"github.com/Ik-cyber/caching-proxy/utils"
)

type Proxy struct {
	config *models.Config
	cache  *Cache
	client *http.Client
}

func NewProxy(cfg *models.Config) *Proxy {
	return &Proxy{
		config: cfg,
		cache:  NewCache(cfg),
		client: &http.Client{},
	}
}

func (p *Proxy) HandleRequest(w http.ResponseWriter, r *http.Request) {
	targetDomain, ttl := p.getTargetDomain(r.Host)
	if targetDomain == "" {
		http.Error(w, "Domain not configured", http.StatusBadGateway)
		return
	}

	cacheKey := r.Method + r.URL.Path + "?" + r.URL.RawQuery

	// Optional: Only cache GET requests, or make it config-driven
	if r.Method == http.MethodGet {
		if cachedResponse, found := p.cache.Get(cacheKey); found {
			utils.Info("[CACHE HIT] %s %s", r.Method, r.URL.String())
			// You could also cache headers if you want, right now we're just caching the body
			w.Write(cachedResponse)
			return
		} else {
			utils.Info("[CACHE MISS] %s %s", r.Method, r.URL.String())
		}
	}

	// Clone the request properly
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
		// Must recreate the Body reader because it's consumed once
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// Build the new request to the target
	proxyReq, err := http.NewRequest(r.Method, "https://"+targetDomain+r.URL.Path+"?"+r.URL.RawQuery, bytes.NewBuffer(bodyBytes))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Copy headers
	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	// Send the request using our http.Client
	resp, err := p.client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Upstream request failed", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// Read response body
	respBody, _ := io.ReadAll(resp.Body)

	// Cache the response body for GET requests
	if r.Method == http.MethodGet {
		p.cache.Set(cacheKey, respBody, ttl)
	}

	// Write status code and body to client
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}

func (p *Proxy) getTargetDomain(host string) (string, int) {
	for _, route := range p.config.Routes {
		if route.Domain == host {
			return route.Domain, route.TTL
		}
	}
	return "", 0
}
