package tests

import (
	"github.com/totallygamerjet/pgl"
	"testing"
	"unsafe"
)

/*
typedef struct hi_uniforms
{
	mat4 mvp_mat;
	vec4 v_color;
} hi_uniforms;


void hi_smooth_vs(float* vs_output, void* vertex_attribs, Shader_Builtins* builtins, void* uniforms)
{
	vec4* v_attribs = (vec4*)vertex_attribs;
	((vec4*)vs_output)[0] = v_attribs[4]; //color

	builtins->gl_Position = mult_mat4_vec4(*((mat4*)uniforms), v_attribs[0]);
}

void hi_smooth_fs(float* fs_input, Shader_Builtins* builtins, void* uniforms)
{
	builtins->gl_FragColor = ((vec4*)fs_input)[0];
}

void hello_interpolation(int argc, char** argv, void* data)
{
	GLenum smooth[4] = { SMOOTH, SMOOTH, SMOOTH, SMOOTH };

	float points_n_colors[] = {
		-0.5, -0.5, 0.0,
		 1.0,  0.0, 0.0,

		 0.5, -0.5, 0.0,
		 0.0,  1.0, 0.0,

		 0.0,  0.5, 0.0,
		 0.0,  0.0, 1.0 };

	hi_uniforms the_uniforms;
	mat4 identity = IDENTITY_MAT4();

	GLuint triangle;
	glGenBuffers(1, &triangle);
	glBindBuffer(GL_ARRAY_BUFFER, triangle);
	glBufferData(GL_ARRAY_BUFFER, sizeof(points_n_colors), points_n_colors, GL_STATIC_DRAW);
	glEnableVertexAttribArray(0);
	glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, sizeof(float)*6, 0);
	glEnableVertexAttribArray(4);
	glVertexAttribPointer(4, 3, GL_FLOAT, GL_FALSE, sizeof(float)*6, sizeof(float)*3);

	GLuint myshader = pglCreateProgram(hi_smooth_vs, hi_smooth_fs, 4, smooth, GL_FALSE);
	glUseProgram(myshader);

	pglSetUniform(&the_uniforms);

	memcpy(the_uniforms.mvp_mat, identity, sizeof(mat4));

	glClearColor(0, 0, 0, 1);

	glClear(GL_COLOR_BUFFER_BIT);
	glDrawArrays(GL_TRIANGLES, 0, 3);

}
*/
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
	pgl.BindBuffer(pgl.GL_ARRAY_BUFFER, triangle)
	pgl.BufferData(pgl.GL_ARRAY_BUFFER, pgl.GLsizei(unsafe.Sizeof(points_n_colors)), unsafe.Pointer(&points_n_colors[0]), pgl.GL_STATIC_DRAW)
	pgl.EnableVertexAttribArray(0)
	pgl.VertexAttribPointer(0, 3, pgl.GL_FLOAT, pgl.GL_FALSE, pgl.GLsizei(unsafe.Sizeof(float32(0))*6), 0)
	pgl.EnableVertexAttribArray(4)
	pgl.VertexAttribPointer(4, 3, pgl.GL_FLOAT, pgl.GL_FALSE, pgl.GLsizei(unsafe.Sizeof(float32(0))*6), pgl.GLsizei(unsafe.Sizeof(float32(0))*3))

	var myshader = pgl.NewProgram(func(vs_output *float32, vertex_attribs unsafe.Pointer, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		var v_attribs = unsafe.Slice((*pgl.Vec4)(vertex_attribs), 5)
		*(*pgl.Vec4)(unsafe.Pointer(vs_output)) = v_attribs[4] //color

		builtins.Gl_Position = pgl.Mult_mat4_vec4(uniforms.(*ht_uniforms).mvp_mat, v_attribs[0])
	}, func(fs_input *float32, builtins *pgl.Shader_Builtins, uniforms interface{}) {
		builtins.Gl_FragColor = *(*pgl.Vec4)(unsafe.Pointer(fs_input))
	}, 4, smooth[:], pgl.GL_FALSE)
	pgl.UseProgram(myshader)

	pgl.SetUniform(&the_uniforms)

	the_uniforms.mvp_mat = identity

	pgl.ClearColor(0, 0, 0, 1)

	pgl.Clear(pgl.GL_COLOR_BUFFER_BIT)
	pgl.DrawArrays(pgl.GL_TRIANGLES, 0, 3)

}
