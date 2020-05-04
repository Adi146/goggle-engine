#version 430 core
#define MAX_POINT_LIGHTS 32

layout (triangles) in;
layout (triangle_strip, max_vertices=18) out;

in VS_OUT {
    vec2 uv;
} gs_in[3];

out GS_OUT {
    vec4 position;
    vec2 uv;
} gs_out;

layout (std140, binding = 2) uniform pointLight {
    int u_numPointLights;
    struct {
        vec3 position;
        float linear;
        float quadratic;

        vec3 ambient;
        vec3 diffuse;
        vec3 specular;

        mat4 viewProjectionMatrix[6];
        float distance;
    } u_pointLights[MAX_POINT_LIGHTS];
};

uniform int u_lightIndex;

void main() {
    for(int face = 0; face < 6; ++face)
    {
        gl_Layer = face;
        for(int i = 0; i < 3; ++i)
        {
            gs_out.position = gl_in[i].gl_Position;
            gs_out.uv = gs_in[i].uv;
            gl_Position = gl_in[i].gl_Position * u_pointLights[u_lightIndex].viewProjectionMatrix[face];
            EmitVertex();
        }
        EndPrimitive();
    }
}
