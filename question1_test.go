package main

import "testing"

type testpair struct {values int
                      fibvalue int}
var tests = []testpair{{ 1, 1 },{ 3, 2 },{4, 3 },{5,5},}

func TestFib(t *testing.T) {
  for _, pair := range tests{
  v := Fib(pair.values)
  if v != pair.fibvalue {
      t.Error("For", pair.values, "expected", pair.fibvalue,"got", v,)}
  }
}
