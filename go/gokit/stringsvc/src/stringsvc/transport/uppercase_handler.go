package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"stringsvc/endpoint"
	"stringsvc/service"

	httptransport "github.com/go-kit/kit/transport/http"
)

// NewUppercaseHandler creates an uppercase handler
func NewUppercaseHandler(svc service.StringService) http.Handler {
	return httptransport.NewServer(
		endpoint.MakeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
