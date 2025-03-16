package performance_examples

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func RunFileExample() {
	// 创建测试文件
	filename := "test_file.txt"
	createTestFile(filename, 1000000) // 创建一个包含100万行的文件
	defer os.Remove(filename)         // 测试完成后删除文件

	fmt.Println("测试不同的文件读取方法...")

	// 方法1: 按行读取
	start := time.Now()
	count1 := readLineByLine(filename)
	duration1 := time.Since(start)
	fmt.Printf("方法1 (按行读取): 处理了 %d 行, 耗时 %v\n", count1, duration1)

	// 方法2: 一次性读取整个文件
	start = time.Now()
	count2 := readEntireFile(filename)
	duration2 := time.Since(start)
	fmt.Printf("方法2 (整体读取): 处理了 %d 字节, 耗时 %v\n", count2, duration2)

	// 方法3: 使用缓冲读取
	start = time.Now()
	count3 := readWithBuffer(filename)
	duration3 := time.Since(start)
	fmt.Printf("方法3 (缓冲读取): 处理了 %d 字节, 耗时 %v\n", count3, duration3)
}

// 创建测试文件
func createTestFile(filename string, lines int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < lines; i++ {
		fmt.Fprintf(writer, "这是测试文件的第 %d 行内容\n", i+1)
	}
	writer.Flush()
}

// 方法1: 按行读取文件
func readLineByLine(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		_ = scanner.Text() // 读取每一行但不做任何处理
		lineCount++
	}

	return lineCount
}

// 方法2: 一次性读取整个文件
func readEntireFile(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return len(data)
}

// 方法3: 使用缓冲读取
func readWithBuffer(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, 4096) // 4KB缓冲区
	totalBytes := 0

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if bytesRead == 0 {
			break
		}

		totalBytes += bytesRead
	}

	return totalBytes
}
