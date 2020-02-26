package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	worker *Worker
	start  time.Time
}

func NewServer(worker *Worker) *Server {
	return &Server{
		worker: worker,
	}
}

func (srv *Server) Start() error {
	srv.start = time.Now()

	http.HandleFunc("/work/", srv.handleWork)
	http.HandleFunc("/", srv.handleCounter)

	log.Print("Listening on port 4000...")
	return http.ListenAndServe(":4000", nil)
}

func (srv *Server) handleCounter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time since last accident: %.0fs\n", time.Since(srv.start).Seconds())
}

func (srv *Server) handleWork(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/work/")
	if id == "" {
		fmt.Fprintf(w, "Select work to do: %s",
			strings.Join(srv.worker.GetIDs(), ", "))
		return
	}

	err := srv.worker.Run(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	fmt.Fprintln(w, "OK")
}
