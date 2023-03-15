package search 

import (
  "math"
)

func Linear(haystack []int, needle int) bool {
  for _, p := range haystack {
    if p == needle {
      return true
    }
  }
  return false
}

func Binary(haystack []int, needle int) bool {
  lo := 0
  hi := len(haystack)
  for lo < hi {
    m := lo + (hi-lo) / 2
    v := haystack[m]
    if v == needle {
      return true
    } else if v > needle {
      hi = m
    } else {
      lo = m + 1
    }
  }
  return false
}

func Two(breaks []bool) int {
  jmp := int(math.Floor(math.Sqrt(float64(len(breaks)))))
  
  i := jmp
  for i < len(breaks) {
    if (breaks[jmp]) {
      break;
    }
    i += jmp
  }
  i -= jmp

  for j := 0; j < jmp && i < len(breaks); {
    if breaks[i] {
      return i
    }
    i++
    j++
  }
  return -1
}
