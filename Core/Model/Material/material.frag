#version 410 core

#define MAX_TEXTURES 4

in vec3 v_normal;
in vec2 v_uv;
in vec3 v_tangent;
in vec3 v_biTangent;

struct MaterialColor {
    vec4 diffuse;
    vec3 specular;
    vec3 emissive;
};

struct Material {
    MaterialColor baseColor;

    float shininess;

    sampler2D texturesDiffuse[MAX_TEXTURES];
    sampler2D texturesSpecular[MAX_TEXTURES];
    sampler2D texturesEmissive[MAX_TEXTURES];
    sampler2D texturesNormals[MAX_TEXTURES];

    int numTextureDiffuse;
    int numTextureSpecular;
    int numTextureEmissive;
    int numTextureNormals;
};

uniform Material u_material;

vec4 GetDiffuseColor() {
    vec4 baseColor = u_material.baseColor.diffuse;

    if (u_material.numTextureDiffuse > 0) {
        vec4 diffuse = vec4(0.0, 0.0, 0.0, 0.0);
        for (int i = 0; i < u_material.numTextureDiffuse; i++){
            diffuse += texture(u_material.texturesDiffuse[i], v_uv);
        }
        baseColor = diffuse;
    }

    return baseColor;
}

vec3 GetSpecularColor() {
    vec3 baseColor = u_material.baseColor.specular;

    if (u_material.numTextureSpecular > 0) {
        vec4 specular = vec4(0, 0, 0, 0);
        for (int i = 0; i < u_material.numTextureSpecular; i++) {
            specular += texture(u_material.texturesSpecular[i], v_uv);
        }
        baseColor = vec3(specular);
    }

    return baseColor;
}

vec3 GetEmissiveColor() {
    vec3 baseColor = u_material.baseColor.emissive;

    if (u_material.numTextureEmissive > 0) {
        vec4 emissive = vec4(0, 0, 0, 0);
        for (int i = 0; i < u_material.numTextureEmissive; i++) {
            emissive += texture(u_material.texturesEmissive[i], v_uv);
        }
        baseColor = vec3(emissive);
    }

    return baseColor;
}

MaterialColor GetMaterialColor() {
    MaterialColor color = u_material.baseColor;
    color.diffuse = GetDiffuseColor();

    if (color.diffuse.a < 0.1) {
        discard;
    }

    color.specular = GetSpecularColor();
    color.emissive = GetEmissiveColor();

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