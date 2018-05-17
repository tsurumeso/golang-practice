package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(a *[]int, left int, right int) {
	if left < right {
		var i, j = left, right
		var pivot = (*a)[left]
		for {
			for (*a)[i] < pivot {
				i++
			}
			for pivot < (*a)[j] {
				j--
			}
			if i >= j {
				break
			}
			(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
			i++
			j--
		}
		quickSort(a, left, i-1)
		quickSort(a, j+1, right)
	}
}

func main() {
	var array []int
	var size = 32
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		array = append(array, rand.Intn(size-1))
	}
	fmt.Println(array)
	quickSort(&array, 0, len(array)-1)
	fmt.Println(array)
}
