package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var x interface{} = time.Now()
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	fmt.Println(v, t, v.Type(), t.Kind())
}
