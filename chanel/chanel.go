package main
import "fmt"


func recv(c chan int) {
    ret := <-c
    fmt.Println("接收成功", ret)
}
func main() {
    ch := make(chan int)
	
	for i:=0;i<10;i++{
		go recv(ch) // 启用goroutine从通道接收值
		ch <- i
	}

    fmt.Println("发送成功")
}