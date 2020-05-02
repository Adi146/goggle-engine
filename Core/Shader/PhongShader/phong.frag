#version 430 core

in VS_OUT {
    vec3 position;
    vec3 normal;
    vec2 uv;
    mat3 tbn;
} fs_in;

layout(location = 0) out vec4 FragColor;

struct MaterialColor {
    vec4 diffuse;
    vec3 specular;
    vec3 emissive;
};

layout (std140, binding = 0) uniform camera {
    mat4 u_projectionMatrix;
    mat4 u_viewMatrix;
    vec3 u_cameraPosition;
};

MaterialColor GetMaterialColor(vec2 uv);
vec3 GetNormalVector(vec3 normal, vec2 uv, mat3 tbn);
float GetShininess(vec2 uv);

vec3 calculateDirectionalLight(vec3 fragPos, vec3 viewVector, vec3 normalVector, MaterialColor color, float shininess);
vec3 calculatePointLight(vec3 fragPos, vec3 viewVector, vec3 normalVector, MaterialColor color, float shininess);
vec3 calculateSpotLight(vec3 fragPos, vec3 viewVector, vec3 normalVector, MaterialColor color, float shininess);

void main() {
    MaterialColor color = GetMaterialColor(fs_in.uv);
    float shininess = GetShininess(fs_in.uv);
    vec3 normal = GetNormalVector(fs_in.normal, fs_in.uv, fs_in.tbn);

    // calculate lights
    vec3 view = normalize(u_cameraPosition - fs_in.position);
    vec3 fragmentColor = vec3(0.0, 0.0, 0.0);
    fragmentColor += calculateDirectionalLight(fs_in.position, view, normal, color, shininess);
    fragmentColor += calculatePointLight(fs_in.position, view, normal, color, shininess);
    fragmentColor += calculateSpotLight(fs_in.position, view, normal, color, shininess);

    FragColor = vec4(fragmentColor + color.emissive, color.diffuse.a);
}

