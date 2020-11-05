// 1. 两数之和
// 题目描述
// 力扣（LeetCode）链接
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]


















package main
import "fmt"

// func towSum(nums []int,target int) []int{
// 	len :=len(nums)
// 	for i:=0;i<len;i++ {
// 		for j:=i+1;j<len;j++{
// 			if nums[i]+nums[j] == target {
// 				return []int{i,j}
// 			}
// 		}
// 	}
// 	return []int{}
// }


func towSum(nums []int,target int) []int{
	mapVal := make(map[int]int)
	len :=len(nums)
	for i:=0;i<len;i++ {
		res := target-nums[i]
		v,ok := mapVal[res]
		if ok {
			return []int{v,i}
		}
		mapVal[nums[i]] = i
	}
	return []int{-1,-1}
}

func main(){
	nums := []int{2,7,11,15}
	fmt.Println(towSum(nums,9))
}