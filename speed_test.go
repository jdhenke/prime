package prime

import "testing"

func BenchmarkSpeed(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		pg := NewGenerator()
		pg.Get(1000000)
	}
}
