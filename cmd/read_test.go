package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func create() {
	filename := "1GB.bin" // 替换为你想要的文件名

	// 创建一个新文件
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 定义文件大小和缓冲区大小
	fileSize := int64(1024 * 1024 * 1024) // 1GB
	bufferSize := 1024 * 1024             // 1MB

	// 创建一个缓冲区
	buffer := make([]byte, bufferSize)

	// 循环写入数据，直到达到文件大小
	for written := int64(0); written < fileSize; {
		n, err := file.Write(buffer)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		written += int64(n)
	}

	fmt.Printf("1GB file '%s' created successfully.\n", filename)
}

func BenchmarkReadAll(b *testing.B) {
	f, _ := os.Open("1GB.bin")
	for i := 0; i <= b.N; i++ {
		_, _ = io.ReadAll(f)
	}
}

func BenchmarkReadFrom(b *testing.B) {
	f, _ := os.Open("1GB.bin")
	for i := 0; i <= b.N; i++ {
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(f)
	}
}
