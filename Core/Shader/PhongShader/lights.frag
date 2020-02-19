#version 410 core
#define MAX_POINT_LIGHTS 64
#define MAX_SPOT_LIGHTS 64

in vec3 v_position;

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
};

struct PointLight{
    vec3 position;

    vec3 ambient;
    vec3 diffuse;
    vec3 specular;

    float linear;
    float quadratic;
};

struct SpotLight{
    vec3 position;
    vec3 direction;

    float innerCone;
    float outerCone;

    vec3 ambient;
    vec3 diffuse;
    vec3 specular;

    float linear;
    float quadratic;
};

layout (std140) uniform Camera {
    mat4 u_projectionMatrix;
    mat4 u_viewMatrix;
};

layout (std140) uniform directionalLight {
    DirectionalLight u_directionalLight;
};

layout (std140) uniform pointLight {
    int u_numPointLights;
    PointLight u_pointLights[MAX_POINT_LIGHTS];
};

uniform SpotLight u_spotLights[MAX_SPOT_LIGHTS];
uniform int u_numSpotLights;

vec3 calculateDirectionalLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color, in float shininess) {
    vec3 lightDirection = vec3(vec4(u_directionalLight.direction, 1.0f) * transpose(inverse(u_viewMatrix)));

    vec3 lightVector = normalize(-lightDirection);
    vec3 reflectionVector = reflect(lightDirection, normalVector);

    vec3 ambientColor = u_directionalLight.ambient * color.diffuse.rgb;
    vec3 diffuseColor = u_directionalLight.diffuse.rgb * max(dot(normalVector, lightVector), 0.0f) * color.diffuse.rgb;
    vec3 specularColor = u_directionalLight.specular * pow(max(dot(reflectionVector, viewVector), 0.00001f), shininess) * color.specular;

    return (ambientColor + diffuseColor + specularColor);
}

vec3 calculatePointLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color, in float shininess) {
    vec3 ambientColor = vec3(0.0, 0.0, 0.0);
    vec3 diffuseColor = vec3(0.0, 0.0, 0.0);
    vec3 specularColor = vec3(0.0, 0.0, 0.0);

    for (int i = 0; i < u_numPointLights; i++){
        vec3 lightPosition = vec3(vec4(u_pointLights[i].position, 1.0) * u_viewMatrix);

        vec3 lightVector = normalize(lightPosition - v_position);
        vec3 reflectionVector = reflect(-lightVector, normalVector);

        float distance = length(lightPosition - v_position);
        float attenuation = 1.0 / ((1.0) + (u_pointLights[i].linear * distance) + (u_pointLights[i].quadratic * pow(distance, 2)));

        ambientColor += attenuation *  u_pointLights[i].ambient * color.diffuse.rgb;
        diffuseColor += attenuation * u_pointLights[i].diffuse.rgb * max(dot(normalVector, lightVector), 0.0f) * color.diffuse.rgb;
        specularColor += attenuation * u_pointLights[i].specular * pow(max(dot(reflectionVector, viewVector), 0.00001f), shininess) * color.specular;
    }

    return (ambientColor + diffuseColor + specularColor);
}

vec3 calculateSpotLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color, in float shininess) {
    vec3 ambientColor = vec3(0.0, 0.0, 0.0);
    vec3 diffuseColor = vec3(0.0, 0.0, 0.0);
    vec3 specularColor = vec3(0.0, 0.0, 0.0);

    for (int i = 0; i < u_numSpotLights; i++){
        vec3 lightPosition = vec3(vec4(u_spotLights[i].position, 1.0) * u_viewMatrix);
        vec3 lightDirection = vec3(vec4(u_spotLights[i].direction, 1.0f) * transpose(inverse(u_viewMatrix)));

        vec3 lightVector = normalize(lightPosition - v_position);
        vec3 reflectionVector = reflect(-lightVector, normalVector);

        float distance = length(lightPosition - v_position);
        float attenuation = 1.0 / ((1.0) + (u_spotLights[i].linear * distance) + (u_spotLights[i].quadratic * pow(distance, 2)));

        float theta = dot(lightVector, normalize(lightDirection));
        if (theta > u_spotLights[i].outerCone) {
            float epsilon = u_spotLights[i].outerCone - u_spotLights[i].innerCone;
            float intensity = clamp((theta - u_spotLights[i].outerCone) / epsilon, 0.0f, 1.0f);

            ambientColor += attenuation * u_spotLights[i].ambient * color.diffuse.rgb;
            diffuseColor += attenuation * intensity * u_spotLights[i].diffuse.rgb * max(dot(normalVector, lightVector), 0.0f) * color.diffuse.rgb;
            specularColor += attenuation * intensity * u_spotLights[i].specular * pow(max(dot(reflectionVector, viewVector), 0.00001f), shininess) * color.specular;
        } else {
            ambientColor += attenuation * u_spotLights[i].ambient * color.diffuse.rgb;
        }
    }

    return (ambientColor + diffuseColor + specularColor);
}
