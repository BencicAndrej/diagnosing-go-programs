package main

import (
	_ "expvar"
	"log"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	runtime.SetBlockProfileRate(1000)
	runtime.SetMutexProfileFraction(1000)

	worker := NewWorker()

	srv := NewServer(worker)

	go runInBackground()
	log.Fatal(srv.Start())
}
