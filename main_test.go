package main

import (
	"testing"
	"math/rand"
	"time"
	"fmt"

	"github.com/stretchr/testify/assert"
)

const N = 10

func BenchmarkSliceSet(b *testing.B) {
	now := time.Now()
	rand.Seed(now.Unix())
	intValues := rand.Perm(N)
	indexOrder := rand.Perm(N)

	for n := 0; n < b.N; n++ {
		ms := NewSliceSet()
		for _, val := range intValues {
			ms.Insert(fmt.Sprint(val))
		}
		for _, index := range indexOrder {
			assert.True(b, ms.Contains(fmt.Sprint(intValues[index])))
		}
	}
}

func BenchmarkSliceSetExp(b *testing.B) {
	now := time.Now()
	rand.Seed(now.Unix())
	intValues := rand.Perm(N)
	indexOrder := rand.Perm(N)

	for n := 0; n < b.N; n++ {
		ms := NewSliceSetExp()
		for _, val := range intValues {
			ms.Insert(fmt.Sprint(val))
		}
		for _, index := range indexOrder {
			assert.True(b, ms.Contains(fmt.Sprint(intValues[index])))
		}
	}
}

func BenchmarkMapSet(b *testing.B) {
	now := time.Now()
	rand.Seed(now.Unix())
	intValues := rand.Perm(N)
	indexOrder := rand.Perm(N)

	for n := 0; n < b.N; n++ {
		ms := NewMapSet()
		for _, val := range intValues {
			ms.Insert(fmt.Sprint(val))
		}
		for _, index := range indexOrder {
			assert.True(b, ms.Contains(fmt.Sprint(intValues[index])))
		}
	}
}
