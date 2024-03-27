package test

import (
	"go-orm-learn/chapter01"
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
			chapter01.DeleteSingle()
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
			chapter01.DeleteById()
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
			chapter01.DeleteBatch()
		})
	}
}
