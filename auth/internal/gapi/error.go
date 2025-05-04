package gapi

import (
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func toBadRequestFieldValidation(validations domain.ValidationErrors) []*errdetails.BadRequest_FieldViolation {
	fieldViolations := make([]*errdetails.BadRequest_FieldViolation, 0, len(validations))
	for _, v := range validations {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{
			Field:       v.Field,
			Description: v.Message,
		})
	}
	return fieldViolations
}

func invalidArgumentError(validations domain.ValidationErrors) error {
	badRequest := &errdetails.BadRequest{FieldViolations: toBadRequestFieldValidation(validations)}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")

	statusDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}

	return statusDetails.Err()
}
