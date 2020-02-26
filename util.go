package main

func SumFirst(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func SumFirstAlt(n int) int {
	return n * (n + 1) / 2
}
