package main

import (
	"testing"
)

func testdriver(path1 string, path2 string) (string, error) {
	destpath, err := NewPath(path2)
	if err != nil {
		return "_", err
	}
	curpath, err := NewPath(path1)
	if err != nil {
		return "_", err
	}
	if destpath.isAbsolute() {
		return destpath.getStringRepresentation(), nil

	} else {
		resultpath := curpath.concatenate(destpath)
		resultpath.shorten()
		return resultpath.getStringRepresentation(), nil
	}
}
func expectsuccess(t *testing.T, path1 string, path2 string, expected string) {
	res, err := testdriver(path1, path2)
	if err != nil {
		t.Fatalf("got unexpected error %v", err)
	}
	if res != expected {
		t.Fatalf("expected %v, got %v", expected, res)
	}
}

func expecterror(t *testing.T, path1 string, path2 string) {
	_, err := testdriver(path1, path2)
	if err == nil {
		t.Fatalf("Expecting error")
	}
}

func TestRelativePath1(t *testing.T) {
	expectsuccess(t, "/", "abc", "/abc")
}

func TestRelativePath2(t *testing.T) {
	expectsuccess(t, "/abc/def", "ghi", "/abc/def/ghi")
}
func TestPreviousPath(t *testing.T) {
	expectsuccess(t, "/abc/def", "..", "/abc")
}
func TestAbsolutePath1(t *testing.T) {
	expectsuccess(t, "/abc/def", "/abc", "/abc")
}
func TestAbsolutePath2(t *testing.T) {
	expectsuccess(t, "/abc/def", "/abc/klm", "/abc/klm")
}

func TestMultipleSlashes(t *testing.T) {
	expectsuccess(t, "/abc/def", "//////", "/")
}
func TestChallenge(t *testing.T) {
	expectsuccess(t, "/abc/def", "../gh///../klm/.", "/abc/klm")
}

func TestError(t *testing.T) {
	expecterror(t, "/abc/def", "..klm")
}
