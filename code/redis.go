package main

import (
  "gopkg.in/redis.v3"
  "fmt"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    //pong, err := client.Ping().Result()
    //fmt.Println(pong, err)
    // Output: PONG <nil>
//}

//func ExampleClient() {
    args := []string{"a","c","d","b","e","f"}
    //args2 := []string{"A","C","D","B","E","F"}


    //err := client.Set("key", "value", 0).Err()
    err := client.HMSet("new", "g", "h",args...).Err()
    if err != nil {
        panic(err)
    }

    val1, err := client.HGet("new","d").Result()
    if err != nil {
      panic(err)
    }
    fmt.Println("d",val1)

    val, err := client.Get("key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := client.Get("key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exists")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
    // Output: key value
    // key2 does not exists
}