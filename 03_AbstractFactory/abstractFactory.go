package _3_AbstractFactory

import "fmt"

/*
Factory：抽象工厂
ProductA、ProductB：抽象产品；

Factory1、Factory2：具体工厂
ProductA1、ProductB1、ProductA2、ProductB2：具体产品；
*/

// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory DAO 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// RDBMainDAO 为关系型数据库的OrderMainDAO的实现
type RDBMainDAO struct {}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Println("rdb main save")
}

// RDBDetailDAO 是关系型数据库的OrderDetailMainDAO的实现
type RDBDetailDAO struct {}

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

// RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct {}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// XMLMainDAO XML主订单存储
type XMLMainDAO struct {}

// SaveOrderMain
func (*XMLMainDAO) SaveOrderMain() {
	fmt.Println("xml main save")
}

// XMLDetailDAO XML详细订单存储
type XMLDetailDAO struct {}

// XMLDetailDAO
func (*XMLDetailDAO) SaveOrderDetail()  {
	fmt.Println("xml detail save")
}

// XMLDAOFactory 是RDB 抽象工厂实现
type XMLDAOFactory struct {}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}



