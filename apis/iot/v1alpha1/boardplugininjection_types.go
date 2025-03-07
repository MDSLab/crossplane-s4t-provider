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

// BoardPluginInjectionParameters are the configurable fields of a BoardPluginInjection.
type BoardPluginInjectionParameters struct {
	// +kubebuilder:validation:Immutable
	BoardUuid string `json:"boardUuid,omitempty"`
	// +kubebuilder:validation:Immutable
	PluginUuid string `json:"pluginUuid,omitempty"`
}

// BoardPluginInjectionObservation are the observable fields of a BoardPluginInjection.
type BoardPluginInjectionObservation struct {
	BoardUuid  string `json:"boardUuid,omitempty"`
	PluginUuid string `json:"pluginUuid,omitempty"`
}

// A BoardPluginInjectionSpec defines the desired state of a BoardPluginInjection.
type BoardPluginInjectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BoardPluginInjectionParameters `json:"forProvider"`
}

// A BoardPluginInjectionStatus represents the observed state of a BoardPluginInjection.
type BoardPluginInjectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BoardPluginInjectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A BoardPluginInjection is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,s4t}
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BoardPluginInjection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BoardPluginInjectionSpec   `json:"spec"`
	Status BoardPluginInjectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BoardPluginInjectionList contains a list of BoardPluginInjection
type BoardPluginInjectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BoardPluginInjection `json:"items"`
}

// BoardPluginInjection type metadata.
var (
	BoardPluginInjectionKind             = reflect.TypeOf(BoardPluginInjection{}).Name()
	BoardPluginInjectionGroupKind        = schema.GroupKind{Group: Group, Kind: BoardPluginInjectionKind}.String()
	BoardPluginInjectionKindAPIVersion   = BoardPluginInjectionKind + "." + SchemeGroupVersion.String()
	BoardPluginInjectionGroupVersionKind = SchemeGroupVersion.WithKind(BoardPluginInjectionKind)
)

func init() {
	SchemeBuilder.Register(&BoardPluginInjection{}, &BoardPluginInjectionList{})
}
