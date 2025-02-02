package utils

import (
	"log"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemcachedUtil interface {
	GetClient() *memcache.Client
	Close()
}

type MemcachedUtilImplementation struct {
	Client *memcache.Client
}

func NewMemcachedUtil() MemcachedUtil {
	println(time.Now().String(), "memcached: connecting")
	client := memcache.New("localhost:11211")
	println(time.Now().String(), "memcached: connected")

	println(time.Now().String(), "memcached: pinging")
	err := client.Ping()
	if err != nil {
		log.Fatalln("error when pinging:", err)
	}
	println(time.Now().String(), "memcached: pinged")
	return &MemcachedUtilImplementation{
		Client: client,
	}
}

func (util *MemcachedUtilImplementation) GetClient() *memcache.Client {
	return util.Client
}

func (util *MemcachedUtilImplementation) Close() {
	println(time.Now().String(), "memcached: closing")
	err := util.Client.Close()
	if err != nil {
		log.Fatalln("error when closing memcached:", err)
	}
	println(time.Now().String(), "memcached: closed")
}
