package model

import (
	"math"

	"github.com/flywave/go3d/vec3"
)

const (
	NOPREFERENCE int32 = 0
	STOREINLINE  int32 = 1
	EXTERNALFILE int32 = 2

	NODELETE      int32 = 0
	USENEWDELETE  int32 = 1
	USEMALLOCFREE int32 = 2

	BOTTOMLEFT int32  = 0
	TOPLEFT    int32  = 1
	IMAGET     string = "osg::Image"
)

type Image struct {
	BufferData
	FileName              string
	WriteHint             int32
	Origin                int32
	S                     int32
	T                     int32
	R                     int32
	RowLength             int32
	InternalTextureFormat int32
	PixelFormat           int32
	DataType              int32
	Packing               int32
	PixelAspectRatio      float32

	AllocationMode int32
	Data           []uint8
}

func NewImage() *Image {
	b := NewBufferData()
	b.Type = IMAGET
	return &Image{BufferData: *b, S: 0, T: 0, R: 0, RowLength: 0, InternalTextureFormat: 0, PixelFormat: 0, DataType: 0, Packing: 4, PixelAspectRatio: 1, AllocationMode: USENEWDELETE}
}

func IsPackedType(t int) bool {
	switch t {
	case (GLUNSIGNEDBYTE332):
	case (GLUNSIGNEDBYTE233REV):
	case (GLUNSIGNEDSHORT565):
	case (GLUNSIGNEDSHORT565REV):
	case (GLUNSIGNEDSHORT4444):
	case (GLUNSIGNEDSHORT4444REV):
	case (GLUNSIGNEDSHORT5551):
	case (GLUNSIGNEDSHORT1555REV):
	case (GLUNSIGNEDINT8888):
	case (GLUNSIGNEDINT8888REV):
	case (GLUNSIGNEDINT1010102):
	case (GLUNSIGNEDINT2101010REV):
		return true
	}
	return false
}

func ComputePixelFormat(formt int) int {
	switch formt {
	case (GLALPHA16FARB):
	case (GLALPHA32FARB):
		return GLALPHA

	case (GLLUMINANCE16FARB):
	case (GLLUMINANCE32FARB):
		return GLLUMINANCE

	case (GLINTENSITY16FARB):
	case (GLINTENSITY32FARB):
		return GLINTENSITY

	case (GLLUMINANCEALPHA16FARB):
	case (GLLUMINANCEALPHA32FARB):
		return GLLUMINANCEALPHA

	case (GLRGB32FARB):
	case (GLRGB16FARB):
		return GLRGB

	case (GLRGBA8):
	case (GLRGBA16):
	case (GLRGBA32FARB):
	case (GLRGBA16FARB):
		return GLRGBA

	case (GLALPHA8IEXT):
	case (GLALPHA16IEXT):
	case (GLALPHA32IEXT):
	case (GLALPHA8UIEXT):
	case (GLALPHA16UIEXT):
	case (GLALPHA32UIEXT):
		return GLALPHAINTEGEREXT

	case (GLLUMINANCE8IEXT):
	case (GLLUMINANCE16IEXT):
	case (GLLUMINANCE32IEXT):
	case (GLLUMINANCE8UIEXT):
	case (GLLUMINANCE16UIEXT):
	case (GLLUMINANCE32UIEXT):
		return GLLUMINANCEINTEGEREXT

	case (GLINTENSITY8IEXT):
	case (GLINTENSITY16IEXT):
	case (GLINTENSITY32IEXT):
	case (GLINTENSITY8UIEXT):
	case (GLINTENSITY16UIEXT):
	case (GLINTENSITY32UIEXT):
		return GLLUMINANCEINTEGEREXT

	case (GLLUMINANCEALPHA8IEXT):
	case (GLLUMINANCEALPHA16IEXT):
	case (GLLUMINANCEALPHA32IEXT):
	case (GLLUMINANCEALPHA8UIEXT):
	case (GLLUMINANCEALPHA16UIEXT):
	case (GLLUMINANCEALPHA32UIEXT):
		return GLLUMINANCEALPHAINTEGEREXT

	case (GLRGB32IEXT):
	case (GLRGB16IEXT):
	case (GLRGB8IEXT):
	case (GLRGB32UIEXT):
	case (GLRGB16UIEXT):
	case (GLRGB8UIEXT):
		return GLRGBINTEGEREXT

	case (GLRGBA32IEXT):
	case (GLRGBA16IEXT):
	case (GLRGBA8IEXT):
	case (GLRGBA32UIEXT):
	case (GLRGBA16UIEXT):
	case (GLRGBA8UIEXT):
		return GLRGBAINTEGEREXT

	case (GLDEPTHCOMPONENT16):
	case (GLDEPTHCOMPONENT24):
	case (GLDEPTHCOMPONENT32):
	case (GLDEPTHCOMPONENT32F):
	case (GLDEPTHCOMPONENT32FNV):
		return GLDEPTHCOMPONENT
	}
	return formt
}

func ComputeNumComponents(pixelFormat int) uint {
	switch pixelFormat {
	case (GLCOMPRESSEDRGBS3TCDXT1EXT):
		return 3
	case (GLCOMPRESSEDRGBAS3TCDXT1EXT):
		return 4
	case (GLCOMPRESSEDRGBAS3TCDXT3EXT):
		return 4
	case (GLCOMPRESSEDRGBAS3TCDXT5EXT):
		return 4
	case (GLCOMPRESSEDSIGNEDREDRGTC1EXT):
		return 1
	case (GLCOMPRESSEDREDRGTC1EXT):
		return 1
	case (GLCOMPRESSEDSIGNEDREDGREENRGTC2EXT):
		return 2
	case (GLCOMPRESSEDREDGREENRGTC2EXT):
		return 2
	case (GLCOMPRESSEDRGBPVRTC4BPPV1IMG):
		return 3
	case (GLCOMPRESSEDRGBPVRTC2BPPV1IMG):
		return 3
	case (GLCOMPRESSEDRGBAPVRTC4BPPV1IMG):
		return 4
	case (GLCOMPRESSEDRGBAPVRTC2BPPV1IMG):
		return 4
	case (GLETC1RGB8OES):
		return 3
	case (GLCOMPRESSEDRGB8ETC2):
		return 3
	case (GLCOMPRESSEDSRGB8ETC2):
		return 3
	case (GLCOMPRESSEDRGB8PUNCHTHROUGHALPHA1ETC2):
		return 4
	case (GLCOMPRESSEDSRGB8PUNCHTHROUGHALPHA1ETC2):
		return 4
	case (GLCOMPRESSEDRGBA8ETC2EAC):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ETC2EAC):
		return 4
	case (GLCOMPRESSEDR11EAC):
		return 1
	case (GLCOMPRESSEDSIGNEDR11EAC):
		return 1
	case (GLCOMPRESSEDRG11EAC):
		return 2
	case (GLCOMPRESSEDSIGNEDRG11EAC):
		return 2
	case (GLCOLORINDEX):
		return 1
	case (GLSTENCILINDEX):
		return 1
	case (GLDEPTHCOMPONENT):
		return 1
	case (GLDEPTHCOMPONENT16):
		return 1
	case (GLDEPTHCOMPONENT24):
		return 1
	case (GLDEPTHCOMPONENT32):
		return 1
	case (GLDEPTHCOMPONENT32F):
		return 1
	case (GLDEPTHCOMPONENT32FNV):
		return 1
	case (GLRED):
		return 1
	case (GLGREEN):
		return 1
	case (GLBLUE):
		return 1
	case (GLALPHA):
		return 1
	case (GLALPHA8IEXT):
		return 1
	case (GLALPHA8UIEXT):
		return 1
	case (GLALPHA16IEXT):
		return 1
	case (GLALPHA16UIEXT):
		return 1
	case (GLALPHA32IEXT):
		return 1
	case (GLALPHA32UIEXT):
		return 1
	case (GLALPHA16FARB):
		return 1
	case (GLALPHA32FARB):
		return 1
	case (GLR32F):
		return 1
	case (GLRG):
		return 2
	case (GLRG32F):
		return 2
	case (GLRGB):
		return 3
	case (GLBGR):
		return 3
	case (GLRGB8IEXT):
		return 3
	case (GLRGB8UIEXT):
		return 3
	case (GLRGB16IEXT):
		return 3
	case (GLRGB16UIEXT):
		return 3
	case (GLRGB32IEXT):
		return 3
	case (GLRGB32UIEXT):
		return 3
	case (GLRGB16FARB):
		return 3
	case (GLRGB32FARB):
		return 3
	case (GLRGBA16FARB):
		return 4
	case (GLRGBA32FARB):
		return 4
	case (GLRGBA):
		return 4
	case (GLBGRA):
		return 4
	case (GLRGBA8):
		return 4
	case (GLLUMINANCE):
		return 1
	case (GLLUMINANCE4):
		return 1
	case (GLLUMINANCE8):
		return 1
	case (GLLUMINANCE12):
		return 1
	case (GLLUMINANCE16):
		return 1
	case (GLLUMINANCE8IEXT):
		return 1
	case (GLLUMINANCE8UIEXT):
		return 1
	case (GLLUMINANCE16IEXT):
		return 1
	case (GLLUMINANCE16UIEXT):
		return 1
	case (GLLUMINANCE32IEXT):
		return 1
	case (GLLUMINANCE32UIEXT):
		return 1
	case (GLLUMINANCE16FARB):
		return 1
	case (GLLUMINANCE32FARB):
		return 1
	case (GLLUMINANCE4ALPHA4):
		return 2
	case (GLLUMINANCE6ALPHA2):
		return 2
	case (GLLUMINANCE8ALPHA8):
		return 2
	case (GLLUMINANCE12ALPHA4):
		return 2
	case (GLLUMINANCE12ALPHA12):
		return 2
	case (GLLUMINANCE16ALPHA16):
		return 2
	case (GLINTENSITY):
		return 1
	case (GLINTENSITY4):
		return 1
	case (GLINTENSITY8):
		return 1
	case (GLINTENSITY12):
		return 1
	case (GLINTENSITY16):
		return 1
	case (GLINTENSITY8UIEXT):
		return 1
	case (GLINTENSITY8IEXT):
		return 1
	case (GLINTENSITY16IEXT):
		return 1
	case (GLINTENSITY16UIEXT):
		return 1
	case (GLINTENSITY32IEXT):
		return 1
	case (GLINTENSITY32UIEXT):
		return 1
	case (GLINTENSITY16FARB):
		return 1
	case (GLINTENSITY32FARB):
		return 1
	case (GLLUMINANCEALPHA):
		return 2
	case (GLLUMINANCEALPHA8IEXT):
		return 2
	case (GLLUMINANCEALPHA8UIEXT):
		return 2
	case (GLLUMINANCEALPHA16IEXT):
		return 2
	case (GLLUMINANCEALPHA16UIEXT):
		return 2
	case (GLLUMINANCEALPHA32IEXT):
		return 2
	case (GLLUMINANCEALPHA32UIEXT):
		return 2
	case (GLLUMINANCEALPHA16FARB):
		return 2
	case (GLLUMINANCEALPHA32FARB):
		return 2
	case (GLHILONV):
		return 2
	case (GLDSDTNV):
		return 2
	case (GLDSDTMAGNV):
		return 3
	case (GLDSDTMAGVIBNV):
		return 4
	case (GLREDINTEGEREXT):
		return 1
	case (GLGREENINTEGEREXT):
		return 1
	case (GLBLUEINTEGEREXT):
		return 1
	case (GLALPHAINTEGEREXT):
		return 1
	case (GLRGBINTEGEREXT):
		return 3
	case (GLRGBAINTEGEREXT):
		return 4
	case (GLBGRINTEGEREXT):
		return 3
	case (GLBGRAINTEGEREXT):
		return 4
	case (GLLUMINANCEINTEGEREXT):
		return 1
	case (GLLUMINANCEALPHAINTEGEREXT):
		return 2
	case (GLCOMPRESSEDRGBAASTC4x4KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC5x4KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC5x5KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC6x5KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC6x6KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC8x5KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC8x6KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC8x8KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC10x5KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC10x6KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC10x8KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC10x10KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC12x10KHR):
		return 4
	case (GLCOMPRESSEDRGBAASTC12x12KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC4x4KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC5x4KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC5x5KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC6x5KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC6x6KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x5KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x6KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x8KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x5KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x6KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x8KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x10KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC12x10KHR):
		return 4
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC12x12KHR):
		return 4
	}
	return 0
}

func ComputePixelSizeInBits(format int, t int) uint {
	switch format {
	case (GLCOMPRESSEDRGBS3TCDXT1EXT):
		return 4
	case (GLCOMPRESSEDRGBAS3TCDXT1EXT):
		return 4
	case (GLCOMPRESSEDRGBAS3TCDXT3EXT):
		return 8
	case (GLCOMPRESSEDRGBAS3TCDXT5EXT):
		return 8
	case (GLCOMPRESSEDSIGNEDREDRGTC1EXT):
		return 4
	case (GLCOMPRESSEDREDRGTC1EXT):
		return 4
	case (GLCOMPRESSEDSIGNEDREDGREENRGTC2EXT):
		return 8
	case (GLCOMPRESSEDREDGREENRGTC2EXT):
		return 8
	case (GLCOMPRESSEDRGBPVRTC4BPPV1IMG):
		return 4
	case (GLCOMPRESSEDRGBPVRTC2BPPV1IMG):
		return 2
	case (GLCOMPRESSEDRGBAPVRTC4BPPV1IMG):
		return 4
	case (GLCOMPRESSEDRGBAPVRTC2BPPV1IMG):
		return 2
	case (GLETC1RGB8OES):
		return 4
	case (GLCOMPRESSEDRGB8ETC2):
		return 4
	case (GLCOMPRESSEDSRGB8ETC2):
		return 4
	case (GLCOMPRESSEDRGB8PUNCHTHROUGHALPHA1ETC2):
		return 4
	case (GLCOMPRESSEDSRGB8PUNCHTHROUGHALPHA1ETC2):
		return 4
	case (GLCOMPRESSEDRGBA8ETC2EAC):
		return 8
	case (GLCOMPRESSEDSRGB8ALPHA8ETC2EAC):
		return 8
	case (GLCOMPRESSEDR11EAC):
		return 4
	case (GLCOMPRESSEDSIGNEDR11EAC):
		return 4
	case (GLCOMPRESSEDRG11EAC):
		return 8
	case (GLCOMPRESSEDSIGNEDRG11EAC):
		return 8
	}

	switch format {
	case (GLCOMPRESSEDALPHA):
	case (GLCOMPRESSEDLUMINANCE):
	case (GLCOMPRESSEDLUMINANCEALPHA):
	case (GLCOMPRESSEDINTENSITY):
	case (GLCOMPRESSEDRGB):
	case (GLCOMPRESSEDRGBA):
		return 0

	}
	switch format {
	case (GLCOMPRESSEDRGBAASTC4x4KHR):
	case (GLCOMPRESSEDRGBAASTC5x4KHR):
	case (GLCOMPRESSEDRGBAASTC5x5KHR):
	case (GLCOMPRESSEDRGBAASTC6x5KHR):
	case (GLCOMPRESSEDRGBAASTC6x6KHR):
	case (GLCOMPRESSEDRGBAASTC8x5KHR):
	case (GLCOMPRESSEDRGBAASTC8x6KHR):
	case (GLCOMPRESSEDRGBAASTC8x8KHR):
	case (GLCOMPRESSEDRGBAASTC10x5KHR):
	case (GLCOMPRESSEDRGBAASTC10x6KHR):
	case (GLCOMPRESSEDRGBAASTC10x8KHR):
	case (GLCOMPRESSEDRGBAASTC10x10KHR):
	case (GLCOMPRESSEDRGBAASTC12x10KHR):
	case (GLCOMPRESSEDRGBAASTC12x12KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC4x4KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC5x4KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC5x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC6x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC6x6KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x6KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x8KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x6KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x8KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x10KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC12x10KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC12x12KHR):
		{
			footprint := ComputeBlockFootprint(format)
			pixelsPerBlock := footprint[0] * footprint[1]
			bitsPerBlock := ComputeBlockSize(format, 0) // 16 x 8 = 128
			bitsPerPixel := bitsPerBlock / uint(pixelsPerBlock)
			if bitsPerBlock == bitsPerPixel*uint(pixelsPerBlock) {
				return bitsPerPixel
			}
			return 0
		}
		return 0
	}

	switch format {
	case (GLLUMINANCE4):
		return 4
	case (GLLUMINANCE8):
		return 8
	case (GLLUMINANCE12):
		return 12
	case (GLLUMINANCE16):
		return 16
	case (GLLUMINANCE4ALPHA4):
		return 8
	case (GLLUMINANCE6ALPHA2):
		return 8
	case (GLLUMINANCE8ALPHA8):
		return 16
	case (GLLUMINANCE12ALPHA4):
		return 16
	case (GLLUMINANCE12ALPHA12):
		return 24
	case (GLLUMINANCE16ALPHA16):
		return 32
	case (GLINTENSITY4):
		return 4
	case (GLINTENSITY8):
		return 8
	case (GLINTENSITY12):
		return 12
	case (GLINTENSITY16):
		return 16
	}
	return 0
}

func ComputeBlockFootprint(pixelFormat int) vec3.T {
	switch pixelFormat {
	case (GLCOMPRESSEDRGBS3TCDXT1EXT):
	case (GLCOMPRESSEDRGBAS3TCDXT1EXT):
	case (GLCOMPRESSEDRGBAS3TCDXT3EXT):
	case (GLCOMPRESSEDRGBAS3TCDXT5EXT):
		return vec3.T{4, 4, 4}

	case (GLCOMPRESSEDSIGNEDREDRGTC1EXT):
	case (GLCOMPRESSEDREDRGTC1EXT):
	case (GLCOMPRESSEDSIGNEDREDGREENRGTC2EXT):
	case (GLCOMPRESSEDREDGREENRGTC2EXT):
	case (GLCOMPRESSEDRGBPVRTC4BPPV1IMG):
	case (GLCOMPRESSEDRGBAPVRTC4BPPV1IMG):
	case (GLETC1RGB8OES):
	case (GLCOMPRESSEDRGB8ETC2):
	case (GLCOMPRESSEDSRGB8ETC2):
	case (GLCOMPRESSEDRGB8PUNCHTHROUGHALPHA1ETC2):
	case (GLCOMPRESSEDSRGB8PUNCHTHROUGHALPHA1ETC2):
	case (GLCOMPRESSEDRGBA8ETC2EAC):
	case (GLCOMPRESSEDSRGB8ALPHA8ETC2EAC):
	case (GLCOMPRESSEDR11EAC):
	case (GLCOMPRESSEDSIGNEDR11EAC):
	case (GLCOMPRESSEDRG11EAC):
	case (GLCOMPRESSEDSIGNEDRG11EAC):
		return vec3.T{4, 4, 1}
	case (GLCOMPRESSEDRGBPVRTC2BPPV1IMG):
	case (GLCOMPRESSEDRGBAPVRTC2BPPV1IMG):
		return vec3.T{8, 4, 1} // no 3d texture support in pvrtc at all
	case (GLCOMPRESSEDRGBAASTC4x4KHR):
		return vec3.T{4, 4, 1}
	case (GLCOMPRESSEDRGBAASTC5x4KHR):
		return vec3.T{5, 4, 1}
	case (GLCOMPRESSEDRGBAASTC5x5KHR):
		return vec3.T{5, 5, 1}
	case (GLCOMPRESSEDRGBAASTC6x5KHR):
		return vec3.T{6, 5, 1}
	case (GLCOMPRESSEDRGBAASTC6x6KHR):
		return vec3.T{6, 6, 1}
	case (GLCOMPRESSEDRGBAASTC8x5KHR):
		return vec3.T{8, 5, 1}
	case (GLCOMPRESSEDRGBAASTC8x6KHR):
		return vec3.T{8, 6, 1}
	case (GLCOMPRESSEDRGBAASTC8x8KHR):
		return vec3.T{8, 8, 1}
	case (GLCOMPRESSEDRGBAASTC10x5KHR):
		return vec3.T{10, 5, 1}
	case (GLCOMPRESSEDRGBAASTC10x6KHR):
		return vec3.T{10, 6, 1}
	case (GLCOMPRESSEDRGBAASTC10x8KHR):
		return vec3.T{10, 8, 1}
	case (GLCOMPRESSEDRGBAASTC10x10KHR):
		return vec3.T{10, 10, 1}
	case (GLCOMPRESSEDRGBAASTC12x10KHR):
		return vec3.T{12, 10, 1}
	case (GLCOMPRESSEDRGBAASTC12x12KHR):
		return vec3.T{12, 12, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC4x4KHR):
		return vec3.T{4, 4, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC5x4KHR):
		return vec3.T{5, 4, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC5x5KHR):
		return vec3.T{5, 5, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC6x5KHR):
		return vec3.T{6, 5, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC6x6KHR):
		return vec3.T{6, 6, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x5KHR):
		return vec3.T{8, 5, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x6KHR):
		return vec3.T{8, 6, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x8KHR):
		return vec3.T{8, 8, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x5KHR):
		return vec3.T{10, 5, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x6KHR):
		return vec3.T{10, 6, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x8KHR):
		return vec3.T{10, 8, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x10KHR):
		return vec3.T{10, 10, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC12x10KHR):
		return vec3.T{12, 10, 1}
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC12x12KHR):
		return vec3.T{12, 12, 1}

	}
	return vec3.T{1, 1, 1}
}

func ComputeBlockSize(pixelFormat int, packing uint) uint {
	switch pixelFormat {
	case (GLCOMPRESSEDRGBS3TCDXT1EXT):
	case (GLCOMPRESSEDRGBAS3TCDXT1EXT):
		if packing > 8 {
			return packing
		}
		return 8

	case (GLCOMPRESSEDRGBAS3TCDXT3EXT):
	case (GLCOMPRESSEDRGBAS3TCDXT5EXT):
	case (GLCOMPRESSEDRGBPVRTC2BPPV1IMG):
	case (GLCOMPRESSEDRGBAPVRTC2BPPV1IMG):
	case (GLCOMPRESSEDRGBPVRTC4BPPV1IMG):
	case (GLCOMPRESSEDRGBAPVRTC4BPPV1IMG):
	case (GLETC1RGB8OES):
		if packing > 16 {
			return packing
		}
		return 16

	case (GLCOMPRESSEDSIGNEDREDRGTC1EXT):
	case (GLCOMPRESSEDREDRGTC1EXT):
		if packing > 8 {
			return packing
		}
		return 8

	case (GLCOMPRESSEDSIGNEDREDGREENRGTC2EXT):
	case (GLCOMPRESSEDREDGREENRGTC2EXT):
		if packing > 16 {
			return packing
		}
		return 16

	case (GLCOMPRESSEDRGB8ETC2):
	case (GLCOMPRESSEDSRGB8ETC2):
	case (GLCOMPRESSEDRGB8PUNCHTHROUGHALPHA1ETC2):
	case (GLCOMPRESSEDSRGB8PUNCHTHROUGHALPHA1ETC2):
	case (GLCOMPRESSEDR11EAC):
	case (GLCOMPRESSEDSIGNEDR11EAC):
		if packing > 8 {
			return packing
		}
		return 8

	case (GLCOMPRESSEDRGBA8ETC2EAC):
	case (GLCOMPRESSEDSRGB8ALPHA8ETC2EAC):
	case (GLCOMPRESSEDRG11EAC):
	case (GLCOMPRESSEDSIGNEDRG11EAC):
		if packing > 16 {
			return packing
		}
		return 16

	case (GLCOMPRESSEDRGBAASTC4x4KHR):
	case (GLCOMPRESSEDRGBAASTC5x4KHR):
	case (GLCOMPRESSEDRGBAASTC5x5KHR):
	case (GLCOMPRESSEDRGBAASTC6x5KHR):
	case (GLCOMPRESSEDRGBAASTC6x6KHR):
	case (GLCOMPRESSEDRGBAASTC8x5KHR):
	case (GLCOMPRESSEDRGBAASTC8x6KHR):
	case (GLCOMPRESSEDRGBAASTC8x8KHR):
	case (GLCOMPRESSEDRGBAASTC10x5KHR):
	case (GLCOMPRESSEDRGBAASTC10x6KHR):
	case (GLCOMPRESSEDRGBAASTC10x8KHR):
	case (GLCOMPRESSEDRGBAASTC10x10KHR):
	case (GLCOMPRESSEDRGBAASTC12x10KHR):
	case (GLCOMPRESSEDRGBAASTC12x12KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC4x4KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC5x4KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC5x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC6x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC6x6KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x6KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x8KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x6KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x8KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x10KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC12x10KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC12x12KHR):
		if packing > 16 {
			return packing
		}
		return 16

	}
	return packing
}

func ComputeRowWidthInBytes(width int, pixelFormat int,
	t int, packing int) uint {
	pixelSize := ComputePixelSizeInBits(pixelFormat, t)
	widthInBits := width * int(pixelSize)
	var packingInBits int = 8
	if packing != 0 {
		packingInBits = packing * 8
	}
	var i int = 0
	if widthInBits%packingInBits > 0 {
		i = 1
	}
	return uint((widthInBits/packingInBits + i) * packing)
}

func ComputeImageSizeInBytes(width int, height int,
	depth int, pixelFormat int,
	ty int, packing int,
	slicepacking int,
	imagepacking int) uint {
	if width <= 0 || height <= 0 || depth <= 0 {
		return 0
	}

	blockSize := ComputeBlockSize(pixelFormat, 0)
	if blockSize > 0 {
		footprint := ComputeBlockFootprint(pixelFormat)
		width = (width + int(footprint[0]) - 1) / int(footprint[0])
		height = (height + int(footprint[1]) - 1) / int(footprint[1])

		size := int(blockSize) * width
		size = RoudUpToMultiple(size, packing)
		size *= height
		size = RoudUpToMultiple(size, slicepacking)
		size *= depth
		size = RoudUpToMultiple(size, imagepacking)
		return uint(size)
	}

	size := int(ComputeRowWidthInBytes(width, pixelFormat, ty, packing))

	size *= height
	size += slicepacking - 1
	size -= size % slicepacking

	size *= depth
	size += imagepacking - 1
	size -= size % imagepacking

	st := int(ComputeBlockSize(pixelFormat, uint(packing)))
	if size > st {
		return uint(size)
	}
	return uint(st)
}

func ComputeNearestPowerOfTwo(s int, bias float32) int {
	if (s & (s - 1)) != 0 {
		p2 := math.Log(float64(s)) / math.Log(2.0)
		roundedp2 := math.Floor(p2 + float64(bias))
		s = (int)(math.Pow(2.0, roundedp2))
	}
	return s
}

func RoudUpToMultiple(s int, pack int) int {
	if pack < 2 {
		return s
	}
	s += pack - 1
	s -= s % pack
	return s
}

func (img *Image) IsCompressed() bool {
	switch img.PixelFormat {
	case (GLCOMPRESSEDALPHAARB):
	case (GLCOMPRESSEDINTENSITYARB):
	case (GLCOMPRESSEDLUMINANCEALPHAARB):
	case (GLCOMPRESSEDLUMINANCEARB):
	case (GLCOMPRESSEDRGBAARB):
	case (GLCOMPRESSEDRGBARB):
	case (GLCOMPRESSEDRGBS3TCDXT1EXT):
	case (GLCOMPRESSEDRGBAS3TCDXT1EXT):
	case (GLCOMPRESSEDRGBAS3TCDXT3EXT):
	case (GLCOMPRESSEDRGBAS3TCDXT5EXT):
	case (GLCOMPRESSEDSIGNEDREDRGTC1EXT):
	case (GLCOMPRESSEDREDRGTC1EXT):
	case (GLCOMPRESSEDSIGNEDREDGREENRGTC2EXT):
	case (GLCOMPRESSEDREDGREENRGTC2EXT):
	case (GLCOMPRESSEDRGBPVRTC4BPPV1IMG):
	case (GLCOMPRESSEDRGBPVRTC2BPPV1IMG):
	case (GLCOMPRESSEDRGBAPVRTC4BPPV1IMG):
	case (GLCOMPRESSEDRGBAPVRTC2BPPV1IMG):
	case (GLETC1RGB8OES):
	case (GLCOMPRESSEDRGB8ETC2):
	case (GLCOMPRESSEDSRGB8ETC2):
	case (GLCOMPRESSEDRGB8PUNCHTHROUGHALPHA1ETC2):
	case (GLCOMPRESSEDSRGB8PUNCHTHROUGHALPHA1ETC2):
	case (GLCOMPRESSEDRGBA8ETC2EAC):
	case (GLCOMPRESSEDSRGB8ALPHA8ETC2EAC):
	case (GLCOMPRESSEDR11EAC):
	case (GLCOMPRESSEDSIGNEDR11EAC):
	case (GLCOMPRESSEDRG11EAC):
	case (GLCOMPRESSEDSIGNEDRG11EAC):
	case (GLCOMPRESSEDRGBAASTC4x4KHR):
	case (GLCOMPRESSEDRGBAASTC5x4KHR):
	case (GLCOMPRESSEDRGBAASTC5x5KHR):
	case (GLCOMPRESSEDRGBAASTC6x5KHR):
	case (GLCOMPRESSEDRGBAASTC6x6KHR):
	case (GLCOMPRESSEDRGBAASTC8x5KHR):
	case (GLCOMPRESSEDRGBAASTC8x6KHR):
	case (GLCOMPRESSEDRGBAASTC8x8KHR):
	case (GLCOMPRESSEDRGBAASTC10x5KHR):
	case (GLCOMPRESSEDRGBAASTC10x6KHR):
	case (GLCOMPRESSEDRGBAASTC10x8KHR):
	case (GLCOMPRESSEDRGBAASTC10x10KHR):
	case (GLCOMPRESSEDRGBAASTC12x10KHR):
	case (GLCOMPRESSEDRGBAASTC12x12KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC4x4KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC5x4KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC5x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC6x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC6x6KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x6KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC8x8KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x5KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x6KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x8KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC10x10KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC12x10KHR):
	case (GLCOMPRESSEDSRGB8ALPHA8ASTC12x12KHR):
		return true
	}
	return false
}
