package main

import (
	"fmt"
	"math/rand"
	"time"

	"./sort"
)

func randomArray(size int) []int {
	var array []int
	for i := 0; i < size; i++ {
		array = append(array, rand.Intn(size-1))
	}
	return array
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var array = randomArray(32)
	fmt.Println(array)
	sort.QuickSort(&array, 0, len(array)-1)
	fmt.Println(array)

	array = randomArray(32)
	fmt.Println(array)
	sort.HeapSort(&array, len(array))
	fmt.Println(array)
}
