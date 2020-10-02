package main

import (
	"fmt"
)

// TODO: Реализовать вычисление Квадратного корня
func Sqrt(x float64) float64 {
	if x <= 0 {
    return 0
  }

  var prev float64 = 0
  var curr float64 = x / 2
  for curr - prev != 0 {
    prev = curr
    curr = 0.5 * (prev + x / prev)
  }

  return curr
}

func main() {
	fmt.Println(Sqrt(2))
}
