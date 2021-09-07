// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2021 Intel Corporation

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFlowAction) DeepCopyInto(out *ClusterFlowAction) {
	*out = *in
	if in.Conf != nil {
		in, out := &in.Conf, &out.Conf
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFlowAction.
func (in *ClusterFlowAction) DeepCopy() *ClusterFlowAction {
	if in == nil {
		return nil
	}
	out := new(ClusterFlowAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFlowConfig) DeepCopyInto(out *ClusterFlowConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFlowConfig.
func (in *ClusterFlowConfig) DeepCopy() *ClusterFlowConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterFlowConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterFlowConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFlowConfigList) DeepCopyInto(out *ClusterFlowConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterFlowConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFlowConfigList.
func (in *ClusterFlowConfigList) DeepCopy() *ClusterFlowConfigList {
	if in == nil {
		return nil
	}
	out := new(ClusterFlowConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterFlowConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFlowConfigSpec) DeepCopyInto(out *ClusterFlowConfigSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]*ClusterFlowRule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ClusterFlowRule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFlowConfigSpec.
func (in *ClusterFlowConfigSpec) DeepCopy() *ClusterFlowConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterFlowConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFlowConfigStatus) DeepCopyInto(out *ClusterFlowConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFlowConfigStatus.
func (in *ClusterFlowConfigStatus) DeepCopy() *ClusterFlowConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterFlowConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFlowRule) DeepCopyInto(out *ClusterFlowRule) {
	*out = *in
	if in.Attr != nil {
		in, out := &in.Attr, &out.Attr
		*out = new(FlowAttr)
		**out = **in
	}
	if in.Pattern != nil {
		in, out := &in.Pattern, &out.Pattern
		*out = make([]*FlowItem, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FlowItem)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = make([]*ClusterFlowAction, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ClusterFlowAction)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFlowRule.
func (in *ClusterFlowRule) DeepCopy() *ClusterFlowRule {
	if in == nil {
		return nil
	}
	out := new(ClusterFlowRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowAction) DeepCopyInto(out *FlowAction) {
	*out = *in
	if in.Conf != nil {
		in, out := &in.Conf, &out.Conf
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowAction.
func (in *FlowAction) DeepCopy() *FlowAction {
	if in == nil {
		return nil
	}
	out := new(FlowAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowAttr) DeepCopyInto(out *FlowAttr) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowAttr.
func (in *FlowAttr) DeepCopy() *FlowAttr {
	if in == nil {
		return nil
	}
	out := new(FlowAttr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowConfigNodeAgentDeployment) DeepCopyInto(out *FlowConfigNodeAgentDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowConfigNodeAgentDeployment.
func (in *FlowConfigNodeAgentDeployment) DeepCopy() *FlowConfigNodeAgentDeployment {
	if in == nil {
		return nil
	}
	out := new(FlowConfigNodeAgentDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlowConfigNodeAgentDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowConfigNodeAgentDeploymentList) DeepCopyInto(out *FlowConfigNodeAgentDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlowConfigNodeAgentDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowConfigNodeAgentDeploymentList.
func (in *FlowConfigNodeAgentDeploymentList) DeepCopy() *FlowConfigNodeAgentDeploymentList {
	if in == nil {
		return nil
	}
	out := new(FlowConfigNodeAgentDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlowConfigNodeAgentDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowConfigNodeAgentDeploymentSpec) DeepCopyInto(out *FlowConfigNodeAgentDeploymentSpec) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(corev1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowConfigNodeAgentDeploymentSpec.
func (in *FlowConfigNodeAgentDeploymentSpec) DeepCopy() *FlowConfigNodeAgentDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(FlowConfigNodeAgentDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowConfigNodeAgentDeploymentStatus) DeepCopyInto(out *FlowConfigNodeAgentDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowConfigNodeAgentDeploymentStatus.
func (in *FlowConfigNodeAgentDeploymentStatus) DeepCopy() *FlowConfigNodeAgentDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(FlowConfigNodeAgentDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowItem) DeepCopyInto(out *FlowItem) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Last != nil {
		in, out := &in.Last, &out.Last
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Mask != nil {
		in, out := &in.Mask, &out.Mask
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowItem.
func (in *FlowItem) DeepCopy() *FlowItem {
	if in == nil {
		return nil
	}
	out := new(FlowItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowRules) DeepCopyInto(out *FlowRules) {
	*out = *in
	if in.Attr != nil {
		in, out := &in.Attr, &out.Attr
		*out = new(FlowAttr)
		**out = **in
	}
	if in.Pattern != nil {
		in, out := &in.Pattern, &out.Pattern
		*out = make([]*FlowItem, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FlowItem)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = make([]*FlowAction, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FlowAction)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowRules.
func (in *FlowRules) DeepCopy() *FlowRules {
	if in == nil {
		return nil
	}
	out := new(FlowRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFlowConfig) DeepCopyInto(out *NodeFlowConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFlowConfig.
func (in *NodeFlowConfig) DeepCopy() *NodeFlowConfig {
	if in == nil {
		return nil
	}
	out := new(NodeFlowConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFlowConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFlowConfigList) DeepCopyInto(out *NodeFlowConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeFlowConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFlowConfigList.
func (in *NodeFlowConfigList) DeepCopy() *NodeFlowConfigList {
	if in == nil {
		return nil
	}
	out := new(NodeFlowConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFlowConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFlowConfigSpec) DeepCopyInto(out *NodeFlowConfigSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]*FlowRules, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FlowRules)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFlowConfigSpec.
func (in *NodeFlowConfigSpec) DeepCopy() *NodeFlowConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NodeFlowConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFlowConfigStatus) DeepCopyInto(out *NodeFlowConfigStatus) {
	*out = *in
	if in.PortInfo != nil {
		in, out := &in.PortInfo, &out.PortInfo
		*out = make([]PortsInformation, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]*FlowRules, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FlowRules)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFlowConfigStatus.
func (in *NodeFlowConfigStatus) DeepCopy() *NodeFlowConfigStatus {
	if in == nil {
		return nil
	}
	out := new(NodeFlowConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortsInformation) DeepCopyInto(out *PortsInformation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortsInformation.
func (in *PortsInformation) DeepCopy() *PortsInformation {
	if in == nil {
		return nil
	}
	out := new(PortsInformation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToPodInterfaceConf) DeepCopyInto(out *ToPodInterfaceConf) {
	*out = *in
	in.PodSelector.DeepCopyInto(&out.PodSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToPodInterfaceConf.
func (in *ToPodInterfaceConf) DeepCopy() *ToPodInterfaceConf {
	if in == nil {
		return nil
	}
	out := new(ToPodInterfaceConf)
	in.DeepCopyInto(out)
	return out
}
