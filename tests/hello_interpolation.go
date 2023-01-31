package tests

import (
	"github.com/totallygamerjet/pgl"
	"testing"
	"unsafe"
)

func hello_interpolation(_ *testing.T) {
	type ht_uniforms struct {
		mvp_mat pgl.Mat4
		v_color pgl.Vec4
	}
	var smooth = [4]pgl.GLenum{pgl.SMOOTH, pgl.SMOOTH, pgl.SMOOTH, pgl.SMOOTH}

	var points_n_colors = [...]float32{
		-0.5, -0.5, 0.0,
		1.0, 0.0, 0.0,

		0.5, -0.5, 0.0,
		0.0, 1.0, 0.0,

		0.0, 0.5, 0.0,
		0.0, 0.0, 1.0,
	}

	var the_uniforms ht_uniforms
	var identity = pgl.Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	var triangle pgl.GLuint
	pgl.GenBuffers(1, &triangle)
	pgl.BindBuffer(pgl.ARRAY_BUFFER, triangle)
	pgl.BufferData(pgl.ARRAY_BUFFER, pgl.GLsizei(unsafe.Sizeof(points_n_colors)), unsafe.Pointer(&points_n_colors[0]), pgl.STATIC_DRAW)
	pgl.EnableVertexAttribArray(0)
	pgl.VertexAttribPointer(0, 3, pgl.FLOAT, pgl.FALSE, pgl.GLsizei(unsafe.Sizeof(float32(0))*6), 0)
	pgl.EnableVertexAttribArray(4)
	pgl.VertexAttribPointer(4, 3, pgl.FLOAT, pgl.FALSE, pgl.GLsizei(unsafe.Sizeof(float32(0))*6), pgl.GLsizei(unsafe.Sizeof(float32(0))*3))

	var myshader = pgl.NewProgram(func(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		var v_attribs = unsafe.Slice((*pgl.Vec4)(vertex_attribs), 5)
		*(*pgl.Vec4)(unsafe.Pointer(vs_output)) = v_attribs[4] //color

		builtins.Gl_Position = pgl.Mult_mat4_vec4(uniforms.(*ht_uniforms).mvp_mat, v_attribs[0])
	}, func(fs_input *float32, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		builtins.Gl_FragColor = *(*pgl.Vec4)(unsafe.Pointer(fs_input))
	}, 4, smooth[:], pgl.FALSE)
	pgl.UseProgram(myshader)

	pgl.SetUniform(&the_uniforms)

	the_uniforms.mvp_mat = identity

	pgl.ClearColor(0, 0, 0, 1)

	pgl.Clear(pgl.COLOR_BUFFER_BIT)
	pgl.DrawArrays(pgl.TRIANGLES, 0, 3)

}
