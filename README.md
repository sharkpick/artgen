# artgen

## about

artgen makes ephemeral art for your enjoyment (or not). it makes it easy to generate some random art and then display it to an enduser. It could also be used to make larger art than it does by default (720p) at the cost of more time/resources.

## usage

artgen is simple to use. you can use html templates to serve the image:
```go
import "github.com/squishd/artgen"

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

## samples

artgen uses some randomness to make fun images like the below. 

![Sample 1](sample1.png)
![Sample 2](sample2.png)
![Sample 3](sample3.png)
![Sample 4](sample4.png)