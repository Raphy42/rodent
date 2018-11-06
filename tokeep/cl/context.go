package cl

// #include "cl.h"
import "C"
import "errors"

type Context struct {
	handle  C.cl_context
	devices []*Device
}

func (c *Context) Dispose() {
	if c.handle != nil {
		C.clReleaseContext(c.handle)
		c.handle = nil
	}
}

func CreateContext(platform *Platform, device *Device) (*Context, error) {
	context := C.interop_context_helper(platform.id, device.id)
	if context == nil {
		return nil, errors.New("(interop) clCreateContext: unable to create context")
	}
	return &Context{handle:context, devices: []*Device{device}}, nil
}