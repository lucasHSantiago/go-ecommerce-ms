package application

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`)
	isValidFullName = regexp.MustCompile(`^[A-Za-z ]+$`)
)

func validateCreateUserParams(arg params.CreateUserApp) error {
	return validation.ValidateStruct(&arg,
		validation.Field(&arg.Username, validateUsername()...),
		validation.Field(&arg.Password, validatePassword()...),
		validation.Field(&arg.FullName, validateFullName()...),
		validation.Field(&arg.Email, validateEmail()...))
}

func validateUpdateUserParams(arg params.UpdateUserApp) error {
	return validation.ValidateStruct(&arg,
		validation.Field(&arg.Username, validateUsername()...),
		validation.Field(&arg.FullName, validateFullName()...),
		validation.Field(&arg.Password, validation.Length(6, 200), validation.NilOrNotEmpty),
		validation.Field(&arg.Email, validateEmail()...))
}

func validateLoginUserParams(arg params.LoginUserApp) error {
	return validation.ValidateStruct(&arg,
		validation.Field(&arg.Username, validateUsername()...),
		validation.Field(&arg.Password, validatePassword()...))
}

func validateUsername() []validation.Rule {
	rules := []validation.Rule{}
	rules = append(rules, validation.Required)
	rules = append(rules, validation.Length(3, 100))
	rules = append(rules, validation.Match(isValidUsername).Error("must contain only letter, digits or underscores"))
	return rules
}

func validateFullName() []validation.Rule {
	rules := []validation.Rule{}
	rules = append(rules, validation.Required)
	rules = append(rules, validation.Length(3, 100))
	rules = append(rules, validation.Match(isValidFullName).Error("must contain only letter and spaces"))
	return rules
}

func validatePassword() []validation.Rule {
	rules := []validation.Rule{}
	rules = append(rules, validation.Required)
	rules = append(rules, validation.Length(6, 100))
	return rules
}

func validateEmail() []validation.Rule {
	rules := []validation.Rule{}
	rules = append(rules, validation.Required)
	rules = append(rules, validation.Length(3, 200))
	rules = append(rules, is.Email)
	return rules
}
