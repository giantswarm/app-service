package mutator

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"k8s.io/api/admission/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	"github.com/giantswarm/app-service/service"
)

const (
	// Method is the HTTP method this endpoint is registered for.
	Method = "POST"
	// Name identifies the endpoint. It is aligned to the package path.
	Name = "webhook/mutator"
	// Path is the HTTP request path this endpoint is registered for.
	Path = "/webhooks/mutator"
)

// Config represents the configuration used to create a version endpoint.
type Config struct {
	// Dependencies.
	Logger  micrologger.Logger
	Service *service.Service
}

// New creates a new configured version endpoint.
func New(config Config) (*Endpoint, error) {
	// Dependencies.
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "logger must not be empty")
	}
	if config.Service == nil {
		return nil, microerror.Maskf(invalidConfigError, "service must not be empty")
	}

	scheme := runtime.NewScheme()
	codecs := serializer.NewCodecFactory(scheme)
	deserializer := codecs.UniversalDeserializer()

	newEndpoint := &Endpoint{
		deserializer: deserializer,
		logger:       config.Logger,
		service:      config.Service,
	}

	return newEndpoint, nil
}

// Endpoint is the endpoint data structure.
type Endpoint struct {
	// Dependencies.
	deserializer runtime.Decoder
	logger       micrologger.Logger
	service      *service.Service
}

// Decoder returns a function to decode requests.
func (e *Endpoint) Decoder() kithttp.DecodeRequestFunc {
	return func(ctx context.Context, request *http.Request) (interface{}, error) {
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		review := v1beta1.AdmissionReview{}

		_, _, err = e.deserializer.Decode(data, nil, &review)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		return review, nil
	}
}

// Endpoint returns a function applying the endoint.
func (e *Endpoint) Endpoint() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error

		review := request.(v1beta1.AdmissionReview)

		newAppCR := &v1alpha1.App{}
		_, _, err = e.deserializer.Decode(review.Request.Object.Raw, nil, newAppCR)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		oldAppCR := &v1alpha1.App{}
		_, _, err = e.deserializer.Decode(review.Request.OldObject.Raw, nil, oldAppCR)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		patches, err := e.service.Mutator.Mutate(ctx, newAppCR, oldAppCR)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		patchData, err := json.Marshal(patches)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		pt := v1beta1.PatchTypeJSONPatch

		response := &v1beta1.AdmissionResponse{
			Allowed:   true,
			UID:       review.Request.UID,
			Patch:     patchData,
			PatchType: &pt,
		}

		return response, nil
	}
}

// Encoder returns a function to encode requests.
func (e *Endpoint) Encoder() kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		admissionResponse := response.(*v1beta1.AdmissionResponse)

		review := &v1beta1.AdmissionReview{
			Response: admissionResponse,
		}

		resp, err := json.Marshal(review)
		if err != nil {
			return microerror.Mask(err)
		}

		_, err = w.Write(resp)
		if err != nil {
			return microerror.Mask(err)
		}

		return nil
	}
}

// Method returns the name of this endpoint
func (e *Endpoint) Method() string {
	return Method
}

// Name returns the nme of this endpoint
func (e *Endpoint) Name() string {
	return Name
}

// Path returns the path of this endpoint
func (e *Endpoint) Path() string {
	return Path
}
