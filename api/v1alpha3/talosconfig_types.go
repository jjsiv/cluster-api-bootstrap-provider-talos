// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha3

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ConfigFinalizer allows us to clean up resources before deletion
	ConfigFinalizer = "talosconfig.bootstrap.cluster.x-k8s.io"
)

// TalosConfigSpec defines the desired state of TalosConfig
type TalosConfigSpec struct {
	TalosVersion  string          `json:"talosVersion,omitempty"` //talos version formatted like v0.8. used for backwards compatibility
	GenerateType  string          `json:"generateType"`           //none,init,controlplane,worker mutually exclusive w/ data
	Data          string          `json:"data,omitempty"`
	ConfigPatches []ConfigPatches `json:"configPatches,omitempty"`
	// Talos Linux machine configuration strategic merge patch list.
	StrategicPatches []string `json:"strategicPatches,omitempty"`
	// Set hostname in the machine configuration to some value.
	Hostname HostnameSpec `json:"hostname,omitempty"`
	// List of variables to expose to the config template.
	// Variables can be used in StrategicPatches and data fields.
	// GoTemplate must be enabled for this to work.
	Variables []Variable `json:"variables,omitempty"`
	// Important: Run "make" to regenerate code after modifying this file
}

// Variable is a definition of a variable to be exposed to the config template.
type Variable struct {
	// Name of the variable to be exposed to the template.
	// +required
	Name string `json:"name"`
	// Value of the variable to be exposed to the template.
	// +optional
	Value string `json:"value,omitempty"`
	// Specifies a source the value of this var should come from.
	// +optional
	ValueFrom *VariableValueSource `json:"valueFrom,omitempty"`
}

type VariableValueSource struct {
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// HostnameSource is the definition of hostname source.
type HostnameSource string

// HostnameSourceMachineName sets the hostname in the generated configuration to the machine name.
const HostnameSourceMachineName HostnameSource = "MachineName"

// HostnameSourceInfrastructureName sets the hostname in the generated configuration to the name of the machine's infrastructure.
const HostnameSourceInfrastructureName HostnameSource = "InfrastructureName"

// HostnameSpec defines the hostname source.
type HostnameSpec struct {
	// Source of the hostname.
	//
	// Allowed values:
	// "MachineName" (use linked Machine's Name).
	// "InfrastructureName" (use linked Machine's infrastructure's name).
	Source HostnameSource `json:"source,omitempty"`
}

// TalosConfigStatus defines the observed state of TalosConfig
type TalosConfigStatus struct {
	// Ready indicates the BootstrapData field is ready to be consumed
	Ready bool `json:"ready,omitempty"`

	// DataSecretName is the name of the secret that stores the bootstrap data script.
	// +optional
	DataSecretName *string `json:"dataSecretName,omitempty"`

	// Talos config will be a string containing the config for download.
	//
	// Deprecated: please use `<cluster>-talosconfig` secret.
	//
	// +optional
	TalosConfig string `json:"talosConfig,omitempty"`

	// FailureReason will be set on non-retryable errors
	// +optional
	FailureReason string `json:"failureReason,omitempty"`

	// FailureMessage will be set on non-retryable errors
	// +optional
	FailureMessage string `json:"failureMessage,omitempty"`

	// ObservedGeneration is the latest generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions defines current service state of the TalosConfig.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=talosconfigs,scope=Namespaced,categories=cluster-api
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// TalosConfig is the Schema for the talosconfigs API
type TalosConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TalosConfigSpec   `json:"spec,omitempty"`
	Status TalosConfigStatus `json:"status,omitempty"`
}

// GetConditions returns the set of conditions for this object.
func (c *TalosConfig) GetConditions() []metav1.Condition {
	return c.Status.Conditions
}

// SetConditions sets the conditions on this object.
func (c *TalosConfig) SetConditions(conditions []metav1.Condition) {
	c.Status.Conditions = conditions
}

// +kubebuilder:object:root=true

// TalosConfigList contains a list of TalosConfig
type TalosConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TalosConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TalosConfig{}, &TalosConfigList{})
}
