package chapter01

import (
	"testing"
)

func TestDeleteSingle(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "DeleteSingle"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteSingle()
		})
	}
}

func TestDeleteById(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "DeleteById"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteById()
		})
	}
}

func TestDeleteByBatch(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "DeleteByBatch"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteBatch()
		})
	}
}
