package traffic

import (
	"fmt"
	"hash/maphash"
	"unsafe"

	"github.com/caicloud/nirvana/log"
	v1 "github.com/dyweb/cloudab/pkg/apis/v1"
)

const (
	defaultBuckets = 1000
)

var (
	defaultSeed = &struct {
		s uint64
	}{
		s: 10000,
	}
	ds              = (*maphash.Seed)(unsafe.Pointer(defaultSeed))
	defaultHashFunc = func(userID string) uint64 {
		var h maphash.Hash
		h.Reset()
		h.SetSeed(*ds)
		h.WriteString(userID)
		return h.Sum64()
	}
)

type HashFunc func(userID string) uint64

type Router struct {
	Bucket   uint64
	HashFunc HashFunc
}

func NewRouter() *Router {
	return &Router{
		Bucket:   100,
		HashFunc: defaultHashFunc,
	}
}

// Route routes userID to one version.
func (r Router) Route(userID string, versions []v1.Version) (int, error) {
	logger := log.DefaultLogger()

	hashNum := r.HashFunc(userID)
	bucketNum := hashNum % r.Bucket

	expected := -1
	for i, v := range versions {
		currentBuckets := uint64(v.Traffic) * (r.Bucket / 100)
		logger.V(log.LevelDebug).Infof(
			"Expected bucket: %d, hash: %d, current bucket: %d", bucketNum, hashNum, currentBuckets)
		if bucketNum > currentBuckets {
			bucketNum -= currentBuckets
		}
		logger.V(log.LevelDebug).Infof("Find the expected version %d", expected)
		expected = i
		break
	}

	if expected == -1 {
		return -1, fmt.Errorf(
			"Failed to get the version in router with bucket %d and total buckets %d",
			bucketNum, defaultBuckets)
	}
	return expected, nil
}
