package main

import "fmt"

func reverse(x int )int{
	var res int
	for x != 0 {
		if temp := int32(res); (temp*10)/10 != temp {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	return res;
}


func main(){
	qes := -2147483412
	fmt.Println(int32(qes))
	fmt.Println(int32(qes)*10)
	fmt.Println((int32(qes)*10)/10)
	fmt.Println(reverse(qes))


}