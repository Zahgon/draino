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
	"time"

	"go.uber.org/zap"
	core "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// Default pod eviction settings.
const (
	DefaultMaxGracePeriod   time.Duration = 8 * time.Minute
	DefaultEvictionOverhead time.Duration = 30 * time.Second

	kindDaemonSet   = "DaemonSet"
	kindStatefulSet = "StatefulSet"

	ConditionDrainedScheduled = "DrainScheduled"
	DefaultSkipDrain          = false
)

type nodeMutatorFn func(*core.Node)

type errTimeout struct{}

func (e errTimeout) Error() string { _ = "STUB: not implemented"; return "" }

func (e errTimeout) Timeout() {
	_ = "STUB: not implemented"

	// IsTimeout returns true if the supplied error was caused by a timeout.
	return
}

func IsTimeout(err error) bool { _ = "STUB: not implemented"; return false }

// A Cordoner cordons nodes.
type Cordoner interface {
	// Cordon the supplied node. Marks it unschedulable for new pods.
	Cordon(n *core.Node, mutators ...nodeMutatorFn) error

	// Uncordon the supplied node. Marks it schedulable for new pods.
	Uncordon(n *core.Node, mutators ...nodeMutatorFn) error
}

// A Drainer drains nodes.
type Drainer interface {
	// Drain the supplied node. Evicts the node of all but mirror and DaemonSet pods.
	Drain(n *core.Node) error
	MarkDrain(n *core.Node, when, finish time.Time, failed bool) error
}

// A CordonDrainer both cordons and drains nodes!
type CordonDrainer interface {
	Cordoner
	Drainer
}

// A NoopCordonDrainer does nothing.
type NoopCordonDrainer struct{}

// Cordon does nothing.
func (d *NoopCordonDrainer) Cordon(n *core.Node, mutators ...nodeMutatorFn) error {
	_ = "STUB: not implemented"

	// Uncordon does nothing.
	return nil
}

func (d *NoopCordonDrainer) Uncordon(n *core.Node, mutators ...nodeMutatorFn) error {
	_ = "STUB: not implemented"

	// Drain does nothing.
	return nil
}

func (d *NoopCordonDrainer) Drain(n *core.Node) error {
	_ = "STUB: not implemented"

	// MarkDrain does nothing.
	return nil
}

func (d *NoopCordonDrainer) MarkDrain(n *core.Node, when, finish time.Time, failed bool) error {
	_ = "STUB: not implemented"

	// APICordonDrainer drains Kubernetes nodes via the Kubernetes API.
	return nil
}

type APICordonDrainer struct {
	c kubernetes.Interface
	l *zap.Logger

	filter PodFilterFunc

	maxGracePeriod   time.Duration
	evictionHeadroom time.Duration
	skipDrain        bool
}

// SuppliedCondition defines the condition will be watched.
type SuppliedCondition struct {
	Type            core.NodeConditionType
	Status          core.ConditionStatus
	MinimumDuration time.Duration
}

// APICordonDrainerOption configures an APICordonDrainer.
type APICordonDrainerOption func(d *APICordonDrainer)

// MaxGracePeriod configures the maximum time to wait for a pod eviction. Pod
// containers will be allowed this much time to shutdown once they receive a
// SIGTERM before they are sent a SIGKILL.
func MaxGracePeriod(m time.Duration) APICordonDrainerOption {
	_ = "STUB: not implemented"
	return *new(APICordonDrainerOption)
}

// EvictionHeadroom configures an amount of time to wait in addition to the
// MaxGracePeriod for the API server to report a pod deleted.
func EvictionHeadroom(h time.Duration) APICordonDrainerOption {
	_ = "STUB: not implemented"
	return *new(APICordonDrainerOption)
}

// WithPodFilter configures a filter that may be used to exclude certain pods
// from eviction when draining.
func WithPodFilter(f PodFilterFunc) APICordonDrainerOption {
	_ = "STUB: not implemented"
	return *new(APICordonDrainerOption)
}

// WithDrain determines if we're actually going to drain nodes
func WithSkipDrain(b bool) APICordonDrainerOption {
	_ = "STUB: not implemented"
	return *new(APICordonDrainerOption)
}

// WithAPICordonDrainerLogger configures a APICordonDrainer to use the supplied
// logger.
func WithAPICordonDrainerLogger(l *zap.Logger) APICordonDrainerOption {
	_ = "STUB: not implemented"
	return *new(APICordonDrainerOption)
}

// NewAPICordonDrainer returns a CordonDrainer that cordons and drains nodes via
// the Kubernetes API.
func NewAPICordonDrainer(c kubernetes.Interface, ao ...APICordonDrainerOption) *APICordonDrainer {
	_ = "STUB: not implemented"
	return nil
}

func (d *APICordonDrainer) deleteTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Cordon the supplied node. Marks it unschedulable for new pods.
func (d *APICordonDrainer) Cordon(n *core.Node, mutators ...nodeMutatorFn) error {
	_ = "STUB: not implemented"
	return nil
}

// Uncordon the supplied node. Marks it schedulable for new pods.
func (d *APICordonDrainer) Uncordon(n *core.Node, mutators ...nodeMutatorFn) error {
	_ = "STUB: not implemented"
	return nil
}

// MarkDrain set a condition on the node to mark that that drain is scheduled.
func (d *APICordonDrainer) MarkDrain(n *core.Node, when, finish time.Time, failed bool) error {
	_ = "STUB: not implemented"

	// Refresh the node object
	return nil
}

// Create or update the condition associated to the monitor

// There was no condition found, let's create one

func IsMarkedForDrain(n *core.Node) bool { _ = "STUB: not implemented"; return false }

// Drain the supplied node. Evicts the node of all but mirror and DaemonSet pods.
func (d *APICordonDrainer) Drain(n *core.Node) error {
	_ = "STUB: not implemented"

	// Do nothing if draining is not enabled.
	return nil
}

// This will _eventually_ abort evictions. Evictions may spend up to
// d.deleteTimeout() in d.awaitDeletion(), or 5 seconds in backoff before
// noticing they've been aborted.

func (d *APICordonDrainer) getPods(node string) ([]core.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *APICordonDrainer) evict(p core.Pod, abort <-chan struct{}, e chan<- error) {
	_ = "STUB: not implemented"
	return
}

// The eviction API returns 429 Too Many Requests if a pod
// cannot currently be evicted, for example due to a pod
// disruption budget.

func (d *APICordonDrainer) awaitDeletion(p core.Pod, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
