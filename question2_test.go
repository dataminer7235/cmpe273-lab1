package main

import "testing"
type rect struct{length int
                breadth int}
type testpair struct {value float64
                      rectangle rect}

var c = &testpair{4, rect{1, 1}}
var tests = []testpair{{{ 1, 1},4},{{ 3, 2},10},{{4, 3},14},{{5,5},10},}

func TestFib(t *testing.T) {
for _, pair := range tests{
  v := pair.rectangle.Perimeter
  if v != tests.value {
      t.Error("For", pair.rectangle, "expected", pair.value,"got", v,)}
  }
}
