package gotester

import "testing"

func Test_generateID(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{"12345", 12345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateID(); got != tt.want {
				t.Errorf("generateID() = %v, want %v", got, tt.want)
			}
		})
	}
}
