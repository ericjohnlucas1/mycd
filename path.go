package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Struct used to represent a path
type path struct {
	absolutePath bool
	parts        []string
}

// Utility method to remove all occurences of an element from a slice.
func removeElement(s []string, val string) []string {
	j := 0
	for _, v := range s {
		if v != val {
			s[j] = v
			j++
		}
	}
	return s[:j]
}

// Returns a new path object given a path string. An error will be returned if the path contains an invalid file/directory name.
func NewPath(pathstring string) (*path, error) {
	var absolutePath bool
	if pathstring[0] == '/' {
		absolutePath = true
	}
	parts := removeElement(strings.Split(pathstring, "/"), "")
	for _, p := range parts {
		if p != ".." && p != "." && !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(p) {
			return nil, fmt.Errorf("%v: No such file or directory", p)
		}
	}

	path := &path{
		absolutePath: absolutePath,
		parts:        parts,
	}
	return path, nil
}

// Returns a boolean value of True if the path is an absolute path, and a boolean value of False if the path is relative
func (p path) isAbsolute() bool {
	return p.absolutePath
}

// Removes segments of the path containing /../ and /./, and resolves such intelligently. Intended for use with absolute paths.
func (p *path) shorten() {
	var shortened []string
	for _, p := range p.parts {
		if p == "." {
			continue
		}
		if p == ".." {
			if len(shortened) > 0 {
				shortened = shortened[:len(shortened)-1]
			}
			continue
		}
		shortened = append(shortened, p)
	}
	p.parts = shortened
}

// Returns the string representation of the path. Absolute paths begin with '/', whereas relative paths do not.
func (p path) getStringRepresentation() string {
	rep := strings.Join(p.parts, "/")
	if p.isAbsolute() {
		rep = "/" + rep
	}

	return rep
}

// Appends the path passed in to the end of the path in context, and returns as a new path struct
func (p path) concatenate(p2 *path) *path {
	return &path{
		absolutePath: p.absolutePath,
		parts:        append(p.parts, p2.parts...),
	}
}
