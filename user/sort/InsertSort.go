package main
import (
	"fmt"
)

func insertSort(arr []int) []int {
	for i, l := 1, len(arr); i < l; i++ {
		temp := arr[i]
		j := i - 1
		for j > -1 && arr[j] > temp {
			arr[j + 1] = arr[j]
			j--
		}
		arr[j + 1] = temp
	}
	return arr
}

func main() {
	arr := []int{34, 3, 45, 32, 75, 88, 24}
	arr = insertSort(arr)
	fmt.Println(arr)