
package ternary

import (
  "strings"
  // "fmt"
)

type rods [64]rune

type point struct {
  x rods
  y rods
}

func intoTwo(x rods) (a, b rods) {
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
  arr := rods{}
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

func interAdd(a, b rods) (c, d rods) {
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

func isZero(a rods) (right bool) {
  right = true
  for _, v := range(a) {
    if v != '5' {
      right = false
    }
  }
  return
}

func addRods(a, b rods) rods {

  c, d := a, b
  // fmt.Println(nice(string(c[:])))
  // fmt.Println(nice(string(d[:])))
  counter := 0
  for {
    c, d = interAdd(c, d)
    // fmt.Println(nice(string(c[:])))
    // fmt.Println(nice(string(d[:])))
    // fmt.Println("end")
    if isZero(d) {
      break
    }
    counter ++
    if counter > 20 {
      break
    }

  }
  return c
}

func nice(a string) (b string) {
  b = strings.Replace(a, "5", "-", -1)
  return
}

func addPoint(a, b point) (c point) {
  c.x = addRods(a.x, b.x)
  c.y = addRods(a.y, b.y)
  return
}

func intoOne(a point) (b string) {
  var c rods
  for i, _ := range(c) {
    x1 := a.x[i]
    x2 := a.y[i]
    switch {
      case x1 == '1' && x2 == '1':
        c[i] = '8'
      case x1 == '1' && x2 == '5':
        c[i] = '1'
      case x1 == '1' && x2 == '9':
        c[i] = '6'
      case x1 == '5' && x2 == '1':
        c[i] = '7'
      case x1 == '5' && x2 == '5':
        c[i] = '5'
      case x1 == '5' && x2 == '9':
        c[i] = '3'
      case x1 == '9' && x2 == '1':
        c[i] = '4'
      case x1 == '9' && x2 == '5':
        c[i] = '9'
      case x1 == '9' && x2 == '9':
        c[i] = '2'
    }
  }
  b = string(c[:])
  return
}

func Add(a string, b string) (c string) {
  pa := readArray(a)
  pb := readArray(b)
  pc := addPoint(pa, pb)
  res := intoOne(pc)
  intPart := strings.TrimLeft(res[:32], "5")
  deciPart := strings.TrimRight(res[32:], "5")
  x := []string{intPart, deciPart}
  c = strings.Join(x, ".")
  return
}