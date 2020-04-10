#version 410 core

layout(location = 0) in vec3 a_position;
layout(location = 1) in vec3 a_normal;
layout(location = 2) in vec2 a_uv;
layout(location = 3) in vec3 a_tangent;

out VS_OUT {
    vec3 position;
    vec3 normal;
    vec2 uv;
    mat3 tbn;
} vs_out;

layout (std140) uniform camera {
    mat4 u_projectionMatrix;
    mat4 u_viewMatrix;
    vec3 u_cameraPosition;
};

uniform mat4 u_modelMatrix;
uniform mat3 u_normalMatrix;

void main() {
    gl_Position = vec4(a_position, 1.0) * (u_modelMatrix * u_viewMatrix * u_projectionMatrix);

    vec3 normal = normalize(a_normal * u_normalMatrix);
    vec3 tangent = normalize(a_tangent * u_normalMatrix);
    //Reorthogonalization with Gram–Schmidt process
    tangent = normalize(tangent - dot(tangent, normal) * normal);
    vec3 biTangent = normalize(cross(normal, tangent) * u_normalMatrix);

    vs_out.position = vec3(vec4(a_position, 1.0) * u_modelMatrix);
    vs_out.normal = normal;
    vs_out.uv = a_uv;
    vs_out.tbn = transpose(mat3(tangent, biTangent, normal));
}