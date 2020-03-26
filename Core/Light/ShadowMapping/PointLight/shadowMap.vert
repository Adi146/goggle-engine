#version 410 core

layout(location = 0) in vec3 a_position;
layout(location = 1) in vec3 a_normal;
layout(location = 2) in vec2 a_uv;
layout(location = 3) in vec3 a_tangent;

out VS_OUT {
    vec2 uv;
} vs_out;

uniform mat4 u_modelMatrix;

void main() {
    gl_Position = vec4(a_position, 1.0) * u_modelMatrix;

    vs_out.uv = a_uv;
}