package main
import "fmt"

func bubbleSort(arr []int) []int {
	for i, l := 0, len(arr); i < l; i = i + 1 {
		for j := l - i - 1; j > 0; j = j - 1 {
			if arr[j] < arr[j - 1] {
				arr[j - 1], arr[j] = arr[j], arr[j - 1]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{2, 3, 5, 1, 13, 41, 22}
	bubbleSort(arr)
	fmt.Print(arr)
}