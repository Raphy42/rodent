#version 330 core

layout (location = 0) in vec3 pos;
layout (location = 1) in vec3 color;

uniform mat4 mvp;

out VS_OUT {
	vec3 color;
} vs_out;

void main() {
	vs_out.color = color;
	gl_PointSize = 10;
	gl_Position = vec4(pos, 1.0);
}