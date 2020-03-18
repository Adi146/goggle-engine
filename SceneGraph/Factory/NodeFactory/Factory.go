package NodeFactory

import (
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

var (
	typeFactory = map[string]reflect.Type{
		"Scene.Node": reflect.TypeOf((*Scene.Node)(nil)).Elem(),
	}
)

func AddType(key string, nodeType reflect.Type) {
	typeFactory[key] = nodeType
}
