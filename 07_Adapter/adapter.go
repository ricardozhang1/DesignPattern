package main

import (
	"fmt"
	"strconv"
)

/*
Target：适配的对象，Adapter：适配；Adaptee：被适配。
*/

//1.被适配对象
// 被适配接口
type Adapter interface {
	SpecRequest(int, int) string
}

// 接口载体
type AdapterImpl struct {}

func (ap *AdapterImpl) SpecRequest(a int, b int) string {
	return "目标对象：SpecRequest " + strconv.Itoa(a + b)
}

func NewAdapter() Adapter {
	return &AdapterImpl{}
}

// 2.适配的目标对象
// 适配的目标接口
type Target interface {
	Request(int, int) string
}

type AdapterImplDsc struct {
	Adapter	// 注意：继承的是接口，非对象（结构体）
}

// 接口包装
func (ad *AdapterImplDsc) Request(a, b int) string {
	return ad.SpecRequest(a, b)
}

func NewTarget(ad Adapter) Target {
	return &AdapterImplDsc{ad}
}

func main() {
	ad := NewAdapter()
	target := NewTarget(ad)
	res := target.Request(1, 2)
	fmt.Println(res)
}






