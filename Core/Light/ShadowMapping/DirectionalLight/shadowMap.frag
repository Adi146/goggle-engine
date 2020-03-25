#version 410 core

vec4 GetDiffuseColor();

void main() {
    vec4 diffuse = GetDiffuseColor();
    if (diffuse.a < 0.5) {
        discard;
    }

    // gl_FragDepth = gl_FragCoord.z;
}
