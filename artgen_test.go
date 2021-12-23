package artgen

import (
	"testing"
)

/*

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

func TestJpg(t *testing.T) {
	p := NewPainting("./")
	p.SetFormat(JPG)
	p.Generate()
	defer p.Cleanup()
}

func TestDefaultSize(t *testing.T) {
	testDir := "./pngs/"
	os.Mkdir(testDir, 0777)
	for i := 0; i < 50; i++ {
		p := NewPainting(testDir)
		p.Generate()
	}
	if f, err := os.ReadDir(testDir); err != nil {
		fmt.Println(err)
	} else {
		var sum uint64 = 0
		for _, e := range f {
			if info, err := e.Info(); err != nil {
				log.Fatalln("Error:", err)
			} else {
				sum += uint64(info.Size())
			}
		}
		avg := sum / uint64(len(f))
		fmt.Println("average png size:", float64(avg)/float64(1000), "kb")
	}
	os.RemoveAll(testDir)
}

func TestDefaultJPGSize(t *testing.T) {
	testDir := "./jpgs/"
	os.Mkdir(testDir, 0777)
	for i := 0; i < 50; i++ {
		p := NewPainting(testDir)
		p.SetFormat(JPG)
		p.Generate()
	}
	if f, err := os.ReadDir(testDir); err != nil {
		fmt.Println(err)
	} else {
		var sum uint64 = 0
		for _, e := range f {
			if info, err := e.Info(); err != nil {
				log.Fatalln("Error:", err)
			} else {
				sum += uint64(info.Size())
			}
		}
		avg := sum / uint64(len(f))
		fmt.Println("average jpg size:", float64(avg)/float64(1000), "kb")
	}
	os.RemoveAll(testDir)
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewPainting()
		p.Generate()
		p.Cleanup()
	}
}
*/
func BenchmarkGenerateDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewPainting("./")
		p.Generate()
		//p.Cleanup()
	}
}

/*
func BenchmarkGenerate1080pPNG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			p := NewPainting("./")
			p.Generate()
			p.SetDimensions(1920, 1080)
			defer p.Cleanup()
		}()
	}
}

func BenchmarkGenerate4kPNG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			p := NewPainting("./")
			p.Generate()
			p.SetDimensions(3840, 2160)
			defer p.Cleanup()
		}()
	}
}

func BenchmarkGenerateFavicon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			p := NewPainting()
			p.Generate()
			p.SetDimensions(32, 32)
			defer p.Cleanup()
		}()
	}
}
*/
