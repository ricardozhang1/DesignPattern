package _4_Singleton

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSingleton(t *testing.T) {
	i1 := getInstance(1)
	i2 := getInstance(2)
	fmt.Println(*i1, *i2)
	require.Equal(t, i1, i2)
}


