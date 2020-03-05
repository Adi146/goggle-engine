#version 330 core

layout(location = 0) in vec3 a_position;
layout(location = 1) in vec3 a_normal;
layout(location = 2) in vec2 a_uv;
layout(location = 3) in vec3 a_tangent;

layout (std140) uniform camera {
    mat4 u_projectionMatrix;
    mat4 u_viewMatrix;
};

out vec3 v_uv;

void main()
{
    //remove translation from u_viewMatrix
    mat4 viewMatrix = transpose(inverse(u_viewMatrix));

    gl_Position = (vec4(a_position, 1.0) * (viewMatrix * u_projectionMatrix)).xyww;
    v_uv = a_position;
}