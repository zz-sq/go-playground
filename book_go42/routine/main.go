package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 1000; i++ {
		Factorial(1000)
	}
	quit <- 1
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

var wg1, wg2 sync.WaitGroup

func main() {
	fmt.Println("1:", time.Now())
	a := 5000
	for i := 1; i <= a; i++ {
		wg1.Add(1)
		go loop()
	}
	for i := 0; i < a; i++ {
		<-quit
		// fmt.Println("receive c:", c)
		wg1.Done()
	}
	fmt.Println("2:", time.Now())
	wg1.Wait()
	fmt.Println("3:", time.Now())
	runtime.GOMAXPROCS(8) // 设置执行使用的核数

	a = 5000
	for i := 1; i <= a; i++ {
		wg2.Add(1)
		go loop()
	}
	for i := 0; i < a; i++ {
		<-quit
		// fmt.Println("receive d:", d)
		wg2.Done()
	}
	fmt.Println("4:", time.Now())
	wg2.Wait()
	fmt.Println("5:", time.Now())

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println("hello:", i)
		}(i)
	}
	fmt.Println("runtime.NUmCPU()", runtime.NumCPU())
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())
	time.Sleep(time.Second * 1) // 等待 5 秒，让协程有机会执行

}
