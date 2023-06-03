package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	dataset, exists_map, notexists_map := generateDataset(1000)

	separator()

	for capacity := 10; capacity <= 10_000; capacity += 10 {
		bloomFilter := NewBloomFilter(capacity)

		for key := range exists_map {
			bloomFilter.Add(key)
		}

		negativesCount := 0

		for i := 0; i < len(dataset); i++ {
			key := dataset[i]
			_, _, exists := bloomFilter.Exists(key)

			if exists {
				_, existsInNotExists := notexists_map[key]

				if existsInNotExists {
					negativesCount++
				}
			}
		}

		fmt.Printf("Capacity: %d, False positive rate: %.3f%%\n", capacity, 100*(float64(negativesCount)/float64((len(dataset)))))
	}

	separator()
}

func generateDataset(size int) ([]string, map[string]bool, map[string]bool) {
	var dataset []string
	exists_map := make(map[string]bool)
	notexists_map := make(map[string]bool)

	half := size / 2

	for i := 0; i < half; i++ {
		dataset = append(dataset, uuid.New().String())
		exists_map[dataset[i]] = true
	}

	for i := half; i < size; i++ {
		dataset = append(dataset, uuid.New().String())
		notexists_map[dataset[i]] = false
	}

	return dataset, exists_map, notexists_map
}

func separator() {
	println("-----------------------------------------------------------------")
}
