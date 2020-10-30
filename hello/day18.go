package main
import "fmt"

func f1() (r int) {
    defer func() {
        r++
	}()
	fmt.Println("f1()")
    return 0
}

func f2() (r int) {
    t := 5
    defer func() {
        t = t + 5
	}()
	fmt.Println("f2()")	
    return t
}

func f3() (r int) {
    defer func(r int) {
        r = r + 5
	}(r)
	fmt.Println("f3()")
    return 1
}

func main(){
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}