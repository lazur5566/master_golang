package client

import (
	"fmt"
	"math/rand"

	config "/Users/Yifeng/MaterGolang/loadBalancer/config"
)

type Request struct {
	Fn func() int // The operation to perform.
	C  chan int   // The channel to return the result.
}

func Requester(work chan<- Request) {
	c := make(chan int)
	for {
		// Kill some time (fake load).
		Sleep(rand.Int63n(config.Test() * 2 * Second))
		work <- Request{workFn, c} // send request
		result := <-c              // wait for answer
		furtherProcess(result)
	}
}

func workFn() int {
	fmt.Println("***************")
	fmt.Println("dispatching work...")
	fmt.Println("***************")
	return rand.Int(10)
}

func furtherProcess(result int) {
	fmt.Println("===============")
	fmt.Println("got number: ", result)
	fmt.Println("===============")
}
