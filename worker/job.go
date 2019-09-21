package worker

import (
	"fmt"
	"time"
)

var (
	MaxWorker = 100 * 100 * 20
	//MaxQueue  = "1000"
)

type Action struct {
	Name string
	Age  int
}

type Job struct {
	AC Action
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				fmt.Println("age", job.AC.Age)
				time.Sleep(time.Second * 1)

			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
