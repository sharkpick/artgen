package artgen

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	p := NewPainting()
	p.Generate()
	defer p.Cleanup()
}

func TestWithWorkspace(t *testing.T) {
	p := NewPainting("./")
	if len(p.File()) < 3 {
		t.Fatalf("Error got %d - want >2\n", len(p.File()))
	}
	if p.File()[:2] != "./" {
		t.Fatalf("Error got %s - want ./\n", p.File()[:2])
	}
	p.Generate()
	if _, err := os.Stat(p.File()); err != nil {
		t.Fatalf("Error got %v but expected to find our finished file...\n", err)
	}
	p.Cleanup()
	if _, err := os.Stat(p.File()); err == nil {
		t.Fatalf("Error file wasn't deleted successfully during Cleanup() - %v", err)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewPainting()
		p.Generate()
		p.Cleanup()
	}
}
