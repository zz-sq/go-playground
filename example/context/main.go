package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

var logs *log.Logger

func doClear(ctx context.Context) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			logs.Println("done")
			return
		default:
			logs.Println("wait for done")
		}
	}
}

func doNothing(ctx context.Context) {
	for {
		time.Sleep(time.Second * 3)
		select {
		case <-ctx.Done():
			logs.Println("doNothing: done")
		default:
			logs.Println("doNothing: wait for done")
		}
	}
}

func A(ctx context.Context) int {
	ctx = context.WithValue(ctx, "AFunction", "Great")

	go B(ctx)

	select {
	// 监测自己上层的ctx ...
	case <-ctx.Done():
		fmt.Println("A Done")
		return -1
	}
	return 1
}

func B(ctx context.Context) int {
	fmt.Println("A value in B:", ctx.Value("AFunction"))
	ctx = context.WithValue(ctx, "BFunction", 999)

	go C(ctx)

	select {
	// 监测自己上层的ctx ...
	case <-ctx.Done():
		fmt.Println("B Done")
		return -2
	}
	return 2
}

func C(ctx context.Context) int {
	fmt.Println("B value in C:", ctx.Value("AFunction"))
	fmt.Println("B value in C:", ctx.Value("BFunction"))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		fmt.Println("C Done")
		return -3
	}
	return 3
}

func main() {
	// logs = log.New(os.Stdout, "", log.Ltime)
	// ctx, cancel := context.WithCancel(context.Background())
	// go doClear(ctx)
	// go doNothing(ctx)

	// time.Sleep(time.Second * 20)
	// cancel()
	// time.Sleep(time.Second * 10)

	// 自动取消(定时取消)
	{
		timeout := 10 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)

		fmt.Println("A 执行完成，返回：", A(ctx))
		select {
		case <-ctx.Done():
			fmt.Println("context Done")
			break
		}
	}
	time.Sleep(20 * time.Second)
}
