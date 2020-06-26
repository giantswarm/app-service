package endpoint

import (
	"github.com/giantswarm/microendpoint/endpoint/healthz"
	"github.com/giantswarm/microendpoint/endpoint/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/app-service/server/endpoint/webhook"
	"github.com/giantswarm/app-service/server/middleware"
	"github.com/giantswarm/app-service/service"
)

// Config represents the configuration used to create a endpoint.
type Config struct {
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Service    *service.Service
}

type Endpoint struct {
	Healthz *healthz.Endpoint
	Version *version.Endpoint
	Webhook *webhook.Endpoint
}

func New(config Config) (*Endpoint, error) {
	var err error

	var healthzEndpoint *healthz.Endpoint
	{
		c := healthz.Config{
			Logger: config.Logger,
		}

		healthzEndpoint, err = healthz.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var webhookEndpoint *webhook.Endpoint
	{
		c := webhook.Config{
			Logger:     config.Logger,
			Middleware: config.Middleware,
			Service:    config.Service.Mutator,
		}

		webhookEndpoint, err = webhook.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var versionEndpoint *version.Endpoint
	{
		c := version.Config{
			Logger:  config.Logger,
			Service: config.Service.Version,
		}

		versionEndpoint, err = version.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	e := &Endpoint{
		Healthz: healthzEndpoint,
		Version: versionEndpoint,
		Webhook: webhookEndpoint,
	}

	return e, nil
}
