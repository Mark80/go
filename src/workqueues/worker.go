package workqueues

import (
  "time"
  "fmt"
)

type WorkRequest struct {
  Name  string
  Delay time.Duration
}

type Worker struct {
  ID        int
  Work      chan WorkRequest
  WorkQueue chan chan WorkRequest
  QuitChan  chan bool
}

func NewWorker(id int, workQueue chan chan WorkRequest) Worker {

  return Worker{
	ID : id,
	Work: make(chan WorkRequest),
	WorkQueue:workQueue,
	QuitChan:make(chan bool),
  }

}

func (worker *Worker) Start() {

  go func() {
	for {

	  worker.WorkQueue <- worker.Work

	  select {
	  case work := <-worker.Work:

		fmt.Printf("worker%d : received work request delay for %s", worker.ID, work.Delay)
		time.Sleep(work.Delay)
		fmt.Printf("worker%d : Hello , %s\n", worker.ID, work.Name)

	  case  <-worker.QuitChan:
        fmt.Printf("stopping worker %d " , worker.ID)
		return

	  }

	}

  }()

}

func (worker *Worker) Stop() {
  go func(){
	worker.QuitChan <- true
  }()

}