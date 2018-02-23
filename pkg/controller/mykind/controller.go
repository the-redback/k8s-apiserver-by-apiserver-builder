
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


package mykind

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/the-redback/apiserver-builded/pkg/apis/batch/v1"
	"github.com/the-redback/apiserver-builded/pkg/controller/sharedinformers"
	listers "github.com/the-redback/apiserver-builded/pkg/client/listers_generated/batch/v1"
)

// +controller:group=batch,version=v1,kind=MyKind,resource=mykinds
type MyKindControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about MyKind
	lister listers.MyKindLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *MyKindControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing mykinds labels
	c.lister = arguments.GetSharedInformers().Factory.Batch().V1().MyKinds().Lister()
}

// Reconcile handles enqueued messages
func (c *MyKindControllerImpl) Reconcile(u *v1.MyKind) error {
	// Implement controller logic here
	log.Printf("Running reconcile MyKind for %s\n", u.Name)
	return nil
}

func (c *MyKindControllerImpl) Get(namespace, name string) (*v1.MyKind, error) {
	return c.lister.MyKinds(namespace).Get(name)
}
