package artgen

import (
	"testing"
)

func TestMain(t *testing.T) {
	config := NewConfiguration("./test/")
	config.Resolution = POSTER11x14
	p := NewPainting(config)
	p.Generate()
	//defer p.Cleanup()
}

/*
func TestWithWorkspace(t *testing.T) {
	config := NewConfiguration("./")
	p := NewPainting(config)
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
	config := NewConfiguration()
	p := NewPainting(config)
	p.SetFormat(JPG)
	p.Generate()
	defer p.Cleanup()
}
*/
/*// the following two tests calculate an average file size for the 2 defaults
func TestDefaultSize(t *testing.T) {
	testDir := "./pngs/"
	os.Mkdir(testDir, 0777)
	for i := 0; i < 50; i++ {
		config := NewConfiguration(testDir)
		p := NewPainting(config)
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
		config := NewConfiguration(testDir)
		p := NewPainting(config)
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
*/

func BenchmarkGenerateDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewPainting()
		p.Generate()
		p.Cleanup()
	}
}

func BenchmarkGenerate1080pPNG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			config := NewConfiguration()
			config.Resolution = HIGHER
			p := NewPainting(config)
			p.Generate()
			defer p.Cleanup()
		}()
	}
}

func BenchmarkGenerate4kPNG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			config := NewConfiguration()
			config.Resolution = HIGHEST
			p := NewPainting(config)
			p.Generate()
			defer p.Cleanup()
		}()
	}
}

func BenchmarkGenerateFavicon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			config := NewConfiguration()
			config.Resolution = FAVICON
			p := NewPainting(config)
			p.Generate()
			defer p.Cleanup()
		}()
	}
}

func BenchmarkGenerateAppleIcons(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			config := NewConfiguration()
			config.Resolution = APPLE
			p := NewPainting(config)
			p.Generate()
			defer p.Cleanup()
		}()
	}
}
