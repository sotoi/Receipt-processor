package validator

import (
	"regexp"

	goValidator "github.com/go-playground/validator/v10"
)
// Validator is an interface for validating structs
type Validator interface {
	Validate(interface{}) error
}

// Validate is a wrapper around the go-playground/validator package
type Validate struct {
	validator *goValidator.Validate

}

// NewValidate creates a new Validate instance
func NewValidate() *Validate {
	validator := goValidator.New()
	validator.RegisterValidation("retailer", retailerValidation)
	validator.RegisterValidation("price", priceValidation)
	validator.RegisterValidation("shortDescription", shortDescriptionValidation)
	return &Validate{
		validator: validator,
	}
}


// Validate validates a struct using the go-playground/validator package
func (v *Validate) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// Custom Validation logic

// retailerValidation validates the retailer field
func retailerValidation(fl goValidator.FieldLevel) bool {
	pattern := `^[\w\s\-&]+$`
	re := regexp.MustCompile(pattern)

	// Get the field value as a string and validate it against the regex
	fieldValue := fl.Field().String()
	return re.MatchString(fieldValue)
}

// priceValidation validates the price field
func priceValidation(fl goValidator.FieldLevel) bool {
	pattern := `^\d+\.\d{2}$`
	re := regexp.MustCompile(pattern)

	// Get the field value as a string and validate it against the regex
	fieldValue := fl.Field().String()
	return re.MatchString(fieldValue)
}

// shortDescriptionValidation validates the shortDescription field
func shortDescriptionValidation(fl goValidator.FieldLevel) bool {
	pattern := `^[\w\s\-]+$`
	re := regexp.MustCompile(pattern)

	// Get the field value as a string and validate it against the regex
	fieldValue := fl.Field().String()
	return re.MatchString(fieldValue)
}

