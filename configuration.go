package artgen

import (
	"math/rand"

	"gonum.org/v1/plot/font"
)

const (
	DefaultDPI = 72
)

type Resolution int

const (
	DEFAULT Resolution = iota // 720p
	HIGHER                    // 1080p
	HIGHEST                   // 4k
	LOWER                     // 360p
	LOWEST                    // 240p
	FAVICON                   // 32x32
	APPLE                     // 120x120

	// walgreens photo-specific dimensions
	PRINT4x6
	PRINT5x7
	PRINT8x10
	PRINTWALLET
	POSTER11x14
	POSTER12x18
	POSTER16x20
	POSTER20x30
	POSTER24x36
	BANNER2x6
	BANNER2x8
	PHOTOGIFT
	CARD4x8
	CARD5x7
	CARDFOLDED5x7
	POSTCARD425x6
	POSTCARD5x7
	PHOTOBOOK
	BRAGBOOK
	NOTEBOOK
	NOTEPAD
	STICKERS
	CANVAS8x10
	CANVAS11x14
	CANVAS12x12
	CANVAS16x20
	CANVAS20x24
	CANVAS20x30
)

type Format int

const (
	PNG Format = iota
	JPG
)

type Configuration struct {
	Format                   Format
	Resolution               Resolution
	Quality, Iterations, DPI int
	WriteFile                bool
	Workspace                string
}

func NewConfiguration(workspace ...string) Configuration {
	c := DefaultConfiguration()
	if len(workspace) > 0 {
		c.Workspace = workspace[0]
	}
	return c
}

func DefaultConfiguration() Configuration {
	return Configuration{
		Format:     PNG,
		Resolution: DEFAULT,
		Quality:    75,
		Iterations: rand.Intn(7) + 1,
		Workspace:  "/dev/shm/",
		WriteFile:  true,
		DPI:        DefaultDPI,
	}
}

func DefaultJPGConfiguration() Configuration {
	return Configuration{
		Format:     JPG,
		Resolution: DEFAULT,
		Quality:    75,
		Iterations: rand.Intn(7) + 1,
		Workspace:  "/dev/shm/",
		WriteFile:  true,
		DPI:        DefaultDPI,
	}
}

func (c *Configuration) WidthInches() font.Length {
	return 14

}
func (c *Configuration) HeightInches() font.Length {
	return 11
}
func (c *Configuration) Width() int {
	w, _ := c.GetResolution()
	return w
}

func (c *Configuration) Height() int {
	_, h := c.GetResolution()
	return h
}

func (c *Configuration) GetResolution() (w, h int) {
	switch c.Resolution {
	case DEFAULT:
		w, h = 1280, 720
	case HIGHER:
		w, h = 1920, 1080
	case HIGHEST:
		w, h = 3840, 2160
	case LOWER:
		w, h = 480, 360
	case LOWEST:
		w, h = 320, 240
	case FAVICON:
		w, h = 32, 32
	case APPLE:
		w, h = 120, 120
	// walgreens print resolutions
	case PRINT4x6:
		w, h = 540, 360
	case PRINT5x7:
		w, h = 630, 450
	case PRINT8x10:
		w, h = 900, 720
	case PRINTWALLET:
		w, h = 270, 180
	case POSTER11x14:
		w, h = 1008, 792
	case POSTER12x18:
		w, h = 2682, 1788
	case POSTER16x20:
		w, h = 2980, 2384
	case POSTER20x30:
		w, h = 4470, 2980
	case POSTER24x36:
		w, h = 5400, 3600
	case BANNER2x6, BANNER2x8:
		w, h = 1800, 1440
	case PHOTOGIFT:
		w, h = 900, 600
	case CARD4x8:
		w, h = 640, 426
	case CARD5x7:
		w, h = 640, 480
	case CARDFOLDED5x7:
		w, h = 840, 600
	case POSTCARD425x6:
		w, h = 720, 480
	case POSTCARD5x7:
		w, h = 840, 600
	case PHOTOBOOK:
		w, h = 1280, 1024
	case BRAGBOOK:
		w, h = 580, 535
	case NOTEBOOK:
		w, h = 1350, 1000
	case NOTEPAD:
		w, h = 370, 370
	case STICKERS:
		w, h = 370, 370
	case CANVAS8x10:
		w, h = 710, 768
	case CANVAS11x14:
		w, h = 994, 781
	case CANVAS12x12:
		w, h = 852, 852
	case CANVAS16x20:
		w, h = 1420, 1136
	case CANVAS20x24:
		w, h = 1740, 1420
	case CANVAS20x30:
		w, h = 3000, 2000
	}
	return w, h
}
