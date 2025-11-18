package config

import (
	"fmt"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

func NewCache() *memcache.Client {
	url := fmt.Sprintf("%s:%s", os.Getenv("CACHE_HOST"), os.Getenv("CACHE_PORT"))
	return memcache.New(url)
}
