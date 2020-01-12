package service

import (
	"github.com/koolalex/mdblog/library/cache"
	"time"
)

var (
	Cache *cache.Cache
)

func init() {
	Cache = cache.NewFrom(cache.DefaultExpiration, 5*time.Minute, make(map[string]cache.Item, 5000))
}
