package godane

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	maxCache := 10
	c := newCache(maxCache)

	for i := 0; i < 30; i++ {
		c.set(fmt.Sprintf("test%d", i), &entry{})
	}

	_, ok := c.get("test29")

	if !ok {
		t.Fatal("want key `test29`")
	}

	if c.len() > maxCache {
		t.Fatalf("want cache len = %d, got %d", maxCache, c.len())
	}
}
