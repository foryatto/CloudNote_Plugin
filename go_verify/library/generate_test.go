package library

import (
	"testing"
)

func TestNewRandomNums(t *testing.T) {
	result := NewRandomNums(8)
	t.Log(result)
}
