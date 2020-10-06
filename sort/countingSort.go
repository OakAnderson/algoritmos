// https://pt.wikipedia.org/wiki/Counting_sort
//
// Performance no melhor caso: ___ 
// Performance no pior caso: ___

package sort

func CountingSort(arr []int) []int {
    contadores := make([]int, len(arr))
    for i := range arr {
        for j := i+1; j < len(arr); j++ {
            if arr[i] > arr[j] {
                contadores[i]++
            } else {
                contadores[j]++
            }
        }
    }
    
    final := make([]int, len(arr))
    for i := range final {
        final[contadores[i]] = arr[i]
    }

    return final 
}
