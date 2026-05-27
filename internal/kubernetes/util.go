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

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
)

// Component is the name of this application.
const Component = "draino"

// BuildConfigFromFlags is clientcmd.BuildConfigFromFlags with no annoying
// dependencies on glog.
// https://godoc.org/k8s.io/client-go/tools/clientcmd#BuildConfigFromFlags
func BuildConfigFromFlags(apiserver, kubecfg string) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewEventRecorder returns a new record.EventRecorder for the given client.
func NewEventRecorder(c kubernetes.Interface) record.EventRecorder {
	_ = "STUB: not implemented"
	return *new(record.EventRecorder)
}

func RetryWithTimeout(f func() error, retryPeriod, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
