#version 410 core

in vec2 v_uv;

uniform sampler2D u_screenTexture;

void main() {
    gl_FragColor = vec4(vec3(1.0 - texture(u_screenTexture, v_uv)), 1.0);
}
