package loadBalancer

import "math/rand"

type Request struct {
	fn func() int // The operation to perform.
	c  chan int   // The channel to return the result.
}

func requester(work chan<- Request) {
	c := make(chan int)
	for {
		// Kill some time (fake load).
		Sleep(rand.Int63n(nWorker * 2 * Second))
		work <- Request{workFn, c} // send request
		result := <-c              // wait for answer
		furtherProcess(result)
	}
}
