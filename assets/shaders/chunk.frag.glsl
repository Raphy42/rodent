#version 410

layout (location = 1) vec3 color;



out vec4 FragColor;

void main() {
    FragColor = vec4(color, 1.0);// + vec4(enabledFaces / 1024.0);
}