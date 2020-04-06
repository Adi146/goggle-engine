#version 410 core
#define MAX_TEXTURES 4

struct MaterialColor {
    vec4 diffuse;
    vec3 specular;
    vec3 emissive;
};

struct Material {
    MaterialColor baseColor;

    float shininess;

    sampler2D textureDiffuse;
    sampler2D textureSpecular;
    sampler2D textureEmissive;
    sampler2D textureNormal;

    bool hasTextureDiffuse;
    bool hasTextureSpecular;
    bool hasTextureEmissive;
    bool hasTextureNormal;
};

uniform Material u_material;

vec4 GetDiffuseColor(vec2 uv) {
    if (u_material.hasTextureDiffuse) {
        return texture(u_material.textureDiffuse, uv);
    } else {
        return u_material.baseColor.diffuse;
    }
}

vec3 GetSpecularColor(vec2 uv) {
    if (u_material.hasTextureSpecular) {
        return texture(u_material.textureSpecular, uv).rgb;
    } else {
        return u_material.baseColor.specular;
    }
}

vec3 GetEmissiveColor(vec2 uv) {
    if (u_material.hasTextureEmissive) {
        return texture(u_material.textureEmissive, uv).rgb;
    } else {
        return u_material.baseColor.emissive;
    }
}

MaterialColor GetMaterialColor(vec2 uv) {
    MaterialColor color = u_material.baseColor;
    color.diffuse = GetDiffuseColor(uv);

    if (color.diffuse.a < 0.1) {
        discard;
    }

    color.specular = GetSpecularColor(uv);
    color.emissive = GetEmissiveColor(uv);

    return color;
}

vec3 GetNormalVector (vec3 normal, vec2 uv, mat3 tbn) {
    if (u_material.hasTextureNormal) {
        //transpose is equal to inverse in this case
        return normalize(texture(u_material.textureNormal, uv).rgb * 2.0 - 1.0f);
    } else {
        return normal;
    }
}

float GetShininess() {
    return u_material.shininess;
}