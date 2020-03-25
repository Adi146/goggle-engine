#version 410 core
#define MAX_POINT_LIGHTS 32

in vec4 FragPos;

struct PointLight{
    vec3 position;
    float linear;
    float quadratic;

    vec3 ambient;
    vec3 diffuse;
    vec3 specular;

    mat4 viewProjectionMatrix[6];
    float distance;
};

layout (std140) uniform pointLight {
    int u_numPointLights;
    PointLight u_pointLights[MAX_POINT_LIGHTS];
};

uniform int u_lightIndex;

vec4 GetDiffuseColor();

void main() {
    vec4 diffuse = GetDiffuseColor();
    if (diffuse.a < 0.5) {
        discard;
    }

    float lightDistance = length(FragPos.xyz - u_pointLights[u_lightIndex].position);

    lightDistance = lightDistance / u_pointLights[u_lightIndex].distance;

    gl_FragDepth = lightDistance;
}
