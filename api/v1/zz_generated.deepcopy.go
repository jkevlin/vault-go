// +build !ignore_autogenerated

/*


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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCLPolicies) DeepCopyInto(out *HCLPolicies) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]PolicyPath, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCLPolicies.
func (in *HCLPolicies) DeepCopy() *HCLPolicies {
	if in == nil {
		return nil
	}
	out := new(HCLPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyPath) DeepCopyInto(out *PolicyPath) {
	*out = *in
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyPath.
func (in *PolicyPath) DeepCopy() *PolicyPath {
	if in == nil {
		return nil
	}
	out := new(PolicyPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuth) DeepCopyInto(out *VaultAuth) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuth.
func (in *VaultAuth) DeepCopy() *VaultAuth {
	if in == nil {
		return nil
	}
	out := new(VaultAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultAuth) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuthList) DeepCopyInto(out *VaultAuthList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultAuth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuthList.
func (in *VaultAuthList) DeepCopy() *VaultAuthList {
	if in == nil {
		return nil
	}
	out := new(VaultAuthList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultAuthList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuthSpec) DeepCopyInto(out *VaultAuthSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuthSpec.
func (in *VaultAuthSpec) DeepCopy() *VaultAuthSpec {
	if in == nil {
		return nil
	}
	out := new(VaultAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultAuthStatus) DeepCopyInto(out *VaultAuthStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultAuthStatus.
func (in *VaultAuthStatus) DeepCopy() *VaultAuthStatus {
	if in == nil {
		return nil
	}
	out := new(VaultAuthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultNamespace) DeepCopyInto(out *VaultNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultNamespace.
func (in *VaultNamespace) DeepCopy() *VaultNamespace {
	if in == nil {
		return nil
	}
	out := new(VaultNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultNamespaceList) DeepCopyInto(out *VaultNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultNamespaceList.
func (in *VaultNamespaceList) DeepCopy() *VaultNamespaceList {
	if in == nil {
		return nil
	}
	out := new(VaultNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultNamespaceSpec) DeepCopyInto(out *VaultNamespaceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultNamespaceSpec.
func (in *VaultNamespaceSpec) DeepCopy() *VaultNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(VaultNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultNamespaceStatus) DeepCopyInto(out *VaultNamespaceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultNamespaceStatus.
func (in *VaultNamespaceStatus) DeepCopy() *VaultNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(VaultNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicy) DeepCopyInto(out *VaultPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicy.
func (in *VaultPolicy) DeepCopy() *VaultPolicy {
	if in == nil {
		return nil
	}
	out := new(VaultPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicyList) DeepCopyInto(out *VaultPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicyList.
func (in *VaultPolicyList) DeepCopy() *VaultPolicyList {
	if in == nil {
		return nil
	}
	out := new(VaultPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicySpec) DeepCopyInto(out *VaultPolicySpec) {
	*out = *in
	in.Policies.DeepCopyInto(&out.Policies)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicySpec.
func (in *VaultPolicySpec) DeepCopy() *VaultPolicySpec {
	if in == nil {
		return nil
	}
	out := new(VaultPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicyStatus) DeepCopyInto(out *VaultPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicyStatus.
func (in *VaultPolicyStatus) DeepCopy() *VaultPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(VaultPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultRole) DeepCopyInto(out *VaultRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultRole.
func (in *VaultRole) DeepCopy() *VaultRole {
	if in == nil {
		return nil
	}
	out := new(VaultRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultRoleData) DeepCopyInto(out *VaultRoleData) {
	*out = *in
	if in.BoundServiceAccountNames != nil {
		in, out := &in.BoundServiceAccountNames, &out.BoundServiceAccountNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BoundServiceAccountNamespaces != nil {
		in, out := &in.BoundServiceAccountNamespaces, &out.BoundServiceAccountNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecretIDBoundCIDRs != nil {
		in, out := &in.SecretIDBoundCIDRs, &out.SecretIDBoundCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TokenBoundCIDRs != nil {
		in, out := &in.TokenBoundCIDRs, &out.TokenBoundCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TokenPolicies != nil {
		in, out := &in.TokenPolicies, &out.TokenPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultRoleData.
func (in *VaultRoleData) DeepCopy() *VaultRoleData {
	if in == nil {
		return nil
	}
	out := new(VaultRoleData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultRoleList) DeepCopyInto(out *VaultRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultRoleList.
func (in *VaultRoleList) DeepCopy() *VaultRoleList {
	if in == nil {
		return nil
	}
	out := new(VaultRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultRoleSpec) DeepCopyInto(out *VaultRoleSpec) {
	*out = *in
	in.Data.DeepCopyInto(&out.Data)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultRoleSpec.
func (in *VaultRoleSpec) DeepCopy() *VaultRoleSpec {
	if in == nil {
		return nil
	}
	out := new(VaultRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultRoleStatus) DeepCopyInto(out *VaultRoleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultRoleStatus.
func (in *VaultRoleStatus) DeepCopy() *VaultRoleStatus {
	if in == nil {
		return nil
	}
	out := new(VaultRoleStatus)
	in.DeepCopyInto(out)
	return out
}
