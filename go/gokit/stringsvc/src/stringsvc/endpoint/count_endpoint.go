package endpoint

import (
	"context"
	"os"

	"stringsvc/service"

	"github.com/go-kit/kit/endpoint"
	kitLog "github.com/go-kit/kit/log"
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
	count := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(ctx, req.S)
		return countResponse{v}, nil
	}
	logger := kitLog.NewLogfmtLogger(os.Stderr)
	count = LoggingMiddleware(kitLog.With(logger, "method", "count"))(count)
	return count
}
