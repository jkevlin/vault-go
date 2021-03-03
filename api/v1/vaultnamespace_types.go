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

// VaultNamespaceSpec defines the desired state of VaultNamespace
type VaultNamespaceSpec struct {
	// NamespaceBase specifies the base path of the namespace. Use "root" for root or no namespace.
	NamespaceBase string `json:"namespaceBase,omitempty" yaml:"namespaceBase,omitempty"`
	NamespaceName string `json:"namespaceName,omitempty" yaml:"namespaceName,omitempty"`
}

// VaultNamespaceStatus defines the observed state of VaultNamespace
type VaultNamespaceStatus struct {
}

// +kubebuilder:object:root=true

// VaultNamespace is the Schema for the vaultnamespaces API
type VaultNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VaultNamespaceSpec   `json:"spec,omitempty"`
	Status VaultNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultNamespaceList contains a list of VaultNamespace
type VaultNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultNamespace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VaultNamespace{}, &VaultNamespaceList{})
}
