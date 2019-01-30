// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "istio.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubetype-gen
// +kubetype-gen:groupVersion=rbac.istio.io/v1alpha1
// +kubetype-gen:kubeType=RbacConfig
// +kubetype-gen:kubeType=ClusterRbacConfig
// +kubetype-gen:ClusterRbacConfig:tag=genclient:nonNamespaced
// +genclient
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RbacConfig defines the global config to control Istio RBAC behavior.
// This Custom Resource is a singleton where only one Custom Resource should be created globally in
// the mesh and the namespace should be the same to other Istio components, which usually is istio-system.
// Note: This is enforced in both istioctl and server side, new Custom Resource will be rejected if found any
// existing one, the user should either delete the existing one or change the existing one directly.
//
// Below is an example of RbacConfig object "istio-rbac-config" which enables Istio RBAC for all
// services in the default namespace.
//
// ```yaml
// apiVersion: "rbac.istio.io/v1alpha1"
// kind: RbacConfig
// metadata:
//   name: default
//   namespace: istio-system
// spec:
//   mode: ON_WITH_INCLUSION
//   inclusion:
//     namespaces: [ "default" ]
// ```
type RbacConfig struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1alpha1.RbacConfig `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RbacConfigList is a collection of RbacConfigs.
type RbacConfigList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []RbacConfig `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +kubetype-gen
// +kubetype-gen:groupVersion=rbac.istio.io/v1alpha1
// +kubetype-gen:kubeType=RbacConfig
// +kubetype-gen:kubeType=ClusterRbacConfig
// +kubetype-gen:ClusterRbacConfig:tag=genclient:nonNamespaced
// +genclient
// +k8s:deepcopy-gen=true
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RbacConfig defines the global config to control Istio RBAC behavior.
// This Custom Resource is a singleton where only one Custom Resource should be created globally in
// the mesh and the namespace should be the same to other Istio components, which usually is istio-system.
// Note: This is enforced in both istioctl and server side, new Custom Resource will be rejected if found any
// existing one, the user should either delete the existing one or change the existing one directly.
//
// Below is an example of RbacConfig object "istio-rbac-config" which enables Istio RBAC for all
// services in the default namespace.
//
// ```yaml
// apiVersion: "rbac.istio.io/v1alpha1"
// kind: RbacConfig
// metadata:
//   name: default
//   namespace: istio-system
// spec:
//   mode: ON_WITH_INCLUSION
//   inclusion:
//     namespaces: [ "default" ]
// ```
type ClusterRbacConfig struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1alpha1.RbacConfig `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterRbacConfigList is a collection of ClusterRbacConfigs.
type ClusterRbacConfigList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []ClusterRbacConfig `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +kubetype-gen
// +kubetype-gen:groupVersion=rbac.istio.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRole specification contains a list of access rules (permissions).
type ServiceRole struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1alpha1.ServiceRole `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRoleList is a collection of ServiceRoles.
type ServiceRoleList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []ServiceRole `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +kubetype-gen
// +kubetype-gen:groupVersion=rbac.istio.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRoleBinding assigns a ServiceRole to a list of subjects.
type ServiceRoleBinding struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1alpha1.ServiceRoleBinding `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRoleBindingList is a collection of ServiceRoleBindings.
type ServiceRoleBindingList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []ServiceRoleBinding `json:"items" protobuf:"bytes,2,rep,name=items"`
}