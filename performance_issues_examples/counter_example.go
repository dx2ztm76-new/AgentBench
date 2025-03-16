package performance_examples

import (
	"fmt"
	"sync"
	"time"
)

// 共享资源
type Counter struct {
	value int
}

// 全局锁和计数器
var mutex sync.Mutex
var counter Counter

func RunCounterExample() {
	// 测试不同的并发方法
	fmt.Println("测试不同的并发方法...")

	// 方法1: 使用互斥锁
	start := time.Now()
	testMethod1(1000000)
	duration1 := time.Since(start)
	fmt.Printf("方法1 (互斥锁): %v\n", duration1)

	// 重置计数器
	counter.value = 0

	// 方法2: 使用分片计数器
	start = time.Now()
	testMethod2(1000000)
	duration2 := time.Since(start)
	fmt.Printf("方法2 (分片计数器): %v\n", duration2)

	// 方法3: 使用通道
	start = time.Now()
	testMethod3(1000000)
	duration3 := time.Since(start)
	fmt.Printf("方法3 (基于通道): %v\n", duration3)
}

// 方法1: 使用互斥锁
func testMethod1(iterations int) {
	var wg sync.WaitGroup

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < iterations/10; j++ {
				// 获取锁
				mutex.Lock()
				counter.value++
				// 模拟一些处理时间
				time.Sleep(time.Nanosecond)
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", counter.value)
}

// 方法2: 使用分片计数器
func testMethod2(iterations int) {
	var wg sync.WaitGroup
	var counters [10]Counter
	var mutexes [10]sync.Mutex

	// 启动10个goroutine，每个使用自己的计数器
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < iterations/10; j++ {
				// 锁定计数器
				mutexes[id].Lock()
				counters[id].value++
				time.Sleep(time.Nanosecond)
				mutexes[id].Unlock()
			}
		}(i)
	}

	wg.Wait()

	// 合并结果
	total := 0
	for i := 0; i < 10; i++ {
		total += counters[i].value
	}
	fmt.Printf("最终计数器值: %d\n", total)
}

// 方法3: 基于通道的方法
func testMethod3(iterations int) {
	// 创建一个通道来接收增量
	increments := make(chan int, iterations)
	done := make(chan bool)

	// 启动一个goroutine来处理所有增量
	go func() {
		total := 0
		for inc := range increments {
			total += inc
		}
		fmt.Printf("最终计数器值: %d\n", total)
		done <- true
	}()

	var wg sync.WaitGroup

	// 启动10个goroutine发送增量
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < iterations/10; j++ {
				increments <- 1
				time.Sleep(time.Nanosecond)
			}
		}()
	}

	wg.Wait()
	close(increments)
	<-done
}
