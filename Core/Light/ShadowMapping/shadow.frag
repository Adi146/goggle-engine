#version 410 core
#define MAX_POINT_LIGHTS 32

in vec3 v_position;
in vec3 v_normal;
in vec4 v_positionLightSpace;

uniform sampler2D u_shadowMapDirectionalLight;
uniform samplerCube u_shadowMapsPointLight[MAX_POINT_LIGHTS];

struct PointLight{
    vec3 position;
    float linear;
    float quadratic;

    vec3 ambient;
    vec3 diffuse;
    vec3 specular;

    mat4 viewProjectionMatrix[6];
};

layout (std140) uniform pointLight {
    int u_numPointLights;
    PointLight u_pointLights[MAX_POINT_LIGHTS];
};

float calculateShadowDirectionalLight() {
    vec3 projCoords = v_positionLightSpace.xyz / v_positionLightSpace.w;
    projCoords = projCoords * 0.5 + 0.5;

    float closestDepth  = texture(u_shadowMapDirectionalLight, projCoords.xy).z;

    float currentDepth = projCoords.z;

    float shadow = 0.0;
    vec2 texelSize = 1.0 / textureSize(u_shadowMapDirectionalLight, 0);
    for(int x = -1; x <= 1; ++x)
    {
        for(int y = -1; y <= 1; ++y)
        {
            float pcfDepth = texture(u_shadowMapDirectionalLight, projCoords.xy + vec2(x, y) * texelSize).r;
            shadow += currentDepth > pcfDepth ? 1.0 : 0.0;
        }
    }

    return shadow / 9.0;
}

float calculateShadowPointLight(int pointLightIndex) {
    vec3 fragToLight = v_position - u_pointLights[pointLightIndex].position;
    float closestDepth = texture(u_shadowMapsPointLight[pointLightIndex], fragToLight).r;

    closestDepth *= 3250;

    float currentDepth = length(fragToLight);

    float bias = 0.05; // we use a much larger bias since depth is now in [near_plane, far_plane] range
    float shadow = currentDepth - bias > closestDepth ? 1.0 : 0.0;

    return shadow;
}
