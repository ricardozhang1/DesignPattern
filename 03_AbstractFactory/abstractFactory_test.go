package _3_AbstractFactory

import "testing"

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderDetailDAO().SaveOrderDetail()
	factory.CreateOrderMainDAO().SaveOrderMain()
}

func TestRdbFactory(t *testing.T) {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
}

func TestXmlFactory(t *testing.T) {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
}
