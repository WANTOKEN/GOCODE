package main

import "fmt"

func SelectMax(arr []int) int{
	if len(arr) <= 1{
		return arr[0]
	}
	max := arr[0]
	for i:=0;i < len(arr);i++{
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}


func SelectSort(arr []int) []int{
	if len(arr)<=1{
		return arr
	}
	length := len(arr)
	for i:=0;i<length;i++{
		min := i
		for j :=i+1;j<length;j++{
			if arr[min]>arr[j]{
			 min = j
			}
		}
		if min != i {
			arr[min],arr[i] = arr[i],arr[min]
		}
	}
	return arr

}


func main(){

	arr := []int{1,3,0,3,7}
	fmt.Println(SelectMax(arr))
	fmt.Println(SelectSort(arr))

}