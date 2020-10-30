package main
import "fmt"

// type Myint int
// func (i Myint) PrintInt ()  {
//     fmt.Println(i)
// }

// func main() {
//     var i Myint = 1
//     i.PrintInt()
// }


type People interface {
    Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
    if think == "speak" {
        talk = "speak"
    } else {
        talk = "hi"
    }
    return
}

func main() {
    var peo People = Student{}
    think := "speak"
    fmt.Println(peo.Speak(think))
}