package main

import (
	"fmt"
	"runtime"
	"sync"
)

func sumOfDigits(n int) int {
	if n < 0 {
		n = -n
	}

	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

func calculateSmallest(vector []int, result, index *int, start, length int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i < start+length; i++ {
		sum := sumOfDigits(vector[i])
		if *result == 0 || *result > sum {
			*result = sum
			*index = i
		} else if *result == sum && *index < i {
			*index = i
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

	var wg sync.WaitGroup

	goroutines := runtime.NumCPU()
	quotient := n / goroutines
	remainder := n % goroutines
	routine, start := 0, 0

	wg.Add(goroutines)

	result, num := 0, -1
	for ; routine < remainder; routine++ {
		go calculateSmallest(vector, &result, &num, start, quotient+1, &wg)
		start += quotient + 1
	}
	for ; routine < goroutines; routine++ {
		go calculateSmallest(vector, &result, &num, start, quotient, &wg)
		start += quotient
	}

	wg.Wait()

	fmt.Println(vector[num])
}
