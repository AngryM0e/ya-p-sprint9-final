package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size == 0 {
		return make([]int, 0)
	}

	slice := make([]int, size)
	for i := 0; i < size; i++ {
		// Generate random positive integers
		slice[i] = rand.Int()
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	// Iterate through slice to find maximum value
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	n := len(data)
	if n == 0 {
		return 0
	}

	// Pre-allocate slice to store maximums from each chunk
	res := make(chan int, CHUNKS)
	var wg sync.WaitGroup

	// Calculate base chunk size
	chunkSize := n / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		// Calculate chunk boundaries
		start := i * chunkSize
		end := start + chunkSize

		// Last chunk takes all remaining elements
		if i == CHUNKS-1 {
			end = n
		}

		// Skip empty chunks that may occur when n < CHUNKS
		if start >= end {
			continue
		}

		wg.Add(1)
		// Launch goroutine to process chunk
		go func(chunk []int) {
			defer wg.Done()

			chunkMax := maximum(chunk)

			// Send result through channel
			res <- chunkMax
		}(data[start:end])
	}

	// Close channel after all goroutines complete
	go func() {
		wg.Wait()
		close(res)
	}()

	// Collect result from channel
	overallMax := <- res
	for max := range res {
		if max > overallMax {
			overallMax = max
		}
	}

	return overallMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	startParallel := time.Now()
	max = maxChunks(data)
	elapsed = time.Since(startParallel).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
