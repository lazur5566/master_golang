package main

import (
	"container/heap"
	"fmt"

	"loadBalancer/client"
	"loadBalancer/config"
	"loadBalancer/server"
)

func main() {
	fmt.Println("vim-go")
	var b server.Balancer
	var r client.Request

	heap.Push(&b.pool, &server.Worker{})
	fmt.Println(config.Test())

}
