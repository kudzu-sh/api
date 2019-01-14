// +build !ignore_autogenerated

// Copyright © 2018 the Kudzu contributors.
// Licensed under the Apache License, Version 2.0; see the NOTICE file.

// Code generated by defaulter. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&API{}, func(obj interface{}) { SetObjectDefaults_API(obj.(*API)) })
	scheme.AddTypeDefaultingFunc(&APIList{}, func(obj interface{}) { SetObjectDefaults_APIList(obj.(*APIList)) })
	return nil
}

func SetObjectDefaults_API(in *API) {
	if in.Spec.Source.Image != nil {
		SetDefaults_ImageSpec(in.Spec.Source.Image)
	}
}

func SetObjectDefaults_APIList(in *APIList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_API(a)
	}
}