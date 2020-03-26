#version 410 core

in VS_OUT {
    vec2 uv;
} fs_in;

uniform sampler2D u_screenTexture;

void main() {
    gl_FragColor = vec4(vec3(1.0 - texture(u_screenTexture, fs_in.uv)), 1.0);
}
