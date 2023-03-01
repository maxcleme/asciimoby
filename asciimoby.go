// Package asciimoby contains a function to generate ASCII art of the lovely Moby.
package asciimoby

import (
	"embed"
	"image"
	"image/color"
	"math"

	"github.com/disintegration/imaging"
	"github.com/qeesung/image2ascii/convert"
)

var (
	//go:embed logo.png
	fs embed.FS

	src image.Image
)

func init() {
	f, err := fs.Open("logo.png")
	if err != nil {
		panic(err)
	}
	src, err = imaging.Decode(f)
	if err != nil {
		panic(err)
	}
}

// Generate returns ASCII art of the lovely Moby.
func Generate(options ...func(*Options)) string {
	opts := &Options{Scale: 1, Size: 32}
	for _, opt := range options {
		opt(opts)
	}
	bg := imaging.New(opts.Size, opts.Size, color.Transparent)
	dst := imaging.Resize(src, int(math.Floor(float64(opts.Size)*opts.Scale)), int(math.Floor(float64(opts.Size)*opts.Scale)), imaging.Lanczos)
	dst = imaging.PasteCenter(bg, dst)
	dst = imaging.Rotate(dst, float64(opts.Angle), color.Transparent)
	dst = imaging.CropAnchor(dst, opts.Size, opts.Size, imaging.Center)

	convertOptions := convert.DefaultOptions
	convertOptions.Colored = opts.Color
	convertOptions.FixedHeight = opts.Size
	convertOptions.FixedWidth = opts.Size
	return convert.NewImageConverter().Image2ASCIIString(dst, &convertOptions)
}
