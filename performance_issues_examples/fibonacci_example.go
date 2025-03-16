package performance_examples

import (
	"fmt"
	"time"
)

func RunFibonacciExample() {
	fmt.Println("计算斐波那契数列...")

	n := 40

	// 测试递归方法
	start := time.Now()
	resultRecursive := fibonacciRecursive(n)
	durationRecursive := time.Since(start)
	fmt.Printf("递归方法 (n=%d): 结果=%d, 耗时=%v\n", n, resultRecursive, durationRecursive)

	// 测试迭代方法
	start = time.Now()
	resultIterative := fibonacciIterative(n)
	durationIterative := time.Since(start)
	fmt.Printf("迭代方法 (n=%d): 结果=%d, 耗时=%v\n", n, resultIterative, durationIterative)

	// 比较性能差异
	fmt.Printf("性能比较: 递归/迭代 = %.2f倍\n", float64(durationRecursive)/float64(durationIterative))
}

// 递归实现
func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

// 迭代实现
func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
