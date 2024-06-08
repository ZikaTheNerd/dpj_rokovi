package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func threeLargestSingle(n int) []int {
	if n < 0 {
		n = -n
	}
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	sort.Sort(sort.Reverse(sort.IntSlice(digits)))
	return digits[0:(min(3, len(digits)))]
}

func formNumber(digits []int) int {
	return digits[0]*100 + digits[1]*10 + digits[2]
}

func calcThreeLargest(vector []int, result *[]int, start, length int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i < start+length; i++ {
		digits := threeLargestSingle(vector[i])
		*result = append(*result, digits...)
		sort.Sort(sort.Reverse(sort.IntSlice(*result)))
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	//fmt.Println(merge([]int{0, 0, 0}, []int{9, 8, 4}))

	vector := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &vector[i])
	}

	goroutines := runtime.NumCPU()
	quotient := n / goroutines
	remainder := n % goroutines

	routine, start := 0, 0
	result := []int{}
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for ; routine < remainder; routine++ {
		go calcThreeLargest(vector, &result, start, quotient+1, &wg)
		start += quotient + 1
	}
	for ; routine < goroutines; routine++ {
		go calcThreeLargest(vector, &result, start, quotient, &wg)
		start += quotient
	}

	wg.Wait()

	if len(result) < 3 {
		fmt.Println(-1)
	}

	fmt.Println(formNumber(result))

	fmt.Println(formNumber(result))

}
