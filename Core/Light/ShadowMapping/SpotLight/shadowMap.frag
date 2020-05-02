#version 430 core
#define MAX_SPOT_LIGHTS 32

in VS_OUT {
    vec2 uv;
} fs_in;

vec4 GetDiffuseColor(vec2 uv);

void main() {
    vec4 diffuse = GetDiffuseColor(fs_in.uv);
    if (diffuse.a < 0.5) {
        discard;
    }

    // gl_FragDepth = gl_FragCoord.z;
}