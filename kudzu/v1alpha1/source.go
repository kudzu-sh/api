// Copyright © 2018 the Kudzu contributors.
// Licensed under the Apache License, Version 2.0; see the NOTICE file.

package v1alpha1

import (
	core "k8s.io/api/core/v1"
)

// SourceSpec specifies the source for an API or Operator.
// +k8s:deepcopy-gen=true
type SourceSpec struct {
	Image *ImageSpec `json:"image,omitempty"`
}

// ImageSpec specifies a Docker image to install an API or Operator.
// +k8s:deepcopy-gen=true
type ImageSpec struct {
	Repository string `json:"repository"`

	// +optional
	Tag string `json:"tag,omitempty"`

	// +optional
	Hash string `json:"hash,omitempty"`

	// +optional
	PullPolicy core.PullPolicy `json:"pullPolicy,omitempty"`
}

// SourceStatus is the status of fetching the source for an API or Operator.
// +k8s:deepcopy-gen=true
type SourceStatus struct {
	Image *ImageStatus `json:"image,omitempty"`
}

// ImageStatus is the status of fetching a Docker image to install an API or Operator.
// +k8s:deepcopy-gen=true
type ImageStatus struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Hash       string `json:"hash"`
}
