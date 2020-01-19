#version 410 core
#define MAX_POINT_LIGHTS 128
#define MAX_SPOT_LIGHTS 128

#define MAX_TEXTURES 24

in vec3 v_position;
in vec3 v_normal;
in vec2 v_uv;

struct MaterialColor {
    vec3 diffuse;
    vec3 specular;
    vec3 emissive;
};

struct Material {
    MaterialColor baseColor;

    float shininess;

    sampler2D texturesDiffuse[MAX_TEXTURES];
    sampler2D texturesNormals[MAX_TEXTURES];

    int numTextureDiffuse;
    int numTextureNormals;
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

vec3 calculateDirectionalLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color) {
    vec3 lightDirection = vec3(vec4(u_directionalLight.direction, 1.0f) * transpose(inverse(u_viewMatrix)));

    vec3 lightVector = normalize(-lightDirection);
    vec3 reflectionVector = reflect(lightDirection, normalVector);

    vec3 ambientColor = u_directionalLight.ambient * color.diffuse;
    vec3 diffuseColor = u_directionalLight.diffuse * max(dot(normalVector, lightVector), 0.0f) * color.diffuse;
    vec3 specularColor = u_directionalLight.specular * pow(max(dot(reflectionVector, viewVector), 0.00001f), u_material.shininess) * color.specular;

    return (ambientColor + diffuseColor + specularColor);
}

vec3 calculatePointLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color) {
    vec3 ambientColor = vec3(0.0, 0.0, 0.0);
    vec3 diffuseColor = vec3(0.0, 0.0, 0.0);
    vec3 specularColor = vec3(0.0, 0.0, 0.0);

    for (int i = 0; i < u_numPointLights; i++){
        vec3 lightPosition = vec3(vec4(u_pointLights[i].position, 1.0) * u_viewMatrix);

        vec3 lightVector = normalize(lightPosition - v_position);
        vec3 reflectionVector = reflect(-lightVector, normalVector);

        float distance = length(lightPosition - v_position);
        float attenuation = 1.0 / ((1.0) + (u_pointLights[i].linear * distance) + (u_pointLights[i].quadratic * pow(distance, 2)));

        ambientColor += attenuation *  u_pointLights[i].ambient * color.diffuse;
        diffuseColor += attenuation * u_pointLights[i].diffuse * max(dot(normalVector, lightVector), 0.0f) * color.diffuse;
        specularColor += attenuation * u_pointLights[i].specular * pow(max(dot(reflectionVector, viewVector), 0.00001f), u_material.shininess) * color.specular;
    }

    return (ambientColor + diffuseColor + specularColor);
}

vec3 calculateSpotLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color) {
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

            ambientColor += attenuation * u_spotLights[i].ambient * color.diffuse;
            diffuseColor += attenuation * intensity * u_spotLights[i].diffuse * max(dot(normalVector, lightVector), 0.0f) * color.diffuse;
            specularColor += attenuation * intensity * u_spotLights[i].specular * pow(max(dot(reflectionVector, viewVector), 0.00001f), u_material.shininess) * color.specular;
        } else {
            ambientColor += attenuation * u_spotLights[i].ambient * color.diffuse;
        }
    }

    return (ambientColor + diffuseColor + specularColor);
}

void main() {
    vec3 view = normalize(-v_position);
    vec3 normal = normalize(v_normal);

    vec3 ambientColor = vec3(0.0, 0.0, 0.0);
    vec3 diffuseColor = vec3(0.0, 0.0, 0.0);
    vec3 specularColor = vec3(0.0, 0.0, 0.0);

    MaterialColor color = u_material.baseColor;
    if (u_material.numTextureDiffuse > 0) {
        vec4 diffuse = vec4(0.0, 0.0, 0.0, 0.0);
        for (int i = 0; i < u_material.numTextureDiffuse; i++){
            diffuse += texture(u_material.texturesDiffuse[i], v_uv);
        }
        if (diffuse.w < 0.9){
            discard;
        }
        color.diffuse = vec3(diffuse);
    }

    vec3 fragmentColor = vec3(0.0, 0.0, 0.0);
    fragmentColor += calculateDirectionalLight(view, normal, color);
    fragmentColor += calculatePointLight(view, normal, color);
    fragmentColor += calculateSpotLight(view, normal, color);

    f_color = vec4(fragmentColor + u_material.baseColor.emissive, 1.0f) ;
}

