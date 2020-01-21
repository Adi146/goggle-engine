package Utils

import "fmt"

type ErrorCollection struct {
	errors []error
}

func (e *ErrorCollection) Error() string {
	switch len(e.errors) {
	case 0:
		return ""
	case 1:
		return e.errors[0].Error()
	default:
		re := fmt.Sprintf("%d errors occured:", len(e.errors))
		for _, err := range e.errors {
			re += " " + err.Error()
		}
		return re
	}
}

func (e *ErrorCollection) Err() error {
	switch len(e.errors) {
	case 0:
		return nil
	case 1:
		return e.errors[0]
	default:
		return e
	}
}

func (e *ErrorCollection) Push(err error) {
	if err == nil {
		return
	}

	switch v := err.(type) {
	case *ErrorCollection:
		e.errors = append(e.errors, v.errors...)
	default:
		e.errors = append(e.errors, err)
	}
}

func (e *ErrorCollection) Len() int {
	return len(e.errors)
}
