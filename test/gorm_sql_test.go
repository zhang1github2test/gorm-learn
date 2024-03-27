package test

import (
	"go-orm-learn/chapter01"
	"testing"
)

func Test_selectBySql(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "selectBySql",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chapter01.SelectBySql()
		})
	}
}
