#version 410 core
#define MAX_POINT_LIGHTS 32
#define MAX_SPOT_LIGHTS 32

layout (std140) uniform directionalLight {
    struct {
        vec3 direction;

        vec3 ambient;
        vec3 diffuse;
        vec3 specular;

        mat4 viewProjectionMatrix;

        float distance;
        float transitionDistance;
    } u_directionalLight;
};

layout (std140) uniform pointLight {
    int u_numPointLights;
    struct {
        vec3 position;
        float linear;
        float quadratic;

        vec3 ambient;
        vec3 diffuse;
        vec3 specular;

        mat4 viewProjectionMatrix[6];
        float distance;
    } u_pointLights[MAX_POINT_LIGHTS];
};

layout (std140) uniform spotLight {
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

layout (std140) uniform camera {
    mat4 u_projectionMatrix;
    mat4 u_viewMatrix;
    vec3 u_cameraPosition;
};

uniform sampler2D u_shadowMapDirectionalLight;
uniform samplerCube u_shadowMapsPointLight[MAX_POINT_LIGHTS];
uniform sampler2D u_shadowMapsSpotLights[MAX_SPOT_LIGHTS];

float calculateShadowDirectionalLight(vec3 fragPos) {
    vec4 positionLightSpace = vec4(fragPos, 1.0) * u_directionalLight.viewProjectionMatrix;

    vec3 projCoords = positionLightSpace.xyz / positionLightSpace.w;
    projCoords = projCoords * 0.5 + 0.5;

    float closestDepth  = texture(u_shadowMapDirectionalLight, projCoords.xy).r;

    float currentDepth = projCoords.z;

    float cameraDistacne = length(u_cameraPosition - fragPos);
    float distance = cameraDistacne - (u_directionalLight.distance - u_directionalLight.transitionDistance);
    distance = distance / u_directionalLight.transitionDistance;
    float transitionFactor = clamp(1.0 - distance, 0.0, 1.0);

    float shadow = 0.0;
    float bias = 0.005;
    vec2 texelSize = 1.0 / textureSize(u_shadowMapDirectionalLight, 0);
    for(int x = -1; x <= 1; ++x)
    {
        for(int y = -1; y <= 1; ++y)
        {
            float pcfDepth = texture(u_shadowMapDirectionalLight, projCoords.xy + vec2(x, y) * texelSize).r;
            shadow += currentDepth - bias > pcfDepth ? transitionFactor * 1.0 : 0.0;
        }
    }

    return shadow / 9.0;
}

float calculateShadowPointLight(int pointLightIndex, vec3 fragPos) {
    vec3 fragToLight = fragPos - u_pointLights[pointLightIndex].position;
    float closestDepth = texture(u_shadowMapsPointLight[pointLightIndex], fragToLight).r;

    closestDepth *=  u_pointLights[pointLightIndex].distance;

    float currentDepth = length(fragToLight);

    float bias = 0.05;
    float shadow = currentDepth - bias > closestDepth ? 1.0 : 0.0;

    return shadow;
}

float calculateShadowSpotLight(int spotLightIndex, vec3 fragPos) {
    vec4 positionLightSpace = vec4(fragPos, 1.0) * u_spotLights[spotLightIndex].viewProjectionMatrix;

    vec3 projCoords = positionLightSpace.xyz / positionLightSpace.w;
    projCoords = projCoords * 0.5 + 0.5;

    float closestDepth  = texture(u_shadowMapsSpotLights[spotLightIndex], projCoords.xy).z;

    float currentDepth = projCoords.z;

    float shadow = 0.0;
    vec2 texelSize = 1.0 / textureSize(u_shadowMapsSpotLights[spotLightIndex], 0);
    for(int x = -1; x <= 1; ++x)
    {
        for(int y = -1; y <= 1; ++y)
        {
            float pcfDepth = texture(u_shadowMapsSpotLights[spotLightIndex], projCoords.xy + vec2(x, y) * texelSize).r;
            shadow += currentDepth > pcfDepth ? 1.0 : 0.0;
        }
    }

    return shadow / 9.0;
}
