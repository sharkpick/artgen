package artgen

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/fogleman/gg"
)

type WebImage struct {
	Configuration
	image *gg.Context
}

func NewRandomPNG(workspace ...string) *WebImage {
	p := &WebImage{
		Configuration: NewPNGConfiguration(workspace...),
	}
	return p
}

func NewRandomJPG(workspace ...string) *WebImage {
	p := &WebImage{
		Configuration: NewJPGConfiguration(workspace...),
	}
	return p
}

func (w WebImage) File() string {
	extension := func() string {
		if w.GetFormat() == JPG {
			return "jpg"
		}
		return "png"
	}()
	return fmt.Sprintf(filename, w.Workspace, w.title, extension)
}

func (w WebImage) GetFormat() Format {
	return w.Format
}

func (w WebImage) GetImage() image.Image {
	return w.image.Image()
}

func (w WebImage) PixelDimensions() (width, height int) {
	if w.image == nil {
		log.Fatalln("Fatal Error: PixelDimensions() called on nil *gg.Context in a WebImage type")
	}
	return w.image.Width(), w.image.Height()
}

func (w *WebImage) Generate() *gg.Context {
	width, height := w.getDimensionsForContext()
	w.image = gg.NewContext(width, height)
	grad := randomLinearGradient(w.image.Width(), w.image.Height())
	w.image.SetFillStyle(grad)
	w.image.MoveTo(0, 0)
	w.image.LineTo(0, float64(w.image.Width()))
	w.image.LineTo(0, float64(w.image.Width()))
	w.image.LineTo(float64(w.image.Width())*25, 0)
	w.image.ClosePath()
	w.image.Fill()
	w.image.Stroke()
	// now iterate and spackle with polygonal noise
	for i := 0; i < w.Iterations; i++ {
		w.image.SetRGBA255(randomRGBA())
		w.image.SetLineWidth(randomLineWidth(w.image.Width()))
		w.image.DrawRegularPolygon(randomPolygon(w.image.Width(), w.image.Height(), i))
		w.image.Stroke()
	}
	if w.WriteFile {
		SaveFile(w)
	}
	return w.image
}

func (w WebImage) getDimensionsForContext() (width, height int) {
	switch w.Resolution {
	case DEFAULT:
		width, height = 1280, 720
	case HIGHER:
		width, height = 1920, 1080
	case HIGHEST:
		width, height = 3840, 2160
	case LOWER:
		width, height = 480, 360
	case LOWEST:
		width, height = 320, 240
	case FAVICON:
		width, height = 32, 32
	case APPLE:
		width, height = 120, 120
	// walgreens print resolutions
	case PRINT4x6:
		width, height = 540, 360
	case PRINT5x7:
		width, height = 630, 450
	case PRINT8x10:
		width, height = 900, 720
	case PRINTWALLET:
		width, height = 270, 180
	case POSTER11x14:
		width, height = 1008, 792
	case POSTER12x18:
		width, height = 2682, 1788
	case POSTER16x20:
		width, height = 2980, 2384
	case POSTER20x30:
		width, height = 4470, 2980
	case POSTER24x36:
		width, height = 5400, 3600
	case BANNER2x6, BANNER2x8:
		width, height = 1800, 1440
	case PHOTOGIFT:
		width, height = 900, 600
	case CARD4x8:
		width, height = 640, 426
	case CARD5x7:
		width, height = 640, 480
	case CARDFOLDED5x7:
		width, height = 840, 600
	case POSTCARD425x6:
		width, height = 720, 480
	case POSTCARD5x7:
		width, height = 840, 600
	case PHOTOBOOK:
		width, height = 1280, 1024
	case BRAGBOOK:
		width, height = 580, 535
	case NOTEBOOK:
		width, height = 1350, 1000
	case NOTEPAD:
		width, height = 370, 370
	case STICKERS:
		width, height = 370, 370
	case CANVAS8x10:
		width, height = 710, 768
	case CANVAS11x14:
		width, height = 994, 781
	case CANVAS12x12:
		width, height = 852, 852
	case CANVAS16x20:
		width, height = 1420, 1136
	case CANVAS20x24:
		width, height = 1740, 1420
	case CANVAS20x30:
		width, height = 3000, 2000
	}
	return width, height
}

func (w WebImage) GetQuality() int {
	return w.Quality
}

func (w *WebImage) Cleanup() {
	if err := os.Remove(w.File()); err != nil {
		log.Println("Warning in WebImage.Cleanup()", err)
	}
}

func (w *WebImage) Image() string {
	f, err := os.Open(w.File())
	if err != nil {
		log.Fatalln("Fatal Error opening image", w.File(), err)
	}
	defer f.Close()
	stats, err := f.Stat()
	if err != nil {
		log.Fatalln("Fatal Error: can't get file size", w.File(), err)
	}
	buf := make([]byte, stats.Size())
	reader := bufio.NewReader(f)
	_, err = reader.Read(buf)
	if err != nil {
		log.Fatalln("Fatal Error:", err)
	}
	return base64.StdEncoding.EncodeToString(buf)
}
