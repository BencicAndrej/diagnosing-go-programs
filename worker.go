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
	w.tasks["allocmem"] = w.runCPU
	w.tasks["log"] = w.runLog

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
	if rand.Int()%10 == 0 {
		SumFirst(10000000000)
	}

	return nil
}

func (w *Worker) runMemory() error {
	reserveBytes()

	return nil
}

func (w *Worker) runAllocatedMemory() error {
	var b []byte
	for i := 0; i < 100; i++ {
		b = make([]byte, 1024*1024) // 1mb
	}
	if b == nil {
		panic("bytes should never be nil")
	}

	return nil
}

func (w *Worker) runLock() error {
	w.lock.Lock()
	defer w.lock.Unlock()

	time.Sleep(10 * time.Millisecond)

	return nil
}
func (w *Worker) runLog() error {
	log.Print("Starting logging work")
	SumFirstAlt(10000000000)
	log.Print("Finished logging work")
	return nil
}
