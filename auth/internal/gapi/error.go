package gapi

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func toBadRequestFieldValidation(validations validation.Errors) []*errdetails.BadRequest_FieldViolation {
	fieldViolations := make([]*errdetails.BadRequest_FieldViolation, 0, len(validations))
	for field, err := range validations {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{
			Field:       field,
			Description: err.Error(),
		})
	}
	return fieldViolations
}

func invalidArgumentError(validations validation.Errors) error {
	badRequest := &errdetails.BadRequest{FieldViolations: toBadRequestFieldValidation(validations)}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")

	statusDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}

	return statusDetails.Err()
}
