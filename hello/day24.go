package main
import "fmt"

func main(){
	m := map[int] string {0:"zereo",1:"one"}
	for k,v := range m{
		fmt.Println(k,v)
	}
}