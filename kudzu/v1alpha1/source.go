// Copyright © 2018 the Kudzu contributors.
// Licensed under the Apache License, Version 2.0; see the NOTICE file.

package v1alpha1

import (
	"fmt"

	core "k8s.io/api/core/v1"
)

// SourceSpec specifies the source for an API or Operator.
// +k8s:deepcopy-gen=true
type SourceSpec struct {
	Image *ImageSpec `json:"image,omitempty"`
}

// Matches verifies that a source's status matches its spec.
func (s SourceSpec) Matches(status *SourceStatus) bool {
	if s.Image == nil || status == nil {
		return false
	}
	return s.Image.Matches(status.Image)
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

func (i ImageSpec) String() string {
	if i.Hash != "" {
		return fmt.Sprintf("%s@%s", i.Repository, i.Hash)
	}
	if i.Tag != "" {
		return fmt.Sprintf("%s:%s", i.Repository, i.Tag)
	}
	return fmt.Sprintf("%s:latest", i.Repository)
}

// Matches verifies that an image's status matches its spec.
func (i ImageSpec) Matches(status *ImageStatus) bool {
	if status == nil || i.Repository != status.Repository {
		return false
	}
	if i.Hash != "" && i.Hash != status.Hash {
		return false
	}
	if (i.Tag != "" && i.Tag != status.Tag) || (i.Tag == "" && status.Tag != "latest") {
		return false
	}
	return true
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
