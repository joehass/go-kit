package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrcmp(t *testing.T) {
	t.Run("normal true", func(t *testing.T) {
		s1 := []int32{1, 2, 3, 4, 5}
		s2 := []int32{2, 3, 6, 7}

		add, dele := Arrcmp(s1, s2)
		assert.Equal(t, true, sliceBool(add, []int32{6, 7}))
		assert.Equal(t, true, sliceBool(dele, []int32{1, 4, 5}))
	})

	t.Run("add empty true", func(t *testing.T) {
		s1 := []int32{1, 2, 3, 4, 5}
		s2 := []int32{}

		add, dele := Arrcmp(s1, s2)
		assert.Equal(t, true, sliceBool(add, nil))
		assert.Equal(t, true, sliceBool(dele, []int32{1,2,3, 4, 5}))
	})

	t.Run("dele empty true", func(t *testing.T) {
		s1 := []int32{}
		s2 := []int32{1, 2, 3, 4, 5}

		add, dele := Arrcmp(s1, s2)
		assert.Equal(t, true, sliceBool(add, []int32{1,2,3,4,5}))
		assert.Equal(t, true, sliceBool(dele,nil ))
	})
}

func TestSliceBool(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		s1 := []int32{1, 2, 3, 4}
		s2 := []int32{1, 3, 4, 2}

		assert.Equal(t, true, sliceBool(s1, s2))
	})
	t.Run("false", func(t *testing.T) {
		s1 := []int32{1, 2, 3, 4}
		s2 := []int32{1, 3, 2}

		assert.Equal(t, true, sliceBool(s1, s2))
	})
}

// sliceBool 比较数组值是否相同
func sliceBool(s1, s2 []int32) bool {
	result := make(map[int32]struct{})
	for _, s := range s1 {
		result[s] = struct{}{}
	}

	for _, s := range s2 {
		if _, ok := result[s]; !ok {
			return false
		}
	}

	return true
}
