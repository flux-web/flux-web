package models

import (
	"github.com/bradfitz/gomemcache/memcache"
)

func memLogin() *memcache.Client {
	l.Println("in memLogin")
	l.Println(GetEnv("MEMCACHE_URL", "memcache") + ":" + GetEnv("MEMCACHE_PORT", "11211"))
	return memcache.New(GetEnv("MEMCACHE_URL", "memcache") + ":" + GetEnv("MEMCACHE_PORT", "11211"))
}

func MemPut(key, value string) {
	l.Println("in MemPut")
	mc := memLogin()
	l.Println("in MemPut - got mem client")
	err := mc.Set(&memcache.Item{Key: key, Value: []byte(value)})
	if err != nil {
		l.Printf(err.Error())
	}
}

func MemGet(key string) string {
	l.Println("in MemGet")
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
