#version 330 core

in vec2 v_uv;

uniform sampler2D u_screenTexture;

void main() {
    gl_FragColor = texture(u_screenTexture, v_uv);
}
