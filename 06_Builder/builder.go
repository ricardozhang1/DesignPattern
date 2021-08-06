package main

import "fmt"

// 1.建造者接口
// Builder是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// 2.建造者对象及操作
type Director struct {
	builder Builder	// 建造者的接口
}

// 创建接口
func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) MakeData() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

// 3.建造者实例
// string建造者
type StringBuilder struct {
	result string
}

func (sb *StringBuilder) Part1() {
	sb.result += "1"
}

func (sb *StringBuilder) Part2() {
	sb.result += "2"
}

func (sb *StringBuilder) Part3() {
	sb.result += "3"
}

func (sb *StringBuilder) GetResult() string {
	return sb.result
}

// int建造者
type IntBuilder struct {
	result int64
}

func (sb *IntBuilder) Part1() {
	sb.result = sb.result + 1
}

func (sb *IntBuilder) Part2() {
	sb.result = sb.result + 2
}

func (sb *IntBuilder) Part3() {
	sb.result = sb.result + 3
}

func (sb *IntBuilder) GetResult() int64 {
	return sb.result
}

func main() {
	b := StringBuilder{}
	//b := IntBuilder{}
	sb := NewDirector(&b)
	sb.MakeData()
	fmt.Println("执行效果：", b.result)
}





