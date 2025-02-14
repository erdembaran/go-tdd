package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	slice1 := []int{2,4,6,8,10}
	slice2 := []int{2,4,6,8,10,12}

	got := SumAll(slice1, slice2)
	want := []int{30, 42}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}