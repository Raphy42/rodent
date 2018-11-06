package cl

// #include "cl.h"
// #cgo linux pkg-config: OpenCL OpenGL
// #cgo darwin LDFLAGS: -framework OpenCL -framework OpenGL
// #cgo windows LDFLAGS: -lOpenCL -lOpenGL
import "C"
