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
	op := C.malloc(C.size_t(len(inputBuffer)))
	ip := (*C.uchar)(C.malloc(C.size_t(len(inputBuffer))))
	//cOutBuf := (*[1 << 30]byte)(op)
	copy(C.GoBytes(unsafe.Pointer(ip), size)[:], inputBuffer)

	C.alac_decode_frame(f.file, ip, op, &size)
	return C.GoBytes(op, size)
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
