// https://pt.wikipedia.org/wiki/Selection_sort
//
// Performance no pior caso: O(n^2)
// Performance no melhor caso: O(n^2)

package sort


func SelectionSort(arr []int) []int {
    for i := range arr {
        menorIdx := i
        for prox := i+1; prox < len(arr); prox++ {
            if arr[menorIdx] > arr[prox] {
                menorIdx = prox
            }
        }

        aux := arr[i]
        arr[i] = arr[menorIdx]
        arr[menorIdx] = aux
    }

   return arr
}
