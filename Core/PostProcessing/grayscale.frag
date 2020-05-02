#version 430 core

in VS_OUT {
    vec2 uv;
} fs_in;

layout(location = 0) out vec4 FragColor;

uniform sampler2D u_screenTexture;

void main() {
    vec4 color = texture(u_screenTexture, fs_in.uv);
    float average = 0.2126 * color.r + 0.7152 * color.g + 0.0722 * color.b;
    FragColor = vec4(average, average, average, 1.0);
}
