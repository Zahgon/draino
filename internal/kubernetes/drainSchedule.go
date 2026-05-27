package kubernetes

import (
	"sync"
	"time"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
)

const (
	SetConditionTimeout     = 10 * time.Second
	SetConditionRetryPeriod = 50 * time.Millisecond
)

type DrainScheduler interface {
	HasSchedule(name string) (has, failed bool)
	Schedule(node *v1.Node) (time.Time, error)
	DeleteSchedule(name string)
}

type DrainSchedules struct {
	sync.Mutex
	schedules map[string]*schedule

	lastDrainScheduledFor time.Time
	period                time.Duration

	logger        *zap.Logger
	drainer       Drainer
	eventRecorder record.EventRecorder
}

func NewDrainSchedules(drainer Drainer, eventRecorder record.EventRecorder, period time.Duration, logger *zap.Logger) DrainScheduler {
	_ = "STUB: not implemented"
	return *new(DrainScheduler)
}

func (d *DrainSchedules) HasSchedule(name string) (has, failed bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (d *DrainSchedules) DeleteSchedule(name string) { _ = "STUB: not implemented"; return }

func (d *DrainSchedules) WhenNextSchedule() time.Time {
	_ = "STUB: not implemented"
	// compute drain schedule time
	return *new(time.Time)
}

func (d *DrainSchedules) Schedule(node *v1.Node) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// we already have a schedule planned

// compute drain schedule time

// Mark the node with the condition stating that drain is scheduled

// if we cannot mark the node, let's remove the schedule

type schedule struct {
	when   time.Time
	failed int32
	finish time.Time
	timer  *time.Timer
}

func (s *schedule) setFailed() { _ = "STUB: not implemented"; return }

func (s *schedule) isFailed() bool { _ = "STUB: not implemented"; return false }

func (d *DrainSchedules) newSchedule(node *v1.Node, when time.Time) *schedule {
	_ = "STUB: not implemented"
	return nil
}

// nolint:gosec

// nolint:gosec

// nolint:gosec

type AlreadyScheduledError struct {
	error
}

func NewAlreadyScheduledError() error { _ = "STUB: not implemented"; return nil }

func IsAlreadyScheduledError(err error) bool { _ = "STUB: not implemented"; return false }
