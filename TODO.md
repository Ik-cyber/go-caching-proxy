# TODO - Caching Proxy Project

## ✅ Completed

- [x] Project folder structure setup
- [x] Configuration parsing using YAML
- [x] Flag parsing for dynamic config paths
- [x] Request proxying and cloning
- [x] Basic in-memory caching with TTL
- [x] Background cache cleanup routine
- [x] Logging system initialized

## 🔨 In Progress

- [ ] Cache hit/miss logging (detailed per request)
- [ ] Test multiple site support from config
- [ ] Improve cache key generation for multi-site
- [ ] Request cloning full validation (headers, body, query params)

## 🚀 Next Up

- [ ] Dynamic port flag support
- [ ] Graceful shutdown with signal handling
- [ ] Configurable cache cleanup interval via YAML
- [ ] Proper README documentation with usage examples
- [ ] Implement basic unit tests

## 💡 Ideas / Future Enhancements

- [ ] Cache statistics endpoint (e.g., total hits, misses, cleanup count)
- [ ] Admin dashboard for cache inspection
- [ ] Prometheus metrics integration
- [ ] Support for Redis-backed cache
- [ ] Multi-threaded request logging
