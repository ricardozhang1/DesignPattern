package main

import "fmt"

// 1.定义接口
type SimpleAPI interface {
	Say(string) string
}

// 2.实现对应接口
// 对象1
type Chinese struct {}

func (c *Chinese) Say(content string) string {
	fmt.Println("Say Chinese")
	return "cn"
}

// 对象2
type English struct {}

func (e *English) Say(content string) string {
	fmt.Println("Say English")
	return "en"
}

// 3.调用对应的方法
func NewAPI(apiName string) SimpleAPI {
	if apiName == "cn" {
		return &Chinese{}
	} else if apiName == "en" {
		return &English{}
	}
	return nil
}

func main() {
	newAPI := NewAPI("cn")
	res := newAPI.Say("你好")
	fmt.Println("调用结果：", res)

	newAPI = NewAPI("en")
	res = newAPI.Say("hello")
	fmt.Println("调用结果：", res)
}


