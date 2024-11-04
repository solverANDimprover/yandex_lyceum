package main

import (
	"fmt"
	"testing"
)

func TestParseValue(t *testing.T) {
	got := "68 / 4"
	expect := "17"
	if ParseValue(got) != expect {
		fmt.Errorf("expected %s got %s", expect, got)
	}

}
