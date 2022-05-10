package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	scannedString := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	count := wc(scannedString, false, false)

	if exp != count {
		t.Errorf("Expected %d, got %d instead !!!\n", exp, count)
	}
}

func TestCountLines(t *testing.T) {
	scannedString := bytes.NewBufferString("word1 word2 word3 word4\n nextline nextline1 \n nextline1 nextline2 \n")
	exp := 3
	count := wc(scannedString, true, false)

	if exp != count {
		t.Errorf("Expected %d, got %d instead !!!\n", exp, count)
	}
}

func TestCountBytes(t *testing.T) {
	scannedString := bytes.NewBufferString("word1\n")
	exp := 6
	count := wc(scannedString, false, true)

	if exp != count {
		t.Errorf("Expected %d, got %d instead !!!\n", exp, count)
	}
}
