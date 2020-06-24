package webhook

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/app-service/server/endpoint/webhook/mutator"
	"github.com/giantswarm/app-service/service"
)

// Config represents the configuration used to create a webhook endpoint.
type Config struct {
	// Dependencies.
	Logger  micrologger.Logger
	Service *service.Service
}

// Endpoint is the webhook endpoint collection.
type Endpoint struct {
	Mutator *mutator.Endpoint
}

// New creates a new configured info endpoint.
func New(config Config) (*Endpoint, error) {
	var err error

	var mutatorEndpoint *mutator.Endpoint
	{
		c := mutator.Config{
			Logger:  config.Logger,
			Service: config.Service,
		}

		mutatorEndpoint, err = mutator.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newEndpoint := &Endpoint{
		Mutator: mutatorEndpoint,
	}

	return newEndpoint, nil
}
