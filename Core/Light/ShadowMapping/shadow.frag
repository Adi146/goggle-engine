#version 410 core

in vec3 v_normal;
in vec4 v_positionLightSpace;

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

uniform sampler2D u_shadowMap;

float calculateShadow() {
    vec3 projCoords = v_positionLightSpace.xyz / v_positionLightSpace.w;
    projCoords = projCoords * 0.5 + 0.5;

    float closestDepth  = texture(u_shadowMap, projCoords.xy).z;

    float currentDepth = projCoords.z;

    float shadow = 0.0;
    vec2 texelSize = 1.0 / textureSize(u_shadowMap, 0);
    for(int x = -1; x <= 1; ++x)
    {
        for(int y = -1; y <= 1; ++y)
        {
            float pcfDepth = texture(u_shadowMap, projCoords.xy + vec2(x, y) * texelSize).r;
            shadow += currentDepth > pcfDepth ? 1.0 : 0.0;
        }
    }

    return shadow / 9.0;
}
