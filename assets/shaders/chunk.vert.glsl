#version 410

uniform mat4 mvp;

lowp vec3 vColor;
vec3 vPos;
int vEnabledFaces;

flat out lowp vec3 gColor;
flat out int gEnabledFaces;

void main() {
    gl_Position = mvp * vec4(vPos, 1.0);

    gColor = vColor;
    gEnabledFaces = vEnabledFaces;
}