package kv_test

import (
	"testing"

	"github.com/ezebunandu/kv"
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

func TestSave__SavesDataPersistently(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/kvtest.store"
	s, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	s.Set("A", "1")
	s.Set("B", "2")
	s.Set("C", "3")
	err = s.Save()
	if err != nil {
		t.Fatal(err)
	}
	s2, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := s2.Get("A"); v != "1" {
		t.Fatalf("want A=1, got A=%s", v)
	}
	if v, _ := s2.Get("B"); v != "2" {
		t.Fatalf("want B=2, got B=%s", v)
	}
	if v, _ := s2.Get("C"); v != "3" {
		t.Fatalf("want C=3, got C=%s", v)
	}
}
