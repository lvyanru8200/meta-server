//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeInfrastructure) DeepCopyInto(out *ComputeInfrastructure) {
	*out = *in
	if in.Hardware != nil {
		in, out := &in.Hardware, &out.Hardware
		*out = new(Hardware)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeInfrastructure.
func (in *ComputeInfrastructure) DeepCopy() *ComputeInfrastructure {
	if in == nil {
		return nil
	}
	out := new(ComputeInfrastructure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hardware) DeepCopyInto(out *Hardware) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hardware.
func (in *Hardware) DeepCopy() *Hardware {
	if in == nil {
		return nil
	}
	out := new(Hardware)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hyperparameter) DeepCopyInto(out *Hyperparameter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hyperparameter.
func (in *Hyperparameter) DeepCopy() *Hyperparameter {
	if in == nil {
		return nil
	}
	out := new(Hyperparameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hyperparameter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperparameterList) DeepCopyInto(out *HyperparameterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hyperparameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperparameterList.
func (in *HyperparameterList) DeepCopy() *HyperparameterList {
	if in == nil {
		return nil
	}
	out := new(HyperparameterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperparameterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperparameterSpec) DeepCopyInto(out *HyperparameterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperparameterSpec.
func (in *HyperparameterSpec) DeepCopy() *HyperparameterSpec {
	if in == nil {
		return nil
	}
	out := new(HyperparameterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperparameterStatus) DeepCopyInto(out *HyperparameterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperparameterStatus.
func (in *HyperparameterStatus) DeepCopy() *HyperparameterStatus {
	if in == nil {
		return nil
	}
	out := new(HyperparameterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLM) DeepCopyInto(out *LLM) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLM.
func (in *LLM) DeepCopy() *LLM {
	if in == nil {
		return nil
	}
	out := new(LLM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LLM) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMCard) DeepCopyInto(out *LLMCard) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMCard.
func (in *LLMCard) DeepCopy() *LLMCard {
	if in == nil {
		return nil
	}
	out := new(LLMCard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMFiles) DeepCopyInto(out *LLMFiles) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMFiles.
func (in *LLMFiles) DeepCopy() *LLMFiles {
	if in == nil {
		return nil
	}
	out := new(LLMFiles)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMImageConfig) DeepCopyInto(out *LLMImageConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMImageConfig.
func (in *LLMImageConfig) DeepCopy() *LLMImageConfig {
	if in == nil {
		return nil
	}
	out := new(LLMImageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMList) DeepCopyInto(out *LLMList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LLM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMList.
func (in *LLMList) DeepCopy() *LLMList {
	if in == nil {
		return nil
	}
	out := new(LLMList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LLMList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMMetdata) DeepCopyInto(out *LLMMetdata) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = make([]LLMDomain, len(*in))
		copy(*out, *in)
	}
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Frameworks != nil {
		in, out := &in.Frameworks, &out.Frameworks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Languages != nil {
		in, out := &in.Languages, &out.Languages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.License != nil {
		in, out := &in.License, &out.License
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Datasets != nil {
		in, out := &in.Datasets, &out.Datasets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.LLMImageConfig = in.LLMImageConfig
	if in.ComputeInfrastructure != nil {
		in, out := &in.ComputeInfrastructure, &out.ComputeInfrastructure
		*out = new(ComputeInfrastructure)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMMetdata.
func (in *LLMMetdata) DeepCopy() *LLMMetdata {
	if in == nil {
		return nil
	}
	out := new(LLMMetdata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMSpec) DeepCopyInto(out *LLMSpec) {
	*out = *in
	in.LLMMetdata.DeepCopyInto(&out.LLMMetdata)
	if in.LLMCard != nil {
		in, out := &in.LLMCard, &out.LLMCard
		*out = new(LLMCard)
		**out = **in
	}
	if in.LLMFiles != nil {
		in, out := &in.LLMFiles, &out.LLMFiles
		*out = new(LLMFiles)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMSpec.
func (in *LLMSpec) DeepCopy() *LLMSpec {
	if in == nil {
		return nil
	}
	out := new(LLMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMStatus) DeepCopyInto(out *LLMStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMStatus.
func (in *LLMStatus) DeepCopy() *LLMStatus {
	if in == nil {
		return nil
	}
	out := new(LLMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Objective) DeepCopyInto(out *Objective) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Objective.
func (in *Objective) DeepCopy() *Objective {
	if in == nil {
		return nil
	}
	out := new(Objective)
	in.DeepCopyInto(out)
	return out
}
