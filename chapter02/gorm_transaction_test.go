package chapter02

import (
	"testing"
)

func TestContextDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ContextDemo()
		})
	}
}
