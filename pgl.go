package pgl

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"math"
	"unsafe"
)

func pglClearScreen() {
	libc.MemSet(unsafe.Pointer(&c.Back_buffer.Buf[0]), math.MaxUint8, int(c.Back_buffer.W*c.Back_buffer.H*4))
}

func pglSetInterp(interpolation []GLenum) {
	n := len(interpolation)
	c.Programs.A[c.Cur_program].Vs_output_size = int64(n)
	c.Vs_output.Size = int64(n)
	copy(c.Programs.A[c.Cur_program].Interpolation[:], interpolation)
	cvec_reserve_float(&c.Vs_output.Output_buf, uint64(n*MAX_VERTICES))
	c.Vs_output.Interpolation = c.Programs.A[c.Cur_program].Interpolation[:]
}

func pglDrawFrame() {
	var frag_shader frag_func = c.Programs.A[c.Cur_program].Fragment_shader
	for y := float32(0.5); y < float32(c.Back_buffer.H); y++ {
		for x := float32(0.5); x < float32(c.Back_buffer.W); x++ {
			c.Builtins.Gl_FragCoord.X = x
			c.Builtins.Gl_FragCoord.Y = y
			c.Builtins.Discard = GL_FALSE
			frag_shader(nil, &c.Builtins, c.Programs.A[c.Cur_program].Uniform)
			if c.Builtins.Discard == 0 {
				draw_pixel(c.Builtins.Gl_FragColor, int64(x), int64(y))
			}
		}
	}
}

func pglBufferData(target GLenum, size GLsizei, data unsafe.Pointer, usage GLenum) {
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
	if (c.Buffers.A[c.Bound_buffers[target]]).User_owned == 0 {
		c.Buffers.A[c.Bound_buffers[target]].Data = nil
	}
	(c.Buffers.A[c.Bound_buffers[target]]).Data = unsafe.Slice((*u8)(data), size)
	(c.Buffers.A[c.Bound_buffers[target]]).User_owned = GL_TRUE
	(c.Buffers.A[c.Bound_buffers[target]]).Size = size
	if target == GLenum(GL_ELEMENT_ARRAY_BUFFER) {
		(c.Vertex_arrays.A[c.Cur_vertex_array]).Element_buffer = c.Bound_buffers[target]
	}
}

func pglTexImage1D(target GLenum, level GLint, internalFormat GLint, width GLsizei, border GLint, format GLenum, type_ GLenum, data unsafe.Pointer) {
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
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W = uint64(width)
	if type_ != GLenum(GL_UNSIGNED_BYTE) {
		return
	}
	var components int64
	_ = components
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
	if (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).User_owned == 0 {
		libc.Free(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
	}
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data = (*u8)(data)
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).User_owned = GL_TRUE
}

func pglTexImage2D(target GLenum, level GLint, internalFormat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, type_ GLenum, data unsafe.Pointer) {
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
	if target == GLenum(GL_TEXTURE_2D) || target == GLenum(GL_TEXTURE_RECTANGLE) {
		cur_tex = int64(c.Bound_textures[target-GLenum(GL_TEXTURE_UNBOUND)-1])
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W = uint64(width)
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).H = uint64(height)
		if (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).User_owned == 0 {
			libc.Free(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
		}
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data = (*u8)(data)
		(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).User_owned = GL_TRUE
	} else {
	}
}

func pglTexImage3D(target GLenum, level GLint, internalFormat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, type_ GLenum, data unsafe.Pointer) {
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
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).W = uint64(width)
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).H = uint64(height)
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).D = uint64(depth)
	if type_ != GLenum(GL_UNSIGNED_BYTE) {
		return
	}
	var components int64
	_ = components
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
	if (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).User_owned == 0 {
		libc.Free(unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data))
	}
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).Data = (*u8)(data)
	(*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(cur_tex)))).User_owned = GL_TRUE
}

func GetBufferData(buffer GLuint, data *unsafe.Pointer) {
	if data == nil {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if buffer != 0 && uint64(buffer) < c.Buffers.Size && (c.Buffers.A[buffer]).Deleted == 0 {
		*data = unsafe.Pointer(&(c.Buffers.A[buffer]).Data[0])
	} else if c.Error == 0 {
		c.Error = GLenum(GL_INVALID_OPERATION)
	}
}

func pglGetTextureData(texture GLuint, data *unsafe.Pointer) {
	if data == nil {
		if c.Error == 0 {
			c.Error = GLenum(GL_INVALID_VALUE)
		}
		return
	}
	if uint64(texture) < c.Textures.Size && (*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(texture)))).Deleted == 0 {
		*data = unsafe.Pointer((*(*glTexture)(unsafe.Add(unsafe.Pointer(c.Textures.A), unsafe.Sizeof(glTexture{})*uintptr(texture)))).Data)
	} else if c.Error == 0 {
		c.Error = GLenum(GL_INVALID_OPERATION)
	}
}
