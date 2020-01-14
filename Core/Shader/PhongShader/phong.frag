#version 410 core
#define MAX_POINT_LIGHTS 128
#define MAX_SPOT_LIGHTS 128

in vec3 v_position;
in vec3 v_normal;

struct Material {
    vec3 diffuse;
    vec3 specular;
    vec3 emissive;
    float shininess;
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

uniform mat4 u_viewMatrix;

uniform Material u_material;

uniform DirectionalLight u_directionalLight;
uniform PointLight u_pointLights[MAX_POINT_LIGHTS];
uniform SpotLight u_spotLights[MAX_SPOT_LIGHTS];
uniform int u_numPointLights;
uniform int u_numSpotLights;

layout(location = 0) out vec4 f_color;

void calculateDirectionalLight(in vec3 viewVector, in vec3 normalVector, inout vec3 ambientColor, inout vec3 diffuseColor, inout vec3 specularColor) {
    vec3 lightDirection = vec3(vec4(u_directionalLight.direction, 1.0f) * transpose(inverse(u_viewMatrix)));

    vec3 lightVector = normalize(-lightDirection);
    vec3 reflectionVector = reflect(lightDirection, normalVector);

    ambientColor += u_directionalLight.ambient * u_material.diffuse;
    diffuseColor += u_directionalLight.diffuse * max(dot(normalVector, lightVector), 0.0f) * u_material.diffuse;
    specularColor += u_directionalLight.specular * pow(max(dot(reflectionVector, viewVector), 0.00001f), u_material.shininess) * u_material.specular;
}

void calculatePointLight(in vec3 viewVector, in vec3 normalVector, inout vec3 ambientColor, inout vec3 diffuseColor, inout vec3 specularColor) {
    for (int i = 0; i < u_numPointLights; i++){
        vec3 lightPosition = vec3(vec4(u_pointLights[i].position, 1.0) * u_viewMatrix);

        vec3 lightVector = normalize(lightPosition - v_position);
        vec3 reflectionVector = reflect(-lightVector, normalVector);

        float distance = length(lightPosition - v_position);
        float attenuation = 1.0 / ((1.0) + (u_pointLights[i].linear * distance) + (u_pointLights[i].quadratic * pow(distance, 2)));

        ambientColor += attenuation *  u_pointLights[i].ambient * u_material.diffuse;
        diffuseColor += attenuation * u_pointLights[i].diffuse * max(dot(normalVector, lightVector), 0.0f) * u_material.diffuse;
        specularColor += attenuation * u_pointLights[i].specular * pow(max(dot(reflectionVector, viewVector), 0.00001f), u_material.shininess) * u_material.specular;
    }
}

void calculateSpotLight(in vec3 viewVector, in vec3 normalVector, inout vec3 ambientColor, inout vec3 diffuseColor, inout vec3 specularColor) {
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

            ambientColor += attenuation * u_spotLights[i].ambient * u_material.diffuse;
            diffuseColor += attenuation * intensity * u_spotLights[i].diffuse * max(dot(normalVector, lightVector), 0.0f) * u_material.diffuse;
            specularColor += attenuation * intensity * u_spotLights[i].specular * pow(max(dot(reflectionVector, viewVector), 0.00001f), u_material.shininess) * u_material.specular;
        } else {
            ambientColor += attenuation * u_spotLights[i].ambient * u_material.diffuse;
        }
    }
}

void main() {
    vec3 view = normalize(-v_position);
    vec3 normal = normalize(v_normal);

    vec3 ambientColor = vec3(0.0, 0.0, 0.0);
    vec3 diffuseColor = vec3(0.0, 0.0, 0.0);
    vec3 specularColor = vec3(0.0, 0.0, 0.0);

    calculateDirectionalLight(view, normal, ambientColor, diffuseColor, specularColor);
    calculatePointLight(view, normal, ambientColor, diffuseColor, specularColor);
    calculateSpotLight(view, normal, ambientColor, diffuseColor, specularColor);

    f_color = vec4(ambientColor + diffuseColor + specularColor + u_material.emissive, 1.0f);
}

