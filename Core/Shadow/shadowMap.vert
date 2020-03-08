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

layout (std140) uniform directionalLight {
    DirectionalLight u_directionalLight;
};

uniform mat4 u_modelMatrix;

void main() {
    gl_Position = vec4(a_position, 1.0) * (u_modelMatrix * u_directionalLight.lightViewMatrix * u_directionalLight.lightProjectionMatrix);
}
