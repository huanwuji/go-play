package main
import "fmt"

func shellSort(arr []int) []int {
	n := len(arr)
	for gap := n / 2; gap > 0; gap = gap / 2 {
		for i := 0; i < gap; i++ {
			for j := i + gap; j < n; j = j + gap {
				temp := arr[j]
				k := j
				for k >= gap && temp < arr[k - gap] {
					arr[k] = arr[k - gap]
					k = k - gap
				}
				arr[k] = temp
			}
		}
	}
	return arr
}

func shellSort1(arr []int) []int {
	n := len(arr)
	for gap := n / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			k := i - gap
			for k > -1 && temp < arr[k] {
				arr[k + gap] = arr[k]
				k = k-gap
			}
			arr[k + gap] = temp
		}
	}
	return arr
}

func shellSort2(arr []int) []int {
	n := len(arr)
	for gap := n / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < n; i++ {
			for j := i - gap; j > -1 && arr[j] > arr[j + gap]; j = j - gap {
				arr[j], arr[j + gap] = arr[j + gap], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{34, 99, 45, 32, 3, 88, 22}
	arr = shellSort1(arr)
	fmt.Println(arr)
}