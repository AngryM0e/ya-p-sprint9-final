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
	// Create deterministic generator with fixed seed for predictable results
	rng := rand.New(rand.NewSource(42))

	slice := make([]int, size)
	for i := 0; i < size; i++ {
		// Generate numbers in range [0, 1000]
		slice[i] = rng.Intn(1000)
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
	maxes := make([]int, CHUNKS)
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
			maxes[i] = data[0]
			continue
		}

		wg.Add(1)
		// Launch goroutine to process chunk
		go func(i, start, end int) {
			defer wg.Done()

			// Find maximum within the assigned chunk
			chunkMax := data[start]
			for j := start + 1; j < end; j++ {
				if data[j] > chunkMax {
					chunkMax = data[j]
				}
			}
			// Store result in pre-allocated slot
			maxes[i] = chunkMax
		}(i, start, end)
	}

	//Wait for all goroutines to complete
	wg.Wait()

	// Find overall maximim from chunk results
	overallMax := maxes[0]
	for i := 1; i < CHUNKS; i++ {
		if maxes[i] > overallMax {
			overallMax = maxes[i]
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
