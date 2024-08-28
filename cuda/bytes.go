package cuda

// This file provides GPU byte slices, used to store regions.

import (
	"fmt"
	"unsafe"

	"github.com/MathieuMoalic/amumax/cuda/cu"
	"github.com/MathieuMoalic/amumax/util"
)

// 3D byte slice, used for region lookup.
type Bytes struct {
	Ptr unsafe.Pointer
	Len int
}

// Construct new byte slice with given length,
// initialised to zeros.
func NewBytes(Len int) *Bytes {
	ptr := cu.MemAlloc(int64(Len))
	cu.MemsetD8(cu.DevicePtr(ptr), 0, int64(Len))
	return &Bytes{unsafe.Pointer(uintptr(ptr)), Len}
}

// Upload src (host) to dst (gpu).
func (dst *Bytes) Upload(src []byte) {
	util.Argument(dst.Len == len(src))
	MemCpyHtoD(dst.Ptr, unsafe.Pointer(&src[0]), int64(dst.Len))
}

// Copy on device: dst = src.
func (dst *Bytes) Copy(src *Bytes) {
	util.Argument(dst.Len == src.Len)
	MemCpy(dst.Ptr, src.Ptr, int64(dst.Len))
}

// Copy to host: dst = src.
func (src *Bytes) Download(dst []byte) {
	util.Argument(src.Len == len(dst))
	MemCpyDtoH(unsafe.Pointer(&dst[0]), src.Ptr, int64(src.Len))
}

// Set one element to value.
// data.Index can be used to find the index for x,y,z.
func (dst *Bytes) Set(index int, value byte) {
	if index < 0 || index >= dst.Len {
		util.Log.PanicIfError(fmt.Errorf("Bytes.Set: index out of range: %d", index))
	}
	src := value
	MemCpyHtoD(unsafe.Pointer(uintptr(dst.Ptr)+uintptr(index)), unsafe.Pointer(&src), 1)
}

// Get one element.
// data.Index can be used to find the index for x,y,z.
func (src *Bytes) Get(index int) byte {
	if index < 0 || index >= src.Len {
		util.Log.PanicIfError(fmt.Errorf("Bytes.Set: index out of range: %v", index))
	}
	var dst byte
	MemCpyDtoH(unsafe.Pointer(&dst), unsafe.Pointer(uintptr(src.Ptr)+uintptr(index)), 1)
	return dst
}

// Frees the GPU memory and disables the slice.
func (b *Bytes) Free() {
	if b.Ptr != nil {
		cu.MemFree(cu.DevicePtr(uintptr(b.Ptr)))
	}
	b.Ptr = nil
	b.Len = 0
}
