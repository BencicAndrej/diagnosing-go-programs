package main

import (
	"testing"
	"time"
)

func BenchmarkSleepSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Second)
	}
}

const NUM = 1000000 // 1 Million

func TestSumsAreEqual(t *testing.T) {
	want := SumFirst(NUM)
	got := SumFirstAlt(NUM)
	if got != want {
		t.Errorf("Sums do not match: got: %d, want %d", got, want)
	}
}

func BenchmarkSumFirstMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumFirst(NUM)
	}
}

func BenchmarkSumFirstMillionAlt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumFirstAlt(NUM)
	}
}
