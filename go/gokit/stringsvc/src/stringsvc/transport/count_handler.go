package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"stringsvc/endpoint"
	"stringsvc/service"

	httptransport "github.com/go-kit/kit/transport/http"
)

// NewCountHandler creates a count handler
func NewCountHandler(svc service.StringService) http.Handler {
	return httptransport.NewServer(
		endpoint.MakeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
