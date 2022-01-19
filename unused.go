//go:build ignore
// +build ignore

package pgl

// This file contains functions that aren't used anywhere else

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
	libc.Free(unsafe.Pointer(&tmp.A))
	libc.Free(unsafe.Pointer(tmp))
}
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
	if vec.A == nil {
		vec.A = unsafe.Slice((*glVertex)(libc.Malloc(int(vec.Capacity*uint64(unsafe.Sizeof(glVertex{}))))), vec.Capacity)
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
