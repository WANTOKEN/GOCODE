package main

import "fmt"

func numWaterBottles(numBottles int,numExchange int) int {
	ans :=  numBottles
	tmp := 0
	for numBottles>=numExchange{
		tmp = numBottles/numExchange    //可换酒数
		numBottles = numBottles%numExchange //剩余空瓶
		numBottles += tmp
		ans += tmp
	}
	return ans


}

func main(){
	fmt.Println(numWaterBottles(15,4))
}