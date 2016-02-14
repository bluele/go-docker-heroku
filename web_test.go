package main

import (
	"os"
	"testing"
)

func TestRedisConnection(t *testing.T) {
	pool, err := newRedisPool(os.Getenv("REDIS_URL"))
	if err != nil {
		t.Error(err)
	}
	conn := pool.Get()
	_, err = conn.Do("set", "test-key", "test-value")
	if err != nil {
		t.Error(err)
	}
}
