package Error

type ErrorWithFields struct {
	error  error
	Fields map[string]interface{}
}

func NewErrorWithFields(err error, fields map[string]interface{}) error {
	switch v := err.(type) {
	case nil:
		return nil
	case *ErrorWithFields:
		return v
	case *ErrorCollection:
		for i := 0; i < len(v.Errors); i++ {
			v.Errors[i] = NewErrorWithFields(v.Errors[i], fields)
		}
		return v
	default:
		return &ErrorWithFields{
			error:  err,
			Fields: fields,
		}
	}
}

func (e *ErrorWithFields) Error() string {
	return e.error.Error()
}
