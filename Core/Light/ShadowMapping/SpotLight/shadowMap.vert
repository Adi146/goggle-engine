#version 430 core
#define MAX_SPOT_LIGHTS 32

layout(location = 0) in vec3 a_position;
layout(location = 1) in vec2 a_uv;
layout(location = 5) in mat4 a_instanceMatrix;

out VS_OUT {
    vec2 uv;
} vs_out;

layout (std140, binding = 3) uniform spotLight {
    int u_numSpotLights;
    struct {
        vec3 position;
        float linear;
        float quadratic;

        vec3 ambient;
        vec3 diffuse;
        vec3 specular;

        vec3 direction;

        float innerCone;
        float outerCone;

        mat4 viewProjectionMatrix;
    } u_spotLights[MAX_SPOT_LIGHTS];
};

uniform mat4 u_modelMatrix;
uniform int u_lightIndex;

void main() {
    gl_Position = vec4(a_position, 1.0) * (u_modelMatrix * a_instanceMatrix * u_spotLights[u_lightIndex].viewProjectionMatrix);

    vs_out.uv = a_uv;
}