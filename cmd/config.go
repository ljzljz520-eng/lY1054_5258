package main

import (
	"os"
	"strconv"
)

type Config struct {
	Addr, DB string
	Rate     int64
}

func loadConfig() Config {
	c := Config{Addr: ":8080", DB: "library.db", Rate: 25}
	if v := os.Getenv("LIBRARY_ADDR"); v != "" {
		c.Addr = v
	}
	if v := os.Getenv("LIBRARY_RATE"); v != "" {
		if n, e := strconv.ParseInt(v, 10, 64); e == nil && n >= 0 {
			c.Rate = n
		}
	}
	return c
}
