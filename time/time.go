package main

import (
	"fmt"
	"time"
)

func main() {
  p := fmt.Println

  now := time.Now()

  p(now)

  then := time.Date(2009, 11, 17, 20, 34, 58, 0, time.Local)

  p(then)

  p(then.Year())
  p(then.Month())
  p(then.Day())
  p(then.Hour())
  p(then.Minute())
  p(then.Second())
  p(then.Nanosecond())
  p(then.Location())

  diff := now.Sub(then)

  p(diff)
  p(diff.Hours() / 24 / 365)
  p(diff.Hours())
  p(diff.Minutes())
  p(diff.Seconds())
  p(diff.Nanoseconds())

  p(then.Add(diff))
  p(then.Add(-diff))
}

