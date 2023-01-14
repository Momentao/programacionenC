
package main

// This file implements the Library data structure.

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// Stores a music Library
type Lib struct {
	clips   []*Clip // all music files
	artists *Tag    // root node for all artists
	tags    []*Tag
}

// Constructs a new Library
func NewLib() *Lib {
	lib := new(Lib)
	lib.init()
	return lib
}

// Initializes the library
func (lib *Lib) init() {
	lib.clips = []*Clip{}
	lib.artists = NewTag("")
	lib.tags = []*Tag{}
}

// Recursively import directory or file into library.
func (lib *Lib) Import(arg string) {
	// rm trailing slash
	if strings.HasSuffix(arg, "/") {
		arg = arg[:len(arg)-1]
	}

	info, err := os.Stat(arg)
	Check(err) // TODO: dontcrash

	if info.IsDir() {
		dir, err := os.OpenFile(arg, os.O_RDONLY, 0777)
		Check(err)
		files, err2 := dir.Readdirnames(-1)
		Check(err2)
		for _, f := range files {
			lib.Import(arg + "/" + f)
		}
		return
	}

	if !info.IsDir() {
		lib.ImportFile(arg)
		return
	}
}

// Non-recursively import file into directory
func (lib *Lib) ImportFile(file string) {
	clip := NewClip(file)
	lib.clips = append(lib.clips, clip)
	artist, ok1 := lib.artists.Child(clip.Artist())
	if !ok1 {
		lib.AddTag(artist)
	}
	album, ok2 := artist.Child(clip.Album())
	if !ok2 {
		lib.AddTag(album)
	}
	title, ok3 := album.Child(clip.Title())
	if !ok3 {
		lib.AddTag(title)
	}
	title.file = file