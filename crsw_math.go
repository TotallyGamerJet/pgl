package pgl

import (
	"fmt"
	"github.com/chewxy/math32"
	"github.com/gotranspile/cxgo/runtime/libc"
	"io"
	"math"
	"os"
)

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
func fprint_vec2(f io.Writer, v vec2, append string) {
	fmt.Fprintf(f, "(%f, %f)%s", v.X, v.Y, append)
}
func fprint_vec3(f io.Writer, v Vec3, append string) {
	fmt.Fprintf(f, "(%f, %f, %f)%s", v.X, v.Y, v.Z, append)
}
func fprint_vec4(f io.Writer, v Vec4, append string) {
	fmt.Fprintf(f, "(%f, %f, %f, %f)%s", v.X, v.Y, v.Z, v.W, append)
}
func print_vec2(v vec2, append string) {
	fmt.Printf("(%f, %f)%s", v.X, v.Y, append)
}
func print_vec3(v Vec3, append string) {
	fmt.Printf("(%f, %f, %f)%s", v.X, v.Y, v.Z, append)
}
func print_vec4(v Vec4, append string) {
	fmt.Printf("(%f, %f, %f, %f)%s", v.X, v.Y, v.Z, v.W, append)
}
func fread_vec2(f io.Reader, v *vec2) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%f, %f)", &v.X, &v.Y)
	return int64(boolToInt(tmp == 2))
}
func fread_vec3(f io.Reader, v *Vec3) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%f, %f, %f)", &v.X, &v.Y, &v.Z)
	return int64(boolToInt(tmp == 3))
}
func fread_vec4(f io.Reader, v *Vec4) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%f, %f, %f, %f)", &v.X, &v.Y, &v.Z, &v.W)
	return int64(boolToInt(tmp == 4))
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

func fprint_dvec2(f io.Writer, v dvec2, append string) {
	fmt.Fprintf(f, "(%f, %f)%s", v.X, v.Y, append)
}
func fprint_dvec3(f io.Writer, v dvec3, append string) {
	fmt.Fprintf(f, "(%f, %f, %f)%s", v.X, v.Y, v.Z, append)
}
func fprint_dvec4(f io.Writer, v dvec4, append string) {
	fmt.Fprintf(f, "(%f, %f, %f, %f)%s", v.X, v.Y, v.Z, v.W, append)
}
func fread_dvec2(f io.Reader, v *dvec2) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%lf, %lf)", &v.X, &v.Y)
	return int64(boolToInt(tmp == 2))
}
func fread_dvec3(f io.Reader, v *dvec3) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%lf, %lf, %lf)", &v.X, &v.Y, &v.Z)
	return int64(boolToInt(tmp == 3))
}
func fread_dvec4(f io.Reader, v *dvec4) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%lf, %lf, %lf, %lf)", &v.X, &v.Y, &v.Z, &v.W)
	return int64(boolToInt(tmp == 4))
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
func fprint_ivec2(f io.Writer, v ivec2, append string) {
	fmt.Fprintf(f, "(%d, %d)%s", v.X, v.Y, append)
}
func fprint_ivec3(f io.Writer, v ivec3, append string) {
	fmt.Fprintf(f, "(%d, %d, %d)%s", v.X, v.Y, v.Z, append)
}
func fprint_ivec4(f io.Writer, v ivec4, append string) {
	fmt.Fprintf(f, "(%d, %d, %d, %d)%s", v.X, v.Y, v.Z, v.W, append)
}
func fread_ivec2(f io.Reader, v *ivec2) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%d, %d)", &v.X, &v.Y)
	return int64(boolToInt(tmp == 2))
}
func fread_ivec3(f io.Reader, v *ivec3) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%d, %d, %d)", &v.X, &v.Y, &v.Z)
	return int64(boolToInt(tmp == 3))
}
func fread_ivec4(f io.Reader, v *ivec4) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%d, %d, %d, %d)", &v.X, &v.Y, &v.Z, &v.W)
	return int64(boolToInt(tmp == 4))
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

func fprint_uvec2(f io.Writer, v uvec2, append string) {
	fmt.Fprintf(f, "(%u, %u)%s", v.X, v.Y, append)
}
func fprint_uvec3(f io.Writer, v uvec3, append string) {
	fmt.Fprintf(f, "(%u, %u, %u)%s", v.X, v.Y, v.Z, append)
}
func fprint_uvec4(f io.Writer, v uvec4, append string) {
	fmt.Fprintf(f, "(%u, %u, %u, %u)%s", v.X, v.Y, v.Z, v.W, append)
}
func fread_uvec2(f io.Reader, v *uvec2) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%u, %u)", &v.X, &v.Y)
	return int64(boolToInt(tmp == 2))
}
func fread_uvec3(f io.Reader, v *uvec3) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%u, %u, %u)", &v.X, &v.Y, &v.Z)
	return int64(boolToInt(tmp == 3))
}
func fread_uvec4(f io.Reader, v *uvec4) int64 {
	var tmp, _ = fmt.Fscanf(f, " (%u, %u, %u, %u)", &v.X, &v.Y, &v.Z, &v.W)
	return int64(boolToInt(tmp == 4))
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
	return int64(boolToInt(a.X == b.X && a.Y == b.Y))
}
func equal_vec3s(a Vec3, b Vec3) int64 {
	return int64(boolToInt(a.X == b.X && a.Y == b.Y && a.Z == b.Z))
}
func equal_vec4s(a Vec4, b Vec4) int64 {
	return int64(boolToInt(a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W))
}
func equal_epsilon_vec2s(a vec2, b vec2, epsilon float32) int64 {
	return int64(boolToInt(math.Abs(float64(a.X-b.X)) < float64(epsilon) && math.Abs(float64(a.Y-b.Y)) < float64(epsilon)))
}
func equal_epsilon_vec3s(a Vec3, b Vec3, epsilon float32) int64 {
	return int64(boolToInt(math.Abs(float64(a.X-b.X)) < float64(epsilon) && math.Abs(float64(a.Y-b.Y)) < float64(epsilon) && math.Abs(float64(a.Z-b.Z)) < float64(epsilon)))
}
func equal_epsilon_vec4s(a Vec4, b Vec4, epsilon float32) int64 {
	return int64(boolToInt(math.Abs(float64(a.X-b.X)) < float64(epsilon) && math.Abs(float64(a.Y-b.Y)) < float64(epsilon) && math.Abs(float64(a.Z-b.Z)) < float64(epsilon) && math.Abs(float64(a.W-b.W)) < float64(epsilon)))
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
func fprint_mat2(f io.Writer, m mat2, append string) {
	fmt.Fprintf(f, "[(%f, %f)\n (%f, %f)]%s", m[0], m[2], m[1], m[3], append)
}
func fprint_mat3(f io.Writer, m mat3, append string) {
	fmt.Fprintf(f, "[(%f, %f, %f)\n (%f, %f, %f)\n (%f, %f, %f)]%s", m[0], m[3], m[6], m[1], m[4], m[7], m[2], m[5], m[8], append)
}
func fprint_mat4(f io.Writer, m Mat4, append string) {
	fmt.Fprintf(f, "[(%f, %f, %f, %f)\n(%f, %f, %f, %f)\n(%f, %f, %f, %f)\n(%f, %f, %f, %f)]%s", m[0], m[4], m[8], m[12], m[1], m[5], m[9], m[13], m[2], m[6], m[10], m[14], m[3], m[7], m[11], m[15], append)
}
func print_mat2(m mat2, append string) {
	fprint_mat2(os.Stdout, m, append)
}
func print_mat3(m mat3, append string) {
	fprint_mat3(os.Stdout, m, append)
}
func print_mat4(m Mat4, append string) {
	fprint_mat4(os.Stdout, m, append)
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
func Scale_mat4(m *Mat4, x float32, y float32, z float32) {
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
func print_Color(c Color, append string) {
	fmt.Printf("(%d, %d, %d, %d)%s", c.R, c.G, c.B, c.A, append)
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
