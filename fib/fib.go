package main

import (
	"fmt"
	"time"
)

func main()  {
	start := time.Now()
	fibG(44)
	fmt.Println("斐波那契递归45: ", time.Since(start))
	start1 := time.Now()
	 fib(10000)
	fmt.Println("斐波那契循环45: ", time.Since(start1))
	// fmt.Printf("%d\n", result1)
}

func fib(n int) int{
	x,y := 0,1
	for i := 0; i<n;i++{
		x, y = y,x+y
	}
	return x
}

func fibG(n int) int {
	if n<2 {
		return n
	}
	return fibG(n-1) + fibG(n-2);
}