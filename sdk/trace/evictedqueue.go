package trace

import (
	"sync"
)

type evictedQueue[T any] struct {
	queue          []T
	capacity       int
	droppedCount   int
	logDroppedMsg  string
	logDroppedOnce sync.Once
}

func newEvictedQueueEvent(capacity int) evictedQueue[Event] { _ = "STUB: not implemented"; return nil }

func newEvictedQueueLink(capacity int) evictedQueue[Link] { _ = "STUB: not implemented"; return nil }

func (eq *evictedQueue[T]) add(value T) { _ = "STUB: not implemented"; return }

func (eq *evictedQueue[T]) logDropped() { _ = "STUB: not implemented"; return }

func (eq *evictedQueue[T]) copy() []T { _ = "STUB: not implemented"; return nil }
