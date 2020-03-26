#version 410 core
#define MAX_POINT_LIGHTS 32
#define MAX_SPOT_LIGHTS 32

in VS_OUT {
    vec3 position;
    vec3 normal;
    vec2 uv;
    mat3 tbn;
    vec4 positionLightSpace;
} fs_in;

struct MaterialColor {
    vec4 diffuse;
    vec3 specular;
    vec3 emissive;
};

struct DirectionalLight {
    vec3 direction;

    vec3 ambient;
    vec3 diffuse;
    vec3 specular;

    mat4 viewProjectionMatrix;
};

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

struct SpotLight{
    vec3 position;
    float linear;
    float quadratic;

    vec3 ambient;
    vec3 diffuse;
    vec3 specular;

    vec3 direction;

    float innerCone;
    float outerCone;
};

layout (std140) uniform directionalLight {
    DirectionalLight u_directionalLight;
};

layout (std140) uniform pointLight {
    int u_numPointLights;
    PointLight u_pointLights[MAX_POINT_LIGHTS];
};

layout (std140) uniform spotLight {
    uniform int u_numSpotLights;
    uniform SpotLight u_spotLights[MAX_SPOT_LIGHTS];
};

float calculateShadowDirectionalLight(vec4 positionLightSpace);
float calculateShadowPointLight(int pointLightIndex, vec3 fragPos);

vec3 calculateDirectionalLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color, in float shininess) {
    vec3 lightDir = normalize(-u_directionalLight.direction);
    float diff = max(dot(normalVector, lightDir), 0.0);

    vec3 reflectDir = reflect(-lightDir, normalVector);
    float spec = pow(max(dot(viewVector, reflectDir), 0.0), shininess);

    vec3 ambientColor = u_directionalLight.ambient * color.diffuse.rgb;
    vec3 diffuseColor = u_directionalLight.diffuse * diff * color.diffuse.rgb;
    vec3 specularColor = u_directionalLight.specular * spec * color.specular;

    float shadow = calculateShadowDirectionalLight(fs_in.positionLightSpace);

    return (ambientColor + ((1.0 - shadow) * diffuseColor + specularColor));
}

vec3 calculatePointLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color, in float shininess) {
    vec3 ambientColor = vec3(0.0, 0.0, 0.0);
    vec3 diffuseColor = vec3(0.0, 0.0, 0.0);
    vec3 specularColor = vec3(0.0, 0.0, 0.0);

    for (int i = 0; i < u_numPointLights; i++){
        vec3 lightDir = normalize(u_pointLights[i].position - fs_in.position);
        float diff = max(dot(normalVector, lightDir), 0.0);

        vec3 reflectDir = reflect(-lightDir, normalVector);
        float spec = pow(max(dot(viewVector, reflectDir), 0.0), shininess);

        float distance = length(u_pointLights[i].position - fs_in.position);
        float attenuation = 1.0 / ((1.0) + (u_pointLights[i].linear * distance) + (u_pointLights[i].quadratic * pow(distance, 2)));

        float shadow = calculateShadowPointLight(i, fs_in.position);

        ambientColor += attenuation *  u_pointLights[i].ambient * color.diffuse.rgb;
        diffuseColor += attenuation * (1 - shadow) * u_pointLights[i].diffuse.rgb * diff * color.diffuse.rgb;
        specularColor += attenuation * (1 - shadow) * u_pointLights[i].specular * spec * color.specular;
    }

    return (ambientColor + diffuseColor + specularColor);
}

vec3 calculateSpotLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color, in float shininess) {
    vec3 ambientColor = vec3(0.0, 0.0, 0.0);
    vec3 diffuseColor = vec3(0.0, 0.0, 0.0);
    vec3 specularColor = vec3(0.0, 0.0, 0.0);

    for (int i = 0; i < u_numSpotLights; i++){
        vec3 lightDir = normalize(u_spotLights[i].position - fs_in.position);
        float diff = max(dot(normalVector, lightDir), 0.0);

        vec3 reflectDir = reflect(-lightDir, normalVector);
        float spec = pow(max(dot(viewVector, reflectDir), 0.0), shininess);

        float distance = length(u_spotLights[i].position - fs_in.position);
        float attenuation = 1.0 / ((1.0) + (u_spotLights[i].linear * distance) + (u_spotLights[i].quadratic * pow(distance, 2)));

        float theta = dot(lightDir, normalize(u_spotLights[i].direction));
        float epsilon = u_spotLights[i].outerCone - u_spotLights[i].innerCone;
        float intensity = clamp((theta - u_spotLights[i].outerCone) / epsilon, 0.0f, 1.0f);

        ambientColor += attenuation * intensity * u_spotLights[i].ambient * color.diffuse.rgb;
        diffuseColor += attenuation * intensity * u_spotLights[i].diffuse * diff * color.diffuse.rgb;
        specularColor += attenuation * intensity * u_spotLights[i].specular * spec* color.specular;
    }

    return (ambientColor + diffuseColor + specularColor);
}
