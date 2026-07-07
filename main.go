package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) ([]int, error) {

	if size <= 0 {
		return []int{}, errors.New("the slice size is 0 or less")
	}
	random := rand.Perm(size)
	return random, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {

	if len(data) == 0 {
		return 0, errors.New("the slice length is 0")
	}

	max := data[0]

	for _, w := range data {
		if w > max {
			max = w
		}
	}

	return max, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {

	if len(data) == 0 {
		return 0, errors.New("the slice length is 0")
	}

	cut := len(data) / CHUNKS
	cutSlice := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	start := 0
	for i := 0; i < CHUNKS; i++ {
		i := i

		localStart := start
		localEnd := start + cut

		if localEnd > len(data) {
			localEnd = len(data)
		}
		start = localEnd

		go func() error {
			defer wg.Done()

			max, err := maximum(data[localStart:localEnd])
			if err != nil {
				return errors.New("number search error")

			}
			cutSlice[i] = max
			return nil
		}()
	}

	wg.Wait()
	max2, err := maximum(cutSlice)
	if err != nil {
		return 0, errors.New("number search error")
	}
	return max2, nil
}

func main() {

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	slice, err := generateRandomElements(SIZE)
	if err != nil {
		log.Fatal("generation error")
	}

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max, err := maximum(slice)
	if err != nil {
		log.Fatal("error searching for a number in the maximum")
	}
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	max, err = maxChunks(slice)
	if err != nil {
		log.Fatal("error searching for a number in the maxChunks")
	}
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
