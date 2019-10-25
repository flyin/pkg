package env

import (
	"log"
	"net/url"
	"os"
)

func GetBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "yes" || v == "on" {
		return true
	}
	return fallback
}

func Get(key string, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func MustGet(key string) string {
	v := Get(key, "")
	if v == "" {
		log.Fatalf("You must define '%s' environment variable", key)
	}
	return v
}

func GetEnvURL(key string, fallback string) *url.URL {
	u, err := url.Parse(Get(key, fallback))
	if err != nil {
		log.Fatalf("Bad url in '%s' environment variable", key)
	}
	return u
}
