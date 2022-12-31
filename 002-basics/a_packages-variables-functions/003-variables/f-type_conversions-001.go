package main

import "fmt"

func main() {
  i := 42
  f := float64(i)
  u := uint(f)

  fmt.Printf("Type: %T Value: %v\n", i, i)
  fmt.Printf("Type: %T Value: %v\n", f, f)
  fmt.Printf("Type: %T Value: %v\n", u, u)
}
