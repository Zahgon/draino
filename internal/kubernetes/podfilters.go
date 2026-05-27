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
)

// A PodFilterFunc returns true if the supplied pod passes the filter.
type PodFilterFunc func(p core.Pod) (bool, error)

// MirrorPodFilter returns true if the supplied pod is not a mirror pod, i.e. a
// pod created by a manifest on the node rather than the API server.
func MirrorPodFilter(p core.Pod) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// LocalStoragePodFilter returns true if the supplied pod does not have local
// storage, i.e. does not use any 'empty dir' volumes.
func LocalStoragePodFilter(p core.Pod) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// UnreplicatedPodFilter returns true if the pod is replicated, i.e. is managed
// by a controller (deployment, daemonset, statefulset, etc) of some sort.
func UnreplicatedPodFilter(p core.Pod) (bool, error) {
	_ = "STUB: not implemented"
	// We're fine with 'evicting' unreplicated pods that aren't actually running.
	return false, nil
}

// NewDaemonSetPodFilter returns a FilterFunc that returns true if the supplied
// pod is not managed by an extant DaemonSet.
func NewDaemonSetPodFilter(client kubernetes.Interface) PodFilterFunc {
	_ = "STUB: not implemented"
	return *new(PodFilterFunc)
}

// Pods pass the filter if they were created by a DaemonSet that no
// longer exists.

// NewStatefulSetPodFilter returns a FilterFunc that returns true if the supplied
// pod is not managed by an extant StatefulSet.
func NewStatefulSetPodFilter(client kubernetes.Interface) PodFilterFunc {
	_ = "STUB: not implemented"
	return *new(PodFilterFunc)
}

// Pods pass the filter if they were created by a StatefulSet that no
// longer exists.

// UnprotectedPodFilter returns a FilterFunc that returns true if the
// supplied pod does not have any of the user-specified annotations for
// protection from eviction
func UnprotectedPodFilter(annotations ...string) PodFilterFunc {
	_ = "STUB: not implemented"
	return *new(PodFilterFunc)
}

// Try to split the annotation into key-value pairs

// If the annotation is a single string, then simply check for
// the existence of the annotation key

// If the annotation is a key-value pair, then check if the
// value for the pod annotation matches that of the
// user-specified value

// NewPodFilters returns a FilterFunc that returns true if all of the supplied
// FilterFuncs return true.
func NewPodFilters(filters ...PodFilterFunc) PodFilterFunc {
	_ = "STUB: not implemented"
	return *new(PodFilterFunc)
}
