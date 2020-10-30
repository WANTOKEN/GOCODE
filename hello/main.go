package main  // 声明 main 包，表明当前是一个可执行程序

import (
    "fmt"  // 导入内置 fmt 
    "sync"
)

func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }

}

// func main() {

//     slice := []int{0,1,2,3}
//     m := make(map[int]*int)

//     for key,val := range slice {
//         value := val
//         m[key] = &value
//     }

//    for k,v := range m {
//        fmt.Println(k,"===>",*v)
//    }
// }

// 1.
// func main() {
//     s := make([]int, 5)
//     s = append(s, 1, 2, 3)
//     fmt.Println(s)
//     fmt.Println(funcMui(1,2))

// }


func funcMui(x,y int)(sum int,err error){
    return x+y,nil
}


// func main() {
//     list := new([]int)
//     list = append(list, 1)
//     fmt.Println(list)
// }

// type inter interface{}
// type Set struct {
//     m map[inter] bool
//     sync.RWMutx
// }

// func New() *Set{
//     return &Set{
//         m: map[inter] bool{},
//     }
// }

// func (s *Set) Add(item inter){
//     s.Lock()
//     defer s.Unlock()
//     s.map[item] = true
// }


// func main() {
//     s1 := []int{1, 2, 3}
//     s2 := []int{4, 5}
//     s1 = append(s1, s2)
//     fmt.Println(s1)
// }

// var(
//     size := 1024
//     max_size = size*2
// )

// func main() {
//     fmt.Println(size,max_size)
// }

// func main() {
//     sn1 := struct {
//         age  int
//         name string
//     }{age: 11, name: "qq"}
//     sn2 := struct {
//         age  int
//         name string
//     }{age: 11, name: "qq"}

//     if sn1 == sn2 {
//         fmt.Println("sn1 == sn2")
//     }

//     sm1 := struct {
//         age int
//         m   map[string]string
//     }{age: 11, m: map[string]string{"a": "1"}}
//     sm2 := struct {
//         age int
//         m   map[string]string
//     }{age: 11, m: map[string]string{"a": "1"}}

//     if sm1 == sm2 {
//         fmt.Println("sm1 == sm2")
//     }
// }

// type MyInt1 int
// type MyInt2 = int

// func main() {
//     var i int =0
//     var i1 MyInt1 = i 
//     var i2 MyInt2 = i
//     fmt.Println(i1,i2)

    
// }

// func main(){
//     str := "abc" + "123";
//     fmt.Println(str)
// }

// func hello() {
//     fmt.Println("Hello Goroutine!")
// }
// func main() {
//     go hello()
//     go hello()
//     go hello()
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     fmt.Println("main goroutine done!")
//     time.Sleep(time.Second)
// }



// var wg sync.WaitGroup

// func hello(i int) {
//     defer wg.Done() // goroutine结束就登记-1
//     fmt.Println("Hello Goroutine!", i)
// }
// func main() {

//     for i := 0; i < 10; i++ {
//         wg.Add(1) // 启动一个goroutine就登记+1
//         go hello(i)
//     }
//     wg.Wait() // 等待所有登记的goroutine都结束
// }