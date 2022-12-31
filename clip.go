
package main

import (
	"path"
	"strings"
	"unicode"
)

// Represents a music clip.
type Clip struct {
	file string
	tags [5]string
}

// Index for Clip.tags
const (
	TAG_TRACK = iota
	TAG_TITLE
	TAG_ALBUM
	TAG_ARTIST
	TAG_GENRE
)

var tagStr []string = []string{"Track", "Title", "Album", "Artist", "Genre"}

func NewClip(file string) *Clip {
	clip := new(Clip)
	clip.file = file
	clip.initTags()
	return clip
}

// Rudimentary way to set clip tags based on file name:
//	artist/album/01_title.ogg