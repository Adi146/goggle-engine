#version 330 core

in vec2 v_uv;

uniform sampler2D u_screenTexture;
uniform float u_kernelOffset;
uniform float u_kernel[9];

void main() {
    vec2 offsets[9] = vec2[](
        vec2(-u_kernelOffset, u_kernelOffset),  // top-left
        vec2( 0.0f, u_kernelOffset),            // top-center
        vec2( u_kernelOffset, u_kernelOffset),  // top-right
        vec2(-u_kernelOffset, 0.0f),            // center-left
        vec2( 0.0f, 0.0f),                      // center-center
        vec2( u_kernelOffset, 0.0f),            // center-right
        vec2(-u_kernelOffset, -u_kernelOffset), // bottom-left
        vec2( 0.0f, -u_kernelOffset),           // bottom-center
        vec2( u_kernelOffset, -u_kernelOffset)  // bottom-right
    );

    vec3 sampleTex[9];
    for(int i = 0; i < 9; i++)
    {
        sampleTex[i] = vec3(texture(u_screenTexture, v_uv + offsets[i]));
    }

    vec3 color = vec3(0.0);
    for(int i = 0; i < 9; i++) {
        color += sampleTex[i] * u_kernel[i];
    }

    gl_FragColor = vec4(color, 1.0);
}
