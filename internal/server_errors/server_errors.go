package server

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var ErrDBRecordNotFound = errors.New("record not found")

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ServerErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ServerErr) Error() string {
	return e.Message
}

func NewServerErr(code int, message string) *ServerErr {
	return &ServerErr{Code: code, Message: message}
}

// func HandleValidationError(err error) (isFatal bool, errorBag []*ErrorDetail) {
// 	if err == nil {
// 		return false, nil
// 	}

// 	var invalidErr *validator.InvalidValidationError
// 	if errors.As(err, &invalidErr) {
// 		return true, nil
// 	}

// 	bag := ValErrToBag(err)
// 	return false, bag
// }

func ValErrToBag(err error) []*ErrorDetail {
	validationErr, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}
	errorBag := make([]*ErrorDetail, len(validationErr))
	for _, v := range validationErr {
		field := v.Field()
		message := v.Error()

		error := &ErrorDetail{
			Field:   field,
			Message: message,
		}

		errorBag = append(errorBag, error)
	}

	return errorBag
}
