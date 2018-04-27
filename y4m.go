package y4m

type Ratio struct {
	Numerator   uint
	Denominator uint
}

type Interlacing byte

const (
	Interlacing_Progressive      = 'p'
	Interlacing_TopFieldFirst    = 't'
	Interlacing_BottomFieldFirst = 'b'
	Interlacing_Mixed            = 'm'
)

type ColorSpace string

const (
	ColorSpace_420jpeg  = "420jpeg"
	ColorSpace_420paldv = "420paldv"
	ColorSpace_420      = "420"
	ColorSpace_422      = "422"
	ColorSpace_444      = "444"
)

type Header struct {
	Width, Height uint
	FrameRate     Ratio
	Interlacing   Interlacing
	PixelAspect   Ratio
	ColorSpace    ColorSpace
	Comment       string
}
