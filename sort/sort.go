package sort

func HeapSort(a *[]int, size int) {
	for i := (size - 1) / 2; i >= 0; i-- {
		downHeap(a, i, size-1)
	}

	for i := size - 1; i > 0; i-- {
		(*a)[0], (*a)[i] = (*a)[i], (*a)[0]
		downHeap(a, 0, i-1)
	}
}

func downHeap(a *[]int, root int, bottom int) {
	var child = 2*root + 1
	var temp = (*a)[root]
	for child <= bottom {
		if child < bottom && (*a)[child+1] > (*a)[child] {
			child++
		}
		if temp > (*a)[child] {
			break
		} else {
			(*a)[(child-1)/2] = (*a)[child]
			child = 2*child + 1
		}
	}
	(*a)[(child-1)/2] = temp
}

func QuickSort(a *[]int, left int, right int) {
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
		QuickSort(a, left, i-1)
		QuickSort(a, j+1, right)
	}
}
