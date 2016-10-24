package main
import "fmt"

func selectSort(arr []int) []int {
	for i,l:=0,len(arr);i<l;i=i+1 {
		temp,idx := arr[i],i
		for j:=i;j<l;j=j+1 {
			if arr[j]<temp {
				temp,idx=arr[j],j
			}
		}
		if i!=idx {
			arr[i],arr[idx] = arr[idx],arr[i]
		}
	}
	return arr
}

func main() {
	arr := []int{45, 64, 34, 8, 4, 588, 7, 99}
	arr = selectSort(arr)
	fmt.Println(arr)
}
