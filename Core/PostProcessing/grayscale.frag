#version 330 core

in vec2 v_uv;

uniform sampler2D u_screenTexture;

void main() {
    vec4 color = texture(u_screenTexture, v_uv);
    float average = 0.2126 * color.r + 0.7152 * color.g + 0.0722 * color.b;
    gl_FragColor = vec4(average, average, average, 1.0);
}
