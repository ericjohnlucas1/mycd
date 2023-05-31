package main

import (
	"fmt"
	"regexp"
	"strings"
)

type path struct {
	absolutePath bool
	parts        []string
}

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

func (p path) isAbsolute() bool {
	return p.absolutePath
}

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

func (p path) getStringRepresentation() string {
	rep := strings.Join(p.parts, "/")
	if p.isAbsolute() {
		rep = "/" + rep
	}

	return rep
}

func (p path) concatenate(p2 *path) *path {
	return &path{
		absolutePath: p.absolutePath,
		parts:        append(p.parts, p2.parts...),
	}
}
