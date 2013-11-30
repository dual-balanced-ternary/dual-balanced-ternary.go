
package ternary

import (
  "testing"
)

func TestAdd(t *testing.T) {
  a := "23.2"
  b := "34.6"
  if "342.7" == Add(a, b) {
    t.Log("right")
  } else {
    t.Error("wrong")
  }
}