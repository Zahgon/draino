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

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.uber.org/zap"
	core "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
)

const (
	// DefaultDrainBuffer is the default minimum time between node drains.
	DefaultDrainBuffer = 10 * time.Minute

	eventReasonCordonStarting  = "CordonStarting"
	eventReasonCordonSucceeded = "CordonSucceeded"
	eventReasonCordonFailed    = "CordonFailed"

	eventReasonUncordonStarting  = "UncordonStarting"
	eventReasonUncordonSucceeded = "UncordonSucceeded"
	eventReasonUncordonFailed    = "UncordonFailed"

	eventReasonDrainScheduled        = "DrainScheduled"
	eventReasonDrainSchedulingFailed = "DrainSchedulingFailed"
	eventReasonDrainStarting         = "DrainStarting"
	eventReasonDrainSucceeded        = "DrainSucceeded"
	eventReasonDrainFailed           = "DrainFailed"

	tagResultSucceeded = "succeeded"
	tagResultFailed    = "failed"

	drainRetryAnnotationKey   = "draino/drain-retry"
	drainRetryAnnotationValue = "true"

	drainoConditionsAnnotationKey = "draino.planet.com/conditions"
)

// Opencensus measurements.
var (
	MeasureNodesCordoned       = stats.Int64("draino/nodes_cordoned", "Number of nodes cordoned.", stats.UnitDimensionless)
	MeasureNodesUncordoned     = stats.Int64("draino/nodes_uncordoned", "Number of nodes uncordoned.", stats.UnitDimensionless)
	MeasureNodesDrained        = stats.Int64("draino/nodes_drained", "Number of nodes drained.", stats.UnitDimensionless)
	MeasureNodesDrainScheduled = stats.Int64("draino/nodes_drainScheduled", "Number of nodes drain scheduled.", stats.UnitDimensionless)

	TagNodeName, _ = tag.NewKey("node_name")
	TagResult, _   = tag.NewKey("result")
)

// A DrainingResourceEventHandler cordons and drains any added or updated nodes.
type DrainingResourceEventHandler struct {
	logger         *zap.Logger
	cordonDrainer  CordonDrainer
	eventRecorder  record.EventRecorder
	drainScheduler DrainScheduler

	lastDrainScheduledFor time.Time
	buffer                time.Duration

	conditions []SuppliedCondition
}

// DrainingResourceEventHandlerOption configures an DrainingResourceEventHandler.
type DrainingResourceEventHandlerOption func(d *DrainingResourceEventHandler)

// WithLogger configures a DrainingResourceEventHandler to use the supplied
// logger.
func WithLogger(l *zap.Logger) DrainingResourceEventHandlerOption {
	_ = "STUB: not implemented"
	return *new(DrainingResourceEventHandlerOption)
}

// WithDrainBuffer configures the minimum time between scheduled drains.
func WithDrainBuffer(d time.Duration) DrainingResourceEventHandlerOption {
	_ = "STUB: not implemented"
	return *new(DrainingResourceEventHandlerOption)
}

// WithConditionsFilter configures which conditions should be handled.
func WithConditionsFilter(conditions []string) DrainingResourceEventHandlerOption {
	_ = "STUB: not implemented"
	return *new(DrainingResourceEventHandlerOption)
}

// NewDrainingResourceEventHandler returns a new DrainingResourceEventHandler.
func NewDrainingResourceEventHandler(d CordonDrainer, e record.EventRecorder, ho ...DrainingResourceEventHandlerOption) *DrainingResourceEventHandler {
	_ = "STUB: not implemented"
	return nil
}

// OnAdd cordons and drains the added node.
func (h *DrainingResourceEventHandler) OnAdd(obj interface{}) { _ = "STUB: not implemented"; return }

// OnUpdate cordons and drains the updated node.
func (h *DrainingResourceEventHandler) OnUpdate(_, newObj interface{}) {
	_ = "STUB: not implemented"

	// OnDelete does nothing. There's no point cordoning or draining deleted nodes.
	return
}

func (h *DrainingResourceEventHandler) OnDelete(obj interface{}) { _ = "STUB: not implemented"; return }

func (h *DrainingResourceEventHandler) HandleNode(n *core.Node) { _ = "STUB: not implemented"; return }

// First cordon the node if it is not yet cordonned

// Let's ensure that a drain is scheduled

// Is there a request to retry a failed drain activity. If yes reschedule drain

func (h *DrainingResourceEventHandler) offendingConditions(n *core.Node) []SuppliedCondition {
	_ = "STUB: not implemented"
	return nil
}

func shouldUncordon(n *core.Node) bool { _ = "STUB: not implemented"; return false }

func parseConditionsFromAnnotation(n *core.Node) []SuppliedCondition {
	_ = "STUB: not implemented"
	return nil
}

func (h *DrainingResourceEventHandler) uncordon(n *core.Node) { _ = "STUB: not implemented"; return }

// nolint:gosec

// nolint:gosec

// nolint:gosec

func removeAnnotationMutator(n *core.Node) { _ = "STUB: not implemented"; return }

func (h *DrainingResourceEventHandler) cordon(n *core.Node, badConditions []SuppliedCondition) {
	_ = "STUB: not implemented"
	return
}

// nolint:gosec
// Events must be associated with this object reference, rather than the
// node itself, in order to appear under `kubectl describe node` due to the
// way that command is implemented.
// https://github.com/kubernetes/kubernetes/blob/17740a2/pkg/printers/internalversion/describe.go#L2711

// nolint:gosec

// nolint:gosec

func conditionAnnotationMutator(conditions []SuppliedCondition) func(*core.Node) {
	_ = "STUB: not implemented"
	return nil
}

// drain schedule the draining activity
func (h *DrainingResourceEventHandler) scheduleDrain(n *core.Node) {
	_ = "STUB: not implemented"
	return
}

// nolint:gosec

// nolint:gosec

// nolint:gosec

func HasDrainRetryAnnotation(n *core.Node) bool { _ = "STUB: not implemented"; return false }
