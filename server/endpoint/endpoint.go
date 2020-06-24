package endpoint

import (
	"github.com/giantswarm/microendpoint/endpoint/healthz"
	"github.com/giantswarm/microendpoint/endpoint/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/app-service/service"
)

// Config represents the configuration used to create a endpoint.
type Config struct {
	// Dependencies.
	Logger  micrologger.Logger
	Service *service.Service
}

type Endpoint struct {
	Healthz *healthz.Endpoint
	Version *version.Endpoint
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

	var versionEndpoint *version.Endpoint
	{
		versionConfig := version.Config{
			Logger:  config.Logger,
			Service: config.Service.Version,
		}

		versionEndpoint, err = version.New(versionConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	e := &Endpoint{
		Healthz: healthzEndpoint,
		Version: versionEndpoint,
	}

	return e, nil
}
