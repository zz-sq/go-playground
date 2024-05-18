package main

import (
	"log"
	"runtime"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Close() {
	p.Name = "NewName"
	log.Println(p)
	log.Println("Close")
}

func (p *Person) NewOpen() {
	log.Println("Init")
	runtime.SetFinalizer(p, (*Person).Close)
}

func Tt(p *Person) {
	p.Name = "NewName"
	log.Println(p)
	log.Println("Tt")
}

// 查看内存情况
func Mem(m *runtime.MemStats) {
	runtime.ReadMemStats(m)
	log.Printf("%d Kb\n", m.Alloc/1024)
}

func main() {
	var a, b int = 20, 30 // 声明实际变量
	var ptra *int         // 声明指针变量
	var ptrb *int = &b

	ptra = &a // 指针变量的存储地址

	log.Printf("a  变量的地址是: %x\n", &a)
	log.Printf("b  变量的地址是: %x\n", &b)

	// 指针变量的存储地址
	log.Printf("ptra  变量的存储地址: %x\n", ptra)
	log.Printf("ptrb  变量的存储地址: %x\n", ptrb)

	// 使用指针访问值
	log.Printf("*ptra  变量的值: %d\n", *ptra)
	log.Printf("*ptrb  变量的值: %d\n", *ptrb)

	var m runtime.MemStats
	Mem(&m)

	var p *Person = &Person{Name: "lee", Age: 4}
	p.NewOpen()
	log.Println("Gc完成第一次")
	log.Println("p:", p)
	runtime.GC()
	time.Sleep(time.Second * 5)
	Mem(&m)

	var p1 *Person = &Person{Name: "Goo", Age: 9}
	runtime.SetFinalizer(p1, Tt)
	log.Println("Gc完成第二次")
	time.Sleep(time.Second * 2)
	runtime.GC()
	time.Sleep(time.Second * 2)
	Mem(&m)
}
