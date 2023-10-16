package validation

import (
	"encoding/json"
	"errors"

	resterr "examplo.com/crud-project/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	trasl    ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ := unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *resterr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterr.NewBadRequestError("Invalid filed type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []resterr.Causes{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := resterr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return resterr.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		resterr.NewBadRequestError("Error trying to convert fieds")
	}
}
