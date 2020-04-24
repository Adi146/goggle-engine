#version 410 core
#define MAX_DIRECTIONAL_LIGHTS 32
#define MAX_POINT_LIGHTS 32
#define MAX_SPOT_LIGHTS 32

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

struct MaterialColor {
    vec4 diffuse;
    vec3 specular;
    vec3 emissive;
};

uniform bool u_directionalLightIsSet;
uniform bool u_pointLightIsSet;
uniform bool u_spotLightIsSet;

float calculateShadowDirectionalLight(int directionalLightIndex, vec3 fragPos);
float calculateShadowPointLight(int pointLightIndex, vec3 fragPos);
float calculateShadowSpotLight(int spotLightIndex, vec3 fragPos);

vec3 calculateDirectionalLight(vec3 fragPos, vec3 viewVector, vec3 normalVector, MaterialColor color, float shininess) {
    vec3 ambientColor = vec3(0.0, 0.0, 0.0);
    vec3 diffuseColor = vec3(0.0, 0.0, 0.0);
    vec3 specularColor = vec3(0.0, 0.0, 0.0);

    if (u_directionalLightIsSet) {
        for (int i = 0; i < u_numDirectionalLights; i++) {
            vec3 lightDir = normalize(-u_directionalLights[i].direction);
            float diff = max(dot(normalVector, lightDir), 0.0);

            vec3 reflectDir = reflect(-lightDir, normalVector);
            float spec = pow(max(dot(viewVector, reflectDir), 0.0), shininess);

            float shadow = calculateShadowDirectionalLight(i, fragPos);

            ambientColor += u_directionalLights[i].ambient * color.diffuse.rgb;
            diffuseColor += (1.0 - shadow) * u_directionalLights[i].diffuse * diff * color.diffuse.rgb;
            specularColor += (1.0 - shadow) * u_directionalLights[i].specular * spec * color.specular;
        }
    }

    return (ambientColor + diffuseColor + specularColor);
}

vec3 calculatePointLight(vec3 fragPos, vec3 viewVector, vec3 normalVector, MaterialColor color, float shininess) {
    vec3 ambientColor = vec3(0.0, 0.0, 0.0);
    vec3 diffuseColor = vec3(0.0, 0.0, 0.0);
    vec3 specularColor = vec3(0.0, 0.0, 0.0);

    if (u_pointLightIsSet) {
        for (int i = 0; i < u_numPointLights; i++){
            vec3 lightDir = normalize(u_pointLights[i].position - fragPos);
            float diff = max(dot(normalVector, lightDir), 0.0);

            vec3 reflectDir = reflect(-lightDir, normalVector);
            float spec = pow(max(dot(viewVector, reflectDir), 0.0), shininess);

            float distance = length(u_pointLights[i].position - fragPos);
            float attenuation = 1.0 / ((1.0) + (u_pointLights[i].linear * distance) + (u_pointLights[i].quadratic * pow(distance, 2)));

            float shadow = calculateShadowPointLight(i, fragPos);

            ambientColor += attenuation *  u_pointLights[i].ambient * color.diffuse.rgb;
            diffuseColor += attenuation * (1 - shadow) * u_pointLights[i].diffuse.rgb * diff * color.diffuse.rgb;
            specularColor += attenuation * (1 - shadow) * u_pointLights[i].specular * spec * color.specular;
        }
    }

    return (ambientColor + diffuseColor + specularColor);
}

vec3 calculateSpotLight(vec3 fragPos, vec3 viewVector, vec3 normalVector, MaterialColor color, float shininess) {
    vec3 ambientColor = vec3(0.0, 0.0, 0.0);
    vec3 diffuseColor = vec3(0.0, 0.0, 0.0);
    vec3 specularColor = vec3(0.0, 0.0, 0.0);

    if (u_spotLightIsSet) {
        for (int i = 0; i < u_numSpotLights; i++){
            vec3 lightDir = normalize(u_spotLights[i].position - fragPos);
            float diff = max(dot(normalVector, lightDir), 0.0);

            vec3 reflectDir = reflect(-lightDir, normalVector);
            float spec = pow(max(dot(viewVector, reflectDir), 0.0), shininess);

            float distance = length(u_spotLights[i].position - fragPos);
            float attenuation = 1.0 / ((1.0) + (u_spotLights[i].linear * distance) + (u_spotLights[i].quadratic * pow(distance, 2)));

            float theta = dot(lightDir, -normalize(u_spotLights[i].direction));
            float epsilon = u_spotLights[i].innerCone - u_spotLights[i].outerCone;
            float intensity = clamp((theta - u_spotLights[i].outerCone) / epsilon, 0.0f, 1.0f);

            float shadow = calculateShadowSpotLight(i, fragPos);

            ambientColor += attenuation * u_spotLights[i].ambient * color.diffuse.rgb;
            diffuseColor += attenuation * intensity * (1 - shadow) * u_spotLights[i].diffuse * diff * color.diffuse.rgb;
            specularColor += attenuation * intensity * (1 - shadow) * u_spotLights[i].specular * spec* color.specular;
        }
    }

    return (ambientColor + diffuseColor + specularColor);
}
