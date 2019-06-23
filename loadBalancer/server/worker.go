package server

import "github.com/afon/MasterGolang/loadBalancer/client"

type Worker struct {
	requests chan client.Request // work to do (buffered channel)
	pending  int                 // count of pending tasks
	index    int                 // index in the heap
}

func (w *Worker) Work(done chan *Worker) {
	for {
		req := <-w.requests // get Request from balancer
		req.C <- req.Fn()   // call fn and send result
		done <- w           // we've finished this request
	}
}
