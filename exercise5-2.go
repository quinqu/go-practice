// Modify the Lissajous program to produce images in
// multiple colors by adding more values to palette2 and then
// displaying them by changing the third argument of
// SetColorIndex in some interesting way

// Lissajous generates GIF animations of random Lissajous figures.
// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette2 = []color.Color{color.White, color.Black, color.RGBA{30, 90, 120, 1}, color.RGBA{200, 130, 76, 1}}

const (
	whiteIndex  = 0 // first color in palette2
	blackIndex2 = 1 // next color in palette2
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex2)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
