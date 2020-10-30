package main
import(
	"fmt"
)

// func main(){
// 	var x string = nil
// 	if x == nil {
// 		x = "default"
// 	}
// 	fmt.Println(x)
// }

var a bool = true
func main() {
    defer func(){
        fmt.Println("1")
	}()

    if a == true {
        fmt.Println("2")
        return
    }
	defer func(){
        fmt.Println("3")
    }()
}