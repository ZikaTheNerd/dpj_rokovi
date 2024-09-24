package main

import (
	"fmt"
	"runtime"
	"sync"
)

func transformNumber(n, scalar int) int {
	if n < 10 {
		return n
	}

	pow, result := 1, 0

	for n > 0 {
		digit := n % 10
		if digit%scalar != 0 {
			result = digit*pow + result
			pow *= 10
		}
		n /= 10
	}

	return result
}

func transform(arr []int, start, length, scalar int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i < start+length; i++ {
		arr[i] = transformNumber(arr[i], scalar)
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	var scalar int
	fmt.Scanf("%d", &scalar)

	var wg sync.WaitGroup
	threads := runtime.NumCPU()
	wg.Add(threads)

	quotient := n / threads
	remainder := n % threads

	thread, start := 0, 0
	for ; thread < remainder; thread++ {
		go transform(arr, start, quotient+1, scalar, &wg)
		start += quotient + 1
	}

	for ; thread < threads; thread++ {
		go transform(arr, start, quotient, scalar, &wg)
		start += quotient
	}

	wg.Wait()

	fmt.Println(arr)

}
