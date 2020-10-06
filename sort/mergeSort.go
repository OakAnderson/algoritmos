// https://pt.wikipedia.org/wiki/Merge_sort
//
// Performance no melhor caso: O(n log n)
// Performance no pior caso: O(n log n)

package sort

func MergeSort(arr []int) {
    if len(arr) > 1 {
        meio := len(arr)/2
        MergeSort(arr[:meio])
        MergeSort(arr[meio:])
        merge(arr[:meio], arr[meio:])

    } 
}

func merge(esq []int, dir []int) {
    var i, j int
    final := append(esq, dir...)
    temp := make([]int, len(esq)+len(dir))
    for i < len(esq) && j < len(dir) {
        if esq[i] < dir[j] {
            temp[i+j] = esq[i]
            i++
        } else {
            temp[i+j] = dir[j]
            j++
        }
    }
    
    for i < len(esq) {
        temp[i+j] = esq[i]
        i++
    }
    for j < len(dir) {
        temp[i+j] = dir[j]
        j++
    }
    
    for k := range final {
        final[k] = temp[k]
    }
}

