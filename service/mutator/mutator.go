package mutator

import (
	"context"
	"fmt"
	"strings"

	"github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/viper"

	"github.com/giantswarm/app-service/flag"
	"github.com/giantswarm/app-service/pkg/label"
)

// Config represents the configuration used to create a new mutator service.
type Config struct {
	Logger micrologger.Logger

	Flag  *flag.Flag
	Viper *viper.Viper
}

// Service is our service data structure
type Service struct {
	Logger micrologger.Logger

	Flag  *flag.Flag
	Viper *viper.Viper
}

// New creates a new configured mutator service.
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

	newService := &Service{
		Logger: config.Logger,

		Flag:  config.Flag,
		Viper: config.Viper,
	}

	return newService, nil
}

// Mutate implements the defaulting logic. It returns a slice of JSON patches
// that will be included in the webhook response.
func (s *Service) Mutate(ctx context.Context, newAppCR, oldAppCR *v1alpha1.App) ([]PatchOperation, error) {
	patches, err := newMetadataPatches(ctx, newAppCR)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return patches, nil
}

func newMetadataPatches(ctx context.Context, appCR *v1alpha1.App) ([]PatchOperation, error) {
	patches := []PatchOperation{}

	// Set the version label if it is missing.
	_, ok := appCR.Labels[label.AppOperatorVersion]
	if !ok {
		// If there are no labels we need to first add this patch.
		if len(appCR.Labels) == 0 {
			patches = append(patches, PatchAdd("/metadata/labels", map[string]string{}))
		}

		path := fmt.Sprintf("/metadata/labels/%s", replaceToEscape(label.AppOperatorVersion))
		patches = append(patches, PatchAdd(path, label.AppOperatorVersionValue))
	}

	return patches, nil
}

func replaceToEscape(from string) string {
	return strings.Replace(from, "/", "~1", -1)
}
