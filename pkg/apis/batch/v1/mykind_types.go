
/*
Copyright 2017 The Kubernetes Authors.

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
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/the-redback/apiserver-builded/pkg/apis/batch"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MyKind
// +k8s:openapi-gen=true
// +resource:path=mykinds,strategy=MyKindStrategy
type MyKind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyKindSpec   `json:"spec,omitempty"`
	Status MyKindStatus `json:"status,omitempty"`
}

// MyKindSpec defines the desired state of MyKind
type MyKindSpec struct {
}

// MyKindStatus defines the observed state of MyKind
type MyKindStatus struct {
}

// Validate checks that an instance of MyKind is well formed
func (MyKindStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*batch.MyKind)
	log.Printf("Validating fields for MyKind %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default MyKind field values
func (MyKindSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*MyKind)
	// set default field values here
	log.Printf("Defaulting fields for MyKind %s\n", obj.Name)
}
