package attribute

import (
	"go.opentelemetry.io/otel/attribute/internal/xxhash"
)

const (
	boolID         uint64 = 7953749933313450591
	int64ID        uint64 = 7592915492740740150
	float64ID      uint64 = 7376742710626956342
	stringID       uint64 = 6874584755375207263
	boolSliceID    uint64 = 6875993255270243167
	int64SliceID   uint64 = 3762322556277578591
	float64SliceID uint64 = 7308324551835016539
	stringSliceID  uint64 = 7453010373645655387
	byteSliceID    uint64 = 6874028470941080415
	sliceID        uint64 = 7883494272577650031
	mapID          uint64 = 6872316492666199903
	emptyID        uint64 = 7305809155345288421
)

func hashKVs(kvs []KeyValue) uint64 { _ = "STUB: not implemented"; return 0 }

func hashKV(h xxhash.Hash, kv KeyValue) xxhash.Hash {
	_ = "STUB: not implemented"
	return *new(xxhash.Hash)
}

func hashValue(h xxhash.Hash, v Value) xxhash.Hash {
	_ = "STUB: not implemented"
	return *new(xxhash.Hash)
}
