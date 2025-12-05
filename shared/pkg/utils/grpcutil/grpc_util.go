package grpcutil

import (
	"buf.build/go/protovalidate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FieldViolation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func BuildValidationError(err error) error {
	ve, ok := err.(*protovalidate.ValidationError)
	if !ok {
		return err
	}

	st := status.New(codes.InvalidArgument, "validation failed")
	badReq := &errdetails.BadRequest{}

	for _, v := range ve.Violations {
		badReq.FieldViolations = append(badReq.FieldViolations,
			&errdetails.BadRequest_FieldViolation{
				Field:       protovalidate.FieldPathString(v.Proto.GetField()),
				Description: v.Proto.GetMessage(),
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
