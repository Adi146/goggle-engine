#version 430 core

layout(location = 0) in vec3 a_position;
layout(location = 1) in vec2 a_uv;
layout(location = 5) in mat4 a_instanceMatrix;

out VS_OUT {
    vec2 uv;
} vs_out;

uniform mat4 u_modelMatrix;

void main() {
    gl_Position = vec4(a_position, 1.0) * (u_modelMatrix * a_instanceMatrix);

    vs_out.uv = a_uv;
}