package main

import (
	"log"
	"os"
)

var cwd string
var cwd2 string

func init() {
	// 这里初始化的是局部变量，不是全局变量，全局变量 cwd 仍然是 ""
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	// cwd2 是全局变量
	cwd2, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}

func main() {
	log.Printf("cwd = %s", cwd)
	log.Printf("cwd2 = %s", cwd2)
}
