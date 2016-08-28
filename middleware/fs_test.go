package middleware

import (
	"os"
	"strings"
	"testing"
)

func TestFsPath(t *testing.T) {
	if actual := FPath(); !strings.HasSuffix(actual, ".coredns") {
		t.Errorf("Expected path to be a .coredns folder, got: %v", actual)
	}

	os.Setenv("COREDNSPATH", "testpath")
	if actual, expected := AssetsPath(), "testpath"; actual != expected {
		t.Errorf("Expected path to be %v, got: %v", expected, actual)
	}
	os.Setenv("COREDNSPATH", "")
}
