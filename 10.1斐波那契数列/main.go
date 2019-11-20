package main

import "fmt"

func main() {
	fmt.Println(fib(6))
	fmt.Println(fib1(6))
	fmt.Println(fib2(6))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func fib1(n int) int {
	if n <= 1 {
		return n
	}
	fib := make([]int, n+1)
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func fib2(n int) int {
	if n <= 1 {
		return n
	}
	pre2, pre1, fib := 0, 1, 0
	for i := 2; i <= n; i++ {
		fib = pre1 + pre2
		pre2 = pre1
		pre1 = fib
	}

	return fib
}
