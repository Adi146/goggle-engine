#version 430 core

layout(location = 0) in vec3 a_position;
layout(location = 1) in vec2 a_uv;
layout(location = 2) in vec3 a_normal;
layout(location = 3) in vec3 a_tangent;
layout(location = 4) in vec3 a_biTangent;
layout(location = 5) in mat4 a_instanceMatrix;

out VS_OUT {
    vec3 position;
    vec3 normal;
    vec2 uv;
    mat3 tbn;
} vs_out;

layout (std140, binding = 0) uniform camera {
    mat4 u_projectionMatrix;
    mat4 u_viewMatrix;
    vec3 u_cameraPosition;
};

uniform mat4 u_modelMatrix;

void main() {
    gl_Position = vec4(a_position, 1.0) * (u_modelMatrix * a_instanceMatrix * u_viewMatrix * u_projectionMatrix);

    mat3 normalMatrix = mat3(transpose(inverse(u_modelMatrix)));
    vec3 normal = normalize(a_normal * normalMatrix);
    vec3 tangent = normalize(a_tangent * normalMatrix);
    vec3 biTangent = normalize(a_biTangent * normalMatrix);

    vs_out.position = vec3(vec4(a_position, 1.0) * u_modelMatrix * a_instanceMatrix);
    vs_out.normal = normal;
    vs_out.uv = a_uv;
    vs_out.tbn = transpose(mat3(tangent, biTangent, normal));
}