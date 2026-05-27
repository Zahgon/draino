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
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/types"
)

// NewNodeLabelFilter returns a filter that returns true if the supplied node satisfies the boolean expression
func NewNodeLabelFilter(expressionStr *string, log *zap.Logger) (func(o interface{}) bool, error) {
	_ = "STUB: not implemented"
	//This feels wrong but this is how the previous behavior worked so I'm only keeping it to maintain compatibility.
	return nil, nil
}

//This feels wrong but this is how the previous behavior worked so I'm only keeping it to maintain compatibility.

// ParseConditions can parse the string array of conditions to a list of
// SuppliedContion to support particular status value and duration.
func ParseConditions(conditions []string) []SuppliedCondition {
	_ = "STUB: not implemented"
	return nil
}

// Keep backward compatibility

// NodeProcessed tracks whether nodes have been processed before using a map.
type NodeProcessed map[types.UID]bool

// NewNodeProcessed returns a new node processed filter.
func NewNodeProcessed() NodeProcessed { _ = "STUB: not implemented"; return *new(NodeProcessed) }

// Filter returns true if the supplied object is a node that this filter has
// not seen before. It is not threadsafe and should always be the last filter
// applied.
func (processed NodeProcessed) Filter(o interface{}) bool { _ = "STUB: not implemented"; return false }

// ConvertLabelsToFilterExpr Convert old list labels into new expression syntax
func ConvertLabelsToFilterExpr(labelsSlice []string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//sort the maps so that the unit tests actually work
