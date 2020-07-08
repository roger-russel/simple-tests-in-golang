package main

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_getSlice(t *testing.T) {

	t.Run("simple", func(t *testing.T) {

		want := []string{
			"two",
			"three",
			"one",
		}

		sort.Strings(want)

		if got := getSlice(); !cmp.Equal(got, want) {
			t.Errorf("getObj() => %v", cmp.Diff(got, want))
		}

	})

}
