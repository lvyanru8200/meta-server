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

type FinetuneExperimentState string

const (
	FinetuneExperimentProcessing FinetuneExperimentState = "PROCESSING"
	FinetuneExperimentSuccess    FinetuneExperimentState = "SUCCESS"
	FinetuneExperimentFailed     FinetuneExperimentState = "FAILED"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FinetuneExperiment is the Schema for the finetuneexperiments API
type FinetuneExperiment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FinetuneExperimentSpec   `json:"spec,omitempty"`
	Status FinetuneExperimentStatus `json:"status,omitempty"`
}

// FinetuneExperimentSpec defines the desired state of FinetuneExperiment
type FinetuneExperimentSpec struct {
	// Defining multiple finetunejobs in a single experiment.
	// +kubebuilder:validation:Required
	FinetuneJobs []FinetuneJobSetting `json:"finetuneJobs"`
	// Define the scoring plugin used for this experiment.
	// +kubebuilder:validation:Required
	ScoringConfig ScoringConfig `json:"scoringConfig"`
}

type FinetuneJobSetting struct {
	// +kubebuilder:validation:Optional
	Name string `json:"name"`
	// +kubebuilder:validation:Required
	Spec FinetuneJobSpec `json:"spec"`
}

// FinetuneExperimentStatus defines the observed state of FinetuneExperiment
type FinetuneExperimentStatus struct {
	BestVersion BestVersion             `json:"bestVersion"`
	JobStates   []FinetuneJobStatus     `json:"jobStates"`
	State       FinetuneExperimentState `json:"state"`
	Conditions  []metav1.Condition      `json:"conditions"`
}

// Describe the highest scoring version of an experiment
type BestVersion struct {
	Score          string `json:"score"`
	Image          string `json:"image"`
	LLM            string `json:"llm"`
	Hyperparameter string `json:"hyperparameter"`
	Dataset        string `json:"dataset"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FinetuneExperimentList contains a list of FinetuneExperiment
type FinetuneExperimentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FinetuneExperiment `json:"items"`
}
