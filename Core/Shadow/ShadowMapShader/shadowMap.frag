#version 410 core

struct MaterialColor {
    vec4 diffuse;
    vec3 specular;
    vec3 emissive;
};

vec4 GetDiffuseColor();

void main() {
    vec4 diffuse = GetDiffuseColor();
    if (diffuse.a < 0.5) {
        discard;
    }

    // gl_FragDepth = gl_FragCoord.z;
}
