package main

import (
	"fmt"
	"time"
)

func main() {
  p := fmt.Println
  now := time.Now()

  p(now)
  p(now.Unix())
  p(now.UnixMilli())
  p(now.UnixMicro())
  p(now.UnixNano())

  p(
    time.Unix(
      now.Unix(),
      0,
    ),
  )

  p(
    time.Unix(
      0, // unix timestamp in seconds
      now.UnixNano(), // unix timestamp in nano seconds
    ),
  )

  p(time.UnixMicro(now.UnixMicro()))
  p(time.UnixMilli(now.UnixMilli()))
}

