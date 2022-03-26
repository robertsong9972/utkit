package util

import "testing"

func Test_printTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printTest()
		})
	}
}
