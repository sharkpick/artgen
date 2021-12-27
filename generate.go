package artgen

import (
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

func randomLinearGradient(width, height int) gg.Gradient {
	grad := gg.NewLinearGradient(0, 0, float64(width)*1.5, float64(height)*1.5)
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
	if max < min {
		max, min = min+5, max
	}
	r = float64(rand.Intn(max-min) + min)
	max, min = int((float64(width)*.9)-(r/2)), int(r)
	if max <= min {
		max, min = min+5, max
	}
	x = float64(rand.Intn(max-min) + min)
	max, min = int((float64(height)*.9)-(r/2)), int(r)
	if max <= min {
		max, min = min+5, max
	}
	y = float64(rand.Intn(max-min) + min)
	rotation = r * 1.5
	return n, x, y, r, rotation
}

func randomLineWidth(width int) float64 {
	min := func() int {
		min := (width / 128)
		if min == 0 {
			return 1
		}
		return min
	}()
	max := func() int {
		max := (width / 64)
		if max <= 1 {
			return 2
		}
		return max
	}()
	return float64(rand.Intn(max-min) + min)
}
