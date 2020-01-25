package AssetImporter

import (
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type ImportResult struct {
	Errors   Error.ErrorCollection
	Warnings Error.ErrorCollection
}

func (result *ImportResult) Success() bool {
	return result.Errors.Len() == 0
}
