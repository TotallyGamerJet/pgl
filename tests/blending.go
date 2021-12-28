package tests

import (
	"pgl"
	"testing"
	"unsafe"
)

func blend_test(_ *testing.T) {
	type blend_uniforms struct {
		v_color pgl.Vec4
	}
	var (
		Red   = pgl.Vec4{1.0, 0.0, 0.0, 1.0}
		Green = pgl.Vec4{0.0, 1.0, 0.0, 1.0}
		Blue  = pgl.Vec4{0.0, 0.0, 1.0, 1.0}
		Black = pgl.Vec4{0.0, 0.0, 0.0, 1.0}
	)

	var points = [...]float32{
		-0.75, 0.75, 0,
		-0.75, 0.25, 0,
		-0.25, 0.75, 0,
		-0.25, 0.25, 0,

		0.25, 0.75, 0,
		0.25, 0.25, 0,
		0.75, 0.75, 0,
		0.75, 0.25, 0,

		-0.75, -0.25, 0,
		-0.75, -0.75, 0,
		-0.25, -0.25, 0,
		-0.25, -0.75, 0,

		0.25, -0.25, 0,
		0.25, -0.75, 0,
		0.75, -0.25, 0,
		0.75, -0.75, 0,

		//mix with white
		-0.15, 0.15, -0.1,
		-0.15, -0.15, -0.1,
		0.15, 0.15, -0.1,
		0.15, -0.15, -0.1,

		// mix with red
		-0.40, 0.65, -0.1,
		-0.40, 0.35, -0.1,
		-0.10, 0.65, -0.1,
		-0.10, 0.35, -0.1,

		// mix with green
		0.10, 0.65, -0.1,
		0.10, 0.35, -0.1,
		0.40, 0.65, -0.1,
		0.40, 0.35, -0.1,

		// mix with blue
		-0.40, -0.35, -0.1,
		-0.40, -0.65, -0.1,
		-0.10, -0.35, -0.1,
		-0.10, -0.65, -0.1,

		// mix with black
		0.10, -0.35, -0.1,
		0.10, -0.65, -0.1,
		0.40, -0.35, -0.1,
		0.40, -0.65, -0.1,
	}

	var the_uniforms blend_uniforms

	var triangle pgl.GLuint
	pgl.GenBuffers(1, &triangle)
	pgl.BindBuffer(pgl.GL_ARRAY_BUFFER, triangle)
	pgl.BufferData(pgl.GL_ARRAY_BUFFER, pgl.GLsizei(unsafe.Sizeof(points)), unsafe.Pointer(&points[0]), pgl.GL_STATIC_DRAW)
	pgl.EnableVertexAttribArray(0)
	pgl.VertexAttribPointer(0, 3, pgl.GL_FLOAT, pgl.GL_FALSE, 0, 0)

	var myshader = pgl.NewProgram(func(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		builtins.Gl_Position = *(*pgl.Vec4)(vertex_attribs)
	}, func(fs_input *float32, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		builtins.Gl_FragColor = uniforms.(*blend_uniforms).v_color
	}, 0, nil, pgl.GL_FALSE)
	pgl.UseProgram(myshader)

	pgl.SetUniform(&the_uniforms)

	the_uniforms.v_color = Red

	pgl.ClearColor(1, 1, 1, 1)

	pgl.Clear(pgl.GL_COLOR_BUFFER_BIT)

	the_uniforms.v_color = Red
	pgl.DrawArrays(pgl.GL_TRIANGLE_STRIP, 0, 4)
	the_uniforms.v_color = Green
	pgl.DrawArrays(pgl.GL_TRIANGLE_STRIP, 4, 4)
	the_uniforms.v_color = Blue
	pgl.DrawArrays(pgl.GL_TRIANGLE_STRIP, 8, 4)
	the_uniforms.v_color = Black
	pgl.DrawArrays(pgl.GL_TRIANGLE_STRIP, 12, 4)

	pgl.Enable(pgl.GL_BLEND)
	pgl.BlendFunc(pgl.GL_SRC_ALPHA, pgl.GL_ONE_MINUS_SRC_ALPHA)
	the_uniforms.v_color = pgl.Vec4{1, 0, 0, 0.5}
	pgl.DrawArrays(pgl.GL_TRIANGLE_STRIP, 16, 4)

	pgl.DrawArrays(pgl.GL_TRIANGLE_STRIP, 20, 4)

	pgl.DrawArrays(pgl.GL_TRIANGLE_STRIP, 24, 4)

	pgl.DrawArrays(pgl.GL_TRIANGLE_STRIP, 28, 4)

	pgl.DrawArrays(pgl.GL_TRIANGLE_STRIP, 32, 4)

	pgl.Disable(pgl.GL_BLEND)

}
