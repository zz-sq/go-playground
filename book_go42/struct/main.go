package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string `field:"name"`
	age  int    `field:"age"`
}

func main() {
	user := User{name: "zhangsan", age: 18}
	fmt.Println(user)
	fmt.Println(user.name)
	fmt.Println(user.age)
	fmt.Printf("type of user: %s\n", reflect.TypeOf(user))
	fmt.Printf("type of user.name: %s\n", reflect.TypeOf(user).Field(0).Tag.Get("field"))
	fmt.Printf("type of user.age: %s\n", reflect.TypeOf(user).Field(1).Tag)
}
