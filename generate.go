package artgen

import (
	"image"
	"image/color"
	"log"
	"math/rand"

	"github.com/fogleman/gg"
)

func (p *Painting) SaveFile(img image.Image) {
	if p.format == PNG {
		if err := gg.SavePNG(p.File(), img); err != nil {
			log.Println("Error saving PNG:", err)
		}
	} else {
		if err := gg.SaveJPG(p.File(), img, p.quality); err != nil {
			log.Println("Error saving PNG:", err)
		}
	}
}

func (p *Painting) Generate() {
	dc := gg.NewContext(p.width, p.height)
	grad := randomLinearGradient()
	dc.SetFillStyle(grad)
	dc.MoveTo(0, 0)
	dc.LineTo(0, float64(p.width))
	dc.LineTo(0, float64(p.width))
	dc.LineTo(4000, 0)
	dc.ClosePath()
	dc.Fill()
	dc.Stroke()
	// now iterate and spackle with polygonal noise
	for i := 0; i < p.iterations; i++ {
		dc.SetRGBA255(randomRGBA())
		dc.SetLineWidth(float64(rand.Intn(40-8) + 8))
		dc.DrawRegularPolygon(p.randomPolygon())
		dc.Stroke()
	}
	p.imageContext = dc
	if p.writeToDisk {
		p.SaveFile(dc.Image())
	}
}

func randomLinearGradient() gg.Gradient {
	x0 := float64(rand.Intn(300-50) + 50)
	y0 := float64(rand.Intn(600-200) + 200)
	x1 := float64(rand.Intn(900-500) + 500)
	y1 := float64(rand.Intn(1200))
	grad := gg.NewLinearGradient(x0, y0, x1, y1)
	grad.AddColorStop(float64(rand.Intn(2-1)+1), randomColor()) // colors inside gradient
	grad.AddColorStop(float64(rand.Intn(4-2)+2), randomColor())
	grad.AddColorStop(float64(rand.Intn(8-4)+4), randomColor())
	return grad
}
func randomRGBA() (r, g, b, a int) {
	r = rand.Intn(255)
	g = rand.Intn(255)
	b = rand.Intn(255)
	a = 255
	return r, g, b, a
}

func randomColor() color.RGBA {
	r, g, b, a := randomRGBA()
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

func (p *Painting) randomPolygon() (n int, x, y, r, rotation float64) {
	max, min := p.width, p.width/4
	n = rand.Intn(5-3) + 3 // defines shape (num points) of polygon
	x = float64(rand.Intn(max-min) + min)
	y = float64(rand.Intn(max-min) + min)
	r = float64(rand.Intn(max-min) + min)
	rotation = rand.Float64()
	return n, x, y, r, rotation
}
