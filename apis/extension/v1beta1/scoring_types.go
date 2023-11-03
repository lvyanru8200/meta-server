/*
Copyright 2023.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Question defines a Scoring CRD's questions and desired anwsers
type Question struct {
	// +kubebuilder:validation:Required
	// Question represents questions or instructions used in the evaluation process
	Question string `json:"question"`
	// +kubebuilder:validation:Required
	// Reference represents answers or outputs for the question in the evaluation process
	Reference string `json:"reference"`
}

// ScoringSpec defines the desired state of Scoring
type ScoringSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +kubebuilder:default:=false
	// LoadPlugin describes a Scoring CR whether uses plugin to do evaluation, if true then Plugin data will be needed,
	// otherwise Questions data will be needed.
	LoadPlugin bool `json:"loadPlugin,omitempty"`
	// Plugin describes the plugin including parameters
	Plugin Plugin `json:"plugin,omitempty"`
	// Questions describes the questions uses for the evaluation in case of none plugin used.
	Questions []Question `json:"questions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Scoring is the Schema for the scorings API
type Scoring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ScoringSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScoringList contains a list of Scoring
type ScoringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scoring `json:"items"`
}
