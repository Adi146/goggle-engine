package Mesh

type IVertexArray interface {
	Bind()
	Unbind()
	EnableUVAttribute()
	DisableUVAttribute()
	EnableNormalAttribute()
	DisableNormalAttribute()
	EnableTangentAttribute()
	DisableTangentAttribute()
	EnableBiTangentAttribute()
	DisableBiTangentAttribute()
}
