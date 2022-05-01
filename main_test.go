package main

import (
	"reflect"
	"testing"
	"math/rand"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	now := time.Now()
	rand.Seed(now.Unix())
	intValues := rand.Perm(100)

	ms := NewMapSet()
	for _, val := range intValues {

	}

	ms.Insert(2)
	ms.Insert("Meow")
	assert.True(t, ms.Contains(2))
	assert.True(t, ms.Contains("Meow"))
	assert.False(t, ms.Contains(3))
	assert.False(t, ms.Contains("adsf"))

	rs := ReflectSliceSet{
		ds: []any{},
	}

	assert.True(t, reflect.DeepEqual(2, 2))

	rs.Insert(2)
	rs.Insert("Meow")
	assert.True(t, rs.Contains(2))
	assert.True(t, rs.Contains("Meow"))
	assert.False(t, rs.Contains(3))
	assert.False(t, rs.Contains("adsf"))
}
