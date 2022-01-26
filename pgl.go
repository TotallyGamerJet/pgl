package pgl

import (
	"math"
)

const (
	sizeOfFloat32 = 4
	sizeOfU32     = 4
)

func NewProgram(vertex_shader vert_func, fragment_shader frag_func, n GLsizei, interpolation []GLenum, fragdepth_or_discard GLboolean) GLuint {
	if vertex_shader == nil || fragment_shader == nil {
		return 0
	}
	if n > GL_MAX_VERTEX_OUTPUT_COMPONENTS {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return 0
	}
	var tmp glProgram = glProgram{Vertex_shader: vertex_shader, Fragment_shader: fragment_shader, Uniform: nil, Vs_output_size: int64(n), Interpolation: [64]GLenum{}, Fragdepth_or_discard: fragdepth_or_discard, Deleted: GL_FALSE}
	copy(tmp.Interpolation[:], interpolation)
	for i := int64(1); uint64(i) < uint64(len(c.Programs)); i++ {
		if c.Programs[i].Deleted != 0 && i != int64(c.Cur_program) {
			c.Programs[i] = tmp
			return GLuint(int32(i))
		}
	}
	c.Programs = append(c.Programs, tmp)
	return GLuint(len(c.Programs) - 1)
}

func SetUniform(uniform interface{}) {
	c.Programs[c.Cur_program].Uniform = uniform
}

func pglResizeFramebuffer(w uint64, h uint64) []u8 {
	var tmp []u8
	tmp = make([]u8, w*h*sizeOfFloat32)
	copy(tmp, c.Zbuf.Buf)
	if tmp == nil {
		if c.Error == GLenum(GL_NO_ERROR) {
			c.Error = GLenum(GL_OUT_OF_MEMORY)
		}
		return nil
	}
	c.Zbuf.Buf = tmp
	c.Zbuf.W = w
	c.Zbuf.H = h
	c.Zbuf.Lastrow = c.Zbuf.Buf[(h-1)*w*sizeOfFloat32:]
	tmp = make([]u8, w*h*sizeOfU32)
	copy(tmp, c.Back_buffer.Buf)
	if tmp == nil {
		if c.Error == GLenum(GL_NO_ERROR) {
			c.Error = GLenum(GL_OUT_OF_MEMORY)
		}
		return nil
	}
	c.Back_buffer.Buf = tmp
	c.Back_buffer.W = w
	c.Back_buffer.H = h
	c.Back_buffer.Lastrow = c.Back_buffer.Buf[(h-1)*w*sizeOfFloat32:]
	return tmp
}

func pglClearScreen() {
	for i := range c.Back_buffer.Buf {
		c.Back_buffer.Buf[i] = math.MaxUint8
	}
}

func pglSetInterp(interpolation []GLenum) {
	n := len(interpolation)
	c.Programs[c.Cur_program].Vs_output_size = int64(n)
	c.Vs_output.Size = int64(n)
	copy(c.Programs[c.Cur_program].Interpolation[:], interpolation)
	c.Vs_output.Output_buf = make([]float32, uint64(n*MAX_VERTICES))
	c.Vs_output.Interpolation = c.Programs[c.Cur_program].Interpolation[:]
}

func pglDrawFrame() {
	var frag_shader frag_func = c.Programs[c.Cur_program].Fragment_shader
	for y := float32(0.5); y < float32(c.Back_buffer.H); y++ {
		for x := float32(0.5); x < float32(c.Back_buffer.W); x++ {
			c.Builtins.Gl_FragCoord.X = x
			c.Builtins.Gl_FragCoord.Y = y
			c.Builtins.Discard = GL_FALSE
			frag_shader(nil, &c.Builtins, c.Programs[c.Cur_program].Uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(x), int64(y))
			}
		}
	}
}

func pglBufferData(target GLenum, size GLsizei, data []u8, usage GLenum) {
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
	if data == nil {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if c.Buffers[c.Bound_buffers[target]].User_owned == false {
		c.Buffers[c.Bound_buffers[target]].Data = nil
	}
	c.Buffers[c.Bound_buffers[target]].Data = data
	c.Buffers[c.Bound_buffers[target]].User_owned = true
	c.Buffers[c.Bound_buffers[target]].Size = size
	if target == GLenum(GL_ELEMENT_ARRAY_BUFFER) {
		c.Vertex_arrays[c.Cur_vertex_array].Element_buffer = c.Bound_buffers[target]
	}
}

func pglTexImage1D(target GLenum, level GLint, internalFormat GLint, width GLsizei, border GLint, format GLenum, type_ GLenum, data []u8) {
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
	if data == nil {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	var cur_tex int64 = int64(c.Bound_textures[target-GLenum(GL_TEXTURE_UNBOUND)-1])
	c.Textures[cur_tex].W = uint64(width)
	if type_ != GLenum(GL_UNSIGNED_BYTE) {
		return
	}
	var components int64
	_ = components
	switch format {
	case GL_RED:
		components = 1
	case GL_RG:
		components = 2
	case GL_RGB, GL_BGR:
		components = 3
	case GL_RGBA, GL_BGRA:
		components = 4
	default:
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if c.Textures[cur_tex].User_owned == 0 {
		c.Textures[cur_tex].Data = nil
	}
	if len(data) != int(width) {
		panic("len(data) != width")
	}
	c.Textures[cur_tex].Data = data
	c.Textures[cur_tex].User_owned = GL_TRUE
}

func pglTexImage2D(target GLenum, level GLint, internalFormat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, type_ GLenum, data []u8) {
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
	if data == nil {
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
	_ = components
	switch format {
	case GL_RED:
		components = 1
	case GL_RG:
		components = 2
	case GL_RGB, GL_BGR:
		components = 3
	case GL_RGBA, GL_BGRA:
		components = 4
	default:
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	var cur_tex int64
	if target == GLenum(GL_TEXTURE_2D) || target == GLenum(GL_TEXTURE_RECTANGLE) {
		cur_tex = int64(c.Bound_textures[target-GLenum(GL_TEXTURE_UNBOUND)-1])
		c.Textures[cur_tex].W = uint64(width)
		c.Textures[cur_tex].H = uint64(height)
		if c.Textures[cur_tex].User_owned == 0 {
			c.Textures[cur_tex].Data = nil
		}
		if len(data) != int(width*height) {
			panic("len(data) != width*height")
		}
		c.Textures[cur_tex].Data = data
		c.Textures[cur_tex].User_owned = GL_TRUE
	} else {
	}
}

func pglTexImage3D(target GLenum, level GLint, internalFormat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, type_ GLenum, data []u8) {
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
	if data == nil {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	var cur_tex int64 = int64(c.Bound_textures[target-GLenum(GL_TEXTURE_UNBOUND)-1])
	c.Textures[cur_tex].W = uint64(width)
	c.Textures[cur_tex].H = uint64(height)
	c.Textures[cur_tex].D = uint64(depth)
	if type_ != GLenum(GL_UNSIGNED_BYTE) {
		return
	}
	var components int64
	_ = components
	switch format {
	case GL_RED:
		components = 1
	case GL_RG:
		components = 2
	case GL_RGB, GL_BGR:
		components = 3
	case GL_RGBA, GL_BGRA:
		components = 4
	default:
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_ENUM)
		}
		return
	}
	if c.Textures[cur_tex].User_owned == 0 {
		c.Textures[cur_tex].Data = nil
	}
	if len(data) != int(width*height*depth) {
		panic("w*h*d != len(data)")
	}
	c.Textures[cur_tex].Data = data
	c.Textures[cur_tex].User_owned = GL_TRUE
}

func pglGetBufferData(buffer GLuint) []u8 {
	if !(buffer != 0 && uint64(buffer) < uint64(len(c.Buffers)) && c.Buffers[buffer].Deleted == false) {
		c.Error = GLenum(GL_INVALID_OPERATION)
		return nil
	}
	return c.Buffers[buffer].Data
}

func pglGetTextureData(texture GLuint) []u8 {
	if !(uint64(texture) < uint64(len(c.Textures)) && c.Textures[texture].Deleted == 0) {
		c.Error = GLenum(GL_INVALID_OPERATION)

	}
	return c.Textures[texture].Data
}
