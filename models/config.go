package models

type Config struct {
	Routes []Route `yaml:"routes"`
}

type Route struct {
	Domain string `yaml:"domain"`
	TTL    int    `yaml:"ttl"`
}
