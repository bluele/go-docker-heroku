package main

import (
	"fmt"
	"github.com/bluele/go-docker-heroku/Godeps/_workspace/src/github.com/garyburd/redigo/redis"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	pool, err := newRedisPool(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok?")
	})
	http.HandleFunc("/incr", func(w http.ResponseWriter, r *http.Request) {
		conn := pool.Get()
		defer conn.Close()
		v, err := redis.Int64(conn.Do("incr", "key"))
		if err != nil {
			log.Panic(err)
			// fmt.Fprint(w, err.Error())
		} else {
			fmt.Fprintf(w, "counter:%v", v)
		}
	})
	log.Println("start server.")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func newRedisPool(us string) (*redis.Pool, error) {
	u, err := url.Parse(us)
	if err != nil {
		return nil, err
	}

	var password string
	if u.User != nil {
		password, _ = u.User.Password()
	}
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", u.Host)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}
