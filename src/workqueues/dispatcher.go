package workqueues

import "fmt"

var WorkerQueue chan chan WorkRequest

func StartDispatcher(nworker int) {

  WorkerQueue = make(chan chan WorkRequest, nworker)

  for i := 0; i < nworker; i++ {
	worker := NewWorker(i, WorkerQueue)
	worker.Start()
  }

  go func() {
	for {

	  select {
	  case work := <-WorkQueue:
		go func() {
		  worker := <-WorkerQueue
		  fmt.Println("Dispatching work request")
		  worker <- work
		}()
	  }
	}
  }()

}