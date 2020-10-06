// https://pt.wikipedia.org/wiki/Insertion_sort
//
// Performace no melhor caso:   O(n)
// Performace no pior caso:     O(n²)
package sort

func InsertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		cursor := arr[i]
		pos := i

		for pos > 0 && arr[pos-1] > cursor {
			arr[pos] = arr[pos-1]
			pos--
		}
		arr[pos] = cursor
	}

	return arr
}
