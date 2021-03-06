package handlers

import "testing"

var dirNameTests = []struct {
	url  string
	want string
}{
	{"https://github.com/foo/bar", "repos/src/github.com/foo/bar"},
}

func TestDirName(t *testing.T) {
	for _, tt := range dirNameTests {
		if got := dirName(tt.url); got != tt.want {
			t.Errorf("dirName(%q) = %q, want %q", tt.url, got, tt.want)
		}
	}
}
