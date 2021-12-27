package artgen

import (
	"fmt"
	"math/rand"
)

type Configuration struct {
	title      string // randomly-generated filename
	Workspace  string // directory to write images to - tries /dev/shm and /tmp before defaulting to working directory
	Format     Format
	Resolution Resolution
	DPI        int  // no native support but can be used to calculate desired size of gg.Context
	Quality    int  // ignored by PNG
	Iterations int  // number of iterations (random polygon generation step)
	WriteFile  bool // used to prepare/generate without writing to disk to store in memory until needed
}

func NewPNGConfiguration(workspace ...string) Configuration {
	return Configuration{
		title:      fmt.Sprintf("%08X", rand.Uint32()),
		Format:     PNG,
		Resolution: DEFAULT,
		DPI:        72,
		Quality:    100,
		Iterations: rand.Intn(7) + 1,
		WriteFile:  true,
		Workspace: func() string {
			if len(workspace) > 0 {
				if workspace[0][0] == '/' {
					return workspace[0]
				}
				return workspace[0]
			}
			return defaultWorkspace
		}(),
	}
}

func NewJPGConfiguration(workspace ...string) Configuration {
	return Configuration{
		title:      fmt.Sprintf("%08X", rand.Uint32()),
		Format:     JPG,
		Resolution: DEFAULT,
		DPI:        72,
		Quality:    80,
		Iterations: rand.Intn(7) + 1,
		WriteFile:  true,
		Workspace: func() string {
			if len(workspace) > 0 {
				if workspace[0][0] == '/' {
					return workspace[0]
				}
				return workspace[0]
			}
			return defaultWorkspace
		}(),
	}
}

type Format int

const (
	PNG Format = iota
	JPG
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

	// print-specific dimensions
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
