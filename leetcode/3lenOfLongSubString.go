// 3. 无重复字符的最长子串
// 题目描述
// 力扣（LeetCode）链接
// 给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
// 示例1：
// 输入: "abcabcbb"
// 输出: 3 
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	i := 0 //计数开始下标
	max := 0
	a := []rune(s)  //转成字符
	fmt.Println(a)
	for m,c := range a {
		for n:= i;n<m;n++{
			if a[n] == c {
				fmt.Println("before:",i)
				i = n+1
			}
		}
		if m-i+1 > max {
			max = m-i+1
		}
	}

	fmt.Println("max:",max)
	return max
}

func main(){
	var s =  "abcabcbb"
	lengthOfLongestSubstring(s)
}