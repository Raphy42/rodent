package cl

// #include "cl.h"
import "C"
import (
	"errors"
	"unsafe"
)

type DeviceType uint

const (
	DeviceTypeCPU         DeviceType = C.CL_DEVICE_TYPE_CPU
	DeviceTypeGPU         DeviceType = C.CL_DEVICE_TYPE_GPU
	DeviceTypeAccelerator DeviceType = C.CL_DEVICE_TYPE_ACCELERATOR
	DeviceTypeDefault     DeviceType = C.CL_DEVICE_TYPE_DEFAULT
	DeviceTypeAll         DeviceType = C.CL_DEVICE_TYPE_ALL
)

type Device struct {
	id C.cl_device_id
}

func (d *Device) infoString(param C.cl_device_info) (string, error) {
	var cstr [1024]C.char
	var cstrLen C.size_t
	if ok := C.clGetDeviceInfo(d.id, param, 1024, unsafe.Pointer(&cstr), &cstrLen); ok != C.CL_SUCCESS {
		return "", errors.New("clGetDeviceInfo: unable to get param")
	}
	return C.GoStringN((*C.char)(unsafe.Pointer(&cstr)), C.int(cstrLen-1)), nil
}

func (d *Device) Name() (string, error) {
	return d.infoString(C.CL_DEVICE_NAME)
}

func (d *Device) Version() (string, error) {
	return d.infoString(C.CL_DEVICE_VERSION)
}

func (d *Device) DriverVersion() (string, error) {
	return d.infoString(C.CL_DRIVER_VERSION)
}
