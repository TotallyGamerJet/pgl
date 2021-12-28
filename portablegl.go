package pgl

import (
	"fmt"
	"github.com/gotranspile/cxgo/runtime/cmath"
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"math"
	"maze.io/x/math32"
	"unsafe"
)

const RM_PI = 3.14159265358979323846
const RM_2PI = 2.0 * RM_PI
const PI_DIV_180 = 0.017453292519943296
const INV_PI_DIV_180 = 57.2957795130823229
const GL_FALSE = 0
const GL_TRUE = 1
const MAX_VERTICES = 500000
const GL_MAX_VERTEX_ATTRIBS = 16
const GL_MAX_VERTEX_OUTPUT_COMPONENTS = 64
const GL_MAX_DRAW_BUFFERS = 8
const CLIP_EPSILON = 1e-5

type u8 uint8
type u16 uint16
type U32 uint32
type u64 uint64
type i8 int8
type i16 int16
type i32 int32
type i64 int64

func rsw_rand_float(min float32, max float32) float32 {
	return (float32(libc.Rand())/(float32(libc.RandMax)-1))*(max-min) + min
}

type vec2 struct {
	X float32
	Y float32
}
type Vec3 struct {
	X float32
	Y float32
	Z float32
}
type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func make_vec2(x float32, y float32) vec2 {
	var v vec2 = vec2{X: x, Y: y}
	return v
}
func make_vec3(x float32, y float32, z float32) Vec3 {
	var v Vec3 = Vec3{X: x, Y: y, Z: z}
	return v
}
func make_vec4(x float32, y float32, z float32, w float32) Vec4 {
	var v Vec4 = Vec4{X: x, Y: y, Z: z, W: w}
	return v
}
func negate_vec2(v vec2) vec2 {
	var r vec2 = vec2{X: -v.X, Y: -v.Y}
	return r
}
func negate_vec3(v Vec3) Vec3 {
	var r Vec3 = Vec3{X: -v.X, Y: -v.Y, Z: -v.Z}
	return r
}
func negate_vec4(v Vec4) Vec4 {
	var r Vec4 = Vec4{X: -v.X, Y: -v.Y, Z: -v.Z, W: -v.W}
	return r
}
func fprint_vec2(f *stdio.File, v vec2, append *byte) {
	stdio.Fprintf(f, "(%f, %f)%s", v.X, v.Y, append)
}
func fprint_vec3(f *stdio.File, v Vec3, append *byte) {
	stdio.Fprintf(f, "(%f, %f, %f)%s", v.X, v.Y, v.Z, append)
}
func fprint_vec4(f *stdio.File, v Vec4, append *byte) {
	stdio.Fprintf(f, "(%f, %f, %f, %f)%s", v.X, v.Y, v.Z, v.W, append)
}
func print_vec2(v vec2, append *byte) {
	stdio.Printf("(%f, %f)%s", v.X, v.Y, append)
}
func print_vec3(v Vec3, append *byte) {
	stdio.Printf("(%f, %f, %f)%s", v.X, v.Y, v.Z, append)
}
func print_vec4(v Vec4, append *byte) {
	stdio.Printf("(%f, %f, %f, %f)%s", v.X, v.Y, v.Z, v.W, append)
}
func fread_vec2(f *stdio.File, v *vec2) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%f, %f)", &v.X, &v.Y))
	return int64(libc.BoolToInt(tmp == 2))
}
func fread_vec3(f *stdio.File, v *Vec3) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%f, %f, %f)", &v.X, &v.Y, &v.Z))
	return int64(libc.BoolToInt(tmp == 3))
}
func fread_vec4(f *stdio.File, v *Vec4) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%f, %f, %f, %f)", &v.X, &v.Y, &v.Z, &v.W))
	return int64(libc.BoolToInt(tmp == 4))
}

type dvec2 struct {
	X float64
	Y float64
}
type dvec3 struct {
	X float64
	Y float64
	Z float64
}
type dvec4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func fprint_dvec2(f *stdio.File, v dvec2, append *byte) {
	stdio.Fprintf(f, "(%f, %f)%s", v.X, v.Y, append)
}
func fprint_dvec3(f *stdio.File, v dvec3, append *byte) {
	stdio.Fprintf(f, "(%f, %f, %f)%s", v.X, v.Y, v.Z, append)
}
func fprint_dvec4(f *stdio.File, v dvec4, append *byte) {
	stdio.Fprintf(f, "(%f, %f, %f, %f)%s", v.X, v.Y, v.Z, v.W, append)
}
func fread_dvec2(f *stdio.File, v *dvec2) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%lf, %lf)", &v.X, &v.Y))
	return int64(libc.BoolToInt(tmp == 2))
}
func fread_dvec3(f *stdio.File, v *dvec3) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%lf, %lf, %lf)", &v.X, &v.Y, &v.Z))
	return int64(libc.BoolToInt(tmp == 3))
}
func fread_dvec4(f *stdio.File, v *dvec4) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%lf, %lf, %lf, %lf)", &v.X, &v.Y, &v.Z, &v.W))
	return int64(libc.BoolToInt(tmp == 4))
}

type ivec2 struct {
	X int64
	Y int64
}
type ivec3 struct {
	X int64
	Y int64
	Z int64
}
type ivec4 struct {
	X int64
	Y int64
	Z int64
	W int64
}

func make_ivec2(x int64, y int64) ivec2 {
	var v ivec2 = ivec2{X: x, Y: y}
	return v
}
func make_ivec3(x int64, y int64, z int64) ivec3 {
	var v ivec3 = ivec3{X: x, Y: y, Z: z}
	return v
}
func make_ivec4(x int64, y int64, z int64, w int64) ivec4 {
	var v ivec4 = ivec4{X: x, Y: y, Z: z, W: w}
	return v
}
func fprint_ivec2(f *stdio.File, v ivec2, append *byte) {
	stdio.Fprintf(f, "(%d, %d)%s", v.X, v.Y, append)
}
func fprint_ivec3(f *stdio.File, v ivec3, append *byte) {
	stdio.Fprintf(f, "(%d, %d, %d)%s", v.X, v.Y, v.Z, append)
}
func fprint_ivec4(f *stdio.File, v ivec4, append *byte) {
	stdio.Fprintf(f, "(%d, %d, %d, %d)%s", v.X, v.Y, v.Z, v.W, append)
}
func fread_ivec2(f *stdio.File, v *ivec2) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%d, %d)", &v.X, &v.Y))
	return int64(libc.BoolToInt(tmp == 2))
}
func fread_ivec3(f *stdio.File, v *ivec3) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%d, %d, %d)", &v.X, &v.Y, &v.Z))
	return int64(libc.BoolToInt(tmp == 3))
}
func fread_ivec4(f *stdio.File, v *ivec4) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%d, %d, %d, %d)", &v.X, &v.Y, &v.Z, &v.W))
	return int64(libc.BoolToInt(tmp == 4))
}

type uvec2 struct {
	X uint64
	Y uint64
}
type uvec3 struct {
	X uint64
	Y uint64
	Z uint64
}
type uvec4 struct {
	X uint64
	Y uint64
	Z uint64
	W uint64
}

func fprint_uvec2(f *stdio.File, v uvec2, append *byte) {
	stdio.Fprintf(f, "(%u, %u)%s", v.X, v.Y, append)
}
func fprint_uvec3(f *stdio.File, v uvec3, append *byte) {
	stdio.Fprintf(f, "(%u, %u, %u)%s", v.X, v.Y, v.Z, append)
}
func fprint_uvec4(f *stdio.File, v uvec4, append *byte) {
	stdio.Fprintf(f, "(%u, %u, %u, %u)%s", v.X, v.Y, v.Z, v.W, append)
}
func fread_uvec2(f *stdio.File, v *uvec2) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%u, %u)", &v.X, &v.Y))
	return int64(libc.BoolToInt(tmp == 2))
}
func fread_uvec3(f *stdio.File, v *uvec3) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%u, %u, %u)", &v.X, &v.Y, &v.Z))
	return int64(libc.BoolToInt(tmp == 3))
}
func fread_uvec4(f *stdio.File, v *uvec4) int64 {
	var tmp int64 = int64(stdio.Fscanf(f, " (%u, %u, %u, %u)", &v.X, &v.Y, &v.Z, &v.W))
	return int64(libc.BoolToInt(tmp == 4))
}
func length_vec2(a vec2) float32 {
	return float32(math.Sqrt(float64(a.X*a.X + a.Y*a.Y)))
}
func length_vec3(a Vec3) float32 {
	return float32(math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z)))
}
func norm_vec2(a vec2) vec2 {
	var (
		l float32 = length_vec2(a)
		c vec2    = vec2{X: a.X / l, Y: a.Y / l}
	)
	return c
}
func Norm_vec3(a Vec3) Vec3 {
	var (
		l float32 = length_vec3(a)
		c Vec3    = Vec3{X: a.X / l, Y: a.Y / l, Z: a.Z / l}
	)
	return c
}
func normalize_vec2(a *vec2) {
	var l float32 = length_vec2(*a)
	a.X /= l
	a.Y /= l
}
func normalize_vec3(a *Vec3) {
	var l float32 = length_vec3(*a)
	a.X /= l
	a.Y /= l
	a.Z /= l
}
func add_vec2s(a vec2, b vec2) vec2 {
	var c vec2 = vec2{X: a.X + b.X, Y: a.Y + b.Y}
	return c
}
func add_vec3s(a Vec3, b Vec3) Vec3 {
	var c Vec3 = Vec3{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
	return c
}
func add_vec4s(a Vec4, b Vec4) Vec4 {
	var c Vec4 = Vec4{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z, W: a.W + b.W}
	return c
}
func sub_vec2s(a vec2, b vec2) vec2 {
	var c vec2 = vec2{X: a.X - b.X, Y: a.Y - b.Y}
	return c
}
func sub_vec3s(a Vec3, b Vec3) Vec3 {
	var c Vec3 = Vec3{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
	return c
}
func sub_vec4s(a Vec4, b Vec4) Vec4 {
	var c Vec4 = Vec4{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z, W: a.W - b.W}
	return c
}
func mult_vec2s(a vec2, b vec2) vec2 {
	var c vec2 = vec2{X: a.X * b.X, Y: a.Y * b.Y}
	return c
}
func mult_vec3s(a Vec3, b Vec3) Vec3 {
	var c Vec3 = Vec3{X: a.X * b.X, Y: a.Y * b.Y, Z: a.Z * b.Z}
	return c
}
func mult_vec4s(a Vec4, b Vec4) Vec4 {
	var c Vec4 = Vec4{X: a.X * b.X, Y: a.Y * b.Y, Z: a.Z * b.Z, W: a.W * b.W}
	return c
}
func div_vec2s(a vec2, b vec2) vec2 {
	var c vec2 = vec2{X: a.X / b.X, Y: a.Y / b.Y}
	return c
}
func div_vec3s(a Vec3, b Vec3) Vec3 {
	var c Vec3 = Vec3{X: a.X / b.X, Y: a.Y / b.Y, Z: a.Z / b.Z}
	return c
}
func div_vec4s(a Vec4, b Vec4) Vec4 {
	var c Vec4 = Vec4{X: a.X / b.X, Y: a.Y / b.Y, Z: a.Z / b.Z, W: a.W / b.W}
	return c
}
func dot_vec2s(a vec2, b vec2) float32 {
	return a.X*b.X + a.Y*b.Y
}
func Dot_vec3s(a Vec3, b Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}
func dot_vec4s(a Vec4, b Vec4) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}
func scale_vec2(a vec2, s float32) vec2 {
	var b vec2 = vec2{X: a.X * s, Y: a.Y * s}
	return b
}
func Scale_vec3(a Vec3, s float32) Vec3 {
	var b Vec3 = Vec3{X: a.X * s, Y: a.Y * s, Z: a.Z * s}
	return b
}
func scale_vec4(a Vec4, s float32) Vec4 {
	var b Vec4 = Vec4{X: a.X * s, Y: a.Y * s, Z: a.Z * s, W: a.W * s}
	return b
}
func equal_vec2s(a vec2, b vec2) int64 {
	return int64(libc.BoolToInt(a.X == b.X && a.Y == b.Y))
}
func equal_vec3s(a Vec3, b Vec3) int64 {
	return int64(libc.BoolToInt(a.X == b.X && a.Y == b.Y && a.Z == b.Z))
}
func equal_vec4s(a Vec4, b Vec4) int64 {
	return int64(libc.BoolToInt(a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W))
}
func equal_epsilon_vec2s(a vec2, b vec2, epsilon float32) int64 {
	return int64(libc.BoolToInt(math.Abs(float64(a.X-b.X)) < float64(epsilon) && math.Abs(float64(a.Y-b.Y)) < float64(epsilon)))
}
func equal_epsilon_vec3s(a Vec3, b Vec3, epsilon float32) int64 {
	return int64(libc.BoolToInt(math.Abs(float64(a.X-b.X)) < float64(epsilon) && math.Abs(float64(a.Y-b.Y)) < float64(epsilon) && math.Abs(float64(a.Z-b.Z)) < float64(epsilon)))
}
func equal_epsilon_vec4s(a Vec4, b Vec4, epsilon float32) int64 {
	return int64(libc.BoolToInt(math.Abs(float64(a.X-b.X)) < float64(epsilon) && math.Abs(float64(a.Y-b.Y)) < float64(epsilon) && math.Abs(float64(a.Z-b.Z)) < float64(epsilon) && math.Abs(float64(a.W-b.W)) < float64(epsilon)))
}
func vec4_to_vec2(a Vec4) vec2 {
	var v vec2 = vec2{X: a.X, Y: a.Y}
	return v
}
func vec4_to_vec3(a Vec4) Vec3 {
	var v Vec3 = Vec3{X: a.X, Y: a.Y, Z: a.Z}
	return v
}
func vec4_to_vec2h(a Vec4) vec2 {
	var v vec2 = vec2{X: a.X / a.W, Y: a.Y / a.W}
	return v
}
func vec4_to_vec3h(a Vec4) Vec3 {
	var v Vec3 = Vec3{X: a.X / a.W, Y: a.Y / a.W, Z: a.Z / a.W}
	return v
}
func cross_product(u Vec3, v Vec3) Vec3 {
	var result Vec3
	result.X = u.Y*v.Z - v.Y*u.Z
	result.Y = -u.X*v.Z + v.X*u.Z
	result.Z = u.X*v.Y - v.X*u.Y
	return result
}
func angle_between_vec3(u Vec3, v Vec3) float32 {
	return float32(math.Acos(float64(Dot_vec3s(u, v))))
}

type mat2 [4]float32
type mat3 [9]float32
type Mat4 [16]float32

func x_mat2(m mat2) vec2 {
	return make_vec2(m[0], m[2])
}
func y_mat2(m mat2) vec2 {
	return make_vec2(m[1], m[3])
}
func c1_mat2(m mat2) vec2 {
	return make_vec2(m[0], m[1])
}
func c2_mat2(m mat2) vec2 {
	return make_vec2(m[2], m[3])
}
func setc1_mat2(m *mat2, v vec2) {
	m[0] = v.X
	m[1] = v.Y
}
func setc2_mat2(m *mat2, v vec2) {
	m[2] = v.X
	m[3] = v.Y
}
func setx_mat2(m *mat2, v vec2) {
	m[0] = v.X
	m[2] = v.Y
}
func sety_mat2(m *mat2, v vec2) {
	m[1] = v.X
	m[3] = v.Y
}
func x_mat3(m mat3) Vec3 {
	return make_vec3(m[0], m[3], m[6])
}
func y_mat3(m mat3) Vec3 {
	return make_vec3(m[1], m[4], m[7])
}
func z_mat3(m mat3) Vec3 {
	return make_vec3(m[2], m[5], m[8])
}
func c1_mat3(m mat3) Vec3 {
	return make_vec3(m[0], m[1], m[2])
}
func c2_mat3(m mat3) Vec3 {
	return make_vec3(m[3], m[4], m[5])
}
func c3_mat3(m mat3) Vec3 {
	return make_vec3(m[6], m[7], m[8])
}
func setc1_mat3(m *mat3, v Vec3) {
	m[0] = v.X
	m[1] = v.Y
	m[2] = v.Z
}
func setc2_mat3(m *mat3, v Vec3) {
	m[3] = v.X
	m[4] = v.Y
	m[5] = v.Z
}
func setc3_mat3(m *mat3, v Vec3) {
	m[6] = v.X
	m[7] = v.Y
	m[8] = v.Z
}
func setx_mat3(m *mat3, v Vec3) {
	m[0] = v.X
	m[3] = v.Y
	m[6] = v.Z
}
func sety_mat3(m *mat3, v Vec3) {
	m[1] = v.X
	m[4] = v.Y
	m[7] = v.Z
}
func setz_mat3(m *mat3, v Vec3) {
	m[2] = v.X
	m[5] = v.Y
	m[8] = v.Z
}
func c1_mat4(m Mat4) Vec4 {
	return make_vec4(m[0], m[1], m[2], m[3])
}
func c2_mat4(m Mat4) Vec4 {
	return make_vec4(m[4], m[5], m[6], m[7])
}
func c3_mat4(m Mat4) Vec4 {
	return make_vec4(m[8], m[9], m[10], m[11])
}
func c4_mat4(m Mat4) Vec4 {
	return make_vec4(m[12], m[13], m[14], m[15])
}
func x_mat4(m Mat4) Vec4 {
	return make_vec4(m[0], m[4], m[8], m[12])
}
func y_mat4(m Mat4) Vec4 {
	return make_vec4(m[1], m[5], m[9], m[13])
}
func z_mat4(m Mat4) Vec4 {
	return make_vec4(m[2], m[6], m[10], m[14])
}
func w_mat4(m Mat4) Vec4 {
	return make_vec4(m[3], m[7], m[11], m[15])
}
func setc1_mat4v3(m *Mat4, v Vec3) {
	m[0] = v.X
	m[1] = v.Y
	m[2] = v.Z
	m[3] = 0
}
func setc2_mat4v3(m *Mat4, v Vec3) {
	m[4] = v.X
	m[5] = v.Y
	m[6] = v.Z
	m[7] = 0
}
func setc3_mat4v3(m *Mat4, v Vec3) {
	m[8] = v.X
	m[9] = v.Y
	m[10] = v.Z
	m[11] = 0
}
func setc4_mat4v3(m *Mat4, v Vec3) {
	m[12] = v.X
	m[13] = v.Y
	m[14] = v.Z
	m[15] = 1
}
func setc1_mat4v4(m *Mat4, v Vec4) {
	m[0] = v.X
	m[1] = v.Y
	m[2] = v.Z
	m[3] = v.W
}
func setc2_mat4v4(m *Mat4, v Vec4) {
	m[4] = v.X
	m[5] = v.Y
	m[6] = v.Z
	m[7] = v.W
}
func setc3_mat4v4(m *Mat4, v Vec4) {
	m[8] = v.X
	m[9] = v.Y
	m[10] = v.Z
	m[11] = v.W
}
func setc4_mat4v4(m *Mat4, v Vec4) {
	m[12] = v.X
	m[13] = v.Y
	m[14] = v.Z
	m[15] = v.W
}
func setx_mat4v3(m *Mat4, v Vec3) {
	m[0] = v.X
	m[4] = v.Y
	m[8] = v.Z
	m[12] = 0
}
func sety_mat4v3(m *Mat4, v Vec3) {
	m[1] = v.X
	m[5] = v.Y
	m[9] = v.Z
	m[13] = 0
}
func setz_mat4v3(m *Mat4, v Vec3) {
	m[2] = v.X
	m[6] = v.Y
	m[10] = v.Z
	m[14] = 0
}
func setw_mat4v3(m *Mat4, v Vec3) {
	m[3] = v.X
	m[7] = v.Y
	m[11] = v.Z
	m[15] = 1
}
func setx_mat4v4(m *Mat4, v Vec4) {
	m[0] = v.X
	m[4] = v.Y
	m[8] = v.Z
	m[12] = v.W
}
func sety_mat4v4(m *Mat4, v Vec4) {
	m[1] = v.X
	m[5] = v.Y
	m[9] = v.Z
	m[13] = v.W
}
func setz_mat4v4(m *Mat4, v Vec4) {
	m[2] = v.X
	m[6] = v.Y
	m[10] = v.Z
	m[14] = v.W
}
func setw_mat4v4(m *Mat4, v Vec4) {
	m[3] = v.X
	m[7] = v.Y
	m[11] = v.Z
	m[15] = v.W
}
func fprint_mat2(f *stdio.File, m mat2, append *byte) {
	stdio.Fprintf(f, "[(%f, %f)\n (%f, %f)]%s", m[0], m[2], m[1], m[3], append)
}
func fprint_mat3(f *stdio.File, m mat3, append *byte) {
	stdio.Fprintf(f, "[(%f, %f, %f)\n (%f, %f, %f)\n (%f, %f, %f)]%s", m[0], m[3], m[6], m[1], m[4], m[7], m[2], m[5], m[8], append)
}
func fprint_mat4(f *stdio.File, m Mat4, append *byte) {
	stdio.Fprintf(f, "[(%f, %f, %f, %f)\n(%f, %f, %f, %f)\n(%f, %f, %f, %f)\n(%f, %f, %f, %f)]%s", m[0], m[4], m[8], m[12], m[1], m[5], m[9], m[13], m[2], m[6], m[10], m[14], m[3], m[7], m[11], m[15], append)
}
func print_mat2(m mat2, append *byte) {
	fprint_mat2(stdio.Stdout(), m, append)
}
func print_mat3(m mat3, append *byte) {
	fprint_mat3(stdio.Stdout(), m, append)
}
func print_mat4(m Mat4, append *byte) {
	fprint_mat4(stdio.Stdout(), m, append)
}
func mult_mat2_vec2(m mat2, v vec2) vec2 {
	var r vec2
	r.X = m[0]*v.X + m[2]*v.Y
	r.Y = m[1]*v.X + m[3]*v.Y
	return r
}
func mult_mat3_vec3(m mat3, v Vec3) Vec3 {
	var r Vec3
	r.X = m[0]*v.X + m[3]*v.Y + m[6]*v.Z
	r.Y = m[1]*v.X + m[4]*v.Y + m[7]*v.Z
	r.Z = m[2]*v.X + m[5]*v.Y + m[8]*v.Z
	return r
}
func Mult_mat4_vec4(m Mat4, v Vec4) Vec4 {
	var r Vec4
	r.X = m[0]*v.X + m[4]*v.Y + m[8]*v.Z + m[12]*v.W
	r.Y = m[1]*v.X + m[5]*v.Y + m[9]*v.Z + m[13]*v.W
	r.Z = m[2]*v.X + m[6]*v.Y + m[10]*v.Z + m[14]*v.W
	r.W = m[3]*v.X + m[7]*v.Y + m[11]*v.Z + m[15]*v.W
	return r
}
func scale_mat3(m *mat3, x float32, y float32, z float32) {
	m[0] = x
	m[3] = 0
	m[6] = 0
	m[1] = 0
	m[4] = y
	m[7] = 0
	m[2] = 0
	m[5] = 0
	m[8] = z
}
func scale_mat4(m *Mat4, x float32, y float32, z float32) {
	m[0] = x
	m[4] = 0
	m[8] = 0
	m[12] = 0
	m[1] = 0
	m[5] = y
	m[9] = 0
	m[13] = 0
	m[2] = 0
	m[6] = 0
	m[10] = z
	m[14] = 0
	m[3] = 0
	m[7] = 0
	m[11] = 0
	m[15] = 1
}
func Translation_mat4(m *Mat4, x float32, y float32, z float32) {
	m[0] = 1
	m[4] = 0
	m[8] = 0
	m[12] = x
	m[1] = 0
	m[5] = 1
	m[9] = 0
	m[13] = y
	m[2] = 0
	m[6] = 0
	m[10] = 1
	m[14] = z
	m[3] = 0
	m[7] = 0
	m[11] = 0
	m[15] = 1
}
func extract_rotation_mat4(dst *mat3, src Mat4, normalize int64) {
	var tmp Vec3
	if normalize != 0 {
		tmp.X = src[0*4+0]
		tmp.Y = src[0*4+1]
		tmp.Z = src[0*4+2]
		normalize_vec3(&tmp)
		dst[0*3+0] = tmp.X
		dst[0*3+1] = tmp.Y
		dst[0*3+2] = tmp.Z
		tmp.X = src[1*4+0]
		tmp.Y = src[1*4+1]
		tmp.Z = src[1*4+2]
		normalize_vec3(&tmp)
		dst[1*3+0] = tmp.X
		dst[1*3+1] = tmp.Y
		dst[1*3+2] = tmp.Z
		tmp.X = src[2*4+0]
		tmp.Y = src[2*4+1]
		tmp.Z = src[2*4+2]
		normalize_vec3(&tmp)
		dst[2*3+0] = tmp.X
		dst[2*3+1] = tmp.Y
		dst[2*3+2] = tmp.Z
	} else {
		dst[0*3+0] = src[0*4+0]
		dst[0*3+1] = src[0*4+1]
		dst[0*3+2] = src[0*4+2]
		dst[1*3+0] = src[1*4+0]
		dst[1*3+1] = src[1*4+1]
		dst[1*3+2] = src[1*4+2]
		dst[2*3+0] = src[2*4+0]
		dst[2*3+1] = src[2*4+1]
		dst[2*3+2] = src[2*4+2]
	}
}
func clamp_01(f float32) float32 {
	if float64(f) < 0.0 {
		return 0.0
	}
	if float64(f) > 1.0 {
		return 1.0
	}
	return f
}
func clamp(x float32, minVal float32, maxVal float32) float32 {
	if x < minVal {
		return minVal
	}
	if x > maxVal {
		return maxVal
	}
	return x
}
func clamp_vec2(x vec2, minVal float32, maxVal float32) vec2 {
	return make_vec2(clamp(x.X, minVal, maxVal), clamp(x.Y, minVal, maxVal))
}
func clamp_vec3(x Vec3, minVal float32, maxVal float32) Vec3 {
	return make_vec3(clamp(x.X, minVal, maxVal), clamp(x.Y, minVal, maxVal), clamp(x.Z, minVal, maxVal))
}
func clamp_vec4(x Vec4, minVal float32, maxVal float32) Vec4 {
	return make_vec4(clamp(x.X, minVal, maxVal), clamp(x.Y, minVal, maxVal), clamp(x.Z, minVal, maxVal), clamp(x.W, minVal, maxVal))
}
func distance_vec2(a vec2, b vec2) float32 {
	return length_vec2(sub_vec2s(a, b))
}
func distance_vec3(a Vec3, b Vec3) float32 {
	return length_vec3(sub_vec3s(a, b))
}
func reflect_vec3(i Vec3, n Vec3) Vec3 {
	return sub_vec3s(i, Scale_vec3(n, Dot_vec3s(i, n)*2))
}
func smoothstep(edge0 float32, edge1 float32, x float32) float32 {
	var t float32 = clamp_01((x - edge0) / (edge1 - edge0))
	return t * t * (3 - t*2)
}
func mix(x float32, y float32, a float32) float32 {
	return x*(1-a) + y*a
}
func mix_vec2s(x vec2, y vec2, a float32) vec2 {
	return add_vec2s(scale_vec2(x, 1-a), scale_vec2(y, a))
}
func mix_vec3s(x Vec3, y Vec3, a float32) Vec3 {
	return add_vec3s(Scale_vec3(x, 1-a), Scale_vec3(y, a))
}
func mix_vec4s(x Vec4, y Vec4, a float32) Vec4 {
	return add_vec4s(scale_vec4(x, 1-a), scale_vec4(y, a))
}
func fabsf_vec2(v vec2) vec2 {
	return make_vec2(math32.Abs(v.X), math32.Abs(v.Y))
}
func fabsf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Abs(v.X), math32.Abs(v.Y), math32.Abs(v.Z))
}
func fabsf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Abs(v.X), math32.Abs(v.Y), math32.Abs(v.Z), math32.Abs(v.W))
}
func floorf_vec2(v vec2) vec2 {
	return make_vec2(math32.Floor(v.X), math32.Floor(v.Y))
}
func floorf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Floor(v.X), math32.Floor(v.Y), math32.Floor(v.Z))
}
func floorf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Floor(v.X), math32.Floor(v.Y), math32.Floor(v.Z), math32.Floor(v.W))
}
func ceilf_vec2(v vec2) vec2 {
	return make_vec2(math32.Ceil(v.X), math32.Ceil(v.Y))
}
func ceilf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Ceil(v.X), math32.Ceil(v.Y), math32.Ceil(v.Z))
}
func ceilf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Ceil(v.X), math32.Ceil(v.Y), math32.Ceil(v.Z), math32.Ceil(v.W))
}
func sinf_vec2(v vec2) vec2 {
	return make_vec2(math32.Sin(v.X), math32.Sin(v.Y))
}
func sinf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Sin(v.X), math32.Sin(v.Y), math32.Sin(v.Z))
}
func sinf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Sin(v.X), math32.Sin(v.Y), math32.Sin(v.Z), math32.Sin(v.W))
}
func cosf_vec2(v vec2) vec2 {
	return make_vec2(math32.Cos(v.X), math32.Cos(v.Y))
}
func cosf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Cos(v.X), math32.Cos(v.Y), math32.Cos(v.Z))
}
func cosf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Cos(v.X), math32.Cos(v.Y), math32.Cos(v.Z), math32.Cos(v.W))
}
func tanf_vec2(v vec2) vec2 {
	return make_vec2(math32.Tan(v.X), math32.Tan(v.Y))
}
func tanf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Tan(v.X), math32.Tan(v.Y), math32.Tan(v.Z))
}
func tanf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Tan(v.X), math32.Tan(v.Y), math32.Tan(v.Z), math32.Tan(v.W))
}
func asinf_vec2(v vec2) vec2 {
	return make_vec2(math32.Asin(v.X), math32.Asin(v.Y))
}
func asinf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Asin(v.X), math32.Asin(v.Y), math32.Asin(v.Z))
}
func asinf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Asin(v.X), math32.Asin(v.Y), math32.Asin(v.Z), math32.Asin(v.W))
}
func acosf_vec2(v vec2) vec2 {
	return make_vec2(math32.Acos(v.X), math32.Acos(v.Y))
}
func acosf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Acos(v.X), math32.Acos(v.Y), math32.Acos(v.Z))
}
func acosf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Acos(v.X), math32.Acos(v.Y), math32.Acos(v.Z), math32.Acos(v.W))
}
func atanf_vec2(v vec2) vec2 {
	return make_vec2(math32.Atan(v.X), math32.Atan(v.Y))
}
func atanf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Atan(v.X), math32.Atan(v.Y), math32.Atan(v.Z))
}
func atanf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Atan(v.X), math32.Atan(v.Y), math32.Atan(v.Z), math32.Atan(v.W))
}
func sinhf_vec2(v vec2) vec2 {
	return make_vec2(math32.Sinh(v.X), math32.Sinh(v.Y))
}
func sinhf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Sinh(v.X), math32.Sinh(v.Y), math32.Sinh(v.Z))
}
func sinhf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Sinh(v.X), math32.Sinh(v.Y), math32.Sinh(v.Z), math32.Sinh(v.W))
}
func coshf_vec2(v vec2) vec2 {
	return make_vec2(math32.Cosh(v.X), math32.Cosh(v.Y))
}
func coshf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Cosh(v.X), math32.Cosh(v.Y), math32.Cosh(v.Z))
}
func coshf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Cosh(v.X), math32.Cosh(v.Y), math32.Cosh(v.Z), math32.Cosh(v.W))
}
func tanhf_vec2(v vec2) vec2 {
	return make_vec2(math32.Tanh(v.X), math32.Tanh(v.Y))
}
func tanhf_vec3(v Vec3) Vec3 {
	return make_vec3(math32.Tanh(v.X), math32.Tanh(v.Y), math32.Tanh(v.Z))
}
func tanhf_vec4(v Vec4) Vec4 {
	return make_vec4(math32.Tanh(v.X), math32.Tanh(v.Y), math32.Tanh(v.Z), math32.Tanh(v.W))
}
func radians(degrees float32) float32 {
	return float32(float64(degrees) * 0.017453292519943295)
}
func degrees(radians float32) float32 {
	return float32(float64(radians) * 57.29577951308232)
}
func fract(x float32) float32 {
	return float32(float64(x) - math.Floor(float64(x)))
}
func radians_vec2(v vec2) vec2 {
	return make_vec2(radians(v.X), radians(v.Y))
}
func radians_vec3(v Vec3) Vec3 {
	return make_vec3(radians(v.X), radians(v.Y), radians(v.Z))
}
func radians_vec4(v Vec4) Vec4 {
	return make_vec4(radians(v.X), radians(v.Y), radians(v.Z), radians(v.W))
}
func degrees_vec2(v vec2) vec2 {
	return make_vec2(degrees(v.X), degrees(v.Y))
}
func degrees_vec3(v Vec3) Vec3 {
	return make_vec3(degrees(v.X), degrees(v.Y), degrees(v.Z))
}
func degrees_vec4(v Vec4) Vec4 {
	return make_vec4(degrees(v.X), degrees(v.Y), degrees(v.Z), degrees(v.W))
}
func fract_vec2(v vec2) vec2 {
	return make_vec2(fract(v.X), fract(v.Y))
}
func fract_vec3(v Vec3) Vec3 {
	return make_vec3(fract(v.X), fract(v.Y), fract(v.Z))
}
func fract_vec4(v Vec4) Vec4 {
	return make_vec4(fract(v.X), fract(v.Y), fract(v.Z), fract(v.W))
}

type Color struct {
	R u8
	G u8
	B u8
	A u8
}

func make_Color(red u8, green u8, blue u8, alpha u8) Color {
	var c Color = Color{R: red, G: green, B: blue, A: alpha}
	return c
}
func print_Color(c Color, append *byte) {
	stdio.Printf("(%d, %d, %d, %d)%s", c.R, c.G, c.B, c.A, append)
}
func vec4_to_Color(v Vec4) Color {
	var c Color
	c.R = u8(v.X * math.MaxUint8)
	c.G = u8(v.Y * math.MaxUint8)
	c.B = u8(v.Z * math.MaxUint8)
	c.A = u8(v.W * math.MaxUint8)
	return c
}
func Color_to_vec4(c Color) Vec4 {
	var v Vec4 = Vec4{X: float32(float64(float32(c.R)) / 255.0), Y: float32(float64(float32(c.G)) / 255.0), Z: float32(float64(float32(c.B)) / 255.0), W: float32(float64(float32(c.A)) / 255.0)}
	return v
}

type Line struct {
	A float32
	B float32
	C float32
}

func make_Line(x1 float32, y1 float32, x2 float32, y2 float32) Line {
	var l Line
	l.A = y1 - y2
	l.B = x2 - x1
	l.C = x1*y2 - x2*y1
	return l
}
func line_func(line *Line, x float32, y float32) float32 {
	return line.A*x + line.B*y + line.C
}
func line_findy(line *Line, x float32) float32 {
	return -(line.A*x + line.C) / line.B
}
func line_findx(line *Line, y float32) float32 {
	return -(line.B*y + line.C) / line.A
}

type Plane struct {
	N Vec3
	D float32
}
type cvector_float struct {
	A        []float32
	Size     uint64
	Capacity uint64
}
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
	GL_NO_ERROR                      = 0
	GL_INVALID_ENUM                  = 1
	GL_INVALID_VALUE                 = 2
	GL_INVALID_OPERATION             = 3
	GL_INVALID_FRAMEBUFFER_OPERATION = 4
	GL_OUT_OF_MEMORY                 = 5
	GL_ARRAY_BUFFER                  = 6
	GL_COPY_READ_BUFFER              = 7
	GL_COPY_WRITE_BUFFER             = 8
	GL_ELEMENT_ARRAY_BUFFER          = 9
	GL_PIXEL_PACK_BUFFER             = 10
	GL_PIXEL_UNPACK_BUFFER           = 11
	GL_TEXTURE_BUFFER                = 12
	GL_TRANSFORM_FEEDBACK_BUFFER     = 13
	GL_UNIFORM_BUFFER                = 14
	GL_NUM_BUFFER_TYPES              = 15
	GL_STREAM_DRAW                   = 16
	GL_STREAM_READ                   = 17
	GL_STREAM_COPY                   = 18
	GL_STATIC_DRAW                   = 19
	GL_STATIC_READ                   = 20
	GL_STATIC_COPY                   = 21
	GL_DYNAMIC_DRAW                  = 22
	GL_DYNAMIC_READ                  = 23
	GL_DYNAMIC_COPY                  = 24
	GL_READ_ONLY                     = 25
	GL_WRITE_ONLY                    = 26
	GL_READ_WRITE                    = 27
	GL_POINT                         = 28
	GL_LINE                          = 29
	GL_FILL                          = 30
	GL_POINTS                        = 31
	GL_LINES                         = 32
	GL_LINE_STRIP                    = 33
	GL_LINE_LOOP                     = 34
	GL_TRIANGLES                     = 35
	GL_TRIANGLE_STRIP                = 36
	GL_TRIANGLE_FAN                  = 37
	GL_LINE_STRIP_AJACENCY           = 38
	GL_LINES_AJACENCY                = 39
	GL_TRIANGLES_AJACENCY            = 40
	GL_TRIANGLE_STRIP_AJACENCY       = 41
	GL_LESS                          = 42
	GL_LEQUAL                        = 43
	GL_GREATER                       = 44
	GL_GEQUAL                        = 45
	GL_EQUAL                         = 46
	GL_NOTEQUAL                      = 47
	GL_ALWAYS                        = 48
	GL_NEVER                         = 49
	GL_ZERO                          = 50
	GL_ONE                           = 51
	GL_SRC_COLOR                     = 52
	GL_ONE_MINUS_SRC_COLOR           = 53
	GL_DST_COLOR                     = 54
	GL_ONE_MINUS_DST_COLOR           = 55
	GL_SRC_ALPHA                     = 56
	GL_ONE_MINUS_SRC_ALPHA           = 57
	GL_DST_ALPHA                     = 58
	GL_ONE_MINUS_DST_ALPHA           = 59
	GL_CONSTANT_COLOR                = 60
	GL_ONE_MINUS_CONSTANT_COLOR      = 61
	GL_CONSTANT_ALPHA                = 62
	GL_ONE_MINUS_CONSTANT_ALPHA      = 63
	GL_SRC_ALPHA_SATURATE            = 64
	NUM_BLEND_FUNCS                  = 65
	GL_SRC1_COLOR                    = 66
	GL_ONE_MINUS_SRC1_COLOR          = 67
	GL_SRC1_ALPHA                    = 68
	GL_ONE_MINUS_SRC1_ALPHA          = 69
	GL_FUNC_ADD                      = 70
	GL_FUNC_SUBTRACT                 = 71
	GL_FUNC_REVERSE_SUBTRACT         = 72
	GL_MIN                           = 73
	GL_MAX                           = 74
	NUM_BLEND_EQUATIONS              = 75
	GL_TEXTURE_UNBOUND               = 76
	GL_TEXTURE_1D                    = 77
	GL_TEXTURE_2D                    = 78
	GL_TEXTURE_3D                    = 79
	GL_TEXTURE_1D_ARRAY              = 80
	GL_TEXTURE_2D_ARRAY              = 81
	GL_TEXTURE_RECTANGLE             = 82
	GL_TEXTURE_CUBE_MAP              = 83
	GL_NUM_TEXTURE_TYPES             = 84
	GL_TEXTURE_CUBE_MAP_POSITIVE_X   = 85
	GL_TEXTURE_CUBE_MAP_NEGATIVE_X   = 86
	GL_TEXTURE_CUBE_MAP_POSITIVE_Y   = 87
	GL_TEXTURE_CUBE_MAP_NEGATIVE_Y   = 88
	GL_TEXTURE_CUBE_MAP_POSITIVE_Z   = 89
	GL_TEXTURE_CUBE_MAP_NEGATIVE_Z   = 90
	GL_TEXTURE_BASE_LEVEL            = 91
	GL_TEXTURE_BORDER_COLOR          = 92
	GL_TEXTURE_COMPARE_FUNC          = 93
	GL_TEXTURE_COMPARE_MODE          = 94
	GL_TEXTURE_LOD_BIAS              = 95
	GL_TEXTURE_MIN_FILTER            = 96
	GL_TEXTURE_MAG_FILTER            = 97
	GL_TEXTURE_MIN_LOD               = 98
	GL_TEXTURE_MAX_LOD               = 99
	GL_TEXTURE_MAX_LEVEL             = 100
	GL_TEXTURE_SWIZZLE_R             = 101
	GL_TEXTURE_SWIZZLE_G             = 102
	GL_TEXTURE_SWIZZLE_B             = 103
	GL_TEXTURE_SWIZZLE_A             = 104
	GL_TEXTURE_SWIZZLE_RGBA          = 105
	GL_TEXTURE_WRAP_S                = 106
	GL_TEXTURE_WRAP_T                = 107
	GL_TEXTURE_WRAP_R                = 108
	GL_REPEAT                        = 109
	GL_CLAMP_TO_EDGE                 = 110
	GL_CLAMP_TO_BORDER               = 111
	GL_MIRRORED_REPEAT               = 112
	GL_NEAREST                       = 113
	GL_LINEAR                        = 114
	GL_NEAREST_MIPMAP_NEAREST        = 115
	GL_NEAREST_MIPMAP_LINEAR         = 116
	GL_LINEAR_MIPMAP_NEAREST         = 117
	GL_LINEAR_MIPMAP_LINEAR          = 118
	GL_RED                           = 119
	GL_RG                            = 120
	GL_RGB                           = 121
	GL_BGR                           = 122
	GL_RGBA                          = 123
	GL_BGRA                          = 124
	GL_COMPRESSED_RED                = 125
	GL_COMPRESSED_RG                 = 126
	GL_COMPRESSED_RGB                = math.MaxInt8
	GL_COMPRESSED_RGBA               = 128
	GL_UNPACK_ALIGNMENT              = 129
	GL_PACK_ALIGNMENT                = 130
	GL_TEXTURE0                      = 131
	GL_TEXTURE1                      = 132
	GL_TEXTURE2                      = 133
	GL_TEXTURE3                      = 134
	GL_TEXTURE4                      = 135
	GL_TEXTURE5                      = 136
	GL_TEXTURE6                      = 137
	GL_TEXTURE7                      = 138
	GL_CULL_FACE                     = 139
	GL_DEPTH_TEST                    = 140
	GL_DEPTH_CLAMP                   = 141
	GL_LINE_SMOOTH                   = 142
	GL_BLEND                         = 143
	GL_COLOR_LOGIC_OP                = 144
	GL_POLYGON_OFFSET_FILL           = 145
	GL_SCISSOR_TEST                  = 146
	GL_STENCIL_TEST                  = 147
	GL_FIRST_VERTEX_CONVENTION       = 148
	GL_LAST_VERTEX_CONVENTION        = 149
	GL_POINT_SPRITE_COORD_ORIGIN     = 150
	GL_UPPER_LEFT                    = 151
	GL_LOWER_LEFT                    = 152
	GL_FRONT                         = 153
	GL_BACK                          = 154
	GL_FRONT_AND_BACK                = 155
	GL_CCW                           = 156
	GL_CW                            = 157
	GL_CLEAR                         = 158
	GL_SET                           = 159
	GL_COPY                          = 160
	GL_COPY_INVERTED                 = 161
	GL_NOOP                          = 162
	GL_AND                           = 163
	GL_NAND                          = 164
	GL_OR                            = 165
	GL_NOR                           = 166
	GL_XOR                           = 167
	GL_EQUIV                         = 168
	GL_AND_REVERSE                   = 169
	GL_AND_INVERTED                  = 170
	GL_OR_REVERSE                    = 171
	GL_OR_INVERTED                   = 172
	GL_INVERT                        = 173
	GL_KEEP                          = 174
	GL_REPLACE                       = 175
	GL_INCR                          = 176
	GL_INCR_WRAP                     = 177
	GL_DECR                          = 178
	GL_DECR_WRAP                     = 179
	GL_UNSIGNED_BYTE                 = 180
	GL_BYTE                          = 181
	GL_BITMAP                        = 182
	GL_UNSIGNED_SHORT                = 183
	GL_SHORT                         = 184
	GL_UNSIGNED_INT                  = 185
	GL_INT                           = 186
	GL_FLOAT                         = 187
	GL_VENDOR                        = 188
	GL_RENDERER                      = 189
	GL_VERSION                       = 190
	GL_SHADING_LANGUAGE_VERSION      = 191
	GL_POLYGON_OFFSET_FACTOR         = 192
	GL_POLYGON_OFFSET_UNITS          = 193
	GL_POINT_SIZE                    = 194
	GL_DEPTH_CLEAR_VALUE             = 195
	GL_DEPTH_RANGE                   = 196
	GL_STENCIL_WRITE_MASK            = 197
	GL_STENCIL_REF                   = 198
	GL_STENCIL_VALUE_MASK            = 199
	GL_STENCIL_FUNC                  = 200
	GL_STENCIL_FAIL                  = 201
	GL_STENCIL_PASS_DEPTH_FAIL       = 202
	GL_STENCIL_PASS_DEPTH_PASS       = 203
	GL_STENCIL_BACK_WRITE_MASK       = 204
	GL_STENCIL_BACK_REF              = 205
	GL_STENCIL_BACK_VALUE_MASK       = 206
	GL_STENCIL_BACK_FUNC             = 207
	GL_STENCIL_BACK_FAIL             = 208
	GL_STENCIL_BACK_PASS_DEPTH_FAIL  = 209
	GL_STENCIL_BACK_PASS_DEPTH_PASS  = 210
	GL_LOGIC_OP_MODE                 = 211
	GL_BLEND_SRC_RGB                 = 212
	GL_BLEND_SRC_ALPHA               = 213
	GL_BLEND_DST_RGB                 = 214
	GL_BLEND_DST_ALPHA               = 215
	GL_BLEND_EQUATION_RGB            = 216
	GL_BLEND_EQUATION_ALPHA          = 217
	GL_CULL_FACE_MODE                = 218
	GL_FRONT_FACE                    = 219
	GL_DEPTH_FUNC                    = 220
	GL_PROVOKING_VERTEX              = 221
	GL_POLYGON_MODE                  = 222
	GL_COMPUTE_SHADER                = 223
	GL_VERTEX_SHADER                 = 224
	GL_TESS_CONTROL_SHADER           = 225
	GL_TESS_EVALUATION_SHADER        = 226
	GL_GEOMETRY_SHADER               = 227
	GL_FRAGMENT_SHADER               = 228
	GL_INFO_LOG_LENGTH               = 229
	GL_COMPILE_STATUS                = 230
	GL_LINK_STATUS                   = 231
	GL_COLOR_BUFFER_BIT              = 1 << 10
	GL_DEPTH_BUFFER_BIT              = 1 << 11
	GL_STENCIL_BUFFER_BIT            = 1 << 12
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
	Interpolation        [64]GLenum
	Fragdepth_or_discard GLboolean
	Deleted              GLboolean
}
type glBuffer struct {
	Size       GLsizei
	Type       GLenum
	Data       []u8
	Deleted    GLboolean
	User_owned GLboolean
}
type glVertex_Attrib struct {
	Size       GLint
	Type       GLenum
	Stride     GLsizei
	Offset     GLsizei
	Normalized GLboolean
	Buf        uint64
	Enabled    GLboolean
	Divisor    GLuint
}
type glVertex_Array struct {
	Vertex_attribs [GL_MAX_VERTEX_ATTRIBS]glVertex_Attrib
	Element_buffer GLuint
	Deleted        GLboolean
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
	Data       *u8
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
	Output_buf    cvector_float
}
type draw_triangle_func func(v0 *glVertex, v1 *glVertex, v2 *glVertex, provoke uint64)
type cvector_glVertex_Array struct {
	A        []glVertex_Array
	Size     uint64
	Capacity uint64
}
type cvector_glBuffer struct {
	A        []glBuffer
	Size     uint64
	Capacity uint64
}
type cvector_glTexture struct {
	A        *glTexture
	Size     uint64
	Capacity uint64
}
type cvector_glProgram struct {
	A        []glProgram
	Size     uint64
	Capacity uint64
}
type cvector_glVertex struct {
	A        []glVertex
	Size     uint64
	Capacity uint64
}
type GlContext struct {
	Vp_mat                 Mat4
	X_min                  int64
	Y_min                  int64
	X_max                  uint64
	Y_max                  uint64
	Vertex_arrays          cvector_glVertex_Array
	Buffers                cvector_glBuffer
	Textures               cvector_glTexture
	Programs               cvector_glProgram
	Cur_vertex_array       GLuint
	Bound_buffers          [9]GLuint
	Bound_textures         [7]GLuint
	Cur_texture2D          GLuint
	Cur_program            GLuint
	Error                  GLenum
	Uniform                unsafe.Pointer
	Vertex_attribs_vs      [16]Vec4
	Builtins               Shader_Builtins
	Vs_output              Vertex_Shader_output
	Fs_input               [64]float32
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
	Glverts                cvector_glVertex
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
	for {
		libc.MemSet(unsafe.Pointer(&mat[0]), 0, int(unsafe.Sizeof(float32(0))*16))
		mat[0] = func() float32 {
			p := &mat[5]
			mat[5] = func() float32 {
				p := &mat[10]
				mat[10] = func() float32 {
					p := &mat[15]
					mat[15] = 1
					return *p
				}()
				return *p
			}()
			return *p
		}()
		if true {
			break
		}
	}
	var f Vec3 = Norm_vec3(sub_vec3s(center, eye))
	var s Vec3 = Norm_vec3(cross_product(f, up))
	var u Vec3 = cross_product(s, f)
	setx_mat4v3(mat, s)
	sety_mat4v3(mat, u)
	setz_mat4v3(mat, negate_vec3(f))
	setc4_mat4v3(mat, make_vec3(-Dot_vec3s(s, eye), -Dot_vec3s(u, eye), Dot_vec3s(f, eye)))
}

var CVEC_glVertex_Array_SZ uint64 = 50

func cvec_glVertex_Array_heap(size uint64, capacity uint64) *cvector_glVertex_Array {
	var vec *cvector_glVertex_Array
	if (func() *cvector_glVertex_Array {
		vec = new(cvector_glVertex_Array)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_glVertex_Array_SZ
	}
	if vec.A = make([]glVertex_Array, vec.Capacity); /*(*glVertex_Array)(libc.Malloc(int(vec.Capacity * uint64(unsafe.Sizeof(glVertex_Array{})))))*/ vec.A == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	return vec
}
func cvec_init_glVertex_Array_heap(vals *glVertex_Array, num uint64) *cvector_glVertex_Array {
	var vec *cvector_glVertex_Array
	if (func() *cvector_glVertex_Array {
		vec = new(cvector_glVertex_Array)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Capacity = num + CVEC_glVertex_Array_SZ
	vec.Size = num
	if vec.A = make([]glVertex_Array, vec.Capacity); /*(*glVertex_Array)(libc.Malloc(int(vec.Capacity * uint64(unsafe.Sizeof(glVertex_Array{})))))*/ vec.A == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	libc.MemMove(unsafe.Pointer(&vec.A[0]), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(glVertex_Array{}))))
	return vec
}
func cvec_glVertex_Array(vec *cvector_glVertex_Array, size uint64, capacity uint64) int64 {
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_glVertex_Array_SZ
	}
	if vec.A = make([]glVertex_Array, vec.Capacity); /*(*glVertex_Array)(libc.Malloc(int(vec.Capacity * uint64(unsafe.Sizeof(glVertex_Array{})))))*/ vec.A == nil {
		libc.Assert(vec.A != nil)
		vec.Capacity = 0
		vec.Size = 0
		return 0
	}
	return 1
}
func cvec_init_glVertex_Array(vec *cvector_glVertex_Array, vals *glVertex_Array, num uint64) int64 {
	vec.Capacity = num + CVEC_glVertex_Array_SZ
	vec.Size = num
	if vec.A = make([]glVertex_Array, vec.Capacity); /*(*glVertex_Array)(libc.Malloc(int(vec.Capacity * uint64(unsafe.Sizeof(glVertex_Array{})))))*/ vec.A == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	libc.MemMove(unsafe.Pointer(&vec.A[0]), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(glVertex_Array{}))))
	return 1
}
func cvec_copyc_glVertex_Array(dest unsafe.Pointer, src unsafe.Pointer) int64 {
	var (
		vec1 *cvector_glVertex_Array = (*cvector_glVertex_Array)(dest)
		vec2 *cvector_glVertex_Array = (*cvector_glVertex_Array)(src)
	)
	vec1.A = nil
	vec1.Size = 0
	vec1.Capacity = 0
	return cvec_copy_glVertex_Array(vec1, vec2)
}
func cvec_copy_glVertex_Array(dest *cvector_glVertex_Array, src *cvector_glVertex_Array) int64 {
	var tmp []glVertex_Array = nil
	if tmp = unsafe.Slice((*glVertex_Array)(libc.Realloc(unsafe.Pointer(&dest.A[0]), int(src.Capacity*uint64(unsafe.Sizeof(glVertex_Array{}))))), src.Capacity); tmp == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	dest.A = tmp
	libc.MemMove(unsafe.Pointer(&dest.A[0]), unsafe.Pointer(&src.A[0]), int(src.Size*uint64(unsafe.Sizeof(glVertex_Array{}))))
	dest.Size = src.Size
	dest.Capacity = src.Capacity
	return 1
}
func cvec_push_glVertex_Array(vec *cvector_glVertex_Array, a glVertex_Array) int64 {
	var (
		tmp    []glVertex_Array
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		*(*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(func() uint64 {
			p := &vec.Size
			x := *p
			*p++
			return x
		}()))) = a
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if tmp = unsafe.Slice((*glVertex_Array)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glVertex_Array{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		*(*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(func() uint64 {
			p := &vec.Size
			x := *p
			*p++
			return x
		}()))) = a
		vec.Capacity = tmp_sz
	}
	return 1
}
func cvec_pop_glVertex_Array(vec *cvector_glVertex_Array) glVertex_Array {
	return *(*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(func() uint64 {
		p := &vec.Size
		*p--
		return *p
	}())))
}
func cvec_back_glVertex_Array(vec *cvector_glVertex_Array) *glVertex_Array {
	return (*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(vec.Size-1)))
}
func cvec_extend_glVertex_Array(vec *cvector_glVertex_Array, num uint64) int64 {
	var (
		tmp    []glVertex_Array
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_glVertex_Array_SZ
		if tmp = unsafe.Slice((*glVertex_Array)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glVertex_Array{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	vec.Size += num
	return 1
}
func cvec_insert_glVertex_Array(vec *cvector_glVertex_Array, i uint64, a glVertex_Array) int64 {
	var (
		tmp    []glVertex_Array
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		libc.MemMove(unsafe.Pointer((*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i+1)))), unsafe.Pointer((*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glVertex_Array{}))))
		*(*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i))) = a
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if tmp = unsafe.Slice((*glVertex_Array)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glVertex_Array{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		libc.MemMove(unsafe.Pointer((*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i+1)))), unsafe.Pointer((*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glVertex_Array{}))))
		*(*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i))) = a
		vec.Capacity = tmp_sz
	}
	vec.Size++
	return 1
}
func cvec_insert_array_glVertex_Array(vec *cvector_glVertex_Array, i uint64, a *glVertex_Array, num uint64) int64 {
	var (
		tmp    []glVertex_Array
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_glVertex_Array_SZ
		if tmp = unsafe.Slice((*glVertex_Array)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glVertex_Array{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	libc.MemMove(unsafe.Pointer((*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i+num)))), unsafe.Pointer((*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glVertex_Array{}))))
	libc.MemMove(unsafe.Pointer((*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i)))), unsafe.Pointer(a), int(num*uint64(unsafe.Sizeof(glVertex_Array{}))))
	vec.Size += num
	return 1
}
func cvec_replace_glVertex_Array(vec *cvector_glVertex_Array, i uint64, a glVertex_Array) glVertex_Array {
	var tmp glVertex_Array = *(*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i)))
	*(*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i))) = a
	return tmp
}
func cvec_erase_glVertex_Array(vec *cvector_glVertex_Array, start uint64, end uint64) {
	var d uint64 = end - start + 1
	libc.MemMove(unsafe.Pointer((*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(start)))), unsafe.Pointer((*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(end+1)))), int((vec.Size-1-end)*uint64(unsafe.Sizeof(glVertex_Array{}))))
	vec.Size -= d
}
func cvec_reserve_glVertex_Array(vec *cvector_glVertex_Array, size uint64) int64 {
	var tmp []glVertex_Array
	if vec.Capacity < size {
		if tmp = unsafe.Slice((*glVertex_Array)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int((size+CVEC_glVertex_Array_SZ)*uint64(unsafe.Sizeof(glVertex_Array{}))))), size+CVEC_glVertex_Array_SZ); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = size + CVEC_glVertex_Array_SZ
	}
	return 1
}
func cvec_set_cap_glVertex_Array(vec *cvector_glVertex_Array, size uint64) int64 {
	var tmp []glVertex_Array
	if size < vec.Size {
		vec.Size = size
	}
	if tmp = unsafe.Slice((*glVertex_Array)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(size*uint64(unsafe.Sizeof(glVertex_Array{}))))), size); tmp == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	vec.A = tmp
	vec.Capacity = size
	return 1
}
func cvec_set_val_sz_glVertex_Array(vec *cvector_glVertex_Array, val glVertex_Array) {
	var i uint64
	for i = 0; i < vec.Size; i++ {
		*(*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i))) = val
	}
}
func cvec_set_val_cap_glVertex_Array(vec *cvector_glVertex_Array, val glVertex_Array) {
	var i uint64
	for i = 0; i < vec.Capacity; i++ {
		*(*glVertex_Array)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex_Array{})*uintptr(i))) = val
	}
}
func cvec_clear_glVertex_Array(vec *cvector_glVertex_Array) {
	vec.Size = 0
}
func cvec_free_glVertex_Array_heap(vec unsafe.Pointer) {
	var tmp *cvector_glVertex_Array = (*cvector_glVertex_Array)(vec)
	if tmp == nil {
		return
	}
	libc.Free(unsafe.Pointer(&tmp.A[0]))
	libc.Free(unsafe.Pointer(tmp))
}
func cvec_free_glVertex_Array(vec unsafe.Pointer) {
	var tmp *cvector_glVertex_Array = (*cvector_glVertex_Array)(vec)
	libc.Free(unsafe.Pointer(&tmp.A[0]))
	tmp.Size = 0
	tmp.Capacity = 0
}

var CVEC_glBuffer_SZ uint64 = 50

func cvec_glBuffer_heap(size uint64, capacity uint64) *cvector_glBuffer {
	var vec *cvector_glBuffer
	if (func() *cvector_glBuffer {
		vec = new(cvector_glBuffer)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_glBuffer_SZ
	}
	if vec.A = unsafe.Slice((*glBuffer)(libc.Malloc(int(vec.Capacity*uint64(unsafe.Sizeof(glBuffer{}))))), vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	return vec
}
func cvec_init_glBuffer_heap(vals *glBuffer, num uint64) *cvector_glBuffer {
	var vec *cvector_glBuffer
	if (func() *cvector_glBuffer {
		vec = new(cvector_glBuffer)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Capacity = num + CVEC_glBuffer_SZ
	vec.Size = num
	if vec.A = unsafe.Slice((*glBuffer)(libc.Malloc(int(vec.Capacity*uint64(unsafe.Sizeof(glBuffer{}))))), vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	libc.MemMove(unsafe.Pointer(&vec.A[0]), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(glBuffer{}))))
	return vec
}
func cvec_glBuffer(vec *cvector_glBuffer, size uint64, capacity uint64) int64 {
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_glBuffer_SZ
	}
	if vec.A = unsafe.Slice((*glBuffer)(libc.Malloc(int(vec.Capacity*uint64(unsafe.Sizeof(glBuffer{}))))), vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	return 1
}
func cvec_init_glBuffer(vec *cvector_glBuffer, vals *glBuffer, num uint64) int64 {
	vec.Capacity = num + CVEC_glBuffer_SZ
	vec.Size = num
	if vec.A = unsafe.Slice((*glBuffer)(libc.Malloc(int(vec.Capacity*uint64(unsafe.Sizeof(glBuffer{}))))), vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	libc.MemMove(unsafe.Pointer(&vec.A[0]), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(glBuffer{}))))
	return 1
}
func cvec_copyc_glBuffer(dest unsafe.Pointer, src unsafe.Pointer) int64 {
	var (
		vec1 *cvector_glBuffer = (*cvector_glBuffer)(dest)
		vec2 *cvector_glBuffer = (*cvector_glBuffer)(src)
	)
	vec1.A = nil
	vec1.Size = 0
	vec1.Capacity = 0
	return cvec_copy_glBuffer(vec1, vec2)
}
func cvec_copy_glBuffer(dest *cvector_glBuffer, src *cvector_glBuffer) int64 {
	var tmp []glBuffer = nil
	if tmp = unsafe.Slice((*glBuffer)(libc.Realloc(unsafe.Pointer(&dest.A[0]), int(src.Capacity*uint64(unsafe.Sizeof(glBuffer{}))))), src.Capacity); tmp == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	dest.A = tmp
	libc.MemMove(unsafe.Pointer(&dest.A[0]), unsafe.Pointer(&src.A[0]), int(src.Size*uint64(unsafe.Sizeof(glBuffer{}))))
	dest.Size = src.Size
	dest.Capacity = src.Capacity
	return 1
}
func cvec_push_glBuffer(vec *cvector_glBuffer, a glBuffer) int64 {
	var (
		tmp    []glBuffer
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		vec.A[vec.Size] = a
		vec.Size++
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if tmp = unsafe.Slice((*glBuffer)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glBuffer{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.A[vec.Size] = a
		vec.Size++
		vec.Capacity = tmp_sz
	}
	return 1
}
func cvec_pop_glBuffer(vec *cvector_glBuffer) glBuffer {
	return *(*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(func() uint64 {
		p := &vec.Size
		*p--
		return *p
	}())))
}
func cvec_back_glBuffer(vec *cvector_glBuffer) *glBuffer {
	return (*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(vec.Size-1)))
}
func cvec_extend_glBuffer(vec *cvector_glBuffer, num uint64) int64 {
	var (
		tmp    []glBuffer
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_glBuffer_SZ
		if tmp = unsafe.Slice((*glBuffer)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glBuffer{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	vec.Size += num
	return 1
}
func cvec_insert_glBuffer(vec *cvector_glBuffer, i uint64, a glBuffer) int64 {
	var (
		tmp    []glBuffer
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		libc.MemMove(unsafe.Pointer((*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i+1)))), unsafe.Pointer((*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glBuffer{}))))
		*(*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i))) = a
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if tmp = unsafe.Slice((*glBuffer)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glBuffer{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		libc.MemMove(unsafe.Pointer((*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i+1)))), unsafe.Pointer((*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glBuffer{}))))
		*(*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i))) = a
		vec.Capacity = tmp_sz
	}
	vec.Size++
	return 1
}
func cvec_insert_array_glBuffer(vec *cvector_glBuffer, i uint64, a *glBuffer, num uint64) int64 {
	var (
		tmp    []glBuffer
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_glBuffer_SZ
		if tmp = unsafe.Slice((*glBuffer)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glBuffer{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	libc.MemMove(unsafe.Pointer((*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i+num)))), unsafe.Pointer((*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glBuffer{}))))
	libc.MemMove(unsafe.Pointer((*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i)))), unsafe.Pointer(a), int(num*uint64(unsafe.Sizeof(glBuffer{}))))
	vec.Size += num
	return 1
}
func cvec_replace_glBuffer(vec *cvector_glBuffer, i uint64, a glBuffer) glBuffer {
	var tmp glBuffer = *(*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i)))
	*(*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i))) = a
	return tmp
}
func cvec_erase_glBuffer(vec *cvector_glBuffer, start uint64, end uint64) {
	var d uint64 = end - start + 1
	libc.MemMove(unsafe.Pointer((*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(start)))), unsafe.Pointer((*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(end+1)))), int((vec.Size-1-end)*uint64(unsafe.Sizeof(glBuffer{}))))
	vec.Size -= d
}
func cvec_reserve_glBuffer(vec *cvector_glBuffer, size uint64) int64 {
	var tmp []glBuffer
	if vec.Capacity < size {
		if tmp = unsafe.Slice((*glBuffer)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int((size+CVEC_glBuffer_SZ)*uint64(unsafe.Sizeof(glBuffer{}))))), size+CVEC_glBuffer_SZ); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = size + CVEC_glBuffer_SZ
	}
	return 1
}
func cvec_set_cap_glBuffer(vec *cvector_glBuffer, size uint64) int64 {
	var tmp []glBuffer
	if size < vec.Size {
		vec.Size = size
	}
	if tmp = unsafe.Slice((*glBuffer)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(size*uint64(unsafe.Sizeof(glBuffer{}))))), size); tmp == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	vec.A = tmp
	vec.Capacity = size
	return 1
}
func cvec_set_val_sz_glBuffer(vec *cvector_glBuffer, val glBuffer) {
	var i uint64
	for i = 0; i < vec.Size; i++ {
		*(*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i))) = val
	}
}
func cvec_set_val_cap_glBuffer(vec *cvector_glBuffer, val glBuffer) {
	var i uint64
	for i = 0; i < vec.Capacity; i++ {
		*(*glBuffer)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glBuffer{})*uintptr(i))) = val
	}
}
func cvec_clear_glBuffer(vec *cvector_glBuffer) {
	vec.Size = 0
}
func cvec_free_glBuffer_heap(vec unsafe.Pointer) {
	var tmp *cvector_glBuffer = (*cvector_glBuffer)(vec)
	if tmp == nil {
		return
	}
	libc.Free(unsafe.Pointer(&tmp.A[0]))
	libc.Free(unsafe.Pointer(tmp))
}
func cvec_free_glBuffer(vec unsafe.Pointer) {
	var tmp *cvector_glBuffer = (*cvector_glBuffer)(vec)
	tmp.A = nil
	tmp.Size = 0
	tmp.Capacity = 0
}

var CVEC_glTexture_SZ uint64 = 50

func cvec_glTexture_heap(size uint64, capacity uint64) *cvector_glTexture {
	var vec *cvector_glTexture
	if (func() *cvector_glTexture {
		vec = new(cvector_glTexture)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_glTexture_SZ
	}
	if (func() *glTexture {
		p := &vec.A
		vec.A = (*glTexture)(libc.Malloc(int(vec.Capacity * uint64(unsafe.Sizeof(glTexture{})))))
		return *p
	}()) == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	return vec
}
func cvec_init_glTexture_heap(vals *glTexture, num uint64) *cvector_glTexture {
	var vec *cvector_glTexture
	if (func() *cvector_glTexture {
		vec = new(cvector_glTexture)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Capacity = num + CVEC_glTexture_SZ
	vec.Size = num
	if (func() *glTexture {
		p := &vec.A
		vec.A = (*glTexture)(libc.Malloc(int(vec.Capacity * uint64(unsafe.Sizeof(glTexture{})))))
		return *p
	}()) == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	libc.MemMove(unsafe.Pointer(vec.A), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(glTexture{}))))
	return vec
}
func cvec_glTexture(vec *cvector_glTexture, size uint64, capacity uint64) int64 {
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_glTexture_SZ
	}
	if (func() *glTexture {
		p := &vec.A
		vec.A = (*glTexture)(libc.Malloc(int(vec.Capacity * uint64(unsafe.Sizeof(glTexture{})))))
		return *p
	}()) == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	return 1
}
func cvec_init_glTexture(vec *cvector_glTexture, vals *glTexture, num uint64) int64 {
	vec.Capacity = num + CVEC_glTexture_SZ
	vec.Size = num
	if (func() *glTexture {
		p := &vec.A
		vec.A = (*glTexture)(libc.Malloc(int(vec.Capacity * uint64(unsafe.Sizeof(glTexture{})))))
		return *p
	}()) == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	libc.MemMove(unsafe.Pointer(vec.A), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(glTexture{}))))
	return 1
}
func cvec_copyc_glTexture(dest unsafe.Pointer, src unsafe.Pointer) int64 {
	var (
		vec1 *cvector_glTexture = (*cvector_glTexture)(dest)
		vec2 *cvector_glTexture = (*cvector_glTexture)(src)
	)
	vec1.A = nil
	vec1.Size = 0
	vec1.Capacity = 0
	return cvec_copy_glTexture(vec1, vec2)
}
func cvec_copy_glTexture(dest *cvector_glTexture, src *cvector_glTexture) int64 {
	var tmp *glTexture = nil
	if (func() *glTexture {
		tmp = (*glTexture)(libc.Realloc(unsafe.Pointer(dest.A), int(src.Capacity*uint64(unsafe.Sizeof(glTexture{})))))
		return tmp
	}()) == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	dest.A = tmp
	libc.MemMove(unsafe.Pointer(dest.A), unsafe.Pointer(src.A), int(src.Size*uint64(unsafe.Sizeof(glTexture{}))))
	dest.Size = src.Size
	dest.Capacity = src.Capacity
	return 1
}
func cvec_push_glTexture(vec *cvector_glTexture, a glTexture) int64 {
	var (
		tmp    *glTexture
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		*(*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(func() uint64 {
			p := &vec.Size
			x := *p
			*p++
			return x
		}()))) = a
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if (func() *glTexture {
			tmp = (*glTexture)(libc.Realloc(unsafe.Pointer(vec.A), int(tmp_sz*uint64(unsafe.Sizeof(glTexture{})))))
			return tmp
		}()) == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		*(*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(func() uint64 {
			p := &vec.Size
			x := *p
			*p++
			return x
		}()))) = a
		vec.Capacity = tmp_sz
	}
	return 1
}
func cvec_pop_glTexture(vec *cvector_glTexture) glTexture {
	return *(*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(func() uint64 {
		p := &vec.Size
		*p--
		return *p
	}())))
}
func cvec_back_glTexture(vec *cvector_glTexture) *glTexture {
	return (*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(vec.Size-1)))
}
func cvec_extend_glTexture(vec *cvector_glTexture, num uint64) int64 {
	var (
		tmp    *glTexture
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_glTexture_SZ
		if (func() *glTexture {
			tmp = (*glTexture)(libc.Realloc(unsafe.Pointer(vec.A), int(tmp_sz*uint64(unsafe.Sizeof(glTexture{})))))
			return tmp
		}()) == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	vec.Size += num
	return 1
}
func cvec_insert_glTexture(vec *cvector_glTexture, i uint64, a glTexture) int64 {
	var (
		tmp    *glTexture
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		libc.MemMove(unsafe.Pointer((*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i+1)))), unsafe.Pointer((*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glTexture{}))))
		*(*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i))) = a
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if (func() *glTexture {
			tmp = (*glTexture)(libc.Realloc(unsafe.Pointer(vec.A), int(tmp_sz*uint64(unsafe.Sizeof(glTexture{})))))
			return tmp
		}()) == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		libc.MemMove(unsafe.Pointer((*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i+1)))), unsafe.Pointer((*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glTexture{}))))
		*(*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i))) = a
		vec.Capacity = tmp_sz
	}
	vec.Size++
	return 1
}
func cvec_insert_array_glTexture(vec *cvector_glTexture, i uint64, a *glTexture, num uint64) int64 {
	var (
		tmp    *glTexture
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_glTexture_SZ
		if (func() *glTexture {
			tmp = (*glTexture)(libc.Realloc(unsafe.Pointer(vec.A), int(tmp_sz*uint64(unsafe.Sizeof(glTexture{})))))
			return tmp
		}()) == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	libc.MemMove(unsafe.Pointer((*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i+num)))), unsafe.Pointer((*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glTexture{}))))
	libc.MemMove(unsafe.Pointer((*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i)))), unsafe.Pointer(a), int(num*uint64(unsafe.Sizeof(glTexture{}))))
	vec.Size += num
	return 1
}
func cvec_replace_glTexture(vec *cvector_glTexture, i uint64, a glTexture) glTexture {
	var tmp glTexture = *(*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i)))
	*(*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i))) = a
	return tmp
}
func cvec_erase_glTexture(vec *cvector_glTexture, start uint64, end uint64) {
	var d uint64 = end - start + 1
	libc.MemMove(unsafe.Pointer((*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(start)))), unsafe.Pointer((*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(end+1)))), int((vec.Size-1-end)*uint64(unsafe.Sizeof(glTexture{}))))
	vec.Size -= d
}
func cvec_reserve_glTexture(vec *cvector_glTexture, size uint64) int64 {
	var tmp *glTexture
	if vec.Capacity < size {
		if (func() *glTexture {
			tmp = (*glTexture)(libc.Realloc(unsafe.Pointer(vec.A), int((size+CVEC_glTexture_SZ)*uint64(unsafe.Sizeof(glTexture{})))))
			return tmp
		}()) == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = size + CVEC_glTexture_SZ
	}
	return 1
}
func cvec_set_cap_glTexture(vec *cvector_glTexture, size uint64) int64 {
	var tmp *glTexture
	if size < vec.Size {
		vec.Size = size
	}
	if (func() *glTexture {
		tmp = (*glTexture)(libc.Realloc(unsafe.Pointer(vec.A), int(size*uint64(unsafe.Sizeof(glTexture{})))))
		return tmp
	}()) == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	vec.A = tmp
	vec.Capacity = size
	return 1
}
func cvec_set_val_sz_glTexture(vec *cvector_glTexture, val glTexture) {
	var i uint64
	for i = 0; i < vec.Size; i++ {
		*(*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i))) = val
	}
}
func cvec_set_val_cap_glTexture(vec *cvector_glTexture, val glTexture) {
	var i uint64
	for i = 0; i < vec.Capacity; i++ {
		*(*glTexture)(unsafe.Add(unsafe.Pointer(vec.A), unsafe.Sizeof(glTexture{})*uintptr(i))) = val
	}
}
func cvec_clear_glTexture(vec *cvector_glTexture) {
	vec.Size = 0
}
func cvec_free_glTexture_heap(vec unsafe.Pointer) {
	var tmp *cvector_glTexture = (*cvector_glTexture)(vec)
	if tmp == nil {
		return
	}
	libc.Free(unsafe.Pointer(tmp.A))
	libc.Free(unsafe.Pointer(tmp))
}
func cvec_free_glTexture(vec unsafe.Pointer) {
	var tmp *cvector_glTexture = (*cvector_glTexture)(vec)
	libc.Free(unsafe.Pointer(tmp.A))
	tmp.Size = 0
	tmp.Capacity = 0
}

var CVEC_glProgram_SZ uint64 = 50

func cvec_glProgram_heap(size uint64, capacity uint64) *cvector_glProgram {
	var vec *cvector_glProgram
	if vec = new(cvector_glProgram); vec == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_glProgram_SZ
	}
	if vec.A = make([]glProgram, vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	return vec
}
func cvec_init_glProgram_heap(vals *glProgram, num uint64) *cvector_glProgram {
	var vec *cvector_glProgram
	if (func() *cvector_glProgram {
		vec = new(cvector_glProgram)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Capacity = num + CVEC_glProgram_SZ
	vec.Size = num
	if vec.A = make([]glProgram, vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	libc.MemMove(unsafe.Pointer(&vec.A[0]), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(glProgram{}))))
	return vec
}
func cvec_glProgram(vec *cvector_glProgram, size uint64, capacity uint64) int64 {
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_glProgram_SZ
	}
	if vec.A = make([]glProgram, vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	return 1
}
func cvec_init_glProgram(vec *cvector_glProgram, vals *glProgram, num uint64) int64 {
	vec.Capacity = num + CVEC_glProgram_SZ
	vec.Size = num
	if vec.A = make([]glProgram, vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	libc.MemMove(unsafe.Pointer(&vec.A[0]), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(glProgram{}))))
	return 1
}
func cvec_copyc_glProgram(dest unsafe.Pointer, src unsafe.Pointer) int64 {
	var (
		vec1 *cvector_glProgram = (*cvector_glProgram)(dest)
		vec2 *cvector_glProgram = (*cvector_glProgram)(src)
	)
	vec1.A = nil
	vec1.Size = 0
	vec1.Capacity = 0
	return cvec_copy_glProgram(vec1, vec2)
}
func cvec_copy_glProgram(dest *cvector_glProgram, src *cvector_glProgram) int64 {
	var tmp []glProgram = nil
	if tmp = unsafe.Slice((*glProgram)(libc.Realloc(unsafe.Pointer(&dest.A[0]), int(src.Capacity*uint64(unsafe.Sizeof(glProgram{}))))), src.Capacity); tmp == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	dest.A = tmp
	libc.MemMove(unsafe.Pointer(&dest.A[0]), unsafe.Pointer(&src.A[0]), int(src.Size*uint64(unsafe.Sizeof(glProgram{}))))
	dest.Size = src.Size
	dest.Capacity = src.Capacity
	return 1
}
func cvec_push_glProgram(vec *cvector_glProgram, a glProgram) int64 {
	var (
		tmp    []glProgram
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		vec.A[vec.Size] = a
		vec.Size++
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if tmp = unsafe.Slice((*glProgram)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glProgram{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.A[vec.Size] = a
		vec.Size++
		vec.Capacity = tmp_sz
	}
	return 1
}
func cvec_pop_glProgram(vec *cvector_glProgram) glProgram {
	return *(*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(func() uint64 {
		p := &vec.Size
		*p--
		return *p
	}())))
}
func cvec_back_glProgram(vec *cvector_glProgram) *glProgram {
	return (*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(vec.Size-1)))
}
func cvec_extend_glProgram(vec *cvector_glProgram, num uint64) int64 {
	var (
		tmp    []glProgram
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_glProgram_SZ
		if tmp = unsafe.Slice((*glProgram)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glProgram{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	vec.Size += num
	return 1
}
func cvec_insert_glProgram(vec *cvector_glProgram, i uint64, a glProgram) int64 {
	var (
		tmp    []glProgram
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		libc.MemMove(unsafe.Pointer((*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i+1)))), unsafe.Pointer((*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glProgram{}))))
		*(*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i))) = a
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if tmp = unsafe.Slice((*glProgram)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glProgram{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		libc.MemMove(unsafe.Pointer((*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i+1)))), unsafe.Pointer((*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glProgram{}))))
		*(*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i))) = a
		vec.Capacity = tmp_sz
	}
	vec.Size++
	return 1
}
func cvec_insert_array_glProgram(vec *cvector_glProgram, i uint64, a *glProgram, num uint64) int64 {
	var (
		tmp    []glProgram
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_glProgram_SZ
		if tmp = unsafe.Slice((*glProgram)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glProgram{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	libc.MemMove(unsafe.Pointer((*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i+num)))), unsafe.Pointer((*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glProgram{}))))
	libc.MemMove(unsafe.Pointer((*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i)))), unsafe.Pointer(a), int(num*uint64(unsafe.Sizeof(glProgram{}))))
	vec.Size += num
	return 1
}
func cvec_replace_glProgram(vec *cvector_glProgram, i uint64, a glProgram) glProgram {
	var tmp glProgram = *(*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i)))
	*(*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i))) = a
	return tmp
}
func cvec_erase_glProgram(vec *cvector_glProgram, start uint64, end uint64) {
	var d uint64 = end - start + 1
	libc.MemMove(unsafe.Pointer((*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(start)))), unsafe.Pointer((*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(end+1)))), int((vec.Size-1-end)*uint64(unsafe.Sizeof(glProgram{}))))
	vec.Size -= d
}
func cvec_reserve_glProgram(vec *cvector_glProgram, size uint64) int64 {
	var tmp []glProgram
	if vec.Capacity < size {
		if tmp = unsafe.Slice((*glProgram)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int((size+CVEC_glProgram_SZ)*uint64(unsafe.Sizeof(glProgram{}))))), size+CVEC_glProgram_SZ); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = size + CVEC_glProgram_SZ
	}
	return 1
}
func cvec_set_cap_glProgram(vec *cvector_glProgram, size uint64) int64 {
	var tmp []glProgram
	if size < vec.Size {
		vec.Size = size
	}
	if tmp = unsafe.Slice((*glProgram)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(size*uint64(unsafe.Sizeof(glProgram{}))))), size); tmp == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	vec.A = tmp
	vec.Capacity = size
	return 1
}
func cvec_set_val_sz_glProgram(vec *cvector_glProgram, val glProgram) {
	var i uint64
	for i = 0; i < vec.Size; i++ {
		*(*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i))) = val
	}
}
func cvec_set_val_cap_glProgram(vec *cvector_glProgram, val glProgram) {
	var i uint64
	for i = 0; i < vec.Capacity; i++ {
		*(*glProgram)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glProgram{})*uintptr(i))) = val
	}
}
func cvec_clear_glProgram(vec *cvector_glProgram) {
	vec.Size = 0
}
func cvec_free_glProgram_heap(vec unsafe.Pointer) {
	var tmp *cvector_glProgram = (*cvector_glProgram)(vec)
	if tmp == nil {
		return
	}
	libc.Free(unsafe.Pointer(&tmp.A[0]))
	libc.Free(unsafe.Pointer(tmp))
}
func cvec_free_glProgram(vec unsafe.Pointer) {
	var tmp *cvector_glProgram = (*cvector_glProgram)(vec)
	libc.Free(unsafe.Pointer(&tmp.A[0]))
	tmp.Size = 0
	tmp.Capacity = 0
}

var CVEC_glVertex_SZ uint64 = 50

func cvec_glVertex_heap(size uint64, capacity uint64) *cvector_glVertex {
	var vec *cvector_glVertex
	if (func() *cvector_glVertex {
		vec = new(cvector_glVertex)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_glVertex_SZ
	}
	if vec.A = unsafe.Slice((*glVertex)(libc.Malloc(int(vec.Capacity*uint64(unsafe.Sizeof(glVertex{}))))), vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	return vec
}
func cvec_init_glVertex_heap(vals *glVertex, num uint64) *cvector_glVertex {
	var vec *cvector_glVertex
	if (func() *cvector_glVertex {
		vec = new(cvector_glVertex)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Capacity = num + CVEC_glVertex_SZ
	vec.Size = num
	if vec.A = unsafe.Slice((*glVertex)(libc.Malloc(int(vec.Capacity*uint64(unsafe.Sizeof(glVertex{}))))), vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	libc.MemMove(unsafe.Pointer(&vec.A[0]), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(glVertex{}))))
	return vec
}
func cvec_glVertex(vec *cvector_glVertex, size uint64, capacity uint64) int64 {
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_glVertex_SZ
	}
	if vec.A = unsafe.Slice((*glVertex)(libc.Malloc(int(vec.Capacity*uint64(unsafe.Sizeof(glVertex{}))))), vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	return 1
}
func cvec_init_glVertex(vec *cvector_glVertex, vals *glVertex, num uint64) int64 {
	vec.Capacity = num + CVEC_glVertex_SZ
	vec.Size = num
	if vec.A = unsafe.Slice((*glVertex)(libc.Malloc(int(vec.Capacity*uint64(unsafe.Sizeof(glVertex{}))))), vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	libc.MemMove(unsafe.Pointer(&vec.A[0]), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(glVertex{}))))
	return 1
}
func cvec_copyc_glVertex(dest unsafe.Pointer, src unsafe.Pointer) int64 {
	var (
		vec1 *cvector_glVertex = (*cvector_glVertex)(dest)
		vec2 *cvector_glVertex = (*cvector_glVertex)(src)
	)
	vec1.A = nil
	vec1.Size = 0
	vec1.Capacity = 0
	return cvec_copy_glVertex(vec1, vec2)
}
func cvec_copy_glVertex(dest *cvector_glVertex, src *cvector_glVertex) int64 {
	var tmp []glVertex = nil
	if tmp = unsafe.Slice((*glVertex)(libc.Realloc(unsafe.Pointer(&dest.A[0]), int(src.Capacity*uint64(unsafe.Sizeof(glVertex{}))))), src.Capacity); tmp == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	dest.A = tmp
	libc.MemMove(unsafe.Pointer(&dest.A[0]), unsafe.Pointer(&src.A[0]), int(src.Size*uint64(unsafe.Sizeof(glVertex{}))))
	dest.Size = src.Size
	dest.Capacity = src.Capacity
	return 1
}
func cvec_push_glVertex(vec *cvector_glVertex, a glVertex) int64 {
	var (
		tmp    []glVertex
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		*(*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(func() uint64 {
			p := &vec.Size
			x := *p
			*p++
			return x
		}()))) = a
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if tmp = unsafe.Slice((*glVertex)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glVertex{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		*(*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(func() uint64 {
			p := &vec.Size
			x := *p
			*p++
			return x
		}()))) = a
		vec.Capacity = tmp_sz
	}
	return 1
}
func cvec_pop_glVertex(vec *cvector_glVertex) glVertex {
	return *(*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(func() uint64 {
		p := &vec.Size
		*p--
		return *p
	}())))
}
func cvec_back_glVertex(vec *cvector_glVertex) *glVertex {
	return (*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(vec.Size-1)))
}
func cvec_extend_glVertex(vec *cvector_glVertex, num uint64) int64 {
	var (
		tmp    []glVertex
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_glVertex_SZ
		if tmp = unsafe.Slice((*glVertex)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glVertex{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	vec.Size += num
	return 1
}
func cvec_insert_glVertex(vec *cvector_glVertex, i uint64, a glVertex) int64 {
	var (
		tmp    []glVertex
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		libc.MemMove(unsafe.Pointer((*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i+1)))), unsafe.Pointer((*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glVertex{}))))
		*(*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i))) = a
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if tmp = unsafe.Slice((*glVertex)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glVertex{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		libc.MemMove(unsafe.Pointer((*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i+1)))), unsafe.Pointer((*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glVertex{}))))
		*(*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i))) = a
		vec.Capacity = tmp_sz
	}
	vec.Size++
	return 1
}
func cvec_insert_array_glVertex(vec *cvector_glVertex, i uint64, a *glVertex, num uint64) int64 {
	var (
		tmp    []glVertex
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_glVertex_SZ
		if tmp = unsafe.Slice((*glVertex)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(glVertex{}))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	libc.MemMove(unsafe.Pointer((*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i+num)))), unsafe.Pointer((*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(glVertex{}))))
	libc.MemMove(unsafe.Pointer((*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i)))), unsafe.Pointer(a), int(num*uint64(unsafe.Sizeof(glVertex{}))))
	vec.Size += num
	return 1
}
func cvec_replace_glVertex(vec *cvector_glVertex, i uint64, a glVertex) glVertex {
	var tmp glVertex = *(*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i)))
	*(*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i))) = a
	return tmp
}
func cvec_erase_glVertex(vec *cvector_glVertex, start uint64, end uint64) {
	var d uint64 = end - start + 1
	libc.MemMove(unsafe.Pointer((*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(start)))), unsafe.Pointer((*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(end+1)))), int((vec.Size-1-end)*uint64(unsafe.Sizeof(glVertex{}))))
	vec.Size -= d
}
func cvec_reserve_glVertex(vec *cvector_glVertex, size uint64) int64 {
	var tmp []glVertex
	if vec.Capacity < size {
		if tmp = unsafe.Slice((*glVertex)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int((size+CVEC_glVertex_SZ)*uint64(unsafe.Sizeof(glVertex{}))))), size+CVEC_glVertex_SZ); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = size + CVEC_glVertex_SZ
	}
	return 1
}
func cvec_set_cap_glVertex(vec *cvector_glVertex, size uint64) int64 {
	var tmp []glVertex
	if size < vec.Size {
		vec.Size = size
	}
	if tmp = unsafe.Slice((*glVertex)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(size*uint64(unsafe.Sizeof(glVertex{}))))), size); tmp == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	vec.A = tmp
	vec.Capacity = size
	return 1
}
func cvec_set_val_sz_glVertex(vec *cvector_glVertex, val glVertex) {
	var i uint64
	for i = 0; i < vec.Size; i++ {
		*(*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i))) = val
	}
}
func cvec_set_val_cap_glVertex(vec *cvector_glVertex, val glVertex) {
	var i uint64
	for i = 0; i < vec.Capacity; i++ {
		*(*glVertex)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(glVertex{})*uintptr(i))) = val
	}
}
func cvec_clear_glVertex(vec *cvector_glVertex) {
	vec.Size = 0
}
func cvec_free_glVertex_heap(vec unsafe.Pointer) {
	var tmp *cvector_glVertex = (*cvector_glVertex)(vec)
	if tmp == nil {
		return
	}
	libc.Free(unsafe.Pointer(&tmp.A[0]))
	libc.Free(unsafe.Pointer(tmp))
}
func cvec_free_glVertex(vec unsafe.Pointer) {
	var tmp *cvector_glVertex = (*cvector_glVertex)(vec)
	libc.Free(unsafe.Pointer(&tmp.A[0]))
	tmp.Size = 0
	tmp.Capacity = 0
}

var CVEC_float_SZ uint64 = 50

func cvec_float_heap(size uint64, capacity uint64) *cvector_float {
	var vec *cvector_float
	if (func() *cvector_float {
		vec = new(cvector_float)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_float_SZ
	}
	if vec.A = make([]float32, vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	return vec
}
func cvec_init_float_heap(vals *float32, num uint64) *cvector_float {
	var vec *cvector_float
	if (func() *cvector_float {
		vec = new(cvector_float)
		return vec
	}()) == nil {
		libc.Assert(vec != nil)
		return nil
	}
	vec.Capacity = num + CVEC_float_SZ
	vec.Size = num
	if vec.A = make([]float32, vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		libc.Free(unsafe.Pointer(vec))
		return nil
	}
	libc.MemMove(unsafe.Pointer(&vec.A[0]), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(float32(0)))))
	return vec
}
func cvec_float(vec *cvector_float, size uint64, capacity uint64) int64 {
	vec.Size = size
	if capacity > vec.Size || vec.Size != 0 && capacity == vec.Size {
		vec.Capacity = capacity
	} else {
		vec.Capacity = vec.Size + CVEC_float_SZ
	}
	if vec.A = make([]float32, vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	return 1
}
func cvec_init_float(vec *cvector_float, vals *float32, num uint64) int64 {
	vec.Capacity = num + CVEC_float_SZ
	vec.Size = num
	if vec.A = make([]float32, vec.Capacity); vec.A == nil {
		libc.Assert(vec.A != nil)
		vec.Size = func() uint64 {
			p := &vec.Capacity
			vec.Capacity = 0
			return *p
		}()
		return 0
	}
	libc.MemMove(unsafe.Pointer(&vec.A[0]), unsafe.Pointer(vals), int(num*uint64(unsafe.Sizeof(float32(0)))))
	return 1
}
func cvec_copyc_float(dest unsafe.Pointer, src unsafe.Pointer) int64 {
	var (
		vec1 *cvector_float = (*cvector_float)(dest)
		vec2 *cvector_float = (*cvector_float)(src)
	)
	vec1.A = nil
	vec1.Size = 0
	vec1.Capacity = 0
	return cvec_copy_float(vec1, vec2)
}
func cvec_copy_float(dest *cvector_float, src *cvector_float) int64 {
	var tmp []float32 = nil
	if tmp = unsafe.Slice((*float32)(libc.Realloc(unsafe.Pointer(&dest.A[0]), int(src.Capacity*uint64(unsafe.Sizeof(float32(0)))))), src.Capacity); tmp == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	dest.A = tmp
	libc.MemMove(unsafe.Pointer(&dest.A[0]), unsafe.Pointer(&src.A[0]), int(src.Size*uint64(unsafe.Sizeof(float32(0)))))
	dest.Size = src.Size
	dest.Capacity = src.Capacity
	return 1
}
func cvec_push_float(vec *cvector_float, a float32) int64 {
	var (
		tmp    []float32
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		*(*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(func() uint64 {
			p := &vec.Size
			x := *p
			*p++
			return x
		}()))) = a
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if tmp = unsafe.Slice((*float32)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(float32(0)))))), vec.Capacity); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		*(*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(func() uint64 {
			p := &vec.Size
			x := *p
			*p++
			return x
		}()))) = a
		vec.Capacity = tmp_sz
	}
	return 1
}
func cvec_pop_float(vec *cvector_float) float32 {
	return *(*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(func() uint64 {
		p := &vec.Size
		*p--
		return *p
	}())))
}
func cvec_back_float(vec *cvector_float) *float32 {
	return (*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(vec.Size-1)))
}
func cvec_extend_float(vec *cvector_float, num uint64) int64 {
	var (
		tmp    []float32
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_float_SZ
		if tmp = unsafe.Slice((*float32)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(float32(0)))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	vec.Size += num
	return 1
}
func cvec_insert_float(vec *cvector_float, i uint64, a float32) int64 {
	var (
		tmp    []float32
		tmp_sz uint64
	)
	if vec.Capacity > vec.Size {
		libc.MemMove(unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i+1)))), unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(float32(0)))))
		*(*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i))) = a
	} else {
		tmp_sz = (vec.Capacity + 1) * 2
		if tmp = unsafe.Slice((*float32)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(float32(0)))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		libc.MemMove(unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i+1)))), unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(float32(0)))))
		*(*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i))) = a
		vec.Capacity = tmp_sz
	}
	vec.Size++
	return 1
}
func cvec_insert_array_float(vec *cvector_float, i uint64, a *float32, num uint64) int64 {
	var (
		tmp    []float32
		tmp_sz uint64
	)
	if vec.Capacity < vec.Size+num {
		tmp_sz = vec.Capacity + num + CVEC_float_SZ
		if tmp = unsafe.Slice((*float32)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(tmp_sz*uint64(unsafe.Sizeof(float32(0)))))), tmp_sz); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = tmp_sz
	}
	libc.MemMove(unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i+num)))), unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i)))), int((vec.Size-i)*uint64(unsafe.Sizeof(float32(0)))))
	libc.MemMove(unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i)))), unsafe.Pointer(a), int(num*uint64(unsafe.Sizeof(float32(0)))))
	vec.Size += num
	return 1
}
func cvec_replace_float(vec *cvector_float, i uint64, a float32) float32 {
	var tmp float32 = *(*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i)))
	*(*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i))) = a
	return tmp
}
func cvec_erase_float(vec *cvector_float, start uint64, end uint64) {
	var d uint64 = end - start + 1
	libc.MemMove(unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(start)))), unsafe.Pointer((*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(end+1)))), int((vec.Size-1-end)*uint64(unsafe.Sizeof(float32(0)))))
	vec.Size -= d
}
func cvec_reserve_float(vec *cvector_float, size uint64) int64 {
	var tmp []float32
	if vec.Capacity < size {
		if tmp = unsafe.Slice((*float32)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int((size+CVEC_float_SZ)*uint64(unsafe.Sizeof(float32(0)))))), size+CVEC_float_SZ); tmp == nil {
			libc.Assert(tmp != nil)
			return 0
		}
		vec.A = tmp
		vec.Capacity = size + CVEC_float_SZ
	}
	return 1
}
func cvec_set_cap_float(vec *cvector_float, size uint64) int64 {
	var tmp []float32
	if size < vec.Size {
		vec.Size = size
	}
	if tmp = unsafe.Slice((*float32)(libc.Realloc(unsafe.Pointer(&vec.A[0]), int(size*uint64(unsafe.Sizeof(float32(0)))))), size); tmp == nil {
		libc.Assert(tmp != nil)
		return 0
	}
	vec.A = tmp
	vec.Capacity = size
	return 1
}
func cvec_set_val_sz_float(vec *cvector_float, val float32) {
	var i uint64
	for i = 0; i < vec.Size; i++ {
		*(*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i))) = val
	}
}
func cvec_set_val_cap_float(vec *cvector_float, val float32) {
	var i uint64
	for i = 0; i < vec.Capacity; i++ {
		*(*float32)(unsafe.Add(unsafe.Pointer(&vec.A[0]), unsafe.Sizeof(float32(0))*uintptr(i))) = val
	}
}
func cvec_clear_float(vec *cvector_float) {
	vec.Size = 0
}
func cvec_free_float_heap(vec unsafe.Pointer) {
	var tmp *cvector_float = (*cvector_float)(vec)
	if tmp == nil {
		return
	}
	libc.Free(unsafe.Pointer(&tmp.A[0]))
	libc.Free(unsafe.Pointer(tmp))
}
func cvec_free_float(vec unsafe.Pointer) {
	var tmp *cvector_float = (*cvector_float)(vec)
	tmp.A = nil
	tmp.Size = 0
	tmp.Capacity = 0
}

var c *GlContext

func gl_clipcode(pt Vec4) int64 {
	var w float32
	w = float32(float64(pt.W) * (1.0 + 1e-05))
	return ((int64(libc.BoolToInt(pt.Z < -w)) | int64(libc.BoolToInt(pt.Z > w))<<1) & (int64(libc.BoolToInt(c.Depth_clamp == 0)) | int64(libc.BoolToInt(c.Depth_clamp == 0))<<1)) | int64(libc.BoolToInt(pt.X < -w))<<2 | int64(libc.BoolToInt(pt.X > w))<<3 | int64(libc.BoolToInt(pt.Y < -w))<<4 | int64(libc.BoolToInt(pt.Y > w))<<5
}
func is_front_facing(v0 *glVertex, v1 *glVertex, v2 *glVertex) int64 {
	var (
		normal  Vec3
		tmpvec3 Vec3 = Vec3{X: 0, Y: 0, Z: 1}
		p0      Vec3 = vec4_to_vec3h(v0.Screen_space)
		p1      Vec3 = vec4_to_vec3h(v1.Screen_space)
		p2      Vec3 = vec4_to_vec3h(v2.Screen_space)
	)
	normal = cross_product(sub_vec3s(p1, p0), sub_vec3s(p2, p0))
	if c.Front_face == GLenum(GL_CW) {
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
		buf_pos = c.Buffers.A[buf].Data[v[enabled[j]].Offset+GLsizei(uint64(v[enabled[j]].Stride)*i):]
		tmpvec4.X = 0.0
		tmpvec4.Y = 0.0
		tmpvec4.Z = 0.0
		tmpvec4.W = 1.0
		libc.MemCpy(unsafe.Pointer(&tmpvec4), unsafe.Pointer(&buf_pos[0]), int(uintptr(v[enabled[j]].Size)*unsafe.Sizeof(float32(0))))
		c.Vertex_attribs_vs[enabled[j]] = tmpvec4
	}
	var vs_out *float32 = &c.Vs_output.Output_buf.A[vert*uint64(c.Vs_output.Size)]
	c.Programs.A[c.Cur_program].Vertex_shader(vs_out, unsafe.Pointer(&c.Vertex_attribs_vs[0]), &c.Builtins, c.Programs.A[c.Cur_program].Uniform)
	c.Glverts.A[vert].Vs_out = unsafe.Slice(vs_out, c.Vs_output.Size)
	c.Glverts.A[vert].Clip_space = c.Builtins.Gl_Position
	c.Glverts.A[vert].Edge_flag = 1
	c.Glverts.A[vert].Clip_code = gl_clipcode(c.Builtins.Gl_Position)
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
		enabled     [GL_MAX_VERTEX_ATTRIBS]int64
	)
	var v []glVertex_Attrib = c.Vertex_arrays.A[c.Cur_vertex_array].Vertex_attribs[:]
	var elem_buffer GLuint = (c.Vertex_arrays.A[c.Cur_vertex_array]).Element_buffer
	for i, j = 0, 0; i < GL_MAX_VERTEX_ATTRIBS; i++ {
		c.Vertex_attribs_vs[i] = vec4_init
		if v[i].Enabled != 0 {
			if v[i].Divisor == 0 {
				enabled[j] = int64(i)
				j++
			} else if (instance_id % GLsizei(v[i].Divisor)) == 0 {
				var n int64 = int64(instance_id/GLsizei(v[i].Divisor) + GLsizei(base_instance))
				buf_pos = c.Buffers.A[v[i].Buf].Data[v[i].Offset+v[i].Stride*GLsizei(n):]
				tmpvec4.X = 0.0
				tmpvec4.Y = 0.0
				tmpvec4.Z = 0.0
				tmpvec4.W = 1.0
				libc.MemCpy(unsafe.Pointer(&tmpvec4), unsafe.Pointer(&buf_pos[0]), int(uintptr(v[enabled[j]].Size)*unsafe.Sizeof(float32(0))))
				c.Vertex_attribs_vs[i] = tmpvec4
			}
		}
	}
	num_enabled = j
	cvec_reserve_glVertex(&c.Glverts, uint64(count))
	c.Builtins.Gl_InstanceID = GLint(instance_id)
	if use_elements == 0 {
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)); vert, i = vert+1, i+1 {
			do_vertex(v, enabled[:], num_enabled, i, vert)
		}
	} else {
		var (
			uint_array   []GLuint   = unsafe.Slice((*GLuint)(unsafe.Pointer(&c.Buffers.A[elem_buffer].Data[0])), first+GLint(count))
			ushort_array []GLushort = unsafe.Slice((*GLushort)(unsafe.Pointer(&c.Buffers.A[elem_buffer].Data[0])), first+GLint(count))
			ubyte_array  []GLubyte  = unsafe.Slice((*GLubyte)(unsafe.Pointer(&c.Buffers.A[elem_buffer].Data[0])), first+GLint(count))
		)
		if c.Buffers.A[elem_buffer].Type == GLenum(GL_UNSIGNED_BYTE) {
			for vert, i = 0, uint64(0); i < uint64(first+GLint(count)); vert, i = vert+1, i+1 {
				do_vertex(v, enabled[:], num_enabled, uint64(ubyte_array[i]), vert)
			}
		} else if c.Buffers.A[elem_buffer].Type == GLenum(GL_UNSIGNED_SHORT) {
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
		fs_input [64]float32
		point    Vec3 = vec4_to_vec3h(vert.Screen_space)
	)
	point.Z = float32((float64(point.Z)-(-1.0))/(1.0-(-1.0))*float64(c.Depth_range_far-c.Depth_range_near) + float64(c.Depth_range_near))
	if c.Depth_clamp != 0 {
		point.Z = clampf_01(point.Z)
	}
	copy(fs_input[:], vert.Vs_out[:c.Vs_output.Size]) //libc.MemCpy(unsafe.Pointer(&fs_input[0]), unsafe.Pointer(vert.Vs_out), int(c.Vs_output.Size*int64(unsafe.Sizeof(float32(0)))))
	var x float32 = float32(float64(point.X) + 0.5)
	var y float32 = float32(float64(point.Y) + 0.5)
	var p_size float32 = float32(c.Point_size)
	var origin float32
	if c.Point_spr_origin == GLenum(GL_UPPER_LEFT) {
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
			for {
				c.Builtins.Gl_FragCoord.X = j
				c.Builtins.Gl_FragCoord.Y = i
				c.Builtins.Gl_FragCoord.Z = point.Z
				c.Builtins.Gl_FragCoord.W = 1 / vert.Screen_space.W
				if true {
					break
				}
			}
			c.Builtins.Discard = GL_FALSE
			c.Builtins.Gl_FragDepth = point.Z
			c.Programs.A[c.Cur_program].Fragment_shader(&fs_input[0], &c.Builtins, c.Programs.A[c.Cur_program].Uniform)
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
	libc.Assert(count <= MAX_VERTICES)
	vertex_stage(first, count, instance, base_instance, use_elements)
	if mode == GLenum(GL_POINTS) {
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)); i, vert = i+1, vert+1 {
			if c.Glverts.A[vert].Clip_code != 0 {
				continue
			}
			c.Glverts.A[vert].Screen_space = Mult_mat4_vec4(c.Vp_mat, c.Glverts.A[vert].Clip_space)
			draw_point((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*uintptr(vert))))
		}
	} else if mode == GLenum(GL_LINES) {
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)-1); func() uint64 {
			i += 2
			return func() uint64 {
				vert += 2
				return vert
			}()
		}() {
			draw_line_clip((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*uintptr(vert))), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*uintptr(vert+1))))
		}
	} else if mode == GLenum(GL_LINE_STRIP) {
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)-1); func() uint64 {
			i++
			return func() uint64 {
				p := &vert
				x := *p
				*p++
				return x
			}()
		}() {
			draw_line_clip((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*uintptr(vert))), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*uintptr(vert+1))))
		}
	} else if mode == GLenum(GL_LINE_LOOP) {
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)-1); func() uint64 {
			i++
			return func() uint64 {
				p := &vert
				x := *p
				*p++
				return x
			}()
		}() {
			draw_line_clip((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*uintptr(vert))), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*uintptr(vert+1))))
		}
		draw_line_clip((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*uintptr(count-1))), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*0)))
	} else if mode == GLenum(GL_TRIANGLES) {
		if c.Provoking_vert == GLenum(GL_LAST_VERTEX_CONVENTION) {
			provoke = 2
		} else {
			provoke = 0
		}
		for vert, i = 0, uint64(first); i < uint64(first+GLint(count)-2); i, vert = i+3, vert+3 {
			draw_triangle(&c.Glverts.A[vert], &c.Glverts.A[vert+1], &c.Glverts.A[vert+2], vert+uint64(provoke))
		}
	} else if mode == GLenum(GL_TRIANGLE_STRIP) {
		var (
			a      uint64 = 0
			b      uint64 = 1
			toggle uint64 = 0
		)
		if c.Provoking_vert == GLenum(GL_LAST_VERTEX_CONVENTION) {
			provoke = 0
		} else {
			provoke = -2
		}
		for vert = 2; vert < uint64(count); vert++ {
			draw_triangle(&c.Glverts.A[a], &c.Glverts.A[b], &c.Glverts.A[vert], vert+uint64(provoke))
			if toggle == 0 {
				a = vert
			} else {
				b = vert
			}
			toggle = uint64(libc.BoolToInt(toggle == 0))
		}
	} else if mode == GLenum(GL_TRIANGLE_FAN) {
		if c.Provoking_vert == GLenum(GL_LAST_VERTEX_CONVENTION) {
			provoke = 0
		} else {
			provoke = -1
		}
		for vert = 2; vert < uint64(count); vert++ {
			draw_triangle((*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*0)), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*uintptr(vert-1))), (*glVertex)(unsafe.Add(unsafe.Pointer(&c.Glverts.A[0]), unsafe.Sizeof(glVertex{})*uintptr(vert))), vert+uint64(provoke))
		}
	}
}
func depthtest(zval float32, zbufval float32) int64 {
	if c.Depth_mask == 0 {
		return 0
	}
	switch c.Depth_func {
	case GL_LESS:
		return int64(libc.BoolToInt(zval < zbufval))
	case GL_LEQUAL:
		return int64(libc.BoolToInt(zval <= zbufval))
	case GL_GREATER:
		return int64(libc.BoolToInt(zval > zbufval))
	case GL_GEQUAL:
		return int64(libc.BoolToInt(zval >= zbufval))
	case GL_EQUAL:
		return int64(libc.BoolToInt(zval == zbufval))
	case GL_NOTEQUAL:
		return int64(libc.BoolToInt(zval != zbufval))
	case GL_ALWAYS:
		return 1
	case GL_NEVER:
		return 0
	}
	return 0
}
func setup_fs_input(t float32, v1_out *float32, v2_out *float32, wa float32, wb float32, provoke uint64) {
	var (
		vs_output *float32 = &c.Vs_output.Output_buf.A[0]
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
	c.Builtins.Discard = GL_FALSE
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
	var v1_out [64]float32
	var v2_out [64]float32
	var provoke uint64
	if c.Provoking_vert == GLenum(GL_LAST_VERTEX_CONVENTION) {
		provoke = uint64((int64(uintptr(unsafe.Pointer(v2)) - uintptr(unsafe.Pointer(&c.Glverts.A[0])))) / int64(unsafe.Sizeof(glVertex{})))
	} else {
		provoke = uint64((int64(uintptr(unsafe.Pointer(v1)) - uintptr(unsafe.Pointer(&c.Glverts.A[0])))) / int64(unsafe.Sizeof(glVertex{})))
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
		hp1     Vec3    = vec4_to_vec3h(v1)
		hp2     Vec3    = vec4_to_vec3h(v2)
		w1      float32 = v1.W
		w2      float32 = v2.W
		x1      float32 = hp1.X
		x2      float32 = hp2.X
		y1      float32 = hp1.Y
		y2      float32 = hp2.Y
		z1      float32 = hp1.Z
		z2      float32 = hp2.Z
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
	var fragment_shader frag_func = c.Programs.A[c.Cur_program].Fragment_shader
	var uniform interface{} = c.Programs.A[c.Cur_program].Uniform
	var fragdepth_or_discard int64 = int64(c.Programs.A[c.Cur_program].Fragdepth_or_discard)
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
			for {
				c.Builtins.Gl_FragCoord.X = x
				c.Builtins.Gl_FragCoord.Y = y
				c.Builtins.Gl_FragCoord.Z = z
				c.Builtins.Gl_FragCoord.W = 1 / w
				if true {
					break
				}
			}
			c.Builtins.Discard = GL_FALSE
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
			for {
				c.Builtins.Gl_FragCoord.X = x
				c.Builtins.Gl_FragCoord.Y = y
				c.Builtins.Gl_FragCoord.Z = z
				c.Builtins.Gl_FragCoord.W = 1 / w
				if true {
					break
				}
			}
			c.Builtins.Discard = GL_FALSE
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
			for {
				c.Builtins.Gl_FragCoord.X = x
				c.Builtins.Gl_FragCoord.Y = y
				c.Builtins.Gl_FragCoord.Z = z
				c.Builtins.Gl_FragCoord.W = 1 / w
				if true {
					break
				}
			}
			c.Builtins.Discard = GL_FALSE
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
			for {
				c.Builtins.Gl_FragCoord.X = x
				c.Builtins.Gl_FragCoord.Y = y
				c.Builtins.Gl_FragCoord.Z = z
				c.Builtins.Gl_FragCoord.W = 1 / w
				if true {
					break
				}
			}
			c.Builtins.Discard = GL_FALSE
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
		fragment_shader      frag_func   = c.Programs.A[c.Cur_program].Fragment_shader
		uniform              interface{} = c.Programs.A[c.Cur_program].Uniform
		fragdepth_or_discard int64       = int64(c.Programs.A[c.Cur_program].Fragdepth_or_discard)
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
		steep                int64 = int64(libc.BoolToInt(math32.Abs(y2-y1) > math32.Abs(x2-x1)))
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
			for {
				c.Builtins.Gl_FragCoord.X = ypxl1
				c.Builtins.Gl_FragCoord.Y = xpxl1
				c.Builtins.Gl_FragCoord.Z = z1
				c.Builtins.Gl_FragCoord.W = 1 / w1
				if true {
					break
				}
			}
			setup_fs_input(0, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(ypxl1), int64(xpxl1))
			}
			for {
				c.Builtins.Gl_FragCoord.X = ypxl1 + 1
				c.Builtins.Gl_FragCoord.Y = xpxl1
				c.Builtins.Gl_FragCoord.Z = z1
				c.Builtins.Gl_FragCoord.W = 1 / w1
				if true {
					break
				}
			}
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
			for {
				c.Builtins.Gl_FragCoord.X = xpxl1
				c.Builtins.Gl_FragCoord.Y = ypxl1
				c.Builtins.Gl_FragCoord.Z = z1
				c.Builtins.Gl_FragCoord.W = 1 / w1
				if true {
					break
				}
			}
			setup_fs_input(0, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(xpxl1), int64(ypxl1))
			}
			for {
				c.Builtins.Gl_FragCoord.X = xpxl1
				c.Builtins.Gl_FragCoord.Y = ypxl1 + 1
				c.Builtins.Gl_FragCoord.Z = z1
				c.Builtins.Gl_FragCoord.W = 1 / w1
				if true {
					break
				}
			}
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
			for {
				c.Builtins.Gl_FragCoord.X = ypxl2
				c.Builtins.Gl_FragCoord.Y = xpxl2
				c.Builtins.Gl_FragCoord.Z = z2
				c.Builtins.Gl_FragCoord.W = 1 / w2
				if true {
					break
				}
			}
			setup_fs_input(1, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(ypxl2), int64(xpxl2))
			}
			for {
				c.Builtins.Gl_FragCoord.X = ypxl2 + 1
				c.Builtins.Gl_FragCoord.Y = xpxl2
				c.Builtins.Gl_FragCoord.Z = z2
				c.Builtins.Gl_FragCoord.W = 1 / w2
				if true {
					break
				}
			}
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
			for {
				c.Builtins.Gl_FragCoord.X = xpxl2
				c.Builtins.Gl_FragCoord.Y = ypxl2
				c.Builtins.Gl_FragCoord.Z = z2
				c.Builtins.Gl_FragCoord.W = 1 / w2
				if true {
					break
				}
			}
			setup_fs_input(1, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(xpxl2), int64(ypxl2))
			}
			for {
				c.Builtins.Gl_FragCoord.X = xpxl2
				c.Builtins.Gl_FragCoord.Y = ypxl2 + 1
				c.Builtins.Gl_FragCoord.Z = z2
				c.Builtins.Gl_FragCoord.W = 1 / w2
				if true {
					break
				}
			}
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
			for {
				c.Builtins.Gl_FragCoord.X = intery
				c.Builtins.Gl_FragCoord.Y = float32(x)
				c.Builtins.Gl_FragCoord.Z = z
				c.Builtins.Gl_FragCoord.W = 1 / w
				if true {
					break
				}
			}
			setup_fs_input(t, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(intery), x)
			}
			for {
				c.Builtins.Gl_FragCoord.X = intery + 1
				c.Builtins.Gl_FragCoord.Y = float32(x)
				c.Builtins.Gl_FragCoord.Z = z
				c.Builtins.Gl_FragCoord.W = 1 / w
				if true {
					break
				}
			}
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
			for {
				c.Builtins.Gl_FragCoord.X = float32(x)
				c.Builtins.Gl_FragCoord.Y = intery
				c.Builtins.Gl_FragCoord.Z = z
				c.Builtins.Gl_FragCoord.W = 1 / w
				if true {
					break
				}
			}
			setup_fs_input(t, v1_out, v2_out, w1, w2, provoke)
			fragment_shader(&c.Fs_input[0], &c.Builtins, uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, x, int64(intery))
			}
			for {
				c.Builtins.Gl_FragCoord.X = float32(x)
				c.Builtins.Gl_FragCoord.Y = intery + 1
				c.Builtins.Gl_FragCoord.Z = z
				c.Builtins.Gl_FragCoord.W = 1 / w
				if true {
					break
				}
			}
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
		if c.Cull_mode == GLenum(GL_FRONT_AND_BACK) {
			return
		}
		if c.Cull_mode == GLenum(GL_BACK) && front_facing == 0 {
			return
		}
		if c.Cull_mode == GLenum(GL_FRONT) && front_facing != 0 {
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

var clip_proc [6]*func(*Vec4, *Vec4, *Vec4) float32 = [6]*func(*Vec4, *Vec4, *Vec4) float32{(*func(*Vec4, *Vec4, *Vec4) float32)(unsafe.Pointer(libc.FuncAddr(clip_zmin))), (*func(*Vec4, *Vec4, *Vec4) float32)(unsafe.Pointer(libc.FuncAddr(clip_zmax))), (*func(*Vec4, *Vec4, *Vec4) float32)(unsafe.Pointer(libc.FuncAddr(clip_xmin))), (*func(*Vec4, *Vec4, *Vec4) float32)(unsafe.Pointer(libc.FuncAddr(clip_xmax))), (*func(*Vec4, *Vec4, *Vec4) float32)(unsafe.Pointer(libc.FuncAddr(clip_ymin))), (*func(*Vec4, *Vec4, *Vec4) float32)(unsafe.Pointer(libc.FuncAddr(clip_ymax)))}

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
		tmp1_out      [64]float32
		tmp2_out      [64]float32
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
			stdio.Printf("Clipping error:\n")
			print_vec4(v0.Clip_space, libc.CString("\n"))
			print_vec4(v1.Clip_space, libc.CString("\n"))
			print_vec4(v2.Clip_space, libc.CString("\n"))
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
			tt = libc.AsFunc(clip_proc[clip_bit], (*func(*Vec4, *Vec4, *Vec4) float32)(nil)).(func(*Vec4, *Vec4, *Vec4) float32)(&tmp1.Clip_space, &q[0].Clip_space, &q[1].Clip_space)
			update_clip_pt(&tmp1, q[0], q[1], tt)
			tt = libc.AsFunc(clip_proc[clip_bit], (*func(*Vec4, *Vec4, *Vec4) float32)(nil)).(func(*Vec4, *Vec4, *Vec4) float32)(&tmp2.Clip_space, &q[0].Clip_space, &q[2].Clip_space)
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
			tt = libc.AsFunc(clip_proc[clip_bit], (*func(*Vec4, *Vec4, *Vec4) float32)(nil)).(func(*Vec4, *Vec4, *Vec4) float32)(&tmp1.Clip_space, &q[0].Clip_space, &q[1].Clip_space)
			update_clip_pt(&tmp1, q[0], q[1], tt)
			tt = libc.AsFunc(clip_proc[clip_bit], (*func(*Vec4, *Vec4, *Vec4) float32)(nil)).(func(*Vec4, *Vec4, *Vec4) float32)(&tmp2.Clip_space, &q[0].Clip_space, &q[2].Clip_space)
			update_clip_pt(&tmp2, q[0], q[2], tt)
			tmp1.Edge_flag = 1
			tmp2.Edge_flag = q[2].Edge_flag
			draw_triangle_clip(q[0], &tmp1, &tmp2, provoke, clip_bit+1)
		}
	}
}
func draw_triangle_point(v0 *glVertex, v1 *glVertex, v2 *glVertex, provoke uint64) {
	var (
		fs_input [64]float32
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
				fs_input[j] = *(*float32)(unsafe.Add(unsafe.Pointer(&c.Vs_output.Output_buf.A[0]), unsafe.Sizeof(float32(0))*uintptr(provoke*uint64(c.Vs_output.Size)+uint64(j))))
			}
		}
		c.Builtins.Discard = GL_FALSE
		c.Programs.A[c.Cur_program].Fragment_shader(&fs_input[0], &c.Builtins, c.Programs.A[c.Cur_program].Uniform)
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
	var l01 Line = make_Line(hp0.X, hp0.Y, hp1.X, hp1.Y)
	var l12 Line = make_Line(hp1.X, hp1.Y, hp2.X, hp2.Y)
	var l20 Line = make_Line(hp2.X, hp2.Y, hp0.X, hp0.Y)
	var alpha float32
	var beta float32
	var gamma float32
	var tmp float32
	var tmp2 float32
	var z float32
	var fs_input [64]float32
	var perspective [192]float32
	var vs_output = &c.Vs_output.Output_buf.A[0]
	for i := int64(0); i < c.Vs_output.Size; i++ {
		perspective[i] = v0.Vs_out[i] / p0.W
		perspective[GL_MAX_VERTEX_OUTPUT_COMPONENTS+i] = v1.Vs_out[i] / p1.W
		perspective[GL_MAX_VERTEX_OUTPUT_COMPONENTS*2+i] = v2.Vs_out[i] / p2.W
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
							tmp = alpha*perspective[i] + beta*perspective[GL_MAX_VERTEX_OUTPUT_COMPONENTS+i] + gamma*perspective[GL_MAX_VERTEX_OUTPUT_COMPONENTS*2+i]
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
					c.Builtins.Discard = GL_FALSE
					c.Builtins.Gl_FragDepth = z
					c.Programs.A[c.Cur_program].Fragment_shader(&fs_input[0], &c.Builtins, c.Programs.A[c.Cur_program].Uniform)
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
	case GL_ZERO:
		Cs.X = 0
		Cs.Y = 0
		Cs.Z = 0
		Cs.W = 0
	case GL_ONE:
		Cs.X = 1
		Cs.Y = 1
		Cs.Z = 1
		Cs.W = 1
	case GL_SRC_COLOR:
		Cs = src
	case GL_ONE_MINUS_SRC_COLOR:
		Cs.X = 1 - src.X
		Cs.Y = 1 - src.Y
		Cs.Z = 1 - src.Z
		Cs.W = 1 - src.W
	case GL_DST_COLOR:
		Cs = dst
	case GL_ONE_MINUS_DST_COLOR:
		Cs.X = 1 - dst.X
		Cs.Y = 1 - dst.Y
		Cs.Z = 1 - dst.Z
		Cs.W = 1 - dst.W
	case GL_SRC_ALPHA:
		Cs.X = src.W
		Cs.Y = src.W
		Cs.Z = src.W
		Cs.W = src.W
	case GL_ONE_MINUS_SRC_ALPHA:
		Cs.X = 1 - src.W
		Cs.Y = 1 - src.W
		Cs.Z = 1 - src.W
		Cs.W = 1 - src.W
	case GL_DST_ALPHA:
		Cs.X = dst.W
		Cs.Y = dst.W
		Cs.Z = dst.W
		Cs.W = dst.W
	case GL_ONE_MINUS_DST_ALPHA:
		Cs.X = 1 - dst.W
		Cs.Y = 1 - dst.W
		Cs.Z = 1 - dst.W
		Cs.W = 1 - dst.W
	case GL_CONSTANT_COLOR:
		Cs = *cnst
	case GL_ONE_MINUS_CONSTANT_COLOR:
		Cs.X = 1 - cnst.X
		Cs.Y = 1 - cnst.Y
		Cs.Z = 1 - cnst.Z
		Cs.W = 1 - cnst.W
	case GL_CONSTANT_ALPHA:
		Cs.X = cnst.W
		Cs.Y = cnst.W
		Cs.Z = cnst.W
		Cs.W = cnst.W
	case GL_ONE_MINUS_CONSTANT_ALPHA:
		Cs.X = 1 - cnst.W
		Cs.Y = 1 - cnst.W
		Cs.Z = 1 - cnst.W
		Cs.W = 1 - cnst.W
	case GL_SRC_ALPHA_SATURATE:
		Cs.X = i
		Cs.Y = i
		Cs.Z = i
		Cs.W = 1
	default:
		stdio.Printf("error unrecognized blend_sfactor!\n")
	}
	switch c.Blend_dfactor {
	case GL_ZERO:
		Cd.X = 0
		Cd.Y = 0
		Cd.Z = 0
		Cd.W = 0
	case GL_ONE:
		Cd.X = 1
		Cd.Y = 1
		Cd.Z = 1
		Cd.W = 1
	case GL_SRC_COLOR:
		Cd = src
	case GL_ONE_MINUS_SRC_COLOR:
		Cd.X = 1 - src.X
		Cd.Y = 1 - src.Y
		Cd.Z = 1 - src.Z
		Cd.W = 1 - src.W
	case GL_DST_COLOR:
		Cd = dst
	case GL_ONE_MINUS_DST_COLOR:
		Cd.X = 1 - dst.X
		Cd.Y = 1 - dst.Y
		Cd.Z = 1 - dst.Z
		Cd.W = 1 - dst.W
	case GL_SRC_ALPHA:
		Cd.X = src.W
		Cd.Y = src.W
		Cd.Z = src.W
		Cd.W = src.W
	case GL_ONE_MINUS_SRC_ALPHA:
		Cd.X = 1 - src.W
		Cd.Y = 1 - src.W
		Cd.Z = 1 - src.W
		Cd.W = 1 - src.W
	case GL_DST_ALPHA:
		Cd.X = dst.W
		Cd.Y = dst.W
		Cd.Z = dst.W
		Cd.W = dst.W
	case GL_ONE_MINUS_DST_ALPHA:
		Cd.X = 1 - dst.W
		Cd.Y = 1 - dst.W
		Cd.Z = 1 - dst.W
		Cd.W = 1 - dst.W
	case GL_CONSTANT_COLOR:
		Cd = *cnst
	case GL_ONE_MINUS_CONSTANT_COLOR:
		Cd.X = 1 - cnst.X
		Cd.Y = 1 - cnst.Y
		Cd.Z = 1 - cnst.Z
		Cd.W = 1 - cnst.W
	case GL_CONSTANT_ALPHA:
		Cd.X = cnst.W
		Cd.Y = cnst.W
		Cd.Z = cnst.W
		Cd.W = cnst.W
	case GL_ONE_MINUS_CONSTANT_ALPHA:
		Cd.X = 1 - cnst.W
		Cd.Y = 1 - cnst.W
		Cd.Z = 1 - cnst.W
		Cd.W = 1 - cnst.W
	case GL_SRC_ALPHA_SATURATE:
		Cd.X = i
		Cd.Y = i
		Cd.Z = i
		Cd.W = 1
	default:
		stdio.Printf("error unrecognized blend_dfactor!\n")
	}
	var result Vec4
	switch c.Blend_equation {
	case GL_FUNC_ADD:
		result = add_vec4s(mult_vec4s(Cs, src), mult_vec4s(Cd, dst))
	case GL_FUNC_SUBTRACT:
		result = sub_vec4s(mult_vec4s(Cs, src), mult_vec4s(Cd, dst))
	case GL_FUNC_REVERSE_SUBTRACT:
		result = sub_vec4s(mult_vec4s(Cd, dst), mult_vec4s(Cs, src))
	case GL_MIN:
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
	case GL_MAX:
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
		stdio.Printf("error unrecognized blend_equation!\n")
	}
	return vec4_to_Color(result)
}
func logic_ops_pixel(s Color, d Color) Color {
	switch c.Logic_func {
	case GL_CLEAR:
		return make_Color(0, 0, 0, 0)
	case GL_SET:
		return make_Color(math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8)
	case GL_COPY:
		return s
	case GL_COPY_INVERTED:
		return make_Color(^s.R, ^s.G, ^s.B, ^s.A)
	case GL_NOOP:
		return d
	case GL_INVERT:
		return make_Color(^d.R, ^d.G, ^d.B, ^d.A)
	case GL_AND:
		return make_Color(s.R&d.R, s.G&d.G, s.B&d.B, s.A&d.A)
	case GL_NAND:
		return make_Color(^(s.R & d.R), ^(s.G & d.G), ^(s.B & d.B), ^(s.A & d.A))
	case GL_OR:
		return make_Color(s.R|d.R, s.G|d.G, s.B|d.B, s.A|d.A)
	case GL_NOR:
		return make_Color(^(s.R | d.R), ^(s.G | d.G), ^(s.B | d.B), ^(s.A | d.A))
	case GL_XOR:
		return make_Color(s.R^d.R, s.G^d.G, s.B^d.B, s.A^d.A)
	case GL_EQUIV:
		return make_Color(^(s.R ^ d.R), ^(s.G ^ d.G), ^(s.B ^ d.B), ^(s.A ^ d.A))
	case GL_AND_REVERSE:
		return make_Color(s.R & ^d.R, s.G & ^d.G, s.B & ^d.B, s.A & ^d.A)
	case GL_AND_INVERTED:
		return make_Color(^s.R&d.R, ^s.G&d.G, ^s.B&d.B, ^s.A&d.A)
	case GL_OR_REVERSE:
		return make_Color(s.R|^d.R, s.G|^d.G, s.B|^d.B, s.A|^d.A)
	case GL_OR_INVERTED:
		return make_Color(^s.R|d.R, ^s.G|d.G, ^s.B|d.B, ^s.A|d.A)
	default:
		fmt.Println(("Unrecognized logic op!, defaulting to GL_COPY"))
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
	case GL_NEVER:
		return 0
	case GL_LESS:
		return int64(libc.BoolToInt((ref & mask) < (int64(stencil) & mask)))
	case GL_LEQUAL:
		return int64(libc.BoolToInt((ref & mask) <= (int64(stencil) & mask)))
	case GL_GREATER:
		return int64(libc.BoolToInt((ref & mask) > (int64(stencil) & mask)))
	case GL_GEQUAL:
		return int64(libc.BoolToInt((ref & mask) >= (int64(stencil) & mask)))
	case GL_EQUAL:
		return int64(libc.BoolToInt((ref & mask) == (int64(stencil) & mask)))
	case GL_NOTEQUAL:
		return int64(libc.BoolToInt((ref & mask) != (int64(stencil) & mask)))
	case GL_ALWAYS:
		return 1
	default:
		fmt.Println(("Error: unrecognized stencil function!"))
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
	case GL_KEEP:
		return
	case GL_ZERO:
		val = 0
	case GL_REPLACE:
		val = u8(int8(ref))
	case GL_INCR:
		if val < math.MaxUint8 {
			val++
		}
	case GL_INCR_WRAP:
		val++
	case GL_DECR:
		if val > 0 {
			val--
		}
	case GL_DECR_WRAP:
		val--
	case GL_INVERT:
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
	*dest = U32(int32(int64(src_color.A)<<c.Ashift | int64(src_color.R)<<c.Rshift | int64(src_color.G)<<c.Gshift | int64(src_color.B)<<c.Bshift))
}
func is_valid(target GLenum, error GLenum, n int64, _rest ...interface{}) int64 {
	var argptr libc.ArgList
	argptr.Start(n, _rest)
	for i := int64(0); i < n; i++ {
		if target == argptr.Arg().(GLenum) {
			return 1
		}
	}
	argptr.End()
	if c.Error == 0 {
		c.Error = error
	}
	return 0
}
func default_vs(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *Shader_Builtins, uniforms interface{}) {
	builtins.Gl_Position = *(*Vec4)(unsafe.Add(unsafe.Pointer((*Vec4)(vertex_attribs)), unsafe.Sizeof(Vec4{})*0))
}
func default_fs(fs_input *float32, builtins *Shader_Builtins, uniforms interface{}) {
	var fragcolor *Vec4 = &builtins.Gl_FragColor
	fragcolor.X = 1.0
	fragcolor.Y = 0.0
	fragcolor.Z = 0.0
	fragcolor.W = 1.0
}
func init_glVertex_Array(v *glVertex_Array) {
	v.Deleted = GL_FALSE
	for i := int64(0); i < GL_MAX_VERTEX_ATTRIBS; i++ {
		init_glVertex_Attrib(&v.Vertex_attribs[i])
	}
}
func init_glVertex_Attrib(v *glVertex_Attrib) {
	v.Buf = 0
	v.Enabled = 0
	v.Divisor = 0
}
func Init_glContext(context *GlContext, back *[]U32, w int64, h int64, bitdepth int64, Rmask U32, Gmask U32, Bmask U32, Amask U32) int64 {
	if bitdepth > 32 || back == nil {
		return 0
	}
	context.User_alloced_backbuf = int64(libc.BoolToInt(*back != nil))
	if *back == nil {
		var bytes_per_pixel int64 = (bitdepth + 8 - 1) / 8
		*back = unsafe.Slice((*U32)(libc.Malloc(int(w*h*bytes_per_pixel))), w*h*bytes_per_pixel)
		if *back == nil {
			return 0
		}
	}
	context.Zbuf.Buf = make([]u8, w*h*int64(unsafe.Sizeof(float32(0)))) //(*u8)(libc.Malloc(int(w * h * int64(unsafe.Sizeof(float32(0))))))
	if context.Zbuf.Buf == nil {
		if context.User_alloced_backbuf == 0 {
			*back = nil
		}
		return 0
	}
	context.Stencil_buf.Buf = make([]u8, w*h) //(*u8)(libc.Malloc(int(w * h)))
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
	cvec_glVertex_Array(&context.Vertex_arrays, 0, 3)
	cvec_glBuffer(&context.Buffers, 0, 3)
	cvec_glProgram(&context.Programs, 0, 3)
	cvec_glTexture(&context.Textures, 0, 1)
	cvec_glVertex(&context.Glverts, 0, 10)
	cvec_float(&context.Vs_output.Output_buf, 0, 0)
	context.Clear_stencil = 0
	context.Clear_color = make_Color(0, 0, 0, 0)
	for {
		context.Blend_color.X = 0
		context.Blend_color.Y = 0
		context.Blend_color.Z = 0
		context.Blend_color.W = 0
		if true {
			break
		}
	}
	context.Point_size = GLfloat(1.0)
	context.Clear_depth = GLfloat(1.0)
	context.Depth_range_near = GLfloat(0.0)
	context.Depth_range_far = GLfloat(1.0)
	make_viewport_matrix(&context.Vp_mat, 0, 0, uint64(w), uint64(h), 1)
	context.Provoking_vert = GLenum(GL_LAST_VERTEX_CONVENTION)
	context.Cull_mode = GLenum(GL_BACK)
	context.Cull_face = GL_FALSE
	context.Front_face = GLenum(GL_CCW)
	context.Depth_test = GL_FALSE
	context.Fragdepth_or_discard = GL_FALSE
	context.Depth_clamp = GL_FALSE
	context.Depth_mask = GL_TRUE
	context.Blend = GL_FALSE
	context.Logic_ops = GL_FALSE
	context.Poly_offset = GL_FALSE
	context.Scissor_test = GL_FALSE
	context.Stencil_test = GL_FALSE
	context.Stencil_writemask = math.MaxUint32
	context.Stencil_writemask_back = math.MaxUint32
	context.Stencil_ref = 0
	context.Stencil_ref_back = 0
	context.Stencil_valuemask = math.MaxUint32
	context.Stencil_valuemask_back = math.MaxUint32
	context.Stencil_func = GLenum(GL_ALWAYS)
	context.Stencil_func_back = GLenum(GL_ALWAYS)
	context.Stencil_sfail = GLenum(GL_KEEP)
	context.Stencil_dpfail = GLenum(GL_KEEP)
	context.Stencil_dppass = GLenum(GL_KEEP)
	context.Stencil_sfail_back = GLenum(GL_KEEP)
	context.Stencil_dpfail_back = GLenum(GL_KEEP)
	context.Stencil_dppass_back = GLenum(GL_KEEP)
	context.Logic_func = GLenum(GL_COPY)
	context.Blend_sfactor = GLenum(GL_ONE)
	context.Blend_dfactor = GLenum(GL_ZERO)
	context.Blend_equation = GLenum(GL_FUNC_ADD)
	context.Depth_func = GLenum(GL_LESS)
	context.Line_smooth = GL_FALSE
	context.Poly_mode_front = GLenum(GL_FILL)
	context.Poly_mode_back = GLenum(GL_FILL)
	context.Point_spr_origin = GLenum(GL_UPPER_LEFT)
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
	context.Error = GLenum(GL_NO_ERROR)
	var tmp_prog glProgram = glProgram{Vertex_shader: default_vs, Fragment_shader: default_fs, Uniform: nil, Vs_output_size: GL_FALSE}
	cvec_push_glProgram(&context.Programs, tmp_prog)
	context.Cur_program = 0
	var tmp_va glVertex_Array
	init_glVertex_Array(&tmp_va)
	cvec_push_glVertex_Array(&context.Vertex_arrays, tmp_va)
	context.Cur_vertex_array = 0
	var tmp_buf glBuffer
	tmp_buf.User_owned = GL_TRUE
	tmp_buf.Deleted = GL_FALSE
	var tmp_tex glTexture
	tmp_tex.User_owned = GL_TRUE
	tmp_tex.Deleted = GL_FALSE
	tmp_tex.Format = GLenum(GL_RGBA)
	tmp_tex.Type = GLenum(GL_TEXTURE_UNBOUND)
	tmp_tex.Data = nil
	tmp_tex.W = 0
	tmp_tex.H = 0
	tmp_tex.D = 0
	cvec_push_glBuffer(&context.Buffers, tmp_buf)
	cvec_push_glTexture(&context.Textures, tmp_tex)
	return 1
}
func Free_glContext(context *GlContext) {
	var i int64
	context.Zbuf.Buf = nil
	context.Stencil_buf.Buf = nil
	if context.User_alloced_backbuf == 0 {
		context.Back_buffer.Buf = nil
	}
	for i = 0; uint64(i) < context.Buffers.Size; i++ {
		if context.Buffers.A[i].User_owned == 0 {
			stdio.Printf("freeing buffer %d\n", i)
			context.Buffers.A[i].Data = nil
		}
	}
	for i = 0; uint64(i) < context.Textures.Size; i++ {
		if (*(*glTexture)(unsafe.Add(unsafe.Pointer(context.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(i)))).User_owned == 0 {
			stdio.Printf("freeing texture %d\n", i)
			libc.Free(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(context.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(i)))).Data))
		}
	}
	cvec_free_glVertex_Array(unsafe.Pointer(&context.Vertex_arrays))
	cvec_free_glBuffer(unsafe.Pointer(&context.Buffers))
	cvec_free_glProgram(unsafe.Pointer(&context.Programs))
	cvec_free_glTexture(unsafe.Pointer(&context.Textures))
	cvec_free_glVertex(unsafe.Pointer(&context.Glverts))
	cvec_free_float(unsafe.Pointer(&context.Vs_output.Output_buf))
}
func Set_glContext(context *GlContext) {
	c = context
}

func GetString(name GLenum) *GLubyte {
	switch name {
	case GL_VENDOR:
		return (*GLubyte)(libc.CString("Robert Winkler"))
	case GL_RENDERER:
		return (*GLubyte)(libc.CString("PortableGL"))
	case GL_VERSION:
		return (*GLubyte)(libc.CString("OpenGL 3.x-ish PortableGL 0.94"))
	case GL_SHADING_LANGUAGE_VERSION:
		return (*GLubyte)(libc.CString("Go"))
	default:
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return nil
	}
}
func GetError() GLenum {
	var err GLenum = c.Error
	c.Error = GLenum(GL_NO_ERROR)
	return err
}
func GenVertexArrays(n GLsizei, arrays *GLuint) {
	a := unsafe.Slice(arrays, n)
	var tmp glVertex_Array
	init_glVertex_Array(&tmp)
	tmp.Deleted = GL_FALSE
	n--
	for i := int64(1); uint64(i) < c.Vertex_arrays.Size && n >= 0; i++ {
		if (c.Vertex_arrays.A[i]).Deleted != 0 {
			c.Vertex_arrays.A[i] = tmp
			a[n] = GLuint(int32(i))
			n--
		}
	}
	for ; n >= 0; n-- {
		cvec_push_glVertex_Array(&c.Vertex_arrays, tmp)
		a[n] = GLuint(c.Vertex_arrays.Size - 1)
	}
}
func DeleteVertexArrays(n GLsizei, arrays *GLuint) {
	a := unsafe.Slice(arrays, n)
	for i := int64(0); i < int64(n); i++ {
		if a[i] == 0 || uint64(a[i]) >= c.Vertex_arrays.Size {
			continue
		}
		if a[i] == c.Cur_vertex_array {
			//TODO check if memcpy isn't enough
			libc.MemCpy(unsafe.Pointer(&c.Vertex_arrays.A[0]), unsafe.Pointer(&c.Vertex_arrays.A[a[i]]), int(unsafe.Sizeof(glVertex_Array{})))
			c.Cur_vertex_array = 0
		}
		c.Vertex_arrays.A[a[i]].Deleted = GL_TRUE
	}
}
func GenBuffers(n GLsizei, buffers *GLuint) {
	b := unsafe.Slice(buffers, n)
	var tmp glBuffer
	tmp.User_owned = GL_TRUE
	tmp.Data = nil
	tmp.Deleted = GL_FALSE
	n--
	for i := int64(1); uint64(i) < c.Buffers.Size && n >= 0; i++ {
		if (c.Buffers.A[i]).Deleted != 0 {
			c.Buffers.A[i] = tmp
			b[n] = GLuint(int32(i))
			n--
		}
	}
	for ; n >= 0; n-- {
		cvec_push_glBuffer(&c.Buffers, tmp)
		b[n] = GLuint(c.Buffers.Size - 1)
	}
}
func DeleteBuffers(n GLsizei, buffers *GLuint) {
	b := unsafe.Slice(buffers, n)
	var type_ GLenum
	for i := int64(0); i < int64(n); i++ {
		if b[i] == 0 || uint64(b[i]) >= c.Buffers.Size {
			continue
		}
		type_ = c.Buffers.A[b[i]].Type
		if b[i] == c.Bound_buffers[type_] {
			c.Bound_buffers[type_] = 0
		}
		if c.Buffers.A[b[i]].User_owned == 0 {
			c.Buffers.A[b[i]].Data = nil
		}
		c.Buffers.A[b[i]].Deleted = GL_TRUE
	}
}
func GenTextures(n GLsizei, textures *GLuint) {
	t := unsafe.Slice(textures, n)
	var tmp glTexture
	tmp.Mag_filter = GLenum(GL_LINEAR)
	tmp.Min_filter = GLenum(GL_LINEAR)
	tmp.Wrap_s = GLenum(GL_REPEAT)
	tmp.Wrap_t = GLenum(GL_REPEAT)
	tmp.Data = nil
	tmp.Deleted = GL_FALSE
	tmp.User_owned = GL_TRUE
	tmp.Format = GLenum(GL_RGBA)
	tmp.Type = GLenum(GL_TEXTURE_UNBOUND)
	tmp.W = 0
	tmp.H = 0
	tmp.D = 0
	n--
	for i := int64(0); uint64(i) < c.Textures.Size && n >= 0; i++ {
		if (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(i)))).Deleted != 0 {
			*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(i))) = tmp
			t[n] = GLuint(int32(i))
			n--
		}
	}
	for ; n >= 0; n-- {
		cvec_push_glTexture(&c.Textures, tmp)
		t[n] = GLuint(c.Textures.Size - 1)
	}
}
func DeleteTextures(n GLsizei, textures *GLuint) {
	t := unsafe.Slice(textures, n)
	var type_ GLenum
	for i := int64(0); i < int64(n); i++ {
		if t[i] == 0 || uint64(t[i]) >= c.Textures.Size {
			continue
		}
		type_ = (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(t[i])))).Type
		if t[i] == c.Bound_textures[type_] {
			c.Bound_textures[type_] = 0
		}
		if (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(t[i])))).User_owned == 0 {
			libc.Free(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(t[i])))).Data))
			(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(t[i])))).Data = nil
		}
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(t[i])))).Deleted = GL_TRUE
	}
}
func BindVertexArray(array GLuint) {
	if uint64(array) < c.Vertex_arrays.Size && c.Vertex_arrays.A[array].Deleted == GL_FALSE {
		c.Cur_vertex_array = array
	} else if c.Error == 0 {
		c.Error = GLenum(GL_INVALID_OPERATION)
	}
}
func BindBuffer(target GLenum, buffer GLuint) {
	if target != GLenum(GL_ARRAY_BUFFER) && target != GLenum(GL_ELEMENT_ARRAY_BUFFER) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	target -= GLenum(GL_ARRAY_BUFFER)
	if uint64(buffer) < c.Buffers.Size && c.Buffers.A[buffer].Deleted == GL_FALSE {
		c.Bound_buffers[target] = buffer
		c.Buffers.A[buffer].Type = target
	} else if c.Error == 0 {
		c.Error = GLenum(GL_INVALID_OPERATION)
	}
}
func BufferData(target GLenum, size GLsizei, data unsafe.Pointer, usage GLenum) {
	if target != GLenum(GL_ARRAY_BUFFER) && target != GLenum(GL_ELEMENT_ARRAY_BUFFER) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	target -= GLenum(GL_ARRAY_BUFFER)
	if c.Bound_buffers[target] == 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_OPERATION)
		}
		return
	}
	c.Buffers.A[c.Bound_buffers[target]].Data = nil
	if (c.Buffers.A[c.Bound_buffers[target]]).Data = make([]u8, size); c.Buffers.A[c.Bound_buffers[target]].Data == nil {
		if c.Error == 0 {
			c.Error = GLenum(GL_OUT_OF_MEMORY)
		}
		return
	}
	if data != nil {
		copy(c.Buffers.A[c.Bound_buffers[target]].Data, unsafe.Slice((*u8)(data), size))
	}
	c.Buffers.A[c.Bound_buffers[target]].User_owned = GL_FALSE
	(c.Buffers.A[c.Bound_buffers[target]]).Size = size
	if target == GLenum(GL_ELEMENT_ARRAY_BUFFER-GL_ARRAY_BUFFER) {
		(c.Vertex_arrays.A[c.Cur_vertex_array]).Element_buffer = c.Bound_buffers[target]
	}
}
func BufferSubData(target GLenum, offset GLsizei, size GLsizei, data unsafe.Pointer) {
	if target != GLenum(GL_ARRAY_BUFFER) && target != GLenum(GL_ELEMENT_ARRAY_BUFFER) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	target -= GLenum(GL_ARRAY_BUFFER)
	if c.Bound_buffers[target] == 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_OPERATION)
		}
		return
	}
	if offset+size > (c.Buffers.A[c.Bound_buffers[target]]).Size {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	libc.MemCpy(unsafe.Add(unsafe.Pointer(&(c.Buffers.A[c.Bound_buffers[target]]).Data[0]), offset), data, int(size))
}
func BindTexture(target GLenum, texture GLuint) {
	if target < GLenum(GL_TEXTURE_1D) || target >= GLenum(GL_NUM_TEXTURE_TYPES) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	target -= GLenum(GL_TEXTURE_UNBOUND + 1)
	if uint64(texture) < c.Textures.Size && (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(texture)))).Deleted == 0 {
		if (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(texture)))).Type == GLenum(GL_TEXTURE_UNBOUND) {
			c.Bound_textures[target] = texture
			(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(texture)))).Type = target
		} else if (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(texture)))).Type == target {
			c.Bound_textures[target] = texture
		} else if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_OPERATION)
		}
	} else if c.Error == 0 {
		c.Error = GLenum(GL_INVALID_VALUE)
	}
}
func TexParameteri(target GLenum, pname GLenum, param GLint) {
	if target != GLenum(GL_TEXTURE_1D) && target != GLenum(GL_TEXTURE_2D) && target != GLenum(GL_TEXTURE_3D) && target != GLenum(GL_TEXTURE_2D_ARRAY) && target != GLenum(GL_TEXTURE_RECTANGLE) && target != GLenum(GL_TEXTURE_CUBE_MAP) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	target -= GLenum(GL_TEXTURE_UNBOUND + 1)
	if pname != GLenum(GL_TEXTURE_MIN_FILTER) && pname != GLenum(GL_TEXTURE_MAG_FILTER) && pname != GLenum(GL_TEXTURE_WRAP_S) && pname != GLenum(GL_TEXTURE_WRAP_T) && pname != GLenum(GL_TEXTURE_WRAP_R) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if pname == GLenum(GL_TEXTURE_MIN_FILTER) {
		if int64(param) != GL_NEAREST && int64(param) != GL_LINEAR && int64(param) != GL_NEAREST_MIPMAP_NEAREST && int64(param) != GL_NEAREST_MIPMAP_LINEAR && int64(param) != GL_LINEAR_MIPMAP_NEAREST && int64(param) != GL_LINEAR_MIPMAP_LINEAR {
			if c.Error == 0 {
				c.Error = GLenum(GL_INVALID_ENUM)
			}
			return
		}
		if int64(param) == GL_NEAREST_MIPMAP_NEAREST || int64(param) == GL_NEAREST_MIPMAP_LINEAR {
			param = GLint(GL_NEAREST)
		}
		if int64(param) == GL_LINEAR_MIPMAP_NEAREST || int64(param) == GL_LINEAR_MIPMAP_LINEAR {
			param = GLint(GL_LINEAR)
		}
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(c.Bound_textures[target])))).Min_filter = GLenum(param)
	} else if pname == GLenum(GL_TEXTURE_MAG_FILTER) {
		if int64(param) != GL_NEAREST && int64(param) != GL_LINEAR {
			if c.Error == 0 {
				c.Error = GLenum(GL_INVALID_ENUM)
			}
			return
		}
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(c.Bound_textures[target])))).Mag_filter = GLenum(param)
	} else if pname == GLenum(GL_TEXTURE_WRAP_S) {
		if int64(param) != GL_REPEAT && int64(param) != GL_CLAMP_TO_EDGE && int64(param) != GL_CLAMP_TO_BORDER && int64(param) != GL_MIRRORED_REPEAT {
			if c.Error == 0 {
				c.Error = GLenum(GL_INVALID_ENUM)
			}
			return
		}
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(c.Bound_textures[target])))).Wrap_s = GLenum(param)
	} else if pname == GLenum(GL_TEXTURE_WRAP_T) {
		if int64(param) != GL_REPEAT && int64(param) != GL_CLAMP_TO_EDGE && int64(param) != GL_CLAMP_TO_BORDER && int64(param) != GL_MIRRORED_REPEAT {
			if c.Error == 0 {
				c.Error = GLenum(GL_INVALID_ENUM)
			}
			return
		}
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(c.Bound_textures[target])))).Wrap_t = GLenum(param)
	} else if pname == GLenum(GL_TEXTURE_WRAP_R) {
		if int64(param) != GL_REPEAT && int64(param) != GL_CLAMP_TO_EDGE && int64(param) != GL_CLAMP_TO_BORDER && int64(param) != GL_MIRRORED_REPEAT {
			if c.Error == 0 {
				c.Error = GLenum(GL_INVALID_ENUM)
			}
			return
		}
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(c.Bound_textures[target])))).Wrap_r = GLenum(param)
	}
}
func PixelStorei(pname GLenum, param GLint) {
	if pname != GLenum(GL_UNPACK_ALIGNMENT) && pname != GLenum(GL_PACK_ALIGNMENT) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if param != 1 && param != 2 && param != 4 && param != 8 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if pname == GLenum(GL_UNPACK_ALIGNMENT) {
		c.Unpack_alignment = param
	} else if pname == GLenum(GL_PACK_ALIGNMENT) {
		c.Pack_alignment = param
	}
}
func GenerateMipmap(target GLenum) {
	if target != GLenum(GL_TEXTURE_1D) && target != GLenum(GL_TEXTURE_2D) && target != GLenum(GL_TEXTURE_3D) && target != GLenum(GL_TEXTURE_CUBE_MAP) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
}
func TexImage1D(target GLenum, level GLint, internalFormat GLint, width GLsizei, border GLint, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(GL_TEXTURE_1D) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if border != 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	var cur_tex int64 = int64(c.Bound_textures[target-GLenum(GL_TEXTURE_UNBOUND)-1])
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W = uint64(width)
	if type_ != GLenum(GL_UNSIGNED_BYTE) {
		return
	}
	var components int64
	if format == GLenum(GL_RED) {
		components = 1
	} else if format == GLenum(GL_RG) {
		components = 2
	} else if format == GLenum(GL_RGB) || format == GLenum(GL_BGR) {
		components = 3
	} else if format == GLenum(GL_RGBA) || format == GLenum(GL_BGRA) {
		components = 4
	} else {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	libc.Free(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
	if (func() *u8 {
		p := &(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data = (*u8)(libc.Malloc(int(int64(width) * components)))
		return *p
	}()) == nil {
		if c.Error == 0 {
			c.Error = GLenum(GL_OUT_OF_MEMORY)
		}
		return
	}
	var texdata *U32 = (*U32)(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
	if data != nil {
		libc.MemCpy(unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(U32(0))*0))), data, int(uintptr(width)*unsafe.Sizeof(U32(0))))
	}
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).User_owned = GL_FALSE
}
func TexImage2D(target GLenum, level GLint, internalFormat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(GL_TEXTURE_2D) && target != GLenum(GL_TEXTURE_RECTANGLE) && target != GLenum(GL_TEXTURE_CUBE_MAP_POSITIVE_X) && target != GLenum(GL_TEXTURE_CUBE_MAP_NEGATIVE_X) && target != GLenum(GL_TEXTURE_CUBE_MAP_POSITIVE_Y) && target != GLenum(GL_TEXTURE_CUBE_MAP_NEGATIVE_Y) && target != GLenum(GL_TEXTURE_CUBE_MAP_POSITIVE_Z) && target != GLenum(GL_TEXTURE_CUBE_MAP_NEGATIVE_Z) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if border != 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if type_ != GLenum(GL_UNSIGNED_BYTE) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	var components int64
	if format == GLenum(GL_RED) {
		components = 1
	} else if format == GLenum(GL_RG) {
		components = 2
	} else if format == GLenum(GL_RGB) || format == GLenum(GL_BGR) {
		components = 3
	} else if format == GLenum(GL_RGBA) || format == GLenum(GL_BGRA) {
		components = 4
	} else {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
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
	if target == GLenum(GL_TEXTURE_2D) || target == GLenum(GL_TEXTURE_RECTANGLE) {
		cur_tex = int64(c.Bound_textures[target-GLenum(GL_TEXTURE_UNBOUND)-1])
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W = uint64(width)
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).H = uint64(height)
		libc.Free(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
		if (func() *u8 {
			p := &(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data
			(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data = (*u8)(libc.Malloc(int(int64(height) * byte_width)))
			return *p
		}()) == nil {
			if c.Error == 0 {
				c.Error = GLenum(GL_OUT_OF_MEMORY)
			}
			return
		}
		if data != nil {
			if padding_needed == 0 {
				libc.MemCpy(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data), data, int(int64(height)*byte_width))
			} else {
				for i := int64(0); i < int64(height); i++ {
					libc.MemCpy(unsafe.Add(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data), i*byte_width), unsafe.Add(unsafe.Pointer((*u8)(data)), i*padded_row_len), int(byte_width))
				}
			}
		}
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).User_owned = GL_FALSE
	} else {
		cur_tex = int64(c.Bound_textures[GL_TEXTURE_CUBE_MAP-GL_TEXTURE_UNBOUND-1])
		if (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W == 0 {
			libc.Free(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
		}
		if width != height {
			if c.Error == 0 {
				c.Error = GLenum(GL_INVALID_VALUE)
			}
			return
		}
		var mem_size int64 = int64(width*height*6) * components
		if (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W == 0 {
			(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W = uint64(width)
			(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).H = uint64(width)
			if (func() *u8 {
				p := &(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data
				(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data = (*u8)(libc.Malloc(int(mem_size)))
				return *p
			}()) == nil {
				if c.Error == 0 {
					c.Error = GLenum(GL_OUT_OF_MEMORY)
				}
				return
			}
		} else if (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W != uint64(width) {
			if c.Error == 0 {
				c.Error = GLenum(GL_INVALID_VALUE)
			}
			return
		}
		target -= GLenum(GL_TEXTURE_CUBE_MAP_POSITIVE_X)
		var p int64 = int64(height) * byte_width
		var texdata *u8 = (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data
		if data != nil {
			if padding_needed == 0 {
				libc.MemCpy(unsafe.Add(unsafe.Pointer(texdata), target*GLenum(p)), data, int(int64(height)*byte_width))
			} else {
				for i := int64(0); i < int64(height); i++ {
					libc.MemCpy(unsafe.Add(unsafe.Pointer(texdata), target*GLenum(p)+GLenum(i*byte_width)), unsafe.Add(unsafe.Pointer((*u8)(data)), i*padded_row_len), int(byte_width))
				}
			}
		}
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).User_owned = GL_FALSE
	}
}
func TexImage3D(target GLenum, level GLint, internalFormat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(GL_TEXTURE_3D) && target != GLenum(GL_TEXTURE_2D_ARRAY) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if border != 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	var cur_tex int64 = int64(c.Bound_textures[target-GLenum(GL_TEXTURE_UNBOUND)-1])
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W = uint64(width)
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).H = uint64(height)
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).D = uint64(depth)
	if type_ != GLenum(GL_UNSIGNED_BYTE) {
		return
	}
	var components int64
	if format == GLenum(GL_RED) {
		components = 1
	} else if format == GLenum(GL_RG) {
		components = 2
	} else if format == GLenum(GL_RGB) || format == GLenum(GL_BGR) {
		components = 3
	} else if format == GLenum(GL_RGBA) || format == GLenum(GL_BGRA) {
		components = 4
	} else {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
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
	libc.Free(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
	if (func() *u8 {
		p := &(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data = (*u8)(libc.Malloc(int(int64(width*height*depth) * components)))
		return *p
	}()) == nil {
		if c.Error == 0 {
			c.Error = GLenum(GL_OUT_OF_MEMORY)
		}
		return
	}
	var texdata *U32 = (*U32)(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
	if data != nil {
		if padding_needed == 0 {
			libc.MemCpy(unsafe.Pointer(texdata), data, int(uintptr(width*height*depth)*unsafe.Sizeof(U32(0))))
		} else {
			for i := int64(0); i < int64(height*depth); i++ {
				libc.MemCpy(unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(U32(0))*uintptr(i*byte_width)))), unsafe.Add(unsafe.Pointer((*u8)(data)), i*padded_row_len), int(byte_width))
			}
		}
	}
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).User_owned = GL_FALSE
}
func TexSubImage1D(target GLenum, level GLint, xoffset GLint, width GLsizei, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(GL_TEXTURE_1D) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	var cur_tex int64 = int64(c.Bound_textures[target-GLenum(GL_TEXTURE_UNBOUND)-1])
	if type_ != GLenum(GL_UNSIGNED_BYTE) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if format != GLenum(GL_RGBA) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if xoffset < 0 || uint64(xoffset+GLint(width)) > (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	var texdata *U32 = (*U32)(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
	libc.MemCpy(unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(U32(0))*uintptr(xoffset)))), data, int(uintptr(width)*unsafe.Sizeof(U32(0))))
}
func TexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(GL_TEXTURE_2D) && target != GLenum(GL_TEXTURE_CUBE_MAP_POSITIVE_X) && target != GLenum(GL_TEXTURE_CUBE_MAP_NEGATIVE_X) && target != GLenum(GL_TEXTURE_CUBE_MAP_POSITIVE_Y) && target != GLenum(GL_TEXTURE_CUBE_MAP_NEGATIVE_Y) && target != GLenum(GL_TEXTURE_CUBE_MAP_POSITIVE_Z) && target != GLenum(GL_TEXTURE_CUBE_MAP_NEGATIVE_Z) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if type_ != GLenum(GL_UNSIGNED_BYTE) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if format != GLenum(GL_RGBA) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	var cur_tex int64
	var d *U32 = (*U32)(data)
	if target == GLenum(GL_TEXTURE_2D) {
		cur_tex = int64(c.Bound_textures[target-GLenum(GL_TEXTURE_UNBOUND)-1])
		var texdata *U32 = (*U32)(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
		if xoffset < 0 || uint64(xoffset+GLint(width)) > (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W || yoffset < 0 || uint64(yoffset+GLint(height)) > (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).H {
			if c.Error == 0 {
				c.Error = GLenum(GL_INVALID_VALUE)
			}
			return
		}
		var w int64 = int64((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W)
		for i := int64(0); i < int64(height); i++ {
			libc.MemCpy(unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(U32(0))*uintptr((int64(yoffset)+i)*w+int64(xoffset))))), unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(d), unsafe.Sizeof(U32(0))*uintptr(i*int64(width))))), int(uintptr(width)*unsafe.Sizeof(U32(0))))
		}
	} else {
		cur_tex = int64(c.Bound_textures[GL_TEXTURE_CUBE_MAP-GL_TEXTURE_UNBOUND-1])
		var texdata *U32 = (*U32)(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
		var w int64 = int64((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W)
		target -= GLenum(GL_TEXTURE_CUBE_MAP_POSITIVE_X)
		var p int64 = w * w
		for i := int64(0); i < int64(height); i++ {
			libc.MemCpy(unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(U32(0))*uintptr(p*int64(target)+(int64(yoffset)+i)*w+int64(xoffset))))), unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(d), unsafe.Sizeof(U32(0))*uintptr(i*int64(width))))), int(uintptr(width)*unsafe.Sizeof(U32(0))))
		}
	}
}
func TexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, type_ GLenum, data unsafe.Pointer) {
	if target != GLenum(GL_TEXTURE_3D) && target != GLenum(GL_TEXTURE_2D_ARRAY) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	var cur_tex int64 = int64(c.Bound_textures[target-GLenum(GL_TEXTURE_UNBOUND)-1])
	if type_ != GLenum(GL_UNSIGNED_BYTE) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if format != GLenum(GL_RGBA) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if xoffset < 0 || uint64(xoffset+GLint(width)) > (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W || yoffset < 0 || uint64(yoffset+GLint(height)) > (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).H || zoffset < 0 || uint64(zoffset+GLint(depth)) > (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).D {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	var w int64 = int64((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W)
	var h int64 = int64((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).H)
	var p int64 = w * h
	var d *U32 = (*U32)(data)
	var texdata *U32 = (*U32)(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
	for j := int64(0); j < int64(depth); j++ {
		for i := int64(0); i < int64(height); i++ {
			libc.MemCpy(unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(texdata), unsafe.Sizeof(U32(0))*uintptr((int64(zoffset)+j)*p+(int64(yoffset)+i)*w+int64(xoffset))))), unsafe.Pointer((*U32)(unsafe.Add(unsafe.Pointer(d), unsafe.Sizeof(U32(0))*uintptr(j*int64(width)*int64(height)+i*int64(width))))), int(uintptr(width)*unsafe.Sizeof(U32(0))))
		}
	}
}
func VertexAttribPointer(index GLuint, size GLint, type_ GLenum, normalized GLboolean, stride GLsizei, offset GLsizei) {
	if index >= GL_MAX_VERTEX_ATTRIBS || size < 1 || size > 4 || c.Bound_buffers[GL_ARRAY_BUFFER-GL_ARRAY_BUFFER] == 0 && offset != 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_OPERATION)
		}
		return
	}
	if type_ != GLenum(GL_FLOAT) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	var v *glVertex_Attrib = &((c.Vertex_arrays.A[c.Cur_vertex_array]).Vertex_attribs[index])
	v.Size = size
	v.Type = type_
	if stride != 0 {
		v.Stride = stride
	} else {
		v.Stride = GLsizei(uint32(uintptr(size) * unsafe.Sizeof(GLfloat(0))))
	}
	v.Offset = offset
	v.Normalized = normalized
	v.Buf = uint64(c.Bound_buffers[GL_ARRAY_BUFFER-GL_ARRAY_BUFFER])
}
func EnableVertexAttribArray(index GLuint) {
	(c.Vertex_arrays.A[c.Cur_vertex_array]).Vertex_attribs[index].Enabled = GL_TRUE
}
func DisableVertexAttribArray(index GLuint) {
	(c.Vertex_arrays.A[c.Cur_vertex_array]).Vertex_attribs[index].Enabled = GL_FALSE
}
func VertexAttribDivisor(index GLuint, divisor GLuint) {
	if index >= GL_MAX_VERTEX_ATTRIBS {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	(c.Vertex_arrays.A[c.Cur_vertex_array]).Vertex_attribs[index].Divisor = divisor
}
func get_vertex_attrib_array(v *glVertex_Attrib, i GLsizei) Vec4 {
	var (
		buf_pos *u8 = (*u8)(unsafe.Add(unsafe.Pointer((*u8)(unsafe.Add(unsafe.Pointer(&(c.Buffers.A[v.Buf]).Data[0]), v.Offset))), v.Stride*i))
		tmpvec4 Vec4
	)
	libc.MemCpy(unsafe.Pointer(&tmpvec4), unsafe.Pointer(buf_pos), int(uintptr(v.Size)*unsafe.Sizeof(float32(0))))
	return tmpvec4
}
func DrawArrays(mode GLenum, first GLint, count GLsizei) {
	if mode < GLenum(GL_POINTS) || mode > GLenum(GL_TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if count < 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if count == 0 {
		return
	}
	run_pipeline(mode, first, count, 0, 0, GL_FALSE)
}
func DrawElements(mode GLenum, count GLsizei, type_ GLenum, offset GLsizei) {
	if mode < GLenum(GL_POINTS) || mode > GLenum(GL_TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if type_ != GLenum(GL_UNSIGNED_BYTE) && type_ != GLenum(GL_UNSIGNED_SHORT) && type_ != GLenum(GL_UNSIGNED_INT) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if count < 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if count == 0 {
		return
	}
	(c.Buffers.A[(c.Vertex_arrays.A[c.Cur_vertex_array]).Element_buffer]).Type = type_
	run_pipeline(mode, GLint(offset), count, 0, 0, GL_TRUE)
}
func DrawArraysInstanced(mode GLenum, first GLint, count GLsizei, instancecount GLsizei) {
	if mode < GLenum(GL_POINTS) || mode > GLenum(GL_TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if count < 0 || instancecount < 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if count == 0 || instancecount == 0 {
		return
	}
	for instance := uint64(0); instance < uint64(instancecount); instance++ {
		run_pipeline(mode, first, count, GLsizei(uint32(instance)), 0, GL_FALSE)
	}
}
func DrawArraysInstancedBaseInstance(mode GLenum, first GLint, count GLsizei, instancecount GLsizei, baseinstance GLuint) {
	if mode < GLenum(GL_POINTS) || mode > GLenum(GL_TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if count < 0 || instancecount < 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if count == 0 || instancecount == 0 {
		return
	}
	for instance := uint64(0); instance < uint64(instancecount); instance++ {
		run_pipeline(mode, first, count, GLsizei(uint32(instance)), baseinstance, GL_FALSE)
	}
}
func DrawElementsInstanced(mode GLenum, count GLsizei, type_ GLenum, offset GLsizei, instancecount GLsizei) {
	if mode < GLenum(GL_POINTS) || mode > GLenum(GL_TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if type_ != GLenum(GL_UNSIGNED_BYTE) && type_ != GLenum(GL_UNSIGNED_SHORT) && type_ != GLenum(GL_UNSIGNED_INT) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if count < 0 || instancecount < 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if count == 0 || instancecount == 0 {
		return
	}
	(c.Buffers.A[(c.Vertex_arrays.A[c.Cur_vertex_array]).Element_buffer]).Type = type_
	for instance := uint64(0); instance < uint64(instancecount); instance++ {
		run_pipeline(mode, GLint(offset), count, GLsizei(uint32(instance)), 0, GL_TRUE)
	}
}
func DrawElementsInstancedBaseInstance(mode GLenum, count GLsizei, type_ GLenum, offset GLsizei, instancecount GLsizei, baseinstance GLuint) {
	if mode < GLenum(GL_POINTS) || mode > GLenum(GL_TRIANGLE_FAN) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if type_ != GLenum(GL_UNSIGNED_BYTE) && type_ != GLenum(GL_UNSIGNED_SHORT) && type_ != GLenum(GL_UNSIGNED_INT) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if count < 0 || instancecount < 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if count == 0 || instancecount == 0 {
		return
	}
	c.Buffers.A[(c.Vertex_arrays.A[c.Cur_vertex_array]).Element_buffer].Type = type_
	for instance := uint64(0); instance < uint64(instancecount); instance++ {
		run_pipeline(mode, GLint(offset), count, GLsizei(uint32(instance)), baseinstance, GL_TRUE)
	}
}
func Viewport(x int64, y int64, width GLsizei, height GLsizei) {
	if width < 0 || height < 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
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
	if func_ < GLenum(GL_LESS) || func_ > GLenum(GL_NEVER) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
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
	if (mask & GLbitfield(GL_COLOR_BUFFER_BIT|GL_DEPTH_BUFFER_BIT|GL_STENCIL_BUFFER_BIT)) == 0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		stdio.Printf("failed to clear\n")
		return
	}
	var col Color = c.Clear_color
	if mask&GLbitfield(GL_COLOR_BUFFER_BIT) != 0 {
		if c.Scissor_test == 0 {
			for i := int64(0); uint64(i) < c.Back_buffer.W*c.Back_buffer.H; i++ {
				*(*U32)(unsafe.Add(unsafe.Pointer((*U32)(unsafe.Pointer(&c.Back_buffer.Buf[0]))), unsafe.Sizeof(U32(0))*uintptr(i))) = U32(int32(int64(col.A)<<c.Ashift | int64(col.R)<<c.Rshift | int64(col.G)<<c.Gshift | int64(col.B)<<c.Bshift))
			}
		} else {
			for y := int64(int64(c.Scissor_ly)); y < int64(c.Scissor_uy); y++ {
				for x := int64(int64(c.Scissor_lx)); x < int64(c.Scissor_ux); x++ {
					*(*U32)(unsafe.Add(unsafe.Pointer((*U32)(unsafe.Pointer(&c.Back_buffer.Lastrow[0]))), unsafe.Sizeof(U32(0))*uintptr(uint64(-y)*c.Back_buffer.W+uint64(x)))) = U32(int32(int64(col.A)<<c.Ashift | int64(col.R)<<c.Rshift | int64(col.G)<<c.Gshift | int64(col.B)<<c.Bshift))
				}
			}
		}
	}
	if mask&GLbitfield(GL_DEPTH_BUFFER_BIT) != 0 {
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
	if mask&GLbitfield(GL_STENCIL_BUFFER_BIT) != 0 {
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
	case GL_CULL_FACE:
		c.Cull_face = GL_TRUE
	case GL_DEPTH_TEST:
		c.Depth_test = GL_TRUE
	case GL_DEPTH_CLAMP:
		c.Depth_clamp = GL_TRUE
	case GL_LINE_SMOOTH:
	case GL_BLEND:
		c.Blend = GL_TRUE
	case GL_COLOR_LOGIC_OP:
		c.Logic_ops = GL_TRUE
	case GL_POLYGON_OFFSET_FILL:
		c.Poly_offset = GL_TRUE
	case GL_SCISSOR_TEST:
		c.Scissor_test = GL_TRUE
	case GL_STENCIL_TEST:
		c.Stencil_test = GL_TRUE
	default:
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
	}
}
func Disable(cap_ GLenum) {
	switch cap_ {
	case GL_CULL_FACE:
		c.Cull_face = GL_FALSE
	case GL_DEPTH_TEST:
		c.Depth_test = GL_FALSE
	case GL_DEPTH_CLAMP:
		c.Depth_clamp = GL_FALSE
	case GL_LINE_SMOOTH:
		c.Line_smooth = GL_FALSE
	case GL_BLEND:
		c.Blend = GL_FALSE
	case GL_COLOR_LOGIC_OP:
		c.Logic_ops = GL_FALSE
	case GL_POLYGON_OFFSET_FILL:
		c.Poly_offset = GL_FALSE
	case GL_SCISSOR_TEST:
		c.Scissor_test = GL_FALSE
	case GL_STENCIL_TEST:
		c.Stencil_test = GL_FALSE
	default:
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
	}
}
func IsEnabled(cap_ GLenum) GLboolean {
	switch cap_ {
	case GL_DEPTH_TEST:
		return c.Depth_test
	case GL_LINE_SMOOTH:
		return c.Line_smooth
	case GL_CULL_FACE:
		return c.Cull_face
	case GL_DEPTH_CLAMP:
		return c.Depth_clamp
	case GL_BLEND:
		return c.Blend
	case GL_COLOR_LOGIC_OP:
		return c.Logic_ops
	case GL_POLYGON_OFFSET_FILL:
		return c.Poly_offset
	case GL_SCISSOR_TEST:
		return c.Scissor_test
	case GL_STENCIL_TEST:
		return c.Stencil_test
	default:
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
	}
	return GL_FALSE
}
func GetBooleanv(pname GLenum, params *GLboolean) {
	switch pname {
	case GL_DEPTH_TEST:
		*params = c.Depth_test
	case GL_LINE_SMOOTH:
		*params = c.Line_smooth
	case GL_CULL_FACE:
		*params = c.Cull_face
	case GL_DEPTH_CLAMP:
		*params = c.Depth_clamp
	case GL_BLEND:
		*params = c.Blend
	case GL_COLOR_LOGIC_OP:
		*params = c.Logic_ops
	case GL_POLYGON_OFFSET_FILL:
		*params = c.Poly_offset
	case GL_SCISSOR_TEST:
		*params = c.Scissor_test
	case GL_STENCIL_TEST:
		*params = c.Stencil_test
	default:
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
	}
}
func GetFloatv(pname GLenum, params *GLfloat) {
	switch pname {
	case GL_POLYGON_OFFSET_FACTOR:
		*params = c.Poly_factor
	case GL_POLYGON_OFFSET_UNITS:
		*params = c.Poly_units
	case GL_POINT_SIZE:
		*params = c.Point_size
	case GL_DEPTH_CLEAR_VALUE:
		*params = c.Clear_depth
	case GL_DEPTH_RANGE:
		*(*GLfloat)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLfloat(0))*0)) = c.Depth_range_near
		*(*GLfloat)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLfloat(0))*1)) = c.Depth_range_near
	default:
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
	}
}
func GetIntegerv(pname GLenum, params *GLint) {
	switch pname {
	case GL_STENCIL_WRITE_MASK:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_writemask)
	case GL_STENCIL_REF:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = c.Stencil_ref
	case GL_STENCIL_VALUE_MASK:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_valuemask)
	case GL_STENCIL_FUNC:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_func)
	case GL_STENCIL_FAIL:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_sfail)
	case GL_STENCIL_PASS_DEPTH_FAIL:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_dpfail)
	case GL_STENCIL_PASS_DEPTH_PASS:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_dppass)
	case GL_STENCIL_BACK_WRITE_MASK:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_writemask_back)
	case GL_STENCIL_BACK_REF:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = c.Stencil_ref_back
	case GL_STENCIL_BACK_VALUE_MASK:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_valuemask_back)
	case GL_STENCIL_BACK_FUNC:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_func_back)
	case GL_STENCIL_BACK_FAIL:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_sfail_back)
	case GL_STENCIL_BACK_PASS_DEPTH_FAIL:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_dpfail_back)
	case GL_STENCIL_BACK_PASS_DEPTH_PASS:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Stencil_dppass_back)
	case GL_LOGIC_OP_MODE:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Logic_func)
	case GL_BLEND_SRC_RGB:
		fallthrough
	case GL_BLEND_SRC_ALPHA:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Blend_sfactor)
	case GL_BLEND_DST_RGB:
		fallthrough
	case GL_BLEND_DST_ALPHA:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Blend_dfactor)
	case GL_BLEND_EQUATION_RGB:
		fallthrough
	case GL_BLEND_EQUATION_ALPHA:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Blend_equation)
	case GL_CULL_FACE_MODE:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Cull_mode)
	case GL_FRONT_FACE:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Front_face)
	case GL_DEPTH_FUNC:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Depth_func)
	case GL_POINT_SPRITE_COORD_ORIGIN:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Point_spr_origin)
		fallthrough
	case GL_PROVOKING_VERTEX:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Provoking_vert)
	case GL_POLYGON_MODE:
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*0)) = GLint(c.Poly_mode_front)
		*(*GLint)(unsafe.Add(unsafe.Pointer(params), unsafe.Sizeof(GLint(0))*1)) = GLint(c.Poly_mode_back)
	default:
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
	}
}
func CullFace(mode GLenum) {
	if mode != GLenum(GL_FRONT) && mode != GLenum(GL_BACK) && mode != GLenum(GL_FRONT_AND_BACK) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	c.Cull_mode = mode
}
func FrontFace(mode GLenum) {
	if mode != GLenum(GL_CCW) && mode != GLenum(GL_CW) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	c.Front_face = mode
}
func PolygonMode(face GLenum, mode GLenum) {
	if face != GLenum(GL_FRONT) && face != GLenum(GL_BACK) && face != GLenum(GL_FRONT_AND_BACK) || mode != GLenum(GL_POINT) && mode != GLenum(GL_LINE) && mode != GLenum(GL_FILL) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if mode == GLenum(GL_POINT) {
		if face == GLenum(GL_FRONT) {
			c.Poly_mode_front = mode
			c.Draw_triangle_front = draw_triangle_point
		} else if face == GLenum(GL_BACK) {
			c.Poly_mode_back = mode
			c.Draw_triangle_back = draw_triangle_point
		} else {
			c.Poly_mode_front = mode
			c.Poly_mode_back = mode
			c.Draw_triangle_front = draw_triangle_point
			c.Draw_triangle_back = draw_triangle_point
		}
	} else if mode == GLenum(GL_LINE) {
		if face == GLenum(GL_FRONT) {
			c.Poly_mode_front = mode
			c.Draw_triangle_front = draw_triangle_line
		} else if face == GLenum(GL_BACK) {
			c.Poly_mode_back = mode
			c.Draw_triangle_back = draw_triangle_line
		} else {
			c.Poly_mode_front = mode
			c.Poly_mode_back = mode
			c.Draw_triangle_front = draw_triangle_line
			c.Draw_triangle_back = draw_triangle_line
		}
	} else {
		if face == GLenum(GL_FRONT) {
			c.Poly_mode_front = mode
			c.Draw_triangle_front = draw_triangle_fill
		} else if face == GLenum(GL_BACK) {
			c.Poly_mode_back = mode
			c.Draw_triangle_back = draw_triangle_fill
		} else {
			c.Poly_mode_front = mode
			c.Poly_mode_back = mode
			c.Draw_triangle_front = draw_triangle_fill

			c.Draw_triangle_back = draw_triangle_fill
		}
	}
}
func PointSize(size GLfloat) {
	if float64(size) <= 0.0 {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	c.Point_size = size
}
func PointParameteri(pname GLenum, param GLint) {
	if pname != GLenum(GL_POINT_SPRITE_COORD_ORIGIN) || int64(param) != GL_LOWER_LEFT && int64(param) != GL_UPPER_LEFT {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	c.Point_spr_origin = GLenum(param)
}
func ProvokingVertex(provokeMode GLenum) {
	if provokeMode != GLenum(GL_FIRST_VERTEX_CONVENTION) && provokeMode != GLenum(GL_LAST_VERTEX_CONVENTION) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	c.Provoking_vert = provokeMode
}

func DeleteProgram(program GLuint) {
	if program == 0 {
		return
	}
	if uint64(program) >= c.Programs.Size {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	c.Programs.A[program].Deleted = GL_TRUE
}
func UseProgram(program GLuint) {
	if uint64(program) >= c.Programs.Size {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	c.Vs_output.Size = c.Programs.A[program].Vs_output_size
	cvec_reserve_float(&c.Vs_output.Output_buf, uint64(c.Vs_output.Size*MAX_VERTICES))
	c.Vs_output.Interpolation = c.Programs.A[program].Interpolation[:]
	c.Fragdepth_or_discard = c.Programs.A[program].Fragdepth_or_discard
	c.Cur_program = program
}

func BlendFunc(sfactor GLenum, dfactor GLenum) {
	if sfactor < GLenum(GL_ZERO) || sfactor >= GLenum(NUM_BLEND_FUNCS) || dfactor < GLenum(GL_ZERO) || dfactor >= GLenum(NUM_BLEND_FUNCS) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	c.Blend_sfactor = sfactor
	c.Blend_dfactor = dfactor
}
func BlendEquation(mode GLenum) {
	if mode < GLenum(GL_FUNC_ADD) || mode >= GLenum(NUM_BLEND_EQUATIONS) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	c.Blend_equation = mode
}
func BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	for {
		c.Blend_color.X = clampf_01(float32(red))
		c.Blend_color.Y = clampf_01(float32(green))
		c.Blend_color.Z = clampf_01(float32(blue))
		c.Blend_color.W = clampf_01(float32(alpha))
		if true {
			break
		}
	}
}
func LogicOp(opcode GLenum) {
	if opcode < GLenum(GL_CLEAR) || opcode > GLenum(GL_INVERT) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
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
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	c.Scissor_lx = x
	c.Scissor_ly = y
	c.Scissor_ux = GLsizei(x + GLint(width))
	c.Scissor_uy = GLsizei(y + GLint(height))
}
func StencilFunc(func_ GLenum, ref GLint, mask GLuint) {
	if func_ < GLenum(GL_LESS) || func_ > GLenum(GL_NEVER) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
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
	if face < GLenum(GL_FRONT) || face > GLenum(GL_FRONT_AND_BACK) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if face == GLenum(GL_FRONT_AND_BACK) {
		StencilFunc(func_, ref, mask)
		return
	}
	if func_ < GLenum(GL_LESS) || func_ > GLenum(GL_NEVER) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if ref > math.MaxUint8 {
		ref = math.MaxUint8
	}
	if ref < 0 {
		ref = 0
	}
	if face == GLenum(GL_FRONT) {
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
	if (sfail < GLenum(GL_INVERT) || sfail > GLenum(GL_DECR_WRAP)) && sfail != GLenum(GL_ZERO) || (dpfail < GLenum(GL_INVERT) || dpfail > GLenum(GL_DECR_WRAP)) && sfail != GLenum(GL_ZERO) || (dppass < GLenum(GL_INVERT) || dppass > GLenum(GL_DECR_WRAP)) && sfail != GLenum(GL_ZERO) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
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
	if face < GLenum(GL_FRONT) || face > GLenum(GL_FRONT_AND_BACK) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if face == GLenum(GL_FRONT_AND_BACK) {
		StencilOp(sfail, dpfail, dppass)
		return
	}
	if (sfail < GLenum(GL_INVERT) || sfail > GLenum(GL_DECR_WRAP)) && sfail != GLenum(GL_ZERO) || (dpfail < GLenum(GL_INVERT) || dpfail > GLenum(GL_DECR_WRAP)) && sfail != GLenum(GL_ZERO) || (dppass < GLenum(GL_INVERT) || dppass > GLenum(GL_DECR_WRAP)) && sfail != GLenum(GL_ZERO) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if face == GLenum(GL_FRONT) {
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
	if face < GLenum(GL_FRONT) || face > GLenum(GL_FRONT_AND_BACK) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if face == GLenum(GL_FRONT_AND_BACK) {
		StencilMask(mask)
		return
	}
	if face == GLenum(GL_FRONT) {
		c.Stencil_writemask = mask
	} else {
		c.Stencil_writemask_back = mask
	}
}
func MapBuffer(target GLenum, access GLenum) unsafe.Pointer {
	if target != GLenum(GL_ARRAY_BUFFER) && target != GLenum(GL_ELEMENT_ARRAY_BUFFER) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return nil
	}
	if access != GLenum(GL_READ_ONLY) && access != GLenum(GL_WRITE_ONLY) && access != GLenum(GL_READ_WRITE) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return nil
	}
	target -= GLenum(GL_ARRAY_BUFFER)
	var data unsafe.Pointer = nil
	pglGetBufferData(c.Bound_buffers[target], &data)
	return data
}
func MapNamedBuffer(buffer GLuint, access GLenum) unsafe.Pointer {
	if access != GLenum(GL_READ_ONLY) && access != GLenum(GL_WRITE_ONLY) && access != GLenum(GL_READ_WRITE) {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return nil
	}
	var data unsafe.Pointer = nil
	pglGetBufferData(buffer, &data)
	return data
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
	return GL_TRUE
}
func UnmapNamedBuffer(buffer GLuint) GLboolean {
	return GL_TRUE
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
	case GL_REPEAT:
		tmp = i - size*(i/size)
		if tmp < 0 {
			tmp = size + tmp
		}
		return tmp
	case GL_CLAMP_TO_BORDER:
		fallthrough
	case GL_CLAMP_TO_EDGE:
		return clampi(i, 0, size-1)
	case GL_MIRRORED_REPEAT:
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
		libc.Assert(false)
		return 0
	}
}
func texture1D(tex GLuint, x float32) Vec4 {
	var (
		i0      int64
		i1      int64
		t       *glTexture = (*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(tex)))
		texdata *Color     = (*Color)(unsafe.Pointer(t.Data))
		w       float64    = float64(t.W) - EPSILON
		xw      float64    = float64(x) * w
	)
	if t.Mag_filter == GLenum(GL_NEAREST) {
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
		t       *glTexture = (*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(tex)))
		texdata *Color     = (*Color)(unsafe.Pointer(t.Data))
		w       int64      = int64(t.W)
		h       int64      = int64(t.H)
		dw      float64    = float64(w) - EPSILON
		dh      float64    = float64(h) - EPSILON
		xw      float64    = float64(x) * dw
		yh      float64    = float64(y) * dh
	)
	if t.Mag_filter == GLenum(GL_NEAREST) {
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
		t       *glTexture = (*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(tex)))
		texdata *Color     = (*Color)(unsafe.Pointer(t.Data))
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
	if t.Mag_filter == GLenum(GL_NEAREST) {
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
		t       *glTexture = (*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(tex)))
		texdata *Color     = (*Color)(unsafe.Pointer(t.Data))
		w       int64      = int64(t.W)
		h       int64      = int64(t.H)
		dw      float64    = float64(w) - EPSILON
		dh      float64    = float64(h) - EPSILON
		plane   int64      = w * h
		xw      float64    = float64(x) * dw
		yh      float64    = float64(y) * dh
	)
	if t.Mag_filter == GLenum(GL_NEAREST) {
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
		t       *glTexture = (*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(tex)))
		texdata *Color     = (*Color)(unsafe.Pointer(t.Data))
		w       int64      = int64(t.W)
		h       int64      = int64(t.H)
		xw      float64    = float64(x)
		yh      float64    = float64(y)
	)
	if t.Mag_filter == GLenum(GL_NEAREST) {
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
		tex     *glTexture = (*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(texture)))
		texdata *Color     = (*Color)(unsafe.Pointer(tex.Data))
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
	if tex.Mag_filter == GLenum(GL_NEAREST) {
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
