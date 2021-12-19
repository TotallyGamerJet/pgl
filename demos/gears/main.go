package main

import (
	"fmt"
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"pgl"
	"unsafe"
)

const M_PI = 3.14159265
const WIDTH = 640
const HEIGHT = 480
const STRIPS_PER_TOOTH = 7
const VERTICES_PER_TOOTH = 34
const GEAR_VERTEX_STRIDE = 6

func sin_cos(x float64, s *float64, c *float64) {
	*s = math.Sin(x)
	*c = math.Cos(x)
	return
}

var window *sdl.Window
var ren *sdl.Renderer
var tex *sdl.Texture
var bbufpix *pgl.U32
var the_context pgl.GlContext
var polygon_mode int

type My_Uniforms struct {
	Mvp_mat        pgl.Mat4
	Normal_mat     pgl.Mat4
	Material_color pgl.Vec3
}
type vertex_strip struct {
	First pgl.GLint
	Count pgl.GLint
}
type GearVertex [6]pgl.GLfloat
type gear struct {
	Vertices  *GearVertex
	Nvertices int
	Strips    *vertex_strip
	Nstrips   int
	Vbo       pgl.GLuint
}

var view_rot [3]pgl.GLfloat = [3]pgl.GLfloat{pgl.GLfloat(20.0), pgl.GLfloat(30.0), pgl.GLfloat(0.0)}
var gear1 *gear
var gear2 *gear
var gear3 *gear
var angle pgl.GLfloat = pgl.GLfloat(0.0)
var ProjectionMatrix [16]pgl.GLfloat
var uniforms My_Uniforms

func vert(v *GearVertex, x pgl.GLfloat, y pgl.GLfloat, z pgl.GLfloat, n [3]pgl.GLfloat) *GearVertex {
	(*(*GearVertex)(unsafe.Add(unsafe.Pointer(v), unsafe.Sizeof(GearVertex{})*0)))[0] = x
	(*(*GearVertex)(unsafe.Add(unsafe.Pointer(v), unsafe.Sizeof(GearVertex{})*0)))[1] = y
	(*(*GearVertex)(unsafe.Add(unsafe.Pointer(v), unsafe.Sizeof(GearVertex{})*0)))[2] = z
	(*(*GearVertex)(unsafe.Add(unsafe.Pointer(v), unsafe.Sizeof(GearVertex{})*0)))[3] = n[0]
	(*(*GearVertex)(unsafe.Add(unsafe.Pointer(v), unsafe.Sizeof(GearVertex{})*0)))[4] = n[1]
	(*(*GearVertex)(unsafe.Add(unsafe.Pointer(v), unsafe.Sizeof(GearVertex{})*0)))[5] = n[2]
	return (*GearVertex)(unsafe.Add(unsafe.Pointer(v), unsafe.Sizeof(GearVertex{})*1))
}
func create_gear(inner_radius pgl.GLfloat, outer_radius pgl.GLfloat, width pgl.GLfloat, teeth pgl.GLint, tooth_depth pgl.GLfloat) *gear {
	var (
		r0        pgl.GLfloat
		r1        pgl.GLfloat
		r2        pgl.GLfloat
		da        pgl.GLfloat
		v         *GearVertex
		gr        *gear
		s         [5]float64
		c         [5]float64
		normal    [3]pgl.GLfloat
		cur_strip int = 0
		i         int
	)
	gr = new(gear)
	if gr == nil {
		return nil
	}
	r0 = inner_radius
	r1 = pgl.GLfloat(float64(outer_radius) - float64(tooth_depth)/2.0)
	r2 = pgl.GLfloat(float64(outer_radius) + float64(tooth_depth)/2.0)
	da = pgl.GLfloat(math.Pi * 2.0 / float64(teeth) / 4.0)
	gr.Nstrips = int(STRIPS_PER_TOOTH * teeth)
	gr.Strips = &make([]vertex_strip, gr.Nstrips)[0]
	gr.Vertices = (*GearVertex)(unsafe.Pointer(&make([]GearVertex, int(VERTICES_PER_TOOTH*teeth))[0][0]))
	v = gr.Vertices
	var sc_val float64
	for i = 0; i < int(teeth); i++ {
		sc_val = float64(i) * 2.0 * math.Pi / float64(teeth)
		sin_cos(sc_val, &s[0], &c[0])
		sin_cos(sc_val+float64(da), &s[1], &c[1])
		sin_cos(sc_val+float64(da*2), &s[2], &c[2])
		sin_cos(sc_val+float64(da*3), &s[3], &c[3])
		sin_cos(sc_val+float64(da*4), &s[4], &c[4])
		type point struct {
			X pgl.GLfloat
			Y pgl.GLfloat
		}
		var p [7]point = [7]point{{X: pgl.GLfloat(float64(r2) * c[1]), Y: pgl.GLfloat(float64(r2) * s[1])}, {X: pgl.GLfloat(float64(r2) * c[2]), Y: pgl.GLfloat(float64(r2) * s[2])}, {X: pgl.GLfloat(float64(r1) * c[0]), Y: pgl.GLfloat(float64(r1) * s[0])}, {X: pgl.GLfloat(float64(r1) * c[3]), Y: pgl.GLfloat(float64(r1) * s[3])}, {X: pgl.GLfloat(float64(r0) * c[0]), Y: pgl.GLfloat(float64(r0) * s[0])}, {X: pgl.GLfloat(float64(r1) * c[4]), Y: pgl.GLfloat(float64(r1) * s[4])}, {X: pgl.GLfloat(float64(r0) * c[4]), Y: pgl.GLfloat(float64(r0) * s[4])}}
		for {
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First = pgl.GLint(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			if true {
				break
			}
		}
		for {
			normal[0] = 0
			normal[1] = 0
			normal[2] = pgl.GLfloat(1.0)
			if true {
				break
			}
		}
		v = vert(v, p[0].X, p[0].Y, pgl.GLfloat(float64(width*(+1))*0.5), normal)
		v = vert(v, p[1].X, p[1].Y, pgl.GLfloat(float64(width*(+1))*0.5), normal)
		v = vert(v, p[2].X, p[2].Y, pgl.GLfloat(float64(width*(+1))*0.5), normal)
		v = vert(v, p[3].X, p[3].Y, pgl.GLfloat(float64(width*(+1))*0.5), normal)
		v = vert(v, p[4].X, p[4].Y, pgl.GLfloat(float64(width*(+1))*0.5), normal)
		v = vert(v, p[5].X, p[5].Y, pgl.GLfloat(float64(width*(+1))*0.5), normal)
		v = vert(v, p[6].X, p[6].Y, pgl.GLfloat(float64(width*(+1))*0.5), normal)
		for {
			var _tmp int = int(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).Count = pgl.GLint(_tmp - int((*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First))
			cur_strip++
			if true {
				break
			}
		}
		for {
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First = pgl.GLint(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			if true {
				break
			}
		}
		for {
			for {
				normal[0] = p[4].Y - p[6].Y
				normal[1] = pgl.GLfloat(float32(-(p[4].X - p[6].X)))
				normal[2] = 0
				if true {
					break
				}
			}
			v = vert(v, p[4].X, p[4].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
			v = vert(v, p[4].X, p[4].Y, pgl.GLfloat(float64(width*1)*0.5), normal)
			v = vert(v, p[6].X, p[6].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
			v = vert(v, p[6].X, p[6].Y, pgl.GLfloat(float64(width*1)*0.5), normal)
			if true {
				break
			}
		}
		for {
			var _tmp int = int(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).Count = pgl.GLint(_tmp - int((*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First))
			cur_strip++
			if true {
				break
			}
		}
		for {
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First = pgl.GLint(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			if true {
				break
			}
		}
		for {
			normal[0] = 0
			normal[1] = 0
			normal[2] = pgl.GLfloat(-1.0)
			if true {
				break
			}
		}
		v = vert(v, p[6].X, p[6].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
		v = vert(v, p[5].X, p[5].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
		v = vert(v, p[4].X, p[4].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
		v = vert(v, p[3].X, p[3].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
		v = vert(v, p[2].X, p[2].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
		v = vert(v, p[1].X, p[1].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
		v = vert(v, p[0].X, p[0].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
		for {
			var _tmp int = int(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).Count = pgl.GLint(_tmp - int((*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First))
			cur_strip++
			if true {
				break
			}
		}
		for {
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First = pgl.GLint(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			if true {
				break
			}
		}
		for {
			for {
				normal[0] = p[0].Y - p[2].Y
				normal[1] = pgl.GLfloat(float32(-(p[0].X - p[2].X)))
				normal[2] = 0
				if true {
					break
				}
			}
			v = vert(v, p[0].X, p[0].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
			v = vert(v, p[0].X, p[0].Y, pgl.GLfloat(float64(width*1)*0.5), normal)
			v = vert(v, p[2].X, p[2].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
			v = vert(v, p[2].X, p[2].Y, pgl.GLfloat(float64(width*1)*0.5), normal)
			if true {
				break
			}
		}
		for {
			var _tmp int = int(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).Count = pgl.GLint(_tmp - int((*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First))
			cur_strip++
			if true {
				break
			}
		}
		for {
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First = pgl.GLint(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			if true {
				break
			}
		}
		for {
			for {
				normal[0] = p[1].Y - p[0].Y
				normal[1] = pgl.GLfloat(float32(-(p[1].X - p[0].X)))
				normal[2] = 0
				if true {
					break
				}
			}
			v = vert(v, p[1].X, p[1].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
			v = vert(v, p[1].X, p[1].Y, pgl.GLfloat(float64(width*1)*0.5), normal)
			v = vert(v, p[0].X, p[0].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
			v = vert(v, p[0].X, p[0].Y, pgl.GLfloat(float64(width*1)*0.5), normal)
			if true {
				break
			}
		}
		for {
			var _tmp int = int(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).Count = pgl.GLint(_tmp - int((*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First))
			cur_strip++
			if true {
				break
			}
		}
		for {
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First = pgl.GLint(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			if true {
				break
			}
		}
		for {
			for {
				normal[0] = p[3].Y - p[1].Y
				normal[1] = pgl.GLfloat(float32(-(p[3].X - p[1].X)))
				normal[2] = 0
				if true {
					break
				}
			}
			v = vert(v, p[3].X, p[3].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
			v = vert(v, p[3].X, p[3].Y, pgl.GLfloat(float64(width*1)*0.5), normal)
			v = vert(v, p[1].X, p[1].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
			v = vert(v, p[1].X, p[1].Y, pgl.GLfloat(float64(width*1)*0.5), normal)
			if true {
				break
			}
		}
		for {
			var _tmp int = int(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).Count = pgl.GLint(_tmp - int((*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First))
			cur_strip++
			if true {
				break
			}
		}
		for {
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First = pgl.GLint(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			if true {
				break
			}
		}
		for {
			for {
				normal[0] = p[5].Y - p[3].Y
				normal[1] = pgl.GLfloat(float32(-(p[5].X - p[3].X)))
				normal[2] = 0
				if true {
					break
				}
			}
			v = vert(v, p[5].X, p[5].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
			v = vert(v, p[5].X, p[5].Y, pgl.GLfloat(float64(width*1)*0.5), normal)
			v = vert(v, p[3].X, p[3].Y, pgl.GLfloat(float64(width*pgl.GLfloat(-1))*0.5), normal)
			v = vert(v, p[3].X, p[3].Y, pgl.GLfloat(float64(width*1)*0.5), normal)
			if true {
				break
			}
		}
		for {
			var _tmp int = int(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
			(*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).Count = pgl.GLint(_tmp - int((*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gr.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(cur_strip)))).First))
			cur_strip++
			if true {
				break
			}
		}
	}
	gr.Nvertices = int(int64(uintptr(unsafe.Pointer(v)) - uintptr(unsafe.Pointer(gr.Vertices))))
	pgl.GlGenBuffers(1, &gr.Vbo)
	pgl.GlBindBuffer(pgl.GLenum(pgl.GL_ARRAY_BUFFER), gr.Vbo)
	pgl.GlBufferData(pgl.GLenum(pgl.GL_ARRAY_BUFFER), pgl.GLsizei(gr.Nvertices*int(unsafe.Sizeof(GearVertex{}))), unsafe.Pointer(gr.Vertices), pgl.GLenum(pgl.GL_STATIC_DRAW))
	return gr
}
func multiply(m *pgl.GLfloat, n *pgl.GLfloat) {
	var (
		tmp    [16]pgl.GLfloat
		row    *pgl.GLfloat
		column *pgl.GLfloat
		//d      div_t
		i int
		j int
	)
	for i = 0; i < 16; i++ {
		tmp[i] = 0
		//d = div(i, 4)
		row = (*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(n), unsafe.Sizeof(pgl.GLfloat(0))*uintptr((i/4)*4)))
		column = (*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*uintptr(i%4)))
		for j = 0; j < 4; j++ {
			tmp[i] += *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(row), unsafe.Sizeof(pgl.GLfloat(0))*uintptr(j))) * *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(column), unsafe.Sizeof(pgl.GLfloat(0))*uintptr(j*4)))
		}
	}
	libc.MemCpy(unsafe.Pointer(m), unsafe.Pointer(&tmp[0]), int(unsafe.Sizeof([16]pgl.GLfloat{})))
}
func rotate(m *pgl.GLfloat, angle pgl.GLfloat, x pgl.GLfloat, y pgl.GLfloat, z pgl.GLfloat) {
	var (
		s float64
		c float64
	)
	sin_cos(float64(angle), &s, &c)
	var r [16]pgl.GLfloat = [16]pgl.GLfloat{pgl.GLfloat(float64(x*x)*(1-c) + c), pgl.GLfloat(float64(y*x)*(1-c) + float64(z)*s), pgl.GLfloat(float64(x*z)*(1-c) - float64(y)*s), 0, pgl.GLfloat(float64(x*y)*(1-c) - float64(z)*s), pgl.GLfloat(float64(y*y)*(1-c) + c), pgl.GLfloat(float64(y*z)*(1-c) + float64(x)*s), 0, pgl.GLfloat(float64(x*z)*(1-c) + float64(y)*s), pgl.GLfloat(float64(y*z)*(1-c) - float64(x)*s), pgl.GLfloat(float64(z*z)*(1-c) + c), 0, 0, 0, 0, 1}
	multiply(m, &r[0])
}
func translate(m *pgl.GLfloat, x pgl.GLfloat, y pgl.GLfloat, z pgl.GLfloat) {
	var t [16]pgl.GLfloat = [16]pgl.GLfloat{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, x, y, z, 1}
	multiply(m, &t[0])
}
func identity(m *pgl.GLfloat) {
	var t [16]pgl.GLfloat = [16]pgl.GLfloat{pgl.GLfloat(1.0), pgl.GLfloat(0.0), pgl.GLfloat(0.0), pgl.GLfloat(0.0), pgl.GLfloat(0.0), pgl.GLfloat(1.0), pgl.GLfloat(0.0), pgl.GLfloat(0.0), pgl.GLfloat(0.0), pgl.GLfloat(0.0), pgl.GLfloat(1.0), pgl.GLfloat(0.0), pgl.GLfloat(0.0), pgl.GLfloat(0.0), pgl.GLfloat(0.0), pgl.GLfloat(1.0)}
	copy(unsafe.Slice(m, 16), t[:])
}
func transpose(m *pgl.GLfloat) {
	var t [16]pgl.GLfloat = [16]pgl.GLfloat{*(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*0)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*4)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*8)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*12)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*1)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*5)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*9)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*13)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*2)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*6)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*10)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*14)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*3)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*7)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*11)), *(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*15))}
	copy(unsafe.Slice(m, 16), t[:])
}

//    GLfloat t[16];
//    identity(t);
//
// Extract and invert the translation part 't'. The inverse of a
// translation matrix can be calculated by negating the translation
// coordinates.
//    t[12] = -m[12]; t[13] = -m[13]; t[14] = -m[14];
//
// Invert the rotation part 'r'. The inverse of a rotation matrix is
// equal to its transpose.
//    m[12] = m[13] = m[14] = 0;
//    transpose(m);
//
//    // inv(m) = inv(r) * inv(t)
//    multiply(m, t);
func invert(m *pgl.GLfloat) {
	var t [16]pgl.GLfloat
	identity(&t[0])
	// Extract and invert the translation part 't'. The inverse of a
	// translation matrix can be calculated by negating the translation
	// coordinates.
	t[12] = -*(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*12))
	t[13] = -*(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*13))
	t[14] = -*(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*14))
	// Invert the rotation part 'r'. The inverse of a rotation matrix is
	// equal to its transpose.
	*(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*12)) = 0
	*(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*13)) = 0
	*(*pgl.GLfloat)(unsafe.Add(unsafe.Pointer(m), unsafe.Sizeof(pgl.GLfloat(0))*14)) = 0
	transpose(m)
	multiply(m, &t[0])
}
func perspective(m *pgl.GLfloat, fovy pgl.GLfloat, aspect pgl.GLfloat, zNear pgl.GLfloat, zFar pgl.GLfloat) {
	var tmp [16]pgl.GLfloat
	identity(&tmp[0])
	var sine float64
	var cosine float64
	var cotangent float64
	var deltaZ float64
	var radians pgl.GLfloat = pgl.GLfloat(float64(fovy/2) * math.Pi / 180)
	deltaZ = float64(zFar - zNear)
	sin_cos(float64(radians), &sine, &cosine)
	if deltaZ == 0 || sine == 0 || aspect == 0 {
		return
	}
	cotangent = cosine / sine
	tmp[0] = pgl.GLfloat(cotangent / float64(aspect))
	tmp[5] = pgl.GLfloat(cotangent)
	tmp[10] = pgl.GLfloat(float64(float32(-(zFar + zNear))) / deltaZ)
	tmp[11] = pgl.GLfloat(-1)
	tmp[14] = pgl.GLfloat(float64(zNear*pgl.GLfloat(-2)*zFar) / deltaZ)
	tmp[15] = 0
	libc.MemCpy(unsafe.Pointer(m), unsafe.Pointer(&tmp[0]), int(unsafe.Sizeof([16]pgl.GLfloat{})))
}
func draw_gear(gear *gear, transform *pgl.GLfloat, x pgl.GLfloat, y pgl.GLfloat, angle pgl.GLfloat, color [4]pgl.GLfloat) {
	var (
		model_view            [16]pgl.GLfloat
		normal_matrix         [16]pgl.GLfloat
		model_view_projection [16]pgl.GLfloat
	)
	libc.MemCpy(unsafe.Pointer(&model_view[0]), unsafe.Pointer(transform), int(unsafe.Sizeof([16]pgl.GLfloat{})))
	translate(&model_view[0], x, y, 0)
	rotate(&model_view[0], pgl.GLfloat(math.Pi*2*float64(angle)/360.0), 0, 0, 1)
	libc.MemCpy(unsafe.Pointer(&model_view_projection[0]), unsafe.Pointer(&ProjectionMatrix[0]), int(unsafe.Sizeof([16]pgl.GLfloat{})))
	multiply(&model_view_projection[0], &model_view[0])
	libc.MemCpy(unsafe.Pointer(&uniforms.Mvp_mat[0]), unsafe.Pointer(&model_view_projection[0]), int(unsafe.Sizeof(pgl.Mat4{})))
	libc.MemCpy(unsafe.Pointer(&normal_matrix[0]), unsafe.Pointer(&model_view[0]), int(unsafe.Sizeof([16]pgl.GLfloat{})))
	invert(&normal_matrix[0])
	transpose(&normal_matrix[0])
	libc.MemCpy(unsafe.Pointer(&uniforms.Normal_mat[0]), unsafe.Pointer(&normal_matrix[0]), int(unsafe.Sizeof(pgl.Mat4{})))
	libc.MemCpy(unsafe.Pointer(&uniforms.Material_color), unsafe.Pointer(&color[0]), int(unsafe.Sizeof(pgl.Vec3{})))
	pgl.GlBindBuffer(pgl.GLenum(pgl.GL_ARRAY_BUFFER), gear.Vbo)
	pgl.GlVertexAttribPointer(0, 3, pgl.GLenum(pgl.GL_FLOAT), pgl.GL_FALSE, pgl.GLsizei(uint32(6*unsafe.Sizeof(pgl.GLfloat(0)))), 0)
	pgl.GlVertexAttribPointer(1, 3, pgl.GLenum(pgl.GL_FLOAT), pgl.GL_FALSE, pgl.GLsizei(uint32(6*unsafe.Sizeof(pgl.GLfloat(0)))), pgl.GLsizei(uint32(0+3*unsafe.Sizeof(pgl.GLfloat(0)))))
	pgl.GlEnableVertexAttribArray(0)
	pgl.GlEnableVertexAttribArray(1)
	var n int
	for n = 0; n < gear.Nstrips; n++ {
		pgl.GlDrawArrays(pgl.GLenum(pgl.GL_TRIANGLE_STRIP), (*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gear.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(n)))).First, pgl.GLsizei((*(*vertex_strip)(unsafe.Add(unsafe.Pointer(gear.Strips), unsafe.Sizeof(vertex_strip{})*uintptr(n)))).Count))
	}
	pgl.GlDisableVertexAttribArray(1)
	pgl.GlDisableVertexAttribArray(0)
}
func gears_draw() {
	var (
		red       [4]pgl.GLfloat = [4]pgl.GLfloat{pgl.GLfloat(0.8), pgl.GLfloat(0.1), pgl.GLfloat(0.0), pgl.GLfloat(1.0)}
		green     [4]pgl.GLfloat = [4]pgl.GLfloat{pgl.GLfloat(0.0), pgl.GLfloat(0.8), pgl.GLfloat(0.2), pgl.GLfloat(1.0)}
		blue      [4]pgl.GLfloat = [4]pgl.GLfloat{pgl.GLfloat(0.2), pgl.GLfloat(0.2), pgl.GLfloat(1.0), pgl.GLfloat(1.0)}
		transform [16]pgl.GLfloat
	)
	identity(&transform[0])
	pgl.GlClearColor(pgl.GLclampf(0.0), pgl.GLclampf(0.0), pgl.GLclampf(0.0), pgl.GLclampf(0.0))
	pgl.GlClear(pgl.GLbitfield(pgl.GL_COLOR_BUFFER_BIT | pgl.GL_DEPTH_BUFFER_BIT))
	translate(&transform[0], 0, 0, pgl.GLfloat(-20))
	rotate(&transform[0], pgl.GLfloat(math.Pi*2*float64(view_rot[0])/360.0), 1, 0, 0)
	rotate(&transform[0], pgl.GLfloat(math.Pi*2*float64(view_rot[1])/360.0), 0, 1, 0)
	rotate(&transform[0], pgl.GLfloat(math.Pi*2*float64(view_rot[2])/360.0), 0, 0, 1)
	draw_gear(gear1, &transform[0], pgl.GLfloat(-3.0), pgl.GLfloat(-2.0), angle, red)
	draw_gear(gear2, &transform[0], pgl.GLfloat(3.1), pgl.GLfloat(-2.0), pgl.GLfloat(float64(angle*pgl.GLfloat(-2))-9.0), green)
	draw_gear(gear3, &transform[0], pgl.GLfloat(-3.1), pgl.GLfloat(4.2), pgl.GLfloat(float64(angle*pgl.GLfloat(-2))-25.0), blue)
}
func gears_idle() {
	var (
		frames int     = 0
		tRot0  float64 = -1.0
		tRate0 float64 = -1.0
		dt     float64
		t      float64 = float64(sdl.GetTicks()) / 1000.0
	)
	if tRot0 < 0.0 {
		tRot0 = t
	}
	dt = t - tRot0
	tRot0 = t
	angle += pgl.GLfloat(dt * 70.0)
	if float64(angle) > 3600.0 {
		angle -= pgl.GLfloat(3600.0)
	}
	frames++
	if tRate0 < 0.0 {
		tRate0 = t
	}
	if t-tRate0 >= 5.0 {
		var (
			seconds pgl.GLfloat = pgl.GLfloat(t - tRate0)
			fps     pgl.GLfloat = pgl.GLfloat(frames) / seconds
		)
		stdio.Printf("%d frames in %3.1f seconds = %6.3f FPS\n", frames, seconds, fps)
		tRate0 = t
		frames = 0
	}
}
func vertex_shader(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *pgl.Shader_Builtins, uniforms interface{}) {
	var (
		v_attribs      *pgl.Vec4    = (*pgl.Vec4)(vertex_attribs)
		vs_out         *pgl.Vec3    = (*pgl.Vec3)(unsafe.Pointer(vs_output))
		u              *My_Uniforms = (uniforms).(*My_Uniforms)
		v4             pgl.Vec4     = pgl.Mult_mat4_vec4(u.Normal_mat, *(*pgl.Vec4)(unsafe.Add(unsafe.Pointer(v_attribs), unsafe.Sizeof(pgl.Vec4{})*1)))
		v3             pgl.Vec3     = pgl.Vec3{X: v4.X, Y: v4.Y, Z: v4.Z}
		N              pgl.Vec3     = pgl.Norm_vec3(v3)
		light_pos      pgl.Vec3     = pgl.Vec3{X: 5.0, Y: 5.0, Z: 10.0}
		L              pgl.Vec3     = pgl.Norm_vec3(light_pos)
		tmp            float32      = pgl.Dot_vec3s(N, L)
		diff_intensity float32
	)
	if float64(tmp) > 0.0 {
		diff_intensity = tmp
	} else {
		diff_intensity = 0.0
	}
	*(*pgl.Vec3)(unsafe.Add(unsafe.Pointer(vs_out), unsafe.Sizeof(pgl.Vec3{})*0)) = pgl.Scale_vec3(u.Material_color, diff_intensity)
	builtins.Gl_Position = pgl.Mult_mat4_vec4(u.Mvp_mat, *(*pgl.Vec4)(unsafe.Add(unsafe.Pointer(v_attribs), unsafe.Sizeof(pgl.Vec4{})*0)))
}
func fragment_shader(fs_input *float32, builtins *pgl.Shader_Builtins, uniforms interface{}) {
	var color pgl.Vec3 = *(*pgl.Vec3)(unsafe.Add(unsafe.Pointer((*pgl.Vec3)(unsafe.Pointer(fs_input))), unsafe.Sizeof(pgl.Vec3{})*0))
	builtins.Gl_FragColor.X = color.X
	builtins.Gl_FragColor.Y = color.Y
	builtins.Gl_FragColor.Z = color.Z
	builtins.Gl_FragColor.W = 1
}
func gears_init() {
	var program pgl.GLuint
	pgl.GlEnable(pgl.GLenum(pgl.GL_CULL_FACE))
	pgl.GlEnable(pgl.GLenum(pgl.GL_DEPTH_TEST))
	var smooth [3]pgl.GLenum = [3]pgl.GLenum{pgl.GLenum(pgl.SMOOTH), pgl.GLenum(pgl.SMOOTH), pgl.GLenum(pgl.SMOOTH)}
	program = pgl.PglCreateProgram(vertex_shader, fragment_shader, 3, &smooth[0], pgl.GL_FALSE)
	pgl.GlUseProgram(program)
	pgl.PglSetUniform(&uniforms)
	perspective(&ProjectionMatrix[0], pgl.GLfloat(60.0), pgl.GLfloat(int(WIDTH/HEIGHT)), pgl.GLfloat(1.0), pgl.GLfloat(1024.0))
	pgl.GlViewport(0, 0, WIDTH, HEIGHT)
	gear1 = create_gear(pgl.GLfloat(1.0), pgl.GLfloat(4.0), pgl.GLfloat(1.0), 20, pgl.GLfloat(0.7))
	gear2 = create_gear(pgl.GLfloat(0.5), pgl.GLfloat(2.0), pgl.GLfloat(2.0), 10, pgl.GLfloat(0.7))
	gear3 = create_gear(pgl.GLfloat(1.3), pgl.GLfloat(2.0), pgl.GLfloat(0.5), 10, pgl.GLfloat(0.7))
}
func check_errors(n int, str *byte) {
	var (
		error pgl.GLenum
		err   int = 0
	)
	for (func() pgl.GLenum {
		error = pgl.GlGetError()
		return error
	}()) != pgl.GLenum(pgl.GL_NO_ERROR) {
		switch error {
		case pgl.GL_INVALID_ENUM:
			stdio.Fprintf(stdio.Stderr(), "invalid enum\n")
		case pgl.GL_INVALID_VALUE:
			stdio.Fprintf(stdio.Stderr(), "invalid value\n")
		case pgl.GL_INVALID_OPERATION:
			stdio.Fprintf(stdio.Stderr(), "invalid operation\n")
		case pgl.GL_INVALID_FRAMEBUFFER_OPERATION:
			stdio.Fprintf(stdio.Stderr(), "invalid framebuffer operation\n")
		case pgl.GL_OUT_OF_MEMORY:
			stdio.Fprintf(stdio.Stderr(), "out of memory\n")
		default:
			stdio.Fprintf(stdio.Stderr(), "wtf?\n")
		}
		err = 1
	}
	if err != 0 {
		stdio.Fprintf(stdio.Stderr(), "%d: %s\n\n", n, func() string {
			if str == nil {
				return "Errors cleared"
			}
			return libc.GoString(str)
		}())
	}
}

func cleanup() {
	pgl.Free_glContext(&the_context)
	tex.Destroy()
	ren.Destroy()
	window.Destroy()

	sdl.Quit()
}

func setup_context() {
	//SDL_SetMainReady();
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	var err error
	window, err = sdl.CreateWindow("c_ex3", 100, 100, WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.Quit()
		panic(err)
	}

	ren, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	tex, _ = ren.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, WIDTH, HEIGHT)

	bbufpix = nil
	if pgl.Init_glContext(&the_context, &bbufpix, WIDTH, HEIGHT, 32, 0x00FF0000, 0x0000FF00, 0x000000FF, 0xFF000000) == 0 {
		panic("failed to init context")
	}

	pgl.Set_glContext(&the_context)
}

func handle_events() bool {
	var e sdl.Event
	var sc, width, height int
	for {
		e = sdl.PollEvent()
		if e == nil {
			break
		}
		if e.GetType() == sdl.QUIT {
			return true
		}
		if e.GetType() == sdl.KEYDOWN {
			sc = int(e.(*sdl.KeyboardEvent).Keysym.Scancode)

			switch sc {
			case sdl.SCANCODE_ESCAPE:
				return true
			case sdl.SCANCODE_P:
				polygon_mode = (polygon_mode + 1) % 3
				if polygon_mode == 0 {
					pgl.GlPolygonMode(pgl.GL_FRONT_AND_BACK, pgl.GL_POINT)
				} else if polygon_mode == 1 {
					pgl.GlPolygonMode(pgl.GL_FRONT_AND_BACK, pgl.GL_LINE)
				} else {
					pgl.GlPolygonMode(pgl.GL_FRONT_AND_BACK, pgl.GL_FILL)
				}
			case sdl.SCANCODE_LEFT:
				view_rot[1] += 5.0
			case sdl.SCANCODE_RIGHT:
				view_rot[1] -= 5.0
			case sdl.SCANCODE_UP:
				view_rot[0] += 5.0
			case sdl.SCANCODE_DOWN:
				view_rot[0] -= 5.0
			}
		} else if e.GetType() == sdl.WINDOWEVENT {
			e := e.(*sdl.WindowEvent)
			switch e.Event {
			case sdl.WINDOWEVENT_RESIZED:
				fmt.Printf("window size %d x %d\n", e.Data1, e.Data2)
				width = int(e.Data1)
				height = int(e.Data2)

				/* Update the projection matrix */
				perspective(&ProjectionMatrix[0], 60.0, pgl.GLfloat(float32(width)/float32(height)), 1.0, 1024.0)

				/* Set the viewport */
				pgl.GlViewport(0, 0, pgl.GLsizei(width), pgl.GLsizei(height))
			}
		}
	}
	return false
}

func main() {
	/* Initialize the window */
	setup_context()
	polygon_mode = 2

	//no default vao in core profile ...
	var vao pgl.GLuint
	pgl.GlGenVertexArrays(1, &vao)
	pgl.GlBindVertexArray(vao)

	/* Initialize the gears */
	gears_init()

	for {
		if handle_events() {
			break
		}
		gears_idle()
		gears_draw()

		tex.Update(nil, unsafe.Slice((*byte)(unsafe.Pointer(bbufpix)), int(HEIGHT*WIDTH*unsafe.Sizeof(pgl.U32(0)))), int(WIDTH*unsafe.Sizeof(pgl.U32(0))))
		//Render the scene
		ren.Copy(tex, nil, nil)
		ren.Present()
	}
}
