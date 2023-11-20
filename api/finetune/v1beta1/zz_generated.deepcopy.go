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
	v1 "kubernetes/staging/src/k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BestVersion) DeepCopyInto(out *BestVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BestVersion.
func (in *BestVersion) DeepCopy() *BestVersion {
	if in == nil {
		return nil
	}
	out := new(BestVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FineTune) DeepCopyInto(out *FineTune) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	in.FinetuneSpec.DeepCopyInto(&out.FinetuneSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FineTune.
func (in *FineTune) DeepCopy() *FineTune {
	if in == nil {
		return nil
	}
	out := new(FineTune)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Finetune) DeepCopyInto(out *Finetune) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Finetune.
func (in *Finetune) DeepCopy() *Finetune {
	if in == nil {
		return nil
	}
	out := new(Finetune)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Finetune) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneExperiment) DeepCopyInto(out *FinetuneExperiment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneExperiment.
func (in *FinetuneExperiment) DeepCopy() *FinetuneExperiment {
	if in == nil {
		return nil
	}
	out := new(FinetuneExperiment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FinetuneExperiment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneExperimentList) DeepCopyInto(out *FinetuneExperimentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FinetuneExperiment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneExperimentList.
func (in *FinetuneExperimentList) DeepCopy() *FinetuneExperimentList {
	if in == nil {
		return nil
	}
	out := new(FinetuneExperimentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FinetuneExperimentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneExperimentSpec) DeepCopyInto(out *FinetuneExperimentSpec) {
	*out = *in
	if in.FinetuneJobs != nil {
		in, out := &in.FinetuneJobs, &out.FinetuneJobs
		*out = make([]FinetuneJobSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ScoringConfig = in.ScoringConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneExperimentSpec.
func (in *FinetuneExperimentSpec) DeepCopy() *FinetuneExperimentSpec {
	if in == nil {
		return nil
	}
	out := new(FinetuneExperimentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneExperimentStatus) DeepCopyInto(out *FinetuneExperimentStatus) {
	*out = *in
	if in.BestVersion != nil {
		in, out := &in.BestVersion, &out.BestVersion
		*out = new(BestVersion)
		**out = **in
	}
	if in.JobsStatus != nil {
		in, out := &in.JobsStatus, &out.JobsStatus
		*out = make([]FinetuneJobStatusSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneExperimentStatus.
func (in *FinetuneExperimentStatus) DeepCopy() *FinetuneExperimentStatus {
	if in == nil {
		return nil
	}
	out := new(FinetuneExperimentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJob) DeepCopyInto(out *FinetuneJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJob.
func (in *FinetuneJob) DeepCopy() *FinetuneJob {
	if in == nil {
		return nil
	}
	out := new(FinetuneJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FinetuneJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJobList) DeepCopyInto(out *FinetuneJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FinetuneJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJobList.
func (in *FinetuneJobList) DeepCopy() *FinetuneJobList {
	if in == nil {
		return nil
	}
	out := new(FinetuneJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FinetuneJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJobResult) DeepCopyInto(out *FinetuneJobResult) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJobResult.
func (in *FinetuneJobResult) DeepCopy() *FinetuneJobResult {
	if in == nil {
		return nil
	}
	out := new(FinetuneJobResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJobSetting) DeepCopyInto(out *FinetuneJobSetting) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJobSetting.
func (in *FinetuneJobSetting) DeepCopy() *FinetuneJobSetting {
	if in == nil {
		return nil
	}
	out := new(FinetuneJobSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJobSpec) DeepCopyInto(out *FinetuneJobSpec) {
	*out = *in
	in.FineTune.DeepCopyInto(&out.FineTune)
	if in.ScoringConfig != nil {
		in, out := &in.ScoringConfig, &out.ScoringConfig
		*out = new(ScoringConfig)
		**out = **in
	}
	in.ServeConfig.DeepCopyInto(&out.ServeConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJobSpec.
func (in *FinetuneJobSpec) DeepCopy() *FinetuneJobSpec {
	if in == nil {
		return nil
	}
	out := new(FinetuneJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJobStatus) DeepCopyInto(out *FinetuneJobStatus) {
	*out = *in
	if in.Result != nil {
		in, out := &in.Result, &out.Result
		*out = new(FinetuneJobResult)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJobStatus.
func (in *FinetuneJobStatus) DeepCopy() *FinetuneJobStatus {
	if in == nil {
		return nil
	}
	out := new(FinetuneJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJobStatusSetting) DeepCopyInto(out *FinetuneJobStatusSetting) {
	*out = *in
	in.FinetuneJobStatus.DeepCopyInto(&out.FinetuneJobStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJobStatusSetting.
func (in *FinetuneJobStatusSetting) DeepCopy() *FinetuneJobStatusSetting {
	if in == nil {
		return nil
	}
	out := new(FinetuneJobStatusSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneList) DeepCopyInto(out *FinetuneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Finetune, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneList.
func (in *FinetuneList) DeepCopy() *FinetuneList {
	if in == nil {
		return nil
	}
	out := new(FinetuneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FinetuneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneSpec) DeepCopyInto(out *FinetuneSpec) {
	*out = *in
	in.Hyperparameter.DeepCopyInto(&out.Hyperparameter)
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(Resource)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSetting)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneSpec.
func (in *FinetuneSpec) DeepCopy() *FinetuneSpec {
	if in == nil {
		return nil
	}
	out := new(FinetuneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneStatus) DeepCopyInto(out *FinetuneStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneStatus.
func (in *FinetuneStatus) DeepCopy() *FinetuneStatus {
	if in == nil {
		return nil
	}
	out := new(FinetuneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hyperparameter) DeepCopyInto(out *Hyperparameter) {
	*out = *in
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = new(Parameters)
		(*in).DeepCopyInto(*out)
	}
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSetting) DeepCopyInto(out *ImageSetting) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSetting.
func (in *ImageSetting) DeepCopy() *ImageSetting {
	if in == nil {
		return nil
	}
	out := new(ImageSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameters) DeepCopyInto(out *Parameters) {
	*out = *in
	if in.Int4 != nil {
		in, out := &in.Int4, &out.Int4
		*out = new(bool)
		**out = **in
	}
	if in.Int8 != nil {
		in, out := &in.Int8, &out.Int8
		*out = new(bool)
		**out = **in
	}
	if in.LoRA_R != nil {
		in, out := &in.LoRA_R, &out.LoRA_R
		*out = new(string)
		**out = **in
	}
	if in.LoRA_Alpha != nil {
		in, out := &in.LoRA_Alpha, &out.LoRA_Alpha
		*out = new(string)
		**out = **in
	}
	if in.LoRA_Dropout != nil {
		in, out := &in.LoRA_Dropout, &out.LoRA_Dropout
		*out = new(string)
		**out = **in
	}
	if in.LearningRate != nil {
		in, out := &in.LearningRate, &out.LearningRate
		*out = new(string)
		**out = **in
	}
	if in.WarmupRatio != nil {
		in, out := &in.WarmupRatio, &out.WarmupRatio
		*out = new(string)
		**out = **in
	}
	if in.WeightDecay != nil {
		in, out := &in.WeightDecay, &out.WeightDecay
		*out = new(string)
		**out = **in
	}
	if in.TrainerType != nil {
		in, out := &in.TrainerType, &out.TrainerType
		*out = new(string)
		**out = **in
	}
	if in.PEFT != nil {
		in, out := &in.PEFT, &out.PEFT
		*out = new(bool)
		**out = **in
	}
	if in.FP16 != nil {
		in, out := &in.FP16, &out.FP16
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameters.
func (in *Parameters) DeepCopy() *Parameters {
	if in == nil {
		return nil
	}
	out := new(Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(ResourceLimits)
		(*in).DeepCopyInto(*out)
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(ResourceLimits)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLimits) DeepCopyInto(out *ResourceLimits) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
	if in.GPU != nil {
		in, out := &in.GPU, &out.GPU
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLimits.
func (in *ResourceLimits) DeepCopy() *ResourceLimits {
	if in == nil {
		return nil
	}
	out := new(ResourceLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoringConfig) DeepCopyInto(out *ScoringConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoringConfig.
func (in *ScoringConfig) DeepCopy() *ScoringConfig {
	if in == nil {
		return nil
	}
	out := new(ScoringConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServeConfig) DeepCopyInto(out *ServeConfig) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServeConfig.
func (in *ServeConfig) DeepCopy() *ServeConfig {
	if in == nil {
		return nil
	}
	out := new(ServeConfig)
	in.DeepCopyInto(out)
	return out
}
