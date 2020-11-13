//洗牌算法

package main

import (
	"fmt"
	"rand"
	"time"
) 

func main(){
	vals := []int{10,12,14,16,18,20}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(r)
}