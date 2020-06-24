// Package label contains common Kubernetes object labels. These are defined in
// https://github.com/giantswarm/fmt/blob/master/kubernetes/annotations_and_labels.md.
package label

const (
	// AppOperatorVersion is used to determine if the custom resource is
	// supported by this version of the operatorkit resource.
	AppOperatorVersion = "app-operator.giantswarm.io/version"

	// AppOperatorVersionValue is the default value to use for the version label.
	//
	// TODO Get this value from the cluster values configmap once we start
	// including the app-operator version in platform releases.
	AppOperatorVersionValue = "1.0.0"
)
