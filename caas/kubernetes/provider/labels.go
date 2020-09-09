// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package provider

import (
	"k8s.io/apimachinery/pkg/labels"
)

const (
	// LabelJujuAppCreatedBy is a Juju application label to apply to objects
	// created by applications managed by Juju. Think istio, kubeflow etc
	// See https://bugs.launchpad.net/juju/+bug/1892285
	LabelJujuAppCreatedBy = "app.juju.is/created-by"
)

// AppendLabels adds the labels defined in src to dest returning the result.
// Overlapping keys in sources maps are overwritten by the very last defined
// value for a duplicate key.
func AppendLabels(dest map[string]string, sources ...map[string]string) map[string]string {
	if dest == nil {
		dest = map[string]string{}
	}
	if sources == nil {
		return dest
	}
	for _, s := range sources {
		for k, v := range s {
			dest[k] = v
		}
	}
	return dest
}

func (k *kubernetesClient) getlabelsForApp(appName string, isNamespaced bool) map[string]string {
	labels := LabelsForApp(appName)
	if !isNamespaced {
		labels = AppendLabels(labels, LabelsForModel(k.CurrentModel()))
	}
	return labels
}

// LabelsForApp returns the labels that should be on a k8s object for a given
// application name
func LabelsForApp(name string) map[string]string {
	return map[string]string{
		labelApplication: name,
	}
}

// LabelsForModel returns the labels that should be on a k8s object for a given
// model name
func LabelsForModel(name string) map[string]string {
	return map[string]string{
		labelModel: name,
	}
}

// LabelForKeyValue returns a Kubernetes label set for the supplied key value.
func LabelForKeyValue(key, value string) labels.Set {
	return labels.Set{
		key: value,
	}
}
