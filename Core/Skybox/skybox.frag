#version 410 core

in vec3 v_uv;

layout(location = 0) out vec4 f_color;

uniform samplerCube u_skybox;

void main()
{    
    f_color = texture(u_skybox, v_uv);
}