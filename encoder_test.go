package y4m_test

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/henkman/y4m"
)

func convertRGBAToYCbCr(rgba *image.RGBA, ycbcr *image.YCbCr) {
	o := 0
	bounds := rgba.Bounds()
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			r, g, b, _ := rgba.At(x, y).RGBA()
			y, cb, cr := color.RGBToYCbCr(byte(r), byte(g), byte(b))
			ycbcr.Y[o] = y
			ycbcr.Cb[o] = cb
			ycbcr.Cr[o] = cr
			o++
		}
	}
}

func TestEncoder(t *testing.T) {
	rand.Seed(time.Now().Unix())
	r := image.Rect(0, 0, 320, 240)
	rgba := image.NewRGBA(r)
	img := image.NewYCbCr(r, image.YCbCrSubsampleRatio444)
	fd, err := os.OpenFile("testme.y4m", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0750)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	y4m.WriteHeader(fd, y4m.Header{Width: 320,
		Height:      240,
		FrameRate:   y4m.Ratio{30, 1},
		Interlacing: y4m.Interlacing_Progressive,
		PixelAspect: y4m.Ratio{1, 1},
		ColorSpace:  y4m.ColorSpace_444,
		Comment:     "testme"})
	for i := 0; i < 30*3; i++ {
		c := color.RGBA{byte(rand.Int31n(255)), byte(rand.Int31n(255)), byte(rand.Int31n(255)), 0xFF}
		draw.Draw(rgba, r, &image.Uniform{c}, image.ZP, draw.Src)
		convertRGBAToYCbCr(rgba, img)
		y4m.WriteFrame(fd, img)
	}
}
