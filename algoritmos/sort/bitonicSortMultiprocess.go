// https://en.wikipedia.org/wiki/Bitonic_sorter
//
// Bitonic sort é um algoritmo de ordenação que usa multiprocessos
// Este algoritmo só pode ordenar listas com tamanho em potência de 2
//
// Performace no pior caso em paralelo: O(log(n)²)
// Performace no pior caso sem paralelamento: O(nlog(n)²)
package sort

import (
	"errors"
	"sync"
)

func bitonicMergeMultiprocess(arr []int, reverse bool) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	arr = compare(arr, reverse)

	var left, right []int
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		left = bitonicMerge(arr[:n/2], reverse)
		wg.Done()
	}()
	go func() {
		right = bitonicMerge(arr[n/2:], reverse)
		wg.Done()
	}()

	wg.Wait()

	return append(left, right...)
}

func BitonicSortMultiprocess(arr []int, reverse bool) ([]int, error) {
	n := len(arr)
	if n <= 1 {
		return arr, nil
	}

	if !(n&(n-1) == 0) {
		return nil, errors.New("O tamanho da lista deve ser potência de 2")
	}

	var left, right []int
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		left, _ = BitonicSort(arr[:n/2], true)
		wg.Done()
	}()
	go func() {
		right, _ = BitonicSort(arr[n/2:], false)
		wg.Done()
	}()

	wg.Wait()

	return bitonicMergeMultiprocess(append(left, right...), reverse), nil
}
