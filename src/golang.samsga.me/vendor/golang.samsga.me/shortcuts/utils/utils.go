package utils

import (
  "fmt"
)

func PanicOnError(msg string, err error) {
  if err != nil {
    panic(fmt.Sprintf("%s: %v", msg, err))
  }
}
