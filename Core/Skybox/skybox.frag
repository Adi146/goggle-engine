#version 410 core

in VS_OUT {
    vec3 uv;
} fs_in;

layout(location = 0) out vec4 f_color;

uniform samplerCube u_skybox;

void main()
{
    f_color = texture(u_skybox, fs_in.uv);
}