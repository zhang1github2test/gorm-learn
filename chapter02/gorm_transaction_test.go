package chapter02

import (
	"testing"
)

func TestTransactionDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TransactionDemo()
		})
	}

	TransactionDemo2()
}

func TestTransactionDemo2(t *testing.T) {
	TransactionDemo2()
}

func TestTransactionDemo3(t *testing.T) {
	TransactionDemo3()
}
