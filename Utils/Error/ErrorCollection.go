package Error

import "fmt"

type ErrorCollection struct {
	Errors []error
}

func (e *ErrorCollection) Error() string {
	switch len(e.Errors) {
	case 0:
		return ""
	case 1:
		return e.Errors[0].Error()
	default:
		re := fmt.Sprintf("%d Errors occured:", len(e.Errors))
		for _, err := range e.Errors {
			re += " " + err.Error()
		}
		return re
	}
}

func (e *ErrorCollection) Err() error {
	switch len(e.Errors) {
	case 0:
		return nil
	case 1:
		return e.Errors[0]
	default:
		return e
	}
}

func (e *ErrorCollection) Push(err error) {
	switch v := err.(type) {
	case nil:
		return
	case *ErrorCollection:
		e.Errors = append(e.Errors, v.Errors...)
	default:
		e.Errors = append(e.Errors, err)
	}
}

func (e *ErrorCollection) Len() int {
	return len(e.Errors)
}
