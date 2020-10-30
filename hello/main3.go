package main
import "fmt"

type person struct {
	name string
}
// func main(){
// 	var m map[person] int
// 	p := person{"mike"}
// 	fmt.Println(m[p])
// }

func hello(num ...int) {  
    num[0] = 18
}

// func main() {  
//     i := []int{5, 6, 7}
//     hello(i...)
//     fmt.Println(i)
// }

// func main() {  
//     a := 5.22
//     b := 8.1
//     fmt.Println(a + b)
// }

// func main() {  
//     a := [5]int{1, 2, 3, 4, 5}
//     t := a[3:4:4]
//     fmt.Println(t)
// }

// func main() {
//     a := [2]int{5, 6}
//     b := [2]int{5, 6}
//     if a == b {
//         fmt.Println("equal")
//     } else {
//         fmt.Println("not equal")
//     }
// }