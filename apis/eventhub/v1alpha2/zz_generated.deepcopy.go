// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPRuleObservation) DeepCopyInto(out *IPRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPRuleObservation.
func (in *IPRuleObservation) DeepCopy() *IPRuleObservation {
	if in == nil {
		return nil
	}
	out := new(IPRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPRuleParameters) DeepCopyInto(out *IPRuleParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.IPMask != nil {
		in, out := &in.IPMask, &out.IPMask
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPRuleParameters.
func (in *IPRuleParameters) DeepCopy() *IPRuleParameters {
	if in == nil {
		return nil
	}
	out := new(IPRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Namespace) DeepCopyInto(out *Namespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Namespace.
func (in *Namespace) DeepCopy() *Namespace {
	if in == nil {
		return nil
	}
	out := new(Namespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Namespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceList) DeepCopyInto(out *NamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Namespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceList.
func (in *NamespaceList) DeepCopy() *NamespaceList {
	if in == nil {
		return nil
	}
	out := new(NamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceObservation) DeepCopyInto(out *NamespaceObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceObservation.
func (in *NamespaceObservation) DeepCopy() *NamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(NamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceParameters) DeepCopyInto(out *NamespaceParameters) {
	*out = *in
	if in.AutoInflateEnabled != nil {
		in, out := &in.AutoInflateEnabled, &out.AutoInflateEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int64)
		**out = **in
	}
	if in.DedicatedClusterID != nil {
		in, out := &in.DedicatedClusterID, &out.DedicatedClusterID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaximumThroughputUnits != nil {
		in, out := &in.MaximumThroughputUnits, &out.MaximumThroughputUnits
		*out = new(int64)
		**out = **in
	}
	if in.NetworkRulesets != nil {
		in, out := &in.NetworkRulesets, &out.NetworkRulesets
		*out = make([]NetworkRulesetsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ZoneRedundant != nil {
		in, out := &in.ZoneRedundant, &out.ZoneRedundant
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceParameters.
func (in *NamespaceParameters) DeepCopy() *NamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(NamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSpec) DeepCopyInto(out *NamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSpec.
func (in *NamespaceSpec) DeepCopy() *NamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceStatus) DeepCopyInto(out *NamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceStatus.
func (in *NamespaceStatus) DeepCopy() *NamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRulesetsObservation) DeepCopyInto(out *NetworkRulesetsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRulesetsObservation.
func (in *NetworkRulesetsObservation) DeepCopy() *NetworkRulesetsObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkRulesetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRulesetsParameters) DeepCopyInto(out *NetworkRulesetsParameters) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IPRule != nil {
		in, out := &in.IPRule, &out.IPRule
		*out = make([]IPRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TrustedServiceAccessEnabled != nil {
		in, out := &in.TrustedServiceAccessEnabled, &out.TrustedServiceAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.VirtualNetworkRule != nil {
		in, out := &in.VirtualNetworkRule, &out.VirtualNetworkRule
		*out = make([]VirtualNetworkRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRulesetsParameters.
func (in *NetworkRulesetsParameters) DeepCopy() *NetworkRulesetsParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkRulesetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleObservation) DeepCopyInto(out *VirtualNetworkRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleObservation.
func (in *VirtualNetworkRuleObservation) DeepCopy() *VirtualNetworkRuleObservation {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleParameters) DeepCopyInto(out *VirtualNetworkRuleParameters) {
	*out = *in
	if in.IgnoreMissingVirtualNetworkServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVirtualNetworkServiceEndpoint, &out.IgnoreMissingVirtualNetworkServiceEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleParameters.
func (in *VirtualNetworkRuleParameters) DeepCopy() *VirtualNetworkRuleParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleParameters)
	in.DeepCopyInto(out)
	return out
}
