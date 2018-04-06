package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Player struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   PlayerSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status PlayerStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PlayerSpec struct {
	Name string `json:"name"`
}

type PlayerStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlayerList is a list of Player objects
type PlayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Player `json:"items" protobuf:"bytes,2,rep,name=items"`
}