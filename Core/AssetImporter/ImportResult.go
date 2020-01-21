package AssetImporter

import "github.com/Adi146/goggle-engine/Utils"

type ImportResult struct {
	Errors   Utils.ErrorCollection
	Warnings Utils.ErrorCollection
}

func (result *ImportResult) Success() bool {
	return result.Errors.Len() == 0
}
