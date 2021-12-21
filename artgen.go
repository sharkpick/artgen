package artgen

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	filename = "%s%s.png"

	defaultWidth  = 1280
	defaultHeight = 720
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Painting struct {
	title         string
	workspace     string
	width, height int
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
		title:     fmt.Sprintf("%08x", rand.Uint32()),
		workspace: workspace,
		width:     defaultWidth,
		height:    defaultHeight,
	}
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

func (p *Painting) File() string {
	return fmt.Sprintf(filename, p.workspace, p.title)
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
