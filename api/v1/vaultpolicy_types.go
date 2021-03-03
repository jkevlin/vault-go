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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PolicyPath defines the vault path and acl
type PolicyPath struct {
	Name         string   `json:"path,omitempty" yaml:"path,omitempty" hcl:",key"`
	Capabilities []string `json:"capabilities,omitempty" yaml:"capabilities,omitempty" hcl:"capabilities"`
}

// HCLPolicies is an array of policies
type HCLPolicies struct {
	Paths []PolicyPath `json:"paths,omitempty" yaml:"paths,omitempty" hcl:"path"`
}

// VaultPolicySpec defines a vault policy
type VaultPolicySpec struct {
	VaultNamespace string      `json:"vaultNamespace,omitempty" yaml:"vaultNamespace,omitempty"`
	PolicyName     string      `json:"policyName,omitempty" yaml:"policyName,omitempty"`
	Policies       HCLPolicies `json:"policies" yaml:"policies"`
}

// VaultPolicyStatus defines the observed state of VaultPolicy
type VaultPolicyStatus struct {
}

// +kubebuilder:object:root=true

// VaultPolicy is the Schema for the vaultpolicies API
type VaultPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VaultPolicySpec   `json:"spec,omitempty"`
	Status VaultPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultPolicyList contains a list of VaultPolicy
type VaultPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VaultPolicy{}, &VaultPolicyList{})
}
