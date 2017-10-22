package endpoint

import (
	"github.com/go-kit/kit/endpoint"
)

// Middleware definition
type Middleware func(endpoint.Endpoint) endpoint.Endpoint
