package main

import (
	_ "expvar"
	"log"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	worker := NewWorker()

	srv := NewServer(worker)

	go runInBackground()
	log.Fatal(srv.Start())
}
