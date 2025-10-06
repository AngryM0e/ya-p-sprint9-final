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
	//create determenism predictable generator
	rng := rand.New(rand.NewSource(42))

	slice := make([]int, size)
	for i := 0; i < size; i++ {
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
	//находим максимальное значение
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

	//Сreate len(8) slice
	maxes := make([]int, CHUNKS)
	var wg sync.WaitGroup

	//Size every piece
	chunkSize := n / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		//Find first & last indexes
		start := i * chunkSize
		end := start + chunkSize

		//Обработка последне чанка (может включать "хвостик")
		if i == CHUNKS-1 {
			end = n // для последнего чанка берем до конца слайса
		}

		// Если start >= end, пропускаем пустой чанк
		if start >= end {
			maxes[i] = data[0] // min value
			continue
		}

		wg.Add(1)
		go func(i, start, end int) {
			defer wg.Done()

			//Находим максимум в своем чанке
			chunkMax := data[start]
			for j := start + 1; j < end; j++ {
				if data[j] > chunkMax {
					chunkMax = data[j]
				}
			}
			//Записываем результат в заранее выделенный слот
			maxes[i] = chunkMax
		}(i, start, end)
	}

	//Wait goroutins end
	wg.Wait()

	//Find overall chunks maximum
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
