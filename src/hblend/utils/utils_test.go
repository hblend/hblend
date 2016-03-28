package utils

import "testing"

func Test_Md5String(t *testing.T) {
	result := Md5String("hello world")
	expected := "5eb63bbbe01eeed093cb22bb8f5acdc3"

	if result != expected {
		t.Error("Result is not the expected")
	}
}
