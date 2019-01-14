// Copyright © 2018 the Kudzu contributors.
// Licensed under the Apache License, Version 2.0; see the NOTICE file.

package v1alpha1

import (
	core "k8s.io/api/core/v1"
)

func SetDefaults_ImageSpec(spec *ImageSpec) {
	if spec.Tag == "" {
		spec.Tag = "latest"
	}

	if spec.PullPolicy == "" {
		if spec.Tag == "latest" {
			spec.PullPolicy = core.PullAlways
		} else {
			spec.PullPolicy = core.PullIfNotPresent
		}
	}
}
