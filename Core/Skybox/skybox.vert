#version 410 core

layout(location = 0) in vec3 a_position;
layout(location = 1) in vec3 a_normal;
layout(location = 2) in vec2 a_uv;
layout(location = 3) in vec3 a_tangent;

layout (std140) uniform Camera {
    mat4 u_projectionMatrix;
    mat4 u_viewMatrix;
};

out vec3 v_uv;

void main()
{
    v_uv = a_position;
    gl_Position = u_projectionMatrix * u_viewMatrix * vec4(a_position, 1.0);
}