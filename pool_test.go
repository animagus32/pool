package pool

import (
	"context"
	"testing"
)

func TestPool(t *testing.T) {
	p := GetPool(context.Background(), 2)
	p.Run()
}
