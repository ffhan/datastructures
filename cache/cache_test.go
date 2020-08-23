package cache

import (
	"algos/lfu"
	"algos/lru"
	"math/rand"
	"testing"
)

type Cache interface {
	Get(key interface{}) interface{}
	Put(key, value interface{})
}

func BenchmarkLru(b *testing.B) {
	s := rand.NewSource(12345)
	r := rand.New(s)
	for i := 0; i < b.N; i++ {
		lru := lru.NewLru(100)
		for j := 0; j < 200; j++ {
			k := r.Int()
			v := r.Int()
			lru.Put(k, v)
			for z := 0; z < 50; z++ {
				index := r.Int()
				lru.Get(index)
			}
		}
	}
}

func BenchmarkLfu(b *testing.B) {
	s := rand.NewSource(12345)
	r := rand.New(s)
	for i := 0; i < b.N; i++ {
		lfu := lfu.NewLfu(100)
		for j := 0; j < 200; j++ {
			k := r.Int()
			v := r.Int()
			lfu.Put(k, v)
			for z := 0; z < 50; z++ {
				index := r.Int()
				lfu.Get(index)
			}
		}
	}
}
