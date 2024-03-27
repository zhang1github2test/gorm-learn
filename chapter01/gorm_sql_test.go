package chapter01

import (
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
			SelectBySql()
		})
	}
}
