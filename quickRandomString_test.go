package quickRandomString

import (
	"testing"
)

func BenchmarkRandStringBytesMaskImprSrcUnsafe4Symbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrcUnsafe(4)
	}
}

func BenchmarkRandStringBytesMaskImprSrcUnsafe32Symbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrcUnsafe(32)
	}
}
