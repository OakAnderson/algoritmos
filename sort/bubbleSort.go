// https://pt.wikipedia.org/wiki/Bubble_sort
//
// Performace no pior caso: O(n²)
// Chamando BubbleSort(arr, true) será possível observar os processos de
// ordenação
package sort

func BubbleSort(arr []int) []int {
	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}

	n := len(arr)
	swapped := true

	for x := 0; swapped; x++ {
		swapped = false
		for i := 1; i < n-x; i++ {
			if arr[i-1] > arr[i] {
				swap(i-1, i)
				swapped = true
			}
		}
	}

	return arr
}
