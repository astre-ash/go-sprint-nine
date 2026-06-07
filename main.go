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

	// Magic values: rand.Intn(100) + (-50) gives the [-50, 49] range.
	minVal = -50
	maxVal = 50
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	slice := make([]int, size)
	rangeSize := maxVal - minVal

	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(rangeSize) + minVal
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, num := range data {
		if num > max {
			max = num
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

	activeChunks := CHUNKS
	if n < CHUNKS {
		activeChunks = n
	}

	maxValues := make([]int, activeChunks)
	var wg sync.WaitGroup

	start := 0

	for i := 0; i < activeChunks; i++ {
		// Dynamically calculate the chunk size to evenly distribute the remainder elements.
		chankSize := (n - start) / (activeChunks - i)
		end := start + chankSize
		chunk := data[start:end]
		start = end

		wg.Add(1)
		go func(c []int, idx int) {
			defer wg.Done()
			max := maximum(c)
			maxValues[idx] = max
		}(chunk, i)
	}

	wg.Wait()
	overallMax := maxValues[0]

	for _, v := range maxValues {
		if v > overallMax {
			overallMax = v
		}
	}

	return overallMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	s := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(s)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальный элемент массива: %d\nВремя поиска: %d us\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	max = maxChunks(s)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальный элемент массива: %d\nВремя поиска: %d us\n", max, elapsed)
}
