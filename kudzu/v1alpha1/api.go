// Copyright © 2018 the Kudzu contributors.
// Licensed under the Apache License, Version 2.0; see the NOTICE file.

package v1alpha1

import (
	duck "github.com/knative/pkg/apis/duck/v1alpha1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// API is a set of custom resource definitions from the same group and version.
type API struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   APISpec   `json:"spec"`
	Status APIStatus `json:"status"`
}

// GetSourceSpec satisfies the Kudzu delegate.Object interface, allowing
// delegate Pods to be created to reify an API.
func (a *API) GetSourceSpec() *SourceSpec {
	return &a.Spec.Source
}

// GetSourceStatus satisfies the Kudzu delegate.Object interface, allowing
// delegate Pods to be created to reify an API.
func (a *API) GetSourceStatus() *SourceStatus {
	return a.Status.Source
}

// SetSourceStatus satisfies the Kudzu delegate.Object interface, allowing
// identity of the image used to reify the API to be saved.
func (a *API) SetSourceStatus(status *SourceStatus) {
	a.Status.Source = status
}

type APISpec struct {
	Source SourceSpec `json:"source"`
}

const (
	APIReady                      = duck.ConditionReady
	APIApplied duck.ConditionType = "Applied"
	APIUpdated duck.ConditionType = "Updated"
)

var apiCondSet = duck.NewLivingConditionSet(APIApplied, APIUpdated)

type APIStatus struct {
	// +optional
	Source *SourceStatus `json:"source,omitempty"`

	// +optional
	Conditions duck.Conditions `json:"conditions,omitempty"`

	// +optional
	ResourceCount int32 `json:"resourceCount,omitempty"`

	// +optional
	Resources []ResourceStatus `json:"resources,omitempty"`
}

func (s *APIStatus) ConditionManager() duck.ConditionManager {
	return apiCondSet.Manage(s)
}

func (s *APIStatus) IsReady() bool {
	return s.ConditionManager().IsHappy()
}

func (s *APIStatus) GetCondition(t duck.ConditionType) *duck.Condition {
	return s.ConditionManager().GetCondition(t)
}

func (s *APIStatus) InitializeConditions() {
	s.ConditionManager().InitializeConditions()
}

func (s *APIStatus) GetConditions() duck.Conditions {
	return s.Conditions
}

func (s *APIStatus) SetConditions(conditions duck.Conditions) {
	s.Conditions = conditions
}

// +k8s:deepcopy-gen=true
type ResourceStatus struct {
	Name    string `json:"name"`
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIList is a list of API resources.
type APIList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata"`

	// Items is the list of API items in this list.
	Items []API `json:"items"`
}
