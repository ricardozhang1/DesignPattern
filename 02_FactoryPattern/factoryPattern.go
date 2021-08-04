package FactoryPattern

// Operator是被封装的实际接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase是Operator接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

//SetA 设置A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

//SetB 设置B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

//PlusOperatorFactory是PlusOperator(加操作)的工厂类
type PlusOperatorFactory struct {}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}

//PlusOperator Operator的实际加法操作
type PlusOperator struct {
	*OperatorBase
}

//Result 获取结果
func (p *PlusOperator) Result() int {
	return p.a + p.b
}

//MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct {}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		&OperatorBase{},
	}
}

//MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

//Result 获取结果
func (m MinusOperator) Result() int {
	return m.a - m.b
}




