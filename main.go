package main

import (
	"fmt"
	"runtime"
)

func main() {
	var memorySlices [][]byte
	for {
		// Allocate a large chunk of memory
		slice := make([]byte, 1024*1024*1000)
		memorySlices = append(memorySlices, slice)

		// Print memory usage
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
	}
}
