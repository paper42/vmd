package markdown

import (
	"testing"
	"bytes"
)

func TestFlatten(t *testing.T) {
	a := flattenSpaces([]byte(" testing leading spaces"))
	b := []byte("testing leading spaces")
	if bytes.Compare(a,b) != 0 {
		t.Errorf("leading spaces failed")
	}

	a = flattenSpaces([]byte("testing trailing spaces  "))
	b = []byte("testing trailing spaces")
	if bytes.Compare(a,b) != 0 {
		t.Errorf("traling spaces failed")
	}

	a = flattenSpaces([]byte("testing   interior spaces"))
	b = []byte("testing interior spaces")
	if bytes.Compare(a,b) != 0 {
		t.Errorf("leading spaces failed")
	}
}
