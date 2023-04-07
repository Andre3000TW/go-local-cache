package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type cacheSuite struct {
	suite.Suite
	cache *cache
}

func (cs *cacheSuite) SetupTest() {
	cs.cache = New().(*cache)
}

func (cs *cacheSuite) TestGet() {
	testCases := []struct {
		desc     string
		key      string
		val      interface{}
		expected interface{}
	}{
		{
			desc:     "Should get right value with type: int",
			key:      "key",
			val:      123456,
			expected: 123456,
		},
		{
			desc:     "Should get right value with type: float",
			key:      "key",
			val:      123.456,
			expected: 123.456,
		},
		{
			desc:     "Should get right value with type: string",
			key:      "key",
			val:      "123456",
			expected: "123456",
		},
	}

	for _, tc := range testCases {
		cs.cache.items[tc.key] = &cacheItem{
			val: tc.val,
		}

		actual := cs.cache.Get(tc.key)

		cs.Require().Equal(tc.expected, actual, tc.desc)
	}
}

func (cs *cacheSuite) TestSet() {
	testCases := []struct {
		desc     string
		key      string
		val      interface{}
		expected interface{}
	}{
		{
			desc:     "Should get the same value as set",
			key:      "key",
			val:      "value",
			expected: "value",
		},
		{
			desc:     "Should get another value when overwrite with the same key",
			key:      "key",
			val:      "value2",
			expected: "value2",
		},
	}

	for _, tc := range testCases {
		cs.cache.Set(tc.key, tc.val)

		item := cs.cache.items[tc.key]

		cs.Require().Equal(tc.expected, item.val, tc.desc)
	}
}

func (cs *cacheSuite) TestSetOnTimerExpired() {
	testCases := []struct {
		desc     string
		key      string
		val      interface{}
		duration time.Duration
	}{
		{
			desc:     "Should get nil when item expired",
			key:      "key",
			val:      "value",
			duration: 31,
		},
	}

	for _, tc := range testCases {
		cs.cache.Set(tc.key, tc.val)

		time.Sleep(tc.duration * time.Second)

		_, ok := cs.cache.items[tc.key]

		cs.Require().False(ok, tc.desc)
	}
}

func (cs *cacheSuite) TestSetOnTimerReset() {
	testCases := []struct {
		desc string
		key  string
		val  interface{}
	}{
		{
			desc: "Should reset timer when overwrite with the same key",
			key:  "key",
			val:  "value",
		},
	}

	for _, tc := range testCases {
		cs.cache.Set(tc.key, tc.val)
		oldTimer := &cs.cache.items[tc.key].timer

		cs.cache.Set(tc.key, tc.val)
		newTimer := &cs.cache.items[tc.key].timer

		cs.Require().NotEqual(oldTimer, newTimer, tc.desc)
	}
}

func TestCacheSuite(t *testing.T) {
	suite.Run(t, new(cacheSuite))
}
