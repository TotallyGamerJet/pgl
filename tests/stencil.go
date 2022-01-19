package tests

import (
	"github.com/totallygamerjet/pgl"
	"testing"
	"unsafe"
)

func stencil_test(_ *testing.T) {
	type stencil_uniforms struct {
		mvp_mat pgl.Mat4
	}
	var smooth = [4]pgl.GLenum{pgl.SMOOTH, pgl.SMOOTH, pgl.SMOOTH, pgl.SMOOTH}

	var points = []float32{
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
		0, 0.5, 0,
	}

	var color_array = []float32{
		1.0, 0.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
	}

	var the_uniforms stencil_uniforms

	var triangle pgl.GLuint
	pgl.GenBuffers(1, &triangle)
	pgl.GenBuffers(1, &triangle)
	pgl.BindBuffer(pgl.GL_ARRAY_BUFFER, triangle)
	pgl.BufferData(pgl.GL_ARRAY_BUFFER, pgl.GLsizei(len(points))*pgl.GLsizei(unsafe.Sizeof(float32(0))), unsafe.Pointer(&points[0]), pgl.GL_STATIC_DRAW)
	pgl.EnableVertexAttribArray(0)
	pgl.VertexAttribPointer(0, 3, pgl.GL_FLOAT, pgl.GL_FALSE, 0, 0)

	var colors pgl.GLuint
	pgl.GenBuffers(1, &colors)
	pgl.BindBuffer(pgl.GL_ARRAY_BUFFER, colors)
	pgl.BufferData(pgl.GL_ARRAY_BUFFER, pgl.GLsizei(len(color_array))*pgl.GLsizei(unsafe.Sizeof(float32(0))), unsafe.Pointer(&color_array[0]), pgl.GL_STATIC_DRAW)
	pgl.EnableVertexAttribArray(4)
	pgl.VertexAttribPointer(4, 4, pgl.GL_FLOAT, pgl.GL_FALSE, 0, 0)

	var myshader = pgl.NewProgram(func(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		*(*pgl.Vec4)(unsafe.Pointer(vs_output)) = unsafe.Slice((*pgl.Vec4)(vertex_attribs), 5)[4]
		builtins.Gl_Position = pgl.Mult_mat4_vec4(uniforms.(*stencil_uniforms).mvp_mat, *(*pgl.Vec4)(vertex_attribs))
	}, func(fs_input *float32, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		builtins.Gl_FragColor = *(*pgl.Vec4)(unsafe.Pointer(fs_input))
	}, 4, smooth[:], pgl.GL_FALSE)
	pgl.UseProgram(myshader)
	pgl.SetUniform(&the_uniforms)

	var basic = pgl.NewProgram(func(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		builtins.Gl_Position = pgl.Mult_mat4_vec4(uniforms.(*stencil_uniforms).mvp_mat, *(*pgl.Vec4)(vertex_attribs))
	}, func(fs_input *float32, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		builtins.Gl_FragColor = pgl.Vec4{1.0, 0.0, 0.0, 1.0}
	}, 0, nil, pgl.GL_FALSE)
	pgl.UseProgram(basic)
	pgl.SetUniform(&the_uniforms)

	pgl.ClearColor(0, 0, 0, 1)
	pgl.Enable(pgl.GL_STENCIL_TEST)
	pgl.StencilFunc(pgl.GL_NOTEQUAL, 1, 0xFF)
	pgl.StencilOp(pgl.GL_KEEP, pgl.GL_REPLACE, pgl.GL_REPLACE)

	// TODO Apparently this is necessary, what's the spec say?
	// should the color buffer and stencil buffer be initialized to 0 on
	// startup automatically or does the user have to do an initial clear?
	pgl.Clear(pgl.GL_COLOR_BUFFER_BIT | pgl.GL_STENCIL_BUFFER_BIT)

	pgl.UseProgram(myshader)

	the_uniforms.mvp_mat = pgl.Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	pgl.StencilFunc(pgl.GL_ALWAYS, 1, 0xFF)
	pgl.StencilMask(0xFF)
	pgl.DrawArrays(pgl.GL_TRIANGLES, 0, 3)

	pgl.UseProgram(basic)
	pgl.StencilFunc(pgl.GL_NOTEQUAL, 1, 0xFF)
	pgl.StencilMask(0x00)

	pgl.Scale_mat4(&the_uniforms.mvp_mat, 1.2, 1.2, 1.2)
	pgl.DrawArrays(pgl.GL_TRIANGLES, 0, 3)
}
