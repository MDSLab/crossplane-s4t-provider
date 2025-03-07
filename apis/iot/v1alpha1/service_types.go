/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ServiceParameters are the configurable fields of a Service.
type ServiceParameters struct {
	// +kubebuilder:validation:Immutable
	Uuid     string `json:"uuid,omitempty"`
	Name     string `json:"name"`
	Project  string `json:"project,omitempty"`
	Port     uint   `json:"port"`
	Protocol string `json:"protocol"`
}

// ServiceObservation are the observable fields of a Service.
type ServiceObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A ServiceSpec defines the desired state of a Service.
type ServiceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServiceParameters `json:"forProvider"`
}

// A ServiceStatus represents the observed state of a Service.
type ServiceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Service is an example API type.
// +kubebuilder:printcolumn:name="Name",type=string,JSONPath=".spec.Name"
// +kubebuilder:printcolumn:name="Port",type=string,JSONPath=".spec.Port"
// +kubebuilder:printcolumn:name="Protocol",type=string,JSONPath=".spec.Protocol"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,s4t}
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceSpec   `json:"spec"`
	Status ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Service
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Service type metadata.
var (
	ServiceKind             = reflect.TypeOf(Service{}).Name()
	ServiceGroupKind        = schema.GroupKind{Group: Group, Kind: ServiceKind}.String()
	ServiceKindAPIVersion   = ServiceKind + "." + SchemeGroupVersion.String()
	ServiceGroupVersionKind = SchemeGroupVersion.WithKind(ServiceKind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
