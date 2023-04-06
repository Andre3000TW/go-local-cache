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
		testName      string
		key           string
		val           any
		sleepDuration time.Duration
		expected      any
		ok            bool
	}{
		{
			testName: "Should get right value with type: int",
			key:      "key",
			val:      123456,
			expected: 123456,
		},
		{
			testName: "Should get right value with type: float",
			key:      "key",
			val:      123.456,
			expected: 123.456,
		},
		{
			testName: "Should get right value with type: string",
			key:      "key",
			val:      "123456",
			expected: "123456",
		},
	}

	for _, tc := range testCases {
		cs.cache.val[tc.key] = tc.val

		actual, _ := cs.cache.Get(tc.key)

		cs.Require().Equal(tc.expected, actual)
	}
}

func (cs *cacheSuite) TestSet() {
	testCases := []struct {
		testName      string
		sleepDuration time.Duration
		key           string
		val           any
		expected      any
	}{
		{
			testName:      "Should get the same value as set",
			sleepDuration: 0,
			key:           "key",
			val:           "value",
			expected:      "value",
		},
		{
			testName:      "Should get another value when overwrite with the same key",
			sleepDuration: 0,
			key:           "key",
			val:           "value",
			expected:      "value",
		},
		{
			testName:      "Should not get the value when expired",
			sleepDuration: 31,
			key:           "key",
			val:           "value",
			expected:      nil,
		},
	}

	for _, tc := range testCases {
		cs.cache.Set(tc.key, tc.val)

		time.Sleep(tc.sleepDuration * time.Second)

		actual := cs.cache.val[tc.key]

		cs.Require().Equal(tc.expected, actual)
	}
}

func TestCacheSuite(t *testing.T) {
	suite.Run(t, new(cacheSuite))
}
