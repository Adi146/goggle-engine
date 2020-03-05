#version 330 core

in vec3 v_position;

struct MaterialColor {
    vec4 diffuse;
    vec3 specular;
    vec3 emissive;
};

layout(location = 0) out vec4 f_color;

MaterialColor GetMaterialColor();
vec3 GetNormalVector();
float GetShininess();

vec3 calculateDirectionalLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color, in float shininess);
vec3 calculatePointLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color, in float shininess);
vec3 calculateSpotLight(in vec3 viewVector, in vec3 normalVector, in MaterialColor color, in float shininess);

void main() {
    MaterialColor color = GetMaterialColor();
    float shininess = GetShininess();
    vec3 normal = GetNormalVector();

    // calculate lights
    vec3 view = normalize(-v_position);
    vec3 fragmentColor = vec3(0.0, 0.0, 0.0);
    fragmentColor += calculateDirectionalLight(view, normal, color, shininess);
    fragmentColor += calculatePointLight(view, normal, color, shininess);
    fragmentColor += calculateSpotLight(view, normal, color, shininess);

    f_color = vec4(fragmentColor + color.emissive, color.diffuse.a) ;
}

