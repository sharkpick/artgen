# artgen

## about

artgen makes ephemeral art for your enjoyment (or not). it makes it easy to generate some random art and then display it to an enduser. It could also be used to make larger art than it does by default (720p) at the cost of more time/resources. 

work is underway to find a good balance of randomization (measured by beauty and variety of images produced) and speed. current default generation speed on a 1.10GHz 4-core laptop is ~370ms. read the suggestion in the usage section below to see how to use a pool of *artgen.Painting and a worker (or a few) to reduce load times.


## Benchmarks 

The average size of files generated with the default PNG settings is 121KB, the average size with the default JPG settings is 41KB. The other sizes have not been benchmarked for average file size but the tests are written for the other two defaults and could be modified.

```
goos: linux
goarch: amd64
pkg: github.com/sharkpick/artgen
cpu: Intel(R) Pentium(R) CPU N4200 @ 1.10GHz
BenchmarkGenerateDefault-4      	      50	 370613990 ns/op
BenchmarkGenerate1080pPNG-4     	      50	1237053033 ns/op
BenchmarkGenerate4kPNG-4        	      50	4021682155 ns/op
BenchmarkGenerateFavicon-4      	      50	   1811865 ns/op
BenchmarkGenerateAppleIcons-4   	      50	   8669384 ns/op
PASS
ok  	github.com/sharkpick/artgen	288.970s
```

## usage

artgen is simple to use. you can use html templates to serve the image:
```go
import "github.com/sharkpick/artgen"

func handleArtgen(w http.ResponseWriter, r *http.Request) {
    p := artgen.NewPainting() // returns a *Painting, ready to generate
    p.Generate() // generates and saves the image
    defer p.Cleanup() // destroy image after serving
    fmt.Println(p.File()) // shows location of the file
    if t, err := template.ParseFiles(templateFile); err != nil {
        // ... handle error 
    } else {
        t.Execute(w, p)
        // inside your template you can find the base64 encoded image in {{ .Image }}
    }
}
```

or change the workspace and write it to an existing directory of your choice
```go
// generate 10 random images
func main() {
    myNewWorkspace := "./"
    for i := 0; i < 10; i++ {
        p := artgen.NewPainting(myNewWorkspace)
        p.Generate()
        // DO NOT Cleanup() here or your images will also be deleted.
    }
}
```

you can also change the size of the end product. you must set your desired result before running Generate()
```go
func main() {
    width, height := 1920, 1080
    for i := 0; i < 10; i++ {
        p := artgen.NewPainting()
        p.SetWidth(width)
        p.SetHeight(height)
        // p.SetDimensions(width, height) // or do both
        p.Generate()
        defer p.Cleanup()
        // do work
    }
}
```

you can also use JPGs instead of PNGs when speed matters. default quality for artget.JPG is 75% but can be adjusted before running Generate().

```go
func main() {
    p := artgen.NewPainting()
    p.SetFormat(artgen.JPG)
    p.SetQuality(50)
    p.Generate()
    defer p.Cleanup()
}
```

if you find your load times take too long, build a channel of *artgen.Painting objects ready to use. build a channel and then fill it when you start your server. use a worker to keep the channel filled (or as filled as you can) and use a timeout to keep a request from hanging forever when a server is too busy or your channel's buffer is not large enough.
```go
done := make(chan interface{}) // close after server shutdown for graceful shutdown
buffer := make(chan *artgen.Painting, 24) // will hold 24 Generated() images for rapid use
func fillBuffer() {
    go func() {
        for {
            select {
            case <-done:
                return
            case buffer <- artgen.NewGeneratePainting():
            }
        }
        
    }()
}
func handleArtGen(w http.ResponseWriter, r *http.Request) {
    switch {
        case p, ok := <- buffer:
            if !ok {
                log.Panicln("Error: channel empty!")
            }
            p.WriteFile()
            defer p.Cleanup()
            // do work...
        case <-r.Context().Done():
            w.Write([]byte("request timed out"))
    }
}
```


## Samples

artgen uses some randomness to make fun images like the below. 

![Sample 1](sample1.png)
![Sample 2](sample2.png)
![Sample 3](sample3.png)
![Sample 4](sample4.png)