package main

import (
	"fmt"
	"time"
)

// 只能向chan里send发送数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {

		fmt.Println("send ready ", i)
		c <- i
		fmt.Println("send ", i)
	}
}

// 只能接收通道中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received ", i)
	}
}

func main() {
	c := make(chan int)
	go send(c)
	go recv(c)
	time.Sleep(time.Second * 1)
	close(c)

	var c1, c2, c3 chan int = make(chan int, 10), make(chan int), make(chan int, 10)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- 10
	}()

	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received from c1, %d", i1)
	case c2 <- i2:
		fmt.Printf("sent to c2, %d", i2)
	case i3, ok := <-c3:
		if ok {
			fmt.Printf("received from c3, %d", i3)
		} else {
			fmt.Printf("c3 is closed")
		}
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	}
}
