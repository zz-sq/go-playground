package main

import "fmt"

type A struct {
	Books int
}

type B interface {
	f()
}

func (a A) f() {
	fmt.Println("A.f(): ", a.Books)
}

type I int

func (i I) f() {
	fmt.Println("I.f(): ", i)
}

func main() {
	var a interface{} = 99
	fmt.Printf("type of a :%T\n", a)
	a = 44.09
	fmt.Printf("type of a :%T\n", a)
	a = "hello"

	// command ok 断言
	if v, ok := a.(string); ok {
		fmt.Printf("断言 type of a :%T\n", v)
		v = v + " world"
		fmt.Printf("v :%s\n", v)
	}

	// command ok 断言
	if v, ok := a.(int); ok {
		fmt.Printf("断言 type of a :%T\n", v)
		v++
		fmt.Printf("v :%d\n", v)
	}

	// switch 断言
	switch str := a.(type) {
	case string:
		fmt.Println("switch 断言 a.(T) = string:", str)
	case int:
		fmt.Println("switch 断言 a.(T) = int:", str)

	default:
		fmt.Println("switch 断言 a.(T) 失败")
	}

	var aa A = A{10}
	aa.f()

	var bb B = A{11}
	bb.f()

	var ii I = 10
	ii.f()

	var bbb B = I(101)
	bbb.f()

}
