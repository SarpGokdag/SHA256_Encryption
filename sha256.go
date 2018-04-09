package main

import (
  "fmt"
  "crypto/sha256"
)

func main()  {
  str := " sarp gokdag "
  encryption := sha256.Sum256([]byte(str))
  fmt.Printf("%X", encryption)
  fmt.Printf("\n")
}
