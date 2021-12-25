package artgen

import "math/rand"

type Resolution int

const (
	DEFAULT Resolution = iota // 720p
	HIGHER                    // 1080p
	HIGHEST                   // 4k
	LOWER                     // 360p
	LOWEST                    // 240p
	FAVICON                   // 32x32
	APPLE                     // 120x120
)

type Format int

const (
	PNG Format = iota
	JPG
)

type Configuration struct {
	Format              Format
	Resolution          Resolution
	Quality, Iterations int
	WriteFile           bool
	Workspace           string
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
	}
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
	}
	return w, h
}
