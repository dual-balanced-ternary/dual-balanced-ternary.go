
package ternary

import (
  "strings"
  "fmt"
)

type point struct {
  x [64]rune
  y [64]rune
}

func intoTwo(x [64]rune) (a, b [64]rune) {
  dict := map[rune] ([2]rune) {
    '1': [2]rune{'1', '5'},
    '2': [2]rune{'9', '9'},
    '3': [2]rune{'5', '9'},
    '4': [2]rune{'9', '1'},
    '5': [2]rune{'5', '5'},
    '6': [2]rune{'1', '9'},
    '7': [2]rune{'5', '1'},
    '8': [2]rune{'1', '1'},
    '9': [2]rune{'9', '5'},
  }
  for index, char := range(x) {
    a[index] = dict[char][0]
    b[index] = dict[char][1]
  }
  return
}

func readArray(x string) point {
  x = strings.Trim(x, " ")
  xs := strings.Split(x, ".")
  arr := [64]rune{}
  for index := range(arr) {
    arr[index] = '5'
  }

  switch {
    case len(xs) == 1:
      for index, char := range(x) {
        cursor := 32 - len(x) + index
        arr[cursor] = char
      }
    case len(xs) == 2:
      for index, char := range(xs[0]) {
        cursor := 32 - len(xs[0]) + index
        arr[cursor] = char
      }
      for index, char := range(xs[1]) {
        cursor := 32 + index
        arr[cursor] = char
      }
  }

  a, b := intoTwo(arr)
  arr2 := point{a, b}
  return arr2
}

func interAdd(a, b [64]rune) (c, d [64]rune) {
  for i := 0; i < 64; i++ {
    x1 := a[i]
    x2 := b[i]
    var ta rune
    var tb rune
    switch {
    case x1 == '5':
      ta = x2
      tb = '5'
    case x2 == '5':
      ta = x1
      tb = '5'
    case x1 == '1' && x2 == '9':
      ta = '5'
      tb = '5'
    case x1 == '9' && x2 == '1':
      ta = '5'
      tb = '5'
    case x1 == '1' && x2 == '1':
      ta = '9'
      tb = '1'
    case x1 == '9' && x2 == '9':
      ta = '1'
      tb = '9'
    }
    c[i] = ta
    if i > 0 {
      d[i - 1] = tb
    } else {
      d[64 - 1] = '5'
    }
  }
  return
}

func addPoint(a, b point) (c point) {
  var next rune
  var curr rune
  return
}

func Add(a string, b string) (c string) {
  pa := readArray(a)
  pb := readArray(b)
  addPoint(pa, pb)
  return
}