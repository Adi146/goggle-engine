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

vec4 GetDiffuseColor(vec2 uv) {
    vec4 baseColor = u_material.baseColor.diffuse;

    if (u_material.numTextureDiffuse > 0) {
        vec4 diffuse = vec4(0.0, 0.0, 0.0, 0.0);
        for (int i = 0; i < u_material.numTextureDiffuse; i++){
            diffuse += texture(u_material.texturesDiffuse[i], uv);
        }
        baseColor = diffuse;
    }

    return baseColor;
}

vec3 GetSpecularColor(vec2 uv) {
    vec3 baseColor = u_material.baseColor.specular;

    if (u_material.numTextureSpecular > 0) {
        vec4 specular = vec4(0, 0, 0, 0);
        for (int i = 0; i < u_material.numTextureSpecular; i++) {
            specular += texture(u_material.texturesSpecular[i], uv);
        }
        baseColor = vec3(specular);
    }

    return baseColor;
}

vec3 GetEmissiveColor(vec2 uv) {
    vec3 baseColor = u_material.baseColor.emissive;

    if (u_material.numTextureEmissive > 0) {
        vec4 emissive = vec4(0, 0, 0, 0);
        for (int i = 0; i < u_material.numTextureEmissive; i++) {
            emissive += texture(u_material.texturesEmissive[i], uv);
        }
        baseColor = vec3(emissive);
    }

    return baseColor;
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
    if (u_material.numTextureNormals > 0) {
        //transpose is equal to inverse in this case
        normal = vec3(0.0, 0.0, 0.0);
        for (int i = 0; i < u_material.numTextureNormals; i++){
            normal += normalize(texture(u_material.texturesNormals[i], uv).rgb * 2.0 - 1.0f);
        }
        normal = normalize(normal * tbn);
    }
    return normal;
}

float GetShininess() {
    return u_material.shininess;
}