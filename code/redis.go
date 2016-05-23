package main

import (
	"gopkg.in/redis.v3"
	"log"
)

func redissend(id string, httpstate string, httpcode string, statecall string, state string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	args := []string{statecall, state}

	err := client.HMSet(id, httpstate, httpcode, args...).Err()
	if err != nil {
		panic(err)
	}

	log.Print(" [x] redis Sent ")
}
