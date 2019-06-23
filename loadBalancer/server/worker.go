package worker

type Worker struct {
	Requests chan Request // work to do (buffered channel)
	Pending  int          // count of pending tasks
	Index    int          // index in the heap
}

func (w *Worker) Work(done chan *Worker) {
	for {
		req := <-w.Requests // get Request from balancer
		req.C <- req.Fn()   // call fn and send result
		done <- w           // we've finished this request
	}
}
