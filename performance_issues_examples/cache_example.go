package performance_examples

import (
	"fmt"
	"runtime"
	"time"
)

// 全局缓存
var cache = make(map[string][]byte)

func RunExample() {
	// 模拟添加数据到缓存
	fmt.Println("开始缓存示例...")

	// 打印初始内存统计
	printMemStats("初始状态")

	// 添加大量数据到缓存
	for i := 0; i < 100000; i++ {
		key := fmt.Sprintf("key-%d", i)
		// 每个值约10KB
		value := make([]byte, 10*1024)
		cache[key] = value

		// 每添加10000项打印一次内存统计
		if (i+1)%10000 == 0 {
			printMemStats(fmt.Sprintf("添加了 %d 项后", i+1))
		}
	}

	// 模拟程序继续运行
	fmt.Println("\n缓存已填充完毕，程序继续运行...")
	time.Sleep(2 * time.Second)

	// 最终内存统计
	printMemStats("最终状态")
}

// 打印内存统计信息
func printMemStats(stage string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("\n--- 内存统计 (%s) ---\n", stage)
	fmt.Printf("分配的内存: %v MB\n", bToMb(m.Alloc))
	fmt.Printf("总分配内存: %v MB\n", bToMb(m.TotalAlloc))
	fmt.Printf("系统内存: %v MB\n", bToMb(m.Sys))
	fmt.Printf("GC次数: %v\n", m.NumGC)
	fmt.Printf("缓存项数: %v\n", len(cache))
}

// 字节转换为MB
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
