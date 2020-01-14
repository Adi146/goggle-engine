#version 410 core

in vec3 v_position;
in vec3 v_normal;
in vec2 v_textureCoordinate;

layout(location = 0) out vec4 f_color;

void main() {
    f_color = vec4(normalize(v_normal), 1.0);
}