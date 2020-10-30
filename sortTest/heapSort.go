package main

import "fmt"

//堆排序
func main(){
	arr :=  [] int{7,3,8,5,1,2} 
	fmt.Println(HeapSort(arr))
}

func HeapSortMax(arr []int,length int) []int{
	if len(arr) <= 1{
		return arr
	}
	depth := length/2 - 1 //最后一个非叶子节点

	for i:= depth ; i>=0;i--{
		topmax := i   
		leftchild := 2*i+1
		rightchild :=2*i+2
		if leftchild <= length-1 && arr[leftchild] > arr[topmax]{
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax]{
			topmax = rightchild			
		}
		if topmax != i{
			arr[i],arr[topmax] = arr[topmax],arr[i]
		}
	}
	return arr
}

func HeapSort(arr []int) []int{
	length :=len(arr)
	for i:=0;i < length; i++ {
		lastlen := length - i //最后的叶子节点
		HeapSortMax(arr,lastlen)
		if i <length{
			arr[0],arr[lastlen-1] = arr[lastlen-1],arr[0]
		}
		fmt.Println(arr)
	}
	return arr
}