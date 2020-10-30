package main

import "fmt"

// func main() {
//     str := "hello"
//     str[0] = 'x'
//     fmt.Println(str)
// }

// func incr(p *int) int {
//     *p++
//     return *p
// }

// func main() {
//     p :=1
//     incr(&p)
//     fmt.Println(p)
// }


func add(args ...int) int {

    sum := 0
    for _, arg := range args {
        sum += arg
    }
    return sum
}

// func main(){
// 	fmt.Println(add([]int{1, 2}...))
// }


// func main() {
//     var s1 []int
//     // var s2 = []int{}
//     if s1 == nil {
//         fmt.Println("yes nil")
//     }else{
//         fmt.Println("no nil")
//     }
// }

// type A interface {
//     ShowA() int
// }

// type B interface {
//     ShowB() int
// }

// type Work struct {
//     i int
// }

// func (w Work) ShowA() int {
//     return w.i + 10
// }

// func (w Work) ShowB() int {
//     return w.i + 20
// }

// // func main() {
// //     c := Work{3}
// //     var a A = c
// //     var b B = c
// //     fmt.Println(a.ShowA())
// //     fmt.Println(b.ShowB())
// // }

// func main() {
//     m := make(map[string]int)//A
// 	m["a"] = 1
//     if v,ok := m["b"];ok {  //B
// 		fmt.Println(v)
		
//     }
// }

type A interface {
    ShowA() int
}

type B interface {
    ShowB() int
}

type Work struct {
    i int
}

func (w Work) ShowA() int {
    return w.i + 10
}

func (w Work) ShowB() int {
    return w.i + 20
}

func main() {
    c := Work{3}
    var a A = c
    var b B = c
    fmt.Println(a.ShowA())
    fmt.Println(b.ShowB())
}