package traffic

import (
	"fmt"

	"github.com/caicloud/nirvana/log"
	v1 "github.com/dyweb/cloudab/pkg/apis/v1"
	"github.com/spaolacci/murmur3"
)

const (
	defaultBuckets = 1000
	defaultSeed    = 10000
)

var (
	defaultHashFunc = func(userID string) uint64 {
		h := murmur3.New64WithSeed(defaultSeed)
		h.Write([]byte(userID))
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
	logger.V(log.LevelDebug).Infof(
		"UserID %s, hash is %d", userID, hashNum)
	bucketNum := hashNum % r.Bucket

	expected := -1
	for i, v := range versions {
		currentBuckets := uint64(v.Traffic) * (r.Bucket / 100)
		logger.V(log.LevelDebug).Infof(
			"Expected bucket: %d, current bucket: %d", bucketNum, currentBuckets)
		if bucketNum > currentBuckets {
			bucketNum -= currentBuckets
		}
		expected = i
		logger.V(log.LevelDebug).Infof("Find the expected version %d", expected)
		break
	}

	if expected == -1 {
		return -1, fmt.Errorf(
			"Failed to get the version in router with bucket %d and total buckets %d",
			bucketNum, defaultBuckets)
	}
	return expected, nil
}
