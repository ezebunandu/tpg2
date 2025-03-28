package kv_test

import (
	"os"
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
		t.Errorf("want: %s, got: %q", want, v)
	}
}

func TestSet__UpdatesExistingKeyToNewValue(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "value")
	s.Set("key", "updated")
	v, ok := s.Get("key")
	if !ok {
		t.Fatal("key not found")
	}
	if v != "updated" {
		t.Errorf("want 'updated', got %q", v)
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

func TestOpenStore__ErrorsWhenPathUnreadable(t *testing.T){
    t.Parallel()
    path := t.TempDir() + "/unreadable.store"
    if _, err := os.Create(path); err != nil {
        t.Fatal(err)
    }
    if err := os.Chmod(path, 0o000); err != nil {
        t.Fatal(err)
    }
    _, err := kv.OpenStore(path)
    if err == nil {
        t.Fatal("want an error but got none")
    }
}


func TestOpenStore__ReturnsErrorOnInvalidData(t *testing.T){
    t.Parallel()
    _, err := kv.OpenStore("testdata/invalid.store")
    if err == nil {
        t.Fatal("no error")
    }
}

func TestSave__ErrorsWhenPathUnwriteable(t *testing.T){
    t.Parallel()
    s, err := kv.OpenStore("bogus/unwritable.store")
    if err != nil {
        t.Fatal(err)
    }
    err = s.Save()
    if err == nil {
        t.Fatal("want error and got none")
    }
}