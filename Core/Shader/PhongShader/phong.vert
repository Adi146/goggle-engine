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
    vec4 positionLightSpace;
} vs_out;

struct DirectionalLight {
    vec3 direction;

    vec3 ambient;
    vec3 diffuse;
    vec3 specular;

    mat4 viewProjectionMatrix;
};

layout (std140) uniform camera {
    mat4 u_projectionMatrix;
    mat4 u_viewMatrix;
    vec3 u_cameraPosition;
};

layout (std140) uniform directionalLight {
    DirectionalLight u_directionalLight;
};

uniform mat4 u_modelMatrix;

void main() {
    gl_Position = vec4(a_position, 1.0) * (u_modelMatrix * u_viewMatrix * u_projectionMatrix);

    mat3 normalMatrix = mat3(transpose(inverse(u_modelMatrix)));
    vec3 normal = normalize(a_normal * normalMatrix);
    vec3 tangent = normalize(a_tangent * normalMatrix);
    //Reorthogonalization with Gram–Schmidt process
    tangent = normalize(tangent - dot(tangent, normal) * normal);
    vec3 biTangent = normalize(cross(normal, tangent) * normalMatrix);

    vs_out.position = vec3(vec4(a_position, 1.0) * u_modelMatrix);
    vs_out.normal = normal;
    vs_out.uv = a_uv;
    vs_out.tbn = transpose(mat3(tangent, biTangent, normal));
    vs_out.positionLightSpace = vec4(a_position, 1.0) * (u_modelMatrix * u_directionalLight.viewProjectionMatrix);
}