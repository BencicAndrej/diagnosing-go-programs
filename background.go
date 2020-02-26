package main

import (
	"time"
)

var s [][]byte

func runInBackground() {
	for {
		s = append(s, make([]byte, 1024))
		time.Sleep(10 * time.Millisecond)
	}
}

func reserveBytes() {
	s = append(s, make([]byte, 1024))
}
