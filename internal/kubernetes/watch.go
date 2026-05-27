/*
Copyright 2018 Planet Labs Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing permissions
and limitations under the License.
*/

package kubernetes

import (
	core "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// An NodeStore is a cache of node resources.
type NodeStore interface {
	// Get an node by name. Returns an error if the node does not exist.
	Get(name string) (*core.Node, error)
}

// An NodeWatch is a cache of node resources that notifies registered
// handlers when its contents change.
type NodeWatch struct {
	cache.SharedInformer
}

// NewNodeWatch creates a watch on node resources. Nodes are cached and the
// provided ResourceEventHandlers are called when the cache changes.
func NewNodeWatch(c kubernetes.Interface, rs ...cache.ResourceEventHandler) *NodeWatch {
	_ = "STUB: not implemented"
	return nil
}

// Get an node by name. Returns an error if the node does not exist.
func (w *NodeWatch) Get(name string) (*core.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
