package main

import (
	"fmt"
	"math/big"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

func main() {
	fmt.Println(YiB / ZiB) // "1024"
	ziB, _ := new(big.Int).SetString("1180591620717411303424", 10)
	// 打印 ZiB 和 YiB
	fmt.Println("ziB:", ziB)
}
