package main

import "testing"

func TestConfig(t *testing.T) {
	if loadConfig().Addr == "" {
		t.Fail()
	}
}
