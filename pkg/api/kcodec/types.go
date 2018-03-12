package kcodec

import (
	kmetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//Item resource
type Item struct {
	kmetav1.TypeMeta   `json:",inline"`
	kmetav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ItemSpec `json:"spec,omitempty"`
}

//ItemSpec spec
type ItemSpec struct {
	Value int `json:"message,omitempty"`
}