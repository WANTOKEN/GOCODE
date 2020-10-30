package main
import "fmt"

func findMedianSortedArrays(nums1 []int,nums2 []int) float64{
	len1 := len(nums1)
	len2 := len(nums2)

	lenSum := len1 + len2
	if lenSum == 0 {
		return float64(0)
	}
	l,r := 0,0
	a := make([]int,0,lenSum)
	for l<len1 && r<len2{
		if nums1[l]<nums2[r]{
			a = append(a,nums1[l])
			l++
		}else{
			a = append(a,nums2[r])
			r++
		}
	}
	fmt.Println(1,a)
	a = append(a,nums1[l:]...)
	fmt.Println(2,a)
	a = append(a,nums2[r:]...)
	fmt.Println(3,a)
	if lenSum%2 != 0{
		return float64(a[lenSum/2])
	}else{
		return (float64(a[lenSum/2-1])+float64(a[lenSum/2]))/2
	}
}

//log(m+n)
// 时间复杂度：O(logmin(m,n))，
// 其中 m 和 n 分别是数组 nums1 和 nums2 的长度。查找的区间是 [0, m][0,m]，
// 而该区间的长度在每次循环之后都会减少为原来的一半。所以，
// 只需要执行 logm 次循环。由于每次循环中的操作次数是常数，所以时间复杂度为O(logm)。
// 由于我们可能需要交换 nums1 和 nums2 使得 m≤n，因此时间复杂度是 O(logmin(m,n))
func findMedianSortedArrays2(nums1 []int,nums2 []int) float64{
	if len(num1) > len(num2) {  //小的放在前面
		return findMedianSortedArrays2(nums2,nums1)
	}
	
}

func main(){
	var nums1 = []int{1,2} 
	nums2 := []int{3,4}
	fmt.Println(findMedianSortedArrays(nums1,nums2))
}