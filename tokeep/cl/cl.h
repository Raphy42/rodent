#pragma once

#include <stdlib.h>

#ifdef WIN32
    # include <CL/cl.h>
	# include <CL/cl_gl.h>
#endif

#ifdef __LINUX__
    # include <CL/cl.h>
	# include <CL/cl_gl.h>
#endif

#ifdef __APPLE__
	#include <OpenGL/OpenGL.h>
	#include <OpenCL/cl.h>
	#include <OpenCL/cl_gl.h>
	#include <OpenCL/cl_gl_ext.h>
#endif


inline cl_context interop_context_helper(const cl_platform_id platform_id, const cl_device_id device_id) {
	cl_context_properties context_properties[] = {
            // We need to add information about the OpenGL context with
            // which we want to exchange information with the OpenCL context.
            #ifdef WIN32
            // We should first check for cl_khr_gl_sharing extension.
            CL_GL_CONTEXT_KHR , (cl_context_properties) wglGetCurrentContext() ,
            CL_WGL_HDC_KHR , (cl_context_properties) wglGetCurrentDC() ,
            #elif __linux__
            // We should first check for cl_khr_gl_sharing extension.
            CL_GL_CONTEXT_KHR , (cl_context_properties) glXGetCurrentContext() ,
            CL_GLX_DISPLAY_KHR , (cl_context_properties) glXGetCurrentDisplay() ,
            #elif __APPLE__
            // We should first check for cl_APPLE_gl_sharing extension.
            #if 0
            // This doesn't work.
            CL_GL_CONTEXT_KHR , (cl_context_properties) CGLGetCurrentContext() ,
            CL_CGL_SHAREGROUP_KHR , (cl_context_properties) CGLGetShareGroup( CGLGetCurrentContext() ) ,
            #else
            CL_CONTEXT_PROPERTY_USE_CGL_SHAREGROUP_APPLE ,
            (cl_context_properties) CGLGetShareGroup( CGLGetCurrentContext() ) ,
            #endif
            #endif
            CL_CONTEXT_PLATFORM , (cl_context_properties)platform_id ,
            0 , 0 ,
	};

    cl_int status = CL_SUCCESS;
	cl_context context = clCreateContext(context_properties, 1, &device_id, 0, 0, &status);

	return context;
}
