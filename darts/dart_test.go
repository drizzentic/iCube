package main

import (
	"testing"
)

func TestScore(t *testing.T) {

testInnerCircle:=score(0,-10)
if testInnerCircle != 1{
	t.Errorf("Failed expected %v but got %v", 1, testInnerCircle)
}
}
