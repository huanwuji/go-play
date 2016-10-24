package main
import (
	"fmt"
)

func binnaryInsertSort(arr []int) []int {
	for i, l := 1, len(arr); i < l; i++ {
		temp := arr[i]
		idx := binnerySearch(arr, 0, i - 1, temp)
		cursor := i
		for cursor > idx {
			arr[cursor] = arr[cursor-1]
			cursor--
		}
		arr[idx] = temp
	}
	return arr
}

func binnerySearch(arr []int, s int, e int, v int) int {
	m := (s + e) / 2
	if arr[m] < v {
		if(m == e) {
			return m + 1
		} else {
			return binnerySearch(arr, m+1, e, v)
		}
	} else {
		if m == 0 {
			return m
		} else {
			return binnerySearch(arr, s, m-1, v)
		}
	}
}

func main() {
	arr := []int{45, 64, 34, 8, 4, 588, 7, 99}
	arr = binnaryInsertSort(arr)
	fmt.Println(arr)
}
