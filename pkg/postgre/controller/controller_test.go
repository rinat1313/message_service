package controller

import (
	"fmt"
	"testing"
)

func Test_connectDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "Test_connectDB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("start")
			connectDB()
		})
	}
}
