#version 410 core

layout(location = 0) in vec3 a_position;
layout(location = 1) in vec3 a_normal;
layout(location = 2) in vec2 a_uv;

uniform mat4 u_projectionMatrix;
uniform mat4 u_viewMatrix;

uniform mat4 u_modelMatrix;

out vec3 v_normal;
out vec2 v_uv;

void main() {
    gl_Position = vec4(a_position, 1.0) * (u_modelMatrix * u_viewMatrix * u_projectionMatrix);

    v_normal = a_normal * mat3(transpose(inverse(u_modelMatrix * u_viewMatrix)));
    v_uv = a_uv;
}