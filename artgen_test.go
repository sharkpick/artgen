package artgen

import (
	"testing"
)

func TestRandomJPG(t *testing.T) {
	p := NewRandomJPG()
	p.Generate()
	defer p.Cleanup()
}

func TestRandomPNG(t *testing.T) {
	p := NewRandomPNG()
	p.Generate()
	defer p.Cleanup()
}

func TestConfigureJPG(t *testing.T) {
	p := NewRandomJPG()
	p.Resolution = HIGHEST
	p.Quality = 100
	p.Generate()
	defer p.Cleanup()
}

func BenchmarkGenerateDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewRandomPNG()
		p.Generate()
		p.Cleanup()
	}
}

func BenchmarkGenerateFavicon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewRandomPNG("./test/")
		p.Resolution = FAVICON
		p.Generate()
		defer p.Cleanup()
	}
}
