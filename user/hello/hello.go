package main

import (
	"fmt"
)

func quickSort(arr []int, low int, high int) {
	if low < high {
		tmp := arr[low]
		left := low
		right := high
		for low < high {
			for arr[high] > tmp && low < high {
				high--
			}
			arr[low] = arr[high]
			for arr[low] < tmp && low < high {
				low++
			}
			arr[high] = arr[low]
		}
		arr[low] = tmp
		quickSort(arr, left, low - 1)
		quickSort(arr, low + 1, right)
	}
}

func main() {
	arr := []int{10, 5, 7, 17, 2, 8, 9, 11}
	quickSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}