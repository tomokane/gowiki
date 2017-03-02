package main

import (
	"testing"
)

func TestLinking(t *testing.T) {
	body := []byte("[test]")
	linked_body := linking(body)
	if string(linked_body) != "<a href=\"test\">test</a>" {
		t.Fatal("can't insert inner-link. produced", string(linked_body))
	}
}
