package main

import (
	"fmt"
	"runtime"
	"sync"
)

func isPerfect(n int) bool {
	sum := 0
	for num := 1; num < n; num++ {
		if n%num == 0 {
			sum += num
		}
	}

	return sum == n
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculateMinimalPerfect(vector []int, start, length int, result *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i < start+length; i++ {
		if isPerfect(vector[i]) {
			if *result == 0 {
				*result = vector[i]
			} else {
				*result = minimum(*result, vector[i])
			}
		}
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	vector := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &vector[i])
	}

	goroutines := runtime.NumCPU()
	quotient := n / goroutines
	remainder := n % goroutines

	var wg sync.WaitGroup
	wg.Add(goroutines)

	start, routine := 0, 0
	minimalPerfect := 0
	for ; routine < remainder; routine++ {
		go calculateMinimalPerfect(vector, start, quotient+1, &minimalPerfect, &wg)
		start += quotient + 1
	}

	for ; routine < goroutines; routine++ {
		go calculateMinimalPerfect(vector, start, quotient, &minimalPerfect, &wg)
		start += quotient
	}

	wg.Wait()
	fmt.Println(minimalPerfect)
}
