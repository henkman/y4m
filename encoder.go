package y4m

import (
	"fmt"
	"image"
	"io"
)

func WriteHeader(w io.Writer, h Header) {
	w.Write([]byte("YUV4MPEG2 "))
	fmt.Fprintf(w, "W%d ", h.Width)
	fmt.Fprintf(w, "H%d ", h.Height)
	fmt.Fprintf(w, "F%d:%d ", h.FrameRate.Numerator, h.FrameRate.Denominator)
	if h.Interlacing != 0 {
		fmt.Fprintf(w, "I%c ", h.Interlacing)
	}
	a := h.PixelAspect
	if a.Numerator != 0 && a.Denominator != 0 {
		fmt.Fprintf(w, "A%d:%d ", a.Numerator, a.Denominator)
	}
	if h.ColorSpace != "" {
		fmt.Fprintf(w, "C%s ", h.ColorSpace)
	}
	if h.Comment != "" {
		fmt.Fprintf(w, "X%s ", h.Comment)
	}
	w.Write([]byte{'\n'})
}

func WriteFrame(w io.Writer, f *image.YCbCr) {
	w.Write([]byte("FRAME \n"))
	w.Write(f.Y)
	w.Write(f.Cb)
	w.Write(f.Cr)
}
