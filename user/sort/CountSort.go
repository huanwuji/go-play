package main
import "fmt"

func countSort(arr []int) []int {
	min,max:=arr[0],arr[0]
	for i,l:=0,len(arr);i<l;i=i+1 {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	countArr:=make([]int, max-min+1)
	for i,l:=0,len(arr);i<l;i=i+1 {
		countArr[arr[i]-min] = 1
	}
	for i,l:=1,len(countArr);i<l;i++ {
		countArr[i] = countArr[i] + countArr[i-1]
	}
	sortArr:=make([]int, len(arr))
	fmt.Println(countArr)
	for i,l:=0,len(sortArr);i<l;i=i+1 {
		sortArr[countArr[arr[i]-min]-1]=arr[i]
	}
	return sortArr
}

func main() {
	arr := []int{2, 3, 5, 1, 13, 41, 22}
	arr=countSort(arr)
	fmt.Print(arr)
}
