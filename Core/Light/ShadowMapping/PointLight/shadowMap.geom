#version 410 core
#define MAX_POINT_LIGHTS 32

layout (triangles) in;
layout (triangle_strip, max_vertices=18) out;

struct PointLight{
    vec3 position;
    float linear;
    float quadratic;

    vec3 ambient;
    vec3 diffuse;
    vec3 specular;

    mat4 viewProjectionMatrix[6];
};

layout (std140) uniform pointLight {
    int u_numPointLights;
    PointLight u_pointLights[MAX_POINT_LIGHTS];
};

uniform int u_lightIndex;

in vec2 g_uv[3];

out vec4 FragPos;
out vec2 v_uv;

void main() {
    for(int face = 0; face < 6; ++face)
    {
        gl_Layer = face;
        for(int i = 0; i < 3; ++i)
        {
            FragPos = gl_in[i].gl_Position;
            v_uv = g_uv[i];
            gl_Position = FragPos * u_pointLights[u_lightIndex].viewProjectionMatrix[face];
            EmitVertex();
        }
        EndPrimitive();
    }
}
