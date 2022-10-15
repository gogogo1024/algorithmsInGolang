package dynamic

import (
	"testing"
)

func TestLongestCommonSubstring(t *testing.T) {
	var data = []struct {
		name, a, b, want string
	}{
		{"hish-vista", "vista", "hish", "is"},
		{"hish-fish", "fish", "hish", "ish"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			actual := substring(d.a, d.b)
			if actual != d.want {
				t.Errorf("want %s but got %s", d.want, actual)
			}
		})
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	var data = []struct {
		name, a, b string
		want       int
	}{
		{"fosh-fish", "fish", "fosh", 3},
		{"fosh-fort", "fort", "fosh", 2},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			actual := subsequence(d.a, d.b)
			if actual != d.want {
				t.Errorf("want %d but got %d", d.want, actual)
			}
		})
	}
}
