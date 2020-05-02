#version 430 core

layout(location = 0) in vec3 a_position;

out VS_OUT {
    vec3 uv;
} vs_out;

layout (std140) uniform camera {
    mat4 u_projectionMatrix;
    mat4 u_viewMatrix;
    vec3 u_cameraPosition;
};

void main()
{
    //remove translation from u_viewMatrix
    mat4 viewMatrix = transpose(inverse(u_viewMatrix));

    gl_Position = (vec4(a_position, 1.0) * (viewMatrix * u_projectionMatrix)).xyww;
    vs_out.uv = a_position;
}