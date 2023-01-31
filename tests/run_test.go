package tests

import (
	"github.com/totallygamerjet/pgl"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

var test_suite = map[string]func(t *testing.T){
	"hello_triangle":      hello_triangle,
	"hello_indexing":      hello_indexing,
	"hello_interpolation": hello_interpolation,
	"blend_test":          blend_test,
	"stencil_test":        stencil_test,
	//	{ "stencil_test", stencil_test },
	//	{ "primitives_test", primitives_test },
	//	{ "zbuf_depthoff", zbuf_test },
	//	{ "zbuf_depthon", zbuf_test, 1 },
	//	{ "zbuf_depthon_greater", zbuf_test, 2 },
	//	{ "zbuf_depthon_fliprange", zbuf_test, 3 },
	//
	//	{ "test_edges", test_edges },
	//
	//	{ "texture2D_nearest", test_texturing, 0 },
	//	{ "texture2D_linear", test_texturing, 1 },
	//	{ "texture2D_repeat", test_texturing, 2 },
	//	{ "texture2D_clamp2edge", test_texturing, 3 },
	//	{ "texture2D_mirroredrepeat", test_texturing, 4 },
	//
	//	{ "texrect_nearest", test_texturing, 5 },
	//	{ "texrect_linear", test_texturing, 6 },
	//	{ "texrect_repeat", test_texturing, 7 },
	//	{ "texrect_clamp2edge", test_texturing, 8 },
	//	{ "texrect_mirroredrepeat", test_texturing, 9 },
	//
	//	{ "texture1D_nearest", test_texture1D, 0 },
	//	{ "texture1D_linear", test_texture1D, 1 },
	//	{ "texture1D_repeat", test_texture1D, 2 },
	//	{ "texture1D_clamp2edge", test_texture1D, 3 },
	//	{ "texture1D_mirroredrepeat", test_texture1D, 4 },
	//
	//	{ "map_vbuffer", mapped_vbuffer },
	//	{ "map_nvbuffer", mapped_nvbuffer },
	//
	//	{ "pglbufferdata", pglbufdata_test },
	//
	//	{ "pglteximage2D", test_pglteximage2D },
	//	{ "pglteximage1D", test_pglteximage1D },
	//
	//	{ "unpack_alignment", test_unpackalignment }
}

const (
	HEIGHT = 640
	WIDTH  = 640
)

var _ image.Image = buffer(nil)

type buffer []pgl.U32

func (b buffer) ColorModel() color.Model {
	return color.RGBAModel //ARGB8888
}

func (b buffer) Bounds() image.Rectangle {
	return image.Rect(0, 0, WIDTH, HEIGHT)
}

func (buf buffer) At(x, y int) color.Color {
	c := buf[x+y*WIDTH]
	return color.NRGBA{uint8(c >> 16), uint8(c >> 8), uint8(c >> 0), uint8(c >> 24)}
}

var bbufpix []pgl.U32
var the_Context pgl.GlContext

const export = false

func Test_Run(t *testing.T) {
	for n, f := range test_suite {
		bbufpix = nil
		if pgl.Init_glContext(&the_Context, &bbufpix, WIDTH, HEIGHT, 32, 0x00FF0000, 0x0000FF00, 0x000000FF, 0xFF000000) == 0 {
			t.Fatal("failed to init context")
		}

		pgl.Set_glContext(&the_Context)

		t.Run(n, func(t2 *testing.T) {
			f(t2)
			if glErr := pgl.GetError(); glErr != pgl.NO_ERROR {
				t.Logf("GL has error: %d", glErr)
			}
			open, err := expected.Open("expected_output/" + n + ".png")
			if err != nil {
				t2.Error(err)
			}
			defer open.Close()
			decode, err := png.Decode(open)
			if err != nil {
				t2.Error(err)
			}
			if export {
				create, err := os.Create(n + ".png")
				if err != nil {
					return
				}
				err = png.Encode(create, buffer(bbufpix))
				if err != nil {
					return
				}
			}
			for i := 0; i < WIDTH; i++ {
				for j := 0; j < HEIGHT; j++ {
					r1, g1, b1, a1 := decode.At(i, j).RGBA()
					r2, g2, b2, a2 := buffer(bbufpix).At(i, j).RGBA()
					if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
						t2.Fatalf("%s is not equal at (%d,%d); (%d,%d,%d,%d) != (%d,%d,%d,%d)", n, i, j, r1, g1, b1, a1, r2, g2, b2, a2)
					}
				}
			}
		})
		pgl.Free_glContext(&the_Context)
	}
}
