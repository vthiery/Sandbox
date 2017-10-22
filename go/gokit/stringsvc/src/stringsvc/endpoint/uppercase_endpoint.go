package endpoint

import (
	"context"
	"os"

	"stringsvc/service"

	"github.com/go-kit/kit/endpoint"
	kitLog "github.com/go-kit/kit/log"
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
	uppercase := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
	logger := kitLog.NewLogfmtLogger(os.Stderr)
	uppercase = LoggingMiddleware(kitLog.With(logger, "method", "uppercase"))(uppercase)
	return uppercase
}
