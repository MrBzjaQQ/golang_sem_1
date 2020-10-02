package main

import (
  "fmt"
  "bytes"
  "strings"
)

type memoizeFunction func(int, ...int) interface{}

// TODO реализовать
var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

//TODO Write memoization function

func memoize(function memoizeFunction) memoizeFunction {
  mem := make(map[string]interface{})
  return func(n int, x ...int) interface{} {
    key := fmt.Sprint(n);
    for _, i := range x {
      key += fmt.Sprintf(",%d", i);
    }
    if value, found := mem[key]; found {
      return value
    }
    value := function(n, x...)
    mem[key] = value
    return value
  }
}

// TODO обернуть функции fibonacci и roman в memoize
func init() {
  fibonacci = memoize(func(x int, _ ...int) interface{} {
    if x < 2 {
      return 1
    }
    return fibonacci(x - 1).(int) + fibonacci(x - 2).(int)
  })

  decimals := []int{ 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1 }
  romans := []string{ "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I" }

  romanForDecimal = memoize(func(x int, xs ...int) interface{} {
    if x < 0 || x > 3999 {
      panic("Wrong value. x should be between 0 and 3999")
    }

    var buffer bytes.Buffer
    for i, decimal := range decimals {
      rem := x / decimal
      x %= decimal
      if rem > 0 {
        buffer.WriteString(strings.Repeat(romans[i], rem));
      }
    }
    return buffer.String();
  })
}

func main() {
	fmt.Println("Fibonacci(45) =", fibonacci(45).(int))
	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	}
}
