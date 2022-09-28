package main

import (
  "fmt"
  "golang.org/x/crypto/md4"
)

func main() {
    h := md4.New()
    fmt.Printf("Random output: %x\n", h.Sum(nil))
}
