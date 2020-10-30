package main

import (
    "fmt"
	"net"
	"time"
	"sync"
	"flag"
)

// func main() {
//     _, err := net.Dial("tcp", "baidu.com:80")
//     if err == nil {
//         fmt.Println("Connection successful")
//     } else {
//         fmt.Println(err)
//     }
// }

// func main(){
// 	for port := 80;port <100;port++{
// 		conn,err := net.Dial("tcp",fmt.Sprintf("baidu.com:%d",port))
// 		if err == nil {
// 			conn.Close()
// 			fmt.Println("Connnection successful")
// 		}else{
// 			fmt.Println(err)
// 		}
// 	}
// }

func isOpen(host string,port int,timeout time.Duration) bool{
	time.Sleep(time.Millisecond * 1)
	conn,err := net.DialTimeout("tcp",fmt.Sprintf("%s:%d",host,port),timeout)
	if err == nil{
		_ = conn.Close()
		return true
	}
	return false
}

// func main() {
// 	ports := []int{}
  
// 	wg := &sync.WaitGroup{}
// 	timeout := time.Millisecond * 500
// 	for port := 1; port < 100000; port++ {
// 	   wg.Add(1)
// 	   go func(p int) {
// 		  opened := isOpen("baidu.com", p,timeout)
// 		  if opened {
// 			 ports = append(ports, port)
// 		  }
// 		  wg.Done()
// 	   }(port)
// 	}
  
// 	wg.Wait()
// 	fmt.Printf("opened ports: %v\n", ports)
//   }

func main() {
    hostname := flag.String("hostname", "", "hostname to test")
    startPort := flag.Int("start-port", 80, "the port on which the scanning starts")
    endPort := flag.Int("end-port", 100, "the port from which the scanning ends")
    timeout := flag.Duration("timeout", time.Millisecond * 500, "timeout")
    flag.Parse()

    ports := []int{}

    wg := &sync.WaitGroup{}
    for port := *startPort; port <= *endPort; port++ {
        wg.Add(1)
        go func(p int) {
            opened := isOpen(*hostname, p, *timeout)
            if opened {
                ports = append(ports, p)
            }
            wg.Done()
        }(port)
    }

    wg.Wait()
    fmt.Printf("opened ports: %v\n", ports)
}