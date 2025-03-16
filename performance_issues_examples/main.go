package performance_examples

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	example := os.Args[1]
	switch example {
	case "cache":
		RunExample()
	case "fibonacci":
		RunFibonacciExample()
	case "file":
		RunFileExample()
	case "counter":
		RunCounterExample()
	case "concurrency":
		RunConcurrencyExample()
	case "database":
		// 注意：运行数据库示例需要安装SQLite驱动
		// go get github.com/mattn/go-sqlite3
		RunDatabaseExample()
	default:
		fmt.Printf("未知示例: %s\n", example)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("使用方法: go run main.go <示例名称>")
	fmt.Println("可用示例:")
	fmt.Println("  cache      - 缓存示例")
	fmt.Println("  fibonacci  - 斐波那契数列")
	fmt.Println("  file       - 文件处理")
	fmt.Println("  counter    - 计数器示例")
	fmt.Println("  concurrency - 并发示例")
	fmt.Println("  database   - 数据库示例 (需要SQLite驱动)")
}
