package AssetImporter

type ImportResult struct {
	errors            []error
	warnings          []error
	NumImportedAssets int
}

func newImportResult() *ImportResult {
	return &ImportResult{
		errors:            []error{},
		warnings:          []error{},
		NumImportedAssets: 0,
	}
}

func (result *ImportResult) addError(err ...error) {
	result.errors = append(result.errors, err...)
}

func (result *ImportResult) addWarning(warning ...error) {
	result.warnings = append(result.warnings, warning...)
}

func (result *ImportResult) addResult(re *ImportResult) {
	result.addError(re.errors...)
	result.addWarning(re.warnings...)
	result.NumImportedAssets += re.NumImportedAssets
}

func (result *ImportResult) addResultAsWarning(re *ImportResult) {
	result.addWarning(re.errors...)
	result.addWarning(re.warnings...)
	result.NumImportedAssets += re.NumImportedAssets
}

func (result *ImportResult) Success() bool {
	return len(result.errors) == 0 && result.NumImportedAssets >= 1
}
