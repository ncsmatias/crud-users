package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/ncsmatias/crud-users/src/configuration/resterr"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationError error) *resterr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonErr) {
		return resterr.BadRequestError("Invalid field type")
	} else if errors.As(validationError, &jsonValidationError) {
		errorCauses := []resterr.Causes{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := resterr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return resterr.BadRequestValidationError("Some fields are invalid", errorCauses)
	}

	return resterr.BadRequestError("Error trying to convert fields")
}
