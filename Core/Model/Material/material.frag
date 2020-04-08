#version 410 core

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

    float uvScale;
};

uniform Material u_materials[4];

uniform sampler2D u_blendMap;
uniform bool u_hasBlendMap;

vec4 getBlendMapColor(vec2 uv) {
    if (u_hasBlendMap) {
        vec4 blendMapColor = texture(u_blendMap, uv);

        float defaultFactor = 1 - (blendMapColor.r + blendMapColor.g + blendMapColor.b);
        return vec4(defaultFactor, blendMapColor.rgb);
    } else {
        return vec4(1, 0, 0, 0);
    }
}

vec4 getDiffuseColor(vec2 uv, vec4 blendMapColor) {
    vec4 color;
    for (int i = 0; i < 4; i++) {
        if (u_materials[i].hasTextureDiffuse) {
            color += texture(u_materials[i].textureDiffuse, uv * u_materials[i].uvScale) * blendMapColor[i];
        } else {
            color += u_materials[i].baseColor.diffuse * blendMapColor[i];
        }
    }

    return color;
}

vec4 GetDiffuseColor(vec2 uv) {
    vec4 blendMapColor = getBlendMapColor(uv);

    return getDiffuseColor(uv, blendMapColor);
}

vec3 getSpecularColor(vec2 uv, vec4 blendMapColor) {
    vec3 color;
    for (int i = 0; i < 4; i++) {
        if (u_materials[i].hasTextureSpecular) {
            color += texture(u_materials[i].textureSpecular, uv * u_materials[i].uvScale).rgb * blendMapColor[i];
        } else {
            color += u_materials[i].baseColor.specular * blendMapColor[i];
        }
    }

    return color;
}

vec3 GetSpecularColor(vec2 uv) {
    vec4 blendMapColor = getBlendMapColor(uv);

    return getSpecularColor(uv, blendMapColor);
}

vec3 getEmissiveColor(vec2 uv, vec4 blendMapColor) {
    vec3 color;
    for (int i = 0; i < 4; i++) {
        if (u_materials[i].hasTextureEmissive) {
            color += texture(u_materials[i].textureEmissive, uv * u_materials[i].uvScale).rgb * blendMapColor[i];
        } else {
            color += u_materials[i].baseColor.emissive * blendMapColor[i];
        }
    }

    return color;
}

vec3 GetEmissiveColor(vec2 uv) {
    vec4 blendMapColor = getBlendMapColor(uv);

    return getEmissiveColor(uv, blendMapColor);
}



MaterialColor GetMaterialColor(vec2 uv) {
    vec4 blendMapColor = getBlendMapColor(uv);

    MaterialColor color;
    color.diffuse = getDiffuseColor(uv, blendMapColor);

    if (color.diffuse.a < 0.1) {
        discard;
    }

    color.specular = getSpecularColor(uv, blendMapColor);
    color.emissive = getEmissiveColor(uv, blendMapColor);

    return color;
}

vec3 getNormalVector(vec3 normal, vec2 uv, mat3 tbn, vec4 blendMapColor) {
    vec3 out_normal;
    for (int i = 0; i < 4; i++) {
        if (u_materials[i].hasTextureNormal) {
            vec3 tmp_normal = texture(u_materials[i].textureNormal, uv * u_materials[i].uvScale).rgb;
            tmp_normal = tmp_normal * 2.0 - 1.0f;
            out_normal += normalize(tmp_normal * tbn) * blendMapColor[i];
        } else {
            out_normal += normal * blendMapColor[i];
        }
    }

    return normalize(out_normal);
}

vec3 GetNormalVector (vec3 normal, vec2 uv, mat3 tbn) {
    vec4 blendMapColor = getBlendMapColor(uv);

    return getNormalVector(normal, uv, tbn, blendMapColor);
}

float getShininess(vec4 blendMapColor) {
    float shininess;
    for (int i = 0; i < 4; i++) {
        shininess += u_materials[i].shininess * blendMapColor[i];
    }

    return shininess;
}

float GetShininess(vec2 uv) {
    vec4 blendMapColor = getBlendMapColor(uv);

    return getShininess(blendMapColor);
}