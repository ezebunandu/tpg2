package kv_test

import (
	"github.com/ezebunandu/kv"
	"testing"
)

func TestGet__ReturnsNotOKIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	_, ok := s.Get("key")
	if ok {
		t.Fatal("unexpected ok")
	}
}

func TestGet__ReturnsValueAndOKIfKeyExists(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "value")
	if err != nil {
		t.Fatal(err)
	}
	want := "value"
	v, ok := s.Get("key")
	if !ok {
		t.Fatal("not ok")
	}
	if want != v {
		t.Errorf("want: %s, got: %s", want, v)
	}
}

func TestSet__UpdatesExistingValue(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "value")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "updated")
	want := "updated"
	v, ok := s.Get("key")
	if !ok {
		t.Fatal("not ok")
	}
	if want != v {
		t.Errorf("want: %s, got: %s", want, v)
	}
}
