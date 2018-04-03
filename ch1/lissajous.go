package ch1

import (
	"image/color"
	"os"
	"math/rand"
	"time"
	"image/gif"
	"image"
	"math"
	"fmt"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	var outFile, err = os.Create("liss1.gif")
	if err != nil {
		fmt.Fprintf(os.Stderr, "lissajous: %v\n", err)
	}

	lissajous(outFile)

	outFile.Close()
}

func lissajous(out *os.File) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8 // 8 * 10ms
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int(x*size + 0.5), size + int(y*size + 0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

