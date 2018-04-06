package tryapi

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Player struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   PlayerSpec
	Status PlayerStatus
}

type PlayerSpec struct {
	Name string
}

type PlayerStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlayerList is a list of Player objects
type PlayerList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Player
}