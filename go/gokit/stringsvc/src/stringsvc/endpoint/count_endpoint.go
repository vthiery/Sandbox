package endpoint

import (
	"context"

	"stringsvc/service"

	"github.com/go-kit/kit/endpoint"
)

// CountRequest defines a count request
type CountRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

// MakeCountEndpoint creates a new endpoint
func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(ctx, req.S)
		return countResponse{v}, nil
	}
}
