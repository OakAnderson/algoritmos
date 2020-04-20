// https://en.wikipedia.org/wiki/Bitonic_sorter
//
// Biotonic sort é um algoritmo de ordenação que usa multiprocessos
// Este algoritmo só pode ordenar listas com tamanho em potência de 2
//
// Performace no pior caso em paralelo: O(log(n)²)
// Performace no pior caso sem paralelamento: O(nlog(n)²)
package sort

import "errors"

func compare(arr []int, reverse bool) []int {
	n := len(arr) / 2
	for i := 0; i < n; i++ {
		if reverse != (arr[i] > arr[i+n]) {
			arr[i], arr[i+n] = arr[i+n], arr[i]
		}
	}

	return arr
}

func biotonicMerge(arr []int, reverse bool) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	arr = compare(arr, reverse)
	left := biotonicMerge(arr[:n/2], reverse)
	right := biotonicMerge(arr[n/2:], reverse)

	return append(left, right...)
}

func BiotonicSort(arr []int, reverse bool) ([]int, error) {
	n := len(arr)
	if n <= 1 {
		return arr, nil
	}

	if !(n&(n-1) == 0) {
		return nil, errors.New("O tamanho da lista deve ser potência de 2")
	}
	left, _ := BiotonicSort(arr[:n/2], true)
	right, _ := BiotonicSort(arr[n/2:], false)

	return biotonicMerge(append(left, right...), reverse), nil
}
