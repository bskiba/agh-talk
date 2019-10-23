package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Shout is a specification for a Foo resource
type Shout struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ShoutSpec   `json:"spec"`
	Status ShoutStatus `json:"status"`
}

// ShoutSpec is the spec for a Shout resource
type ShoutSpec struct {
	Text  string `json:"text"`
	Count       *int32 `json:"count"`
}

// ShoutStatus is the status for a Shout resource
type ShoutStatus struct {
	Written bool `json:"written"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type ShoutList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Shout `json:"items"`
}
