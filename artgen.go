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
	workspace = "/dev/shm/"
	filename  = "%s%s.png"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Painting struct {
	title     string
	workspace string
}

func NewPainting(useWorkspace ...string) *Painting {
	workspace := func() string {
		if len(useWorkspace) > 0 {
			return useWorkspace[0]
		}
		return workspace
	}()
	return &Painting{
		title:     fmt.Sprintf("%08x", rand.Uint32()),
		workspace: workspace,
	}
}

func (p *Painting) File() string {
	return fmt.Sprintf(filename, workspace, p.title)
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
