package service

import (
	"time"

	"github.com/giantswarm/microendpoint/service/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/viper"

	"github.com/giantswarm/app-service/flag"
	"github.com/giantswarm/app-service/pkg/project"
	"github.com/giantswarm/app-service/service/mutator"
)

const (
	// DefaultRetryCount is the number of times to retry a failed network call.
	DefaultRetryCount = 5
	// DefaultTimeout is the timeout for network calls.
	DefaultTimeout = 5 * time.Second
)

// Config represents the configuration used to create a new service.
type Config struct {
	Logger micrologger.Logger
	Flag   *flag.Flag

	Viper *viper.Viper
}

// Service is a type providing implementation of microkit service interface.
type Service struct {
	Mutator *mutator.Service
	Version *version.Service
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Logger must not be empty")
	}

	if config.Flag == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Flag must not be empty")
	}
	if config.Viper == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Viper must not be empty")
	}

	var err error

	var mutatorService *mutator.Service
	{
		c := mutator.Config{
			Logger: config.Logger,

			Flag:  config.Flag,
			Viper: config.Viper,
		}

		mutatorService, err = mutator.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var versionService *version.Service
	{
		c := version.Config{
			Description: project.Description(),
			GitCommit:   project.GitSHA(),
			Name:        project.Name(),
			Source:      project.Source(),
			Version:     project.Version(),
		}

		versionService, err = version.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newService := &Service{
		Mutator: mutatorService,
		Version: versionService,
	}

	return newService, nil
}
