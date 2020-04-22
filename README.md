Golang Data Structures and Algorithms
=========================================

Exemplos simples de implementação de estrutura de dados e algoritmos em Go.

Você pode testar criando um arquivo go: (Ex: uso de `BubbleSort` em `sort`)

```go
package main

import (
    "github.com/OakAnderson/algoritmos/algoritmos/sort"
)

func main() {
    myList := []int{1, 8, 3, 5, 6}
    myList = sort.BubbleSort(myList)
    print(myList)
}
```


## List of Implementations

- [math](algoritmos/math)
    - [DecimalToBinaryIP](algoritmos/math/decimalToBinaryIP.go)
- [sort](algoritmos/sort)
    - [BitonicSort](algoritmos/sort/bitonicSort.go)
    - [BitonicSortMultiprocess](algoritmos/sort/bitonicSortMultiprocess.go)
    - [BubbleSort](algoritmos/sort/bubbleSort.go)
    - [InsertionSort](algoritmos/sort/insertionSort.go)