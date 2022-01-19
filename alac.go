package alac

import (
	// #include "alac.h"
	"C"
	"unsafe"
)

// File is a wrapper type around the C type.
type File struct {
	file *C.alac_file
}

// New Allocates new AlacFile.
func New(sampleSize, numberOfChannels int) File {
	alac := File{}
	alac.file = C.alac_create(C.int(sampleSize), C.int(numberOfChannels))
	return alac
}

// DecodeFrame Decodes a frame from inputBuffer and puts it in the outputBuffer.
func (f *File) DecodeFrame(inputBuffer []byte) []byte {
	size := C.int(len(inputBuffer))
	p := unsafe.Pointer(C.malloc(C.size_t(len(inputBuffer))))
	
	cBuf := (*[1 << 30]byte)(p)
	copy(cBuf[:], inputBuffer)

	C.alac_decode_frame(f.file, (*C.uchar)(unsafe.Pointer(&inputBuffer)), p, &size)
	return C.GoBytes(p, size)
}

// SetInfo Set's the "info" for our AlacFile.
func (f *File) SetInfo(inputBuffer []byte) {
	C.alac_set_info(f.file, (*C.char)(unsafe.Pointer(&inputBuffer)))
}

// AllocateBuffers Allocates the C buffers for our AlacFile.
func (f *File) AllocateBuffers() {
	C.alac_allocate_buffers(f.file)
}

// Free frees the C buffers we wrap in our AlacFile type.
func (f *File) Free() {
	C.alac_free(f.file)
}
