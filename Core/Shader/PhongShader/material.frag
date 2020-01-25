#version 410 core

#define MAX_TEXTURES 24

in vec3 v_normal;
in vec2 v_uv;
in vec3 v_tangent;
in vec3 v_biTangent;

struct MaterialColor {
    vec3 diffuse;
    vec3 specular;
    vec3 emissive;
};

struct Material {
    MaterialColor baseColor;

    float shininess;

    sampler2D texturesDiffuse[MAX_TEXTURES];
    sampler2D texturesNormals[MAX_TEXTURES];

    int numTextureDiffuse;
    int numTextureNormals;
};

uniform Material u_material;

MaterialColor GetMaterialColor() {
    MaterialColor color = u_material.baseColor;
    if (u_material.numTextureDiffuse > 0) {
        vec4 diffuse = vec4(0.0, 0.0, 0.0, 0.0);
        for (int i = 0; i < u_material.numTextureDiffuse; i++){
            diffuse += texture(u_material.texturesDiffuse[i], v_uv);
        }
        // check if material is transparent
        if (diffuse.w < 0.9){
            discard;
        }
        color.diffuse = vec3(diffuse);
    }

    return color;
}

vec3 GetNormalVector () {
    vec3 normal = v_normal;
    if (u_material.numTextureNormals > 0) {
        //transpose is equal to inverse in this case
        mat3 tbn = transpose(mat3(v_tangent, v_biTangent, v_normal));
        normal = vec3(0.0, 0.0, 0.0);
        for (int i = 0; i < u_material.numTextureNormals; i++){
            normal += normalize(texture(u_material.texturesNormals[i], v_uv).rgb * 2.0 - 1.0f);
        }
        normal = normalize(normal * tbn);
    }
    return normal;
}

float GetShininess() {
    return u_material.shininess;
}