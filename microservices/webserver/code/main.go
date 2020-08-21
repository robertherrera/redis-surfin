package main

import (
        "fmt"
        "net/http"
        "log"
    
    	"github.com/go-redis/redis"
    )
    
    func main() {
        fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
        })
    
        http.ListenAndServe(":80", nil)
    }