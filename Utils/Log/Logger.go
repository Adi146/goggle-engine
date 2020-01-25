package Log

import (
	ErrUtils "github.com/Adi146/goggle-engine/Utils/Error"
	"github.com/sirupsen/logrus"
)

func Error(err error, msg string) {
	switch v := err.(type) {
	case nil:
		return
	case *ErrUtils.ErrorCollection:
		for _, err := range v.Errors {
			Error(err, msg)
		}
	case *ErrUtils.ErrorWithFields:
		logrus.WithFields(v.Fields).WithError(v).Error(msg)
	default:
		logrus.WithError(v).Error(msg)
	}
}

func Warn(err error, msg string) {
	switch v := err.(type) {
	case nil:
		return
	case *ErrUtils.ErrorCollection:
		for _, err := range v.Errors {
			Warn(err, msg)
		}
	case *ErrUtils.ErrorWithFields:
		logrus.WithFields(v.Fields).WithError(v).Warn(msg)
	default:
		logrus.WithError(v).Warn(msg)
	}
}
