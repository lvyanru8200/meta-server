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

// HyperparameterSpec defines the desired state of Hyperparameter
type HyperparameterSpec struct {
	// +kubebuilder:validation:Required
	// Over-senator adjustments to achieve targets.
	Objective Objective `json:"objective"`
	// +kubebuilder:validation:Required
	// Finetune paramenter config.
	Parameters Parameters `json:"parameters"`
}

// The goal is usually a specific objective that is desired to be achieved through hyperparameter tuning,
// such as minimising losses or maximising some evaluation metric (e.g. accuracy, recall, etc.)
type Objective struct {
	/*
		The type (Type) of fine-tuning may refer to a specific method or strategy of hyperparameter tuning. Common hyperparameter tuning methods include:
		1. Grid Search
		2. Random Search
		3. Bayesian Optimisation
		4. Genetic Algorithms
		5. Hyperband et al.
	*/
	Type string `json:"type"`
}

type Parameters struct {
	// Scheduler specifies the learning rate scheduler.
	Scheduler string `json:"scheduler"`
	// Optimizer specifies the optimization algorithm.
	Optimizer string `json:"optimizer"`
	// Int4 indicates whether to use 4-bit integer quantization.
	Int4 bool `json:"int4"`
	// Int8 indicates whether to use 8-bit integer quantization.
	Int8 bool `json:"int8"`
	// LoRA_R represents the radius parameter for Localized Receptive Attention.
	LoRA_R string `json:"loRA_R"`
	// LoRA_Alpha represents the alpha parameter for Localized Receptive Attention.
	LoRA_Alpha string `json:"loRA_Alpha"`
	// LoRA_Dropout specifies the dropout rate for Localized Receptive Attention.
	LoRA_Dropout string `json:"loRA_Dropout"`
	// LearningRate specifies the initial learning rate.
	LearningRate string `json:"learningRate"`
	// Epochs specifies the number of training epochs.
	Epochs int `json:"epochs"`
	// BlockSize specifies the block size.
	BlockSize int `json:"blockSize"`
	// BatchSize specifies the size of mini-batches.
	BatchSize int `json:"batchSize"`
	// WarmupRatio specifies the ratio of warmup steps.
	WarmupRatio string `json:"warmupRatio"`
	// WeightDecay specifies the weight decay factor.
	WeightDecay string `json:"weightDecay"`
	// GradAccSteps specifies the number of gradient accumulation steps.
	GradAccSteps int `json:"gradAccSteps"`
	// TrainerType specifies the type of trainer to use.
	TrainerType string `json:"trainerType"`
	// PEFT indicates whether to enable Performance Evaluation and Forecasting Tool.
	PEFT bool `json:"PEFT"`
	// FP16 indicates whether to use 16-bit floating point precision.
	FP16 bool `json:"FP16"`
}

// HyperparameterStatus defines the observed state of Hyperparameter
type HyperparameterStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Hyperparameter is the Schema for the hyperparameters API
type Hyperparameter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HyperparameterSpec   `json:"spec,omitempty"`
	Status HyperparameterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HyperparameterList contains a list of Hyperparameter
type HyperparameterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hyperparameter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Hyperparameter{}, &HyperparameterList{})
}
