package main

import (
	"fmt"

	"github.com/afon/MasterGolang/loadBalancer/client"
	"github.com/afon/MasterGolang/loadBalancer/config"
	"github.com/afon/MasterGolang/loadBalancer/server"
)

func main() {
	fmt.Println("vim-go")
	work := make(chan client.Request)
	b := server.NewBalancer(config.NWorker)
	for i := 0; i < config.NRequester; i++ {
		go client.Requester(work)
	}
	b.Balance(work)
	//time.Sleep(100 * time.Second)
}
