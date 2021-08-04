package FactoryPattern

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()  //根据不同的工厂接口，创建不同的工厂
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var factory OperatorFactory
	factory = PlusOperatorFactory{}
	require.Equal(t, 3, compute(factory, 1, 2))

	factory = MinusOperatorFactory{}
	require.Equal(t, 4, compute(factory, 5,1))
}



