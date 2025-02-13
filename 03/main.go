package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

// argument type declaration
func add2(x, y int) int {
  return x + y
}

// return a tuple
func swap(x, y string) (string, string) {
  return y, x
}

// named results
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(add(42, 13))
  fmt.Println(add2(42, 13))
  c, a := swap("aaa", "ccc")
  fmt.Println(a, c)
  x, y := split(9)
  fmt.Println(x, y)
}
