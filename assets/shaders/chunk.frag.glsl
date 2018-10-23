#version 410

flat in lowp vec3 fColor;
//varying int fEnabledFaces;

out vec4 FragColor;

void main() {
    FragColor = vec4(fColor, 1.0);// + vec4(enabledFaces / 1024.0);
}