package chap1

import (
	_ "image"
	"image/color"
	_ "image/gif"
	"io"
	_ "math"
	_ "math/rand"
	_ "os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer) {
}
