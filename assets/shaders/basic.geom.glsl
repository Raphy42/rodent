#version 330 core

layout (points) in;
layout (triangle_strip, max_vertices = 24) out;

in VS_OUT {
    vec3 color;
} gs_in[];

uniform mat4 mvp;

out vec3 color;

const float size = .3;
const vec3 lightDirection = normalize(vec3(0.4, -1, 0.8));

void createVertex(vec3 offset, vec3 normal) {
    vec4 actualOffset = vec4(offset * size, 0.0);
    vec4 worldPosition = gl_in[0].gl_Position + actualOffset;
    gl_Position = mvp * worldPosition;
    float brightness = max(dot(-lightDirection, normal), 0.3);
    EmitVertex();
}

void main() {
    color = gs_in[0].color;

    vec3 faceNormal = vec3(0.0, 0.0, 1.0);
	createVertex(vec3(-1.0, 1.0, 1.0), faceNormal);
	createVertex(vec3(-1.0, -1.0, 1.0), faceNormal);
	createVertex(vec3(1.0, 1.0, 1.0), faceNormal);
	createVertex(vec3(1.0, -1.0, 1.0), faceNormal);

	EndPrimitive();

	faceNormal = vec3(1.0, 0.0, 0.0);
	createVertex(vec3(1.0, 1.0, 1.0), faceNormal);
	createVertex(vec3(1.0, -1.0, 1.0), faceNormal);
	createVertex(vec3(1.0, 1.0, -1.0), faceNormal);
	createVertex(vec3(1.0, -1.0, -1.0), faceNormal);

	EndPrimitive();

	faceNormal = vec3(0.0, 0.0, -1.0);
	createVertex(vec3(1.0, 1.0, -1.0), faceNormal);
	createVertex(vec3(1.0, -1.0, -1.0), faceNormal);
	createVertex(vec3(-1.0, 1.0, -1.0), faceNormal);
	createVertex(vec3(-1.0, -1.0, -1.0), faceNormal);

	EndPrimitive();

	faceNormal = vec3(-1.0, 0.0, 0.0);
	createVertex(vec3(-1.0, 1.0, -1.0), faceNormal);
	createVertex(vec3(-1.0, -1.0, -1.0), faceNormal);
	createVertex(vec3(-1.0, 1.0, 1.0), faceNormal);
	createVertex(vec3(-1.0, -1.0, 1.0), faceNormal);

	EndPrimitive();

	faceNormal = vec3(0.0, 1.0, 0.0);
	createVertex(vec3(1.0, 1.0, 1.0), faceNormal);
	createVertex(vec3(1.0, 1.0, -1.0), faceNormal);
	createVertex(vec3(-1.0, 1.0, 1.0), faceNormal);
	createVertex(vec3(-1.0, 1.0, -1.0), faceNormal);

	EndPrimitive();

	faceNormal = vec3(0.0, -1.0, 0.0);
	createVertex(vec3(-1.0, -1.0, 1.0), faceNormal);
	createVertex(vec3(-1.0, -1.0, -1.0), faceNormal);
	createVertex(vec3(1.0, -1.0, 1.0), faceNormal);
	createVertex(vec3(1.0, -1.0, -1.0), faceNormal);

	EndPrimitive();
}