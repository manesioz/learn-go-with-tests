package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("calling the sum function", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4, 5}
		arr2 := []int{0, 0, 0, 0}
		arr3 := []int{2, 4, 6}

		got := Sum(arr1, arr2, arr3)
		want := []int{15, 0, 12}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error: got %q, but want %q", got, want)
		}
	})
}
