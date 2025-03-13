package main

import (
	"testing"
)

// go test -bench=.

func TestGenerateRandomPassword(t *testing.T) {
	for i := range 50 {
		result := len(generateRandomPassword(i))
		if result != i {
			t.Errorf("generateRandomPassword failed for length %v. got: %v", i, result)
		}
	}
}

func BenchmarkGenerateRandomPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateRandomPassword(25)
	}
}
