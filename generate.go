package artgen

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/fogleman/gg"
)

func doFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return doFibonacci(n-1) + doFibonacci(n-2)
}

func (p *Painting) SaveFile() {
	if p.Format == PNG {
		if err := gg.SavePNG(p.File(), p.imageContext.Image()); err != nil {
			log.Println("Error saving PNG:", err)
		}
	} else {
		if err := gg.SaveJPG(p.File(), p.imageContext.Image(), p.Quality); err != nil {
			log.Println("Error saving PNG:", err)
		}
	}
}

func (p *Painting) Generate() {
	dc := gg.NewContext(p.GetResolution())
	grad := randomLinearGradient()
	dc.SetFillStyle(grad)
	dc.MoveTo(0, 0)
	dc.LineTo(0, float64(p.Width()))
	dc.LineTo(0, float64(p.Width()))
	dc.LineTo(float64(p.Width()*8), 0)
	dc.ClosePath()
	dc.Fill()
	dc.Stroke()
	// now iterate and spackle with polygonal noise
	for i := 0; i < p.Iterations; i++ {
		dc.SetRGBA255(randomRGBA())
		dc.SetLineWidth(randomLineWidth(p.Width()))
		dc.DrawRegularPolygon(randomPolygon(p.Width(), p.Height(), i))
		dc.Stroke()
	}
	p.imageContext = dc
	if p.WriteFile {
		p.SaveFile()
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
func randomRGBA(transparent ...bool) (r, g, b, a int) {
	r = rand.Intn(255)
	g = rand.Intn(255)
	b = rand.Intn(255)
	a = func() int {
		if len(transparent) > 0 && transparent[0] {
			return rand.Intn(255-128) + 128
		}
		return 255
	}()
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

func randomPolygon(width, height int, it ...int) (n int, x, y, r, rotation float64) {
	iteration := func() int {
		if len(it) > 0 {
			return it[0] + 1
		}
		return 1
	}()
	n = rand.Intn(7-3) + 3
	max, min := height/16*(iteration*2), height/26*iteration
	r = float64(rand.Intn(max-min) + min)
	max, min = int((float64(width)*.95)-(r/2)), int(r)/2
	x = float64(rand.Intn(max-min) + min)
	max, min = int((float64(height)*.95)-(r/2)), int(r)/2
	y = float64(rand.Intn(max-min) + min)
	rotation = r * 1.5
	return n, x, y, r, rotation
}

func randomLineWidth(width int) float64 {
	min, max := width/128, width/64
	return float64(rand.Intn(max-min) + min)
}
