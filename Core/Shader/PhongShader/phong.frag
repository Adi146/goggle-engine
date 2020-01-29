#version 410 core

in vec3 v_position;
in vec2 v_uv;

struct MaterialColor {
    vec3 diffuse;
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

uniform sampler2D u_depthMap;

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

    f_color = vec4(fragmentColor + color.emissive, 1.0f) ;

    //float depthValue = texture(u_depthMap, v_uv).r;
    //f_color = vec4(vec2(depthValue), 1.0, 1.0);
}

