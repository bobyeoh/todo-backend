package validation

import (
	"todo/app/utils"

	"github.com/go-playground/validator/v10"
)

type (
	// CustomValidator godoc
	CustomValidator struct {
		Validator *validator.Validate
	}
)

// Validate godoc
func (v *CustomValidator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

// ProcessError godocs
func ProcessError(err error) utils.ErrorCode {
	error := err.(validator.ValidationErrors)[0]
	switch nameSpace := error.Namespace(); nameSpace {
	case "TaskRequest.Name":
		if error.ActualTag() == "lte" {
			return *utils.TheTaskNameLengthExceed
		}
		return *utils.TheTaskNameRequired
	case "TaskRequest.ColumnID":
		return *utils.TheColumnRequired
	case "LoginRequest.Name":
		return *utils.TheNameRequired
	case "LoginRequest.Password":
		return *utils.ThePasswordRequired
	default:
		return *utils.UnknownError
	}
}
