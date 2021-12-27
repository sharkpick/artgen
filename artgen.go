package artgen

import (
	"image"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	filename = "%s%s.%s"
)

var defaultWorkspace = func() string {
	if _, err := os.Stat("/dev/shm/"); !os.IsNotExist(err) {
		log.Println("using /dev/shm/ as default workspace")
		return "/dev/shm/"
	} else if _, err = os.Stat("/tmp/"); !os.IsNotExist(err) {
		log.Println("using /tmp/ as default workspace")
		return "/tmp/"
	} else {
		log.Println("Warning: using ./ as default workspace")
		return "./"
	}
}()

type Painting interface {
	File() string
	GetImage() image.Image
	GetQuality() int
	GetFormat() Format
}

func SaveFile(p Painting) {
	if p.GetFormat() == PNG {
		if err := gg.SavePNG(p.File(), p.GetImage()); err != nil {
			log.Println("Error saving PNG:", err)
		}
	} else {
		if err := gg.SaveJPG(p.File(), p.GetImage(), p.GetQuality()); err != nil {
			log.Println("Error saving PNG:", err)
		}
	}
}
