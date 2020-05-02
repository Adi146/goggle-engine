#version 430 core

in VS_OUT {
    vec2 uv;
} fs_in;

layout(location = 0) out vec4 FragColor;

uniform sampler2D u_screenTexture;

void main() {
    FragColor = vec4(vec3(1.0 - texture(u_screenTexture, fs_in.uv)), 1.0);
}
