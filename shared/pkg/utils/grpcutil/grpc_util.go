package grpcutil

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FieldViolation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func BuildValidationError(err error) error {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	st := status.New(codes.InvalidArgument, "validation failed")
	badReq := &errdetails.BadRequest{}

	for _, e := range errs {
		badReq.FieldViolations = append(badReq.FieldViolations,
			&errdetails.BadRequest_FieldViolation{
				Field:       e.Field(),
				Description: fmt.Sprintf("%s_%s", e.Tag(), e.Value()),
			},
		)
	}

	stWithDetails, err := st.WithDetails(badReq)
	if err == nil {
		st = stWithDetails
	}

	return st.Err()
}

func ParseValidationError(err error) ([]*FieldViolation, bool) {
	st, ok := status.FromError(err)
	if !ok {
		return nil, false
	}

	var result []*FieldViolation
	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *errdetails.BadRequest:
			for _, v := range t.FieldViolations {
				result = append(result, &FieldViolation{
					Field:   v.Field,
					Message: v.Description,
				})
			}
		}
	}

	return result, len(result) > 0
}
