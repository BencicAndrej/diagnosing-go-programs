package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Worker struct {
	lock  sync.Mutex
	tasks map[string]func() error
}

func NewWorker() *Worker {
	w := &Worker{
		tasks: make(map[string]func() error),
	}

	w.tasks["cpu"] = w.runCPU
	w.tasks["mem"] = w.runMemory
	w.tasks["lock"] = w.runLock

	return w
}

func (w *Worker) Run(id string) error {
	task := w.tasks[id]

	return task()
}

func (w *Worker) GetIDs() []string {
	var ids []string
	for id := range w.tasks {
		ids = append(ids, id)
	}

	return ids
}

func (w *Worker) runCPU() error {
	log.Print("Starting work")
	defer log.Print("Finished work")

	if rand.Int()%10 == 0 {
		SumFirst(1000000000)
	}

	return nil
}

func (w *Worker) runMemory() error {
	reserveBytes()

	return nil
}

func (w *Worker) runLock() error {
	w.lock.Lock()
	defer w.lock.Unlock()

	time.Sleep(10 * time.Millisecond)

	return nil
}
