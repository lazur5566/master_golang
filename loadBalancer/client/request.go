package client

import (
	"fmt"
	"math/rand"
	"time"
)

type Request struct {
	Fn func() int // The operation to perform.
	C  chan int   // The channel to return the result.
}

func Requester(work chan<- Request) {
	c := make(chan int)
	for {
		// Kill some time (fake load).
		//time.Sleep(rand.Int63n(config.NWorker*2) * time.Millisecond)
		time.Sleep(time.Second)
		work <- Request{workFn, c} // send request
		fmt.Println("Client: request sent...")
		result := <-c // wait for answer
		fmt.Println("Client: result received...")
		furtherProcess(result)
	}
}

func workFn() int {
	fmt.Println("***************")
	fmt.Println("dispatching work...")
	fmt.Println("***************")
	return rand.Intn(10)
}

func furtherProcess(result int) {
	fmt.Println("===============")
	fmt.Println("got number: ", result)
	fmt.Println("===============")
}
