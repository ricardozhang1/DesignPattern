package main

import "fmt"

type Subject interface {
	Do() string
}

type RealSubject struct {}

func (r *RealSubject) Do() string {
	return "执行以太坊智能合约"
}

type ProxySubject struct {
	RealSubject
	money int
}

func (p *ProxySubject) Do() string {
	if p.money > 0 {
		return p.RealSubject.Do()
	} else {
		return "费用不足，请充值"
	}
}

func main() {
	var sub Subject
	data := &ProxySubject{
		RealSubject: RealSubject{},
		money: 10,
	}
	sub = data
	for i:=0; i<13; i++ {
		fmt.Println(sub.Do())
		data.money--
	}
}


