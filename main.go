package main

import (
	"etcd-demo/config"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello")
	config.Load()

	http.ListenAndServe(":8090", nil)
}
