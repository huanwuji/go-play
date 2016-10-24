package main
import "fmt"

func quickSort(arr []int, l int, r int) {
	s, e := l, r
	fmt.Println(arr, l, r, s, e)
	temp := arr[r]
	for l < r {
		for arr[l] <= temp && l < r {
			l = l + 1
		}
		for arr[r] > temp && l < r {
			r = r - 1
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		} else if l == r {
			arr[s], arr[l] = arr[s], arr[l]
		}
	}
	fmt.Println(arr, l, r, s, e)
	if s < l {
		fmt.Println(arr, l, r, s, e)
		quickSort(arr, s, l - 1)
	}
	if l < e {
		quickSort(arr, l, e)
	}
}

func main() {
	arr := []int{45, 64, 34, 8, 4, 588, 7, 99}
	quickSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}
