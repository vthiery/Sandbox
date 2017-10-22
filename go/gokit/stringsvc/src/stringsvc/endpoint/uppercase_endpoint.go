package endpoint

import (
	"context"

	"stringsvc/service"

	"github.com/go-kit/kit/endpoint"
)

// UppercaseRequest defines an uppercase request
type UppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

// MakeUppercaseEndpoint creates a new endpoint
func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}
