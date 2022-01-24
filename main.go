package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

var redisclient *redis.Client

func main() {
	redishost := os.Getenv("REDISHOST")
	if redishost == "" {
		redishost = "localhost:6379"
	}
	redisclient = redis.NewClient(&redis.Options{
		Addr:     redishost,
		Password: "",
		DB:       0,
	})
	err := redisclient.Set("name", "Elliot", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	name, err := redisclient.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Hello from this simple server, the name stored is %s", name)
}
