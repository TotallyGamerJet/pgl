package tests

import (
	"github.com/totallygamerjet/pgl"
	"testing"
	"unsafe"
)

func hello_triangle(_ *testing.T) {
	type ht_uniforms struct {
		mvp_mat pgl.Mat4
		v_color pgl.Vec4
	}
	var Red = pgl.Vec4{1, 0, 0, 0}

	var points = []float32{
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
		0, 0.5, 0,
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
	pgl.BindBuffer(pgl.GL_ARRAY_BUFFER, triangle)
	pgl.BufferData(pgl.GL_ARRAY_BUFFER, pgl.GLsizei(unsafe.Sizeof(pgl.GLfloat(0))*9), unsafe.Pointer(&points[0]), pgl.GL_STATIC_DRAW)
	pgl.EnableVertexAttribArray(0)
	pgl.VertexAttribPointer(0, 3, pgl.GL_FLOAT, pgl.GL_FALSE, 0, 0)

	var myshader = pgl.NewProgram(func(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		builtins.Gl_Position = pgl.Mult_mat4_vec4(uniforms.(*ht_uniforms).mvp_mat, *(*pgl.Vec4)(vertex_attribs))
	}, func(fs_input *float32, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		builtins.Gl_FragColor = uniforms.(*ht_uniforms).v_color
	}, 0, nil, pgl.GL_FALSE)
	pgl.UseProgram(myshader)

	pgl.SetUniform(&the_uniforms)

	the_uniforms.v_color = Red

	the_uniforms.mvp_mat = identity

	pgl.ClearColor(0, 0, 0, 1)
	pgl.Clear(pgl.GL_COLOR_BUFFER_BIT)
	pgl.DrawArrays(pgl.GL_TRIANGLES, 0, 3)
}
