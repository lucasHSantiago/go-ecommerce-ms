package application

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/validator"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[A-Za-z ]+$`).MatchString
)

func validateCreateUserParams(arg params.CreateUserApp) *validator.ValidationErrors {
	var errs validator.ValidationErrors

	if err := validateUsername(arg.Username); err != nil {
		errs = append(errs, validator.NewFieldValidation("username", err))
	}

	if err := validatePassword(arg.Password); err != nil {
		errs = append(errs, validator.NewFieldValidation("password", err))
	}

	if err := validateFullName(arg.FullName); err != nil {
		errs = append(errs, validator.NewFieldValidation("full_name", err))
	}

	if err := validateEmail(arg.Email); err != nil {
		errs = append(errs, validator.NewFieldValidation("email", err))
	}

	if len(errs) > 0 {
		return &errs
	}

	return nil
}

func validateUpdateUserParams(arg params.UpdateUserApp) *validator.ValidationErrors {
	var errs validator.ValidationErrors

	if err := validateUsername(arg.Username); err != nil {
		errs = append(errs, validator.NewFieldValidation("username", err))
	}

	if arg.FullName != nil {
		if err := validateFullName(*arg.FullName); err != nil {
			errs = append(errs, validator.NewFieldValidation("full_name", err))
		}
	}

	if arg.Password != nil {
		if err := validatePassword(*arg.Password); err != nil {
			errs = append(errs, validator.NewFieldValidation("password", err))
		}
	}

	if arg.Email != nil {
		if err := validateEmail(*arg.Email); err != nil {
			errs = append(errs, validator.NewFieldValidation("email", err))
		}
	}

	if len(errs) > 0 {
		return &errs
	}

	return nil
}

func validateLoginUserParams(arg params.LoginUserApp) *validator.ValidationErrors {
	var errs validator.ValidationErrors

	if err := validateUsername(arg.Username); err != nil {
		errs = append(errs, validator.NewFieldValidation("username", err))
	}

	if err := validatePassword(arg.Password); err != nil {
		errs = append(errs, validator.NewFieldValidation("password", err))
	}

	if len(errs) > 0 {
		return &errs
	}

	return nil
}

func validateUsername(value string) error {
	if err := validator.ValidateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("must contain only letter, digits or underscores")
	}

	return nil
}

func validateFullName(value string) error {
	if err := validator.ValidateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letter and spaces")
	}

	return nil
}

func validatePassword(value string) error {
	return validator.ValidateString(value, 6, 100)
}

func validateEmail(value string) error {
	if err := validator.ValidateString(value, 3, 200); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}

	return nil
}
