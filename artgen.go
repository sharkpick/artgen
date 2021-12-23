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

	defaultWidth   = 1280
	defaultHeight  = 720
	defaultFormat  = PNG
	defaultQuality = 75 // default used by jpg. ignored for png
)

var defaultIterations = (rand.Intn(7) + 1)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Format int

const (
	PNG Format = iota
	JPG
)

type Painting struct {
	title               string
	workspace           string
	width, height       int
	format              Format
	quality, iterations int
	imageContext        *gg.Context
	writeToDisk         bool
}

func NewPainting(useWorkspace ...string) *Painting {
	workspace := func() string {
		if len(useWorkspace) > 0 {
			return useWorkspace[0]
		} else {
			return "/dev/shm/"
		}
	}()
	return &Painting{
		title:       fmt.Sprintf("%08x", rand.Uint32()),
		workspace:   workspace,
		width:       defaultWidth,
		height:      defaultHeight,
		format:      defaultFormat,
		quality:     defaultQuality,
		iterations:  defaultIterations,
		writeToDisk: true,
	}
}

func NewGeneratedPainting(useWorkspace ...string) *Painting {
	workspace := func() string {
		if len(useWorkspace) > 0 {
			return useWorkspace[0]
		} else {
			return "/dev/shm/"
		}
	}()
	p := &Painting{
		title:       fmt.Sprintf("%08x", rand.Uint32()),
		workspace:   workspace,
		width:       defaultWidth,
		height:      defaultHeight,
		format:      defaultFormat,
		quality:     defaultQuality,
		iterations:  defaultIterations,
		writeToDisk: true,
	}
	p.Generate()
	return p
}

func (p *Painting) SetIterations(iterations int) {
	p.iterations = iterations
}

func (p *Painting) SetWidth(width int) {
	p.width = width
}

func (p *Painting) SetHeight(height int) {
	p.height = height
}

func (p *Painting) SetDimensions(width, height int) {
	p.width = width
	p.height = height
}

func (p *Painting) SetFormat(format Format) {
	p.format = format
}

func (p *Painting) GetFormat() Format {
	return p.format
}

func (p *Painting) SetJPGQuality(quality int) {
	p.quality = func() int {
		if quality > 100 {
			return 100
		}
		return quality
	}()
}

func (p *Painting) SetWriteToDisk(write bool) {
	p.writeToDisk = write
}

func (p *Painting) File() string {
	extension := func() string {
		if p.format == PNG {
			return "png"
		} else {
			return "jpg"
		}
	}()
	return fmt.Sprintf(filename, p.workspace, p.title, extension)
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
