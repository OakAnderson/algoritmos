// https://pt.wikipedia.org/wiki/Insertion_sort
//
// Performace no melhor caso:   O(n)
// Performace no pior caso:     O(n²)
package sort

import "fmt"

func InsertionSort(arr []int, simulation bool) []int {
	iteration := 0
	if simulation {
		fmt.Printf("iteração %d : %v\n", iteration, arr)
	}

	for i := 0; i < len(arr); i++ {
		cursor := arr[i]
		pos := i

		for pos > 0 && arr[pos-1] > cursor {
			arr[pos] = arr[pos-1]
			pos = pos - 1
		}
		arr[pos] = cursor

		if simulation {
			iteration++
			fmt.Printf("iteração %d : %v\n", iteration, arr)
		}
	}

	return arr
}
