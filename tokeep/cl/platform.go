package cl

// #include "cl.h"
import "C"
import "github.com/pkg/errors"

const maxPlatforms = 32
const maxDeviceCount = 64

type Platform struct {
	id C.cl_platform_id
}

func ListPlatforms() ([]*Platform, error) {
	var platformIds [maxPlatforms]C.cl_platform_id
	var platformCount C.cl_uint
	if ok := C.clGetPlatformIDs(C.cl_uint(maxPlatforms), &platformIds[0], &platformCount); ok != C.CL_SUCCESS {
		return nil, errors.New("cl.GetPlatformIDs: unable to find id")
	}
	platforms := make([]*Platform, platformCount)
	for i := 0; i < int(platformCount); i++ {
		platforms[i] = &Platform{id: platformIds[i]}
	}
	return platforms, nil
}

func (p *Platform) GetDevices(xtype DeviceType) ([]*Device, error) {
	var deviceIds [maxDeviceCount]C.cl_device_id
	var deviceCount C.cl_uint

	if ok := C.clGetDeviceIDs(p.id, C.cl_device_type(xtype), C.cl_uint(maxDeviceCount), &deviceIds[0], &deviceCount); ok != C.CL_SUCCESS {
		return nil, errors.New("cl.GetDeviceIDs: unable to find devices")
	}

	if deviceCount > maxDeviceCount {
		deviceCount = maxDeviceCount
	}

	devices := make([]*Device, deviceCount)
	for i := 0; i < int(deviceCount); i++ {
		devices[i] = &Device{id: deviceIds[i]}
	}
	return devices, nil
}

