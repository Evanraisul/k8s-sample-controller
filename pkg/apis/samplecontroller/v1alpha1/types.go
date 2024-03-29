/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="AvailableReplicas",type="integer",JSONPath=".status.availableReplicas"

// Evan is a specification for a Evan resource
type Evan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EvanSpec   `json:"spec,omitempty"`
	Status EvanStatus `json:"status,omitempty"`
}

type DeploymentConfig struct {
	Name     string `json:"name,omitempty"`
	Replicas *int32 `json:"replicas,omitempty"`
	Image    string `json:"image"`
}

type ServiceConfig struct {
	Name       string             `json:"name,omitempty"`
	Type       corev1.ServiceType `json:"type,omitempty"`
	Port       int32              `json:"port,omitempty"`
	TargetPort int32              `json:"targetPort,omitempty"`
	NodePort   int32              `json:"nodePort,omitempty"`
}

type DeletionPolicy string

const (
	DeletionPolicyDelete  DeletionPolicy = "Delete"
	DeletionPolicyWipeOut DeletionPolicy = "WipeOut"
)

// EvanSpec is the spec for an Evan resource
type EvanSpec struct {
	DeploymentConfig DeploymentConfig `json:"deploymentConfig,omitempty"`
	ServiceConfig    ServiceConfig    `json:"serviceConfig,omitempty"`
	DeletionPolicy   DeletionPolicy   `json:"deletionPolicy,omitempty"`
}

// EvanStatus is the status for an Evan resource
type EvanStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EvanList is a list of Evan resources
type EvanList struct {
	metav1.TypeMeta `json:",inline,omitempty"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Evan `json:"items,omitempty"`
}
