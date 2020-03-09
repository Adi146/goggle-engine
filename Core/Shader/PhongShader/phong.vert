#version 410 core

layout(location = 0) in vec3 a_position;
layout(location = 1) in vec3 a_normal;
layout(location = 2) in vec2 a_uv;
layout(location = 3) in vec3 a_tangent;

struct DirectionalLight {
    vec3 direction;

    vec3 ambient;
    vec3 diffuse;
    vec3 specular;

    mat4 lightProjectionMatrix;
    mat4 lightViewMatrix;
};

layout (std140) uniform camera {
    mat4 u_projectionMatrix;
    mat4 u_viewMatrix;
};

layout (std140) uniform directionalLight {
    DirectionalLight u_directionalLight;
};

uniform mat4 u_modelMatrix;

out vec3 v_position;
out vec3 v_normal;
out vec2 v_uv;
out vec3 v_tangent;
out vec3 v_biTangent;
out vec4 v_positionLightSpace;

void main() {
    gl_Position = vec4(a_position, 1.0) * (u_modelMatrix * u_viewMatrix * u_projectionMatrix);

    mat3 invModelView = mat3(transpose(inverse(u_modelMatrix * u_viewMatrix)));
    vec3 normal = normalize(a_normal * invModelView);
    vec3 tangent = normalize(a_tangent * invModelView);
    //Reorthogonalization with Gram–Schmidt process
    tangent = normalize(tangent - dot(tangent, normal) * normal);
    vec3 biTangent = normalize(cross(normal, tangent) * invModelView);

    v_position = vec3(vec4(a_position, 1.0) * (u_modelMatrix * u_viewMatrix));
    v_normal = normal;
    v_uv = a_uv;
    v_tangent = tangent;
    v_biTangent = biTangent;
    v_positionLightSpace = vec4(a_position, 1.0) * (u_modelMatrix * u_directionalLight.lightViewMatrix * u_directionalLight.lightProjectionMatrix);
}