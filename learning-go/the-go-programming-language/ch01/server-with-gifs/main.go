// minimal HTTP reporting server

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	
)

var mu sync.Mutex
var count int

var palette = []color.Color{color.White, color.Black}
// []color.Color{...} is a composite literal (more specifically a slice), which is just compact notation for creating any composite data type in Go
// all pixels in the palette are initialized to be the first color in the palette, which ours is White.

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // second color in palette
)

// the main function connects any URLs with a path beginning with "/" to a handler and starts a server which is listening for requests on port 8000
func main() {
	// when a request arrives, its given to the handeler
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8200", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
	// we add these mu.Lock() and mu.Unlock() so that we do not face a race condition and that only one goroutine accesses a variable at a time
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	// this is also a composite literal (more specifically a struct)
	// anim is a struct of type gif.GIF
	// this declaration creates a struc that sets LoopCount to the value of nframes. all other fields in the struct are set to 0 of their type
	phase := 0.0

	// this is 2 nested for loops.
	for i := 0; i < nframes; i++ {
		// we initialize nframes to be 64, so this outer for loop iterates 64 times, which produces each frame.

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// this rect creates a white palette that is 201x201
		img := image.NewPaletted(rect, palette)
		// we place this palette, which is white and black, onto this white 201x201 rectangle
		
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// this inner for loop generates a new image and sets some of our pixels to black
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		// we append the black pixels to the image using append and we append it to our anim struct. we then also append 80ms delay.
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
	// anim is encoded into GIF format and written to the output stream "out".
	// out is of type io.Writer which lets us write out to destinations
}