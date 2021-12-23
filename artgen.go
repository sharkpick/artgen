package artgen

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
)

const (
	filename = "%s%s.%s"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Painting struct {
	Configuration
	title        string
	imageContext *gg.Context
}

func NewPainting(configuration ...Configuration) *Painting {
	config := func() Configuration {
		if len(configuration) > 0 {
			return configuration[0]
		} else {
			return DefaultConfiguration()
		}
	}()
	return &Painting{
		Configuration: config,
		title:         fmt.Sprintf("%08x", rand.Uint32()),
	}
}

func NewGeneratedPainting(configuration ...Configuration) *Painting {
	config := func() Configuration {
		if len(configuration) > 0 {
			return configuration[0]
		} else {
			tmp := DefaultConfiguration()
			tmp.WriteFile = false
			return tmp
		}
	}()
	p := &Painting{
		Configuration: config,
		title:         fmt.Sprintf("%08x", rand.Uint32()),
	}
	p.Generate()
	return p
}

func (p *Painting) SetIterations(iterations int) {
	p.Iterations = iterations
}

func (p *Painting) SetResolution(resolution Resolution) {
	p.Resolution = resolution
}

func (p *Painting) SetFormat(format Format) {
	p.Format = format
}

func (p *Painting) GetFormat() Format {
	return p.Format
}

func (p *Painting) SetJPGQuality(quality int) {
	p.Quality = func() int {
		if quality > 100 {
			return 100
		}
		return quality
	}()
}

func (p *Painting) SetWriteToDisk(write bool) {
	p.WriteFile = write
}

func (p *Painting) File() string {
	extension := func() string {
		if p.Format == PNG {
			return "png"
		} else {
			return "jpg"
		}
	}()
	return fmt.Sprintf(filename, p.Workspace, p.title, extension)
}

func (p *Painting) Image() string {
	f, err := os.Open(p.File())
	if err != nil {
		log.Fatalln("Fatal Error opening image", p.File(), err)
	}
	defer f.Close()
	stats, err := f.Stat()
	if err != nil {
		log.Fatalln("Error: can't get file size", err)
	}
	buf := make([]byte, stats.Size())
	reader := bufio.NewReader(f)
	_, err = reader.Read(buf)
	if err != nil {
		log.Fatalln("Error reading to buffer", err)
	}
	return base64.StdEncoding.EncodeToString(buf)
}

func (p *Painting) Cleanup() {
	if err := os.Remove(p.File()); err != nil {
		log.Printf("Error cleaning up painting %s - %v\n", p.title, err)
	}
}
