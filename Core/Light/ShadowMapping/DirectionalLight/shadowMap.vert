#version 410 core
#define MAX_DIRECTIONAL_LIGHTS 32

layout(location = 0) in vec3 a_position;
layout(location = 1) in vec2 a_uv;
layout(location = 5) in mat4 a_instanceMatrix;

out VS_OUT {
    vec2 uv;
} vs_out;

layout (std140) uniform directionalLight {
    int u_numDirectionalLights;
    struct {
        vec3 direction;

        vec3 ambient;
        vec3 diffuse;
        vec3 specular;

        mat4 viewProjectionMatrix;

        float distance;
        float transitionDistance;
    } u_directionalLights[MAX_DIRECTIONAL_LIGHTS];
};

uniform mat4 u_modelMatrix;
uniform int u_lightIndex;

void main() {
    gl_Position = vec4(a_position, 1.0) * (u_modelMatrix * a_instanceMatrix * u_directionalLights[u_lightIndex].viewProjectionMatrix);

    vs_out.uv = a_uv;
}
