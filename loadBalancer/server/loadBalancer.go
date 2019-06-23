package server

import (
	"container/heap"
	"fmt"

	"github.com/afon/MasterGolang/loadBalancer/client"
	"github.com/afon/MasterGolang/loadBalancer/config"
)

type Pool []*Worker

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Len() int {
	return len(p)
}

func (p *Pool) Push(x interface{}) {
	n := len(*p)
	worker := x.(*Worker)
	worker.index = n
	*p = append(*p, worker)
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	worker := old[n-1]
	worker.index = -1 // for safety
	*p = old[0 : n-1]
	return worker
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) Balance(work chan client.Request) {
	for {
		select {
		case req := <-work: // received a Request...
			fmt.Println("Balancer: received a Request...")
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished ...
			fmt.Println("Balancer: a worker has finished...")
			b.completed(w) // ...so update its info
		}
	}
}

// Send Request to worker
func (b *Balancer) dispatch(req client.Request) {
	// Grab the least loaded worker...
	w := heap.Pop(&b.pool).(*Worker)
	// ...send it the task.
	w.requests <- req
	// One more in its work queue.
	w.pending++
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

// Job is complete; update heap
func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue.
	w.pending--
	// Remove it from heap.
	heap.Remove(&b.pool, w.index)
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

func NewBalancer(nWorker int) *Balancer {
	done := make(chan *Worker, nWorker)
	// create nWorker WOK channels
	b := &Balancer{make(Pool, 0, nWorker), done}
	for i := 0; i < nWorker; i++ {
		w := &Worker{requests: make(chan client.Request, config.NRequester)}
		heap.Push(&b.pool, w)
		go w.Work(b.done)
	}
	fmt.Println("New Balancer Ends...")
	fmt.Println("Pool's size is ", len(b.pool))
	return b
}
