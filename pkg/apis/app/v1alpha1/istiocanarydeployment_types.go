package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IstioCanaryDeploymentSpec defines the desired state of IstioCanaryDeployment
type IstioCanaryDeploymentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	DeploymentSpec appsv1.DeploymentSpec `json:"deploymentSpec"`
	ServiceSpec    corev1.ServiceSpec    `json:"serviceSpec"`
	VSName         string                `json:"vsName"`
}

// IstioCanaryDeploymentStatus defines the observed state of IstioCanaryDeployment
type IstioCanaryDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IstioCanaryDeployment is the Schema for the istiocanarydeployments API
// +k8s:openapi-gen=true
type IstioCanaryDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IstioCanaryDeploymentSpec   `json:"spec,omitempty"`
	Status IstioCanaryDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IstioCanaryDeploymentList contains a list of IstioCanaryDeployment
type IstioCanaryDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IstioCanaryDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IstioCanaryDeployment{}, &IstioCanaryDeploymentList{})
}
