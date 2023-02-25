package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	length := 10
	result := RandString(length)
	assert.Len(t, result, length)
}

func BenchmarkGenerateShortURL(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RandString(10)
	}
}
