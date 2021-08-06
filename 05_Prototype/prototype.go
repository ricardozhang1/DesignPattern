package main

import "fmt"

// 1.原型对象需要实现的接口
type Cloneable interface {
	Clone()	Cloneable
}

// 2.原型对象的类
type PrototypeManager struct {
	prototypes	map[string]Cloneable
}

// 3.原型对象操作
func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes: map[string]Cloneable{}}
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

// 用户对象1
type Type1 struct {
	name	string
}

// 深拷贝，另外开辟一块内存空间
func (t *Type1) Clone() Cloneable {
	t1 := *t	// 开辟新地址，存储数据（变量赋值）
	return &t1	// 返回变量的地址
}

// 用户对象2
type Type2 struct {
	age	uint8
}

// 浅拷贝，返回源数据地址
func (t *Type2) Clone() Cloneable {
	return t
}

func main() {
	manager := NewPrototypeManager()
	// 深拷贝
	manager.Set("jz", &Type1{"江州"})
	clone1 := manager.Get("jz")
	clone2 := clone1.Clone()
	if clone1 == clone2 {
		fmt.Println("浅拷贝")
	} else {
		fmt.Println("深拷贝")
	}

	// 浅拷贝
	manager.Set("OK", &Type2{age: 23})
	c1 := manager.Get("OK")
	c2 := c1.Clone()
	if c2 == c1 {
		fmt.Println("浅拷贝，共用同一个内存地址")
	} else {
		fmt.Println("深拷贝，不同的内存地址")
	}
}


