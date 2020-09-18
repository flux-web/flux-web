package models

import (
	"github.com/bradfitz/gomemcache/memcache"
)

func memLogin() *memcache.Client {
	return memcache.New(GetEnv("MEMCACHE_URL", "memcache") + ":" + GetEnv("MEMCACHE_PORT", "11211"))
}

func MemPut(key, value string) {
	mc := memLogin()
	err := mc.Set(&memcache.Item{Key: key, Value: []byte(value)})
	if err != nil {
		l.Printf(err.Error())
	}
}

func MemGet(key string) string {
	value := ""
	mc := memLogin()
	item, err := mc.Get(key)
	if err != nil {
		l.Printf(err.Error())
	} else {
		value = string(item.Value)
	}
	return value
}
