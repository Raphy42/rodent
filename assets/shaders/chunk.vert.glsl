#version 410

layout (location = 0) in vec3 pos;
layout (location = 1) in vec3 color;
layout (location = 2) in int enabledFaces;

uniform mat4 mvp;

out vec3 gColor;
out int gEnabledFaces;

void main() {
    gl_Positionw = mvp * vec4(vPos, 1.0);

    gColor = vColor;
    gEnabledFaces = vEnabledFaces;
}