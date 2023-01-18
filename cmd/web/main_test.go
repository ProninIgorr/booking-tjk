package main

import "testing"

func TestRun(t *testing.T) {
	db, err := run()
	if err != nil {
		t.Error("failed run()")
	}
	if db == nil {
		t.Error("failed run()")
	}
}
