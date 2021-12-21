package artgen

import (
	"testing"
)

func TestMain(t *testing.T) {
	p := NewPainting()
	p.Generate()
	defer p.Cleanup()
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewPainting()
		p.Generate()
		p.Cleanup()
	}
}
