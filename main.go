package main

import (
	"timewheel/wheel"
	"timewheel/worker"
)

func main() {
	dispatcher := worker.NewDispatcher(worker.MaxWorker)
	dispatcher.Run()
	worker.JobQueue = make(chan worker.Job)
	wheel.Exec()
}
