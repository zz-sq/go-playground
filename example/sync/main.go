package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}

func main() {
	wg := sync.WaitGroup{}
	var mutex sync.Mutex
	fmt.Println("Lock G0")
	mutex.Lock()
	wg.Add(3)

	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Lock G(%d)\n", i)
			mutex.Lock()
			fmt.Printf("Locked G(%d)\n", i)
			time.Sleep(time.Second * 2)

			fmt.Printf("Unlock G(%d)\n", i)
			mutex.Unlock()
			fmt.Printf("Unlocked G(%d)\n", i)
			wg.Done()
		}(i)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("Unlock G0")
	mutex.Unlock()
	fmt.Println("Unlocked G0")
	wg.Wait()
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("v:", v, "---i:", i)
	}

	for i := 0; i < 10; i++ {

		go func(i int) {
			once.Do(onced)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 10)
}
