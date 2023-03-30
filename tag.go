package main

// The tag tree is used to look up music clips
// using fuzzy matching (so that the name does not
// have to be typed 100% correct on the command line)
// A tag may represent, e.g., an artist, album, clip...
// Some tags, like albums, have child Tags while
// others, like clips, are leaf nodes.
// The leaf nodes point to a file to be played.

import (
	"strings"
)

type Tag struct {
	fuzzy    string // fuzzyfied tag
	children []*Tag // children, if any
	file     string // music file, in case of leaf node
}

func NewTag(tag string) *Tag {
	return &Tag{Fuzzy(tag), []*Tag{}, ""}
}

// Get a child by fuzzy tag matching.
// If the child does not exist yet, it is added.
func (this *Tag) Ch