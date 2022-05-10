package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	scannedString := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	count := wc(scannedString)

	if exp != count {
		t.Errorf("Expected %d, got %d instead !!!\n", exp, count)
	}
}
