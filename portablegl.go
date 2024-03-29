/*
	Package pgl is a cpu implementation of OpenGL 3.3ish written entirely in Go

MIT License
Copyright (c) 2011-2022 Robert Winkler
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and
to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED
TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
Clipping code copyright (c) Fabrice Bellard from TinyGL
https://bellard.org/TinyGL/

	(C) 1997-1998 Fabrice Bellard
	 This software is provided 'as-is', without any express or implied
	 warranty.  In no event will the authors be held liable for any damages
	 arising from the use of this software.
	 Permission is granted to anyone to use this software for any purpose,
	 including commercial applications, and to alter it and redistribute it
	 freely, subject to the following restrictions:
	 1. The origin of this software must not be misrepresented; you must not
	    claim that you wrote the original software. If you use this software
	    in a product, an acknowledgment in the product and its documentation
	    *is* required.
	 2. Altered source versions must be plainly marked as such, and must not be
	    misrepresented as being the original software.
	 3. This notice may not be removed or altered from any source distribution.

If you redistribute modified sources, I would appreciate that you
include in the files history information documenting your changes.
*/
package pgl

import (
	"encoding/binary"
	"github.com/chewxy/math32"
	"github.com/gotranspile/cxgo/runtime/cmath"
	"math"
	"unsafe"
)

const RM_PI = 3.14159265358979323846
const RM_2PI = 2.0 * RM_PI
const PI_DIV_180 = 0.017453292519943296
const INV_PI_DIV_180 = 57.2957795130823229
const FALSE = 0
const TRUE = 1
const MAX_VERTICES = 500000
const MAX_VERTEX_ATTRIBS = 16
const MAX_VERTEX_OUTPUT_COMPONENTS = 64
const MAX_DRAW_BUFFERS = 8
const CLIP_EPSILON = 1e-5

type GLuint uint32
type GLint int32
type GLint64 int64
type GLuint64 uint64
type GLushort uint16
type GLshort int16
type GLubyte uint8
type GLbyte int8
type GLchar int8
type GLsizei int32
type GLenum int64
type GLbitfield int64
type GLfloat float32
type GLclampf float32
type GLdouble float64
type GLboolean uint8

const (
	NO_ERROR                      = 0
	INVALID_ENUM                  = 1
	INVALID_VALUE                 = 2
	INVALID_OPERATION             = 3
	INVALID_FRAMEBUFFER_OPERATION = 4
	OUT_OF_MEMORY                 = 5
	ARRAY_BUFFER                  = 6
	COPY_READ_BUFFER              = 7
	COPY_WRITE_BUFFER             = 8
	ELEMENT_ARRAY_BUFFER          = 9
	PIXEL_PACK_BUFFER             = 10
	PIXEL_UNPACK_BUFFER           = 11
	TEXTURE_BUFFER                = 12
	TRANSFORM_FEEDBACK_BUFFER     = 13
	UNIFORM_BUFFER                = 14
	NUM_BUFFER_TYPES              = 15
	STREAM_DRAW                   = 16
	STREAM_READ                   = 17
	STREAM_COPY                   = 18
	STATIC_DRAW                   = 19
	STATIC_READ                   = 20
	STATIC_COPY                   = 21
	DYNAMIC_DRAW                  = 22
	DYNAMIC_READ                  = 23
	DYNAMIC_COPY                  = 24
	READ_ONLY                     = 25
	WRITE_ONLY                    = 26
	READ_WRITE                    = 27
	POINT                         = 28
	LINE                          = 29
	FILL                          = 30
	POINTS                        = 31
	LINES                         = 32
	LINE_STRIP                    = 33
	LINE_LOOP                     = 34
	TRIANGLES                     = 35
	TRIANGLE_STRIP                = 36
	TRIANGLE_FAN                  = 37
	LINE_STRIP_AJACENCY           = 38
	LINES_AJACENCY                = 39
	TRIANGLES_AJACENCY            = 40
	TRIANGLE_STRIP_AJACENCY       = 41
	LESS                          = 42
	LEQUAL                        = 43
	GREATER                       = 44
	GEQUAL                        = 45
	EQUAL                         = 46
	NOTEQUAL                      = 47
	ALWAYS                        = 48
	NEVER                         = 49
	ZERO                          = 50
	ONE                           = 51
	SRC_COLOR                     = 52
	ONE_MINUS_SRC_COLOR           = 53
	DST_COLOR                     = 54
	ONE_MINUS_DST_COLOR           = 55
	SRC_ALPHA                     = 56
	ONE_MINUS_SRC_ALPHA           = 57
	DST_ALPHA                     = 58
	ONE_MINUS_DST_ALPHA           = 59
	CONSTANT_COLOR                = 60
	ONE_MINUS_CONSTANT_COLOR      = 61
	CONSTANT_ALPHA                = 62
	ONE_MINUS_CONSTANT_ALPHA      = 63
	SRC_ALPHA_SATURATE            = 64
	NUM_BLEND_FUNCS               = 65
	SRC1_COLOR                    = 66
	ONE_MINUS_SRC1_COLOR          = 67
	SRC1_ALPHA                    = 68
	ONE_MINUS_SRC1_ALPHA          = 69
	FUNC_ADD                      = 70
	FUNC_SUBTRACT                 = 71
	FUNC_REVERSE_SUBTRACT         = 72
	MIN                           = 73
	MAX                           = 74
	NUM_BLEND_EQUATIONS           = 75
	TEXTURE_UNBOUND               = 76
	TEXTURE_1D                    = 77
	TEXTURE_2D                    = 78
	TEXTURE_3D                    = 79
	TEXTURE_1D_ARRAY              = 80
	TEXTURE_2D_ARRAY              = 81
	TEXTURE_RECTANGLE             = 82
	TEXTURE_CUBE_MAP              = 83
	NUM_TEXTURE_TYPES             = 84
	TEXTURE_CUBE_MAP_POSITIVE_X   = 85
	TEXTURE_CUBE_MAP_NEGATIVE_X   = 86
	TEXTURE_CUBE_MAP_POSITIVE_Y   = 87
	TEXTURE_CUBE_MAP_NEGATIVE_Y   = 88
	TEXTURE_CUBE_MAP_POSITIVE_Z   = 89
	TEXTURE_CUBE_MAP_NEGATIVE_Z   = 90
	TEXTURE_BASE_LEVEL            = 91
	TEXTURE_BORDER_COLOR          = 92
	TEXTURE_COMPARE_FUNC          = 93
	TEXTURE_COMPARE_MODE          = 94
	TEXTURE_LOD_BIAS              = 95
	TEXTURE_MIN_FILTER            = 96
	TEXTURE_MAG_FILTER            = 97
	TEXTURE_MIN_LOD               = 98
	TEXTURE_MAX_LOD               = 99
	TEXTURE_MAX_LEVEL             = 100
	TEXTURE_SWIZZLE_R             = 101
	TEXTURE_SWIZZLE_G             = 102
	TEXTURE_SWIZZLE_B             = 103
	TEXTURE_SWIZZLE_A             = 104
	TEXTURE_SWIZZLE_RGBA          = 105
	TEXTURE_WRAP_S                = 106
	TEXTURE_WRAP_T                = 107
	TEXTURE_WRAP_R                = 108
	REPEAT                        = 109
	CLAMP_TO_EDGE                 = 110
	CLAMP_TO_BORDER               = 111
	MIRRORED_REPEAT               = 112
	NEAREST                       = 113
	LINEAR                        = 114
	NEAREST_MIPMAP_NEAREST        = 115
	NEAREST_MIPMAP_LINEAR         = 116
	LINEAR_MIPMAP_NEAREST         = 117
	LINEAR_MIPMAP_LINEAR          = 118
	RED                           = 119
	RG                            = 120
	RGB                           = 121
	BGR                           = 122
	RGBA                          = 123
	BGRA                          = 124
	COMPRESSED_RED                = 125
	COMPRESSED_RG                 = 126
	COMPRESSED_RGB                = math.MaxInt8
	COMPRESSED_RGBA               = 128
	UNPACK_ALIGNMENT              = 129
	PACK_ALIGNMENT                = 130
	TEXTURE0                      = 131
	TEXTURE1                      = 132
	TEXTURE2                      = 133
	TEXTURE3                      = 134
	TEXTURE4                      = 135
	TEXTURE5                      = 136
	TEXTURE6                      = 137
	TEXTURE7                      = 138
	CULL_FACE                     = 139
	DEPTH_TEST                    = 140
	DEPTH_CLAMP                   = 141
	LINE_SMOOTH                   = 142
	BLEND                         = 143
	COLOR_LOGIC_OP                = 144
	POLYGON_OFFSET_FILL           = 145
	SCISSOR_TEST                  = 146
	STENCIL_TEST                  = 147
	FIRST_VERTEX_CONVENTION       = 148
	LAST_VERTEX_CONVENTION        = 149
	POINT_SPRITE_COORD_ORIGIN     = 150
	UPPER_LEFT                    = 151
	LOWER_LEFT                    = 152
	FRONT                         = 153
	BACK                          = 154
	FRONT_AND_BACK                = 155
	CCW                           = 156
	CW                            = 157
	CLEAR                         = 158
	SET                           = 159
	COPY                          = 160
	COPY_INVERTED                 = 161
	NOOP                          = 162
	AND                           = 163
	NAND                          = 164
	OR                            = 165
	NOR                           = 166
	XOR                           = 167
	EQUIV                         = 168
	AND_REVERSE                   = 169
	AND_INVERTED                  = 170
	OR_REVERSE                    = 171
	OR_INVERTED                   = 172
	INVERT                        = 173
	KEEP                          = 174
	REPLACE                       = 175
	INCR                          = 176
	INCR_WRAP                     = 177
	DECR                          = 178
	DECR_WRAP                     = 179
	UNSIGNED_BYTE                 = 180
	BYTE                          = 181
	BITMAP                        = 182
	UNSIGNED_SHORT                = 183
	SHORT                         = 184
	UNSIGNED_INT                  = 185
	INT                           = 186
	FLOAT                         = 187
	VENDOR                        = 188
	RENDERER                      = 189
	VERSION                       = 190
	SHADING_LANGUAGE_VERSION      = 191
	POLYGON_OFFSET_FACTOR         = 192
	POLYGON_OFFSET_UNITS          = 193
	POINT_SIZE                    = 194
	DEPTH_CLEAR_VALUE             = 195
	DEPTH_RANGE                   = 196
	STENCIL_WRITE_MASK            = 197
	STENCIL_REF                   = 198
	STENCIL_VALUE_MASK            = 199
	STENCIL_FUNC                  = 200
	STENCIL_FAIL                  = 201
	STENCIL_PASS_DEPTH_FAIL       = 202
	STENCIL_PASS_DEPTH_PASS       = 203
	STENCIL_BACK_WRITE_MASK       = 204
	STENCIL_BACK_REF              = 205
	STENCIL_BACK_VALUE_MASK       = 206
	STENCIL_BACK_FUNC             = 207
	STENCIL_BACK_FAIL             = 208
	STENCIL_BACK_PASS_DEPTH_FAIL  = 209
	STENCIL_BACK_PASS_DEPTH_PASS  = 210
	LOGIC_OP_MODE                 = 211
	BLEND_SRC_RGB                 = 212
	BLEND_SRC_ALPHA               = 213
	BLEND_DST_RGB                 = 214
	BLEND_DST_ALPHA               = 215
	BLEND_EQUATION_RGB            = 216
	BLEND_EQUATION_ALPHA          = 217
	CULL_FACE_MODE                = 218
	FRONT_FACE                    = 219
	DEPTH_FUNC                    = 220
	PROVOKING_VERTEX              = 221
	POLYGON_MODE                  = 222
	COMPUTE_SHADER                = 223
	VERTEX_SHADER                 = 224
	TESS_CONTROL_SHADER           = 225
	TESS_EVALUATION_SHADER        = 226
	GEOMETRY_SHADER               = 227
	FRAGMENT_SHADER               = 228
	INFO_LOG_LENGTH               = 229
	COMPILE_STATUS                = 230
	LINK_STATUS                   = 231
	COLOR_BUFFER_BIT              = 1 << 10
	DEPTH_BUFFER_BIT              = 1 << 11
	STENCIL_BUFFER_BIT            = 1 << 12
)
const (
	SMOOTH = iota
	FLAT
	NOPERSPECTIVE
)

type PerVertex struct {
	Gl_Position     Vec4
	Gl_PointSize    float32
	Gl_ClipDistance [6]float32
}
type Shader_Builtins struct {
	Gl_Position    Vec4
	Gl_InstanceID  GLint
	Gl_PointCoord  vec2
	Gl_FrontFacing GLboolean
	Gl_FragCoord   Vec4
	Gl_FragColor   Vec4
	Gl_FragDepth   float32
	Discard        GLboolean
}
type vert_func func(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *Shader_Builtins, uniforms interface{})
type frag_func func(fs_input *float32, builtins *Shader_Builtins, uniforms interface{})
type glProgram struct {
	Vertex_shader        vert_func
	Fragment_shader      frag_func
	Uniform              interface{}
	Vs_output_size       int64
	Interpolation        [MAX_VERTEX_OUTPUT_COMPONENTS]GLenum
	Fragdepth_or_discard GLboolean
	Deleted              GLboolean
}
type glBuffer struct {
	Size       GLsizei
	Type       GLenum
	Data       []u8
	Deleted    bool
	User_owned bool
}
type glVertex_Attrib struct {
	Size       GLint
	Type       GLenum
	Stride     GLsizei
	Offset     GLsizei
	Normalized bool
	Buf        uint64
	Enabled    bool
	Divisor    GLuint
}
type glVertex_Array struct {
	Vertex_attribs [MAX_VERTEX_ATTRIBS]glVertex_Attrib
	Element_buffer GLuint
	Deleted        bool
}
type glTexture struct {
	W          uint64
	H          uint64
	D          uint64
	Base_level int64
	Mag_filter GLenum
	Min_filter GLenum
	Wrap_s     GLenum
	Wrap_t     GLenum
	Wrap_r     GLenum
	Format     GLenum
	Type       GLenum
	Deleted    GLboolean
	User_owned GLboolean
	Data       []u8
}
type glVertex struct {
	Clip_space   Vec4
	Screen_space Vec4
	Clip_code    int64
	Edge_flag    int64
	Vs_out       []float32
}
type glFramebuffer struct {
	Buf     []u8
	Lastrow []u8
	W       uint64
	H       uint64
}
type Vertex_Shader_output struct {
	Size          int64
	Interpolation []GLenum
	Output_buf    []float32
}
type draw_triangle_func func(v0 *glVertex, v1 *glVertex, v2 *glVertex, provoke uint64)

type GlContext struct {
	Vp_mat                 Mat4
	X_min                  int64
	Y_min                  int64
	X_max                  uint64
	Y_max                  uint64
	Vertex_arrays          []glVertex_Array
	Buffers                []glBuffer
	Textures               []glTexture
	Programs               []glProgram
	Cur_vertex_array       GLuint
	Bound_buffers          [9]GLuint
	Bound_textures         [7]GLuint
	Cur_texture2D          GLuint
	Cur_program            GLuint
	Error                  GLenum
	Vertex_attribs_vs      [MAX_VERTEX_ATTRIBS]Vec4
	Builtins               Shader_Builtins
	Vs_output              Vertex_Shader_output
	Fs_input               [MAX_VERTEX_OUTPUT_COMPONENTS]float32
	Depth_test             GLboolean
	Line_smooth            GLboolean
	Cull_face              GLboolean
	Fragdepth_or_discard   GLboolean
	Depth_clamp            GLboolean
	Depth_mask             GLboolean
	Blend                  GLboolean
	Logic_ops              GLboolean
	Poly_offset            GLboolean
	Scissor_test           GLboolean
	Stencil_test           GLboolean
	Stencil_writemask      GLuint
	Stencil_writemask_back GLuint
	Stencil_ref            GLint
	Stencil_ref_back       GLint
	Stencil_valuemask      GLuint
	Stencil_valuemask_back GLuint
	Stencil_func           GLenum
	Stencil_func_back      GLenum
	Stencil_sfail          GLenum
	Stencil_dpfail         GLenum
	Stencil_dppass         GLenum
	Stencil_sfail_back     GLenum
	Stencil_dpfail_back    GLenum
	Stencil_dppass_back    GLenum
	Logic_func             GLenum
	Blend_sfactor          GLenum
	Blend_dfactor          GLenum
	Blend_equation         GLenum
	Cull_mode              GLenum
	Front_face             GLenum
	Poly_mode_front        GLenum
	Poly_mode_back         GLenum
	Depth_func             GLenum
	Point_spr_origin       GLenum
	Provoking_vert         GLenum
	Poly_factor            GLfloat
	Poly_units             GLfloat
	Scissor_lx             GLint
	Scissor_ly             GLint
	Scissor_ux             GLsizei
	Scissor_uy             GLsizei
	Unpack_alignment       GLint
	Pack_alignment         GLint
	Clear_stencil          GLint
	Clear_color            Color
	Blend_color            Vec4
	Point_size             GLfloat
	Clear_depth            GLfloat
	Depth_range_near       GLfloat
	Depth_range_far        GLfloat
	Draw_triangle_front    draw_triangle_func
	Draw_triangle_back     draw_triangle_func
	Zbuf                   glFramebuffer
	Back_buffer            glFramebuffer
	Stencil_buf            glFramebuffer
	User_alloced_backbuf   int64
	Bitdepth               int64
	Rmask                  U32
	Gmask                  U32
	Bmask                  U32
	Amask                  U32
	Rshift                 int64
	Gshift                 int64
	Bshift                 int64
	Ashift                 int64
	Glverts                []glVertex
}

func load_rotation_mat3(mat *mat3, v Vec3, angle float32) {
	var (
		s     float32
		c     float32
		xx    float32
		yy    float32
		zz    float32
		xy    float32
		yz    float32
		zx    float32
		xs    float32
		ys    float32
		zs    float32
		one_c float32
	)
	s = float32(math.Sin(float64(angle)))
	c = float32(math.Cos(float64(angle)))
	normalize_vec3(&v)
	xx = v.X * v.X
	yy = v.Y * v.Y
	zz = v.Z * v.Z
	xy = v.X * v.Y
	yz = v.Y * v.Z
	zx = v.Z * v.X
	xs = v.X * s
	ys = v.Y * s
	zs = v.Z * s
	one_c = float32(1.0 - float64(c))
	mat[0] = (one_c * xx) + c
	mat[3] = (one_c * xy) - zs
	mat[6] = (one_c * zx) + ys
	mat[1] = (one_c * xy) + zs
	mat[4] = (one_c * yy) + c
	mat[7] = (one_c * yz) - xs
	mat[2] = (one_c * zx) - ys
	mat[5] = (one_c * yz) + xs
	mat[8] = (one_c * zz) + c
}
func Mult_mat4_mat4(c *Mat4, a, b Mat4) {
	c[0] = a[0]*b[0] + a[4]*b[1] + a[8]*b[2] + a[12]*b[3]
	c[4] = a[0]*b[4] + a[4]*b[5] + a[8]*b[6] + a[12]*b[7]
	c[8] = a[0]*b[8] + a[4]*b[9] + a[8]*b[10] + a[12]*b[11]
	c[12] = a[0]*b[12] + a[4]*b[13] + a[8]*b[14] + a[12]*b[15]
	c[1] = a[1]*b[0] + a[5]*b[1] + a[9]*b[2] + a[13]*b[3]
	c[5] = a[1]*b[4] + a[5]*b[5] + a[9]*b[6] + a[13]*b[7]
	c[9] = a[1]*b[8] + a[5]*b[9] + a[9]*b[10] + a[13]*b[11]
	c[13] = a[1]*b[12] + a[5]*b[13] + a[9]*b[14] + a[13]*b[15]
	c[2] = a[2]*b[0] + a[6]*b[1] + a[10]*b[2] + a[14]*b[3]
	c[6] = a[2]*b[4] + a[6]*b[5] + a[10]*b[6] + a[14]*b[7]
	c[10] = a[2]*b[8] + a[6]*b[9] + a[10]*b[10] + a[14]*b[11]
	c[14] = a[2]*b[12] + a[6]*b[13] + a[10]*b[14] + a[14]*b[15]
	c[3] = a[3]*b[0] + a[7]*b[1] + a[11]*b[2] + a[15]*b[3]
	c[7] = a[3]*b[4] + a[7]*b[5] + a[11]*b[6] + a[15]*b[7]
	c[11] = a[3]*b[8] + a[7]*b[9] + a[11]*b[10] + a[15]*b[11]
	c[15] = a[3]*b[12] + a[7]*b[13] + a[11]*b[14] + a[15]*b[15]
}
func Load_rotation_mat4(mat *Mat4, v Vec3, angle float32) {
	var (
		s     float32
		c     float32
		xx    float32
		yy    float32
		zz    float32
		xy    float32
		yz    float32
		zx    float32
		xs    float32
		ys    float32
		zs    float32
		one_c float32
	)
	s = float32(math.Sin(float64(angle)))
	c = float32(math.Cos(float64(angle)))
	normalize_vec3(&v)
	xx = v.X * v.X
	yy = v.Y * v.Y
	zz = v.Z * v.Z
	xy = v.X * v.Y
	yz = v.Y * v.Z
	zx = v.Z * v.X
	xs = v.X * s
	ys = v.Y * s
	zs = v.Z * s
	one_c = float32(1.0 - float64(c))
	mat[0] = (one_c * xx) + c
	mat[4] = (one_c * xy) - zs
	mat[8] = (one_c * zx) + ys
	mat[12] = 0.0
	mat[1] = (one_c * xy) + zs
	mat[5] = (one_c * yy) + c
	mat[9] = (one_c * yz) - xs
	mat[13] = 0.0
	mat[2] = (one_c * zx) - ys
	mat[6] = (one_c * yz) + xs
	mat[10] = (one_c * zz) + c
	mat[14] = 0.0
	mat[3] = 0.0
	mat[7] = 0.0
	mat[11] = 0.0
	mat[15] = 1.0
}
func make_viewport_matrix(mat *Mat4, x int64, y int64, width uint64, height uint64, opengl int64) {
	var (
		w float32
		h float32
		l float32
		t float32
		b float32
		r float32
	)
	if opengl != 0 {
		w = float32(width)
		h = float32(height)
		l = float32(x)
		b = float32(y)
		r = float32(float64(l+w) - 0.01)
		t = float32(float64(b+h) - 0.01)
		mat[0] = (r - l) / 2
		mat[4] = 0
		mat[8] = 0
		mat[12] = (l + r) / 2
		mat[1] = 0
		mat[5] = (t - b) / 2
		mat[9] = 0
		mat[13] = (b + t) / 2
		mat[2] = 0
		mat[6] = 0
		mat[10] = 1
		mat[14] = 0
		mat[3] = 0
		mat[7] = 0
		mat[11] = 0
		mat[15] = 1
	} else {
		w = float32(width)
		h = float32(height)
		l = float32(float64(x) - 0.5)
		b = float32(float64(y) - 0.5)
		r = l + w
		t = b + h
		mat[0] = (r - l) / 2
		mat[4] = 0
		mat[8] = 0
		mat[12] = (l + r) / 2
		mat[1] = 0
		mat[5] = (t - b) / 2
		mat[9] = 0
		mat[13] = (b + t) / 2
		mat[2] = 0
		mat[6] = 0
		mat[10] = 1
		mat[14] = 0
		mat[3] = 0
		mat[7] = 0
		mat[11] = 0
		mat[15] = 1
	}
}
func make_pers_matrix(mat *Mat4, z_near float32, z_far float32) {
	mat[0] = z_near
	mat[4] = 0
	mat[8] = 0
	mat[12] = 0
	mat[1] = 0
	mat[5] = z_near
	mat[9] = 0
	mat[13] = 0
	mat[2] = 0
	mat[6] = 0
	mat[10] = z_near + z_far
	mat[14] = z_far * z_near
	mat[3] = 0
	mat[7] = 0
	mat[11] = float32(-1)
	mat[15] = 0
}
func Make_perspective_matrix(mat *Mat4, fov float32, aspect float32, n float32, f float32) {
	var (
		t float32 = n * float32(math.Tan(float64(fov)*0.5))
		b float32 = -t
		l float32 = b * aspect
		r float32 = -l
	)
	make_perspective_proj_matrix(mat, l, r, b, t, n, f)
}
func make_perspective_proj_matrix(mat *Mat4, l float32, r float32, b float32, t float32, n float32, f float32) {
	mat[0] = float32((float64(n) * 2.0) / float64(r-l))
	mat[4] = 0.0
	mat[8] = (r + l) / (r - l)
	mat[12] = 0.0
	mat[1] = 0.0
	mat[5] = float32((float64(n) * 2.0) / float64(t-b))
	mat[9] = (t + b) / (t - b)
	mat[13] = 0.0
	mat[2] = 0.0
	mat[6] = 0.0
	mat[10] = -((f + n) / (f - n))
	mat[14] = float32(-((float64(f*n) * 2.0) / float64(f-n)))
	mat[3] = 0.0
	mat[7] = 0.0
	mat[11] = -1.0
	mat[15] = 0.0
}
func make_orthographic_matrix(mat *Mat4, l float32, r float32, b float32, t float32, n float32, f float32) {
	mat[0] = float32(2.0 / float64(r-l))
	mat[4] = 0
	mat[8] = 0
	mat[12] = -((r + l) / (r - l))
	mat[1] = 0
	mat[5] = float32(2.0 / float64(t-b))
	mat[9] = 0
	mat[13] = -((t + b) / (t - b))
	mat[2] = 0
	mat[6] = 0
	mat[10] = float32(2.0 / float64(f-n))
	mat[14] = -((n + f) / (f - n))
	mat[3] = 0
	mat[7] = 0
	mat[11] = 0
	mat[15] = 1
}
func lookAt(mat *Mat4, eye Vec3, center Vec3, up Vec3) {
	*mat = Mat4{}
	mat[0] = 1
	mat[5] = 1
	mat[10] = 1
	mat[15] = 1
	var f Vec3 = Norm_vec3(sub_vec3s(center, eye))
	var s Vec3 = Norm_vec3(cross_product(f, up))
	var u Vec3 = cross_product(s, f)
	setx_mat4v3(mat, s)
	sety_mat4v3(mat, u)
	setz_mat4v3(mat, negate_vec3(f))
	setc4_mat4v3(mat, make_vec3(-Dot_vec3s(s, eye), -Dot_vec3s(u, eye), Dot_vec3s(f, eye)))
}

const CVEC_float_SZ uint64 = 50

var c *GlContext

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func gl_clipcode(pt Vec4) int64 {
	var w float32
	w = float32(float64(pt.W) * (1.0 + 1e-05))
	return ((int64(boolToInt(pt.Z < -w)) | int64(boolToInt(pt.Z > w))<<1) & (int64(boolToInt(c.Depth_clamp == 0)) | int64(boolToInt(c.Depth_clamp == 0))<<1)) | int64(boolToInt(pt.X < -w))<<2 | int64(boolToInt(pt.X > w))<<3 | int64(boolToInt(pt.Y < -w))<<4 | int64(boolToInt(pt.Y > w))<<5
}
func is_front_facing(v0 *glVertex, v1 *glVertex, v2 *glVertex) int64 {
	var (
		normal  Vec3
		tmpvec3 = Vec3{X: 0, Y: 0, Z: 1}
		p0      = vec4_to_vec3h(v0.Screen_space)
		p1      = vec4_to_vec3h(v1.Screen_space)
		p2      = vec4_to_vec3h(v2.Screen_space)
	)
	normal = cross_product(sub_vec3s(p1, p0), sub_vec3s(p2, p0))
	if c.Front_face == GLenum(CW) {
		normal = negate_vec3(normal)
	}
	if Dot_vec3s(normal, tmpvec3) <= 0 {
		return 0
	}
	return 1
}
func do_vertex(v []glVertex_Attrib, enabled []int64, num_enabled uint64, i uint64, vert uint64) {
	var (
		buf     GLuint
		buf_pos []u8
		tmpvec4 Vec4
	)
	for j := int64(0); uint64(j) < num_enabled; j++ {
		buf = GLuint(v[enabled[j]].Buf)
		buf_pos = c.Buffers[buf].Data[v[enabled[j]].Offset+GLsizei(uint64(v[enabled[j]].Stride)*i):]
		tmpvec4.X = 0.0
		tmpvec4.Y = 0.0
		tmpvec4.Z = 0.0
		tmpvec4.W = 1.0
		var b = *(*[]byte)(unsafe.Pointer(&buf_pos))
		switch v[enabled[j]].Size {
		case 4:
			tmpvec4.W = math.Float32frombits(binary.LittleEndian.Uint32(b[12:]))
			fallthrough
		case 3:
			tmpvec4.Z = math.Float32frombits(binary.LittleEndian.Uint32(b[8:]))
			fallthrough
		case 2:
			tmpvec4.Y = math.Float32frombits(binary.LittleEndian.Uint32(b[4:]))
			fallthrough
		case 1:
			tmpvec4.X = math.Float32frombits(binary.LittleEndian.Uint32(b))
		}
		c.Vertex_attribs_vs[enabled[j]] = tmpvec4
	}
	var vs_out = &c.Vs_output.Output_buf[vert*uint64(c.Vs_output.Size)]
	c.Programs[c.Cur_program].Vertex_shader(vs_out, unsafe.Pointer(&c.Vertex_attribs_vs[0]), &c.Builtins, c.Programs[c.Cur_program].Uniform)
	c.Glverts[vert].Vs_out = unsafe.Slice(vs_out, c.Vs_output.Size)
	c.Glverts[vert].Clip_space = c.Builtins.Gl_Position
	c.Glverts[vert].Edge_flag = 1
	c.Glverts[vert].Clip_code = gl_clipcode(c.Builtins.Gl_Position)
}
func vertex_stage(first GLint, count GLsizei, instance_id GLsizei, base_instance GLuint, use_elements GLboolean) {
	var (
		i           uint64
		j           uint64
		vert        uint64
		num_enabled uint64
		tmpvec4     Vec4
		buf_pos     []u8
		vec4_init   = Vec4{0.0, 0.0, 0.0, 1.0}
		enabled     [MAX_VERTEX_ATTRIBS]int64
	)
	var v = c.Vertex_arrays[c.Cur_vertex_array].Vertex_attribs[:]
	var elem_buffer = c.Vertex_arrays[c.Cur_vertex_array].Element_buffer
	for i, j = 0, 0; i < MAX_VERTEX_ATTRIBS; i++ {
		c.Vertex_attribs_vs[i] = vec4_init
		if v[i].Enabled {
			if v[i].Divisor == 0 {
				enabled[j] = int64(i)
				j++
			} else if (instance_id % GLsizei(v[i].Divisor)) == 0 {
				var n = int64(instance_id/GLsizei(v[i].Divisor) + GLsizei(base_instance))
				buf_pos = c.Buffers[v[i].Buf].Data[v[i].Offset+v[i].Stride*GLsizei(n):]
				tmpvec4.X = 0.0
				tmpvec4.Y = 0.0
				tmpvec4.Z = 0.0
				tmpvec4.W = 1.0
				var b = *(*[]byte)(unsafe.Pointer(&buf_pos))
				switch v[enabled[j]].Size {
				case 4:
					tmpvec4.W = math.Float32frombits(binary.LittleEndian.Uint32(b[12:]))
					fallthrough
				case 3:
					tmpvec4.Z = math.Float32frombits(binary.LittleEndian.Uint32(b[8:]))
					fallthrough
				case 2:
					tmpvec4.Y = math.Float32frombits(binary.LittleEndian.Uint32(b[4:]))
					fallthrough
				case 1:
					tmpvec4.X = math.Float32frombits(binary.LittleEndian.Uint32(b))
				}
				c.Vertex_attribs_vs[i] = tmpvec4
			}
		}
	}
	num_enabled = j
	if GLsizei(len(c.Glverts)) < count {
		var tmp = make([]glVertex, count)
		copy(tmp, c.Glverts)
		c.Glverts = tmp
	}
	c.Builtins.Gl_InstanceID = GLint(instance_id)
	if use_elements == 0 {
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)); vert, i = vert+1, i+1 {
			do_vertex(v, enabled[:], num_enabled, i, vert)
		}
	} else {
		var (
			uint_array   []GLuint   = unsafe.Slice((*GLuint)(unsafe.Pointer(&c.Buffers[elem_buffer].Data[0])), first+GLint(count))
			ushort_array []GLushort = unsafe.Slice((*GLushort)(unsafe.Pointer(&c.Buffers[elem_buffer].Data[0])), first+GLint(count))
			ubyte_array  []GLubyte  = unsafe.Slice((*GLubyte)(unsafe.Pointer(&c.Buffers[elem_buffer].Data[0])), first+GLint(count))
		)
		if c.Buffers[elem_buffer].Type == GLenum(UNSIGNED_BYTE) {
			for vert, i = 0, uint64(0); i < uint64(first+GLint(count)); vert, i = vert+1, i+1 {
				do_vertex(v, enabled[:], num_enabled, uint64(ubyte_array[i]), vert)
			}
		} else if c.Buffers[elem_buffer].Type == GLenum(UNSIGNED_SHORT) {
			for vert, i = 0, uint64(0); i < uint64(first+GLint(count)); vert, i = vert+1, i+1 {
				do_vertex(v, enabled[:], num_enabled, uint64(ushort_array[i]), vert)
			}
		} else {
			for vert, i = 0, uint64(0); i < uint64(first+GLint(count)); vert, i = vert+1, i+1 {
				do_vertex(v, enabled[:], num_enabled, uint64(uint_array[i]), vert)
			}
		}
	}
}
func draw_point(vert *glVertex) {
	var (
		fs_input [MAX_VERTEX_OUTPUT_COMPONENTS]float32
		point    Vec3 = vec4_to_vec3h(vert.Screen_space)
	)
	point.Z = float32((float64(point.Z)-(-1.0))/(1.0-(-1.0))*float64(c.Depth_range_far-c.Depth_range_near) + float64(c.Depth_range_near))
	if c.Depth_clamp != 0 {
		point.Z = clampf_01(point.Z)
	}
	copy(fs_input[:], vert.Vs_out[:c.Vs_output.Size])
	var x float32 = float32(float64(point.X) + 0.5)
	var y float32 = float32(float64(point.Y) + 0.5)
	var p_size float32 = float32(c.Point_size)
	var origin float32
	if c.Point_spr_origin == GLenum(UPPER_LEFT) {
		origin = -1.0
	} else {
		origin = 1.0
	}
	if p_size <= 1 {
		if x < 0 || y < 0 || x >= float32(c.Back_buffer.W) || y >= float32(c.Back_buffer.H) {
			return
		}
	}
	for i := float32(y - p_size/2); i < y+p_size/2; i++ {
		if i < 0 || i >= float32(c.Back_buffer.H) {
			continue
		}
		for j := float32(x - p_size/2); j < x+p_size/2; j++ {
			if j < 0 || j >= float32(c.Back_buffer.W) {
				continue
			}
			c.Builtins.Gl_PointCoord.X = float32((float64(int64(j))+0.5-float64(point.X))/float64(p_size) + 0.5)
			c.Builtins.Gl_PointCoord.Y = float32(float64(origin)*(float64(int64(i))+0.5-float64(point.Y))/float64(p_size) + 0.5)

			c.Builtins.Gl_FragCoord.X = j
			c.Builtins.Gl_FragCoord.Y = i
			c.Builtins.Gl_FragCoord.Z = point.Z
			c.Builtins.Gl_FragCoord.W = 1 / vert.Screen_space.W

			c.Builtins.Discard = FALSE
			c.Builtins.Gl_FragDepth = point.Z
			c.Programs[c.Cur_program].Fragment_shader(&fs_input[0], &c.Builtins, c.Programs[c.Cur_program].Uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(j), int64(i))
			}
		}
	}
}
func run_pipeline(mode GLenum, first GLint, count GLsizei, instance GLsizei, base_instance GLuint, use_elements GLboolean) {
	var (
		i       uint64
		vert    uint64
		provoke int64
	)
	if count > MAX_VERTICES {
		panic("assert failed")
	}
	vertex_stage(first, count, instance, base_instance, use_elements)
	if mode == GLenum(POINTS) {
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)); i, vert = i+1, vert+1 {
			if c.Glverts[vert].Clip_code != 0 {
				continue
			}
			c.Glverts[vert].Screen_space = Mult_mat4_vec4(c.Vp_mat, c.Glverts[vert].Clip_space)
			draw_point((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*uintptr(vert))))
		}
	} else if mode == GLenum(LINES) {
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)-1); func() uint64 {
			i += 2
			return func() uint64 {
				vert += 2
				return vert
			}()
		}() {
			draw_line_clip((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*uintptr(vert))), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*uintptr(vert+1))))
		}
	} else if mode == GLenum(LINE_STRIP) {
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)-1); func() uint64 {
			i++
			return func() uint64 {
				p := &vert
				x := *p
				*p++
				return x
			}()
		}() {
			draw_line_clip((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*uintptr(vert))), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*uintptr(vert+1))))
		}
	} else if mode == GLenum(LINE_LOOP) {
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)-1); func() uint64 {
			i++
			return func() uint64 {
				p := &vert
				x := *p
				*p++
				return x
			}()
		}() {
			draw_line_clip((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*uintptr(vert))), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*uintptr(vert+1))))
		}
		draw_line_clip((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*uintptr(count-1))), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*0)))
	} else if mode == GLenum(TRIANGLES) {
		if c.Provoking_vert == GLenum(LAST_VERTEX_CONVENTION) {
			provoke = 2
		} else {
			provoke = 0
		}
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)-2); i, vert = i+3, vert+3 {
			draw_triangle(&c.Glverts[vert], &c.Glverts[vert+1], &c.Glverts[vert+2], vert+uint64(provoke))
		}
	} else if mode == GLenum(TRIANGLE_STRIP) {
		var (
			a      uint64 = 0
			b      uint64 = 1
			toggle uint64 = 0
		)
		if c.Provoking_vert == GLenum(LAST_VERTEX_CONVENTION) {
			provoke = 0
		} else {
			provoke = -2
		}
		for vert = 2; vert < uint64(count); vert++ {
			draw_triangle(&c.Glverts[a], &c.Glverts[b], &c.Glverts[vert], vert+uint64(provoke))
			if toggle == 0 {
				a = vert
			} else {
				b = vert
			}
			toggle = uint64(boolToInt(toggle == 0))
		}
	} else if mode == GLenum(TRIANGLE_FAN) {
		if c.Provoking_vert == GLenum(LAST_VERTEX_CONVENTION) {
			provoke = 0
		} else {
			provoke = -1
		}
		for vert = 2; vert < uint64(count); vert++ {
			draw_triangle((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*0)), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*uintptr(vert-1))), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts[0]), unsafe.Sizeof(glVertex{})*uintptr(vert))), vert+uint64(provoke))
		}
	}
}
func depthtest(zval float32, zbufval float32) int64 {
	if c.Depth_mask == 0 {
		return 0
	}
	switch c.Depth_func {
	case LESS:
		return int64(boolToInt(zval < zbufval))
	case LEQUAL:
		return int64(boolToInt(zval <= zbufval))
	case GREATER:
		return int64(boolToInt(zval > zbufval))
	case GEQUAL:
		return int64(boolToInt(zval >= zbufval))
	case EQUAL:
		return int64(boolToInt(zval == zbufval))
	case NOTEQUAL:
		return int64(boolToInt(zval != zbufval))
	case ALWAYS:
		return 1
	case NEVER:
		return 0
	}
	return 0
}
func setup_fs_input(t float32, v1_out *float32, v2_out *float32, wa float32, wb float32, provoke uint64) {
	var (
		vs_output *float32 = &c.Vs_output.Output_buf[0]
		inv_wa    float32  = float32(1.0 / float64(wa))
		inv_wb    float32  = float32(1.0 / float64(wb))
	)
	for i := int64(0); i < c.Vs_output.Size; i++ {
		if c.Vs_output.Interpolation[i] == GLenum(SMOOTH) {
			c.Fs_input[i] = (*(*float32)(unsafe.Add(unsafe.Pointer(v1_out), unsafe.Sizeof(float32(0))*uintptr(i)))*inv_wa + t*(*(*float32)(unsafe.Add(unsafe.Pointer(v2_out), unsafe.Sizeof(float32(0))*uintptr(i)))*inv_wb-*(*float32)(unsafe.Add(unsafe.Pointer(v1_out), unsafe.Sizeof(float32(0))*uintptr(i)))*inv_wa)) / (inv_wa + t*(inv_wb-inv_wa))
		} else if c.Vs_output.Interpolation[i] == GLenum(NOPERSPECTIVE) {
			c.Fs_input[i] = *(*float32)(unsafe.Add(unsafe.Pointer(v1_out), unsafe.Sizeof(float32(0))*uintptr(i))) + t*(*(*float32)(unsafe.Add(unsafe.Pointer(v2_out), unsafe.Sizeof(float32(0))*uintptr(i)))-*(*float32)(unsafe.Add(unsafe.Pointer(v1_out), unsafe.Sizeof(float32(0))*uintptr(i))))
		} else {
			c.Fs_input[i] = *(*float32)(unsafe.Add(unsafe.Pointer(vs_output), unsafe.Sizeof(float32(0))*uintptr(provoke*uint64(c.Vs_output.Size)+uint64(i))))
		}
	}
	c.Builtins.Discard = FALSE
}
func clip_line(denom float32, num float32, tmin *float32, tmax *float32) int64 {
	var t float32
	if denom > 0 {
		t = num / denom
		if t > *tmax {
			return 0
		}
		if t > *tmin {
			*tmin = t
		}
	} else if denom < 0 {
		t = num / denom
		if t < *tmin {
			return 0
		}
		if t < *tmax {
			*tmax = t
		}
	} else if num > 0 {
		return 0
	}
	return 1
}
func interpolate_clipped_line(v1 *glVertex, v2 *glVertex, v1_out *float32, v2_out *float32, tmin float32, tmax float32) {
	for i := int64(0); i < c.Vs_output.Size; i++ {
		*(*float32)(unsafe.Add(unsafe.Pointer(v1_out), unsafe.Sizeof(float32(0))*uintptr(i))) = v1.Vs_out[i] + (v2.Vs_out[i]-v1.Vs_out[i])*tmin
		*(*float32)(unsafe.Add(unsafe.Pointer(v2_out), unsafe.Sizeof(float32(0))*uintptr(i))) = v1.Vs_out[i] + (v2.Vs_out[i]-v1.Vs_out[i])*tmax
	}
}
func draw_line_clip(v1 *glVertex, v2 *glVertex) {
	var (
		cc1  int64
		cc2  int64
		d    Vec4
		p1   Vec4
		p2   Vec4
		t1   Vec4
		t2   Vec4
		tmin float32
		tmax float32
	)
	cc1 = v1.Clip_code
	cc2 = v2.Clip_code
	p1 = v1.Clip_space
	p2 = v2.Clip_space
	var v1_out [MAX_VERTEX_OUTPUT_COMPONENTS]float32
	var v2_out [MAX_VERTEX_OUTPUT_COMPONENTS]float32
	var provoke uint64
	if c.Provoking_vert == GLenum(LAST_VERTEX_CONVENTION) {
		provoke = uint64((int64(uintptr(unsafe.Pointer(v2)) - uintptr(unsafe.Pointer(&c.Glverts[0])))) / int64(unsafe.Sizeof(glVertex{})))
	} else {
		provoke = uint64((int64(uintptr(unsafe.Pointer(v1)) - uintptr(unsafe.Pointer(&c.Glverts[0])))) / int64(unsafe.Sizeof(glVertex{})))
	}
	if cc1&cc2 != 0 {
		return
	} else if (cc1 | cc2) == 0 {
		t1 = Mult_mat4_vec4(c.Vp_mat, p1)
		t2 = Mult_mat4_vec4(c.Vp_mat, p2)
		if c.Line_smooth == 0 {
			draw_line_shader(t1, t2, &v1.Vs_out[0], &v2.Vs_out[0], provoke)
		} else {
			draw_line_smooth_shader(t1, t2, &v1.Vs_out[0], &v2.Vs_out[0], provoke)
		}
	} else {
		d = sub_vec4s(p2, p1)
		tmin = 0
		tmax = 1
		if clip_line(d.X+d.W, -p1.X-p1.W, &tmin, &tmax) != 0 && clip_line(-d.X+d.W, p1.X-p1.W, &tmin, &tmax) != 0 && clip_line(d.Y+d.W, -p1.Y-p1.W, &tmin, &tmax) != 0 && clip_line(-d.Y+d.W, p1.Y-p1.W, &tmin, &tmax) != 0 && clip_line(d.Z+d.W, -p1.Z-p1.W, &tmin, &tmax) != 0 && clip_line(-d.Z+d.W, p1.Z-p1.W, &tmin, &tmax) != 0 {
			t1 = add_vec4s(p1, scale_vec4(d, tmin))
			t2 = add_vec4s(p1, scale_vec4(d, tmax))
			t1 = Mult_mat4_vec4(c.Vp_mat, t1)
			t2 = Mult_mat4_vec4(c.Vp_mat, t2)
			interpolate_clipped_line(v1, v2, &v1_out[0], &v2_out[0], tmin, tmax)
			if c.Line_smooth == 0 {
				draw_line_shader(t1, t2, &v1_out[0], &v2_out[0], provoke)
			} else {
				draw_line_smooth_shader(t1, t2, &v1_out[0], &v2_out[0], provoke)
			}
		}
	}
}
func draw_line_shader(v1 Vec4, v2 Vec4, v1_out *float32, v2_out *float32, provoke uint64) {
	var (
		tmp     float32
		tmp_ptr *float32
		hp1     = vec4_to_vec3h(v1)
		hp2     = vec4_to_vec3h(v2)
		w1      = v1.W
		w2      = v2.W
		x1      = hp1.X
		x2      = hp2.X
		y1      = hp1.Y
		y2      = hp2.Y
		z1      = hp1.Z
		z2      = hp2.Z
	)
	if x2 < x1 {
		tmp = x1
		x1 = x2
		x2 = tmp
		tmp = y1
		y1 = y2
		y2 = tmp
		tmp = z1
		z1 = z2
		z2 = tmp
		tmp = w1
		w1 = w2
		w2 = tmp
		tmp_ptr = v1_out
		v1_out = v2_out
		v2_out = tmp_ptr
	}
	var m float32 = (y2 - y1) / (x2 - x1)
	var line Line = make_Line(x1, y1, x2, y2)
	var t float32
	var x float32
	var y float32
	var z float32
	var w float32
	var p1 vec2 = vec2{X: x1, Y: y1}
	var p2 vec2 = vec2{X: x2, Y: y2}
	var pr vec2
	var sub_p2p1 vec2 = sub_vec2s(p2, p1)
	var line_length_squared float32 = length_vec2(sub_p2p1)
	line_length_squared *= line_length_squared
	var fragment_shader frag_func = c.Programs[c.Cur_program].Fragment_shader
	var uniform interface{} = c.Programs[c.Cur_program].Uniform
	var fragdepth_or_discard int64 = int64(c.Programs[c.Cur_program].Fragdepth_or_discard)
	_ = fragdepth_or_discard
	var i_x1 float32
	var i_y1 float32
	var i_x2 float32
	var i_y2 float32
	i_x1 = float32(math.Floor(float64(p1.X)) + 0.5)
	i_y1 = float32(math.Floor(float64(p1.Y)) + 0.5)
	i_x2 = float32(math.Floor(float64(p2.X)) + 0.5)
	i_y2 = float32(math.Floor(float64(p2.Y)) + 0.5)
	var x_min float32
	var x_max float32
	var y_min float32
	var y_max float32
	x_min = i_x1
	x_max = i_x2
	if m <= 0 {
		y_min = i_y2
		y_max = i_y1
	} else {
		y_min = i_y1
		y_max = i_y2
	}
	z1 = float32((float64(z1)-(-1.0))/(1.0-(-1.0))*float64(c.Depth_range_far-c.Depth_range_near) + float64(c.Depth_range_near))
	z2 = float32((float64(z2)-(-1.0))/(1.0-(-1.0))*float64(c.Depth_range_far-c.Depth_range_near) + float64(c.Depth_range_near))
	if m <= float32(-1) {
		for func() float32 {
			x = x_min
			return func() float32 {
				y = y_max
				return y
			}()
		}(); y >= y_min && x <= x_max; y-- {
			pr.X = x
			pr.Y = y
			t = dot_vec2s(sub_vec2s(pr, p1), sub_p2p1) / line_length_squared
			z = (1-t)*z1 + t*z2
			w = (1-t)*w1 + t*w2
			c.Builtins.Gl_FragCoord.X = x
			c.Builtins.Gl_FragCoord.Y = y
			c.Builtins.Gl_FragCoord.Z = z
			c.Builtins.Gl_FragCoord.W = 1 / w
			c.Builtins.Discard = FALSE
			c.Builtins.Gl_FragDepth = z
			setup_fs_input(t, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(x), int64(y))
			}
			//line_1:
			if line_func(&line, float32(float64(x)+0.5), y-1) < 0 {
				x++
			}
		}
	} else if m <= 0 {
		for func() float32 {
			x = x_min
			return func() float32 {
				y = y_max
				return y
			}()
		}(); x <= x_max && y >= y_min; x++ {
			pr.X = x
			pr.Y = y
			t = dot_vec2s(sub_vec2s(pr, p1), sub_p2p1) / line_length_squared
			z = (1-t)*z1 + t*z2
			w = (1-t)*w1 + t*w2

			c.Builtins.Gl_FragCoord.X = x
			c.Builtins.Gl_FragCoord.Y = y
			c.Builtins.Gl_FragCoord.Z = z
			c.Builtins.Gl_FragCoord.W = 1 / w

			c.Builtins.Discard = FALSE
			c.Builtins.Gl_FragDepth = z
			setup_fs_input(t, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(x), int64(y))
			}
			//line_2:
			if line_func(&line, x+1, float32(float64(y)-0.5)) > 0 {
				y--
			}
		}
	} else if m <= 1 {
		for func() float32 {
			x = x_min
			return func() float32 {
				y = y_min
				return y
			}()
		}(); x <= x_max && y <= y_max; x++ {
			pr.X = x
			pr.Y = y
			t = dot_vec2s(sub_vec2s(pr, p1), sub_p2p1) / line_length_squared
			z = (1-t)*z1 + t*z2
			w = (1-t)*w1 + t*w2

			c.Builtins.Gl_FragCoord.X = x
			c.Builtins.Gl_FragCoord.Y = y
			c.Builtins.Gl_FragCoord.Z = z
			c.Builtins.Gl_FragCoord.W = 1 / w

			c.Builtins.Discard = FALSE
			c.Builtins.Gl_FragDepth = z
			setup_fs_input(t, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(x), int64(y))
			}
			//line_3:
			if line_func(&line, x+1, float32(float64(y)+0.5)) < 0 {
				y++
			}
		}
	} else {
		for func() float32 {
			x = x_min
			return func() float32 {
				y = y_min
				return y
			}()
		}(); y <= y_max && x <= x_max; y++ {
			pr.X = x
			pr.Y = y
			t = dot_vec2s(sub_vec2s(pr, p1), sub_p2p1) / line_length_squared
			z = (1-t)*z1 + t*z2
			w = (1-t)*w1 + t*w2

			c.Builtins.Gl_FragCoord.X = x
			c.Builtins.Gl_FragCoord.Y = y
			c.Builtins.Gl_FragCoord.Z = z
			c.Builtins.Gl_FragCoord.W = 1 / w

			c.Builtins.Discard = FALSE
			c.Builtins.Gl_FragDepth = z
			setup_fs_input(t, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(x), int64(y))
			}
			//line_4:
			if line_func(&line, float32(float64(x)+0.5), y+1) > 0 {
				x++
			}
		}
	}
}
func draw_line_smooth_shader(v1 Vec4, v2 Vec4, v1_out *float32, v2_out *float32, provoke uint64) {
	var (
		tmp                  float32
		tmp_ptr              *float32
		fragment_shader      frag_func   = c.Programs[c.Cur_program].Fragment_shader
		uniform              interface{} = c.Programs[c.Cur_program].Uniform
		fragdepth_or_discard int64       = int64(c.Programs[c.Cur_program].Fragdepth_or_discard)
		hp1                  Vec3        = vec4_to_vec3h(v1)
		hp2                  Vec3        = vec4_to_vec3h(v2)
		x1                   float32     = hp1.X
		x2                   float32     = hp2.X
		y1                   float32     = hp1.Y
		y2                   float32     = hp2.Y
		z1                   float32     = hp1.Z
		z2                   float32     = hp2.Z
		w1                   float32     = v1.W
		w2                   float32     = v2.W
		x                    int64
		j                    int64
		steep                int64 = int64(boolToInt(math32.Abs(y2-y1) > math32.Abs(x2-x1)))
	)
	if steep != 0 {
		tmp = x1
		x1 = y1
		y1 = tmp
		tmp = x2
		x2 = y2
		y2 = tmp
	}
	if x1 > x2 {
		tmp = x1
		x1 = x2
		x2 = tmp
		tmp = y1
		y1 = y2
		y2 = tmp
		tmp = z1
		z1 = z2
		z2 = tmp
		tmp = w1
		w1 = w2
		w2 = tmp
		tmp_ptr = v1_out
		v1_out = v2_out
		v2_out = tmp_ptr
	}
	var dx float32 = x2 - x1
	var dy float32 = y2 - y1
	var gradient float32 = dy / dx
	var xend float32 = float32(float64(x1) + 0.5)
	var yend float32 = y1 + gradient*(xend-x1)
	var xgap float32 = float32(1.0 - float64(cmath.Modff(float32(float64(x1)+0.5), &tmp)))
	_ = xgap
	var xpxl1 float32 = xend
	var ypxl1 float32
	cmath.Modff(yend, &ypxl1)
	z1 = float32((float64(z1)-(-1.0))/(1.0-(-1.0))*float64(c.Depth_range_far-c.Depth_range_near) + float64(c.Depth_range_near))
	if steep != 0 {
		if c.Depth_test == 0 || fragdepth_or_discard == 0 && depthtest(z1, *(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(xpxl1))*c.Zbuf.W+uint64(int64(ypxl1)))))) != 0 {
			if c.Fragdepth_or_discard == 0 && c.Depth_test != 0 {
				*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(xpxl1))*c.Zbuf.W+uint64(int64(ypxl1))))) = z1
				*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(xpxl1))*c.Zbuf.W+uint64(int64(ypxl1+1))))) = z1
			}

			c.Builtins.Gl_FragCoord.X = ypxl1
			c.Builtins.Gl_FragCoord.Y = xpxl1
			c.Builtins.Gl_FragCoord.Z = z1
			c.Builtins.Gl_FragCoord.W = 1 / w1

			setup_fs_input(0, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(ypxl1), int64(xpxl1))
			}

			c.Builtins.Gl_FragCoord.X = ypxl1 + 1
			c.Builtins.Gl_FragCoord.Y = xpxl1
			c.Builtins.Gl_FragCoord.Z = z1
			c.Builtins.Gl_FragCoord.W = 1 / w1

			setup_fs_input(0, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(ypxl1+1), int64(xpxl1))
			}
		}
	} else {
		if c.Depth_test == 0 || fragdepth_or_discard == 0 && depthtest(z1, *(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(ypxl1))*c.Zbuf.W+uint64(int64(xpxl1)))))) != 0 {
			if c.Fragdepth_or_discard == 0 && c.Depth_test != 0 {
				*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(ypxl1))*c.Zbuf.W+uint64(int64(xpxl1))))) = z1
				*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(ypxl1+1))*c.Zbuf.W+uint64(int64(xpxl1))))) = z1
			}

			c.Builtins.Gl_FragCoord.X = xpxl1
			c.Builtins.Gl_FragCoord.Y = ypxl1
			c.Builtins.Gl_FragCoord.Z = z1
			c.Builtins.Gl_FragCoord.W = 1 / w1

			setup_fs_input(0, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(xpxl1), int64(ypxl1))
			}

			c.Builtins.Gl_FragCoord.X = xpxl1
			c.Builtins.Gl_FragCoord.Y = ypxl1 + 1
			c.Builtins.Gl_FragCoord.Z = z1
			c.Builtins.Gl_FragCoord.W = 1 / w1

			setup_fs_input(0, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(xpxl1), int64(ypxl1+1))
			}
		}
	}
	var intery float32 = yend + gradient
	xend = float32(float64(x2) + 0.5)
	yend = y2 + gradient*(xend-x2)
	xgap = cmath.Modff(float32(float64(x2)+0.5), &tmp)
	var xpxl2 float32 = xend
	var ypxl2 float32
	cmath.Modff(yend, &ypxl2)
	z2 = float32((float64(z2)-(-1.0))/(1.0-(-1.0))*float64(c.Depth_range_far-c.Depth_range_near) + float64(c.Depth_range_near))
	if steep != 0 {
		if c.Depth_test == 0 || fragdepth_or_discard == 0 && depthtest(z2, *(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(xpxl2))*c.Zbuf.W+uint64(int64(ypxl2)))))) != 0 {
			if c.Fragdepth_or_discard == 0 && c.Depth_test != 0 {
				*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(xpxl2))*c.Zbuf.W+uint64(int64(ypxl2))))) = z2
				*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(xpxl2))*c.Zbuf.W+uint64(int64(ypxl2+1))))) = z2
			}

			c.Builtins.Gl_FragCoord.X = ypxl2
			c.Builtins.Gl_FragCoord.Y = xpxl2
			c.Builtins.Gl_FragCoord.Z = z2
			c.Builtins.Gl_FragCoord.W = 1 / w2

			setup_fs_input(1, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(ypxl2), int64(xpxl2))
			}

			c.Builtins.Gl_FragCoord.X = ypxl2 + 1
			c.Builtins.Gl_FragCoord.Y = xpxl2
			c.Builtins.Gl_FragCoord.Z = z2
			c.Builtins.Gl_FragCoord.W = 1 / w2

			setup_fs_input(1, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(ypxl2+1), int64(xpxl2))
			}
		}
	} else {
		if c.Depth_test == 0 || fragdepth_or_discard == 0 && depthtest(z2, *(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(ypxl2))*c.Zbuf.W+uint64(int64(xpxl2)))))) != 0 {
			if c.Fragdepth_or_discard == 0 && c.Depth_test != 0 {
				*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(ypxl2))*c.Zbuf.W+uint64(int64(xpxl2))))) = z2
				*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(ypxl2+1))*c.Zbuf.W+uint64(int64(xpxl2))))) = z2
			}

			c.Builtins.Gl_FragCoord.X = xpxl2
			c.Builtins.Gl_FragCoord.Y = ypxl2
			c.Builtins.Gl_FragCoord.Z = z2
			c.Builtins.Gl_FragCoord.W = 1 / w2

			setup_fs_input(1, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(xpxl2), int64(ypxl2))
			}

			c.Builtins.Gl_FragCoord.X = xpxl2
			c.Builtins.Gl_FragCoord.Y = ypxl2 + 1
			c.Builtins.Gl_FragCoord.Z = z2
			c.Builtins.Gl_FragCoord.W = 1 / w2

			setup_fs_input(1, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(xpxl2), int64(ypxl2+1))
			}
		}
	}
	var range_ float32 = float32(math.Ceil(float64(x2 - x1)))
	var t float32
	var z float32
	var w float32
	for func() int64 {
		j = 1
		return func() int64 {
			x = int64(xpxl1 + 1)
			return x
		}()
	}(); float32(x) < xpxl2; func() float32 {
		x++
		j++
		return func() float32 {
			intery += gradient
			return intery
		}()
	}() {
		t = float32(j) / range_
		z = (1-t)*z1 + t*z2
		w = (1-t)*w1 + t*w2
		if steep != 0 {
			if c.Fragdepth_or_discard == 0 && c.Depth_test != 0 {
				if depthtest(z, *(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-x)*c.Zbuf.W+uint64(int64(intery)))))) == 0 {
					continue
				} else {
					*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-x)*c.Zbuf.W+uint64(int64(intery))))) = z
					*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-x)*c.Zbuf.W+uint64(int64(intery+1))))) = z
				}
			}

			c.Builtins.Gl_FragCoord.X = intery
			c.Builtins.Gl_FragCoord.Y = float32(x)
			c.Builtins.Gl_FragCoord.Z = z
			c.Builtins.Gl_FragCoord.W = 1 / w

			setup_fs_input(t, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(intery), x)
			}

			c.Builtins.Gl_FragCoord.X = intery + 1
			c.Builtins.Gl_FragCoord.Y = float32(x)
			c.Builtins.Gl_FragCoord.Z = z
			c.Builtins.Gl_FragCoord.W = 1 / w

			setup_fs_input(t, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(intery+1), x)
			}
		} else {
			if c.Fragdepth_or_discard == 0 && c.Depth_test != 0 {
				if depthtest(z, *(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(intery))*c.Zbuf.W+uint64(x))))) == 0 {
					continue
				} else {
					*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(intery))*c.Zbuf.W+uint64(x)))) = z
					*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-int64(intery+1))*c.Zbuf.W+uint64(x)))) = z
				}
			}

			c.Builtins.Gl_FragCoord.X = float32(x)
			c.Builtins.Gl_FragCoord.Y = intery
			c.Builtins.Gl_FragCoord.Z = z
			c.Builtins.Gl_FragCoord.W = 1 / w

			setup_fs_input(t, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, x, int64(intery))
			}

			c.Builtins.Gl_FragCoord.X = float32(x)
			c.Builtins.Gl_FragCoord.Y = intery + 1
			c.Builtins.Gl_FragCoord.Z = z
			c.Builtins.Gl_FragCoord.W = 1 / w

			setup_fs_input(t, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, x, int64(intery+1))
			}
		}
	}
}
func draw_triangle(v0 *glVertex, v1 *glVertex, v2 *glVertex, provoke uint64) {
	var (
		c_or  int64
		c_and int64
	)
	c_and = v0.Clip_code & v1.Clip_code & v2.Clip_code
	if c_and != 0 {
		return
	}
	c_or = v0.Clip_code | v1.Clip_code | v2.Clip_code
	if c_or == 0 {
		draw_triangle_final(v0, v1, v2, provoke)
	} else {
		draw_triangle_clip(v0, v1, v2, provoke, 0)
	}
}
func draw_triangle_final(v0 *glVertex, v1 *glVertex, v2 *glVertex, provoke uint64) {
	var front_facing int64
	v0.Screen_space = Mult_mat4_vec4(c.Vp_mat, v0.Clip_space)
	v1.Screen_space = Mult_mat4_vec4(c.Vp_mat, v1.Clip_space)
	v2.Screen_space = Mult_mat4_vec4(c.Vp_mat, v2.Clip_space)
	front_facing = is_front_facing(v0, v1, v2)
	if c.Cull_face != 0 {
		if c.Cull_mode == GLenum(FRONT_AND_BACK) {
			return
		}
		if c.Cull_mode == GLenum(BACK) && front_facing == 0 {
			return
		}
		if c.Cull_mode == GLenum(FRONT) && front_facing != 0 {
			return
		}
	}
	c.Builtins.Gl_FrontFacing = GLboolean(int8(front_facing))
	if front_facing != 0 {
		c.Draw_triangle_front(v0, v1, v2, provoke)
	} else {
		c.Draw_triangle_back(v0, v1, v2, provoke)
	}
}
func clip_xmin(c *Vec4, a *Vec4, b *Vec4) float32 {
	var (
		t   float32
		dx  float32
		dy  float32
		dz  float32
		dw  float32
		den float32
	)
	dx = b.X - a.X
	dy = b.Y - a.Y
	dz = b.Z - a.Z
	dw = b.W - a.W
	den = -(-dx) + dw
	if den == 0 {
		t = 0
	} else {
		t = (-a.X - a.W) / den
	}
	c.Y = a.Y + t*dy
	c.Z = a.Z + t*dz
	c.W = a.W + t*dw
	c.X = -c.W
	return t
}
func clip_xmax(c *Vec4, a *Vec4, b *Vec4) float32 {
	var (
		t   float32
		dx  float32
		dy  float32
		dz  float32
		dw  float32
		den float32
	)
	dx = b.X - a.X
	dy = b.Y - a.Y
	dz = b.Z - a.Z
	dw = b.W - a.W
	den = -(+dx) + dw
	if den == 0 {
		t = 0
	} else {
		t = (+a.X - a.W) / den
	}
	c.Y = a.Y + t*dy
	c.Z = a.Z + t*dz
	c.W = a.W + t*dw
	c.X = +c.W
	return t
}
func clip_ymin(c *Vec4, a *Vec4, b *Vec4) float32 {
	var (
		t   float32
		dx  float32
		dy  float32
		dz  float32
		dw  float32
		den float32
	)
	dx = b.X - a.X
	dy = b.Y - a.Y
	dz = b.Z - a.Z
	dw = b.W - a.W
	den = -(-dy) + dw
	if den == 0 {
		t = 0
	} else {
		t = (-a.Y - a.W) / den
	}
	c.X = a.X + t*dx
	c.Z = a.Z + t*dz
	c.W = a.W + t*dw
	c.Y = -c.W
	return t
}
func clip_ymax(c *Vec4, a *Vec4, b *Vec4) float32 {
	var (
		t   float32
		dx  float32
		dy  float32
		dz  float32
		dw  float32
		den float32
	)
	dx = b.X - a.X
	dy = b.Y - a.Y
	dz = b.Z - a.Z
	dw = b.W - a.W
	den = -(+dy) + dw
	if den == 0 {
		t = 0
	} else {
		t = (+a.Y - a.W) / den
	}
	c.X = a.X + t*dx
	c.Z = a.Z + t*dz
	c.W = a.W + t*dw
	c.Y = +c.W
	return t
}
func clip_zmin(c *Vec4, a *Vec4, b *Vec4) float32 {
	var (
		t   float32
		dx  float32
		dy  float32
		dz  float32
		dw  float32
		den float32
	)
	dx = b.X - a.X
	dy = b.Y - a.Y
	dz = b.Z - a.Z
	dw = b.W - a.W
	den = -(-dz) + dw
	if den == 0 {
		t = 0
	} else {
		t = (-a.Z - a.W) / den
	}
	c.X = a.X + t*dx
	c.Y = a.Y + t*dy
	c.W = a.W + t*dw
	c.Z = -c.W
	return t
}
func clip_zmax(c *Vec4, a *Vec4, b *Vec4) float32 {
	var (
		t   float32
		dx  float32
		dy  float32
		dz  float32
		dw  float32
		den float32
	)
	dx = b.X - a.X
	dy = b.Y - a.Y
	dz = b.Z - a.Z
	dw = b.W - a.W
	den = -(+dz) + dw
	if den == 0 {
		t = 0
	} else {
		t = (+a.Z - a.W) / den
	}
	c.X = a.X + t*dx
	c.Y = a.Y + t*dy
	c.W = a.W + t*dw
	c.Z = +c.W
	return t
}

var clip_proc = [6]func(*Vec4, *Vec4, *Vec4) float32{clip_zmin, clip_zmax, clip_xmin, clip_xmax, clip_ymin, clip_ymax}

func update_clip_pt(q *glVertex, v0 *glVertex, v1 *glVertex, t float32) {
	for i := int64(0); i < c.Vs_output.Size; i++ {
		//why is this correct for both SMOOTH and NOPERSPECTIVE?
		q.Vs_out[i] = v0.Vs_out[i] + (v1.Vs_out[i]-v0.Vs_out[i])*t
		//FLAT should be handled indirectly by the provoke index
		//nothing to do here unless I change that
	}
	q.Clip_code = gl_clipcode(q.Clip_space)
}
func draw_triangle_clip(v0 *glVertex, v1 *glVertex, v2 *glVertex, provoke uint64, clip_bit int64) {
	var (
		c_or          int64
		c_and         int64
		c_ex_or       int64
		cc            [3]int64
		edge_flag_tmp int64
		clip_mask     int64
		tmp1          glVertex
		tmp2          glVertex
		q             [3]*glVertex
		tt            float32
		tmp1_out      [MAX_VERTEX_OUTPUT_COMPONENTS]float32
		tmp2_out      [MAX_VERTEX_OUTPUT_COMPONENTS]float32
	)
	tmp1.Vs_out = tmp1_out[:]
	tmp2.Vs_out = tmp2_out[:]
	cc[0] = v0.Clip_code
	cc[1] = v1.Clip_code
	cc[2] = v2.Clip_code
	c_or = cc[0] | cc[1] | cc[2]
	if c_or == 0 {
		draw_triangle_final(v0, v1, v2, provoke)
	} else {
		c_and = cc[0] & cc[1] & cc[2]
		if c_and != 0 {
			return
		}
		for clip_bit < 6 && (c_or&(1<<clip_bit)) == 0 {
			clip_bit++
		}
		if clip_bit == 6 {
			println("Clipping error:")
			print_vec4(v0.Clip_space, "\n")
			print_vec4(v1.Clip_space, "\n")
			print_vec4(v2.Clip_space, "\n")
			return
		}
		clip_mask = 1 << clip_bit
		c_ex_or = (cc[0] ^ cc[1] ^ cc[2]) & clip_mask
		if c_ex_or != 0 {
			if cc[0]&clip_mask != 0 {
				q[0] = v0
				q[1] = v1
				q[2] = v2
			} else if cc[1]&clip_mask != 0 {
				q[0] = v1
				q[1] = v2
				q[2] = v0
			} else {
				q[0] = v2
				q[1] = v0
				q[2] = v1
			}
			tt = clip_proc[clip_bit](&tmp1.Clip_space, &q[0].Clip_space, &q[1].Clip_space)
			update_clip_pt(&tmp1, q[0], q[1], tt)
			tt = clip_proc[clip_bit](&tmp2.Clip_space, &q[0].Clip_space, &q[2].Clip_space)
			update_clip_pt(&tmp2, q[0], q[2], tt)
			tmp1.Edge_flag = q[0].Edge_flag
			edge_flag_tmp = q[2].Edge_flag
			q[2].Edge_flag = 0
			draw_triangle_clip(&tmp1, q[1], q[2], provoke, clip_bit+1)
			tmp2.Edge_flag = 1
			tmp1.Edge_flag = 0
			q[2].Edge_flag = edge_flag_tmp
			draw_triangle_clip(&tmp2, &tmp1, q[2], provoke, clip_bit+1)
		} else {
			if (cc[0] & clip_mask) == 0 {
				q[0] = v0
				q[1] = v1
				q[2] = v2
			} else if (cc[1] & clip_mask) == 0 {
				q[0] = v1
				q[1] = v2
				q[2] = v0
			} else {
				q[0] = v2
				q[1] = v0
				q[2] = v1
			}
			tt = clip_proc[clip_bit](&tmp1.Clip_space, &q[0].Clip_space, &q[1].Clip_space)
			update_clip_pt(&tmp1, q[0], q[1], tt)
			tt = clip_proc[clip_bit](&tmp2.Clip_space, &q[0].Clip_space, &q[2].Clip_space)
			update_clip_pt(&tmp2, q[0], q[2], tt)
			tmp1.Edge_flag = 1
			tmp2.Edge_flag = q[2].Edge_flag
			draw_triangle_clip(q[0], &tmp1, &tmp2, provoke, clip_bit+1)
		}
	}
}
func draw_triangle_point(v0 *glVertex, v1 *glVertex, v2 *glVertex, provoke uint64) {
	var (
		fs_input [MAX_VERTEX_OUTPUT_COMPONENTS]float32
		point    Vec3
		vert     [3]*glVertex = [3]*glVertex{v0, v1, v2}
	)
	for i := int64(0); i < 3; i++ {
		if vert[i].Edge_flag == 0 {
			continue
		}
		point = vec4_to_vec3h(vert[i].Screen_space)
		point.Z = float32((float64(point.Z)-(-1.0))/(1.0-(-1.0))*float64(c.Depth_range_far-c.Depth_range_near) + float64(c.Depth_range_near))
		if c.Depth_clamp != 0 {
			point.Z = clampf_01(point.Z)
		}
		for j := int64(0); j < c.Vs_output.Size; j++ {
			if c.Vs_output.Interpolation[j] != GLenum(FLAT) {
				fs_input[j] = vert[i].Vs_out[j]
			} else {
				fs_input[j] = *(*float32)(unsafe.Add(unsafe.Pointer(&c.Vs_output.Output_buf[0]), unsafe.Sizeof(float32(0))*uintptr(provoke*uint64(c.Vs_output.Size)+uint64(j))))
			}
		}
		c.Builtins.Discard = FALSE
		c.Programs[c.Cur_program].Fragment_shader(&fs_input[0], &c.Builtins, c.Programs[c.Cur_program].Uniform)
		if c.Builtins.Discard == 0 {
			draw_pixel(c.Builtins.Gl_FragColor, int64(point.X), int64(point.Y))
		}
	}
}
func draw_triangle_line(v0 *glVertex, v1 *glVertex, v2 *glVertex, provoke uint64) {
	if v0.Edge_flag != 0 {
		draw_line_shader(v0.Screen_space, v1.Screen_space, &v0.Vs_out[0], &v1.Vs_out[0], provoke)
	}
	if v1.Edge_flag != 0 {
		draw_line_shader(v1.Screen_space, v2.Screen_space, &v1.Vs_out[0], &v2.Vs_out[0], provoke)
	}
	if v2.Edge_flag != 0 {
		draw_line_shader(v2.Screen_space, v0.Screen_space, &v2.Vs_out[0], &v0.Vs_out[0], provoke)
	}
}
func draw_triangle_fill(v0 *glVertex, v1 *glVertex, v2 *glVertex, provoke uint64) {
	var (
		p0              Vec4    = v0.Screen_space
		p1              Vec4    = v1.Screen_space
		p2              Vec4    = v2.Screen_space
		hp0             Vec3    = vec4_to_vec3h(p0)
		hp1             Vec3    = vec4_to_vec3h(p1)
		hp2             Vec3    = vec4_to_vec3h(p2)
		max_depth_slope float32 = 0
		poly_offset     float32 = 0
	)
	if c.Poly_offset != 0 {
		var dzxy [6]float32
		dzxy[0] = math32.Abs((hp1.Z - hp0.Z) / (hp1.X - hp0.X))
		dzxy[1] = math32.Abs((hp1.Z - hp0.Z) / (hp1.Y - hp0.Y))
		dzxy[2] = math32.Abs((hp2.Z - hp1.Z) / (hp2.X - hp1.X))
		dzxy[3] = math32.Abs((hp2.Z - hp1.Z) / (hp2.Y - hp1.Y))
		dzxy[4] = math32.Abs((hp0.Z - hp2.Z) / (hp0.X - hp2.X))
		dzxy[5] = math32.Abs((hp0.Z - hp2.Z) / (hp0.Y - hp2.Y))
		max_depth_slope = dzxy[0]
		for i := int64(1); i < 6; i++ {
			if dzxy[i] > max_depth_slope {
				max_depth_slope = dzxy[i]
			}
		}
		poly_offset = float32(float64(max_depth_slope*float32(c.Poly_factor)) + float64(c.Poly_units)*1e-06)
	}
	var x_min float32
	if hp0.X < hp1.X {
		x_min = hp0.X
	} else {
		x_min = hp1.X
	}
	var x_max float32
	if hp0.X > hp1.X {
		x_max = hp0.X
	} else {
		x_max = hp1.X
	}
	var y_min float32
	if hp0.Y < hp1.Y {
		y_min = hp0.Y
	} else {
		y_min = hp1.Y
	}
	var y_max float32
	if hp0.Y > hp1.Y {
		y_max = hp0.Y
	} else {
		y_max = hp1.Y
	}
	if hp2.X < x_min {
		x_min = hp2.X
	} else {
		x_min = x_min
	}
	if hp2.X > x_max {
		x_max = hp2.X
	} else {
		x_max = x_max
	}
	if hp2.Y < y_min {
		y_min = hp2.Y
	} else {
		y_min = y_min
	}
	if hp2.Y > y_max {
		y_max = hp2.Y
	} else {
		y_max = y_max
	}
	var l01 = make_Line(hp0.X, hp0.Y, hp1.X, hp1.Y)
	var l12 = make_Line(hp1.X, hp1.Y, hp2.X, hp2.Y)
	var l20 = make_Line(hp2.X, hp2.Y, hp0.X, hp0.Y)
	var alpha float32
	var beta float32
	var gamma float32
	var tmp float32
	var tmp2 float32
	var z float32
	var fs_input [MAX_VERTEX_OUTPUT_COMPONENTS]float32
	var perspective [MAX_VERTEX_OUTPUT_COMPONENTS * 3]float32
	var vs_output = &c.Vs_output.Output_buf[0]
	for i := int64(0); i < c.Vs_output.Size; i++ {
		perspective[i] = v0.Vs_out[i] / p0.W
		perspective[MAX_VERTEX_OUTPUT_COMPONENTS+i] = v1.Vs_out[i] / p1.W
		perspective[MAX_VERTEX_OUTPUT_COMPONENTS*2+i] = v2.Vs_out[i] / p2.W
	}
	var inv_w0 float32 = 1 / p0.W
	var inv_w1 float32 = 1 / p1.W
	var inv_w2 float32 = 1 / p2.W
	var x float32
	var y float32 = float32(math.Floor(float64(y_min)) + 0.5)
	for ; y <= y_max; y++ {
		x = float32(math.Floor(float64(x_min)) + 0.5)
		for ; x <= x_max; x++ {
			gamma = line_func(&l01, x, y) / line_func(&l01, hp2.X, hp2.Y)
			beta = line_func(&l20, x, y) / line_func(&l20, hp1.X, hp1.Y)
			alpha = 1 - beta - gamma
			if alpha >= 0 && beta >= 0 && gamma >= 0 {
				if (alpha > 0 || line_func(&l12, hp0.X, hp0.Y)*line_func(&l12, float32(-1), -2.5) > 0) && (beta > 0 || line_func(&l20, hp1.X, hp1.Y)*line_func(&l20, float32(-1), -2.5) > 0) && (gamma > 0 || line_func(&l01, hp2.X, hp2.Y)*line_func(&l01, float32(-1), -2.5) > 0) {
					tmp2 = alpha*inv_w0 + beta*inv_w1 + gamma*inv_w2
					z = alpha*hp0.Z + beta*hp1.Z + gamma*hp2.Z
					z = float32((float64(z)-(-1.0))/(1.0-(-1.0))*float64(c.Depth_range_far-c.Depth_range_near) + float64(c.Depth_range_near))
					z += poly_offset
					for i := int64(0); i < c.Vs_output.Size; i++ {
						if c.Vs_output.Interpolation[i] == GLenum(SMOOTH) {
							tmp = alpha*perspective[i] + beta*perspective[MAX_VERTEX_OUTPUT_COMPONENTS+i] + gamma*perspective[MAX_VERTEX_OUTPUT_COMPONENTS*2+i]
							fs_input[i] = tmp / tmp2
						} else if c.Vs_output.Interpolation[i] == GLenum(NOPERSPECTIVE) {
							fs_input[i] = alpha*v0.Vs_out[i] + beta*v1.Vs_out[i] + gamma*v2.Vs_out[i]
						} else {
							fs_input[i] = *(*float32)(unsafe.Add(unsafe.Pointer(vs_output), unsafe.Sizeof(float32(0))*uintptr(provoke*uint64(c.Vs_output.Size)+uint64(i))))
						}
					}
					c.Builtins.Gl_FragCoord.X = x
					c.Builtins.Gl_FragCoord.Y = y
					c.Builtins.Gl_FragCoord.Z = z
					c.Builtins.Gl_FragCoord.W = tmp2
					c.Builtins.Discard = FALSE
					c.Builtins.Gl_FragDepth = z
					c.Programs[c.Cur_program].Fragment_shader(&fs_input[0], &c.Builtins, c.Programs[c.Cur_program].Uniform)
					if c.Builtins.Discard == 0 {
						draw_pixel(c.Builtins.Gl_FragColor, int64(x), int64(y))
					}
				}
			}
		}
	}
}
func blend_pixel(src Vec4, dst Vec4) Color {
	var (
		cnst *Vec4 = &c.Blend_color
		i    float32
	)
	if src.W < (1 - dst.W) {
		i = src.W
	} else {
		i = 1 - dst.W
	}
	var Cs Vec4
	var Cd Vec4
	switch c.Blend_sfactor {
	case ZERO:
		Cs.X = 0
		Cs.Y = 0
		Cs.Z = 0
		Cs.W = 0
	case ONE:
		Cs.X = 1
		Cs.Y = 1
		Cs.Z = 1
		Cs.W = 1
	case SRC_COLOR:
		Cs = src
	case ONE_MINUS_SRC_COLOR:
		Cs.X = 1 - src.X
		Cs.Y = 1 - src.Y
		Cs.Z = 1 - src.Z
		Cs.W = 1 - src.W
	case DST_COLOR:
		Cs = dst
	case ONE_MINUS_DST_COLOR:
		Cs.X = 1 - dst.X
		Cs.Y = 1 - dst.Y
		Cs.Z = 1 - dst.Z
		Cs.W = 1 - dst.W
	case SRC_ALPHA:
		Cs.X = src.W
		Cs.Y = src.W
		Cs.Z = src.W
		Cs.W = src.W
	case ONE_MINUS_SRC_ALPHA:
		Cs.X = 1 - src.W
		Cs.Y = 1 - src.W
		Cs.Z = 1 - src.W
		Cs.W = 1 - src.W
	case DST_ALPHA:
		Cs.X = dst.W
		Cs.Y = dst.W
		Cs.Z = dst.W
		Cs.W = dst.W
	case ONE_MINUS_DST_ALPHA:
		Cs.X = 1 - dst.W
		Cs.Y = 1 - dst.W
		Cs.Z = 1 - dst.W
		Cs.W = 1 - dst.W
	case CONSTANT_COLOR:
		Cs = *cnst
	case ONE_MINUS_CONSTANT_COLOR:
		Cs.X = 1 - cnst.X
		Cs.Y = 1 - cnst.Y
		Cs.Z = 1 - cnst.Z
		Cs.W = 1 - cnst.W
	case CONSTANT_ALPHA:
		Cs.X = cnst.W
		Cs.Y = cnst.W
		Cs.Z = cnst.W
		Cs.W = cnst.W
	case ONE_MINUS_CONSTANT_ALPHA:
		Cs.X = 1 - cnst.W
		Cs.Y = 1 - cnst.W
		Cs.Z = 1 - cnst.W
		Cs.W = 1 - cnst.W
	case SRC_ALPHA_SATURATE:
		Cs.X = i
		Cs.Y = i
		Cs.Z = i
		Cs.W = 1
	default:
		print("error unrecognized blend_sfactor!\n")
	}
	switch c.Blend_dfactor {
	case ZERO:
		Cd.X = 0
		Cd.Y = 0
		Cd.Z = 0
		Cd.W = 0
	case ONE:
		Cd.X = 1
		Cd.Y = 1
		Cd.Z = 1
		Cd.W = 1
	case SRC_COLOR:
		Cd = src
	case ONE_MINUS_SRC_COLOR:
		Cd.X = 1 - src.X
		Cd.Y = 1 - src.Y
		Cd.Z = 1 - src.Z
		Cd.W = 1 - src.W
	case DST_COLOR:
		Cd = dst
	case ONE_MINUS_DST_COLOR:
		Cd.X = 1 - dst.X
		Cd.Y = 1 - dst.Y
		Cd.Z = 1 - dst.Z
		Cd.W = 1 - dst.W
	case SRC_ALPHA:
		Cd.X = src.W
		Cd.Y = src.W
		Cd.Z = src.W
		Cd.W = src.W
	case ONE_MINUS_SRC_ALPHA:
		Cd.X = 1 - src.W
		Cd.Y = 1 - src.W
		Cd.Z = 1 - src.W
		Cd.W = 1 - src.W
	case DST_ALPHA:
		Cd.X = dst.W
		Cd.Y = dst.W
		Cd.Z = dst.W
		Cd.W = dst.W
	case ONE_MINUS_DST_ALPHA:
		Cd.X = 1 - dst.W
		Cd.Y = 1 - dst.W
		Cd.Z = 1 - dst.W
		Cd.W = 1 - dst.W
	case CONSTANT_COLOR:
		Cd = *cnst
	case ONE_MINUS_CONSTANT_COLOR:
		Cd.X = 1 - cnst.X
		Cd.Y = 1 - cnst.Y
		Cd.Z = 1 - cnst.Z
		Cd.W = 1 - cnst.W
	case CONSTANT_ALPHA:
		Cd.X = cnst.W
		Cd.Y = cnst.W
		Cd.Z = cnst.W
		Cd.W = cnst.W
	case ONE_MINUS_CONSTANT_ALPHA:
		Cd.X = 1 - cnst.W
		Cd.Y = 1 - cnst.W
		Cd.Z = 1 - cnst.W
		Cd.W = 1 - cnst.W
	case SRC_ALPHA_SATURATE:
		Cd.X = i
		Cd.Y = i
		Cd.Z = i
		Cd.W = 1
	default:
		print("error unrecognized blend_dfactor!\n")
	}
	var result Vec4
	switch c.Blend_equation {
	case FUNC_ADD:
		result = add_vec4s(mult_vec4s(Cs, src), mult_vec4s(Cd, dst))
	case FUNC_SUBTRACT:
		result = sub_vec4s(mult_vec4s(Cs, src), mult_vec4s(Cd, dst))
	case FUNC_REVERSE_SUBTRACT:
		result = sub_vec4s(mult_vec4s(Cd, dst), mult_vec4s(Cs, src))
	case MIN:
		if src.X < dst.X {
			result.X = src.X
		} else {
			result.X = dst.X
		}
		if src.Y < dst.Y {
			result.Y = src.Y
		} else {
			result.Y = dst.Y
		}
		if src.Z < dst.Z {
			result.Z = src.Z
		} else {
			result.Z = dst.Z
		}
		if src.W < dst.W {
			result.W = src.W
		} else {
			result.W = dst.W
		}
	case MAX:
		if src.X > dst.X {
			result.X = src.X
		} else {
			result.X = dst.X
		}
		if src.Y > dst.Y {
			result.Y = src.Y
		} else {
			result.Y = dst.Y
		}
		if src.Z > dst.Z {
			result.Z = src.Z
		} else {
			result.Z = dst.Z
		}
		if src.W > dst.W {
			result.W = src.W
		} else {
			result.W = dst.W
		}
	default:
		print("error unrecognized blend_equation!\n")
	}
	return vec4_to_Color(result)
}
func logic_ops_pixel(s Color, d Color) Color {
	switch c.Logic_func {
	case CLEAR:
		return make_Color(0, 0, 0, 0)
	case SET:
		return make_Color(math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8)
	case COPY:
		return s
	case COPY_INVERTED:
		return make_Color(^s.R, ^s.G, ^s.B, ^s.A)
	case NOOP:
		return d
	case INVERT:
		return make_Color(^d.R, ^d.G, ^d.B, ^d.A)
	case AND:
		return make_Color(s.R&d.R, s.G&d.G, s.B&d.B, s.A&d.A)
	case NAND:
		return make_Color(^(s.R & d.R), ^(s.G & d.G), ^(s.B & d.B), ^(s.A & d.A))
	case OR:
		return make_Color(s.R|d.R, s.G|d.G, s.B|d.B, s.A|d.A)
	case NOR:
		return make_Color(^(s.R | d.R), ^(s.G | d.G), ^(s.B | d.B), ^(s.A | d.A))
	case XOR:
		return make_Color(s.R^d.R, s.G^d.G, s.B^d.B, s.A^d.A)
	case EQUIV:
		return make_Color(^(s.R ^ d.R), ^(s.G ^ d.G), ^(s.B ^ d.B), ^(s.A ^ d.A))
	case AND_REVERSE:
		return make_Color(s.R & ^d.R, s.G & ^d.G, s.B & ^d.B, s.A & ^d.A)
	case AND_INVERTED:
		return make_Color(^s.R&d.R, ^s.G&d.G, ^s.B&d.B, ^s.A&d.A)
	case OR_REVERSE:
		return make_Color(s.R|^d.R, s.G|^d.G, s.B|^d.B, s.A|^d.A)
	case OR_INVERTED:
		return make_Color(^s.R|d.R, ^s.G|d.G, ^s.B|d.B, ^s.A|d.A)
	default:
		println("Unrecognized logic op!, defaulting to COPY")
		return s
	}
}
func stencil_test(stencil u8) int64 {
	var (
		func_ int64
		ref   int64
		mask  int64
	)
	if c.Builtins.Gl_FrontFacing != 0 {
		func_ = int64(c.Stencil_func)
		ref = int64(c.Stencil_ref)
		mask = int64(c.Stencil_valuemask)
	} else {
		func_ = int64(c.Stencil_func_back)
		ref = int64(c.Stencil_ref_back)
		mask = int64(c.Stencil_valuemask_back)
	}
	switch func_ {
	case NEVER:
		return 0
	case LESS:
		return int64(boolToInt((ref & mask) < (int64(stencil) & mask)))
	case LEQUAL:
		return int64(boolToInt((ref & mask) <= (int64(stencil) & mask)))
	case GREATER:
		return int64(boolToInt((ref & mask) > (int64(stencil) & mask)))
	case GEQUAL:
		return int64(boolToInt((ref & mask) >= (int64(stencil) & mask)))
	case EQUAL:
		return int64(boolToInt((ref & mask) == (int64(stencil) & mask)))
	case NOTEQUAL:
		return int64(boolToInt((ref & mask) != (int64(stencil) & mask)))
	case ALWAYS:
		return 1
	default:
		println(("Error: unrecognized stencil function!"))
		return 0
	}
}
func stencil_op(stencil int64, depth int64, dest *u8) {
	var (
		op   int64
		ref  int64
		mask int64
		ops  *GLenum
	)
	if c.Builtins.Gl_FrontFacing != 0 {
		ops = &c.Stencil_sfail
		ref = int64(c.Stencil_ref)
		mask = int64(c.Stencil_writemask)
	} else {
		ops = &c.Stencil_sfail_back
		ref = int64(c.Stencil_ref_back)
		mask = int64(c.Stencil_writemask_back)
	}
	if stencil == 0 {
		op = int64(*(*GLenum)(unsafe.Add(unsafe.Pointer(ops), unsafe.Sizeof(GLenum(0))*0)))
	} else if depth == 0 {
		op = int64(*(*GLenum)(unsafe.Add(unsafe.Pointer(ops), unsafe.Sizeof(GLenum(0))*1)))
	} else {
		op = int64(*(*GLenum)(unsafe.Add(unsafe.Pointer(ops), unsafe.Sizeof(GLenum(0))*2)))
	}
	var val u8 = *dest
	switch op {
	case KEEP:
		return
	case ZERO:
		val = 0
	case REPLACE:
		val = u8(int8(ref))
	case INCR:
		if val < math.MaxUint8 {
			val++
		}
	case INCR_WRAP:
		val++
	case DECR:
		if val > 0 {
			val--
		}
	case DECR_WRAP:
		val--
	case INVERT:
		val = ^val
	}
	*dest = u8(int8(int64(val) & mask))
}
func draw_pixel_vec2(cf Vec4, pos vec2) {
	draw_pixel(cf, int64(pos.X), int64(pos.Y))
}
func draw_pixel(cf Vec4, x int64, y int64) {
	if c.Scissor_test != 0 {
		if x < int64(c.Scissor_lx) || y < int64(c.Scissor_ly) || x >= int64(c.Scissor_ux) || y >= int64(c.Scissor_uy) {
			return
		}
	}
	var stencil_dest *u8 = (*u8)(unsafe.Add(unsafe.Pointer(&c.Stencil_buf.Lastrow[0]), uint64(-y)*c.Stencil_buf.W+uint64(x)))
	if c.Stencil_test != 0 {
		if stencil_test(*stencil_dest) == 0 {
			stencil_op(0, 1, stencil_dest)
			return
		}
	}
	if c.Depth_test != 0 {
		var (
			dest_depth   float32 = *(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-y)*c.Zbuf.W+uint64(x))))
			src_depth    float32 = c.Builtins.Gl_FragDepth
			depth_result int64   = depthtest(src_depth, dest_depth)
		)
		if c.Stencil_test != 0 {
			stencil_op(1, depth_result, stencil_dest)
		}
		if depth_result == 0 {
			return
		}
		*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-y)*c.Zbuf.W+uint64(x)))) = src_depth
	} else if c.Stencil_test != 0 {
		stencil_op(1, 1, stencil_dest)
	}
	var dest_color Color
	var src_color Color
	var dest *U32 = (*U32)(unsafe.Add(unsafe.Pointer((*U32)(unsafe.Pointer(&c.Back_buffer.Lastrow[0]))), unsafe.Sizeof(U32(0))*uintptr(uint64(-y)*c.Back_buffer.W+uint64(x))))
	dest_color = make_Color(u8(int8(int64(*dest&c.Rmask)>>c.Rshift)), u8(int8(int64(*dest&c.Gmask)>>c.Gshift)), u8(int8(int64(*dest&c.Bmask)>>c.Bshift)), u8(int8(int64(*dest&c.Amask)>>c.Ashift)))
	if c.Blend != 0 {
		src_color = blend_pixel(cf, Color_to_vec4(dest_color))
	} else {
		cf.X = clampf_01(cf.X)
		cf.Y = clampf_01(cf.Y)
		cf.Z = clampf_01(cf.Z)
		cf.W = clampf_01(cf.W)
		src_color = vec4_to_Color(cf)
	}
	if c.Logic_ops != 0 {
		src_color = logic_ops_pixel(src_color, dest_color)
	}
	*dest = U32(src_color.A)<<c.Ashift | U32(src_color.R)<<c.Rshift | U32(src_color.G)<<c.Gshift | U32(src_color.B)<<c.Bshift
}
func default_vs(_ *float32, vertex_attribs unsafe.Pointer, builtins *Shader_Builtins, _ interface{}) {
	builtins.Gl_Position = *(*Vec4)(unsafe.Add(unsafe.Pointer((*Vec4)(vertex_attribs)), unsafe.Sizeof(Vec4{})*0))
}
func default_fs(_ *float32, builtins *Shader_Builtins, _ interface{}) {
	builtins.Gl_FragColor = Vec4{
		X: 1, Y: 0, Z: 0, W: 1,
	}
}
func init_glVertex_Array(v *glVertex_Array) {
	v.Deleted = false
	for i := int64(0); i < MAX_VERTEX_ATTRIBS; i++ {
		init_glVertex_Attrib(&v.Vertex_attribs[i])
	}
}
func init_glVertex_Attrib(v *glVertex_Attrib) {
	v.Buf = 0
	v.Enabled = false
	v.Divisor = 0
}
func Init_glContext(context *GlContext, back *[]U32, w int64, h int64, bitdepth int64, Rmask U32, Gmask U32, Bmask U32, Amask U32) int64 {
	if bitdepth > 32 || back == nil {
		return 0
	}
	context.User_alloced_backbuf = int64(boolToInt(*back != nil))
	if *back == nil {
		var bytes_per_pixel int64 = (bitdepth + 8 - 1) / 8
		*back = make([]U32, w*h*bytes_per_pixel)
		if *back == nil {
			return 0
		}
	}
	context.Zbuf.Buf = make([]u8, w*h*int64(unsafe.Sizeof(float32(0))))
	if context.Zbuf.Buf == nil {
		if context.User_alloced_backbuf == 0 {
			*back = nil
		}
		return 0
	}
	context.Stencil_buf.Buf = make([]u8, w*h)
	if context.Stencil_buf.Buf == nil {
		if context.User_alloced_backbuf == 0 {
			*back = nil
		}
		context.Zbuf.Buf = nil
		return 0
	}
	context.X_min = 0
	context.Y_min = 0
	context.X_max = uint64(w)
	context.Y_max = uint64(h)
	context.Zbuf.W = uint64(w)
	context.Zbuf.H = uint64(h)
	context.Zbuf.Lastrow = context.Zbuf.Buf[(h-1)*w*int64(unsafe.Sizeof(float32(0))):]
	context.Stencil_buf.W = uint64(w)
	context.Stencil_buf.H = uint64(h)
	context.Stencil_buf.Lastrow = context.Stencil_buf.Buf[:(h-1)*w] //(*u8)(unsafe.Add(unsafe.Pointer(&context.Stencil_buf.Buf[0]), (h-1)*w))
	context.Back_buffer.W = uint64(w)
	context.Back_buffer.H = uint64(h)
	context.Back_buffer.Buf = unsafe.Slice((*u8)(unsafe.Pointer(&(*back)[0])), w*h*int64(unsafe.Sizeof(U32(0))))
	context.Back_buffer.Lastrow = context.Back_buffer.Buf[(h-1)*w*int64(unsafe.Sizeof(U32(0))):] //(*u8)(unsafe.Add(unsafe.Pointer(&context.Back_buffer.Buf[0]), (h-1)*w*int64(unsafe.Sizeof(U32(0)))))
	context.Bitdepth = bitdepth
	context.Rmask = Rmask
	context.Gmask = Gmask
	context.Bmask = Bmask
	context.Amask = Amask
	context.Rshift = 0
	for (Rmask & 1) == 0 {
		Rmask >>= 1
		context.Rshift++
	}
	context.Gshift = 0
	for (Gmask & 1) == 0 {
		Gmask >>= 1
		context.Gshift++
	}
	context.Bshift = 0
	for (Bmask & 1) == 0 {
		Bmask >>= 1
		context.Bshift++
	}
	context.Ashift = 0
	for (Amask & 1) == 0 {
		Amask >>= 1
		context.Ashift++
	}
	context.Vertex_arrays = make([]glVertex_Array, 0, 3)
	context.Buffers = make([]glBuffer, 0, 3)
	context.Programs = make([]glProgram, 0, 3)
	context.Textures = make([]glTexture, 0, 1)
	context.Glverts = make([]glVertex, 10)
	context.Vs_output.Output_buf = make([]float32, CVEC_float_SZ)
	context.Clear_stencil = 0
	context.Clear_color = make_Color(0, 0, 0, 0)
	context.Blend_color.X = 0
	context.Blend_color.Y = 0
	context.Blend_color.Z = 0
	context.Blend_color.W = 0
	context.Point_size = GLfloat(1.0)
	context.Clear_depth = GLfloat(1.0)
	context.Depth_range_near = GLfloat(0.0)
	context.Depth_range_far = GLfloat(1.0)
	make_viewport_matrix(&context.Vp_mat, 0, 0, uint64(w), uint64(h), 1)
	context.Provoking_vert = GLenum(LAST_VERTEX_CONVENTION)
	context.Cull_mode = GLenum(BACK)
	context.Cull_face = FALSE
	context.Front_face = GLenum(CCW)
	context.Depth_test = FALSE
	context.Fragdepth_or_discard = FALSE
	context.Depth_clamp = FALSE
	context.Depth_mask = TRUE
	context.Blend = FALSE
	context.Logic_ops = FALSE
	context.Poly_offset = FALSE
	context.Scissor_test = FALSE
	context.Stencil_test = FALSE
	context.Stencil_writemask = math.MaxUint32
	context.Stencil_writemask_back = math.MaxUint32
	context.Stencil_ref = 0
	context.Stencil_ref_back = 0
	context.Stencil_valuemask = math.MaxUint32
	context.Stencil_valuemask_back = math.MaxUint32
	context.Stencil_func = GLenum(ALWAYS)
	context.Stencil_func_back = GLenum(ALWAYS)
	context.Stencil_sfail = GLenum(KEEP)
	context.Stencil_dpfail = GLenum(KEEP)
	context.Stencil_dppass = GLenum(KEEP)
	context.Stencil_sfail_back = GLenum(KEEP)
	context.Stencil_dpfail_back = GLenum(KEEP)
	context.Stencil_dppass_back = GLenum(KEEP)
	context.Logic_func = GLenum(COPY)
	context.Blend_sfactor = GLenum(ONE)
	context.Blend_dfactor = GLenum(ZERO)
	context.Blend_equation = GLenum(FUNC_ADD)
	context.Depth_func = GLenum(LESS)
	context.Line_smooth = FALSE
	context.Poly_mode_front = GLenum(FILL)
	context.Poly_mode_back = GLenum(FILL)
	context.Point_spr_origin = GLenum(UPPER_LEFT)
	context.Poly_factor = GLfloat(0.0)
	context.Poly_units = GLfloat(0.0)
	context.Scissor_lx = 0
	context.Scissor_ly = 0
	context.Scissor_ux = GLsizei(w)
	context.Scissor_uy = GLsizei(h)
	context.Unpack_alignment = 4
	context.Pack_alignment = 4
	context.Draw_triangle_front = draw_triangle_fill
	context.Draw_triangle_back = draw_triangle_fill
	context.Error = GLenum(NO_ERROR)
	var tmp_prog glProgram = glProgram{Vertex_shader: default_vs, Fragment_shader: default_fs, Uniform: nil, Vs_output_size: FALSE}
	context.Programs = append(context.Programs, tmp_prog)
	context.Cur_program = 0
	var tmp_va glVertex_Array
	init_glVertex_Array(&tmp_va)
	context.Vertex_arrays = append(context.Vertex_arrays, tmp_va)
	context.Cur_vertex_array = 0
	var tmp_buf glBuffer
	tmp_buf.User_owned = true
	tmp_buf.Deleted = false
	var tmp_tex glTexture
	tmp_tex.User_owned = TRUE
	tmp_tex.Deleted = FALSE
	tmp_tex.Format = GLenum(RGBA)
	tmp_tex.Type = GLenum(TEXTURE_UNBOUND)
	tmp_tex.Data = nil
	tmp_tex.W = 0
	tmp_tex.H = 0
	tmp_tex.D = 0
	context.Buffers = append(context.Buffers, tmp_buf)
	context.Textures = append(context.Textures, tmp_tex)
	return 1
}
func Free_glContext(context *GlContext) {
	var i int64
	context.Zbuf.Buf = nil
	context.Stencil_buf.Buf = nil
	if context.User_alloced_backbuf == 0 {
		context.Back_buffer.Buf = nil
	}
	for i = 0; uint64(i) < uint64(len(context.Buffers)); i++ {
		if context.Buffers[i].User_owned == false {
			println("freeing buffer", i)
			context.Buffers[i].Data = nil
		}
	}
	for i = 0; uint64(i) < uint64(len(context.Textures)); i++ {
		if (context.Textures[i]).User_owned == 0 {
			println("freeing texture", i)
			(context.Textures[i]).Data = nil
		}
	}
	context.Vertex_arrays = nil
	context.Buffers = nil
	context.Programs = nil
	context.Textures = nil
	context.Glverts = nil
	context.Vs_output.Output_buf = nil
}
func Set_glContext(context *GlContext) {
	c = context
}

func GetString(name GLenum) *GLubyte {
	switch name {
	case VENDOR:
		return &[]GLubyte("Robert Winkler\x00")[0]
	case RENDERER:
		return &[]GLubyte("PortableGL\x00")[0]
	case VERSION:
		return &[]GLubyte("OpenGL 3.x-ish PortableGL 0.94\x00")[0]
	case SHADING_LANGUAGE_VERSION:
		return &[]GLubyte("Go\x00")[0]
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return nil
	}
}
func GetError() GLenum {
	var err GLenum = c.Error
	c.Error = GLenum(NO_ERROR)
	return err
}
func GenVertexArrays(n GLsizei, arrays *GLuint) {
	a := unsafe.Slice(arrays, n)
	var tmp glVertex_Array
	init_glVertex_Array(&tmp)
	tmp.Deleted = false
	n--
	for i := int64(1); uint64(i) < uint64(len(c.Vertex_arrays)) && n >= 0; i++ {
		if c.Vertex_arrays[i].Deleted != false {
			c.Vertex_arrays[i] = tmp
			a[n] = GLuint(int32(i))
			n--
		}
	}
	for ; n >= 0; n-- {
		c.Vertex_arrays = append(c.Vertex_arrays, tmp)
		a[n] = GLuint(len(c.Vertex_arrays) - 1)
	}
}
func DeleteVertexArrays(n GLsizei, arrays *GLuint) {
	a := unsafe.Slice(arrays, n)
	for i := int64(0); i < int64(n); i++ {
		if a[i] == 0 || uint64(a[i]) >= uint64(len(c.Vertex_arrays)) {
			continue
		}
		if a[i] == c.Cur_vertex_array {
			c.Vertex_arrays[0] = c.Vertex_arrays[a[i]]
			c.Cur_vertex_array = 0
		}
		c.Vertex_arrays[a[i]].Deleted = true
	}
}
func GenBuffers(n GLsizei, buffers *GLuint) {
	b := unsafe.Slice(buffers, n)
	var tmp glBuffer
	tmp.User_owned = true
	tmp.Data = nil
	tmp.Deleted = false
	n--
	for i := int64(1); uint64(i) < uint64(len(c.Buffers)) && n >= 0; i++ {
		if c.Buffers[i].Deleted {
			c.Buffers[i] = tmp
			b[n] = GLuint(int32(i))
			n--
		}
	}
	for ; n >= 0; n-- {
		c.Buffers = append(c.Buffers, tmp)
		b[n] = GLuint(len(c.Buffers) - 1)
	}
}
func DeleteBuffers(n GLsizei, buffers *GLuint) {
	b := unsafe.Slice(buffers, n)
	var type_ GLenum
	for i := int64(0); i < int64(n); i++ {
		if b[i] == 0 || uint64(b[i]) >= uint64(len(c.Buffers)) {
			continue
		}
		type_ = c.Buffers[b[i]].Type
		if b[i] == c.Bound_buffers[type_] {
			c.Bound_buffers[type_] = 0
		}
		if c.Buffers[b[i]].User_owned == false {
			c.Buffers[b[i]].Data = nil
		}
		c.Buffers[b[i]].Deleted = true
	}
}
func GenTextures(n GLsizei, textures *GLuint) {
	t := unsafe.Slice(textures, n)
	var tmp glTexture
	tmp.Mag_filter = GLenum(LINEAR)
	tmp.Min_filter = GLenum(LINEAR)
	tmp.Wrap_s = GLenum(REPEAT)
	tmp.Wrap_t = GLenum(REPEAT)
	tmp.Data = nil
	tmp.Deleted = FALSE
	tmp.User_owned = TRUE
	tmp.Format = GLenum(RGBA)
	tmp.Type = GLenum(TEXTURE_UNBOUND)
	tmp.W = 0
	tmp.H = 0
	tmp.D = 0
	n--
	for i := int64(0); uint64(i) < uint64(len(c.Textures)) && n >= 0; i++ {
		if (c.Textures[i]).Deleted != 0 {
			c.Textures[i] = tmp
			t[n] = GLuint(int32(i))
			n--
		}
	}
	for ; n >= 0; n-- {
		c.Textures = append(c.Textures, tmp)
		t[n] = GLuint(len(c.Textures) - 1)
	}
}
func DeleteTextures(n GLsizei, textures *GLuint) {
	t := unsafe.Slice(textures, n)
	var type_ GLenum
	for i := int64(0); i < int64(n); i++ {
		if t[i] == 0 || uint64(t[i]) >= uint64(len(c.Textures)) {
			continue
		}
		type_ = (c.Textures[t[i]]).Type
		if t[i] == c.Bound_textures[type_] {
			c.Bound_textures[type_] = 0
		}
		if (c.Textures[t[i]]).User_owned == 0 {
			(c.Textures[t[i]]).Data = nil
		}
		(c.Textures[t[i]]).Deleted = TRUE
	}
}
func BindVertexArray(array GLuint) {
	if uint64(array) < uint64(len(c.Vertex_arrays)) && c.Vertex_arrays[array].Deleted == false {
		c.Cur_vertex_array = array
	} else if c.Error == 0 {
		c.Error = GLenum(INVALID_OPERATION)
	}
}
func BindBuffer(target GLenum, buffer GLuint) {
	if target != GLenum(ARRAY_BUFFER) && target != GLenum(ELEMENT_ARRAY_BUFFER) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	target -= GLenum(ARRAY_BUFFER)
	if uint64(buffer) < uint64(len(c.Buffers)) && c.Buffers[buffer].Deleted == false {
		c.Bound_buffers[target] = buffer
		c.Buffers[buffer].Type = target
	} else if c.Error == 0 {
		c.Error = GLenum(INVALID_OPERATION)
	}
}
func BufferData(target GLenum, size GLsizei, data unsafe.Pointer, usage GLenum) {
	if target != GLenum(ARRAY_BUFFER) && target != GLenum(ELEMENT_ARRAY_BUFFER) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	target -= GLenum(ARRAY_BUFFER)
	if c.Bound_buffers[target] == 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_OPERATION)
		}
		return
	}
	c.Buffers[c.Bound_buffers[target]].Data = nil
	if (c.Buffers[c.Bound_buffers[target]]).Data = make([]u8, size); c.Buffers[c.Bound_buffers[target]].Data == nil {
		if c.Error == 0 {
			c.Error = GLenum(OUT_OF_MEMORY)
		}
		return
	}
	if data != nil {
		copy(c.Buffers[c.Bound_buffers[target]].Data, unsafe.Slice((*u8)(data), size))
	}
	c.Buffers[c.Bound_buffers[target]].User_owned = false
	c.Buffers[c.Bound_buffers[target]].Size = size
	if target == GLenum(ELEMENT_ARRAY_BUFFER-ARRAY_BUFFER) {
		(c.Vertex_arrays[c.Cur_vertex_array]).Element_buffer = c.Bound_buffers[target]
	}
}
func BufferSubData(target GLenum, offset GLsizei, size GLsizei, data unsafe.Pointer) {
	if target != GLenum(ARRAY_BUFFER) && target != GLenum(ELEMENT_ARRAY_BUFFER) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	target -= GLenum(ARRAY_BUFFER)
	if c.Bound_buffers[target] == 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_OPERATION)
		}
		return
	}
	if offset+size > (c.Buffers[c.Bound_buffers[target]]).Size {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	copy(c.Buffers[c.Bound_buffers[target]].Data[offset:], unsafe.Slice((*u8)(data), size))
}
func BindTexture(target GLenum, texture GLuint) {
	if target < GLenum(TEXTURE_1D) || target >= GLenum(NUM_TEXTURE_TYPES) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	target -= GLenum(TEXTURE_UNBOUND + 1)
	if uint64(texture) < uint64(len(c.Textures)) && (c.Textures[texture]).Deleted == 0 {
		if (c.Textures[texture]).Type == GLenum(TEXTURE_UNBOUND) {
			c.Bound_textures[target] = texture
			(c.Textures[texture]).Type = target
		} else if (c.Textures[texture]).Type == target {
			c.Bound_textures[target] = texture
		} else if c.Error == 0 {
			c.Error = GLenum(INVALID_OPERATION)
		}
	} else if c.Error == 0 {
		c.Error = GLenum(INVALID_VALUE)
	}
}
func TexParameteri(target GLenum, pname GLenum, param GLint) {
	switch target {
	case TEXTURE_1D, TEXTURE_2D, TEXTURE_3D, TEXTURE_2D_ARRAY, TEXTURE_RECTANGLE, TEXTURE_CUBE_MAP:
		target -= GLenum(TEXTURE_UNBOUND + 1)
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	switch pname {
	case GLenum(TEXTURE_MIN_FILTER):
		if int64(param) != NEAREST && int64(param) != LINEAR && int64(param) != NEAREST_MIPMAP_NEAREST && int64(param) != NEAREST_MIPMAP_LINEAR && int64(param) != LINEAR_MIPMAP_NEAREST && int64(param) != LINEAR_MIPMAP_LINEAR {
			if c.Error == 0 {
				c.Error = GLenum(INVALID_ENUM)
			}
			return
		}
		if int64(param) == NEAREST_MIPMAP_NEAREST || int64(param) == NEAREST_MIPMAP_LINEAR {
			param = GLint(NEAREST)
		}
		if int64(param) == LINEAR_MIPMAP_NEAREST || int64(param) == LINEAR_MIPMAP_LINEAR {
			param = GLint(LINEAR)
		}
		c.Textures[c.Bound_textures[target]].Min_filter = GLenum(param)
	case GLenum(TEXTURE_MAG_FILTER):
		if int64(param) != NEAREST && int64(param) != LINEAR {
			if c.Error == 0 {
				c.Error = GLenum(INVALID_ENUM)
			}
			return
		}
		c.Textures[c.Bound_textures[target]].Mag_filter = GLenum(param)
	case GLenum(TEXTURE_WRAP_S):
		if int64(param) != REPEAT && int64(param) != CLAMP_TO_EDGE && int64(param) != CLAMP_TO_BORDER && int64(param) != MIRRORED_REPEAT {
			if c.Error == 0 {
				c.Error = GLenum(INVALID_ENUM)
			}
			return
		}
		c.Textures[c.Bound_textures[target]].Wrap_s = GLenum(param)
	case GLenum(TEXTURE_WRAP_T):
		if int64(param) != REPEAT && int64(param) != CLAMP_TO_EDGE && int64(param) != CLAMP_TO_BORDER && int64(param) != MIRRORED_REPEAT {
			if c.Error == 0 {
				c.Error = GLenum(INVALID_ENUM)
			}
			return
		}
		c.Textures[c.Bound_textures[target]].Wrap_t = GLenum(param)
	case GLenum(TEXTURE_WRAP_R):
		if int64(param) != REPEAT && int64(param) != CLAMP_TO_EDGE && int64(param) != CLAMP_TO_BORDER && int64(param) != MIRRORED_REPEAT {
			if c.Error == 0 {
				c.Error = GLenum(INVALID_ENUM)
			}
			return
		}
		c.Textures[c.Bound_textures[target]].Wrap_r = GLenum(param)
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
}
func PixelStorei(pname GLenum, param GLint) {
	if pname != GLenum(UNPACK_ALIGNMENT) && pname != GLenum(PACK_ALIGNMENT) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	switch param {
	case 1, 2, 4, 8:
	// all good here
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	if pname == GLenum(UNPACK_ALIGNMENT) {
		c.Unpack_alignment = param
	} else if pname == GLenum(PACK_ALIGNMENT) {
		c.Pack_alignment = param
	}
}
func GenerateMipmap(target GLenum) {
	if target != GLenum(TEXTURE_1D) && target != GLenum(TEXTURE_2D) && target != GLenum(TEXTURE_3D) && target != GLenum(TEXTURE_CUBE_MAP) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
}
func TexImage1D(target GLenum, level GLint, internalFormat GLint, width GLsizei, border GLint, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(TEXTURE_1D) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if border != 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	var cur_tex int64 = int64(c.Bound_textures[target-GLenum(TEXTURE_UNBOUND)-1])
	(c.Textures[cur_tex]).W = uint64(width)
	if type_ != GLenum(UNSIGNED_BYTE) {
		return
	}
	var components int64
	switch format {
	case RED:
		components = 1
	case RG:
		components = 2
	case RGB, BGR:
		components = 3
	case RGBA, BGRA:
		components = 4
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	c.Textures[cur_tex].Data = make([]u8, int64(width)*components)
	var texdata = c.Textures[cur_tex].Data
	if data != nil {
		copy(texdata, unsafe.Slice((*u8)(data), uintptr(width)*unsafe.Sizeof(U32(0))))
	}
	c.Textures[cur_tex].User_owned = FALSE
}
func TexImage2D(target GLenum, level GLint, internalFormat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(TEXTURE_2D) && target != GLenum(TEXTURE_RECTANGLE) && target != GLenum(TEXTURE_CUBE_MAP_POSITIVE_X) && target != GLenum(TEXTURE_CUBE_MAP_NEGATIVE_X) && target != GLenum(TEXTURE_CUBE_MAP_POSITIVE_Y) && target != GLenum(TEXTURE_CUBE_MAP_NEGATIVE_Y) && target != GLenum(TEXTURE_CUBE_MAP_POSITIVE_Z) && target != GLenum(TEXTURE_CUBE_MAP_NEGATIVE_Z) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if border != 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	if type_ != GLenum(UNSIGNED_BYTE) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	var components int64
	switch format {
	case GLenum(RED):
		components = 1
	case GLenum(RG):
		components = 2
	case GLenum(RGB), GLenum(BGR):
		components = 3
	case GLenum(RGBA), GLenum(BGRA):
		components = 4
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	var cur_tex int64
	var byte_width int64 = int64(width) * components
	var padding_needed int64 = byte_width % int64(c.Unpack_alignment)
	var padded_row_len int64
	if padding_needed == 0 {
		padded_row_len = byte_width
	} else {
		padded_row_len = byte_width + int64(c.Unpack_alignment) - padding_needed
	}
	if target == GLenum(TEXTURE_2D) || target == GLenum(TEXTURE_RECTANGLE) {
		cur_tex = int64(c.Bound_textures[target-GLenum(TEXTURE_UNBOUND)-1])
		c.Textures[cur_tex].W = uint64(width)
		c.Textures[cur_tex].H = uint64(height)
		c.Textures[cur_tex].Data = make([]u8, int64(height)*byte_width)
		if data != nil {
			if padding_needed == 0 {
				copy(c.Textures[cur_tex].Data, unsafe.Slice((*u8)(data), int64(height)*byte_width))
			} else {
				for i := int64(0); i < int64(height); i++ {
					copy(c.Textures[cur_tex].Data[i*byte_width:], unsafe.Slice((*u8)(unsafe.Add(unsafe.Pointer((*u8)(data)), i*padded_row_len)), byte_width))
				}
			}
		}
		(c.Textures[cur_tex]).User_owned = FALSE
	} else {
		cur_tex = int64(c.Bound_textures[TEXTURE_CUBE_MAP-TEXTURE_UNBOUND-1])
		if (c.Textures[cur_tex]).W == 0 {
			(c.Textures[cur_tex]).Data = nil
		}
		if width != height {
			if c.Error == 0 {
				c.Error = GLenum(INVALID_VALUE)
			}
			return
		}
		var mem_size int64 = int64(width*height*6) * components
		if (c.Textures[cur_tex]).W == 0 {
			(c.Textures[cur_tex]).W = uint64(width)
			(c.Textures[cur_tex]).H = uint64(width)
			if (func() *u8 {
				p := &(c.Textures[cur_tex]).Data
				c.Textures[cur_tex].Data = make([]u8, mem_size)
				return &(*p)[0]
			}()) == nil {
				if c.Error == 0 {
					c.Error = GLenum(OUT_OF_MEMORY)
				}
				return
			}
		} else if (c.Textures[cur_tex]).W != uint64(width) {
			if c.Error == 0 {
				c.Error = GLenum(INVALID_VALUE)
			}
			return
		}
		target -= GLenum(TEXTURE_CUBE_MAP_POSITIVE_X)
		var p int64 = int64(height) * byte_width
		var texdata = (c.Textures[cur_tex]).Data
		if data != nil {
			if padding_needed == 0 {
				copy(texdata[target*GLenum(p):], unsafe.Slice((*u8)(data), int64(height)*byte_width))
			} else {
				for i := int64(0); i < int64(height); i++ {
					copy(texdata[target*GLenum(p)+GLenum(i*byte_width):], unsafe.Slice((*u8)(unsafe.Add(unsafe.Pointer((*u8)(data)), i*padded_row_len)), byte_width))
				}
			}
		}
		c.Textures[cur_tex].User_owned = FALSE
	}
}
func TexImage3D(target GLenum, level GLint, internalFormat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(TEXTURE_3D) && target != GLenum(TEXTURE_2D_ARRAY) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if border != 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	var cur_tex int64 = int64(c.Bound_textures[target-GLenum(TEXTURE_UNBOUND)-1])
	c.Textures[cur_tex].W = uint64(width)
	c.Textures[cur_tex].H = uint64(height)
	c.Textures[cur_tex].D = uint64(depth)
	if type_ != GLenum(UNSIGNED_BYTE) {
		return
	}
	var components int64
	switch format {
	case GLenum(RED):
		components = 1
	case GLenum(RG):
		components = 2
	case GLenum(RGB), GLenum(BGR):
		components = 3
	case GLenum(RGBA), GLenum(BGRA):
		components = 4
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	var byte_width int64 = int64(width) * components
	var padding_needed int64 = byte_width % int64(c.Unpack_alignment)
	var padded_row_len int64
	if padding_needed == 0 {
		padded_row_len = byte_width
	} else {
		padded_row_len = byte_width + int64(c.Unpack_alignment) - padding_needed
	}
	c.Textures[cur_tex].Data = make([]u8, int64(width*height*depth)*components)
	var texdata = c.Textures[cur_tex].Data
	if data != nil {
		if padding_needed == 0 {
			copy(texdata, unsafe.Slice((*u8)(data), int(uintptr(width*height*depth)*unsafe.Sizeof(U32(0)))))
		} else {
			for i := int64(0); i < int64(height*depth); i++ {
				copy(texdata[unsafe.Sizeof(U32(0))*uintptr(i*byte_width):], unsafe.Slice((*u8)(unsafe.Add(unsafe.Pointer((*u8)(data)), i*padded_row_len)), byte_width))
			}
		}
	}
	c.Textures[cur_tex].User_owned = FALSE
}
func TexSubImage1D(target GLenum, level GLint, xoffset GLint, width GLsizei, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(TEXTURE_1D) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	var cur_tex int64 = int64(c.Bound_textures[target-GLenum(TEXTURE_UNBOUND)-1])
	if type_ != GLenum(UNSIGNED_BYTE) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if format != GLenum(RGBA) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if xoffset < 0 || uint64(xoffset+GLint(width)) > (c.Textures[cur_tex]).W {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	var texdata = (c.Textures[cur_tex]).Data
	copy(texdata[unsafe.Sizeof(U32(0))*uintptr(xoffset):], unsafe.Slice((*u8)(data), int(uintptr(width)*unsafe.Sizeof(U32(0)))))
}
func TexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(TEXTURE_2D) && target != GLenum(TEXTURE_CUBE_MAP_POSITIVE_X) && target != GLenum(TEXTURE_CUBE_MAP_NEGATIVE_X) && target != GLenum(TEXTURE_CUBE_MAP_POSITIVE_Y) && target != GLenum(TEXTURE_CUBE_MAP_NEGATIVE_Y) && target != GLenum(TEXTURE_CUBE_MAP_POSITIVE_Z) && target != GLenum(TEXTURE_CUBE_MAP_NEGATIVE_Z) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if type_ != GLenum(UNSIGNED_BYTE) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if format != GLenum(RGBA) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	var cur_tex int64
	var d *U32 = (*U32)(data)
	if target == GLenum(TEXTURE_2D) {
		cur_tex = int64(c.Bound_textures[target-GLenum(TEXTURE_UNBOUND)-1])
		var texdata = (c.Textures[cur_tex]).Data
		if xoffset < 0 || uint64(xoffset+GLint(width)) > (c.Textures[cur_tex]).W || yoffset < 0 || uint64(yoffset+GLint(height)) > (c.Textures[cur_tex]).H {
			if c.Error == 0 {
				c.Error = GLenum(INVALID_VALUE)
			}
			return
		}
		var w int64 = int64((c.Textures[cur_tex]).W)
		for i := int64(0); i < int64(height); i++ {
			copy(texdata[unsafe.Sizeof(U32(0))*uintptr((int64(yoffset)+i)*w+int64(xoffset)):], unsafe.Slice((*u8)(unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(d), unsafe.Sizeof(U32(0))*uintptr(i*int64(width)))))), int(uintptr(width)*unsafe.Sizeof(U32(0)))))
		}
	} else {
		cur_tex = int64(c.Bound_textures[TEXTURE_CUBE_MAP-TEXTURE_UNBOUND-1])
		var texdata = c.Textures[cur_tex].Data
		var w int64 = int64((c.Textures[cur_tex]).W)
		target -= GLenum(TEXTURE_CUBE_MAP_POSITIVE_X)
		var p int64 = w * w
		for i := int64(0); i < int64(height); i++ {
			copy(texdata[unsafe.Sizeof(U32(0))*uintptr(p*int64(target)+(int64(yoffset)+i)*w+int64(xoffset)):], unsafe.Slice((*u8)(unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(d), unsafe.Sizeof(U32(0))*uintptr(i*int64(width)))))), int(uintptr(width)*unsafe.Sizeof(U32(0)))))
		}
	}
}
func TexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(TEXTURE_3D) && target != GLenum(TEXTURE_2D_ARRAY) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	var cur_tex int64 = int64(c.Bound_textures[target-GLenum(TEXTURE_UNBOUND)-1])
	if type_ != GLenum(UNSIGNED_BYTE) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if format != GLenum(RGBA) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if xoffset < 0 || uint64(xoffset+GLint(width)) > (c.Textures[cur_tex]).W || yoffset < 0 || uint64(yoffset+GLint(height)) > (c.Textures[cur_tex]).H || zoffset < 0 || uint64(zoffset+GLint(depth)) > (c.Textures[cur_tex]).D {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	var w int64 = int64((c.Textures[cur_tex]).W)
	var h int64 = int64((c.Textures[cur_tex]).H)
	var p int64 = w * h
	var d *U32 = (*U32)(data)
	var texdata = c.Textures[cur_tex].Data
	for j := int64(0); j < int64(depth); j++ {
		for i := int64(0); i < int64(height); i++ {
			copy(texdata[(int64(zoffset)+j)*p+(int64(yoffset)+i)*w+int64(xoffset):], unsafe.Slice((*u8)(unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(d), unsafe.Sizeof(U32(0))*uintptr(j*int64(width)*int64(height)+i*int64(width)))))), int(uintptr(width)*unsafe.Sizeof(U32(0)))))
		}
	}
}
func VertexAttribPointer(index GLuint, size GLint, type_ GLenum, normalized GLboolean, stride GLsizei, offset GLsizei) {
	if index >= MAX_VERTEX_ATTRIBS || size < 1 || size > 4 || c.Bound_buffers[ARRAY_BUFFER-ARRAY_BUFFER] == 0 && offset != 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_OPERATION)
		}
		return
	}
	if type_ != GLenum(FLOAT) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	var v = &((c.Vertex_arrays[c.Cur_vertex_array]).Vertex_attribs[index])
	v.Size = size
	v.Type = type_
	if stride != 0 {
		v.Stride = stride
	} else {
		v.Stride = GLsizei(uint32(uintptr(size) * unsafe.Sizeof(GLfloat(0))))
	}
	v.Offset = offset
	v.Normalized = normalized != 0
	v.Buf = uint64(c.Bound_buffers[ARRAY_BUFFER-ARRAY_BUFFER])
}
func EnableVertexAttribArray(index GLuint) {
	c.Vertex_arrays[c.Cur_vertex_array].Vertex_attribs[index].Enabled = true
}
func DisableVertexAttribArray(index GLuint) {
	c.Vertex_arrays[c.Cur_vertex_array].Vertex_attribs[index].Enabled = false
}
func VertexAttribDivisor(index GLuint, divisor GLuint) {
	if index >= MAX_VERTEX_ATTRIBS {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	c.Vertex_arrays[c.Cur_vertex_array].Vertex_attribs[index].Divisor = divisor
}
func get_vertex_attrib_array(v *glVertex_Attrib, i GLsizei) Vec4 {
	var (
		buf_pos *u8 = (*u8)(unsafe.Add(unsafe.Pointer((*u8)(unsafe.Add(unsafe.Pointer(&(c.Buffers[v.Buf]).Data[0]), v.Offset))), v.Stride*i))
		tmpvec4 Vec4
	)
	var b = *(*[]byte)(unsafe.Pointer(&buf_pos))
	switch v.Size {
	case 4:
		tmpvec4.W = math.Float32frombits(binary.LittleEndian.Uint32(b[12:]))
		fallthrough
	case 3:
		tmpvec4.Z = math.Float32frombits(binary.LittleEndian.Uint32(b[8:]))
		fallthrough
	case 2:
		tmpvec4.Y = math.Float32frombits(binary.LittleEndian.Uint32(b[4:]))
		fallthrough
	case 1:
		tmpvec4.X = math.Float32frombits(binary.LittleEndian.Uint32(b))
	}
	return tmpvec4
}
func DrawArrays(mode GLenum, first GLint, count GLsizei) {
	if mode < GLenum(POINTS) || mode > GLenum(TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if count < 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	if count == 0 {
		return
	}
	run_pipeline(mode, first, count, 0, 0, FALSE)
}
func DrawElements(mode GLenum, count GLsizei, type_ GLenum, offset GLsizei) {
	if mode < GLenum(POINTS) || mode > GLenum(TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if type_ != GLenum(UNSIGNED_BYTE) && type_ != GLenum(UNSIGNED_SHORT) && type_ != GLenum(UNSIGNED_INT) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if count < 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	if count == 0 {
		return
	}
	(c.Buffers[(c.Vertex_arrays[c.Cur_vertex_array]).Element_buffer]).Type = type_
	run_pipeline(mode, GLint(offset), count, 0, 0, TRUE)
}
func DrawArraysInstanced(mode GLenum, first GLint, count GLsizei, instancecount GLsizei) {
	if mode < GLenum(POINTS) || mode > GLenum(TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if count < 0 || instancecount < 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	if count == 0 || instancecount == 0 {
		return
	}
	for instance := uint64(0); instance < uint64(instancecount); instance++ {
		run_pipeline(mode, first, count, GLsizei(uint32(instance)), 0, FALSE)
	}
}
func DrawArraysInstancedBaseInstance(mode GLenum, first GLint, count GLsizei, instancecount GLsizei, baseinstance GLuint) {
	if mode < GLenum(POINTS) || mode > GLenum(TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if count < 0 || instancecount < 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	if count == 0 || instancecount == 0 {
		return
	}
	for instance := uint64(0); instance < uint64(instancecount); instance++ {
		run_pipeline(mode, first, count, GLsizei(uint32(instance)), baseinstance, FALSE)
	}
}
func DrawElementsInstanced(mode GLenum, count GLsizei, type_ GLenum, offset GLsizei, instancecount GLsizei) {
	if mode < GLenum(POINTS) || mode > GLenum(TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if type_ != GLenum(UNSIGNED_BYTE) && type_ != GLenum(UNSIGNED_SHORT) && type_ != GLenum(UNSIGNED_INT) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if count < 0 || instancecount < 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	if count == 0 || instancecount == 0 {
		return
	}
	c.Buffers[c.Vertex_arrays[c.Cur_vertex_array].Element_buffer].Type = type_
	for instance := uint64(0); instance < uint64(instancecount); instance++ {
		run_pipeline(mode, GLint(offset), count, GLsizei(uint32(instance)), 0, TRUE)
	}
}
func DrawElementsInstancedBaseInstance(mode GLenum, count GLsizei, type_ GLenum, offset GLsizei, instancecount GLsizei, baseinstance GLuint) {
	if mode < GLenum(POINTS) || mode > GLenum(TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if type_ != GLenum(UNSIGNED_BYTE) && type_ != GLenum(UNSIGNED_SHORT) && type_ != GLenum(UNSIGNED_INT) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if count < 0 || instancecount < 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	if count == 0 || instancecount == 0 {
		return
	}
	c.Buffers[(c.Vertex_arrays[c.Cur_vertex_array]).Element_buffer].Type = type_
	for instance := uint64(0); instance < uint64(instancecount); instance++ {
		run_pipeline(mode, GLint(offset), count, GLsizei(uint32(instance)), baseinstance, TRUE)
	}
}
func Viewport(x int64, y int64, width GLsizei, height GLsizei) {
	if width < 0 || height < 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	make_viewport_matrix(&c.Vp_mat, x, y, uint64(width), uint64(height), 1)
	c.X_min = x
	c.Y_min = y
	c.X_max = uint64(x + int64(width))
	c.Y_max = uint64(y + int64(height))
}
func ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	red = GLclampf(clampf_01(float32(red)))
	green = GLclampf(clampf_01(float32(green)))
	blue = GLclampf(clampf_01(float32(blue)))
	alpha = GLclampf(clampf_01(float32(alpha)))
	var tmp Vec4 = Vec4{X: float32(red), Y: float32(green), Z: float32(blue), W: float32(alpha)}
	c.Clear_color = vec4_to_Color(tmp)
}
func ClearDepth(depth GLclampf) {
	c.Clear_depth = GLfloat(clampf_01(float32(depth)))
}
func DepthFunc(func_ GLenum) {
	if func_ < GLenum(LESS) || func_ > GLenum(NEVER) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	c.Depth_func = func_
}
func DepthRange(nearVal GLclampf, farVal GLclampf) {
	c.Depth_range_near = GLfloat(clampf_01(float32(nearVal)))
	c.Depth_range_far = GLfloat(clampf_01(float32(farVal)))
}
func DepthMask(flag GLboolean) {
	c.Depth_mask = flag
}
func Clear(mask GLbitfield) {
	if (mask & GLbitfield(COLOR_BUFFER_BIT|DEPTH_BUFFER_BIT|STENCIL_BUFFER_BIT)) == 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		println("failed to clear")
		return
	}
	var col Color = c.Clear_color
	if mask&GLbitfield(COLOR_BUFFER_BIT) != 0 {
		if c.Scissor_test == 0 {
			for i := int64(0); uint64(i) < c.Back_buffer.W*c.Back_buffer.H; i++ {
				*(*U32)(unsafe.Add(unsafe.Pointer((*U32)(unsafe.Pointer(&c.Back_buffer.Buf[0]))), unsafe.Sizeof(U32(0))*uintptr(i))) = U32(col.A)<<c.Ashift | U32(col.R)<<c.Rshift | U32(col.G)<<c.Gshift | U32(col.B)<<c.Bshift
			}
		} else {
			for y := int64(int64(c.Scissor_ly)); y < int64(c.Scissor_uy); y++ {
				for x := int64(int64(c.Scissor_lx)); x < int64(c.Scissor_ux); x++ {
					*(*U32)(unsafe.Add(unsafe.Pointer((*U32)(unsafe.Pointer(&c.Back_buffer.Lastrow[0]))), unsafe.Sizeof(U32(0))*uintptr(uint64(-y)*c.Back_buffer.W+uint64(x)))) = U32(col.A)<<c.Ashift | U32(col.R)<<c.Rshift | U32(col.G)<<c.Gshift | U32(col.B)<<c.Bshift
				}
			}
		}
	}
	if mask&GLbitfield(DEPTH_BUFFER_BIT) != 0 {
		if c.Scissor_test == 0 {
			for i := int64(0); uint64(i) < c.Zbuf.W*c.Zbuf.H; i++ {
				*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Buf[0]))), unsafe.Sizeof(float32(0))*uintptr(i))) = float32(c.Clear_depth)
			}
		} else {
			for y := int64(int64(c.Scissor_ly)); y < int64(c.Scissor_uy); y++ {
				for x := int64(int64(c.Scissor_lx)); x < int64(c.Scissor_ux); x++ {
					*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(&c.Zbuf.Lastrow[0]))), unsafe.Sizeof(float32(0))*uintptr(uint64(-y)*c.Zbuf.W+uint64(x)))) = float32(c.Clear_depth)
				}
			}
		}
	}
	if mask&GLbitfield(STENCIL_BUFFER_BIT) != 0 {
		if c.Scissor_test == 0 {
			for i := int64(0); uint64(i) < c.Stencil_buf.W*c.Stencil_buf.H; i++ {
				*(*u8)(unsafe.Add(unsafe.Pointer(&c.Stencil_buf.Buf[0]), i)) = u8(int8(c.Clear_stencil))
			}
		} else {
			for y := int64(int64(c.Scissor_ly)); y < int64(c.Scissor_uy); y++ {
				for x := int64(int64(c.Scissor_lx)); x < int64(c.Scissor_ux); x++ {
					*(*u8)(unsafe.Add(unsafe.Pointer(&c.Stencil_buf.Lastrow[0]), uint64(-y)*c.Stencil_buf.W+uint64(x))) = u8(int8(c.Clear_stencil))
				}
			}
		}
	}
}
func Enable(cap_ GLenum) {
	switch cap_ {
	case CULL_FACE:
		c.Cull_face = TRUE
	case DEPTH_TEST:
		c.Depth_test = TRUE
	case DEPTH_CLAMP:
		c.Depth_clamp = TRUE
	case LINE_SMOOTH:
	case BLEND:
		c.Blend = TRUE
	case COLOR_LOGIC_OP:
		c.Logic_ops = TRUE
	case POLYGON_OFFSET_FILL:
		c.Poly_offset = TRUE
	case SCISSOR_TEST:
		c.Scissor_test = TRUE
	case STENCIL_TEST:
		c.Stencil_test = TRUE
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
	}
}
func Disable(cap_ GLenum) {
	switch cap_ {
	case CULL_FACE:
		c.Cull_face = FALSE
	case DEPTH_TEST:
		c.Depth_test = FALSE
	case DEPTH_CLAMP:
		c.Depth_clamp = FALSE
	case LINE_SMOOTH:
		c.Line_smooth = FALSE
	case BLEND:
		c.Blend = FALSE
	case COLOR_LOGIC_OP:
		c.Logic_ops = FALSE
	case POLYGON_OFFSET_FILL:
		c.Poly_offset = FALSE
	case SCISSOR_TEST:
		c.Scissor_test = FALSE
	case STENCIL_TEST:
		c.Stencil_test = FALSE
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
	}
}
func IsEnabled(cap_ GLenum) GLboolean {
	switch cap_ {
	case DEPTH_TEST:
		return c.Depth_test
	case LINE_SMOOTH:
		return c.Line_smooth
	case CULL_FACE:
		return c.Cull_face
	case DEPTH_CLAMP:
		return c.Depth_clamp
	case BLEND:
		return c.Blend
	case COLOR_LOGIC_OP:
		return c.Logic_ops
	case POLYGON_OFFSET_FILL:
		return c.Poly_offset
	case SCISSOR_TEST:
		return c.Scissor_test
	case STENCIL_TEST:
		return c.Stencil_test
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
	}
	return FALSE
}
func GetBooleanv(pname GLenum, params *GLboolean) {
	switch pname {
	case DEPTH_TEST:
		*params = c.Depth_test
	case LINE_SMOOTH:
		*params = c.Line_smooth
	case CULL_FACE:
		*params = c.Cull_face
	case DEPTH_CLAMP:
		*params = c.Depth_clamp
	case BLEND:
		*params = c.Blend
	case COLOR_LOGIC_OP:
		*params = c.Logic_ops
	case POLYGON_OFFSET_FILL:
		*params = c.Poly_offset
	case SCISSOR_TEST:
		*params = c.Scissor_test
	case STENCIL_TEST:
		*params = c.Stencil_test
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
	}
}
func GetFloatv(pname GLenum, params *GLfloat) {
	switch pname {
	case POLYGON_OFFSET_FACTOR:
		*params = c.Poly_factor
	case POLYGON_OFFSET_UNITS:
		*params = c.Poly_units
	case POINT_SIZE:
		*params = c.Point_size
	case DEPTH_CLEAR_VALUE:
		*params = c.Clear_depth
	case DEPTH_RANGE:
		*(*GLfloat)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLfloat(0))*0)) = c.Depth_range_near
		*(*GLfloat)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLfloat(0))*1)) = c.Depth_range_near
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
	}
}
func GetIntegerv(pname GLenum, params *GLint) {
	switch pname {
	case STENCIL_WRITE_MASK:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_writemask)
	case STENCIL_REF:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = c.Stencil_ref
	case STENCIL_VALUE_MASK:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_valuemask)
	case STENCIL_FUNC:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_func)
	case STENCIL_FAIL:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_sfail)
	case STENCIL_PASS_DEPTH_FAIL:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_dpfail)
	case STENCIL_PASS_DEPTH_PASS:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_dppass)
	case STENCIL_BACK_WRITE_MASK:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_writemask_back)
	case STENCIL_BACK_REF:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = c.Stencil_ref_back
	case STENCIL_BACK_VALUE_MASK:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_valuemask_back)
	case STENCIL_BACK_FUNC:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_func_back)
	case STENCIL_BACK_FAIL:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_sfail_back)
	case STENCIL_BACK_PASS_DEPTH_FAIL:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_dpfail_back)
	case STENCIL_BACK_PASS_DEPTH_PASS:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_dppass_back)
	case LOGIC_OP_MODE:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Logic_func)
	case BLEND_SRC_RGB:
		fallthrough
	case BLEND_SRC_ALPHA:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Blend_sfactor)
	case BLEND_DST_RGB:
		fallthrough
	case BLEND_DST_ALPHA:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Blend_dfactor)
	case BLEND_EQUATION_RGB:
		fallthrough
	case BLEND_EQUATION_ALPHA:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Blend_equation)
	case CULL_FACE_MODE:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Cull_mode)
	case FRONT_FACE:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Front_face)
	case DEPTH_FUNC:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Depth_func)
	case POINT_SPRITE_COORD_ORIGIN:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Point_spr_origin)
		fallthrough
	case PROVOKING_VERTEX:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Provoking_vert)
	case POLYGON_MODE:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Poly_mode_front)
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*1)) = GLint(c.Poly_mode_back)
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
	}
}
func CullFace(mode GLenum) {
	switch mode {
	case FRONT, BACK, FRONT_AND_BACK:
		c.Cull_mode = mode
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
}
func FrontFace(mode GLenum) {
	switch mode {
	case CCW, CW:
		c.Front_face = mode
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
}
func PolygonMode(face GLenum, mode GLenum) {
	switch mode {
	case POINT:
		switch face {
		case FRONT:
			c.Poly_mode_front = mode
			c.Draw_triangle_front = draw_triangle_point
		case BACK:
			c.Poly_mode_back = mode
			c.Draw_triangle_back = draw_triangle_point
		case FRONT_AND_BACK:
			c.Poly_mode_front = mode
			c.Poly_mode_back = mode
			c.Draw_triangle_front = draw_triangle_point
			c.Draw_triangle_back = draw_triangle_point
		default:
			if c.Error == 0 {
				c.Error = GLenum(INVALID_ENUM)
			}
			return
		}
	case LINE:
		switch face {
		case FRONT:
			c.Poly_mode_front = mode
			c.Draw_triangle_front = draw_triangle_line
		case BACK:
			c.Poly_mode_back = mode
			c.Draw_triangle_back = draw_triangle_line
		case FRONT_AND_BACK:
			c.Poly_mode_front = mode
			c.Poly_mode_back = mode
			c.Draw_triangle_front = draw_triangle_line
			c.Draw_triangle_back = draw_triangle_line
		default:
			if c.Error == 0 {
				c.Error = GLenum(INVALID_ENUM)
			}
			return
		}
	case FILL:
		switch face {
		case FRONT:
			c.Poly_mode_front = mode
			c.Draw_triangle_front = draw_triangle_fill
		case BACK:
			c.Poly_mode_back = mode
			c.Draw_triangle_back = draw_triangle_fill
		case FRONT_AND_BACK:
			c.Poly_mode_front = mode
			c.Poly_mode_back = mode
			c.Draw_triangle_front = draw_triangle_fill

			c.Draw_triangle_back = draw_triangle_fill
		default:
			if c.Error == 0 {
				c.Error = GLenum(INVALID_ENUM)
			}
			return
		}
	default:
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}

}
func PointSize(size GLfloat) {
	if float64(size) <= 0.0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	c.Point_size = size
}
func PointParameteri(pname GLenum, param GLint) {
	if pname != GLenum(POINT_SPRITE_COORD_ORIGIN) || int64(param) != LOWER_LEFT && int64(param) != UPPER_LEFT {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	c.Point_spr_origin = GLenum(param)
}
func ProvokingVertex(provokeMode GLenum) {
	if provokeMode != GLenum(FIRST_VERTEX_CONVENTION) && provokeMode != GLenum(LAST_VERTEX_CONVENTION) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	c.Provoking_vert = provokeMode
}

func DeleteProgram(program GLuint) {
	if program == 0 {
		return
	}
	if uint64(program) >= uint64(len(c.Programs)) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	c.Programs[program].Deleted = TRUE
}
func UseProgram(program GLuint) {
	if uint64(program) >= uint64(len(c.Programs)) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	c.Vs_output.Size = c.Programs[program].Vs_output_size
	if int64(len(c.Vs_output.Output_buf)) < c.Vs_output.Size*MAX_VERTICES {
		var tmp = make([]float32, c.Vs_output.Size*MAX_VERTICES)
		copy(tmp, c.Vs_output.Output_buf)
		c.Vs_output.Output_buf = tmp
	}
	c.Vs_output.Interpolation = c.Programs[program].Interpolation[:]
	c.Fragdepth_or_discard = c.Programs[program].Fragdepth_or_discard
	c.Cur_program = program
}

func BlendFunc(sfactor GLenum, dfactor GLenum) {
	if sfactor < GLenum(ZERO) || sfactor >= GLenum(NUM_BLEND_FUNCS) || dfactor < GLenum(ZERO) || dfactor >= GLenum(NUM_BLEND_FUNCS) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	c.Blend_sfactor = sfactor
	c.Blend_dfactor = dfactor
}
func BlendEquation(mode GLenum) {
	if mode < GLenum(FUNC_ADD) || mode >= GLenum(NUM_BLEND_EQUATIONS) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	c.Blend_equation = mode
}
func BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	c.Blend_color.X = clampf_01(float32(red))
	c.Blend_color.Y = clampf_01(float32(green))
	c.Blend_color.Z = clampf_01(float32(blue))
	c.Blend_color.W = clampf_01(float32(alpha))
}
func LogicOp(opcode GLenum) {
	if opcode < GLenum(CLEAR) || opcode > GLenum(INVERT) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	c.Logic_func = opcode
}
func PolygonOffset(factor GLfloat, units GLfloat) {
	c.Poly_factor = factor
	c.Poly_units = units
}
func Scissor(x GLint, y GLint, width GLsizei, height GLsizei) {
	if width < 0 || height < 0 {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_VALUE)
		}
		return
	}
	c.Scissor_lx = x
	c.Scissor_ly = y
	c.Scissor_ux = GLsizei(x + GLint(width))
	c.Scissor_uy = GLsizei(y + GLint(height))
}
func StencilFunc(func_ GLenum, ref GLint, mask GLuint) {
	if func_ < GLenum(LESS) || func_ > GLenum(NEVER) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	c.Stencil_func = func_
	c.Stencil_func_back = func_
	if ref > math.MaxUint8 {
		ref = math.MaxUint8
	}
	if ref < 0 {
		ref = 0
	}
	c.Stencil_ref = ref
	c.Stencil_ref_back = ref
	c.Stencil_valuemask = mask
	c.Stencil_valuemask_back = mask
}
func StencilFuncSeparate(face GLenum, func_ GLenum, ref GLint, mask GLuint) {
	if face < GLenum(FRONT) || face > GLenum(FRONT_AND_BACK) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if face == GLenum(FRONT_AND_BACK) {
		StencilFunc(func_, ref, mask)
		return
	}
	if func_ < GLenum(LESS) || func_ > GLenum(NEVER) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if ref > math.MaxUint8 {
		ref = math.MaxUint8
	}
	if ref < 0 {
		ref = 0
	}
	if face == GLenum(FRONT) {
		c.Stencil_func = func_
		c.Stencil_ref = ref
		c.Stencil_valuemask = mask
	} else {
		c.Stencil_func_back = func_
		c.Stencil_ref_back = ref
		c.Stencil_valuemask_back = mask
	}
}
func StencilOp(sfail GLenum, dpfail GLenum, dppass GLenum) {
	if (sfail < GLenum(INVERT) || sfail > GLenum(DECR_WRAP)) && sfail != GLenum(ZERO) || (dpfail < GLenum(INVERT) || dpfail > GLenum(DECR_WRAP)) && sfail != GLenum(ZERO) || (dppass < GLenum(INVERT) || dppass > GLenum(DECR_WRAP)) && sfail != GLenum(ZERO) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	c.Stencil_sfail = sfail
	c.Stencil_dpfail = dpfail
	c.Stencil_dppass = dppass
	c.Stencil_sfail_back = sfail
	c.Stencil_dpfail_back = dpfail
	c.Stencil_dppass_back = dppass
}
func StencilOpSeparate(face GLenum, sfail GLenum, dpfail GLenum, dppass GLenum) {
	if face < GLenum(FRONT) || face > GLenum(FRONT_AND_BACK) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if face == GLenum(FRONT_AND_BACK) {
		StencilOp(sfail, dpfail, dppass)
		return
	}
	if (sfail < GLenum(INVERT) || sfail > GLenum(DECR_WRAP)) && sfail != GLenum(ZERO) || (dpfail < GLenum(INVERT) || dpfail > GLenum(DECR_WRAP)) && sfail != GLenum(ZERO) || (dppass < GLenum(INVERT) || dppass > GLenum(DECR_WRAP)) && sfail != GLenum(ZERO) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if face == GLenum(FRONT) {
		c.Stencil_sfail = sfail
		c.Stencil_dpfail = dpfail
		c.Stencil_dppass = dppass
	} else {
		c.Stencil_sfail_back = sfail
		c.Stencil_dpfail_back = dpfail
		c.Stencil_dppass_back = dppass
	}
}
func ClearStencil(s GLint) {
	c.Clear_stencil = s & math.MaxUint8
}
func StencilMask(mask GLuint) {
	c.Stencil_writemask = mask
	c.Stencil_writemask_back = mask
}
func StencilMaskSeparate(face GLenum, mask GLuint) {
	if face < GLenum(FRONT) || face > GLenum(FRONT_AND_BACK) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return
	}
	if face == GLenum(FRONT_AND_BACK) {
		StencilMask(mask)
		return
	}
	if face == GLenum(FRONT) {
		c.Stencil_writemask = mask
	} else {
		c.Stencil_writemask_back = mask
	}
}
func MapBuffer(target GLenum, access GLenum) unsafe.Pointer {
	if target != GLenum(ARRAY_BUFFER) && target != GLenum(ELEMENT_ARRAY_BUFFER) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return nil
	}
	if access != GLenum(READ_ONLY) && access != GLenum(WRITE_ONLY) && access != GLenum(READ_WRITE) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return nil
	}
	target -= GLenum(ARRAY_BUFFER)
	var data = pglGetBufferData(c.Bound_buffers[target])
	return unsafe.Pointer(&data[0])
}
func MapNamedBuffer(buffer GLuint, access GLenum) unsafe.Pointer {
	if access != GLenum(READ_ONLY) && access != GLenum(WRITE_ONLY) && access != GLenum(READ_WRITE) {
		if c.Error == 0 {
			c.Error = GLenum(INVALID_ENUM)
		}
		return nil
	}
	var data = pglGetBufferData(buffer)
	return unsafe.Pointer(&data[0])
}
func GetDoublev(pname GLenum, params *GLdouble) {
}
func GetInteger64v(pname GLenum, params *GLint64) {
}
func GetProgramiv(program GLuint, pname GLenum, params *GLint) {
}
func GetProgramInfoLog(program GLuint, maxLength GLsizei, length *GLsizei, infoLog *byte) {
}
func AttachShader(program GLuint, shader GLuint) {
}
func CompileShader(shader GLuint) {
}
func GetShaderInfoLog(shader GLuint, maxLength GLsizei, length *GLsizei, infoLog *byte) {
}
func LinkProgram(program GLuint) {
}
func ShaderSource(shader GLuint, count GLsizei, string_ **byte, length *GLint) {
}
func GetShaderiv(shader GLuint, pname GLenum, params *GLint) {
}
func DeleteShader(shader GLuint) {
}
func DetachShader(program GLuint, shader GLuint) {
}
func CreateProgram() GLuint {
	return 0
}
func CreateShader(shaderType GLenum) GLuint {
	return 0
}
func GetUniformLocation(program GLuint, name *byte) GLint {
	return 0
}
func GetAttribLocation(program GLuint, name *byte) GLint {
	return 0
}
func UnmapBuffer(target GLenum) GLboolean {
	return TRUE
}
func UnmapNamedBuffer(buffer GLuint) GLboolean {
	return TRUE
}
func LineWidth(width GLfloat) {
}
func ActiveTexture(texture GLenum) {
}
func TexParameterfv(target GLenum, pname GLenum, params *GLfloat) {
}
func Uniform1f(location GLint, v0 GLfloat) {
}
func Uniform2f(location GLint, v0 GLfloat, v1 GLfloat) {
}
func Uniform3f(location GLint, v0 GLfloat, v1 GLfloat, v2 GLfloat) {
}
func Uniform4f(location GLint, v0 GLfloat, v1 GLfloat, v2 GLfloat, v3 GLfloat) {
}
func Uniform1i(location GLint, v0 GLint) {
}
func Uniform2i(location GLint, v0 GLint, v1 GLint) {
}
func Uniform3i(location GLint, v0 GLint, v1 GLint, v2 GLint) {
}
func Uniform4i(location GLint, v0 GLint, v1 GLint, v2 GLint, v3 GLint) {
}
func Uniform1ui(location GLuint, v0 GLuint) {
}
func Uniform2ui(location GLuint, v0 GLuint, v1 GLuint) {
}
func Uniform3ui(location GLuint, v0 GLuint, v1 GLuint, v2 GLuint) {
}
func Uniform4ui(location GLuint, v0 GLuint, v1 GLuint, v2 GLuint, v3 GLuint) {
}
func Uniform1fv(location GLint, count GLsizei, value *GLfloat) {
}
func Uniform2fv(location GLint, count GLsizei, value *GLfloat) {
}
func Uniform3fv(location GLint, count GLsizei, value *GLfloat) {
}
func Uniform4fv(location GLint, count GLsizei, value *GLfloat) {
}
func Uniform1iv(location GLint, count GLsizei, value *GLint) {
}
func Uniform2iv(location GLint, count GLsizei, value *GLint) {
}
func Uniform3iv(location GLint, count GLsizei, value *GLint) {
}
func Uniform4iv(location GLint, count GLsizei, value *GLint) {
}
func Uniform1uiv(location GLint, count GLsizei, value *GLuint) {
}
func Uniform2uiv(location GLint, count GLsizei, value *GLuint) {
}
func Uniform3uiv(location GLint, count GLsizei, value *GLuint) {
}
func Uniform4uiv(location GLint, count GLsizei, value *GLuint) {
}
func UniformMatrix2fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) {
}
func UniformMatrix3fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) {
}
func UniformMatrix4fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) {
}
func UniformMatrix2x3fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) {
}
func UniformMatrix3x2fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) {
}
func UniformMatrix2x4fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) {
}
func UniformMatrix4x2fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) {
}
func UniformMatrix3x4fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) {
}
func UniformMatrix4x3fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) {
}
func clampf_01(f float32) float32 {
	if float64(f) < 0.0 {
		return 0.0
	}
	if float64(f) > 1.0 {
		return 1.0
	}
	return f
}
func clampf(f float32, min float32, max float32) float32 {
	if f < min {
		return min
	}
	if f > max {
		return max
	}
	return f
}
func clampi(i int64, min int64, max int64) int64 {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}
func wrap(i int64, size int64, mode GLenum) int64 {
	var (
		tmp  int64
		tmp2 int64
	)
	_ = tmp2
	switch mode {
	case REPEAT:
		tmp = i - size*(i/size)
		if tmp < 0 {
			tmp = size + tmp
		}
		return tmp
	case CLAMP_TO_BORDER:
		fallthrough
	case CLAMP_TO_EDGE:
		return clampi(i, 0, size-1)
	case MIRRORED_REPEAT:
		if i < 0 {
			i = -i
		}
		tmp = i / size
		tmp2 = i / (size * 2)
		if tmp%2 != 0 {
			return (size - 1) - (i - tmp*size)
		} else {
			return i - tmp*size
		}
		return tmp
	default:
		panic("shouldn't be here")
		return 0
	}
}
func texture1D(tex GLuint, x float32) Vec4 {
	var (
		i0      int64
		i1      int64
		t       *glTexture = &c.Textures[tex]
		texdata *Color     = (*Color)(unsafe.Pointer(&t.Data[0]))
		w       float64    = float64(t.W) - EPSILON
		xw      float64    = float64(x) * w
	)
	if t.Mag_filter == GLenum(NEAREST) {
		i0 = wrap(int64(math.Floor(xw)), int64(t.W), t.Wrap_s)
		return Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(i0))))
	} else {
		i0 = wrap(int64(math.Floor(xw-0.5)), int64(t.W), t.Wrap_s)
		i1 = wrap(int64(math.Floor(xw+0.499999)), int64(t.W), t.Wrap_s)
		var tmp2 float64
		var alpha float64 = cmath.Modf(xw+0.5, &tmp2)
		if alpha < 0 {
			alpha++
		}
		var ci Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(i0))))
		var ci1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(i1))))
		ci = scale_vec4(ci, float32(1-alpha))
		ci1 = scale_vec4(ci1, float32(alpha))
		ci = add_vec4s(ci, ci1)
		return ci
	}
}
func texture2D(tex GLuint, x float32, y float32) Vec4 {
	var (
		i0      int64
		j0      int64
		i1      int64
		j1      int64
		t       *glTexture = &c.Textures[tex]
		texdata *Color     = (*Color)(unsafe.Pointer(&t.Data[0]))
		w       int64      = int64(t.W)
		h       int64      = int64(t.H)
		dw      float64    = float64(w) - EPSILON
		dh      float64    = float64(h) - EPSILON
		xw      float64    = float64(x) * dw
		yh      float64    = float64(y) * dh
	)
	if t.Mag_filter == GLenum(NEAREST) {
		i0 = wrap(int64(math.Floor(xw)), w, t.Wrap_s)
		j0 = wrap(int64(math.Floor(yh)), h, t.Wrap_t)
		return Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(j0*w+i0))))
	} else {
		i0 = wrap(int64(math.Floor(xw-0.5)), w, t.Wrap_s)
		j0 = wrap(int64(math.Floor(yh-0.5)), h, t.Wrap_t)
		i1 = wrap(int64(math.Floor(xw+0.499999)), w, t.Wrap_s)
		j1 = wrap(int64(math.Floor(yh+0.499999)), h, t.Wrap_t)
		var tmp2 float64
		var alpha float64 = cmath.Modf(xw+0.5, &tmp2)
		var beta float64 = cmath.Modf(yh+0.5, &tmp2)
		if alpha < 0 {
			alpha++
		}
		if beta < 0 {
			beta++
		}
		var cij Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(j0*w+i0))))
		var ci1j Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(j0*w+i1))))
		var cij1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(j1*w+i0))))
		var ci1j1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(j1*w+i1))))
		cij = scale_vec4(cij, float32((1-alpha)*(1-beta)))
		ci1j = scale_vec4(ci1j, float32(alpha*(1-beta)))
		cij1 = scale_vec4(cij1, float32((1-alpha)*beta))
		ci1j1 = scale_vec4(ci1j1, float32(alpha*beta))
		cij = add_vec4s(cij, ci1j)
		cij = add_vec4s(cij, cij1)
		cij = add_vec4s(cij, ci1j1)
		return cij
	}
}
func texture3D(tex GLuint, x float32, y float32, z float32) Vec4 {
	var (
		i0      int64
		j0      int64
		i1      int64
		j1      int64
		k0      int64
		k1      int64
		t       *glTexture = &c.Textures[tex]
		texdata *Color     = (*Color)(unsafe.Pointer(&t.Data[0]))
		dw      float64    = float64(t.W) - EPSILON
		dh      float64    = float64(t.H) - EPSILON
		dd      float64    = float64(t.D) - EPSILON
		w       int64      = int64(t.W)
		h       int64      = int64(t.H)
		d       int64      = int64(t.D)
		plane   int64      = int64(uint64(w) * t.H)
		xw      float64    = float64(x) * dw
		yh      float64    = float64(y) * dh
		zd      float64    = float64(z) * dd
	)
	if t.Mag_filter == GLenum(NEAREST) {
		i0 = wrap(int64(math.Floor(xw)), w, t.Wrap_s)
		j0 = wrap(int64(math.Floor(yh)), h, t.Wrap_t)
		k0 = wrap(int64(math.Floor(zd)), d, t.Wrap_r)
		return Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(k0*plane+j0*w+i0))))
	} else {
		i0 = wrap(int64(math.Floor(xw-0.5)), w, t.Wrap_s)
		j0 = wrap(int64(math.Floor(yh-0.5)), h, t.Wrap_t)
		k0 = wrap(int64(math.Floor(zd-0.5)), d, t.Wrap_r)
		i1 = wrap(int64(math.Floor(xw+0.499999)), w, t.Wrap_s)
		j1 = wrap(int64(math.Floor(yh+0.499999)), h, t.Wrap_t)
		k1 = wrap(int64(math.Floor(zd+0.499999)), d, t.Wrap_r)
		var tmp2 float64
		var alpha float64 = cmath.Modf(xw+0.5, &tmp2)
		var beta float64 = cmath.Modf(yh+0.5, &tmp2)
		var gamma float64 = cmath.Modf(zd+0.5, &tmp2)
		if alpha < 0 {
			alpha++
		}
		if beta < 0 {
			beta++
		}
		if gamma < 0 {
			gamma++
		}
		var cijk Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(k0*plane+j0*w+i0))))
		var ci1jk Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(k0*plane+j0*w+i1))))
		var cij1k Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(k0*plane+j1*w+i0))))
		var ci1j1k Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(k0*plane+j1*w+i1))))
		var cijk1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(k1*plane+j0*w+i0))))
		var ci1jk1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(k1*plane+j0*w+i1))))
		var cij1k1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(k1*plane+j1*w+i0))))
		var ci1j1k1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(k1*plane+j1*w+i1))))
		cijk = scale_vec4(cijk, float32((1-alpha)*(1-beta)*(1-gamma)))
		ci1jk = scale_vec4(ci1jk, float32(alpha*(1-beta)*(1-gamma)))
		cij1k = scale_vec4(cij1k, float32((1-alpha)*beta*(1-gamma)))
		ci1j1k = scale_vec4(ci1j1k, float32(alpha*beta*(1-gamma)))
		cijk1 = scale_vec4(cijk1, float32((1-alpha)*(1-beta)*gamma))
		ci1jk1 = scale_vec4(ci1jk1, float32(alpha*(1-beta)*gamma))
		cij1k1 = scale_vec4(cij1k1, float32((1-alpha)*beta*gamma))
		ci1j1k1 = scale_vec4(ci1j1k1, float32(alpha*beta*gamma))
		cijk = add_vec4s(cijk, ci1jk)
		cijk = add_vec4s(cijk, cij1k)
		cijk = add_vec4s(cijk, ci1j1k)
		cijk = add_vec4s(cijk, cijk1)
		cijk = add_vec4s(cijk, ci1jk1)
		cijk = add_vec4s(cijk, cij1k1)
		cijk = add_vec4s(cijk, ci1j1k1)
		return cijk
	}
}
func texture2DArray(tex GLuint, x float32, y float32, z int64) Vec4 {
	var (
		i0      int64
		j0      int64
		i1      int64
		j1      int64
		t       *glTexture = &c.Textures[tex]
		texdata *Color     = (*Color)(unsafe.Pointer(&t.Data[0]))
		w       int64      = int64(t.W)
		h       int64      = int64(t.H)
		dw      float64    = float64(w) - EPSILON
		dh      float64    = float64(h) - EPSILON
		plane   int64      = w * h
		xw      float64    = float64(x) * dw
		yh      float64    = float64(y) * dh
	)
	if t.Mag_filter == GLenum(NEAREST) {
		i0 = wrap(int64(math.Floor(xw)), w, t.Wrap_s)
		j0 = wrap(int64(math.Floor(yh)), h, t.Wrap_t)
		return Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(z*plane+j0*w+i0))))
	} else {
		i0 = wrap(int64(math.Floor(xw-0.5)), w, t.Wrap_s)
		j0 = wrap(int64(math.Floor(yh-0.5)), h, t.Wrap_t)
		i1 = wrap(int64(math.Floor(xw+0.499999)), w, t.Wrap_s)
		j1 = wrap(int64(math.Floor(yh+0.499999)), h, t.Wrap_t)
		var tmp2 float64
		var alpha float64 = cmath.Modf(xw+0.5, &tmp2)
		var beta float64 = cmath.Modf(yh+0.5, &tmp2)
		if alpha < 0 {
			alpha++
		}
		if beta < 0 {
			beta++
		}
		var cij Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(z*plane+j0*w+i0))))
		var ci1j Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(z*plane+j0*w+i1))))
		var cij1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(z*plane+j1*w+i0))))
		var ci1j1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(z*plane+j1*w+i1))))
		cij = scale_vec4(cij, float32((1-alpha)*(1-beta)))
		ci1j = scale_vec4(ci1j, float32(alpha*(1-beta)))
		cij1 = scale_vec4(cij1, float32((1-alpha)*beta))
		ci1j1 = scale_vec4(ci1j1, float32(alpha*beta))
		cij = add_vec4s(cij, ci1j)
		cij = add_vec4s(cij, cij1)
		cij = add_vec4s(cij, ci1j1)
		return cij
	}
}
func texture_rect(tex GLuint, x float32, y float32) Vec4 {
	var (
		i0      int64
		j0      int64
		i1      int64
		j1      int64
		t       *glTexture = &c.Textures[tex]
		texdata *Color     = (*Color)(unsafe.Pointer(&t.Data[0]))
		w       int64      = int64(t.W)
		h       int64      = int64(t.H)
		xw      float64    = float64(x)
		yh      float64    = float64(y)
	)
	if t.Mag_filter == GLenum(NEAREST) {
		i0 = wrap(int64(math.Floor(xw)), w, t.Wrap_s)
		j0 = wrap(int64(math.Floor(yh)), h, t.Wrap_t)
		return Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(j0*w+i0))))
	} else {
		i0 = wrap(int64(math.Floor(xw-0.5)), w, t.Wrap_s)
		j0 = wrap(int64(math.Floor(yh-0.5)), h, t.Wrap_t)
		i1 = wrap(int64(math.Floor(xw+0.499999)), w, t.Wrap_s)
		j1 = wrap(int64(math.Floor(yh+0.499999)), h, t.Wrap_t)
		var tmp2 float64
		var alpha float64 = cmath.Modf(xw+0.5, &tmp2)
		var beta float64 = cmath.Modf(yh+0.5, &tmp2)
		if alpha < 0 {
			alpha++
		}
		if beta < 0 {
			beta++
		}
		var cij Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(j0*w+i0))))
		var ci1j Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(j0*w+i1))))
		var cij1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(j1*w+i0))))
		var ci1j1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(j1*w+i1))))
		cij = scale_vec4(cij, float32((1-alpha)*(1-beta)))
		ci1j = scale_vec4(ci1j, float32(alpha*(1-beta)))
		cij1 = scale_vec4(cij1, float32((1-alpha)*beta))
		ci1j1 = scale_vec4(ci1j1, float32(alpha*beta))
		cij = add_vec4s(cij, ci1j)
		cij = add_vec4s(cij, cij1)
		cij = add_vec4s(cij, ci1j1)
		return cij
	}
}
func texture_cubemap(texture GLuint, x float32, y float32, z float32) Vec4 {
	var (
		tex     *glTexture = &c.Textures[texture]
		texdata *Color     = (*Color)(unsafe.Pointer(&tex.Data[0]))
		x_mag   float32
	)
	if x < 0 {
		x_mag = -x
	} else {
		x_mag = x
	}
	var y_mag float32
	if y < 0 {
		y_mag = -y
	} else {
		y_mag = y
	}
	var z_mag float32
	if z < 0 {
		z_mag = -z
	} else {
		z_mag = z
	}
	var s float32
	var t float32
	var max float32
	var p int64
	var i0 int64
	var j0 int64
	var i1 int64
	var j1 int64
	if x_mag > y_mag {
		if x_mag > z_mag {
			max = x_mag
			t = -y
			if x_mag == x {
				p = 0
				s = -z
			} else {
				p = 1
				s = z
			}
		} else {
			max = z_mag
			t = -y
			if z_mag == z {
				p = 4
				s = x
			} else {
				p = 5
				s = -x
			}
		}
	} else {
		if y_mag > z_mag {
			max = y_mag
			s = x
			if y_mag == y {
				p = 2
				t = z
			} else {
				p = 3
				t = -z
			}
		} else {
			max = z_mag
			t = -y
			if z_mag == z {
				p = 4
				s = x
			} else {
				p = 5
				s = -x
			}
		}
	}
	x = float32((float64(s/max) + 1.0) / 2.0)
	y = float32((float64(t/max) + 1.0) / 2.0)
	var w int64 = int64(tex.W)
	var h int64 = int64(tex.H)
	var dw float64 = float64(w) - EPSILON
	var dh float64 = float64(h) - EPSILON
	var plane int64 = w * w
	var xw float64 = float64(x) * dw
	var yh float64 = float64(y) * dh
	if tex.Mag_filter == GLenum(NEAREST) {
		i0 = wrap(int64(math.Floor(xw)), w, tex.Wrap_s)
		j0 = wrap(int64(math.Floor(yh)), h, tex.Wrap_t)
		var tmpvec4 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(p*plane+j0*w+i0))))
		return tmpvec4
	} else {
		i0 = wrap(int64(math.Floor(xw-0.5)), int64(tex.W), tex.Wrap_s)
		j0 = wrap(int64(math.Floor(yh-0.5)), int64(tex.H), tex.Wrap_t)
		i1 = wrap(int64(math.Floor(xw+0.499999)), int64(tex.W), tex.Wrap_s)
		j1 = wrap(int64(math.Floor(yh+0.499999)), int64(tex.H), tex.Wrap_t)
		var tmp2 float64
		var alpha float64 = cmath.Modf(xw+0.5, &tmp2)
		var beta float64 = cmath.Modf(yh+0.5, &tmp2)
		if alpha < 0 {
			alpha++
		}
		if beta < 0 {
			beta++
		}
		var cij Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(p*plane+j0*w+i0))))
		var ci1j Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(p*plane+j0*w+i1))))
		var cij1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(p*plane+j1*w+i0))))
		var ci1j1 Vec4 = Color_to_vec4(*(*Color)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(Color{})*uintptr(p*plane+j1*w+i1))))
		cij = scale_vec4(cij, float32((1-alpha)*(1-beta)))
		ci1j = scale_vec4(ci1j, float32(alpha*(1-beta)))
		cij1 = scale_vec4(cij1, float32((1-alpha)*beta))
		ci1j1 = scale_vec4(ci1j1, float32(alpha*beta))
		cij = add_vec4s(cij, ci1j)
		cij = add_vec4s(cij, cij1)
		cij = add_vec4s(cij, ci1j1)
		return cij
	}
}

func put_pixel(color Color, x int64, y int64) {
	var dest *U32 = (*U32)(unsafe.Add(unsafe.Pointer((*U32)(unsafe.Pointer(&c.Back_buffer.Lastrow[0]))), unsafe.Sizeof(U32(0))*uintptr(uint64(-y)*c.Back_buffer.W+uint64(x))))
	_ = dest
	*dest = U32(int32(int64(color.A)<<c.Ashift | int64(color.R)<<c.Rshift | int64(color.G)<<c.Gshift | int64(color.B)<<c.Bshift))
}
func put_line(the_color Color, x1 float32, y1 float32, x2 float32, y2 float32) {
	var tmp float32
	if x2 < x1 {
		tmp = x1
		x1 = x2
		x2 = tmp
		tmp = y1
		y1 = y2
		y2 = tmp
	}
	var m float32 = (y2 - y1) / (x2 - x1)
	var A float32 = y1 - y2
	var B float32 = x2 - x1
	var C float32 = x1*y2 - x2*y1
	var x int64
	var y int64
	var x_min float32
	if 0 > (func() float32 {
		if x1 < x2 {
			return x1
		}
		return x2
	}()) {
		x_min = 0
	} else if x1 < x2 {
		x_min = x1
	} else {
		x_min = x2
	}
	var x_max float32
	if float32(c.Back_buffer.W-1) < (func() float32 {
		if x1 > x2 {
			return x1
		}
		return x2
	}()) {
		x_max = float32(c.Back_buffer.W - 1)
	} else if x1 > x2 {
		x_max = x1
	} else {
		x_max = x2
	}
	var y_min float32
	if 0 > (func() float32 {
		if y1 < y2 {
			return y1
		}
		return y2
	}()) {
		y_min = 0
	} else if y1 < y2 {
		y_min = y1
	} else {
		y_min = y2
	}
	var y_max float32
	if float32(c.Back_buffer.H-1) < (func() float32 {
		if y1 > y2 {
			return y1
		}
		return y2
	}()) {
		y_max = float32(c.Back_buffer.H - 1)
	} else if y1 > y2 {
		y_max = y1
	} else {
		y_max = y2
	}
	if m <= float32(-1) {
		x = int64(x1)
		for y = int64(y_max); float32(y) >= y_min; y-- {
			put_pixel(the_color, x, y)
			if float64(A)*(float64(x)+0.5)+float64(B*float32(y-1))+float64(C) < 0 {
				x++
			}
		}
	} else if m <= 0 {
		y = int64(y1)
		for x = int64(x_min); float32(x) <= x_max; x++ {
			put_pixel(the_color, x, y)
			if float64(A*float32(x+1))+float64(B)*(float64(y)-0.5)+float64(C) > 0 {
				y--
			}
		}
	} else if m <= 1 {
		y = int64(y1)
		for x = int64(x_min); float32(x) <= x_max; x++ {
			put_pixel(the_color, x, y)
			if float64(A*float32(x+1))+float64(B)*(float64(y)+0.5)+float64(C) < 0 {
				y++
			}
		}
	} else {
		x = int64(x1)
		for y = int64(y_min); float32(y) <= y_max; y++ {
			put_pixel(the_color, x, y)
			if float64(A)*(float64(x)+0.5)+float64(B*float32(y+1))+float64(C) > 0 {
				x++
			}
		}
	}
}
func put_triangle(c1 Color, c2 Color, c3 Color, p1 vec2, p2 vec2, p3 vec2) {
	var x_min float32
	if math.Floor(float64(p1.X)) < math.Floor(float64(p2.X)) {
		x_min = float32(math.Floor(float64(p1.X)))
	} else {
		x_min = float32(math.Floor(float64(p2.X)))
	}
	var x_max float32
	if math.Ceil(float64(p1.X)) > math.Ceil(float64(p2.X)) {
		x_max = float32(math.Ceil(float64(p1.X)))
	} else {
		x_max = float32(math.Ceil(float64(p2.X)))
	}
	var y_min float32
	if math.Floor(float64(p1.Y)) < math.Floor(float64(p2.Y)) {
		y_min = float32(math.Floor(float64(p1.Y)))
	} else {
		y_min = float32(math.Floor(float64(p2.Y)))
	}
	var y_max float32
	if math.Ceil(float64(p1.Y)) > math.Ceil(float64(p2.Y)) {
		y_max = float32(math.Ceil(float64(p1.Y)))
	} else {
		y_max = float32(math.Ceil(float64(p2.Y)))
	}
	if math.Floor(float64(p3.X)) < float64(x_min) {
		x_min = float32(math.Floor(float64(p3.X)))
	} else {
		x_min = x_min
	}
	if math.Ceil(float64(p3.X)) > float64(x_max) {
		x_max = float32(math.Ceil(float64(p3.X)))
	} else {
		x_max = x_max
	}
	if math.Floor(float64(p3.Y)) < float64(y_min) {
		y_min = float32(math.Floor(float64(p3.Y)))
	} else {
		y_min = y_min
	}
	if math.Ceil(float64(p3.Y)) > float64(y_max) {
		y_max = float32(math.Ceil(float64(p3.Y)))
	} else {
		y_max = y_max
	}
	if 0 > x_min {
		x_min = 0
	} else {
		x_min = x_min
	}
	if float32(c.Back_buffer.W-1) < x_max {
		x_max = float32(c.Back_buffer.W - 1)
	} else {
		x_max = x_max
	}
	if 0 > y_min {
		y_min = 0
	} else {
		y_min = y_min
	}
	if float32(c.Back_buffer.H-1) < y_max {
		y_max = float32(c.Back_buffer.H - 1)
	} else {
		y_max = y_max
	}
	var l12 Line = make_Line(p1.X, p1.Y, p2.X, p2.Y)
	var l23 Line = make_Line(p2.X, p2.Y, p3.X, p3.Y)
	var l31 Line = make_Line(p3.X, p3.Y, p1.X, p1.Y)
	var alpha float32
	var beta float32
	var gamma float32
	var c Color
	var x float32
	var y float32
	for y = y_min; y <= y_max; y++ {
		for x = x_min; x <= x_max; x++ {
			gamma = line_func(&l12, x, y) / line_func(&l12, p3.X, p3.Y)
			beta = line_func(&l31, x, y) / line_func(&l31, p2.X, p2.Y)
			alpha = 1 - beta - gamma
			if alpha >= 0 && beta >= 0 && gamma >= 0 {
				if (alpha > 0 || line_func(&l23, p1.X, p1.Y)*line_func(&l23, float32(-1), float32(-1)) > 0) && (beta > 0 || line_func(&l31, p2.X, p2.Y)*line_func(&l31, float32(-1), float32(-1)) > 0) && (gamma > 0 || line_func(&l12, p3.X, p3.Y)*line_func(&l12, float32(-1), float32(-1)) > 0) {
					c.R = u8(alpha*float32(c1.R) + beta*float32(c2.R) + gamma*float32(c3.R))
					c.G = u8(alpha*float32(c1.G) + beta*float32(c2.G) + gamma*float32(c3.G))
					c.B = u8(alpha*float32(c1.B) + beta*float32(c2.B) + gamma*float32(c3.B))
					put_pixel(c, int64(x), int64(y))
				}
			}
		}
	}
}
